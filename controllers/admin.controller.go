package controllers

import (
	"fmt"
	"strconv"
	"switches/database"
	"switches/models"
	"switches/templates/components"
	"switches/templates/pages"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func GetAdminHome(c *fiber.Ctx) error {
	user := c.Locals("User").(models.User)
	return Render(pages.Admin(user))(c)
}

func GetAdminSwitchEdit(c *fiber.Ctx) error {
	user := c.Locals("User").(models.User)
	return Render(pages.SwitchEdit(user))(c)
}

func PostAdminSwitchCreate(c *fiber.Ctx) error {
	user := c.Locals("User").(models.User)
	log.Info().Msg("Creating a new switch")
	return Render(pages.SwitchEdit(user))(c)
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

	var imageLinks []pages.ImageLinksInput
	for i := 0; i < linkCount; i++ {
		if i == imageLinkIndex {
			continue
		}
		linkAddress := c.FormValue(fmt.Sprintf("link-address-%d", i))
		altText := c.FormValue(fmt.Sprintf("link-alt-text-%d", i))

		imageLinks = append(imageLinks, pages.ImageLinksInput{
			LinkAddress: linkAddress,
			AltText:     altText,
		})
	}

	return Render(pages.ImageLinks(imageLinks))(c)
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

	var imageLinks []pages.ImageLinksInput
	for i := 0; i <= linkCount; i++ {
		linkAddress := c.FormValue(fmt.Sprintf("link-address-%d", i))
		altText := c.FormValue(fmt.Sprintf("link-alt-text-%d", i))

		imageLinks = append(imageLinks, pages.ImageLinksInput{
			LinkAddress: linkAddress,
			AltText:     altText,
		})
	}

	return Render(pages.ImageLinks(imageLinks))(c)
}

type SwitchQueryParams struct {
	Search string `json:"search"`
}

func GetAdminSwitches(c *fiber.Ctx) error {
	var request SwitchQueryParams
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
	var brands, manufacturers []models.Producer
	if err := database.DB.Find(&brands).Error; err != nil {
		log.Error().Err(err).Msg("Error getting the brands")
		return c.Status(fiber.StatusBadRequest).Next()
	}

	if err := database.DB.Find(&manufacturers).Error; err != nil {
		log.Error().Err(err).Msg("Error getting the manufacturers")
		return c.Status(fiber.StatusBadRequest).Next()
	}

	var switchTypes []models.Type
	if err := database.DB.Where("category = 'switch_type'").Find(&switchTypes).Error; err != nil {
		log.Error().Err(err).Msg("Error getting the switch types")
		return c.Status(fiber.StatusBadRequest).Next()
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

	return Render(pages.SwitchCreate(pages.SwitchCreateInput{
		Brands:        brandOptions,
		Manufacturers: manufacturerOptions,
		SwitchTypes:   switchOptions,
	}))(c)
}
