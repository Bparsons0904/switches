package routes

import (
	"switches/controllers"

	"github.com/gofiber/fiber/v2"
)

func ImageRoutes(app *fiber.App) {
	imagesApi := app.Group("/images")

	imagesApi.Get("/:imageLinkID", controllers.GetImage)
}
