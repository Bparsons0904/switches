package controllers

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/google/uuid"
)

func Render(component templ.Component) fiber.Handler {
	return func(c *fiber.Ctx) error {
		componentHandler := templ.Handler(component)

		err := adaptor.HTTPHandler(componentHandler)(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
		return nil
	}
}

func GetSwitchIDQuery(c *fiber.Ctx) (uuid.UUID, error) {
	switchID, err := uuid.Parse(c.Query("switchID"))
	if err != nil {
		return uuid.Nil, c.Status(fiber.StatusInternalServerError).Redirect("/404")
	}
	return switchID, nil
}

func GetSwitchIDParam(c *fiber.Ctx) (uuid.UUID, error) {
	switchID, err := uuid.Parse(c.Params("switchID"))
	if err != nil {
		return uuid.Nil, c.Status(fiber.StatusInternalServerError).Redirect("/404")
	}
	return switchID, nil
}
