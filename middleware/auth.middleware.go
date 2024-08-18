package middleware

import (
	"log"
	"switches/config"
	"switches/controllers"
	"switches/database"
	"switches/models"

	"github.com/gofiber/fiber/v2"
)

func AuthenticateUser(config config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sessionID := c.Cookies("sessionID")
		var user models.User
		if sessionID != "" {
			retrievedSession, err := database.GetJSONKeyDB[controllers.Session](
				"session",
				sessionID,
			)
			if err != nil {
				log.Println("Error getting session from keydb", err)
				c.ClearCookie("sessionID")
			}

			user, err = database.GetJSONKeyDB[models.User]("user", retrievedSession.UserID.String())
			if err != nil {
				if err := database.DB.First(&user, retrievedSession.UserID).Error; err != nil {
					log.Println("Error getting user from keydb", err)
				} else {
					err := database.SetJSONKeyDB("user", user.ID.String(), user)
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
