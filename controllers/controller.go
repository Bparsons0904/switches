package controllers

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func RenderPage(fullPage, partialContent templ.Component) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var component templ.Component
		if c.Get("HX-Request") == "true" {
			component = partialContent
		} else {
			component = fullPage
		}

		componentHandler := templ.Handler(component)

		err := adaptor.HTTPHandler(componentHandler)(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
		return nil
	}
}

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
