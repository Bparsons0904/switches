package controllers

import (
	"switches/models"
	"switches/templates/pages"

	"github.com/gofiber/fiber/v2"
)

func GetAiRecommendations(c *fiber.Ctx) error {
	return c.SendString("AI Recommendations")
}

func GetGuidedRecommendations(c *fiber.Ctx) error {
	user := c.Locals("User").(models.User)
	return Render(pages.GuidedRecommendation(user))(c)
}

func GetUserBasedRecommendations(c *fiber.Ctx) error {
	return c.SendString("User Based Recommendations")
}

func GetFeelingLuckyRecommendations(c *fiber.Ctx) error {
	return c.SendString("Feeling Lucky Recommendations")
}
