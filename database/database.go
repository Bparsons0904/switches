package database

import (
	"fmt"
	"log"
	"os"
	"switches/config"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
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

	server.Use(DBMiddleware(DB, KeyDB))
	return DB, KeyDB
}

func DBMiddleware(db *gorm.DB, keydb *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("db", db)
		c.Locals("keydb", keydb)
		return c.Next()
	}
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
	SetDatabase(db)
	log.Println("Connected Successfully to the PostgreSQL Database")

	return db, nil
}

func SetDatabase(db *gorm.DB) {
	DB = db
}

func GetDatabase() *gorm.DB {
	return DB
}

func StartKeyDB(config config.Config) (*redis.Client, error) {
	keydb := redis.NewClient(&redis.Options{
		Addr:     config.KeyDBHost,
		Password: config.KeyDBPassword,
		DB:       int(config.KeyDB),
	})

	_, err := keydb.Ping(keydb.Context()).Result()
	if err != nil {
		return nil, err
	}

	SetKeyDatabase(keydb)
	log.Println("Connected Successfully to the KeyDB Database")

	return KeyDB, nil
}

func SetKeyDatabase(keydb *redis.Client) {
	KeyDB = keydb
}

func GetKeyDatabase() *redis.Client {
	return KeyDB
}

func GetDebugDatabase() *gorm.DB {
	debugDB := DB.Session(&gorm.Session{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				LogLevel: logger.Info,
			}),
	})
	return debugDB
}
