package seed

import (
	"switches/models"

	"gorm.io/gorm"
)

func seedOutemus(tx *gorm.DB, admin models.User) error {
	outemus := []models.Switch{
		{
			Name:             "Blue",
			ShortDescription: "Clicky switch with a tactile bump and distinct sound.",
			LongDescription:  "Outemu Blue switches are known for their clicky sound and tactile bump, making them a popular choice for typists who enjoy audible feedback. With an actuation force of 50 grams, these switches provide a satisfying click and resistance, ideal for typists who value strong tactile and auditory feedback.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2015-01-01"),
			Available:        true,
			PricePoint:       1, // Value
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Ice Blue",
			ShortDescription: "Transparent clicky switch for RGB setups.",
			LongDescription:  "Outemu Ice Blue switches feature a transparent housing, making them perfect for RGB lighting setups in mechanical keyboards. These switches offer a clicky and tactile experience with an actuation force of 50 grams, delivering a satisfying typing feel while enhancing the visual aesthetics of your keyboard.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2017-06-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Purple",
			ShortDescription: "Clicky switch with a louder sound profile.",
			LongDescription:  "Outemu Purple switches are designed to deliver a louder click sound with a strong tactile bump. With an actuation force of 55 grams, these switches appeal to users who enjoy pronounced feedback, making them ideal for those who seek a more noticeable auditory experience during typing.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2016-03-01"),
			Available:        true,
			PricePoint:       1, // Value
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(70)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(55), // Purple
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(32), // Very High
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Black Dustproof",
			ShortDescription: "Clicky switch with dustproof features.",
			LongDescription:  "Outemu Black Dustproof switches offer a clicky typing experience combined with dustproof features, enhancing durability and consistent performance even in dusty environments. With an actuation force of 55 grams, these switches provide a tactile and audible feedback while ensuring longevity.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2018-05-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(70)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Green",
			ShortDescription: "Heavy clicky switch with strong tactile feedback.",
			LongDescription:  "Outemu Green switches are heavy clicky switches that require more force to actuate, offering a strong tactile bump with an actuation force of 80 grams. Perfect for those who prefer a more resistant key press, these switches provide a satisfying and deliberate typing experience.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2017-09-01"),
			Available:        true,
			PricePoint:       1, // Value
			Details: models.Details{
				SpringForce:           ptr(float32(80)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(80)),
				BottomOutForce:        ptr(float32(95)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(54), // Green
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "White",
			ShortDescription: "Smooth clicky switch with a clear sound.",
			LongDescription:  "Outemu White switches offer a smooth clicky experience with a clear and sharp sound. With an actuation force of 50 grams, these switches provide an enjoyable typing experience for those who appreciate crisp auditory feedback combined with smooth keystrokes.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2019-11-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(51), // White
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Sky Blue",
			ShortDescription: "Clicky switch with a balanced feel.",
			LongDescription:  "Outemu Sky Blue switches provide a balanced clicky experience with an actuation force of 55 grams, offering a satisfying tactile bump and moderate resistance for a comfortable typing experience. These switches are ideal for users who want a balance between tactile feedback and smooth operation.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2020-03-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(70)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Ocean Blue",
			ShortDescription: "Clicky switch with a light actuation force.",
			LongDescription:  "Outemu Ocean Blue switches are designed for users who prefer a clicky switch with a lighter actuation force of 45 grams. These switches provide a quick response, making them suitable for fast typing and gaming, while still delivering satisfying tactile feedback.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2021-07-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Polar Blue",
			ShortDescription: "Clicky switch with enhanced tactile feedback.",
			LongDescription:  "Outemu Polar Blue switches provide enhanced tactile feedback with a clicky sound and an actuation force of 60 grams. These switches are ideal for users who value strong tactile sensations and a pronounced click, offering a robust and satisfying typing experience.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2022-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(60)),
				BottomOutForce:        ptr(float32(75)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Brown",
			ShortDescription: "Tactile switch with a subtle bump and no click.",
			LongDescription:  "Outemu Brown switches offer a smooth tactile bump without a click sound, making them versatile for both typing and gaming enthusiasts. With an actuation force of 45 grams, these switches are perfect for users who prefer quieter switches with a subtle tactile response.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2015-01-01"),
			Available:        true,
			PricePoint:       1, // Value
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(48), // Brown
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Ice Brown",
			ShortDescription: "Tactile switch with transparent housing.",
			LongDescription:  "Outemu Ice Brown switches provide a tactile experience with a transparent housing, allowing for enhanced RGB lighting while maintaining a tactile feel. With an actuation force of 45 grams, these switches are ideal for users who seek both performance and aesthetics in their keyboard setups.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2017-06-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(48), // Brown
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Purple Tactile",
			ShortDescription: "Tactile switch with a distinct bump.",
			LongDescription:  "Outemu Purple Tactile switches offer a pronounced tactile bump with no click, providing a satisfying feedback with an actuation force of 50 grams. These switches are ideal for typing without the audible click sound, catering to users who prefer a quieter but responsive typing experience.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2018-03-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(55), // Purple
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Orange",
			ShortDescription: "Medium tactile switch with a light tactile bump.",
			LongDescription:  "Outemu Orange switches offer a medium tactile experience with a lighter bump and an actuation force of 50 grams. These switches are ideal for users who prefer a softer touch with tactile feedback, providing a smooth and balanced typing feel.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2019-05-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(56), // Orange
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Sky Tactile",
			ShortDescription: "Tactile switch with a clear, satisfying bump.",
			LongDescription:  "Outemu Sky Tactile switches provide a satisfying tactile bump with a clear actuation point and an actuation force of 55 grams. These switches are perfect for users who want precise feedback during typing and gaming, offering a balanced and responsive experience.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2020-02-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(70)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Aqua",
			ShortDescription: "Tactile switch with a smooth bump and light actuation.",
			LongDescription:  "Outemu Aqua switches are tactile switches that offer a smooth bump with a light actuation force of 45 grams. Ideal for users who want a tactile feel without too much resistance, these switches provide a comfortable and responsive typing experience.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2021-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(65), // Cyan
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Milk Brown",
			ShortDescription: "Tactile switch with a muted sound profile.",
			LongDescription:  "Outemu Milk Brown switches provide a tactile typing experience with a muted sound profile and an actuation force of 50 grams. Perfect for quiet environments where tactile feedback is desired without noise, these switches offer a smooth and satisfying typing feel.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2022-04-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(48), // Brown
				TopHousingColorID:     ptr(51), // White
				BottomHousingColorID:  ptr(51), // White
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Jade",
			ShortDescription: "Heavy tactile switch with a strong bump.",
			LongDescription:  "Outemu Jade switches are heavy tactile switches with an actuation force of 60 grams, offering a strong tactile bump and resistance. Ideal for those who enjoy pronounced tactile feedback, these switches provide a deliberate and satisfying typing experience.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2022-08-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(60)),
				BottomOutForce:        ptr(float32(75)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(54), // Green
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(31), // High
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Black",
			ShortDescription: "Heavy linear switch with no tactile bump.",
			LongDescription:  "Outemu Black switches are heavier linear switches with an actuation force of 60 grams, requiring more force to actuate. These switches offer a satisfying feel for those who prefer a more resistant key press with smooth and consistent keystrokes.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2016-01-01"),
			Available:        true,
			PricePoint:       1, // Value
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(60)),
				BottomOutForce:        ptr(float32(75)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Ice Black",
			ShortDescription: "Linear switch with a transparent housing.",
			LongDescription:  "Outemu Ice Black switches offer a heavy linear feel with a transparent housing, allowing for enhanced RGB lighting effects. With an actuation force of 60 grams, these switches provide a robust typing experience with smooth and consistent performance.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2017-07-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(60)),
				BottomOutForce:        ptr(float32(75)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Yellow",
			ShortDescription: "Medium-weight linear switch with smooth keystroke.",
			LongDescription:  "Outemu Yellow switches provide a smooth linear experience with medium actuation force of 50 grams, suitable for both typing and gaming. These switches offer a balanced typing feel with no tactile bump, ideal for users who prefer a consistent and smooth keystroke.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2018-02-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(50), // Yellow
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Silver",
			ShortDescription: "Speed linear switch with reduced travel distance.",
			LongDescription:  "Outemu Silver switches are speed linear switches designed for fast typing and gaming, featuring an actuation force of 45 grams and reduced travel distance for quicker actuation. These switches are perfect for users who prioritize speed and responsiveness in their keyboards.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2019-06-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(1.5)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(3.5)),
				PreTravel:             ptr(float32(1.5)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(59), // Silver
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Red Dustproof",
			ShortDescription: "Dustproof linear switch with smooth actuation.",
			LongDescription:  "Outemu Red Dustproof switches are linear switches with dustproof features, offering smooth actuation and an actuation force of 45 grams. These switches are ideal for users who want reliable performance in various environments, including those prone to dust.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2020-04-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Pink",
			ShortDescription: "Linear switch with light actuation force.",
			LongDescription:  "Outemu Pink switches offer a linear typing experience with light actuation force of 45 grams, making them suitable for users who prefer minimal resistance. These switches provide a smooth and gentle typing feel, perfect for long typing sessions.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2020-11-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(57), // Pink
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Peach",
			ShortDescription: "Smooth linear switch with medium resistance.",
			LongDescription:  "Outemu Peach switches provide a smooth linear experience with medium resistance and an actuation force of 50 grams. Perfect for typists who enjoy a balanced typing feel, these switches offer a consistent and satisfying keystroke.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-03-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(71), // Peach
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Milk White",
			ShortDescription: "Silent linear switch with noise-reducing features.",
			LongDescription:  "Outemu Milk White switches are silent linear switches with noise-reducing features, offering a quiet typing experience for those who value silence. With an actuation force of 45 grams, these switches provide smooth and consistent keystrokes without the noise.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     4,      // Silent Linear
			ReleaseDate:      parseDate("2021-08-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(51), // White
				TopHousingColorID:     ptr(51), // White
				BottomHousingColorID:  ptr(51), // White
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Cream",
			ShortDescription: "Smooth linear switch with a creamy feel.",
			LongDescription:  "Outemu Cream switches offer a linear experience with a creamy, smooth feel and an actuation force of 50 grams. Ideal for users who enjoy a luxurious typing sensation, these switches provide a consistent and satisfying keystroke.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2022-02-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(66), // Ivory
				TopHousingColorID:     ptr(66), // Ivory
				BottomHousingColorID:  ptr(66), // Ivory
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Purple Linear",
			ShortDescription: "Linear switch with a smooth keystroke and no bump.",
			LongDescription:  "Outemu Purple Linear switches offer a smooth typing experience with no tactile bump and an actuation force of 45 grams. These switches provide a consistent feel for fast typing and gaming, ideal for users who prefer a seamless and uninterrupted keystroke.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2022-09-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(55), // Purple
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Silent White",
			ShortDescription: "Silent tactile switch with noise reduction.",
			LongDescription:  "Outemu Silent White switches offer a tactile typing experience with built-in noise-reducing features, perfect for quiet environments. With an actuation force of 45 grams, these switches provide a smooth tactile bump with minimal noise, catering to users who prefer a quieter typing experience.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     5,      // Silent Tactile
			ReleaseDate:      parseDate("2019-10-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(51), // White
				TopHousingColorID:     ptr(51), // White
				BottomHousingColorID:  ptr(51), // White
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Silent Red",
			ShortDescription: "Silent linear switch with quiet operation.",
			LongDescription:  "Outemu Silent Red switches provide a smooth linear action with noise-dampening technology and an actuation force of 45 grams. These switches are ideal for environments where silence is essential, offering a quiet and responsive typing experience.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     4,      // Silent Linear
			ReleaseDate:      parseDate("2020-03-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Silent Black",
			ShortDescription: "Heavy silent linear switch with quiet performance.",
			LongDescription:  "Outemu Silent Black switches are designed for those who prefer a heavier linear switch with an actuation force of 60 grams. These switches offer smooth operation with noise-dampening technology, providing resistance without the noise, perfect for quiet environments.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     4,      // Silent Linear
			ReleaseDate:      parseDate("2020-08-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(60)),
				BottomOutForce:        ptr(float32(75)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Ice Purple",
			ShortDescription: "Tactile switch with enhanced feedback and RGB support.",
			LongDescription:  "Outemu Ice Purple switches provide tactile feedback with a clear housing and an actuation force of 50 grams, allowing for optimal RGB lighting effects. These switches are perfect for users who desire both feedback and aesthetics, offering a balance between tactile sensation and visual appeal.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2018-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(55), // Purple
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Polar Ice",
			ShortDescription: "Specialty switch with a smooth linear feel.",
			LongDescription:  "Outemu Polar Ice switches offer a unique smooth linear experience with an actuation force of 50 grams. Featuring a custom stem design that enhances stability and reduces wobble, these switches are ideal for precise typing, offering a consistent and stable keystroke.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-02-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Fire",
			ShortDescription: "Specialty switch with fast actuation.",
			LongDescription:  "Outemu Fire switches are designed for rapid actuation with an actuation force of 45 grams, providing a quick response suitable for gaming and fast typing. Ideal for users who prioritize speed and responsiveness, these switches offer a fast and precise typing experience.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-05-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(1.5)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(3.5)),
				PreTravel:             ptr(float32(1.5)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(56), // Orange
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Shadow",
			ShortDescription: "Silent tactile switch with subtle feedback.",
			LongDescription:  "Outemu Shadow switches offer a silent tactile experience with a subtle bump and an actuation force of 50 grams, making them perfect for users who enjoy tactile feedback without noise. These switches are ideal for quiet environments where a balance between feedback and silence is desired.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     5,      // Silent Tactile
			ReleaseDate:      parseDate("2021-09-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Crystal",
			ShortDescription: "Transparent switch for RGB enthusiasts.",
			LongDescription:  "Outemu Crystal switches feature a fully transparent housing, designed to maximize RGB lighting while providing a consistent linear feel for typing and gaming. With an actuation force of 50 grams, these switches are perfect for users who want both performance and aesthetics in their keyboards.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2022-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(49), // Clear
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Teal",
			ShortDescription: "Tactile switch with a unique stem design.",
			LongDescription:  "Outemu Teal switches offer a tactile experience with a specially designed stem that provides a unique feedback and enhances overall stability. With an actuation force of 50 grams, these switches are perfect for users who appreciate distinctive tactile sensations and a consistent typing feel.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2022-06-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(62), // Teal
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Moonlight",
			ShortDescription: "Silent linear switch with a smooth keystroke.",
			LongDescription:  "Outemu Moonlight switches provide a silent linear experience with a smooth keystroke and an actuation force of 45 grams. These switches offer a whisper-quiet operation, making them perfect for users who need silence while typing without sacrificing smoothness.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     4,      // Silent Linear
			ReleaseDate:      parseDate("2022-11-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(51), // White
				TopHousingColorID:     ptr(51), // White
				BottomHousingColorID:  ptr(51), // White
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Red",
			ShortDescription: "Light linear switch for smooth typing and gaming.",
			LongDescription:  "Outemu Red switches are light linear switches with an actuation force of 45 grams. These switches offer a smooth keystroke without tactile feedback, making them suitable for both typing and gaming where a seamless keystroke is desired.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2015-01-01"),
			Available:        true,
			PricePoint:       1, // Value
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Ice Light Purple",
			ShortDescription: "Light tactile switch with transparent housing.",
			LongDescription:  "Outemu Ice Light Purple switches offer a lighter tactile experience with an actuation force of 45 grams. These switches feature a transparent housing for enhanced RGB lighting effects, making them ideal for users who seek both a lighter touch and visual appeal.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2018-06-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(55), // Purple
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Dustproof Red",
			ShortDescription: "Linear switch with dustproof design.",
			LongDescription:  "Outemu Dustproof Red switches are linear switches with a dustproof design and an actuation force of 45 grams. These switches offer consistent performance and increased durability, making them ideal for various environments, including those prone to dust.",
			ManufacturerID:   ptr(4), // Outemu
			BrandID:          ptr(4), // Outemu
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2018-09-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(22), // ABS
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(false),
			},
		},
	}
	err := processSwitches(tx, outemus, admin)
	if err != nil {
		return err
	}

	return nil
}
