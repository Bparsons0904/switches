package config

import (
	"log"

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

// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)
func LoadConfig() (Config, error) {
	// Set the file name of the configurations file
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	// Read configuration
	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("Error reading config file, %s", err)
	}

	// Set environment variables (if they are not set already)
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
