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
	return Render(pages.Home())(c)
}

func GetAdminHome(c *fiber.Ctx) error {
	user := c.Locals("User").(models.User)
	return Render(pages.Admin(user))(c)
}

func GetAdminSwitchEdit(c *fiber.Ctx) error {
	user := c.Locals("User").(models.User)
	return Render(pages.SwitchEdit(user))(c)
}

func GetSwitchPage(c *fiber.Ctx) error {
	timer := utils.StartTimer("Get Switch Page")
	defer timer.LogTotalTime()

	var clickyClacks []models.Switch
	err := database.DB.
		Preload("ImageLinks").
		Preload("Brand").
		Preload("SwitchType").
		Find(&clickyClacks).Error
	if err != nil {
		// log.Println("Error getting the user", err)
		return c.Status(fiber.StatusBadRequest).Next()
	}
	timer.LogTime("Get Switches")

	var switchTypes []models.Type
	err = database.DB.
		Select("id", "name").
		Order("id").
		Where("category = ?", "switch_type").
		Find(&switchTypes).Error
	if err != nil {
		// log.Println("Error getting the switch types", err)
		return c.Status(fiber.StatusBadRequest).Next()
	}
	timer.LogTime("Get Switch Types")

	var switchBrands []models.Producer
	if err := database.DB.
		Find(&switchBrands).Error; err != nil {
		// log.Println("Error getting the switch brands", err)
		return c.Status(fiber.StatusBadRequest).Next()
	}
	timer.LogTime("Get Switch Brands")

	props := pages.SwitchesPageProps{
		ClickyClacks: clickyClacks,
		SwitchTypes:  switchTypes,
		SwitchBrands: switchBrands,
		User:         c.Locals("User").(models.User),
	}

	return Render(
		pages.Switches(props),
	)(
		c,
	)
}

func GetSwitchDetailPage(c *fiber.Ctx) error {
	timer := utils.StartTimer("getSwitchDetailPage")
	defer timer.LogTotalTime()

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

	return Render(
		pages.SwitchDetail(user, clickyClack),
	)(
		c,
	)
}

func NotFound(c *fiber.Ctx) error {
	return Render(pages.NotFound())(c)
}
