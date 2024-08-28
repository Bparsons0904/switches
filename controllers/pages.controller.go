package controllers

import (
	"switches/database"
	"switches/models"
	"switches/templates/pages"
	"switches/utils"

	"github.com/gofiber/fiber/v2"
)

func GetHomePage(c *fiber.Ctx) error {
	user := c.Locals("User").(models.User)
	return RenderPage(pages.HomePage(user), pages.Home(user))(c)
}

func GetSwitchPage(c *fiber.Ctx) error {
	user := c.Locals("User").(models.User)
	return RenderPage(pages.SwitchesPage(user), pages.Switches(user))(c)
}

func GetSwitchDetailPage(c *fiber.Ctx) error {
	timer := utils.StartTimer("getSwitchDetailPage")
	user := c.Locals("User").(models.User)
	switchID, err := GetSwitchIDParam(c)
	if err != nil {
		return err
	}

	var clickyClack models.Switch
	if err := database.DB.
		Preload("ImageLinks").
		Preload("Brand").
		Preload("SwitchType").
		First(&clickyClack, switchID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).Redirect("/404")
	}

	timer.LogTotalTime()
	return RenderPage(
		pages.SwitchDetailPage(user, clickyClack),
		pages.SwitchDetail(user, clickyClack),
	)(
		c,
	)
}

func NotFound(c *fiber.Ctx) error {
	return RenderPage(pages.NotFound(), pages.NotFound())(c)
}
