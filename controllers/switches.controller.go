package controllers

import (
	"fmt"
	"log"
	"switches/database"
	"switches/models"
	"switches/templates/components"
	"switches/utils"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	defer timer.LogTotalTime()

	userID, switchID, err := getParams(c)
	if err != nil {
		log.Println("Error getting the uuid of Switch", err)
		return err
	}

	var clickyClack models.Switch
	database.DB.
		Preload("ImageLinks").
		Preload("Brand").
		Preload("SwitchType").
		First(&clickyClack, switchID)

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

	component := components.SwitchDetailCard(user, clickyClack)
	return Render(component)(c)
}

func getParams(c *fiber.Ctx) (uuid.UUID, uuid.UUID, error) {
	userID := c.Locals("UserID").(uuid.UUID)
	switchID, err := GetSwitchIDParam(c)
	if err != nil {
		log.Println("Error getting the uuid of Switch", err)
	}

	return userID, switchID, err
}

func getSwitchAndUser(
	userID uuid.UUID,
	switchID uuid.UUID,
	preload string,
	timer *utils.Timer,
) (models.Switch, models.User, error) {
	var clickyClack models.Switch
	if err := database.DB.First(&clickyClack, switchID).Error; err != nil {
		return models.Switch{}, models.User{}, err
	}
	timer.LogTime("Get Switch")

	var user models.User
	if userID != uuid.Nil {
		if err := database.DB.Preload(preload).First(&user, userID).Error; err != nil {
			return models.Switch{}, models.User{}, err
		}
	}
	timer.LogTime("Get User")

	return clickyClack, user, nil
}

func processUserSwitch(
	c *fiber.Ctx,
	create bool,
	table interface{},
	preload string,
	selectorFunc func(models.User, models.Switch) templ.Component,
	timer *utils.Timer,
) error {
	userID, switchID, err := getParams(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Next()
	}

	log.Println("User ID", userID)
	if userID == uuid.Nil {
		log.Println("User not logged in")
		return c.Redirect("/auth/login", fiber.StatusTemporaryRedirect)
	}

	if create {
		if err := database.DB.Model(table).Create(map[string]interface{}{
			"UserID": userID, "SwitchID": switchID,
		}).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).Next()
		}
		timer.LogTime("Create Association")
	} else {
		if err := database.DB.Delete(table, "user_id = ? AND switch_id = ?", userID, switchID).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).Next()
		}
		timer.LogTime("Delete Association")
	}

	clickyClack, user, err := getSwitchAndUser(userID, switchID, preload, timer)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Next()
	}

	component := selectorFunc(user, clickyClack)
	return Render(component)(c)
}

func CreateUserOwnedSwitch(c *fiber.Ctx) error {
	timer := utils.StartTimer("Create User Owned Switch")
	defer timer.LogTotalTime()

	return processUserSwitch(
		c,
		true,
		&models.UserOwnedSwitches{},
		"OwnedSwitches",
		components.OwnedSelector,
		&timer,
	)
}

func DeleteUserOwnedSwitch(c *fiber.Ctx) error {
	timer := utils.StartTimer("Delete User Owned Switch")
	defer timer.LogTotalTime()

	return processUserSwitch(
		c,
		false,
		&models.UserOwnedSwitches{},
		"OwnedSwitches",
		components.OwnedSelector,
		&timer,
	)
}

func CreateUserLikedSwitch(c *fiber.Ctx) error {
	timer := utils.StartTimer("Create User Liked Switch")
	defer timer.LogTotalTime()

	return processUserSwitch(
		c,
		true,
		&models.UserLikedSwitches{},
		"LikedSwitches",
		components.LikedSelector,
		&timer,
	)
}

func DeleteUserLikedSwitch(c *fiber.Ctx) error {
	timer := utils.StartTimer("Delete User Liked Switch")
	defer timer.LogTotalTime()

	return processUserSwitch(
		c,
		false,
		&models.UserLikedSwitches{},
		"LikedSwitches",
		components.LikedSelector,
		&timer,
	)
}
