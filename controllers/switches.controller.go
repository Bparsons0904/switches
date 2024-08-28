package controllers

import (
	"fmt"
	"log"
	"switches/database"
	"switches/models"
	"switches/templates/components"
	"switches/templates/components/icons"
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
	switchID, err := GetSwitchIDQuery(c)
	if err != nil {
		log.Println("Error getting the uuid of Switch", err)
		return err
	}

	var switchModel models.Switch
	database.DB.Preload("ImageLinks").First(&switchModel, switchID)

	timer.LogTotalTime()
	component := components.SwitchDetailCard(switchModel)
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

	err = database.DB.Model(&user).Association("OwnedSwitches").Append(&models.Switch{ID: switchID})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Next()
	}

	timer.LogTotalTime()
	component := icons.BookmarkCheckFilled()
	return Render(component)(c)
}

func DeleteUserOwnedSwitch(c *fiber.Ctx) error {
	return c.SendString("Delete User Owned Switch")
}

func CreateUserLikedSwitch(c *fiber.Ctx) error {
	return c.SendString("Handle User Owned Switch")
}

func DeleteUserLikedSwitch(c *fiber.Ctx) error {
	return c.SendString("Add User Owned Switch")
}
