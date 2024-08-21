package controllers

import (
	"switches/models"
	"switches/templates/pages"

	"github.com/gofiber/fiber/v2"
)

func GetHome(c *fiber.Ctx) error {
	user := c.Locals("User").(models.User)
	return Render(pages.HomePage(user), pages.Home(user))(c)
}

func GetSwitches(c *fiber.Ctx) error {
	user := c.Locals("User").(models.User)
	return Render(pages.SwitchesPage(user), pages.Switches(user))(c)
}
