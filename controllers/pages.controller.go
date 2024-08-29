package controllers

import (
	"switches/database"
	"switches/models"
	"switches/templates/pages"
	"switches/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetHomePage(c *fiber.Ctx) error {
	user := c.Locals("User").(models.User)
	return RenderPage(pages.HomePage(user), pages.Home(user))(c)
}

func GetSwitchPage(c *fiber.Ctx) error {
	user := c.Locals("User").(models.User)

	var clickyClacks []models.Switch
	if err := database.DB.
		Order("RANDOM()").
		Find(&clickyClacks).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).Next()
	}

	return RenderPage(pages.SwitchesPage(user, clickyClacks), pages.Switches(user, clickyClacks))(c)
}

func GetSwitchDetailPage(c *fiber.Ctx) error {
	timer := utils.StartTimer("getSwitchDetailPage")
	userID := c.Locals("UserID").(uuid.UUID)
	switchID, err := GetSwitchIDParam(c)
	if err != nil {
		return err
	}

	var user models.User
	if userID != uuid.Nil {
		if err := database.DB.
			Preload("OwnedSwitches").
			Preload("LikedSwitches").
			First(&user, userID).Error; err != nil {
			return c.Status(fiber.StatusBadRequest).Next()
		}
	}
	timer.LogTime("Get User")

	var clickyClack models.Switch
	if err := database.DB.
		Preload("ImageLinks").
		Preload("Brand").
		Preload("SwitchType").
		First(&clickyClack, switchID).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).Next()
	}
	timer.LogTime("Get Switch")

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
