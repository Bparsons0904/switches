package controllers

import (
	"switches/database"
	"switches/models"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

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

	if err := database.DB.Delete(&models.ImageLink{}, "link_address = ?", imageLog.Src).Error; err != nil {
		log.Error().Err(err).Msg("Error deleting image link")
		return c.JSON(fiber.Map{"status": "error"})
	}

	log.Warn().Str("src", imageLog.Src).Msg("Client Error")
	return c.JSON(fiber.Map{"status": "success"})
}
