package routes

import (
	"switches/controllers"

	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(app *fiber.App) {
	// adminApi := app.Group("/admin", middleware.IsAdmin())
	adminApi := app.Group("/admin")

	adminApi.Get("/", controllers.GetAdminHome)
	adminApi.Get("/switches", controllers.GetAdminSwitches)
	adminApi.Get("/switches/create", controllers.GetAdminSwitchCreate)
	adminApi.Post("/switches/create", controllers.PostAdminSwitchCreate)
	adminApi.Get("/images/create", controllers.GetImageLinkToList)
	adminApi.Delete("/images/:imageLinkIndex", controllers.DeleteImageLinkToList)
	adminApi.Get("/switches/:switchID/edit", controllers.GetAdminSwitchEdit)
}

func AdminPlaceHolder(c *fiber.Ctx) error {
	return c.SendString("Admin Placeholder")
}
