package main

import (
	"fmt"
	"log"
	"os"

	"switches/config"
	_ "switches/migrator/migrations"

	"github.com/Bparsons0904/deadigations"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Could not load environment variables:", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)

	migrationTool := deadigations.NewMigrationTool(dsn)

	migrationTool.Run(os.Args)
}
