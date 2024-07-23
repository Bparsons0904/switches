package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Tier       string
	BaseURL    string
	ServerPort string
}

func LoadConfig() (Config, error) {
	// Set the file name of the configurations file
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	// Read configuration
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}

	// Set environment variables (if they are not set already)
	viper.AutomaticEnv()

	config := Config{
		Tier:       viper.GetString("TIER"),
		BaseURL:    viper.GetString("BASE_URL"),
		ServerPort: viper.GetString("SERVER_PORT"),
	}

	return config, nil
}
