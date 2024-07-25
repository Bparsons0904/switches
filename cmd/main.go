package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"switches/database"
	"switches/routes"

	env "switches/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/gorm"
)

var (
	server *fiber.App
	config env.Config
)

func init() {
	log.SetOutput(os.Stdout)

	var err error
	config, err = env.LoadConfig()
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	log.Println("Environment variables loaded successfully", config)

	if config.Tier == "development" {
		server = fiber.New(fiber.Config{
			ReadBufferSize:    16384,
			StreamRequestBody: true,
		})
	} else {
		server = fiber.New(fiber.Config{
			DisableStartupMessage: true,
			StreamRequestBody:     true,
			ReadBufferSize:        16384,
		})
	}

	// go scheduler.InitScheduler(db)
}

func LoadEnvMiddleware(c *fiber.Ctx) error {
	c.Locals("clientOrigin", config.BaseURL)
	return c.Next()
}

func DBMiddleware(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	}
}

func main() {
	log.Println("Starting server...", config.BaseURL)
	server.Use(cors.New(cors.Config{
		AllowOrigins:     config.BaseURL,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, withCredentials, X-Response-Type",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: true,
	}))
	server.Use(LoadEnvMiddleware)
	db := database.ConnectDB(config)
	server.Use(DBMiddleware(db))

	if config.Tier == "local" {
		server.Use(logger.New())
	}

	server.Use(recover.New())
	server.Use(func(c *fiber.Ctx) error {
		// Handle Preflight Request
		if c.Method() == "OPTIONS" {
			c.Set("Access-Control-Allow-Origin", config.BaseURL)
			c.Set(
				"Access-Control-Allow-Headers",
				"Origin, Content-Type, Accept, Authorization, withCredentials, X-Response-Type",
			)
			c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Set("Access-Control-Allow-Credentials", "true")
			c.Status(fiber.StatusNoContent)
			return nil
		}

		return c.Next()
	})

	server.Static("/styles", "./static/styles")
	server.Static("/scripts", "./static/scripts")
	server.Static("/images", "./static/assets/images")
	routes.SetupRoutes(server, config)

	// Create a channel to listen for a shutdown signal
	go setupGracefulShutdown(server)

	log.Println("Server is running in", config.Tier, "mode on port", config.ServerPort)
	log.Fatal(server.Listen(":" + config.ServerPort))
}

func setupGracefulShutdown(server *fiber.App) {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-channel
		log.Println("Received termination signal, gracefully shutting down...")

		if err := server.Shutdown(); err != nil {
			log.Fatal("Server forced to shutdown:", err)
		}
	}()
}
