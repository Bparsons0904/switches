package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"switches/config"
	"switches/database"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	Config config.Config
)

func init() {
	log.SetOutput(os.Stdout)

	var err error
	Config, err = config.LoadConfig()
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	DB = database.ConnectDB(Config)
}

func main() {
	args := os.Args

	for i, arg := range args[1:] {
		switch {
		case arg == "-up":
			log.Println("Running Up")
			migrateUp()
			log.Println("All Migrations Completed")
		case arg == "-down":
			log.Println("Running Down")
			migrateDown()
			log.Println("All Migrations Completed")
		case strings.HasPrefix(arg, "-create"):
			// If -create is provided, expect the next argument to be the migration name
			if len(args) <= i+1 {
				log.Fatal("Expected migration name after -create")
			}

			if len(args) > i+2 { // i+2 because the arguments slice includes the program name and we're 1-indexed here
				migrationName := args[i+2]
				err := createMigrationFile(migrationName)
				if err != nil {
					log.Fatalf("Error creating migration file: %v", err)
				}
			} else {
				log.Fatal("Expected migration name after -create")
			}
		}
	}
}

var MigrationOptions = &gormigrate.Options{
	TableName:                 "migrations",
	IDColumnName:              "id",
	IDColumnSize:              255,
	UseTransaction:            true,
	ValidateUnknownMigrations: false,
}

func migrateUp() {
	// if len(migrations.RegisteredMigrations) == 0 {
	// 	log.Println("No migrations registered")
	// 	return
	// }
	//
	// sort.SliceStable(migrations.RegisteredMigrations, func(i, j int) bool {
	// 	return migrations.RegisteredMigrations[i].ID < migrations.RegisteredMigrations[j].ID
	// })
	//
	// m := gormigrate.New(DB, MigrationOptions, migrations.RegisteredMigrations)
	//
	// if err := m.Migrate(); err != nil {
	// 	log.Fatalf("Migration failed: %v", err)
	// }
}

func migrateDown() {
	// if len(migrations.RegisteredMigrations) == 0 {
	// 	log.Println("No migrations registered")
	// 	return
	// }
	//
	// // Sort in reverse order
	// sort.SliceStable(migrations.RegisteredMigrations, func(i, j int) bool {
	// 	return migrations.RegisteredMigrations[i].ID > migrations.RegisteredMigrations[j].ID
	// })
	//
	// m := gormigrate.New(DB, MigrationOptions, migrations.RegisteredMigrations)
	//
	// if err := m.RollbackLast(); err != nil {
	// 	log.Fatalf("Migration failed: %v", err)
	// }
}

func createMigrationFile(name string) error {
	// Create a timestamp prefix for the file to ensure they're processed in the order they were created
	timestamp := time.Now().Format("20060102150405") // YYYYMMDDHHMMSS
	filename := fmt.Sprintf("%s_%s.go", timestamp, strings.Replace(name, " ", "_", -1))
	dir := "./migrator/migrations"

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	filePath := fmt.Sprintf("%s/%s", dir, filename)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	migrationTemplate := `package migrations

import (
	"gorm.io/gorm"
)

func init() {
	RegisterMigration(Migration{
		ID:          "%s",
		Description: "Add description of changes",
		Migrate: func(tx *gorm.DB) error {
			// Your migration logic goes here.
			// return tx.AddColumn(&YourModel{}, "name").Error
			return nil // Replace with actual code
		},
		Rollback: func(tx *gorm.DB) error {
			// Your rollback logic goes here.
			// return tx.Migrator().AlterColumn(&YourModel{}, "name")
			return nil // Replace with actual code
		},
	})
}`

	migrationContent := fmt.Sprintf(migrationTemplate, timestamp)
	_, err = file.WriteString(migrationContent)
	if err != nil {
		return err
	}

	log.Printf("Migration file created: %s", filePath)
	return nil
}
