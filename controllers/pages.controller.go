package controllers

import (
	"switches/models"
	"switches/templates/pages"

	"github.com/gofiber/fiber/v2"
)

func GetHome(c *fiber.Ctx) error {
	user := c.Locals("User").(models.User)
	return RenderPage(pages.HomePage(user), pages.Home(user))(c)
}
