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
	Avaliable        string    `form:"available"`
	PricePoint       int       `form:"pricePoint"`
	SiteURL          string    `form:"site-url"`
	LinkCount        int       `form:"link-count"`
	// Details
	DetailsID           uuid.UUID `form:"details-id"`
	SpringForce         float32   `form:"spring-force"`
	TactilityType       int       `form:"tactility-type"`
	TactilityFeedback   int       `form:"tactility-feedback"`
	FactoryLubed        string    `form:"factory-lubed"`
	PreTravel           float32   `form:"pre-travel"`
	TotalTravel         float32   `form:"total-travel"`
	InitialForce        float32   `form:"initial-force"`
	ActuationForce      float32   `form:"actuation-force"`
	ActuationPoint      float32   `form:"actuation-point"`
	ResetPoint          float32   `form:"reset-point"`
	BottomOutPoint      float32   `form:"bottom-out-point"`
	BottomOutForce      float32   `form:"bottom-out-force"`
	PinConfiguration    int       `form:"pin-configuration"`
	SpringMaterialType  int       `form:"spring-material-type"`
	SoundLevel          int       `form:"sound-level"`
	SoundType           int       `form:"sound-type"`
	TopMaterialHousing  int       `form:"top-material-housing"`
	BaseMaterialHousing int       `form:"base-material-housing"`
	StemMaterial        int       `form:"stem-material"`
	TopHousingColor     int       `form:"top-housing-color"`
	BottomHousingColor  int       `form:"bottom-housing-color"`
	StemColor           int       `form:"stem-color"`
	ShineThrough        string    `form:"shine-through"`
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

	isFactoryLubed := request.FactoryLubed == "on"
	hasShineThrough := request.ShineThrough == "on"
	details := models.Details{
		ID:              request.DetailsID,
		FactoryLubed:    &isFactoryLubed,
		HasShineThrough: &hasShineThrough,
	}

	if request.SpringMaterialType != 0 {
		details.SpringMaterialTypeID = &request.SpringMaterialType
	}
	if request.TopMaterialHousing != 0 {
		details.TopHousingMaterialID = &request.TopMaterialHousing
	}
	if request.BaseMaterialHousing != 0 {
		details.BaseHousingMaterialID = &request.BaseMaterialHousing
	}
	if request.StemMaterial != 0 {
		details.StemMaterialID = &request.StemMaterial
	}
	if request.TopHousingColor != 0 {
		details.TopHousingColorID = &request.TopHousingColor
	}
	if request.BottomHousingColor != 0 {
		details.BottomHousingColorID = &request.BottomHousingColor
	}
	if request.StemColor != 0 {
		details.StemColorID = &request.StemColor
	}
	if request.SpringForce != 0 {
		details.SpringForce = &request.SpringForce
	}
	if request.PreTravel != 0 {
		details.PreTravel = &request.PreTravel
	}
	if request.TotalTravel != 0 {
		details.TotalTravel = &request.TotalTravel
	}
	if request.InitialForce != 0 {
		details.InitialForce = &request.InitialForce
	}
	if request.ActuationForce != 0 {
		details.ActuationForce = &request.ActuationForce
	}
	if request.ActuationPoint != 0 {
		details.ActuationPoint = &request.ActuationPoint
	}
	if request.ResetPoint != 0 {
		details.ResetPoint = &request.ResetPoint
	}
	if request.BottomOutPoint != 0 {
		details.BottomOutForcePoint = &request.BottomOutPoint
	}
	if request.BottomOutForce != 0 {
		details.BottomOutForce = &request.BottomOutForce
	}
	if request.PinConfiguration != 0 {
		details.PinConfigurationID = &request.PinConfiguration
	}
	if request.SoundLevel != 0 {
		details.SoundLevelID = &request.SoundLevel
	}
	if request.TactilityType != 0 {
		details.TactilityTypeID = &request.TactilityType
	}
	if request.TactilityFeedback != 0 {
		details.TactilityFeedbackID = &request.TactilityFeedback
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
		Available:        request.Avaliable == "on",
		PricePoint:       request.PricePoint,
		SiteURL:          request.SiteURL,
		CreatedByID:      user.ID,
		UpdatedByID:      user.ID,
	}

	tx := database.DB.Begin()

	if err := tx.Debug().Model(&clickyClack).Updates(clickyClack).Update("available", request.Avaliable == "on").Error; err != nil {
		tx.Rollback()
		log.Error().Err(err).Msg("Error updating the switch")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error updating the switch", "error": err,
		})
	}
	timer.LogTime("Update Switch")

	if err := tx.Debug().Model(&details).Updates(details).Error; err != nil {
		log.Error().Err(err).Msg("Error updating the switch details")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error updating the switch details", "error": err,
		})
	}

	var currentImageLinkIDs []uuid.UUID
	if err := tx.Debug().Model(&models.ImageLink{}).
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
		materials,
		colors []components.Option
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
		case "color":
			colors = append(colors, components.Option{
				Value: strconv.Itoa(switchType.ID),
				Label: switchType.Name,
			})
		}
	}
	timer.LogTime("Types Generation")

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
		Colors:             colors,
	}
	timer.LogTime("Set Switch Data")

	return switchInput, nil
}
