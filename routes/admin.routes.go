package routes

import (
	"switches/controllers"
	"switches/middleware"

	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(app *fiber.App) {
	adminApi := app.Group("/admin", middleware.IsAdmin())

	adminApi.Get("/", controllers.GetAdminHome)
	adminApi.Get("/switches", controllers.GetAdminSwitches)
	adminApi.Get("/switches/create", controllers.GetAdminSwitchCreate)
	adminApi.Post("/switches/create", controllers.PostAdminSwitchCreate)
	adminApi.Get("/images/create", controllers.GetImageLinkToList)
	adminApi.Delete("/images/:imageLinkIndex", controllers.DeleteImageLinkToList)
	adminApi.Get("/switches/:switchID/edit", controllers.GetAdminSwitchEdit)
	adminApi.Put("/switches/:switchID/edit", controllers.PutSwitch)
}

func AdminPlaceHolder(c *fiber.Ctx) error {
	return c.SendString("Admin Placeholder")
}
