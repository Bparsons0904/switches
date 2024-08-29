package routes

import (
	"switches/controllers"

	"github.com/gofiber/fiber/v2"
)

func SwitchRoutes(app *fiber.App) {
	switches := app.Group("/switches")

	switches.Get("/", controllers.GetSwitchPage)
	switches.Get("/featured", controllers.GetFeaturedSwitches)
	switches.Get("/:switchID", controllers.GetSwitchDetailPage)
	switches.Get("/:switchID/modal", controllers.GetSwitchDetailCard)
	switches.Post("/:switchID/owned", controllers.CreateUserOwnedSwitch)
	switches.Delete("/:switchID/owned", controllers.DeleteUserOwnedSwitch)
	switches.Post("/:switchID/liked", controllers.CreateUserLikedSwitch)
	switches.Delete("/:switchID/liked", controllers.DeleteUserLikedSwitch)
}
