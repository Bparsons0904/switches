package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"log"
	"sync"
	"time"
)

type Auth struct {
	State         string    `json:"state"`
	Nonce         string    `json:"nonce"`
	CodeChallenge string    `json:"codeChallenge"`
	ExpiresAt     time.Time `json:"expiresAt"`
}

type AuthResponse struct {
	Code  string `json:"code"`
	State string `json:"state"`
}

var (
	activeAuths = make(map[string]Auth)
	authMutex   sync.RWMutex
)

func generateRandomString(length int) (string, error) {
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

func GenerateAuth() (Auth, error) {
	state, err := generateRandomString(32)
	if err != nil {
		log.Println("Error generating state", err)
		return Auth{}, err
	}

	nonce, err := generateRandomString(32)
	if err != nil {
		log.Println("Error generating nonce", err)
		return Auth{}, err
	}

	codeVerifier, err := generateRandomString(32)
	if err != nil {
		log.Println("Error generating code verifier", err)
		return Auth{}, err
	}

	codeChallenge := generateCodeChallenge(codeVerifier)

	expiresAt := time.Now().Add(15 * time.Minute)

	auth := Auth{
		State:         state,
		Nonce:         nonce,
		CodeChallenge: codeChallenge,
		ExpiresAt:     expiresAt,
	}

	activeAuths[state] = auth

	return auth, nil
}

func GetAuth(state string) (Auth, bool) {
	authMutex.Lock()
	defer authMutex.Unlock()

	auth, ok := activeAuths[state]
	if !ok {
		log.Println("Auth not found for state", state)
		return Auth{}, false
	}

	if time.Now().After(auth.ExpiresAt) {
		log.Println("Auth expired for state", state)
		delete(activeAuths, state)
		return Auth{}, false
	}

	delete(activeAuths, state)
	return auth, ok
}

func CleanUpAuths() {
	authMutex.Lock()
	defer authMutex.Unlock()

	newActiveAuths := make(map[string]Auth, len(activeAuths))
	now := time.Now()
	for state, auth := range activeAuths {
		if now.Before(auth.ExpiresAt) {
			newActiveAuths[state] = auth
		}
	}

	activeAuths = newActiveAuths
}
