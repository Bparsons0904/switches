package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func UID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		UID := c.Cookies("UID")

		if UID != "" {
			_, err := uuid.Parse(UID)
			if err != nil {
				log.Error().Err(err).Msg("Error parsing UID cookie")
				UID = ""
			}
		}

		if UID == "" {
			UID = uuid.NewString()
			expiredIn := time.Now().AddDate(1, 0, 0)
			cookie := fiber.Cookie{
				Name:     "UID",
				Value:    UID,
				Expires:  expiredIn,
				HTTPOnly: true,
				Secure:   true,
				SameSite: fiber.CookieSameSiteNoneMode,
			}
			c.Cookie(&cookie)
		}

		c.Locals("UID", UID)
		return c.Next()
	}
}
