package controllers

import (
	"log"
	"switches/database"
	"switches/models"
	"switches/templates/pages"
	"switches/utils"

	"github.com/gofiber/fiber/v2"
)

func GetHome(c *fiber.Ctx) error {
	timer := utils.StartTimer("Get home")
	sessionID := c.Cookies("sessionID")
	var retrievedSession Session
	if sessionID != "" {
		log.Println("Session ID found", sessionID)
		var err error
		retrievedSession, err = database.GetJSONKeyDB[Session]("session", sessionID)
		timer.LogTime("Get session from keydb")
		log.Println("Retrieved session", retrievedSession)
		if err != nil {
			log.Println("Error getting session from keydb", err)
			c.ClearCookie("sessionID")
		}
	}

	var user models.User
	if err := database.DB.First(&user, retrievedSession.UserID).Error; err != nil {
		log.Println("Error getting user from keydb", err)
	}

	timer.LogTime("Get User from DB")
	log.Println("GetHome")
	return Render(pages.HomePage(), pages.Home())(c)
}

func GetSwitches(c *fiber.Ctx) error {
	log.Println("GetSwitches")
	return Render(pages.SwitchesPage(), pages.Switches())(c)
}
