package database

import (
	"encoding/json"
	"fmt"
	"switches/config"
	"switches/utils"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	KeyDB *redis.Client
)

func ConnectDB(config config.Config, server *fiber.App) (*gorm.DB, *redis.Client) {
	DB, err := StartPostgresDB(config)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to the PostgreSQL Database")
	}

	KeyDB, err := StartKeyDB(config)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to the KeyDB Database")
	}

	return DB, KeyDB
}

func StartPostgresDB(config config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)

	zlog := utils.GetLogger()
	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{Logger: utils.NewGormLogger(zlog), PrepareStmt: true},
	)
	if err != nil {
		return nil, err
	}

	log.Info().Msg("Connected Successfully to the PostgreSQL Database")

	DB = db
	return db, nil
}

func StartKeyDB(config config.Config) (*redis.Client, error) {
	if KeyDB != nil {
		log.Warn().Msg("KeyDB already initialized. Reusing existing connection.")
		return KeyDB, nil
	}

	keydb := redis.NewClient(&redis.Options{
		Addr:     config.KeyDBHost,
		Password: config.KeyDBPassword,
		DB:       int(config.KeyDB),
	})

	_, err := keydb.Ping(keydb.Context()).Result()
	if err != nil {
		message := fmt.Sprintf("Failed to ping KeyDB at %s: %v", config.KeyDBHost, err)
		log.Error().Err(err).Msg(message)
		return nil, err
	}

	KeyDB = keydb
	log.Info().Msg("Connected Successfully to the KeyDB Database")

	return KeyDB, nil
}

func SetUUIDJSONKeyDB[T any](hash string, key uuid.UUID, value T, duration time.Duration) error {
	err := SetJSONKeyDB(hash, key.String(), value, duration)
	return err
}

func SetJSONKeyDB[T any](hash, key string, value T, duration time.Duration) error {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		log.Error().Err(err).Msg("Error marshalling JSON for key/value store")
		return err
	}

	query := fmt.Sprintf("%s:%s", hash, key)
	if err := KeyDB.Set(KeyDB.Context(), query, jsonValue, duration).Err(); err != nil {
		log.Error().Err(err).Msg("Error setting key/value store for JSON value")
		return err
	}

	return nil
}

func DeleteUUIDKeyDB(hash string, key uuid.UUID) error {
	err := DeleteStringKeyDB(hash, key.String())
	return err
}

func DeleteStringKeyDB(hash, key string) error {
	if hash == "" || key == "" {
		log.Warn().Msg("Key is empty, cannot delete key/value store")
		return fmt.Errorf("Key is empty, cannot delete key/value store")
	}

	query := fmt.Sprintf("%s:%s", hash, key)
	err := KeyDB.Del(KeyDB.Context(), query).Err()

	return err
}

func GetUUIDJSONKeyDB[T any](hash string, key uuid.UUID) (T, error) {
	stringKey := key.String()
	value, err := GetJSONKeyDB[T](hash, stringKey)
	if err != nil {
		return value, err
	}
	return value, nil
}

func GetJSONKeyDB[T any](hash, key string) (T, error) {
	var value T
	if hash == "" || key == "" {
		log.Warn().Msg("Key is empty, cannot get JDON value from key/value store,")
		return value, fmt.Errorf("Key is empty, cannot get JDON value from key/value store")
	}

	query := fmt.Sprintf("%s:%s", hash, key)
	jsonValue, err := KeyDB.Get(KeyDB.Context(), query).Result()
	if err != nil {
		log.Error().Err(err).Msg("Error getting key/value store for JSON value")
		return value, err
	}

	err = json.Unmarshal([]byte(jsonValue), &value)
	if err != nil {
		log.Error().Err(err).Msg("Error unmarshalling JSON value")
		return value, err
	}
	return value, nil
}
