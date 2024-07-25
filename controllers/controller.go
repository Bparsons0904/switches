package controllers

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func Render(fullPage, partialContent templ.Component) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if viper.GetString("TIER") == "development" {
			c.Set("Cache-Control", "no-store")
		}

		var component templ.Component
		if c.Get("HX-Request") == "true" {
			component = partialContent
		} else {
			component = fullPage
		}

		err := component.Render(c.Context(), c.Response().BodyWriter())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
		return nil
	}
}
