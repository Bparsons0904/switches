package controllers

import "github.com/gofiber/fiber/v2"

func HandleUserOwnedSwitch(c *fiber.Ctx) error {
	return c.SendString("Create User Owned Switch")
}

func DeleteUserOwnedSwitch(c *fiber.Ctx) error {
	return c.SendString("Create User Owned Switch")
}

func HandleUserLikedSwitch(c *fiber.Ctx) error {
	return c.SendString("Create User Owned Switch")
}

func DeleteUserLikedSwitch(c *fiber.Ctx) error {
	return c.SendString("Create User Owned Switch")
}
