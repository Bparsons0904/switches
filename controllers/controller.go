package controllers

import (
	"switches/models"
	"switches/templates/layouts"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func Render(selectedComponent templ.Component) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Get("HX-Request") != "true" {
			user := c.Locals("User").(models.User)
			selectedComponent = layouts.FullPage(user, selectedComponent)
		}

		componentHandler := templ.Handler(selectedComponent)
		err := adaptor.HTTPHandler(componentHandler)(c)
		if err != nil {
			log.Error().Err(err).Msg("Error rendering component")
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
		return nil
	}
}

func GetSwitchIDQuery(c *fiber.Ctx) (uuid.UUID, error) {
	switchID, err := uuid.Parse(c.Query("switchID"))
	if err != nil {
		log.Error().Err(err).Msg("Error parsing switchID")
		return uuid.Nil, c.Status(fiber.StatusInternalServerError).Redirect("/404")
	}
	return switchID, nil
}

func GetSwitchIDParam(c *fiber.Ctx) (uuid.UUID, error) {
	switchID, err := uuid.Parse(c.Params("switchID"))
	if err != nil {
		log.Error().Err(err).Msg("Error parsing switchID")
		return uuid.Nil, c.Status(fiber.StatusInternalServerError).Redirect("/404")
	}
	return switchID, nil
}
