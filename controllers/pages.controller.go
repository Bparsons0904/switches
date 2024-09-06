package controllers

import (
	"switches/templates/pages"

	"github.com/gofiber/fiber/v2"
)

func GetHomePage(c *fiber.Ctx) error {
	return Render(pages.Home())(c)
}

func NotFound(c *fiber.Ctx) error {
	return Render(pages.NotFound())(c)
}
