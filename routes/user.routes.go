package routes

import (
	"switches/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	users := app.Group("/users")

	users.Get("/", controllers.UserPlaceHolder)
}
