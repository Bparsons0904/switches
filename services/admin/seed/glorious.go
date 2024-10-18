package seed

import (
	"switches/models"

	"gorm.io/gorm"
)

func seedGlorious(tx *gorm.DB, admin models.User) error {
	glorious := []models.Switch{
		{
			Name:             "Panda",
			ShortDescription: "Tactile switch with a crisp bump and smooth travel.",
			LongDescription:  "Glorious Panda switches are renowned for their crisp tactile bump and smooth travel, providing a premium typing experience with notable tactile feedback. These switches are favored by enthusiasts for their responsive and satisfying feel.",
			ManufacturerID:   ptr(11), // Glorious
			BrandID:          ptr(11), // Glorious
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2020-07-01"),
			Available:        true,
			PricePoint:       3, // Expensive

			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(53)),
				BottomOutForce:        ptr(float32(67)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(51), // White
				BottomHousingColorID:  ptr(51), // White
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(41), // Sharp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Lynx",
			ShortDescription: "Linear switch with a smooth, fast feel for gaming.",
			LongDescription:  "The Glorious Lynx switch is designed for gaming enthusiasts who seek a smooth and fast linear feel. Its consistent actuation and quick response make it an ideal choice for competitive gaming and fast typing.",
			ManufacturerID:   ptr(11), // Glorious
			BrandID:          ptr(11), // Glorious
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2021-11-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Fox",
			ShortDescription: "Linear switch with a medium actuation force.",
			LongDescription:  "Glorious Fox switches offer a balanced linear experience with a medium actuation force, making them versatile for both gaming and everyday typing. Their smooth and consistent performance is ideal for a wide range of users.",
			ManufacturerID:   ptr(11), // Glorious
			BrandID:          ptr(11), // Glorious
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2022-06-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(63)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(48)),
				BottomOutForce:        ptr(float32(63)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(56), // Orange
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Ice",
			ShortDescription: "Linear switch with a smooth, icy feel and quiet operation.",
			LongDescription:  "The Glorious Ice switch offers a smooth linear action with quiet operation, designed to provide a cool, 'icy' feel. It's perfect for those who prefer a silent typing experience without compromising on smoothness.",
			ManufacturerID:   ptr(11), // Glorious
			BrandID:          ptr(11), // Glorious
			SwitchTypeID:     4,       // Silent Linear
			ReleaseDate:      parseDate("2023-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(62)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(47)),
				BottomOutForce:        ptr(float32(62)),
				TotalTravel:           ptr(float32(3.8)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(52), // Transparent
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Holy Panda",
			ShortDescription: "Tactile switch with a distinct bump and premium feel.",
			LongDescription:  "The Glorious Holy Panda switch is a premium tactile switch that delivers a distinct bump with a satisfying feel. Loved by enthusiasts, it offers a top-tier typing experience with excellent feedback.",
			ManufacturerID:   ptr(11), // Glorious
			BrandID:          ptr(11), // Glorious
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2021-03-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(53)),
				BottomOutForce:        ptr(float32(67)),
				TotalTravel:           ptr(float32(3.8)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(51), // White
				BottomHousingColorID:  ptr(51), // White
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(41), // Sharp
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Goat Clicky",
			ShortDescription: "Clicky switch with a medium actuation force and goat logo.",
			LongDescription:  "Glorious Goat Clicky switches feature a crisp click sound with a medium actuation force, delivering an engaging typing experience. The unique goat logo and sharp auditory feedback make them a standout choice for clicky switch fans.",
			ManufacturerID:   ptr(11), // Glorious
			BrandID:          ptr(11), // Glorious
			SwitchTypeID:     3,       // Clicky
			ReleaseDate:      parseDate("2022-02-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.2)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.2)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				TactilityTypeID:       ptr(40), // Click Bar
				TactilityFeedbackID:   ptr(43), // Crisp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Fox Clicky",
			ShortDescription: "Clicky switch with a smooth action and audible feedback.",
			LongDescription:  "The Glorious Fox Clicky switch offers smooth action paired with satisfying clicky feedback. It is designed for users who enjoy a balance of tactile feel and audible satisfaction, making it perfect for typing and gaming alike.",
			ManufacturerID:   ptr(11), // Glorious
			BrandID:          ptr(11), // Glorious
			SwitchTypeID:     3,       // Clicky
			ReleaseDate:      parseDate("2021-12-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(63)),
				ActuationPoint:        ptr(float32(2.2)),
				ActuationForce:        ptr(float32(48)),
				BottomOutForce:        ptr(float32(63)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.2)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(56), // Orange
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				TactilityTypeID:       ptr(40), // Click Bar
				TactilityFeedbackID:   ptr(43), // Crisp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Lynx Clicky",
			ShortDescription: "Clicky switch with a light actuation force for fast typing.",
			LongDescription:  "Glorious Lynx Clicky switches provide a light actuation force, ideal for fast typing and gaming. The crisp clicky sound enhances the typing experience, delivering both auditory and tactile feedback.",
			ManufacturerID:   ptr(11), // Glorious
			BrandID:          ptr(11), // Glorious
			SwitchTypeID:     3,       // Clicky
			ReleaseDate:      parseDate("2022-05-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				ActuationPoint:        ptr(float32(2.2)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.2)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				TactilityTypeID:       ptr(40), // Click Bar
				TactilityFeedbackID:   ptr(43), // Crisp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Panda Linear",
			ShortDescription: "Linear switch with a smooth and consistent feel.",
			LongDescription:  "Glorious Panda Linear switches provide a smooth and consistent linear feel, offering a premium typing experience suited for both gaming and everyday use. They are designed to deliver a high-quality, reliable keystroke.",
			ManufacturerID:   ptr(11), // Glorious
			BrandID:          ptr(11), // Glorious
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2021-09-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(53)),
				BottomOutForce:        ptr(float32(67)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(51), // White
				BottomHousingColorID:  ptr(51), // White
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Fox Linear V2",
			ShortDescription: "Linear switch with a light actuation and refined design.",
			LongDescription:  "The Glorious Fox Linear V2 switch features a light actuation force and a refined design, delivering a seamless typing experience with enhanced smoothness and aesthetics. It's perfect for users seeking a responsive and elegant switch.",
			ManufacturerID:   ptr(11), // Glorious
			BrandID:          ptr(11), // Glorious
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2023-02-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(58)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(43)),
				BottomOutForce:        ptr(float32(58)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(56), // Orange
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Lynx Linear V2",
			ShortDescription: "Linear switch with improved smoothness and sound.",
			LongDescription:  "Glorious Lynx Linear V2 switches feature enhanced smoothness and sound, providing a refined typing experience for those who prefer a quieter and more fluid keystroke. These switches are a great upgrade for linear switch fans.",
			ManufacturerID:   ptr(11), // Glorious
			BrandID:          ptr(11), // Glorious
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2022-08-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Ice Linear V2",
			ShortDescription: "Linear switch with a quiet operation and icy aesthetics.",
			LongDescription:  "The Glorious Ice Linear V2 switch offers a quiet and smooth typing experience, featuring icy aesthetics that add a cool visual appeal to any keyboard. Its silent operation makes it ideal for office use or shared spaces.",
			ManufacturerID:   ptr(11), // Glorious
			BrandID:          ptr(11), // Glorious
			SwitchTypeID:     4,       // Silent Linear
			ReleaseDate:      parseDate("2023-05-01"),
			Available:        true,
			PricePoint:       3, // Expensive

			Details: models.Details{
				SpringForce:           ptr(float32(62)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(47)),
				BottomOutForce:        ptr(float32(62)),
				TotalTravel:           ptr(float32(3.8)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(52), // Transparent
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Panda Silent",
			ShortDescription: "Silent tactile switch with a quiet bump and smooth feel.",
			LongDescription:  "Glorious Panda Silent switches offer a quiet tactile bump with a smooth feel, providing a premium typing experience without the noise. These switches are perfect for those who enjoy tactile feedback without disturbing others.",
			ManufacturerID:   ptr(11), // Glorious
			BrandID:          ptr(11), // Glorious
			SwitchTypeID:     5,       // Silent Tactile
			ReleaseDate:      parseDate("2021-11-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(53)),
				BottomOutForce:        ptr(float32(67)),
				TotalTravel:           ptr(float32(3.8)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(51), // White
				BottomHousingColorID:  ptr(51), // White
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(42), // Rounded
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Lynx Silent",
			ShortDescription: "Silent linear switch with a soft feel and quiet operation.",
			LongDescription:  "The Glorious Lynx Silent switch offers a soft linear action with silent operation, making it a great choice for offices or shared spaces. Its understated aesthetics and smooth performance make it a versatile switch for any build.",
			ManufacturerID:   ptr(11), // Glorious
			BrandID:          ptr(11), // Glorious
			SwitchTypeID:     4,       // Silent Linear
			ReleaseDate:      parseDate("2022-07-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(3.8)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Crystal Ice",
			ShortDescription: "Specialty linear switch with transparent housing and icy smoothness.",
			LongDescription:  "The Glorious Crystal Ice switch features a smooth linear action with a transparent housing, allowing RGB lighting to shine through. Its unique design and consistent feel make it a popular choice for custom builds that prioritize both aesthetics and performance.",
			ManufacturerID:   ptr(11), // Glorious
			BrandID:          ptr(11), // Glorious
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2023-04-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(62)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(47)),
				BottomOutForce:        ptr(float32(62)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(52), // Transparent
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Aurora",
			ShortDescription: "Specialty tactile switch with RGB diffusion and enhanced tactility.",
			LongDescription:  "The Glorious Aurora switch offers enhanced tactile feedback combined with excellent RGB diffusion, providing both visual appeal and a satisfying typing experience. Its innovative design makes it a standout choice for enthusiasts who value both aesthetics and performance.",
			ManufacturerID:   ptr(11), // Glorious
			BrandID:          ptr(11), // Glorious
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2023-03-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(52)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(52), // Transparent
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(41), // Sharp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Summer Sun",
			ShortDescription: "Specialty linear switch with a bright design and smooth feel.",
			LongDescription:  "The Glorious Summer Sun switch offers a bright and vibrant design paired with a smooth linear action, capturing the essence of summer in both aesthetics and typing experience. It's a popular choice for those who love cheerful keyboard setups.",
			ManufacturerID:   ptr(11), // Glorious
			BrandID:          ptr(11), // Glorious
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2023-07-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(62)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(47)),
				BottomOutForce:        ptr(float32(62)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(50), // Yellow
				TopHousingColorID:     ptr(56), // Orange
				BottomHousingColorID:  ptr(56), // Orange
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Bobcat",
			ShortDescription: "Tactile switch with a pronounced bump and medium actuation force.",
			LongDescription:  "The Glorious Bobcat switch offers a pronounced tactile bump with a medium actuation force, delivering satisfying feedback for those who prefer a more noticeable tactile response. This switch is ideal for typists seeking tactile feedback without the noise of a clicky switch.",
			ManufacturerID:   ptr(11), // Glorious
			BrandID:          ptr(11), // Glorious
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2023-09-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(68)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(68)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(48), // Brown
				BottomHousingColorID:  ptr(48), // Brown
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(41), // Sharp
				HasShineThrough:       ptr(false),
			},
		},
	}

	err := processSwitches(tx, glorious, admin)
	if err != nil {
		return err
	}

	return nil
}
