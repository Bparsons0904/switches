package routes

import "github.com/gofiber/fiber/v2"

func UserRoutes(app *fiber.App) {
	user := app.Group("/user")

	user.Get("/", UserPlaceHolder)
}

func UserPlaceHolder(c *fiber.Ctx) error {
	return c.SendString("User Placeholder")
}
