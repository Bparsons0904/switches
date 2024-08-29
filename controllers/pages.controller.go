package controllers

import (
	"log"
	"switches/database"
	"switches/models"
	"switches/templates/pages"
	"switches/utils"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetHomePage(c *fiber.Ctx) error {
	user := c.Locals("User").(models.User)
	return RenderPage(pages.HomePage(user), pages.Home(user))(c)
}

func GetSwitchPage(c *fiber.Ctx) error {
	timer := utils.StartTimer("Get Switch Page")
	defer timer.LogTotalTime()

	userID := c.Locals("UserID").(uuid.UUID)

	wg := sync.WaitGroup{}
	wg.Add(2)

	var err error
	var user models.User
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if userID != uuid.Nil {
			err = database.DB.
				Preload("OwnedSwitches").
				Preload("LikedSwitches").
				First(&user, userID).Error
		}
	}(&wg)

	var clickyClacks []models.Switch
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		err = database.DB.
			Preload("ImageLinks").
			Preload("Brand").
			Preload("SwitchType").
			Order("RANDOM()").
			Find(&clickyClacks).Error
	}(&wg)

	wg.Wait()
	if err != nil {
		log.Println("Error getting the user or switches", err)
		return c.Status(fiber.StatusBadRequest).Next()
	}

	return RenderPage(pages.SwitchesPage(user, clickyClacks), pages.Switches(user, clickyClacks))(c)
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
