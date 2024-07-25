package controllers

import (
	"switches/templates/pages"

	"github.com/gofiber/fiber/v2"
)

func GetHome(c *fiber.Ctx) error {
	return Render(pages.HomePage(), pages.Home())(c)
}
