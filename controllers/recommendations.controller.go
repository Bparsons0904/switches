package controllers

import (
	"fmt"
	"switches/database"
	"switches/models"
	"switches/templates/pages"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

const (
	guided = "GuidedRecommendation"
	week   = time.Second * 60
	// week   = time.Hour * 24 * 7
)

func GetAiRecommendations(c *fiber.Ctx) error {
	return c.SendString("AI Recommendations")
}

func GetGuidedRecommendations(c *fiber.Ctx) error {
	UID := c.Locals("UID").(string)
	User := c.Locals("User").(models.User)

	currentProgress, err := database.GetJSONKeyDB[pages.Progress](guided, UID)
	if err != nil {
		log.Err(err).Msg("Error getting progress")
	}

	if currentProgress.Step == 0 {
		currentProgress = pages.Progress{
			Step:      1,
			Questions: make(map[int]int),
		}
		err := database.SetJSONKeyDB(guided, UID, currentProgress, week)
		if err != nil {
			log.Error().Err(err).Msg("Error setting initial progress")
		}
	}

	return Render(pages.GuidedRecommendation(User, currentProgress))(c)
}

type ProgressBody struct {
	GotoStep int `json:"step"`
	Question int `json:"question"`
	Option   int `json:"value"`
}

func PostGuidedRecommendation(c *fiber.Ctx) error {
	User := c.Locals("User").(models.User)
	UID := c.Locals("UID").(string)

	var request ProgressBody
	err := c.BodyParser(&request)
	if err != nil {
		log.Err(err).Msg("Error parsing request")
	}

	currentProgress, err := database.GetJSONKeyDB[pages.Progress](guided, UID)
	if err != nil {
		log.Err(err).Msg("Error getting progress")
	}

	currentProgress.Step = request.Question + 1
	currentProgress.Questions[request.Question] = request.Option
	if err := database.SetJSONKeyDB(guided, UID, currentProgress, week); err != nil {
		log.Err(err).Msg("Error setting progress")
		return c.Status(fiber.StatusInternalServerError).Next()
	}

	return Render(pages.GuidedRecommendation(User, currentProgress))(c)
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
