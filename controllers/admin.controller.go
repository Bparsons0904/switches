package controllers

import (
	"fmt"
	"strconv"
	"switches/database"
	"switches/models"
	"switches/services/admin/seed"
	"switches/templates/components"
	"switches/templates/pages"
	"switches/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func PostSeedSwitches(c *fiber.Ctx) error {
	err := seed.SeedDatabase(database.DB)
	component := components.ToastNotification(components.ToastNotificationProps{
		Title: "Seeded Switches",
		Type:  "success",
	})
	if err != nil {
		log.Error().Err(err).Msg("Error seeing the database")
		component = components.ToastNotification(components.ToastNotificationProps{
			Title: "Failed to Seed Database",
			Type:  "error",
		})
	}

	return Render(component)(c)
}

func GetAdminHome(c *fiber.Ctx) error {
	user := c.Locals("User").(models.User)
	return Render(pages.Admin(user))(c)
}

func GetAdminSwitchEdit(c *fiber.Ctx) error { // {{{
	timer := utils.StartTimer("Get Admin Switch Edit")
	defer timer.LogTotalTime()

	switchID, err := GetSwitchIDParam(c)
	if err != nil {
		return err
	}

	switchInput, err := getSwitchInputData()
	if err != nil {
		log.Error().Err(err).Msg("Error getting the switch input data for edit")
		return c.Status(fiber.StatusBadRequest).Next()
	}

	var clickyClack models.Switch
	err = database.DB.
		Preload("ImageLinks").
		Preload("Brand").
		Preload("SwitchType").
		Preload("Details").
		First(&clickyClack, switchID).Error
	if err != nil {
		log.Error().Err(err).Msg("Error getting the switche")
		return c.Status(fiber.StatusBadRequest).Next()
	}

	input := components.SwitchInput{
		FormData: switchInput,
		Switch:   clickyClack,
	}

	return Render(pages.SwitchEdit(input))(c)
} // }}}

type SwitchQueryParams struct {
	ID               uuid.UUID `form:"switch-id"`
	Name             string    `form:"switch-name"`
	ShortDescription string    `form:"short-desc"`
	LongDescription  string    `form:"long-desc"`
	SwitchTypeID     int       `form:"switch-type-id"`
	BrandID          *int      `form:"brand-id"`
	ManufacturerID   *int      `form:"manufacturer-id"`
	ReleaseDate      *string   `form:"release-date"`
	Avaliable        string    `form:"avaliable"`
	PricePoint       int       `form:"pricePoint"`
	SiteURL          string    `form:"site-url"`
	LinkCount        int       `form:"link-count"`
}

func PutSwitch(c *fiber.Ctx) error { // {{{
	timer := utils.StartTimer("Put Switches")
	defer timer.LogTotalTime()

	user := c.Locals("User").(models.User)

	var request SwitchQueryParams
	if err := c.BodyParser(&request); err != nil {
		log.Error().Err(err).Msg("Error parsing request")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid link count",
		})
	}

	releaseDate, err := time.Parse("2006-01-02", *request.ReleaseDate)
	if err != nil {
		log.Error().Err(err).Msg("Error parsing release date")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing release date", "error": err,
		})
	}

	clickyClack := models.Switch{
		ID:               request.ID,
		Name:             request.Name,
		ShortDescription: request.ShortDescription,
		LongDescription:  request.LongDescription,
		SwitchTypeID:     request.SwitchTypeID,
		BrandID:          request.BrandID,
		ManufacturerID:   request.ManufacturerID,
		ReleaseDate:      &releaseDate,
		Available:        request.Avaliable == "true",
		PricePoint:       request.PricePoint,
		SiteURL:          request.SiteURL,
		CreatedByID:      user.ID,
		UpdatedByID:      user.ID,
	}

	tx := database.DB.Begin()

	if err := tx.Model(&clickyClack).Updates(clickyClack).Error; err != nil {
		tx.Rollback()
		log.Error().Err(err).Msg("Error updating the switch")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error updating the switch", "error": err,
		})
	}
	timer.LogTime("Update Switch")

	var currentImageLinkIDs []uuid.UUID
	if err := tx.Model(&models.ImageLink{}).
		Where("owner_id = ? AND owner_type = 'switches'", clickyClack.ID).
		Pluck("id", &currentImageLinkIDs).Error; err != nil {
		tx.Rollback()
		log.Error().Err(err).Msg("Error getting the current image links")
	}
	timer.LogTime("Get Current Image Links")

	var includedImageLinkIDS []uuid.UUID
	for i := 0; i < request.LinkCount; i++ {
		id := c.FormValue(fmt.Sprintf("link-id-%d", i))
		linkAddress := c.FormValue(fmt.Sprintf("link-address-%d", i))
		altText := c.FormValue(fmt.Sprintf("link-alt-text-%d", i))
		if id == uuid.Nil.String() {
			if err := tx.Create(&models.ImageLink{
				LinkAddress: linkAddress,
				AltText:     altText,
				OwnerID:     clickyClack.ID,
				OwnerType:   "switches",
				CreatedByID: user.ID,
				UpdatedByID: user.ID,
				IsPrimary:   i == 0,
			}).Error; err != nil {
				tx.Rollback()
				log.Error().Err(err).Msg("Error creating the image link")
			}
			timer.LogTime("Create New Image Link")
		} else {
			includedImageLinkIDS = append(includedImageLinkIDS, uuid.MustParse(id))
			if err := tx.
				Model(&models.ImageLink{}).
				Where("id = ?", id).
				Updates(models.ImageLink{
					LinkAddress: linkAddress,
					AltText:     altText,
					UpdatedByID: user.ID,
					IsPrimary:   i == 0,
				}).Error; err != nil {
				tx.Rollback()
				log.Error().Err(err).Msg("Error updating the image link")
			}
			timer.LogTime(fmt.Sprintf("Update Image Link %s", id))
		}
	}

	for _, currentID := range currentImageLinkIDs {
		var found bool
		for _, includedImageLinkID := range includedImageLinkIDS {
			if currentID == includedImageLinkID {
				found = true
				continue
			}
		}
		if !found {
			log.Info().Str("ID", currentID.String()).Msg("Deleting the image link")
			if err := tx.Where("id = ?", currentID).Delete(&models.ImageLink{}).Error; err != nil {
				tx.Rollback()
				log.Error().Err(err).Msg("Error deleting the image link")
			}
		}
	}

	tx.Commit()

	return GetAdminSwitchEdit(c)
}

func PostAdminSwitchCreate(c *fiber.Ctx) error {
	timer := utils.StartTimer("Create Switch")
	defer timer.LogTotalTime()

	user := c.Locals("User").(models.User)

	var request SwitchQueryParams
	if err := c.BodyParser(&request); err != nil {
		log.Error().Err(err).Msg("Error parsing request")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid link count",
		})
	}

	releaseDate, err := time.Parse("2006-01-02", *request.ReleaseDate)
	if err != nil {
		log.Error().Err(err).Msg("Error parsing release date")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing release date", "error": err,
		})
	}

	var imageLinks []models.ImageLink
	for i := 0; i < request.LinkCount; i++ {
		linkAddress := c.FormValue(fmt.Sprintf("link-address-%d", i))
		altText := c.FormValue(fmt.Sprintf("link-alt-text-%d", i))
		if linkAddress == "" || altText == "" {
			log.Warn().Msg("Empty link address or alt text")
			continue
		}

		imageLinks = append(imageLinks, models.ImageLink{
			LinkAddress: linkAddress,
			AltText:     altText,
			CreatedByID: user.ID,
			UpdatedByID: user.ID,
			IsPrimary:   i == 0,
			OwnerType:   "switch",
		})
	}

	releaseDate, err = time.Parse("2006-01-02", *request.ReleaseDate)
	if err != nil {
		log.Error().Err(err).Msg("Error parsing release date")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing release date", "error": err,
		})
	}

	clickyClack := models.Switch{
		Name:             request.Name,
		ShortDescription: request.ShortDescription,
		LongDescription:  request.LongDescription,
		SwitchTypeID:     request.SwitchTypeID,
		BrandID:          request.BrandID,
		ManufacturerID:   request.ManufacturerID,
		ReleaseDate:      &releaseDate,
		Available:        request.Avaliable == "true",
		PricePoint:       request.PricePoint,
		SiteURL:          request.SiteURL,
		CreatedByID:      user.ID,
		UpdatedByID:      user.ID,
		ImageLinks:       imageLinks,
	}

	timer.LogTime("Set data")
	if err := database.DB.Create(&clickyClack).Error; err != nil {
		log.Error().Err(err).Msg("Error creating the switch")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error creating the switch", "error": err,
		})
	}
	timer.LogTime("Create switch")

	if err := database.DB.
		Preload("ImageLinks").
		Preload("Brand").
		Preload("SwitchType").
		First(&clickyClack, clickyClack.ID).Error; err != nil {
		log.Error().Err(err).Msg("Error getting the switch")
		return c.Status(fiber.StatusBadRequest).Next()
	}
	timer.LogTime("Get Switch")
	return Render(pages.SwitchDetail(user, clickyClack))(c)
} // }}}

func DeleteImageLinkToList(c *fiber.Ctx) error {
	timer := utils.StartTimer("Delete Image Link to list")
	defer timer.LogTotalTime()

	imageLinkIndex, err := c.ParamsInt("imageLinkIndex")
	if err != nil {
		log.Error().Msg("No image link index provided")
		return c.Status(fiber.StatusBadRequest).Next()
	}

	linkCountStr := c.FormValue("link-count")
	linkCount, err := strconv.Atoi(linkCountStr)
	if err != nil {
		log.Error().Err(err).Msg("Error parsing link count")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid link count",
		})
	}

	var imageLinks []models.ImageLink
	for i := 0; i < linkCount; i++ {
		if i == imageLinkIndex {
			continue
		}

		idSting := c.FormValue(fmt.Sprintf("link-id-%d", i))
		id, err := uuid.Parse(idSting)
		linkAddress := c.FormValue(fmt.Sprintf("link-address-%d", i))
		altText := c.FormValue(fmt.Sprintf("link-alt-text-%d", i))

		imageLink := models.ImageLink{
			LinkAddress: linkAddress,
			AltText:     altText,
		}
		if err == nil {
			imageLink.ID = id
		}

		imageLinks = append(imageLinks, imageLink)
	}

	return Render(components.ImageLinksForm(imageLinks))(c)
}

func GetImageLinkToList(c *fiber.Ctx) error {
	timer := utils.StartTimer("Get Image Link to list")
	defer timer.LogTotalTime()

	linkCountStr := c.FormValue("link-count")
	linkCount, err := strconv.Atoi(linkCountStr)
	if err != nil {
		log.Error().Err(err).Msg("Error parsing link count")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid link count",
		})
	}

	var imageLinks []models.ImageLink
	for i := 0; i <= linkCount; i++ {
		idSting := c.FormValue(fmt.Sprintf("link-id-%d", i))
		id, err := uuid.Parse(idSting)
		linkAddress := c.FormValue(fmt.Sprintf("link-address-%d", i))
		altText := c.FormValue(fmt.Sprintf("link-alt-text-%d", i))

		imageLink := models.ImageLink{
			LinkAddress: linkAddress,
			AltText:     altText,
		}
		if err == nil {
			imageLink.ID = id
		}

		imageLinks = append(imageLinks, imageLink)
	}

	return Render(components.ImageLinksForm(imageLinks))(c)
}

type SwitchSearchQueryParams struct {
	Search string `json:"search"`
}

func GetAdminSwitches(c *fiber.Ctx) error {
	timer := utils.StartTimer("Get Admin Switches")
	defer timer.LogTotalTime()

	var request SwitchSearchQueryParams
	if err := c.QueryParser(&request); err != nil {
		log.Warn().Err(err).Msg("Error parsing query params")
	}

	var clickyClacks []models.Switch
	clickyClackQuery := database.DB.
		Preload("ImageLinks").
		Preload("Brand").
		Preload("SwitchType")

	if len(request.Search) > 2 {
		clickyClackQuery.Where("LOWER(name) LIKE LOWER(?)", fmt.Sprintf("%%%s%%", request.Search))
	}

	if err := clickyClackQuery.
		Find(&clickyClacks).Error; err != nil {
		log.Error().Err(err).Msg("Error getting the switches")
		return c.Status(fiber.StatusBadRequest).Next()
	}
	return Render(components.AdminSwitches(clickyClacks, request.Search))(c)
}

func GetAdminSwitchCreate(c *fiber.Ctx) error {
	timer := utils.StartTimer("Get Admin Switch Create")
	defer timer.LogTotalTime()

	switchData, err := getSwitchInputData()
	if err != nil {
		log.Error().Err(err).Msg("Error getting the switch input data for create")
		return c.Status(fiber.StatusBadRequest).Next()
	}

	switchInput := components.SwitchInput{
		FormData: switchData,
		Switch:   models.Switch{},
	}

	return Render(pages.SwitchCreate(switchInput))(c)
}

func getSwitchInputData() (components.SwitchData, error) {
	timer := utils.StartTimer("Get Admin Switch Create")
	defer timer.LogTotalTime()

	switchInput := components.SwitchData{}
	var brands, manufacturers []models.Producer
	if err := database.DB.Find(&brands).Error; err != nil {
		log.Error().Err(err).Msg("Error getting the brands")
		return switchInput, err
	}
	timer.LogTime("Get Brands")

	if err := database.DB.Find(&manufacturers).Error; err != nil {
		log.Error().Err(err).Msg("Error getting the manufacturers")
		return switchInput, err
	}
	timer.LogTime("Get Manufacturers")

	// var switchTypes []models.Type
	// if err := database.DB.Where("category = 'switch_type'").Find(&switchTypes).Error; err != nil {
	// 	log.Error().Err(err).Msg("Error getting the switch types")
	// 	return switchInput, err
	// }
	// timer.LogTime("Get Switch Types")

	var availableTypes []models.Type
	if err := database.DB.Find(&availableTypes).Error; err != nil {
		log.Error().Err(err).Msg("Error getting the switch types")
		return switchInput, err
	}
	timer.LogTime("Get Switch Types")

	var brandOptions []components.Option
	for _, brand := range brands {
		brandOptions = append(brandOptions, components.Option{
			Value: strconv.Itoa(brand.ID),
			Label: brand.Name,
		})
	}

	var manufacturerOptions []components.Option
	for _, manufacturer := range manufacturers {
		manufacturerOptions = append(manufacturerOptions, components.Option{
			Value: strconv.Itoa(manufacturer.ID),
			Label: manufacturer.Name,
		})
	}

	var switchOptions,
		tactilityTypes,
		tactilityFeedbacks,
		pinConfigurations,
		springMaterials,
		soundLevels,
		soundTypes,
		materials []components.Option
	for _, switchType := range availableTypes {
		switch switchType.Category {
		case "switch_type":
			switchOptions = append(switchOptions, components.Option{
				Value: strconv.Itoa(switchType.ID),
				Label: switchType.Name,
			})
		case "tactility_type":
			tactilityTypes = append(tactilityTypes, components.Option{
				Value: strconv.Itoa(switchType.ID),
				Label: switchType.Name,
			})
		case "tactility_feedback":
			tactilityFeedbacks = append(tactilityFeedbacks, components.Option{
				Value: strconv.Itoa(switchType.ID),
				Label: switchType.Name,
			})
		case "pin_configuration":
			pinConfigurations = append(pinConfigurations, components.Option{
				Value: strconv.Itoa(switchType.ID),
				Label: switchType.Name,
			})
		case "spring_material":
			springMaterials = append(springMaterials, components.Option{
				Value: strconv.Itoa(switchType.ID),
				Label: switchType.Name,
			})
		case "sound_level":
			soundLevels = append(soundLevels, components.Option{
				Value: strconv.Itoa(switchType.ID),
				Label: switchType.Name,
			})
		case "sound_type":
			soundTypes = append(soundTypes, components.Option{
				Value: strconv.Itoa(switchType.ID),
				Label: switchType.Name,
			})
		case "material":
			materials = append(materials, components.Option{
				Value: strconv.Itoa(switchType.ID),
				Label: switchType.Name,
			})
		}
	}

	switchInput = components.SwitchData{
		Brands:             brandOptions,
		Manufacturers:      manufacturerOptions,
		SwitchTypes:        switchOptions,
		TactilityTypes:     tactilityTypes,
		TactilityFeedbacks: tactilityFeedbacks,
		PinConfigurations:  pinConfigurations,
		SpringMaterials:    springMaterials,
		SoundLevels:        soundLevels,
		SoundTypes:         soundTypes,
		Materials:          materials,
	}
	timer.LogTime("Set Switch Data")

	return switchInput, nil
}
