package services

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"switches/database"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

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

type AuthResponse struct {
	Code  string `json:"code"`
	State string `json:"state"`
}

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
