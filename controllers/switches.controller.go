package controllers

import (
	"switches/models"
	"switches/templates/components"
	"switches/templates/pages"

	"github.com/gofiber/fiber/v2"
)

func GetSwitches(c *fiber.Ctx) error {
	user := c.Locals("User").(models.User)
	return RenderPage(pages.SwitchesPage(user), pages.Switches(user))(c)
}

func GetFeaturedSwitches(c *fiber.Ctx) error {
	component := components.FeaturedSwitches()
	return Render(component)(c)
}
