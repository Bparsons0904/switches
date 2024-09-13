package controllers

import (
	"fmt"
	"strings"
	"switches/database"
	"switches/models"
	"switches/templates/components"
	"switches/templates/pages"
	"switches/utils"
	"time"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetReviewForm(c *fiber.Ctx) error {
	switchID := c.Query("switch-id")
	ratingID := c.Query("rating-id")

	postPath := fmt.Sprintf("/switches/%s/ratings/%s/review", switchID, ratingID)
	return Render(components.UserReviewForm(postPath))(c)
}

func PostUserSwitchReview(c *fiber.Ctx) error {
	ratingID := c.Params("ratingID")

	var userRating models.Rating
	if err := database.DB.
		Model(&userRating).
		Clauses(clause.Returning{}).
		Where("id = ?", ratingID).
		Updates(models.Rating{
			Review: c.FormValue("user-review"),
		}).Error; err != nil {
		log.Error().Err(err).Msg("Error getting the user rating")
		return c.Status(fiber.StatusBadRequest).Next()
	}

	return Render(components.UserReview(userRating))(c)
}

func PutUserSwitch(c *fiber.Ctx) error {
	timer := utils.StartTimer("Put User Switch")
	defer timer.LogTotalTime()

	userID := c.Locals("UserID").(uuid.UUID)
	if userID == uuid.Nil {
		log.Warn().Msg("User not logged in")
		c.Set("HX-Redirect", "/auth/login")
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	switchID, err := uuid.Parse(c.Params("switchID"))
	if err != nil {
		log.Error().Err(err).Msg("Error parsing the switch id")
		return c.Status(fiber.StatusBadRequest).Next()
	}
	rating, err := c.ParamsInt("rating")
	if err != nil {
		log.Error().Err(err).Msg("Error parsing the rating")
		return c.Status(fiber.StatusBadRequest).Next()
	}

	tx := database.DB.Begin()
	var userRating models.Rating
	err = tx.
		Where(models.Rating{
			UserID:   userID,
			SwitchID: switchID,
		}).
		Attrs(models.Rating{
			Rating: rating + 1,
		}).
		FirstOrCreate(&userRating).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		tx.Rollback()
		log.Error().Err(err).Msg("Error creating the rating")
		return c.Status(fiber.StatusBadRequest).Next()
	} else {
		userRating.Rating = rating + 1
		if err := tx.Save(&userRating).Error; err != nil {
			tx.Rollback()
			log.Error().Err(err).Msg("Error saving the rating with new user value")
			return c.Status(fiber.StatusBadRequest).Next()
		}
	}

	var clickyClack models.Switch
	if err := tx.
		Preload("Ratings").
		First(&clickyClack, switchID).Error; err != nil {
		log.Error().Err(err).Msg("Error getting the switch")
		tx.Rollback()
		return c.Status(fiber.StatusBadRequest).Next()
	}
	clickyClack.GetUserRating(userID)

	totalRating := 0
	for _, rating := range clickyClack.Ratings {
		totalRating += rating.Rating
	}
	clickyClack.RatingsCount = len(clickyClack.Ratings)
	clickyClack.AverageRating = float64(totalRating) / float64(clickyClack.RatingsCount)
	if err := tx.Model(&models.Switch{}).Where("id = ?", clickyClack.ID).Updates(models.Switch{
		RatingsCount:  clickyClack.RatingsCount,
		AverageRating: clickyClack.AverageRating,
	}).Error; err != nil {
		log.Error().Err(err).Msg("Error saving the switch after ratings update")
		tx.Rollback()
		return c.Status(fiber.StatusBadRequest).Next()
	}

	tx.Commit()

	showReviewButton := true
	currentURL := c.Get("Hx-Current-Url")
	currentURL = currentURL[strings.LastIndex(currentURL, "/")+1:]
	if currentURL == "switches" || strings.Contains(currentURL, "modal") {
		showReviewButton = false
	}
	return Render(components.Ratings(clickyClack, showReviewButton))(c)
}

func GetSwitchPage(c *fiber.Ctx) error {
	timer := utils.StartTimer("Get Switch Page")
	defer timer.LogTotalTime()

	var clickyClacks []models.Switch
	err := database.DB.
		Preload("ImageLinks").
		Preload("Brand").
		Preload("SwitchType").
		Find(&clickyClacks).Error
	if err != nil {
		log.Error().Err(err).Msg("Error getting the switches")
		return c.Status(fiber.StatusBadRequest).Next()
	}
	timer.LogTime("Get Switches")

	var switchTypes []models.Type
	err = database.DB.
		Select("id", "name").
		Order("id").
		Where("category = ?", "switch_type").
		Find(&switchTypes).Error
	if err != nil {
		log.Error().Err(err).Msg("Error getting the switch types")
		return c.Status(fiber.StatusBadRequest).Next()
	}
	timer.LogTime("Get Switch Types")

	var switchBrands []models.Producer
	if err := database.DB.
		Find(&switchBrands).Error; err != nil {
		log.Error().Err(err).Msg("Error getting the switch brands")
		return c.Status(fiber.StatusBadRequest).Next()
	}
	timer.LogTime("Get Switch Brands")

	props := pages.SwitchesPageProps{
		ClickyClacks: clickyClacks,
		SwitchTypes:  switchTypes,
		SwitchBrands: switchBrands,
		User:         c.Locals("User").(models.User),
	}

	return Render(
		pages.Switches(props),
	)(
		c,
	)
}

func GetSwitchDetailPage(c *fiber.Ctx) error {
	timer := utils.StartTimer("getSwitchDetailPage")
	defer timer.LogTotalTime()

	userID := c.Locals("UserID").(uuid.UUID)
	switchID, err := GetSwitchIDParam(c)
	if err != nil {
		return err
	}

	var user models.User
	if userID != uuid.Nil {
		if err := database.DB.
			Preload("OwnedSwitches").
			Preload("LikedSwitches").
			First(&user, userID).Error; err != nil {
			return c.Status(fiber.StatusBadRequest).Next()
		}
	}
	timer.LogTime("Get User")

	var clickyClack models.Switch
	if err := database.DB.
		Preload("ImageLinks").
		Preload("Brand").
		Preload("SwitchType").
		Preload("Ratings").
		First(&clickyClack, switchID).Error; err != nil {
		log.Error().Err(err).Msg("Error getting the switch")
		return c.Status(fiber.StatusBadRequest).Next()
	}
	clickyClack.GetUserRating(userID)
	timer.LogTime("Get Switch")

	return Render(pages.SwitchDetail(user, clickyClack))(c)
}

func GetSwitchList(c *fiber.Ctx) error {
	timer := utils.StartTimer("Get Switch List")
	defer timer.LogTotalTime()

	user := c.Locals("User").(models.User)

	type SwitchQueryParams struct {
		SwitchTypeIDs   []int  `json:"switchTypeIDs"`
		BrandIDs        []int  `json:"brandIDs"`
		Pricepoints     []int  `json:"pricepoints"`
		Search          string `json:"search"`
		SwitchFavorites bool   `json:"switchFavorites"`
		SwitchOwned     bool   `json:"switchOwned"`
	}
	request := new(SwitchQueryParams)

	if err := c.QueryParser(request); err != nil {
		log.Warn().Err(err).Msg("Error parsing query params")
	}

	var clickyClacks []models.Switch
	clickyClackQuery := database.DB.
		Preload("ImageLinks").
		Preload("Brand").
		Preload("SwitchType").
		Preload("Ratings")

	if len(request.SwitchTypeIDs) > 0 {
		clickyClackQuery.Where("switch_type_id IN (?)", request.SwitchTypeIDs)
	}

	if len(request.Search) > 2 {
		clickyClackQuery.Where("LOWER(name) LIKE LOWER(?)", fmt.Sprintf("%%%s%%", request.Search))
	}

	if len(request.BrandIDs) > 0 {
		clickyClackQuery.Where("brand_id IN (?)", request.BrandIDs)
	}

	if len(request.Pricepoints) > 0 {
		clickyClackQuery.Where("price_point IN (?)", request.Pricepoints)
	}

	if request.SwitchFavorites || request.SwitchOwned {
		var idsToInclude []uuid.UUID
		if request.SwitchOwned {
			var userOwnedSwitches []uuid.UUID
			if err := database.DB.Model(&models.UserOwnedSwitches{}).
				Where("user_id = ?", user.ID).
				Pluck("switch_id", &userOwnedSwitches).Error; err != nil {
				log.Error().Err(err).Msg("Error getting the user owned switches")
				return c.Status(fiber.StatusBadRequest).Next()
			}
			idsToInclude = append(idsToInclude, userOwnedSwitches...)
		}

		if request.SwitchFavorites {
			var userLikedSwitches []uuid.UUID
			if err := database.DB.Model(&models.UserLikedSwitches{}).
				Where("user_id = ?", user.ID).
				Pluck("switch_id", &userLikedSwitches).Error; err != nil {
				log.Error().Err(err).Msg("Error getting the user liked switches")
				return c.Status(fiber.StatusBadRequest).Next()
			}
			idsToInclude = append(idsToInclude, userLikedSwitches...)
		}

		clickyClackQuery.Where("id IN (?)", idsToInclude)
	}

	err := clickyClackQuery.Find(&clickyClacks).Error
	if err != nil {
		log.Error().Err(err).Msg("Error getting the switches")
		return c.Status(fiber.StatusBadRequest).Next()
	}

	if user.ID != uuid.Nil {
		for i, clickyClack := range clickyClacks {
			clickyClack.GetUserRating(user.ID)
			clickyClacks[i] = clickyClack
		}
	}
	timer.LogTime("Get Switches")

	component := components.SwitchList(user, clickyClacks)
	return Render(component)(c)
}

func GetFeaturedSwitches(c *fiber.Ctx) error {
	timer := utils.StartTimer("Get Featured Switches")
	defer timer.LogTotalTime()
	userID := c.Locals("UserID").(uuid.UUID)

	var clickyClacks []models.Switch
	database.DB.
		Joins("INNER JOIN image_links ON image_links.owner_id = switches.id").
		Preload("ImageLinks").
		Preload("SwitchType").
		Preload("Ratings").
		Limit(4).
		Order("RANDOM()").
		Find(&clickyClacks)
	timer.LogTime("Get Switches")

	if userID != uuid.Nil {
		for i, clickyClack := range clickyClacks {
			log.Info().Msgf("Getting user rating for %s", clickyClack.Name)
			clickyClack.GetUserRating(userID)
			clickyClacks[i] = clickyClack
		}
	}

	component := components.FeaturedSwitches(clickyClacks)
	return Render(component)(c)
}

func GetSwitchDetailCard(c *fiber.Ctx) error {
	timer := utils.StartTimer("getSwitchModal")
	defer timer.LogTotalTime()

	user, switchID, err := getParams(c)
	if err != nil {
		log.Error().Err(err).Msg("Error getting the params")
		return err
	}

	var clickyClack models.Switch
	database.DB.
		Preload("ImageLinks").
		Preload("Brand").
		Preload("SwitchType").
		Preload("Ratings").
		First(&clickyClack, switchID)

	clickyClack.GetUserRating(user.ID)
	component := components.SwitchDetailCard(user, clickyClack)
	return Render(component)(c)
}

func getParams(c *fiber.Ctx) (models.User, uuid.UUID, error) {
	user := c.Locals("User").(models.User)
	switchID, err := GetSwitchIDParam(c)
	if err != nil {
		log.Error().Err(err).Msg("Error getting the switch id")
	}

	return user, switchID, err
}

func getSwitches(switchID uuid.UUID, timer *utils.Timer) (models.Switch, error) {
	var clickyClack models.Switch
	if err := database.DB.First(&clickyClack, switchID).Error; err != nil {
		return models.Switch{}, err
	}
	timer.LogTime("Get Switch")

	return clickyClack, nil
}

func getUpdatedUser(user models.User, timer *utils.Timer, c *fiber.Ctx) (models.User, error) {
	var updatedUser models.User
	if err := database.DB.
		Preload("OwnedSwitches").
		Preload("LikedSwitches").
		First(&updatedUser, user.ID).Error; err != nil {
		return models.User{}, err
	}
	timer.LogTime("Get User")

	c.Locals("User", updatedUser)
	if err := database.SetUUIDJSONKeyDB("user", updatedUser.ID, updatedUser, 30*time.Hour); err != nil {
		log.Error().Err(err).Msg("Error setting user in keydb")
		if err := database.DeleteUUIDKeyDB("user", updatedUser.ID); err != nil {
			log.Error().Err(err).Msg("Error deleting user in keydb")
		}
	}
	timer.LogTime("Set User in KeyDB")

	return updatedUser, nil
}

func processUserSwitch(
	c *fiber.Ctx,
	create bool,
	table interface{},
	selectorFunc func(models.User, models.Switch) templ.Component,
	timer *utils.Timer,
) error {
	user, switchID, err := getParams(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Next()
	}

	if user.ID == uuid.Nil {
		log.Warn().Msg("User not logged in")
		c.Set("HX-Redirect", "/auth/login")
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if create {
		if err := database.DB.Model(table).Create(map[string]interface{}{
			"UserID": user.ID, "SwitchID": switchID,
		}).Error; err != nil {
			log.Error().Err(err).Msg("Error creating the association")
			return c.Status(fiber.StatusInternalServerError).Next()
		}
		timer.LogTime("Create Association")
	} else {
		if err := database.DB.Delete(table, "user_id = ? AND switch_id = ?", user.ID, switchID).Error; err != nil {
			log.Error().Err(err).Msg("Error deleting the association")
			return c.Status(fiber.StatusInternalServerError).Next()
		}
		timer.LogTime("Delete Association")
	}

	clickyClack, err := getSwitches(switchID, timer)
	if err != nil {
		log.Error().Err(err).Msg("Error getting the switch")
		return c.Status(fiber.StatusInternalServerError).Next()
	}

	updatedUser, err := getUpdatedUser(user, timer, c)
	if err != nil {
		log.Error().Err(err).Msg("Error getting the updated user")
		return c.Status(fiber.StatusInternalServerError).Next()
	}

	component := selectorFunc(updatedUser, clickyClack)
	return Render(component)(c)
}

func CreateUserOwnedSwitch(c *fiber.Ctx) error {
	timer := utils.StartTimer("Create User Owned Switch")
	defer timer.LogTotalTime()

	return processUserSwitch(
		c,
		true,
		&models.UserOwnedSwitches{},
		components.OwnedSelector,
		&timer,
	)
}

func DeleteUserOwnedSwitch(c *fiber.Ctx) error {
	timer := utils.StartTimer("Delete User Owned Switch")
	defer timer.LogTotalTime()

	return processUserSwitch(
		c,
		false,
		&models.UserOwnedSwitches{},
		components.OwnedSelector,
		&timer,
	)
}

func CreateUserLikedSwitch(c *fiber.Ctx) error {
	timer := utils.StartTimer("Create User Liked Switch")
	defer timer.LogTotalTime()

	return processUserSwitch(
		c,
		true,
		&models.UserLikedSwitches{},
		components.LikedSelector,
		&timer,
	)
}

func DeleteUserLikedSwitch(c *fiber.Ctx) error {
	timer := utils.StartTimer("Delete User Liked Switch")
	defer timer.LogTotalTime()

	return processUserSwitch(
		c,
		false,
		&models.UserLikedSwitches{},
		components.LikedSelector,
		&timer,
	)
}
