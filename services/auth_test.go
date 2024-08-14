package services

import (
	"encoding/base64"
	"testing"
	"time"
)

func TestGenerateAuth(t *testing.T) {
	auth, err := GenerateAuthString()
	if err != nil {
		t.Fatalf("GenerateAuth returned an error: %v", err)
	}

	if auth.State == "" || auth.Nonce == "" || auth.CodeChallenge == "" {
		t.Fatalf("Generated auth has empty fields: %v", auth)
	}

	if auth.ExpiresAt.Before(time.Now()) {
		t.Fatalf("Generated auth has an expiration time in the past: %v", auth.ExpiresAt)
	}
}

func TestGetAuth(t *testing.T) {
	auth, err := GenerateAuthString()
	if err != nil {
		t.Fatalf("GenerateAuth returned an error: %v", err)
	}

	retrievedAuth, ok := GetAuth(auth.State)
	if !ok {
		t.Fatalf("GetAuth did not retrieve the auth: %v", auth.State)
	}

	if retrievedAuth.State != auth.State || retrievedAuth.Nonce != auth.Nonce ||
		retrievedAuth.CodeChallenge != auth.CodeChallenge {
		t.Fatalf("Retrieved auth does not match the generated one")
	}

	_, ok = GetAuth(auth.State)
	if ok {
		t.Fatalf("GetAuth should not retrieve an auth after it has been used")
	}
}

func TestCleanUpAuths(t *testing.T) {
	auth, err := GenerateAuthString()
	if err != nil {
		t.Fatalf("GenerateAuth returned an error: %v", err)
	}

	// Set the auth to expire in the past
	auth.ExpiresAt = time.Now().Add(-1 * time.Minute)
	activeAuths[auth.State] = auth

	CleanUpAuths()

	_, ok := activeAuths[auth.State]
	if ok {
		t.Fatalf("CleanUpAuths did not remove the expired auth")
	}
}

func TestGenerateRandomString(t *testing.T) {
	str, err := generateRandomString(32)
	if err != nil {
		t.Fatalf("generateRandomString returned an error: %v", err)
	}

	if len(str) == 0 {
		t.Fatalf("generateRandomString returned an empty string")
	}
}

func TestGenerateRandomStringLength(t *testing.T) {
	str, err := generateRandomString(32)
	if err != nil {
		t.Fatalf("generateRandomString returned an error: %v", err)
	}

	decoded, err := base64.RawURLEncoding.DecodeString(str)
	if err != nil {
		t.Fatalf("generateRandomString returned an invalid base64 string: %v", err)
	}

	if len(decoded) != 32 {
		t.Fatalf("generateRandomString did not return a string of the correct length")
	}
}

func TestGenerateCodeChallenge(t *testing.T) {
	verifier := "test_verifier"
	challenge := generateCodeChallenge(verifier)

	if challenge == "" {
		t.Fatalf("generateCodeChallenge returned an empty challenge")
	}
}
