package controllers

import (
	"log"

	"switches/templates/pages"

	"github.com/gofiber/fiber/v2"
)

func GetHome(c *fiber.Ctx) error {
	log.Println("GetHome")
	return Render(pages.HomePage(), pages.Home())(c)
}

func GetSwitches(c *fiber.Ctx) error {
	log.Println("GetSwitches")
	return Render(pages.SwitchesPage(), pages.Switches())(c)
}
