package controllers

import (
	"log"
	"switches/database"
	"switches/models"
	"switches/templates/pages"

	"github.com/gofiber/fiber/v2"
)

func GetHome(c *fiber.Ctx) error {
	sessionID := c.Cookies("sessionID")
	var user models.User
	if sessionID != "" {
		retrievedSession, err := database.GetJSONKeyDB[Session]("session", sessionID)
		if err != nil {
			log.Println("Error getting session from keydb", err)
			c.ClearCookie("sessionID")
		}

		user, err = database.GetJSONKeyDB[models.User]("user", retrievedSession.UserID.String())
		if err != nil {
			if err := database.DB.First(&user, retrievedSession.UserID).Error; err != nil {
				log.Println("Error getting user from keydb", err)
			} else {
				database.SetJSONKeyDB("user", user.ID.String(), user)
			}
		}
	}

	log.Println("GetHome")
	return Render(pages.HomePage(user), pages.Home(user))(c)
}

func GetSwitches(c *fiber.Ctx) error {
	log.Println("GetSwitches")
	return Render(pages.SwitchesPage(), pages.Switches())(c)
}
