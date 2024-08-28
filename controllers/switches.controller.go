package controllers

import (
	"fmt"
	"log"
	"switches/database"
	"switches/models"
	"switches/templates/components"
	"switches/utils"

	"github.com/gofiber/fiber/v2"
)

func GetFeaturedSwitches(c *fiber.Ctx) error {
	timer := utils.StartTimer("getFeaturedSwitches")
	var switchesFromDB []models.Switch
	database.DB.
		Joins("INNER JOIN image_links ON image_links.owner_id = switches.id").
		Preload("ImageLinks").
		Limit(4).
		Order("RANDOM()").
		Find(&switchesFromDB)

	timer.LogTotalTime()

	var switches []components.SwitchCardProps
	for _, clickyClack := range switchesFromDB {
		imageLink := clickyClack.ImageLinks[0].LinkAddress
		imageAlt := fmt.Sprintf("Image of %s switch", clickyClack.Name)
		newSwitch := components.SwitchCardProps{
			ID:          clickyClack.ID.String(),
			Title:       clickyClack.Name,
			Description: clickyClack.ShortDescription,
			ImageSrc:    imageLink,
			ImageAlt:    imageAlt,
		}

		switches = append(switches, newSwitch)
	}

	component := components.FeaturedSwitches(switches)
	return Render(component)(c)
}

func GetSwitchDetailCard(c *fiber.Ctx) error {
	timer := utils.StartTimer("getSwitchModal")
	user := c.Locals("User").(models.User)
	switchID, err := GetSwitchIDQuery(c)
	if err != nil {
		log.Println("Error getting the uuid of Switch", err)
		return err
	}

	var switchModel models.Switch
	database.DB.
		Preload("ImageLinks").
		Preload("Brand").
		Preload("SwitchType").
		Preload("OwnedSwitches").
		Preload("LikedSwitches").
		First(&switchModel, switchID)

	timer.LogTotalTime()
	component := components.SwitchDetailCard(user, switchModel)
	return Render(component)(c)
}

func CreateUserOwnedSwitch(c *fiber.Ctx) error {
	timer := utils.StartTimer("Create user owned Switch")
	user := c.Locals("User").(models.User)
	switchID, err := GetSwitchIDParam(c)
	if err != nil {
		log.Println("Error getting the uuid of Switch", err)
		return err
	}

	var clickyClack models.Switch
	if err := database.DB.First(&clickyClack, switchID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).Next()
	}

	err = database.DB.Model(&user).Association("OwnedSwitches").Append(&models.Switch{ID: switchID})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Next()
	}

	timer.LogTotalTime()
	component := components.OwnedSelector(user, clickyClack)
	return Render(component)(c)
}

func DeleteUserOwnedSwitch(c *fiber.Ctx) error {
	timer := utils.StartTimer("Delete user owned Switch")
	user := c.Locals("User").(models.User)
	switchID, err := GetSwitchIDParam(c)
	if err != nil {
		log.Println("Error getting the uuid of Switch", err)
		return err
	}

	var clickyClack models.Switch
	if err := database.DB.First(&clickyClack, switchID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).Next()
	}

	err = database.DB.Model(&user).Association("OwnedSwitches").Delete(&models.Switch{ID: switchID})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Next()
	}

	timer.LogTotalTime()
	component := components.OwnedSelector(user, clickyClack)
	return Render(component)(c)
}

func CreateUserLikedSwitch(c *fiber.Ctx) error {
	timer := utils.StartTimer("Create user liked Switch")
	user := c.Locals("User").(models.User)
	switchID, err := GetSwitchIDParam(c)
	if err != nil {
		log.Println("Error getting the uuid of Switch", err)
		return err
	}

	var clickyClack models.Switch
	if err := database.DB.First(&clickyClack, switchID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).Next()
	}

	err = database.DB.Model(&user).Association("LikedSwitches").Append(&models.Switch{ID: switchID})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Next()
	}

	timer.LogTotalTime()
	component := components.LikedSelector(user, clickyClack)
	return Render(component)(c)
}

func DeleteUserLikedSwitch(c *fiber.Ctx) error {
	timer := utils.StartTimer("Delete user liked Switch")
	user := c.Locals("User").(models.User)
	switchID, err := GetSwitchIDParam(c)
	if err != nil {
		log.Println("Error getting the uuid of Switch", err)
		return err
	}

	var clickyClack models.Switch
	if err := database.DB.First(&clickyClack, switchID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).Next()
	}

	err = database.DB.Model(&user).Association("LikedSwitches").Delete(&models.Switch{ID: switchID})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Next()
	}

	timer.LogTotalTime()
	component := components.LikedSelector(user, clickyClack)
	return Render(component)(c)
}
