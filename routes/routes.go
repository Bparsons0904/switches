package routes

import (
	"os"
	"runtime"
	"switches/config"
	"switches/controllers"
	"switches/middleware"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

type Health struct {
	Uptime        string `json:"uptime"`
	AppVersion    string `json:"app_version"`
	MemoryUsage   uint64 `json:"memory_usage"`
	NumGoroutine  int    `json:"num_goroutine"`
	NumCPU        int    `json:"num_cpu"`
	DatabaseAlive bool   `json:"database_alive"`
}

var startTime = time.Now().UTC()

func SetupRoutes(app *fiber.App, config config.Config) {
	app.Use(compress.New(compress.Config{
		Level: compress.LevelDefault,
	}))
	app.Use(
		middleware.AuthenticateUser(config),
	)

	app.Get("/", controllers.GetHome)
	AdminRoutes(app)
	AuthRoutes(app)
	SwitchRoutes(app)
	UserRoutes(app)

	api := app.Group("/api")
	HealthRoutes(api)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Route Not found"})
	})
}

func HealthRoutes(app fiber.Router) {
	app.Get("/health", getHealth)
	app.Get("/health/monitor", monitor.New(monitor.Config{
		Title: "Bob's Next Great Thing Health Monitor",
	}))
}

func getHealth(c *fiber.Ctx) error {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	version := os.Getenv("IMAGE_VERSION_TAG")
	health := Health{
		Uptime:       time.Since(startTime).String(),
		AppVersion:   version,
		MemoryUsage:  memStats.Alloc,
		NumGoroutine: runtime.NumGoroutine(),
		NumCPU:       runtime.NumCPU(),
	}

	return c.JSON(health)
}
