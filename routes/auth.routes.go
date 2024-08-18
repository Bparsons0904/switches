package routes

import (
	"switches/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	auth := app.Group("/auth")

	auth.Get("/callback", controllers.GetAuthCallback)
	auth.Get("/logout", controllers.UserLogout)
}
