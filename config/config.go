package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Tier       string
	BaseURL    string
	ServerPort string
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
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

	config := Config{
		Tier:       viper.GetString("TIER"),
		BaseURL:    viper.GetString("BASE_URL"),
		ServerPort: viper.GetString("SERVER_PORT"),
		DBHost:     viper.GetString("DB_HOST"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName:     viper.GetString("DB_NAME"),
		DBPort:     viper.GetString("DB_PORT"),
	}

	return config, nil
}
