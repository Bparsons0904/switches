package controllers

import (
	"log/slog"
	"switches/database"
	"switches/models"
	"switches/templates/components"
	"switches/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetImage(c *fiber.Ctx) error {
	timer := utils.StartTimer("Get Image")
	defer timer.LogTotalTime()

	imageLinkID, err := uuid.Parse(c.Params("imageLinkID"))
	if err != nil {
		slog.Error("Error parsing imageLinkID", "error", err)
		return c.Status(fiber.StatusInternalServerError).Next()
	}

	var imageLink models.ImageLink
	if err := database.DB.First(&imageLink, imageLinkID).Error; err != nil {
		slog.Error("Error getting imageLink", "error", err)
	}

	component := components.ImageNew(&imageLink, true)
	return Render(component)(c)
}
