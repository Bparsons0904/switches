package config

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Tier            string
	BaseURL         string
	ServerPort      string
	DBHost          string
	DBUser          string
	DBPassword      string
	DBName          string
	DBPort          string
	AuthClientID    string
	AuthURL         string
	AuthRedirectURL string
	KeyDBHost       string
	KeyDBPassword   string
	KeyDB           int64
	AppendNumber    int64
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
	log.Println("Current time in seconds", appendNumber)

	config := Config{
		Tier:            viper.GetString("TIER"),
		BaseURL:         viper.GetString("BASE_URL"),
		ServerPort:      viper.GetString("SERVER_PORT"),
		DBHost:          viper.GetString("DB_HOST"),
		DBUser:          viper.GetString("DB_USER"),
		DBPassword:      viper.GetString("DB_PASSWORD"),
		DBName:          viper.GetString("DB_NAME"),
		DBPort:          viper.GetString("DB_PORT"),
		AuthClientID:    viper.GetString("AUTH_CLIENT_ID"),
		AuthURL:         viper.GetString("AUTH_URL"),
		AuthRedirectURL: viper.GetString("AUTH_REDIRECT_URL"),
		KeyDBHost:       viper.GetString("KEYDB_HOST"),
		KeyDBPassword:   viper.GetString("KEYDB_PASSWORD"),
		KeyDB:           viper.GetInt64("KEYDB_DB"),
		AppendNumber:    appendNumber,
	}

	err := validateConfig(config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func validateConfig(config Config) error {
	v := reflect.ValueOf(config)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := v.Type().Field(i).Name

		switch field.Kind() {
		case reflect.String:
			if field.String() == "" {
				return fmt.Errorf("config string field %s is empty", fieldName)
			}
		case reflect.Int64:
			if field.Int() == 0 {
				return fmt.Errorf("config int64 field %s is not set", fieldName)
			}
		case reflect.Bool:
		case reflect.Float64:
			if field.Float() == 0.0 {
				return fmt.Errorf("config bool field %s is not set", fieldName)
			}
		default:
			return fmt.Errorf("not field type %s is not set", fieldName)
		}
	}
	return nil
}
