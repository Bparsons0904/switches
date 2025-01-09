package controllers

import (
	"fmt"
	"switches/database"
	"switches/models"
	"switches/templates/pages"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
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
	user := c.Locals("User").(models.User)

	existingSwitches := []uuid.UUID{}
	if user.ID != uuid.Nil {
		for _, favoriteSwitch := range user.LikedSwitches {
			existingSwitches = append(existingSwitches, favoriteSwitch.ID)
		}
		for _, ownedSwitch := range user.OwnedSwitches {
			existingSwitches = append(existingSwitches, ownedSwitch.ID)
		}
	}

	var selectedSwitch models.Switch
	if err := database.DB.Debug().
		Order("RANDOM()").
		Limit(1).
		Where("id NOT IN (?)", existingSwitches).
		Pluck("id", &selectedSwitch).Error; err != nil {
		log.Error().Err(err).Msg("Error finding a switch")
		return c.Status(fiber.StatusInternalServerError).Next()
	}

	selectedSwitchPath := fmt.Sprintf(
		"/switches/%s",
		selectedSwitch.ID,
	)
	return c.Redirect(selectedSwitchPath)
}
