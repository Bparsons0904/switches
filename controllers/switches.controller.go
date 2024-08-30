package controllers

import (
	"fmt"
	"log"
	"switches/database"
	"switches/models"
	"switches/templates/components"
	"switches/utils"
	"time"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetSwitchList(c *fiber.Ctx) error {
	timer := utils.StartTimer("Get Switch List")
	defer timer.LogTotalTime()

	user := c.Locals("User").(models.User)

	type SwitchQueryParams struct {
		SwitchTypeIDs []int  `json:"switchTypeIDs"`
		Search        string `json:"search"`
	}
	request := new(SwitchQueryParams)

	if err := c.QueryParser(request); err != nil {
		log.Println("Error parsing query params", err)
	}

	var clickyClacks []models.Switch
	clickyClackQuery := database.DB.
		Preload("ImageLinks").
		Preload("Brand").
		Preload("SwitchType")

	if len(request.SwitchTypeIDs) > 0 {
		clickyClackQuery.Where("switch_type_id IN (?)", request.SwitchTypeIDs)
	}

	if len(request.Search) > 2 {
		clickyClackQuery.Where("LOWER(name) LIKE LOWER(?)", fmt.Sprintf("%%%s%%", request.Search))
	}

	err := clickyClackQuery.Find(&clickyClacks).Error
	if err != nil {
		log.Println("Error getting the user", err)
		return c.Status(fiber.StatusBadRequest).Next()
	}
	timer.LogTime("Get Switches")

	component := components.SwitchList(user, clickyClacks)
	return Render(component)(c)
}

func GetFeaturedSwitches(c *fiber.Ctx) error {
	timer := utils.StartTimer("Get Featured Switches")
	defer timer.LogTotalTime()

	var switchesFromDB []models.Switch
	database.DB.
		Joins("INNER JOIN image_links ON image_links.owner_id = switches.id").
		Preload("ImageLinks").
		Limit(4).
		Order("RANDOM()").
		Find(&switchesFromDB)
	timer.LogTime("Get Switches")

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

	user, switchID, err := getParams(c)
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

	component := components.SwitchDetailCard(user, clickyClack)
	return Render(component)(c)
}

func getParams(c *fiber.Ctx) (models.User, uuid.UUID, error) {
	user := c.Locals("User").(models.User)
	log.Println("User", user)
	switchID, err := GetSwitchIDParam(c)
	if err != nil {
		log.Println("Error getting the uuid of Switch", err)
	}

	return user, switchID, err
}

func getSwitches(switchID uuid.UUID, timer *utils.Timer) (models.Switch, error) {
	var clickyClack models.Switch
	if err := database.DB.First(&clickyClack, switchID).Error; err != nil {
		return models.Switch{}, err
	}
	timer.LogTime("Get Switch")

	return clickyClack, nil
}

func getUpdatedUser(user models.User, timer *utils.Timer, c *fiber.Ctx) (models.User, error) {
	var updatedUser models.User
	if err := database.DB.
		Preload("OwnedSwitches").
		Preload("LikedSwitches").
		First(&updatedUser, user.ID).Error; err != nil {
		return models.User{}, err
	}
	timer.LogTime("Get User")

	c.Locals("User", updatedUser)
	if err := database.SetUUIDJSONKeyDB("user", updatedUser.ID, updatedUser, 30*time.Hour); err != nil {
		log.Println("Error setting user in keydb", err)
		if err := database.DeleteUUIDKeyDB("user", updatedUser.ID); err != nil {
			log.Println("Error deleting user in keydb", err)
		}
	}
	timer.LogTime("Set User in KeyDB")

	return updatedUser, nil
}

func processUserSwitch(
	c *fiber.Ctx,
	create bool,
	table interface{},
	selectorFunc func(models.User, models.Switch) templ.Component,
	timer *utils.Timer,
) error {
	user, switchID, err := getParams(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Next()
	}

	if user.ID == uuid.Nil {
		log.Println("User not logged in")
		c.Set("HX-Redirect", "/auth/login")
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if create {
		if err := database.DB.Model(table).Create(map[string]interface{}{
			"UserID": user.ID, "SwitchID": switchID,
		}).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).Next()
		}
		timer.LogTime("Create Association")
	} else {
		if err := database.DB.Delete(table, "user_id = ? AND switch_id = ?", user.ID, switchID).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).Next()
		}
		timer.LogTime("Delete Association")
	}

	clickyClack, err := getSwitches(switchID, timer)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Next()
	}

	updatedUser, err := getUpdatedUser(user, timer, c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Next()
	}

	component := selectorFunc(updatedUser, clickyClack)
	return Render(component)(c)
}

func CreateUserOwnedSwitch(c *fiber.Ctx) error {
	timer := utils.StartTimer("Create User Owned Switch")
	defer timer.LogTotalTime()

	return processUserSwitch(
		c,
		true,
		&models.UserOwnedSwitches{},
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
		components.LikedSelector,
		&timer,
	)
}
