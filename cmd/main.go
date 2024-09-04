package main

import (
	"os"
	"os/signal"
	"switches/database"
	"switches/middleware"
	"switches/routes"
	"switches/utils"
	"syscall"

	env "switches/config"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog/log"
)

var (
	server *fiber.App
	config env.Config
)

func init() {
	var err error
	config, err = env.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("? Could not load environment variables")
	}

	if config.Tier == "development" {
		server = fiber.New(fiber.Config{
			ReadBufferSize:           16384,
			StreamRequestBody:        true,
			EnableSplittingOnParsers: true,
		})
	} else {
		server = fiber.New(fiber.Config{
			DisableStartupMessage:    true,
			StreamRequestBody:        true,
			ReadBufferSize:           16384,
			EnableSplittingOnParsers: true,
		})
	}
	// go scheduler.InitScheduler(db)
}

func main() {
	log.Info().Str("domain", config.BaseURL).Msg("Starting server...")
	utils.SetStatic(config.AppendNumber, server)
	database.ConnectDB(config, server)
	zlog := utils.GetLogger()
	server.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &zlog,
	}))

	// server.Use(limiter.New())
	server.Use(recover.New())
	// server.Use(helmet.New())
	// server.Use(csrf.New())
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
	server.Use(cors.New(cors.Config{
		AllowOrigins:     config.BaseURL,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, withCredentials, X-Response-Type",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: true,
	}))

	server.Use(compress.New(compress.Config{
		Level: compress.LevelDefault,
	}))
	server.Use(
		middleware.AuthenticateUser(config),
	)

	routes.SetupRoutes(server, config)

	// Create a channel to listen for a shutdown signal
	go setupGracefulShutdown(server)

	log.Info().Msg("Server is running in " + config.Tier + " mode on port " + config.ServerPort)
	serverErr := server.Listen(":" + config.ServerPort)
	log.Error().Err(serverErr).Msg("Server error")
}

func setupGracefulShutdown(server *fiber.App) {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-channel
		log.Info().Msg("Received termination signal, gracefully shutting down...")

		if err := server.Shutdown(); err != nil {
			log.Error().Err(err).Msg("Server forced to shutdown")
		}
	}()
}
