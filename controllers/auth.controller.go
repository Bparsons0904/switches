package controllers

import (
	"log"
	"switches/templates/pages"

	"github.com/gofiber/fiber/v2"
)

func GetAuthCallback(c *fiber.Ctx) error {
	log.Println("Get auth callback")
	return Render(pages.SwitchesPage(), pages.Switches())(c)
}
