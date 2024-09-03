package controllers

import "github.com/gofiber/fiber/v2"

func UserPlaceHolder(c *fiber.Ctx) error {
	return c.SendString("User Placeholder")
}
