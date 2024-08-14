package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	os.Setenv("TIER", "test")
	os.Setenv("BASE_URL", "http://localhost:8080")
	os.Setenv("SERVER_PORT", "8080")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_USER", "testuser")
	os.Setenv("DB_PASSWORD", "testpassword")
	os.Setenv("DB_NAME", "testdb")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("AUTH_CLIENT_ID", "testclientid")
	os.Setenv("AUTH_URL", "http://auth.url")
	os.Setenv("AUTH_REDIRECT_URL", "http://auth.redirect")

	config, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig returned an error: %v", err)
	}

	if config.Tier != "test" {
		t.Errorf("Expected Tier to be 'test', got '%s'", config.Tier)
	}

	if config.BaseURL != "http://localhost:8080" {
		t.Errorf("Expected BaseURL to be 'http://localhost:8080', got '%s'", config.BaseURL)
	}

	if config.ServerPort != "8080" {
		t.Errorf("Expected ServerPort to be '8080', got '%s'", config.ServerPort)
	}

	if config.DBHost != "localhost" {
		t.Errorf("Expected DBHost to be 'localhost', got '%s'", config.DBHost)
	}

	if config.DBUser != "testuser" {
		t.Errorf("Expected DBUser to be 'testuser', got '%s'", config.DBUser)
	}

	if config.DBPassword != "testpassword" {
		t.Errorf("Expected DBPassword to be 'testpassword', got '%s'", config.DBPassword)
	}

	if config.DBName != "testdb" {
		t.Errorf("Expected DBName to be 'testdb', got '%s'", config.DBName)
	}

	if config.DBPort != "5432" {
		t.Errorf("Expected DBPort to be '5432', got '%s'", config.DBPort)
	}

	if config.AuthClientID != "testclientid" {
		t.Errorf("Expected AuthClientID to be 'testclientid', got '%s'", config.AuthClientID)
	}

	if config.AuthURL != "http://auth.url" {
		t.Errorf("Expected AuthURL to be 'http://auth.url', got '%s'", config.AuthURL)
	}

	if config.AuthRedirectURL != "http://auth.redirect" {
		t.Errorf(
			"Expected AuthRedirectURL to be 'http://auth.redirect', got '%s'",
			config.AuthRedirectURL,
		)
	}

	if config.AppendNumber == 0 {
		t.Error("Expected AppendNumber to be non-zero, but it was 0")
	}

	os.Unsetenv("TIER")
	os.Unsetenv("BASE_URL")
	os.Unsetenv("SERVER_PORT")
	os.Unsetenv("DB_HOST")
	os.Unsetenv("DB_USER")
	os.Unsetenv("DB_PASSWORD")
	os.Unsetenv("DB_NAME")
	os.Unsetenv("DB_PORT")
	os.Unsetenv("AUTH_CLIENT_ID")
	os.Unsetenv("AUTH_URL")
	os.Unsetenv("AUTH_REDIRECT_URL")
}

func TestValidateConfig_MissingString(t *testing.T) {
	config := Config{
		Tier:            "",
		BaseURL:         "http://localhost:8080",
		ServerPort:      "8080",
		DBHost:          "localhost",
		DBUser:          "testuser",
		DBPassword:      "testpassword",
		DBName:          "testdb",
		DBPort:          "5432",
		AuthClientID:    "testclientid",
		AuthURL:         "http://auth.url",
		AuthRedirectURL: "http://auth.redirect",
		AppendNumber:    123456789,
	}

	err := validateConfig(config)
	if err == nil {
		t.Error("Expected validation to fail due to missing Tier string, but it succeeded")
	}
}

func TestValidateConfig_MissingInt64(t *testing.T) {
	config := Config{
		Tier:            "development",
		BaseURL:         "http://localhost:8080",
		ServerPort:      "8080",
		DBHost:          "localhost",
		DBUser:          "testuser",
		DBPassword:      "testpassword",
		DBName:          "testdb",
		DBPort:          "5432",
		AuthClientID:    "testclientid",
		AuthURL:         "http://auth.url",
		AuthRedirectURL: "http://auth.redirect",
		AppendNumber:    0, // Intentionally left to trigger validation error
	}

	err := validateConfig(config)
	if err == nil {
		t.Error("Expected validation to fail due to missing AppendNumber, but it succeeded")
	}
}
