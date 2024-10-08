package services

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"switches/database"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
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

func GenerateAuthString(originalURL string) (Auth, error) {
	state, err := GenerateRandomString(32)
	if err != nil {
		log.Error().Err(err).Msg("Error generating state")
		return Auth{}, err
	}

	nonce, err := GenerateRandomString(32)
	if err != nil {
		log.Error().Err(err).Msg("Error generating nonce")
		return Auth{}, err
	}

	codeVerifier, err := GenerateRandomString(32)
	if err != nil {
		log.Error().Err(err).Msg("Error generating code verifier")
		return Auth{}, err
	}

	codeChallenge := generateCodeChallenge(codeVerifier)

	expiresAt := time.Now().Add(15 * time.Minute)
	callbackUrl := fmt.Sprintf("%s/auth/callback", viper.GetString("BASE_URL"))
	auth := Auth{
		ClientID:      viper.GetString("AUTH_CLIENT_ID"),
		AuthURL:       viper.GetString("AUTH_URL"),
		CallbackURL:   callbackUrl,
		State:         state,
		Nonce:         nonce,
		CodeChallenge: codeChallenge,
		CodeVerifier:  codeVerifier,
		ExpiresAt:     expiresAt,
		OriginalURL:   originalURL,
	}

	if err := database.SetJSONKeyDB("auth", state, auth, 15*time.Minute); err != nil {
		log.Error().Err(err).Msg("Error saving auth")
		return Auth{}, err
	}

	return auth, nil
}

func GetAuth(state string) (Auth, error) {
	auth, err := database.GetJSONKeyDB[Auth]("auth", state)
	if err != nil {
		log.Error().Err(err).Msg("Error retrieving auth from Redis")
		return Auth{}, err
	}

	err = database.DeleteStringKeyDB("auth", state)
	if err != nil {
		log.Error().Err(err).Msg("Error deleting auth from Redis")
	}

	if time.Now().After(auth.ExpiresAt) {
		log.Error().Err(err).Any("state", state).Msg("Auth expired for state")
		return Auth{}, fmt.Errorf("Auth expired for state: %s", state)
	}

	return auth, nil
}

func SessionFlow(sessionID string) (Session, error) {
	session, err := database.GetJSONKeyDB[Session](
		"session",
		sessionID,
	)
	if err != nil {
		log.Error().Err(err).Msg("Error getting session from keydb")
		return Session{}, err
	}

	if !session.IsLoggedIn {
		log.Warn().Any("session", session).Msg("Session not logged in")
		err = fmt.Errorf("Session not logged in")
	}

	if time.Now().After(session.ExpiresAt) {
		log.Warn().Any("session", session).Msg("Session expired")
		err = refreshToken(session)
	}

	if err != nil {
		_ = database.DeleteStringKeyDB("session", session.AccessToken)
		log.Error().Err(err).Msg("Error refreshing token")
		return Session{}, fmt.Errorf("Problem with the session %f", err)
	}

	return session, nil
}

func refreshToken(session Session) error {
	clientID := viper.GetString("AUTH_CLIENT_ID")
	authURL := viper.GetString("AUTH_URL")
	tokenURL := fmt.Sprintf("https://%s/oauth/v2/token", authURL)
	resp, err := http.PostForm(tokenURL, url.Values{
		"grant_type":    {"refresh_token"},
		"client_id":     {clientID},
		"refresh_token": {session.RefreshToken}, // Use the refresh token here
	})
	if err != nil {
		log.Error().Err(err).Msg("Error getting token")
		return err
	}
	defer resp.Body.Close()

	var tokenResponse TokenResponse
	if err = json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		log.Error().Err(err).Msg("Error decoding token response")
		return err
	}

	session.AccessToken = tokenResponse.AccessToken
	session.ExpiresAt = time.Now().Add(time.Duration(tokenResponse.ExpiresIn) * time.Second)

	if tokenResponse.RefreshToken != session.RefreshToken {
		session.RefreshToken = tokenResponse.RefreshToken
		session.RefreshBy = time.Now().Add(30 * 24 * time.Hour)
	}

	err = database.SetJSONKeyDB("session", session.AccessToken, session, 30*24*time.Hour)
	if err != nil {
		log.Error().Err(err).Msg("Error setting session in keydb")
		return err
	}

	return nil
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
	OriginalURL   string    `json:"originalUrl"`
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
	AccessToken  string    `json:"accessToken"`
	UserID       uuid.UUID `json:"userId"`
	Sub          int       `json:"sub"`
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
