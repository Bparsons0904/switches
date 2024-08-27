package controllers

import (
	"fmt"
	"log"
	"switches/database"
	"switches/models"
	"switches/templates/components"
	"switches/templates/pages"
	"switches/utils"

	"github.com/gofiber/fiber/v2"
)

func GetSwitches(c *fiber.Ctx) error {
	user := c.Locals("User").(models.User)
	return RenderPage(pages.SwitchesPage(user), pages.Switches(user))(c)
}

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

func GetSwitchDetail(c *fiber.Ctx) error {
	timer := utils.StartTimer("getSwitchModal")
	switchID := c.Query("switchID")

	log.Println(switchID)
	var switchModel models.Switch
	database.DB.Preload("ImageLinks").First(&switchModel, "id = ?", switchID)

	timer.LogTotalTime()
	component := components.SwitchDetailCard(switchModel)
	return Render(component)(c)
}
