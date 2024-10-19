package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func LogRoutes(app fiber.Router) {
	logsApi := app.Group("/logs")

	logsApi.Post("/errors/images", PostImageLog)
}

func PostImageLog(c *fiber.Ctx) error {
	type ImageLog struct {
		Src string `json:"src"`
	}
	var imageLog ImageLog
	err := c.BodyParser(&imageLog)
	if err != nil {
		log.Error().Err(err).Msg("Error trying to parse image log body")
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Unable to parse src from body", "error": err})
	}

	log.Error().Str("src", imageLog.Src).Msg("Client Error")
	return c.JSON(fiber.Map{"status": "success"})
}
