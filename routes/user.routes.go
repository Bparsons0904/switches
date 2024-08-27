package routes

import (
	"switches/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	users := app.Group("/users")

	users.Get("/", UserPlaceHolder)
	users.Post("/switch/owned", controllers.HandleUserOwnedSwitch)
	users.Delete("/switch/owned", controllers.DeleteUserOwnedSwitch)
	users.Post("/switch/liked", controllers.HandleUserLikedSwitch)
	users.Delete("/switch/liked", controllers.DeleteUserLikedSwitch)
}

func UserPlaceHolder(c *fiber.Ctx) error {
	return c.SendString("User Placeholder")
}
