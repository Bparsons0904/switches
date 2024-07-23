package database

import (
	"fmt"
	"log"
	"os"

	"switches/config"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB(config config.Config) *gorm.DB {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)

	var logLevel logger.LogLevel
	if viper.Get("DB_LOGGING") == "true" {
		logLevel = logger.Info
	} else {
		logLevel = logger.Warn
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logLevel), PrepareStmt: true})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}

	SetDatabase(db)

	log.Println("Connected Successfully to the Database")
	return DB
}

// func DBIsInitialized(config utils.Config) (bool, bool) {
// 	var err error
// 	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)
//
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return false, false
// 	}
//
// 	var serviamUser models.User
// 	db.Where("external_id = ?", "serviam").Find(&serviamUser)
// 	if serviamUser.ID == uuid.Nil {
// 		return true, false
// 	}
//
// 	return true, true
// }

func SetDatabase(db *gorm.DB) {
	DB = db
}

func GetDatabase() *gorm.DB {
	return DB
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
