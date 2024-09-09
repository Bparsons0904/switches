package controllers

import (
	"fmt"
	"strconv"
	"switches/database"
	"switches/models"
	"switches/templates/components"
	"switches/templates/pages"
	"switches/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func GetAdminHome(c *fiber.Ctx) error {
	user := c.Locals("User").(models.User)
	return Render(pages.Admin(user))(c)
}

func GetAdminSwitchEdit(c *fiber.Ctx) error {
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
}

type SwitchQueryParams struct {
	Name             string  `form:"switch-name"`
	ShortDescription string  `form:"short-desc"`
	LongDescription  string  `form:"long-desc"`
	SwitchTypeID     string  `form:"switch-type-id"`
	BrandID          *string `form:"brand-id"`
	ManufacturerID   *string `form:"manufacturer-id"`
	ReleaseDate      *string `form:"release-date"`
	Avaliable        string  `form:"avaliable"`
	Price            string  `form:"price"`
	SiteURL          string  `form:"site-url"`
	LinkCount        string  `form:"link-count"`
}

func PostAdminSwitchCreate(c *fiber.Ctx) error {
	timer := utils.StartTimer("Create Switch")
	defer timer.LogTotalTime()
	user := c.Locals("User").(models.User)
	log.Info().Msg("Creating a new switch")

	var request SwitchQueryParams
	if err := c.BodyParser(&request); err != nil {
		log.Error().Err(err).Msg("Error parsing request")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid link count",
		})
	}
	timer.LogTime("Parse Request")

	var imageLinks []*models.ImageLink
	linkCount, err := strconv.Atoi(request.LinkCount)
	if err != nil {
		log.Error().Err(err).Msg("Error parsing link count")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid link count", "error": err,
		})
	}

	for i := 0; i < linkCount; i++ {
		linkAddress := c.FormValue(fmt.Sprintf("link-address-%d", i))
		altText := c.FormValue(fmt.Sprintf("link-alt-text-%d", i))
		if linkAddress == "" || altText == "" {
			log.Warn().Msg("Empty link address or alt text")
			continue
		}

		imageLinks = append(imageLinks, &models.ImageLink{
			LinkAddress: linkAddress,
			AltText:     altText,
			CreatedByID: user.ID,
			UpdatedByID: user.ID,
			IsPrimary:   i == 0,
			OwnerType:   "switch",
		})
	}

	switchTypeID, err := strconv.Atoi(request.SwitchTypeID)
	if err != nil {
		log.Error().Err(err).Msg("Error parsing switch type ID")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing switch type ID", "error": err,
		})
	}
	brandID, err := strconv.Atoi(*request.BrandID)
	if err != nil {
		log.Error().Err(err).Msg("Error parsing brand ID")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing brand ID", "error": err,
		})
	}
	manufacturerID, err := strconv.Atoi(*request.ManufacturerID)
	if err != nil {
		log.Error().Err(err).Msg("Error parsing manufacturer ID")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing manufacturer ID", "error": err,
		})
	}
	releaseDate, err := time.Parse("2006-01-02", *request.ReleaseDate)
	if err != nil {
		log.Error().Err(err).Msg("Error parsing release date")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing release date", "error": err,
		})
	}
	pricePoint, err := strconv.Atoi(request.Price)
	if err != nil {
		log.Error().Err(err).Msg("Error parsing price")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing price", "error": err,
		})
	}

	clickyClack := models.Switch{
		Name:             request.Name,
		ShortDescription: request.ShortDescription,
		LongDescription:  request.LongDescription,
		SwitchTypeID:     switchTypeID,
		BrandID:          &brandID,
		ManufacturerID:   &manufacturerID,
		ReleaseDate:      &releaseDate,
		Available:        request.Avaliable == "true",
		PricePoint:       pricePoint,
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
}

func DeleteImageLinkToList(c *fiber.Ctx) error {
	imageLinkIndex, err := c.ParamsInt("imageLinkIndex")
	if err != nil {
		log.Error().Msg("No image link index provided")
		return c.Status(fiber.StatusBadRequest).Next()
	}

	log.Info().Int("imageLinkIndex", imageLinkIndex).Msg("Deleting image link")
	linkCountStr := c.FormValue("link-count")
	linkCount, err := strconv.Atoi(linkCountStr)
	if err != nil {
		log.Error().Err(err).Msg("Error parsing link count")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid link count",
		})
	}

	var imageLinks []components.ImageLinksInput
	for i := 0; i < linkCount; i++ {
		if i == imageLinkIndex {
			continue
		}
		linkAddress := c.FormValue(fmt.Sprintf("link-address-%d", i))
		altText := c.FormValue(fmt.Sprintf("link-alt-text-%d", i))

		imageLinks = append(imageLinks, components.ImageLinksInput{
			LinkAddress: linkAddress,
			AltText:     altText,
		})
	}

	return Render(components.ImageLinksForm(imageLinks))(c)
}

func GetImageLinkToList(c *fiber.Ctx) error {
	log.Info().Msg("Adding a new image link")
	test := c.Query("link-count")
	log.Info().Str("link-count", test).Msg("Link count")
	linkCountStr := c.FormValue("link-count")
	linkCount, err := strconv.Atoi(linkCountStr)
	if err != nil {
		log.Error().Err(err).Msg("Error parsing link count")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid link count",
		})
	}

	var imageLinks []components.ImageLinksInput
	for i := 0; i <= linkCount; i++ {
		linkAddress := c.FormValue(fmt.Sprintf("link-address-%d", i))
		altText := c.FormValue(fmt.Sprintf("link-alt-text-%d", i))

		imageLinks = append(imageLinks, components.ImageLinksInput{
			LinkAddress: linkAddress,
			AltText:     altText,
		})
	}

	return Render(components.ImageLinksForm(imageLinks))(c)
}

type SwitchSearchQueryParams struct {
	Search string `json:"search"`
}

func GetAdminSwitches(c *fiber.Ctx) error {
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
	switchInput := components.SwitchData{}
	var brands, manufacturers []models.Producer
	if err := database.DB.Find(&brands).Error; err != nil {
		log.Error().Err(err).Msg("Error getting the brands")
		return switchInput, err
	}

	if err := database.DB.Find(&manufacturers).Error; err != nil {
		log.Error().Err(err).Msg("Error getting the manufacturers")
		return switchInput, err
	}

	var switchTypes []models.Type
	if err := database.DB.Where("category = 'switch_type'").Find(&switchTypes).Error; err != nil {
		log.Error().Err(err).Msg("Error getting the switch types")
		return switchInput, err
	}

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

	var switchOptions []components.Option
	for _, switchType := range switchTypes {
		switchOptions = append(switchOptions, components.Option{
			Value: strconv.Itoa(switchType.ID),
			Label: switchType.Name,
		})
	}

	switchInput = components.SwitchData{
		Brands:        brandOptions,
		Manufacturers: manufacturerOptions,
		SwitchTypes:   switchOptions,
	}

	return switchInput, nil
}
