package controllers

import (
	"log"
	"switches/models"
	"switches/templates/pages"

	"github.com/gofiber/fiber/v2"
)

func GetHome(c *fiber.Ctx) error {
	user := c.Locals("User").(models.User)

	log.Println("GetHome", user)
	return Render(pages.HomePage(user), pages.Home(user))(c)
}

func GetSwitches(c *fiber.Ctx) error {
	log.Println("GetSwitches")
	return Render(pages.SwitchesPage(), pages.Switches())(c)
}
