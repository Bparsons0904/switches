package services

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"switches/database"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func GenerateRandomString(length int) (string, error) {
	b := make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func generateCodeChallenge(verifier string) string {
	h := sha256.New()
	h.Write([]byte(verifier))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}

func GenerateAuthString() (Auth, error) {
	state, err := GenerateRandomString(32)
	if err != nil {
		log.Println("Error generating state", err)
		return Auth{}, err
	}

	nonce, err := GenerateRandomString(32)
	if err != nil {
		log.Println("Error generating nonce", err)
		return Auth{}, err
	}

	codeVerifier, err := GenerateRandomString(32)
	if err != nil {
		log.Println("Error generating code verifier", err)
		return Auth{}, err
	}

	codeChallenge := generateCodeChallenge(codeVerifier)

	expiresAt := time.Now().Add(15 * time.Minute)
	auth := Auth{
		ClientID:      viper.GetString("AUTH_CLIENT_ID"),
		AuthURL:       viper.GetString("AUTH_URL"),
		CallbackURL:   viper.GetString("AUTH_REDIRECT_URL"),
		State:         state,
		Nonce:         nonce,
		CodeChallenge: codeChallenge,
		CodeVerifier:  codeVerifier,
		ExpiresAt:     expiresAt,
	}

	authJSON, err := json.Marshal(auth)
	if err != nil {
		log.Println("Error marshalling auth:", err)
		return Auth{}, err
	}

	if err := database.KeyDB.Set(database.KeyDB.Context(), state, authJSON, 15*time.Minute).Err(); err != nil {
		log.Println("Error saving auth", err)
		return Auth{}, err
	}

	return auth, nil
}

func GetAuth(state string) (Auth, error) {
	keyDB := database.KeyDB

	authJSON, err := keyDB.Get(keyDB.Context(), state).Result()
	if err == redis.Nil {
		log.Println("Auth not found for state:", state)
		return Auth{}, err
	} else if err != nil {
		log.Println("Error retrieving auth from Redis:", err)
		return Auth{}, err
	}

	var auth Auth
	err = json.Unmarshal([]byte(authJSON), &auth)
	if err != nil {
		log.Println("Error unmarshalling auth:", err)
		return Auth{}, err
	}

	if time.Now().After(auth.ExpiresAt) {
		log.Println("Auth expired for state:", state)
		keyDB.Del(keyDB.Context(), state)
		return Auth{}, fmt.Errorf("Auth expired for state: %s", state)
	}

	keyDB.Del(keyDB.Context(), state)
	return auth, nil
}

func SessionFlow(sessionID string) (Session, error) {
	session, err := database.GetJSONKeyDB[Session](
		"session",
		sessionID,
	)
	if err != nil {
		log.Println("Error getting session from keydb", err)
		return Session{}, err
	}

	log.Println("Session found", session.ExpiresAt)

	if !session.IsLoggedIn {
		log.Println("Session not logged in", session)
		err = fmt.Errorf("Session not logged in")
	}

	if time.Now().After(session.ExpiresAt) {
		log.Println("Session expired", session)
		if time.Now().After(session.RefreshBy) {
			err = fmt.Errorf("Session expired")
		} else {
			go refreshToken(session)
		}
	}

	if err != nil {
		err = database.DeleteUUIDKeyDB("session", session.SessionID)
		return Session{}, fmt.Errorf("Problem with the session %f", err)
	}

	log.Println("Session flow completed", session)
	return session, nil
}

func refreshToken(session Session) {
	log.Println("Refreshing token for session", session.SessionID)

	clientID := viper.GetString("AUTH_CLIENT_ID")
	authURL := viper.GetString("AUTH_URL")

	tokenURL := fmt.Sprintf("https://%s/oauth/v2/token", authURL)
	resp, err := http.PostForm(tokenURL, url.Values{
		"grant_type":    {"refresh_token"},
		"client_id":     {clientID},
		"refresh_token": {session.RefreshToken}, // Use the refresh token here
	})
	if err != nil {
		log.Println("Error getting token", err)
		EndUserSession(session)
		return
	}
	defer resp.Body.Close()

	var tokenResponse TokenResponse
	if err = json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		log.Println("Error decoding token response", err)
		EndUserSession(session)
		return
	}

	session.AccessToken = tokenResponse.AccessToken
	session.ExpiresAt = time.Now().Add(time.Duration(tokenResponse.ExpiresIn) * time.Second)

	if tokenResponse.RefreshToken != session.RefreshToken {
		session.RefreshToken = tokenResponse.RefreshToken
		session.RefreshBy = time.Now().Add(30 * 24 * time.Hour)
	}

	err = database.SetJSONKeyDB("session", session.SessionID.String(), session)
	if err != nil {
		EndUserSession(session)
		log.Println("Error setting session in keydb", err)
	}
}

func EndUserSession(session Session) {
	session.IsLoggedIn = false
	err := database.SetUUIDJSONKeyDB("session", session.SessionID, session)
	if err != nil {
		log.Println("Error setting session in keydb", err)
	}
}

type Auth struct {
	ClientID      string    `json:"clientId"`
	AuthURL       string    `json:"authUrl"`
	CallbackURL   string    `json:"callbackUrl"`
	State         string    `json:"state"`
	Nonce         string    `json:"nonce"`
	CodeChallenge string    `json:"codeChallenge"`
	CodeVerifier  string    `json:"codeVerifier"`
	ExpiresAt     time.Time `json:"expiresAt"`
}

type TokenResponse struct {
	IDToken      string `json:"id_token"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

type AuthResponse struct {
	Code  string `json:"code"`
	State string `json:"state"`
}

type Session struct {
	SessionID    uuid.UUID `json:"sessionId"`
	UserID       uuid.UUID `json:"userId"`
	Sub          int       `json:"sub"`
	AccessToken  string    `json:"accessToken"`
	ExpiresAt    time.Time `json:"expiresAt"`
	RefreshToken string    `json:"refreshToken"`
	RefreshBy    time.Time `json:"refreshBy"`
	IDToken      string    `json:"idToken"`
	IsLoggedIn   bool      `json:"isLoggedIn"`
}

type Claims struct {
	Sub               int      `json:"sub"`
	Email             string   `json:"email"`
	EmailVerified     bool     `json:"email_verified"`
	FamilyName        string   `json:"family_name"`
	Name              string   `json:"name"`
	PreferredUsername string   `json:"preferred_username"`
	GivenName         string   `json:"given_name"`
	Roles             []string `json:"roles"`
}
