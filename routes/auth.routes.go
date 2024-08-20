package routes

import (
	"switches/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	auth := app.Group("/auth")

	auth.Get("/login", controllers.AuthLogin)
	auth.Get("/callback", controllers.AuthCallback)
	auth.Get("/logout", controllers.UserLogout)
	auth.Get("/logout/callback", controllers.UserLogoutCallback)
}
