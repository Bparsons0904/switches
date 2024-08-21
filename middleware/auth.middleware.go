package middleware

import (
	"log"
	"switches/config"
	"switches/database"
	"switches/models"
	"switches/services"
	"time"

	"github.com/gofiber/fiber/v2"
)

func AuthenticateUser(config config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Println("AuthenticateUser started")
		sessionID := c.Cookies("sessionID")
		var user models.User
		if sessionID != "" {
			log.Println("sessionID was found in the cookies", sessionID)
			retrievedSession, err := services.SessionFlow(sessionID)
			if err != nil {
				c.ClearCookie("sessionID")
				c.Locals("User", user)
				log.Println("Error getting session from keydb", err)
				return c.Next()
			}

			user, err = database.GetJSONKeyDB[models.User]("user", retrievedSession.UserID.String())
			if err != nil {
				if err := database.DB.First(&user, retrievedSession.UserID).Error; err != nil {
					log.Println("Error getting user from keydb", err)
				} else {
					err := database.SetJSONKeyDB("user", user.ID.String(), user, 30*time.Hour)
					if err != nil {
						log.Println("Error setting user in keydb", err)
					}
				}
			}
		}

		c.Locals("User", user)
		log.Println("setting user in local: ", user.ID)
		return c.Next()
	}
}

func IsAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Println("IsAdmin started")
		user := c.Locals("User").(models.User)
		log.Println("IsAdmin", user)
		if !user.IsAdmin {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

		return c.Next()
	}
}
