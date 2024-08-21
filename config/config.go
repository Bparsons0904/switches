package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Tier          string
	BaseURL       string
	ServerPort    string
	DBHost        string
	DBUser        string
	DBPassword    string
	DBName        string
	DBPort        string
	AuthClientID  string
	AuthURL       string
	KeyDBHost     string
	KeyDBPassword string
	KeyDB         int64
	AppendNumber  int64
	RandonString  string
}

func LoadConfig() (Config, error) {
	env := os.Getenv("TIER")
	if env == "" {
		env = "development"
	}

	if env == "development" {
		viper.SetConfigName(".env")
		viper.SetConfigType("env")
		viper.AddConfigPath(".")

		if err := viper.ReadInConfig(); err != nil {
			log.Panicf("Error reading config file, %s", err)
		}
	}

	viper.AutomaticEnv()

	appendNumber := time.Now().Unix()
	viper.Set("appendNumber", appendNumber)
	requiredEnvs := []string{
		"TIER",
		"BASE_URL",
		"SERVER_PORT",
		"DB_HOST",
		"DB_USER",
		"DB_PASSWORD",
		"DB_NAME",
		"DB_PORT",
		"AUTH_CLIENT_ID",
		"AUTH_URL",
		"KEYDB_HOST",
		"KEYDB_PASSWORD",
	}

	if err := validateEnvVars(requiredEnvs); err != nil {
		return Config{}, err
	}

	config := Config{
		Tier:          viper.GetString("TIER"),
		BaseURL:       viper.GetString("BASE_URL"),
		ServerPort:    viper.GetString("SERVER_PORT"),
		DBHost:        viper.GetString("DB_HOST"),
		DBUser:        viper.GetString("DB_USER"),
		DBPassword:    viper.GetString("DB_PASSWORD"),
		DBName:        viper.GetString("DB_NAME"),
		DBPort:        viper.GetString("DB_PORT"),
		AuthClientID:  viper.GetString("AUTH_CLIENT_ID"),
		AuthURL:       viper.GetString("AUTH_URL"),
		KeyDBHost:     viper.GetString("KEYDB_HOST"),
		KeyDBPassword: viper.GetString("KEYDB_PASSWORD"),
		KeyDB:         viper.GetInt64("KEYDB_DB"),
		RandonString:  "7pdme2QREQqq15JZ",
		AppendNumber:  appendNumber,
	}

	return config, nil
}

func validateEnvVars(envVars []string) error {
	for _, envVar := range envVars {
		if viper.GetString(envVar) == "" {
			return fmt.Errorf("required environment variable %s is not set or empty", envVar)
		}
	}
	return nil
}
