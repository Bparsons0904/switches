package routes

import "github.com/gofiber/fiber/v2"

func AdminRoutes(app *fiber.App) {
	admin := app.Group("/admin")

	admin.Get("/", AdminPlaceHolder)
}

func AdminPlaceHolder(c *fiber.Ctx) error {
	return c.SendString("Admin Placeholder")
}
