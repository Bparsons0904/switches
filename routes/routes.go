package routes

import (
	"log"
	"os"
	"runtime"
	"switches/config"
	"switches/controllers"
	"switches/database"
	"time"

	"github.com/gofiber/fiber/v2"
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
	// HealthRoutes(app)
	// ProxyRoutes(app, config)

	// Websocket routes
	// service := websockets.NewHigherPathWebSocketService(oidcProvider)
	// Add service to context
	// app.Use(func(c *fiber.Ctx) error {
	// 	c.Locals("service", service)
	// 	return c.Next()
	// })
	// app.Use("/thp", websocket.New(func(c *websocket.Conn) {
	// 	service.HandleWebSocketConnection(c)
	// }))

	// app.Get("/version", func(c *fiber.Ctx) error {
	// 	return c.JSON(fiber.Map{
	// 		"version": config.Version,
	// 	})
	// })

	// app.Use(compress.New(compress.Config{
	// 	Level: compress.LevelDefault,
	// }))
	app.Get("/auth/callback", controllers.GetAuthCallback)
	app.Get("/switches", controllers.GetSwitches)
	app.Get("/", controllers.GetHome)

	api := app.Group("/api")
	HealthRoutes(api)
	// AuthRoutes(api, config)
	// api.Use(middleware.AuthenticateUser(config))
	// api.Use(compress.New(compress.Config{
	// 	Level: compress.LevelDefault,
	// }))

	// AccountRoutes(api)
	// SubAccountRoutes(api)
	// UserRoutes(api)
	// LeadRoutes(api)
	// NotificationRoutes(api)
	// ReportsRoutes(api)
	// DevToolsRoutes(api)
	// IntegrationRoutes(api)
	// Return 404 on all internal routes not implemented
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

	sqlDB, err := database.GetDatabase().DB()
	if err != nil {
		log.Println("Error getting database: ", err)
	}
	dbAlive := sqlDB.Ping() == nil

	version := os.Getenv("IMAGE_VERSION_TAG")

	health := Health{
		Uptime:        time.Since(startTime).String(),
		AppVersion:    version,
		MemoryUsage:   memStats.Alloc,
		NumGoroutine:  runtime.NumGoroutine(),
		NumCPU:        runtime.NumCPU(),
		DatabaseAlive: dbAlive,
	}

	return c.JSON(health)
}
