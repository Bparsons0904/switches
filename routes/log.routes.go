package routes

import (
	"switches/controllers"

	"github.com/gofiber/fiber/v2"
)

func LogRoutes(app fiber.Router) {
	logsApi := app.Group("/logs")

	logsApi.Post("/errors/images", controllers.PostImageLog)
}
