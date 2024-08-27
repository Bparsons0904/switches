package routes

import (
	"switches/controllers"

	"github.com/gofiber/fiber/v2"
)

func SwitchRoutes(app *fiber.App) {
	switches := app.Group("/switches")

	switches.Get("/", controllers.GetSwitchPage)
	switches.Get("/featured", controllers.GetFeaturedSwitches)
	switches.Get("/modal", controllers.GetSwitchDetailCard)
	switches.Get("/:switchID", controllers.GetSwitchDetailPage)
}
