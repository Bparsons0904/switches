package main

import (
	"fmt"
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
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
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"

	"github.com/tdewolff/minify/v2/js"
)

var (
	server *fiber.App
	config env.Config
	m      *minify.M
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
	m = minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/js", js.Minify)
	// go scheduler.InitScheduler(db)
}

func main() {
	log.Info().Str("domain", config.BaseURL).Msg("Starting server...")
	setStatic(config.AppendNumber, server)
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

func setStatic(appendNumber int64, server *fiber.App) {
	stylesDir := "./static/styles"
	scriptsDir := "./static/scripts"
	destDir := "./assets"

	if err := os.RemoveAll(destDir); err != nil {
		fmt.Printf("Error removing existing destination directory: %v\n", err)
		return
	}

	if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
		fmt.Printf("Error creating destination directory: %v\n", err)
		return
	}

	// log.Println("Copying static files...")
	err := copyStaticFiles(stylesDir, destDir, appendNumber)
	if err != nil {
		log.Fatal().Err(err).Msg("Error copying static style files")
	}

	err = copyStaticFiles(scriptsDir, destDir, appendNumber)
	if err != nil {
		log.Fatal().Err(err).Msg("Error copying static script files")
	}

	server.Static("/assets", "./assets")
	server.Static("/images", "./static/assets/images")
}

func copyStaticFiles(sourceDir, destDir string, appendNumber int64) error {
	files, err := os.ReadDir(sourceDir)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		oldPath := filepath.Join(sourceDir, file.Name())

		ext := filepath.Ext(file.Name())
		baseName := strings.TrimSuffix(file.Name(), ext)
		newFileName := fmt.Sprintf("%s_%d.min%s", baseName, appendNumber, ext)
		newPath := filepath.Join(destDir, newFileName)

		if err := copyAndMinifyFile(oldPath, newPath); err != nil {
			log.Error().
				Err(err).
				Str("filename", file.Name()).
				Msg("Error copying and minifying file")
			continue
		}
	}
	return nil
}

func copyAndMinifyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	if strings.HasSuffix(strings.ToLower(src), ".css") {
		err = m.Minify("text/css", destFile, sourceFile)
		if err != nil {
			return fmt.Errorf("minification error: %v", err)
		}
	} else if strings.HasSuffix(strings.ToLower(src), ".js") {
		err = m.Minify("text/js", destFile, sourceFile)
		if err != nil {
			return fmt.Errorf("minification error: %v", err)
		}
	} else {
		_, err = io.Copy(destFile, sourceFile)
		if err != nil {
			return err
		}
	}

	return nil
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
