package routes

import (
	"switches/controllers"
	"switches/middleware"

	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(app *fiber.App) {
	adminApi := app.Group("/admin", middleware.IsAdmin())

	adminApi.Get("/", controllers.GetAdminHome)
	adminApi.Get("/switches/:switchID/edit", controllers.GetAdminSwitchEdit)
}

func AdminPlaceHolder(c *fiber.Ctx) error {
	return c.SendString("Admin Placeholder")
}
