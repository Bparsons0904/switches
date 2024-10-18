package controllers

import (
	"fmt"
	"regexp"
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

func GetReviewForm(c *fiber.Ctx) error { // {{{
	switchID := c.Query("switch-id")
	ratingID := c.Query("rating-id")

	postPath := fmt.Sprintf("/switches/%s/ratings/%s/review", switchID, ratingID)
	return Render(components.UserReviewForm(postPath))(c)
} // }}}

func getDetailPageData(c *fiber.Ctx) (models.User, models.Switch, error) { // {{{
	user := c.Locals("User").(models.User)
	switchID, err := GetSwitchIDParam(c)
	if err != nil {
		return models.User{}, models.Switch{}, err
	}

	timer := utils.StartTimer("Get Switch Details")
	var clickyClack models.Switch
	if err := database.DB.
		Preload("ImageLinks").
		Preload("Brand").
		Preload("SwitchType").
		Preload("Details.SpringMaterialType").
		Preload("Details.TopHousingMaterial").
		Preload("Details.BaseHousingMaterial").
		Preload("Details.StemMaterial").
		Preload("Details.SoundLevel").
		Preload("Details.SoundType").
		Preload("Details.TactilityType").
		Preload("Details.TactilityFeedback").
		Preload("Details.StemColor").
		Preload("Details.TopHousingColor").
		Preload("Details.BottomHousingColor").
		Preload("Details.PinConfiguration").
		Preload("Ratings", func(db *gorm.DB) *gorm.DB {
			return db.Where("admin_review_required = false").Preload("User")
		}).
		First(&clickyClack, switchID).Error; err != nil {
		log.Error().Err(err).Msg("Error getting the switch")
		return models.User{}, models.Switch{}, err
	}
	clickyClack.GetUserRating(user.ID)
	timer.LogTotalTime()

	return user, clickyClack, nil
} // }}}

func PostUserSwitchReview(c *fiber.Ctx) error { // {{{
	ratingID := c.Params("ratingID")

	var userRating models.Rating
	if err := database.DB.
		Model(&userRating).
		Clauses(clause.Returning{}).
		Where("id = ?", ratingID).
		Updates(models.Rating{
			Review:         c.FormValue("user-review"),
			OriginalReview: c.FormValue("user-review"),
		}).Error; err != nil {
		log.Error().Err(err).Msg("Error getting the user rating")
		return c.Status(fiber.StatusBadRequest).Next()
	}

	user, clickyClack, err := getDetailPageData(c)
	if err != nil {
		log.Error().Err(err).Msg("Error getting the detail page data")
		return c.Status(fiber.StatusBadRequest).Next()
	}

	return Render(pages.SwitchDetail(user, clickyClack))(c)
} // }}}

func PutUserSwitch(c *fiber.Ctx) error { // {{{
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

	if err := models.UpdateSwitchRating(switchID, tx); err != nil {
		tx.Rollback()
		log.Error().Err(err).Msg("Error updating the switch rating on update")
		return c.Status(fiber.StatusBadRequest).Next()
	}

	tx.Commit()

	var clickyClack models.Switch
	if err := database.DB.
		Preload("Ratings.User").
		First(&clickyClack, switchID).Error; err != nil {
		log.Error().Err(err).Msg("Error getting the switch")
		return c.Status(fiber.StatusBadRequest).Next()
	}
	clickyClack.GetUserRating(userID)

	currentURL := c.Get("Hx-Current-Url")
	regex := `switches\/[a-f0-9]{8}-([a-f0-9]{4}-){3}[a-f0-9]{12}`
	matched, err := regexp.MatchString(regex, currentURL)
	if err != nil {
		log.Error().Err(err).Msg("Error trying to complete regex check on current url")
	}

	return Render(components.Ratings(clickyClack, matched))(c)
} // }}}

func GetSwitchPage(c *fiber.Ctx) error { // {{{
	timer := utils.StartTimer("Get Switch Page")
	defer timer.LogTotalTime()

	var clickyClacks []models.Switch
	err := database.DB.
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

	props := components.SwitchesFilterProps{
		ClickyClacks:         []models.Switch{},
		FilteredClickyClacks: clickyClacks,
		SwitchTypes:          switchTypes,
		SwitchBrands:         switchBrands,
		User:                 c.Locals("User").(models.User),
		Params: components.SwitchQueryParams{
			SwitchTypeIDs:   []int{},
			BrandIDs:        []int{},
			Pricepoints:     []int{},
			Search:          "",
			SwitchFavorites: false,
			SwitchOwned:     false,
		},
	}

	return Render(pages.Switches(props))(c)
} // }}}

func GetSwitchDetailPage(c *fiber.Ctx) error { // {{{
	timer := utils.StartTimer("getSwitchDetailPage")
	defer timer.LogTotalTime()

	user, clickyClack, err := getDetailPageData(c)
	if err != nil {
		log.Error().Err(err).Msg("Error getting the detail page data")
		return c.Status(fiber.StatusBadRequest).Next()
	}

	return Render(pages.SwitchDetail(user, clickyClack))(c)
} // }}}

func GetSwitchListMore(c *fiber.Ctx) error { // {{{
	timer := utils.StartTimer("Get More Switch List")
	defer timer.LogTotalTime()

	user := c.Locals("User").(models.User)

	_, filteredQuery, request, err := getListQuery(c, user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).Next()
	}

	var filteredClickyClacks []models.Switch
	err = filteredQuery.Offset(20).Find(&filteredClickyClacks).Error
	if err != nil {
		log.Error().Err(err).Msg("Error getting the switches")
		return c.Status(fiber.StatusBadRequest).Next()
	}

	for i, clickyClack := range filteredClickyClacks {
		clickyClack.GetUserRating(user.ID)
		filteredClickyClacks[i] = clickyClack
	}

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

	var switchBrands []models.Producer
	if err := database.DB.
		Find(&switchBrands).Error; err != nil {
		log.Error().Err(err).Msg("Error getting the switch brands")
		return c.Status(fiber.StatusBadRequest).Next()
	}

	props := components.SwitchesFilterProps{
		User:                 user,
		FilteredClickyClacks: filteredClickyClacks,
		ClickyClacks:         []models.Switch{},
		SwitchTypes:          switchTypes,
		SwitchBrands:         switchBrands,
		Params:               *request,
	}
	component := components.SwitchListMore(props)
	return Render(component)(c)
} // }}}

func getListQuery( // {{{
	c *fiber.Ctx,
	user models.User,
) (*gorm.DB, *gorm.DB, *components.SwitchQueryParams, error) {
	request := new(components.SwitchQueryParams)

	if err := c.QueryParser(request); err != nil {
		log.Error().Err(err).Msg("Error parsing query params")
		return &gorm.DB{}, &gorm.DB{}, request, err
	}

	clickyClackQuery := database.DB.
		Select("id, name, short_description, price_point, switch_type_id, brand_id")

	filteredClickyClackQuery := database.DB.
		Select(
			"id, name, short_description, price_point, average_rating, ratings_count, switch_type_id, brand_id",
		).
		Preload("ImageLinks", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, owner_id, owner_type, link_address, alt_text ")
		}).
		Preload("Brand", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name")
		}).
		Preload("SwitchType", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name")
		}).
		Preload("Ratings", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, review, rating, switch_id, admin_review_required, user_id")
		})

	if len(request.SwitchTypeIDs) > 0 {
		filteredClickyClackQuery.Where("switch_type_id IN (?)", request.SwitchTypeIDs)
	}

	if len(request.Search) > 2 {
		filteredClickyClackQuery.Where(
			"LOWER(name) LIKE LOWER(?)",
			fmt.Sprintf("%%%s%%", request.Search),
		)
		clickyClackQuery.Where(
			"LOWER(name) LIKE LOWER(?)",
			fmt.Sprintf("%%%s%%", request.Search),
		)
	}

	if len(request.BrandIDs) > 0 {
		filteredClickyClackQuery.Where("brand_id IN (?)", request.BrandIDs)
	}

	if len(request.Pricepoints) > 0 {
		filteredClickyClackQuery.Where("price_point IN (?)", request.Pricepoints)
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
				return &gorm.DB{}, &gorm.DB{}, request, err
			}
			idsToInclude = append(idsToInclude, userOwnedSwitches...)
		}

		if request.SwitchFavorites {
			var userLikedSwitches []uuid.UUID
			if err := database.DB.Model(&models.UserLikedSwitches{}).
				Where("user_id = ?", user.ID).
				Pluck("switch_id", &userLikedSwitches).Error; err != nil {
				log.Error().Err(err).Msg("Error getting the user liked switches")
				return &gorm.DB{}, &gorm.DB{}, request, err
			}
			idsToInclude = append(idsToInclude, userLikedSwitches...)
		}

		clickyClackQuery.Where("id IN (?)", idsToInclude)
		filteredClickyClackQuery.Where("id IN (?)", idsToInclude)
	}

	return clickyClackQuery, filteredClickyClackQuery, request, nil
} // }}}

func GetSwitchList(c *fiber.Ctx) error { // {{{
	timer := utils.StartTimer("Get Switch List")
	defer timer.LogTotalTime()
	user := c.Locals("User").(models.User)

	query, filteredQuery, request, err := getListQuery(c, user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).Next()
	}

	var clickyClacks []models.Switch
	err = query.Find(&clickyClacks).Error
	if err != nil {
		log.Error().Err(err).Msg("Error getting the switches")
		return c.Status(fiber.StatusBadRequest).Next()
	}

	var filteredClickyClacks []models.Switch
	err = filteredQuery.Limit(20).Find(&filteredClickyClacks).Error
	if err != nil {
		log.Error().Err(err).Msg("Error getting the filtered switches")
		return c.Status(fiber.StatusBadRequest).Next()
	}

	for i, clickyClack := range filteredClickyClacks {
		clickyClack.GetUserRating(user.ID)
		filteredClickyClacks[i] = clickyClack
	}

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

	var switchBrands []models.Producer
	if err := database.DB.
		Find(&switchBrands).Error; err != nil {
		log.Error().Err(err).Msg("Error getting the switch brands")
		return c.Status(fiber.StatusBadRequest).Next()
	}

	props := components.SwitchesFilterProps{
		User:                 user,
		FilteredClickyClacks: filteredClickyClacks,
		ClickyClacks:         clickyClacks,
		SwitchTypes:          switchTypes,
		SwitchBrands:         switchBrands,
		Params:               *request,
	}

	component := components.SwitchList(props)
	return Render(component)(c)
} // }}}

func GetFeaturedSwitches(c *fiber.Ctx) error { // {{{
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

	for i, clickyClack := range clickyClacks {
		clickyClack.GetUserRating(userID)
		clickyClacks[i] = clickyClack
	}

	component := components.FeaturedSwitches(clickyClacks)
	return Render(component)(c)
} // }}}

func GetSwitchDetailCard(c *fiber.Ctx) error { // {{{
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
} // }}}

func getParams(c *fiber.Ctx) (models.User, uuid.UUID, error) { // {{{
	user := c.Locals("User").(models.User)
	switchID, err := GetSwitchIDParam(c)
	if err != nil {
		log.Error().Err(err).Msg("Error getting the switch id")
	}

	return user, switchID, err
} // }}}

func getSwitches(switchID uuid.UUID, timer *utils.Timer) (models.Switch, error) { // {{{
	var clickyClack models.Switch
	if err := database.DB.First(&clickyClack, switchID).Error; err != nil {
		return models.Switch{}, err
	}
	timer.LogTime("Get Switch")

	return clickyClack, nil
} // }}}

func getUpdatedUser( // {{{
	user models.User,
	timer *utils.Timer,
	c *fiber.Ctx,
) (models.User, error) {
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
} // }}}

func processUserSwitch( // {{{
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
} // }}}

func CreateUserOwnedSwitch(c *fiber.Ctx) error { // {{{
	timer := utils.StartTimer("Create User Owned Switch")
	defer timer.LogTotalTime()

	return processUserSwitch(
		c,
		true,
		&models.UserOwnedSwitches{},
		components.OwnedSelector,
		&timer,
	)
} // }}}

func DeleteUserOwnedSwitch(c *fiber.Ctx) error { // {{{
	timer := utils.StartTimer("Delete User Owned Switch")
	defer timer.LogTotalTime()

	return processUserSwitch(
		c,
		false,
		&models.UserOwnedSwitches{},
		components.OwnedSelector,
		&timer,
	)
} // }}}

func CreateUserLikedSwitch(c *fiber.Ctx) error { // {{{
	timer := utils.StartTimer("Create User Liked Switch")
	defer timer.LogTotalTime()

	return processUserSwitch(
		c,
		true,
		&models.UserLikedSwitches{},
		components.LikedSelector,
		&timer,
	)
} // }}}

func DeleteUserLikedSwitch(c *fiber.Ctx) error { // {{{
	timer := utils.StartTimer("Delete User Liked Switch")
	defer timer.LogTotalTime()

	return processUserSwitch(
		c,
		false,
		&models.UserLikedSwitches{},
		components.LikedSelector,
		&timer,
	)
} // }}}
