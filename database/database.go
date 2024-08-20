package database

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"switches/config"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB    *gorm.DB
	KeyDB *redis.Client
)

func ConnectDB(config config.Config, server *fiber.App) (*gorm.DB, *redis.Client) {
	DB, err := StartPostgresDB(config)
	if err != nil {
		log.Fatal("Failed to connect to the PostgreSQL Database")
	}

	KeyDB, err := StartKeyDB(config)
	if err != nil {
		log.Fatal("Failed to connect to the KeyDB Database")
	}

	// server.Use(DBMiddleware(DB, KeyDB))
	return DB, KeyDB
}

// func DBMiddleware(db *gorm.DB, keydb *redis.Client) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		c.Locals("db", db)
// 		// c.Locals("keydb", keydb)
// 		return c.Next()
// 	}
// }

func StartPostgresDB(config config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)

	var logLevel logger.LogLevel
	if viper.Get("DB_LOGGING") == "true" {
		logLevel = logger.Info
	} else {
		logLevel = logger.Warn
	}

	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{Logger: logger.Default.LogMode(logLevel), PrepareStmt: true},
	)
	if err != nil {
		return nil, err
	}

	log.Println("Connected Successfully to the PostgreSQL Database")

	DB = db
	return db, nil
}

func StartKeyDB(config config.Config) (*redis.Client, error) {
	if KeyDB != nil {
		log.Println("KeyDB already initialized. Reusing existing connection.")
		return KeyDB, nil
	}

	keydb := redis.NewClient(&redis.Options{
		Addr:     config.KeyDBHost,
		Password: config.KeyDBPassword,
		DB:       int(config.KeyDB),
	})

	_, err := keydb.Ping(keydb.Context()).Result()
	if err != nil {
		return nil, err
	}

	KeyDB = keydb
	log.Println("Connected Successfully to the KeyDB Database")

	return KeyDB, nil
}

func SetUUIDJSONKeyDB[T any](hash string, key uuid.UUID, value T) error {
	err := SetJSONKeyDB(hash, key.String(), value)
	return err
}

func SetJSONKeyDB[T any](hash, key string, value T) error {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		log.Println("Error marshalling JSON for key/value store", err)
		return err
	}

	query := fmt.Sprintf("%s:%s", hash, key)
	if err := KeyDB.Set(KeyDB.Context(), query, jsonValue, 0).Err(); err != nil {
		log.Println("Error setting key/value store for JSON value", err)
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
		log.Println("Key is empty, cannot delete key/value store,")
		return fmt.Errorf("Key is empty, cannot delete key/value store")
	}

	query := fmt.Sprintf("%s:%s", hash, key)
	err := KeyDB.Del(KeyDB.Context(), query).Err()

	return err
}

func GetJSONKeyDB[T any](hash, key string) (T, error) {
	var value T
	if hash == "" || key == "" {
		log.Println("Key is empty, cannot get JDON value from key/value store,")
		return value, fmt.Errorf("Key is empty, cannot get JDON value from key/value store")
	}

	query := fmt.Sprintf("%s:%s", hash, key)
	jsonValue, err := KeyDB.Get(KeyDB.Context(), query).Result()
	if err != nil {
		log.Println("Error getting key/value store for JSON value", err)
		return value, err
	}

	err = json.Unmarshal([]byte(jsonValue), &value)
	if err != nil {
		log.Println("Error unmarshalling JSON value", err)
		return value, err
	}
	return value, nil
}

func GetDebugDatabase() *gorm.DB {
	debugDB := DB.Session(&gorm.Session{
		Logger: logger.New(
			log.New(os.Stdout, "", log.LstdFlags),
			logger.Config{
				LogLevel: logger.Info,
			}),
	})
	return debugDB
}
