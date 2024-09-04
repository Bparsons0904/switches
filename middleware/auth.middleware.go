package middleware

import (
	"switches/config"
	"switches/database"
	"switches/models"
	"switches/services"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func AuthenticateUser(config config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Info().Msg("AuthenticateUser started")
		sessionID := c.Cookies("sessionID")
		var user models.User
		if sessionID != "" {
			retrievedSession, err := services.SessionFlow(sessionID)
			if err != nil {
				c.ClearCookie("sessionID")
				c.Locals("User", user)
				log.Error().Err(err).Msg("Error getting session from keydb")
				return c.Next()
			}

			user, err = database.GetUUIDJSONKeyDB[models.User]("user", retrievedSession.UserID)
			if err != nil {
				if err := database.DB.
					Preload("OwnedSwitches").
					Preload("LikedSwitches").
					First(&user, retrievedSession.UserID).Error; err != nil {
					log.Warn().Err(err).Msg("Unable to get the user in KeyDB")
				} else {
					err := database.SetUUIDJSONKeyDB("user", user.ID, user, 30*time.Hour)
					if err != nil {
						log.Error().Msg("Error setting user in keydb")
					}
				}
			}
		}

		c.Locals("User", user)
		c.Locals("UserID", user.ID)
		if user.ID != uuid.Nil {
			log.Info().Str("userID", user.ID.String()).Msg("User is authenticated")
		}
		return c.Next()
	}
}

func IsAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("User").(models.User)
		if !user.IsAdmin {
			log.Warn().Str("userID", user.ID.String()).Msg("Unauthorized attempted access")
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

		return c.Next()
	}
}
