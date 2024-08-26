package controllers

import (
	"switches/database"
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

func GetSwitchModal(c *fiber.Ctx) error {
	switchID := c.Params("switchID")

	var switchModel models.Switch
	database.DB.First(&switchModel, switchID)

	component := components.SwitchModalCard(switchModel)
	return Render(component)(c)
}
