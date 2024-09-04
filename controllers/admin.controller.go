package controllers

import (
	"fmt"
	"switches/database"
	"switches/models"
	"switches/templates/components"
	"switches/templates/pages"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func GetAdminHome(c *fiber.Ctx) error {
	user := c.Locals("User").(models.User)
	return Render(pages.Admin(user))(c)
}

func GetAdminSwitchEdit(c *fiber.Ctx) error {
	user := c.Locals("User").(models.User)
	return Render(pages.SwitchEdit(user))(c)
}

type SwitchQueryParams struct {
	Search string `json:"search"`
}

func GetAdminSwitches(c *fiber.Ctx) error {
	var request SwitchQueryParams
	if err := c.QueryParser(&request); err != nil {
		log.Warn().Err(err).Msg("Error parsing query params")
	}

	var clickyClacks []models.Switch
	clickyClackQuery := database.DB.
		Preload("ImageLinks").
		Preload("Brand").
		Preload("SwitchType")

	if len(request.Search) > 2 {
		clickyClackQuery.Where("LOWER(name) LIKE LOWER(?)", fmt.Sprintf("%%%s%%", request.Search))
	}

	if err := clickyClackQuery.
		Find(&clickyClacks).Error; err != nil {
		log.Error().Err(err).Msg("Error getting the switches")
		return c.Status(fiber.StatusBadRequest).Next()
	}
	return Render(components.AdminSwitches(clickyClacks, request.Search))(c)
}

func GetAdminSwitchCreate(c *fiber.Ctx) error {
	return Render(pages.SwitchCreate())(c)
}
