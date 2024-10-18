package seed

import (
	"switches/models"

	"gorm.io/gorm"
)

func seedNovelkey(tx *gorm.DB, admin models.User) error {
	novelKey := []models.Switch{
		{
			Name:             "Cream",
			ShortDescription: "Smooth linear switch with a POM housing.",
			LongDescription:  "The NovelKeys Cream switch is a highly regarded linear switch that features a full POM (polyoxymethylene) housing, offering a unique smoothness and a deep, satisfying thock sound. Known for its buttery feel and distinctive aesthetic, the Cream switch has become a favorite among keyboard enthusiasts who seek a premium typing experience with excellent durability.",
			ManufacturerID:   ptr(9), // NovelKeys
			BrandID:          ptr(9), // NovelKeys
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2019-05-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(23), // POM
				BaseHousingMaterialID: ptr(23), // POM
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(51), // White
				TopHousingColorID:     ptr(51), // White
				BottomHousingColorID:  ptr(51), // White
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Cream Tactile",
			ShortDescription: "Tactile switch with a smooth bump and POM housing.",
			LongDescription:  "The NovelKeys Cream Tactile switch combines the renowned smoothness of the original Cream switch with a tactile bump, delivering a unique typing experience. The full POM housing not only enhances the smoothness but also provides a distinct thock sound, making it an excellent choice for users who desire both tactile feedback and a satisfying auditory experience.",
			ManufacturerID:   ptr(9), // NovelKeys
			BrandID:          ptr(9), // NovelKeys
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2020-07-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(23), // POM
				BaseHousingMaterialID: ptr(23), // POM
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(51), // White
				TopHousingColorID:     ptr(51), // White
				BottomHousingColorID:  ptr(51), // White
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(42), // Rounded
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "x Kailh Box Jade",
			ShortDescription: "Clicky switch with a crisp and loud click.",
			LongDescription:  "The NovelKeys x Kailh Box Jade switch is a highly tactile and clicky switch, known for its crisp, loud click and satisfying feedback. Designed with a reinforced click bar, the Box Jade delivers a sharp and consistent keystroke, making it ideal for users who love a pronounced auditory and tactile experience.",
			ManufacturerID:   ptr(3), // Kailh
			BrandID:          ptr(9), // NovelKeys
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2018-04-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				ActuationPoint:        ptr(float32(1.8)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(3.6)),
				PreTravel:             ptr(float32(1.8)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(54), // Green
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(51), // White
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(32), // Very High
				TactilityTypeID:       ptr(40), // Click Bar
				TactilityFeedbackID:   ptr(41), // Sharp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "x Kailh Box Royal",
			ShortDescription: "Tactile switch with a heavy bump and smooth action.",
			LongDescription:  "The NovelKeys x Kailh Box Royal switch is a tactile switch that features a heavy and defined bump, paired with a smooth keystroke. It is particularly favored by typists who prefer strong tactile feedback without the click, providing a responsive and satisfying typing experience that excels in both gaming and general use.",
			ManufacturerID:   ptr(3), // Kailh
			BrandID:          ptr(9), // NovelKeys
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2018-04-01"),
			Available:        true,
			PricePoint:       2, // Average

			Details: models.Details{
				SpringForce:           ptr(float32(75)),
				ActuationPoint:        ptr(float32(1.8)),
				ActuationForce:        ptr(float32(60)),
				BottomOutForce:        ptr(float32(75)),
				TotalTravel:           ptr(float32(3.6)),
				PreTravel:             ptr(float32(1.8)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(55), // Purple
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(51), // White
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(41), // Sharp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Sherbet",
			ShortDescription: "Linear switch with a light actuation force and smooth keystrokes.",
			LongDescription:  "The NovelKeys Sherbet switch is a linear switch designed for users who prefer a light actuation force paired with smooth, effortless keystrokes. The Sherbet switch features a vibrant color scheme and delivers a gentle typing experience, making it an excellent choice for both casual users and those who appreciate a softer feel.",
			ManufacturerID:   ptr(9), // NovelKeys
			BrandID:          ptr(9), // NovelKeys
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2020-11-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(40)),
				BottomOutForce:        ptr(float32(55)),
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
			Name:             "Silk Yellow",
			ShortDescription: "Silky smooth linear switch with a light actuation force.",
			LongDescription:  "The NovelKeys Silk Yellow switch is a linear switch that offers an exceptionally smooth typing experience with a light actuation force. Built with pre-lubed components, the Silk Yellow provides minimal resistance and a near-silent keystroke, making it ideal for fast typists and those who value a seamless, silky smooth feel.",
			ManufacturerID:   ptr(9), // NovelKeys
			BrandID:          ptr(9), // NovelKeys
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-02-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(63.5)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(63.5)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(50), // Yellow
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Silk Black",
			ShortDescription: "Linear switch with a heavier actuation force and smooth feel.",
			LongDescription:  "The NovelKeys Silk Black switch offers a heavier actuation force than its yellow counterpart, while still delivering the same smooth, pre-lubed keystrokes. Ideal for users who prefer more resistance in their switches, the Silk Black provides a satisfying and consistent linear typing experience, perfect for extended typing sessions or gaming.",
			ManufacturerID:   ptr(9), // NovelKeys
			BrandID:          ptr(9), // NovelKeys
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-02-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(70)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(70)),
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
			Name:             "Silk Red",
			ShortDescription: "Smooth linear switch with a light actuation force.",
			LongDescription:  "The NovelKeys Silk Red switch is a smooth linear switch with a light actuation force, designed for users who value speed and fluidity in their typing. Pre-lubed from the factory, the Silk Red offers a consistent and quiet keystroke, making it ideal for gaming and fast-paced typing.",
			ManufacturerID:   ptr(9), // NovelKeys
			BrandID:          ptr(9), // NovelKeys
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-02-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(62)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(62)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "x Kailh Box Pale Blue",
			ShortDescription: "Clicky switch with a bright sound and crisp feel.",
			LongDescription:  "The NovelKeys x Kailh Box Pale Blue switch offers a bright and crisp clicky sound, providing a tactile and auditory experience that is both satisfying and engaging. Designed with a reinforced click bar, the Pale Blue switch is ideal for users who enjoy a sharp, precise keystroke and clear auditory feedback.",
			ManufacturerID:   ptr(3), // Kailh
			BrandID:          ptr(9), // NovelKeys
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2018-04-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				ActuationPoint:        ptr(float32(1.8)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(3.6)),
				PreTravel:             ptr(float32(1.8)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(51), // White
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				TactilityTypeID:       ptr(40), // Click Bar
				TactilityFeedbackID:   ptr(43), // Crisp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "x Kailh Box Red",
			ShortDescription: "Linear switch with a light feel and smooth action.",
			LongDescription:  "The NovelKeys x Kailh Box Red switch is a linear switch known for its light feel and smooth action. Ideal for both gaming and typing, this switch offers a responsive keystroke with minimal resistance, making it a popular choice for users who prefer a fluid and effortless typing experience.",
			ManufacturerID:   ptr(3), // Kailh
			BrandID:          ptr(9), // NovelKeys
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2018-04-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				ActuationPoint:        ptr(float32(1.8)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(3.6)),
				PreTravel:             ptr(float32(1.8)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(51), // White
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "x Kailh Box Black",
			ShortDescription: "Linear switch with a heavy actuation force and smooth feel.",
			LongDescription:  "The NovelKeys x Kailh Box Black switch offers a linear typing experience with a heavier actuation force, making it a great choice for users who prefer more resistance in their switches. With its smooth and consistent feel, the Box Black switch provides a satisfying keystroke for both heavy typists and gamers.",
			ManufacturerID:   ptr(3), // Kailh
			BrandID:          ptr(9), // NovelKeys
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2018-04-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(70)),
				ActuationPoint:        ptr(float32(1.8)),
				ActuationForce:        ptr(float32(60)),
				BottomOutForce:        ptr(float32(70)),
				TotalTravel:           ptr(float32(3.6)),
				PreTravel:             ptr(float32(1.8)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(51), // White
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "x Kailh Box Brown",
			ShortDescription: "Tactile switch with a soft bump and smooth action.",
			LongDescription:  "The NovelKeys x Kailh Box Brown switch is a tactile switch that features a soft bump with smooth action. This switch provides a balanced typing experience, ideal for users who enjoy tactile feedback without the click, making it a versatile option for both typing and gaming.",
			ManufacturerID:   ptr(3), // Kailh
			BrandID:          ptr(9), // NovelKeys
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2018-04-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				ActuationPoint:        ptr(float32(1.8)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(3.6)),
				PreTravel:             ptr(float32(1.8)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(48), // Brown
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(51), // White
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(29), // Low
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(42), // Rounded
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "x Kailh Box Pink",
			ShortDescription: "Clicky switch with a unique pink color and crisp feel.",
			LongDescription:  "The NovelKeys x Kailh Box Pink switch offers a unique aesthetic with its pink color and crisp clicky feel. This switch provides both visual appeal and a satisfying auditory experience, making it a standout choice for users who value both style and performance.",
			ManufacturerID:   ptr(3), // Kailh
			BrandID:          ptr(9), // NovelKeys
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2018-04-01"),
			Available:        true,
			PricePoint:       2, // Average

			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(1.8)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(3.6)),
				PreTravel:             ptr(float32(1.8)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(57), // Pink
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(51), // White
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				TactilityTypeID:       ptr(40), // Click Bar
				TactilityFeedbackID:   ptr(43), // Crisp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "x Kailh Box Navy",
			ShortDescription: "Heavy clicky switch with a deep, satisfying click.",
			LongDescription:  "The NovelKeys x Kailh Box Navy switch is a heavy clicky switch that provides a deep and satisfying click, making it perfect for typists who crave a pronounced tactile and auditory experience. Known for its strong actuation force and durable construction, the Box Navy switch is a favorite among enthusiasts.",
			ManufacturerID:   ptr(3), // Kailh
			BrandID:          ptr(9), // NovelKeys
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2018-04-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(70)),
				ActuationPoint:        ptr(float32(1.8)),
				ActuationForce:        ptr(float32(60)),
				BottomOutForce:        ptr(float32(70)),
				TotalTravel:           ptr(float32(3.6)),
				PreTravel:             ptr(float32(1.8)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(51), // White
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(32), // Very High
				TactilityTypeID:       ptr(40), // Click Bar
				TactilityFeedbackID:   ptr(41), // Sharp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "x Kailh Box Heavy Pale Blue",
			ShortDescription: "Heavy clicky switch with a strong tactile bump and crisp click.",
			LongDescription:  "The NovelKeys x Kailh Box Heavy Pale Blue switch is designed for those who crave a strong tactile bump and crisp click sound. Offering a satisfying and engaging typing experience, the Heavy Pale Blue switch is an excellent choice for users who prefer a heavier, more deliberate keystroke.",
			ManufacturerID:   ptr(3), // Kailh
			BrandID:          ptr(9), // NovelKeys
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2018-04-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(75)),
				ActuationPoint:        ptr(float32(1.8)),
				ActuationForce:        ptr(float32(65)),
				BottomOutForce:        ptr(float32(75)),
				TotalTravel:           ptr(float32(3.6)),
				PreTravel:             ptr(float32(1.8)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(51), // White
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(32), // Very High
				TactilityTypeID:       ptr(40), // Click Bar
				TactilityFeedbackID:   ptr(41), // Sharp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Blueberry",
			ShortDescription: "Tactile switch with a firm bump and distinct feedback.",
			LongDescription:  "The NovelKeys Blueberry switch offers a firm tactile bump with distinct feedback, providing a unique typing experience with a pronounced and satisfying keystroke. Ideal for those who enjoy a strong tactile sensation, the Blueberry switch stands out for its bold feel and exceptional build quality.",
			ManufacturerID:   ptr(9), // NovelKeys
			BrandID:          ptr(9), // NovelKeys
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2019-12-01"),
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
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(41), // Sharp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Box Royal",
			ShortDescription: "Tactile clicky switch with a heavy bump and satisfying click.",
			LongDescription:  "The NovelKeys Box Royal switch is a tactile clicky switch that offers a heavy bump and satisfying click, making it ideal for those who enjoy strong tactile feedback and a pronounced auditory experience. Its durable construction and consistent performance make it a popular choice for enthusiasts.",
			ManufacturerID:   ptr(3), // Kailh
			BrandID:          ptr(9), // NovelKeys
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2018-04-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(75)),
				ActuationPoint:        ptr(float32(1.8)),
				ActuationForce:        ptr(float32(60)),
				BottomOutForce:        ptr(float32(75)),
				TotalTravel:           ptr(float32(3.6)),
				PreTravel:             ptr(float32(1.8)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(55), // Purple
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(51), // White
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(41), // Sharp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Sherbet Tactile",
			ShortDescription: "Tactile switch with a light bump and colorful design.",
			LongDescription:  "The NovelKeys Sherbet Tactile switch offers a light tactile bump and a colorful design, providing a playful typing experience with a unique and enjoyable feel. Perfect for those who appreciate a softer tactile switch, the Sherbet Tactile adds both performance and a splash of color to any keyboard setup.",
			ManufacturerID:   ptr(9), // NovelKeys
			BrandID:          ptr(9), // NovelKeys
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2020-11-01"),
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
				StemColorID:           ptr(56), // Orange
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(42), // Rounded
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "T1",
			ShortDescription: "Tactile switch with a pronounced bump and consistent feel.",
			LongDescription:  "The NovelKeys T1 switch is known for its pronounced tactile bump and consistent feel, offering a satisfying typing experience with a strong and distinctive tactile sensation. Ideal for those who enjoy pronounced feedback, the T1 switch delivers reliability and durability for both typing and gaming.",
			ManufacturerID:   ptr(9), // NovelKeys
			BrandID:          ptr(9), // NovelKeys
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2019-10-01"),
			Available:        true,
			PricePoint:       2, // Average
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
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(41), // Sharp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "x Kailh Cream",
			ShortDescription: "Linear switch with a unique feel and smooth action.",
			LongDescription:  "The NovelKeys x Kailh Cream switch is a collaboration with Kailh, offering a unique linear feel and smooth action with a POM housing. Known for its deep, thocky sound and consistent performance, the Cream switch has become a staple for enthusiasts seeking a premium and reliable typing experience.",
			ManufacturerID:   ptr(3), // Kailh
			BrandID:          ptr(9), // NovelKeys
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2019-05-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(23), // POM
				BaseHousingMaterialID: ptr(23), // POM
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(51), // White
				TopHousingColorID:     ptr(51), // White
				BottomHousingColorID:  ptr(51), // White
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Dry Yellow",
			ShortDescription: "Linear switch with a dry feel and quick actuation.",
			LongDescription:  "The NovelKeys Dry Yellow switch provides a unique linear feel with quick actuation, offering a dry and satisfying keystroke for those who prefer a different linear experience. This switch stands out for its distinct tactile feel and is ideal for users who want something different from the typical smooth switches.",
			ManufacturerID:   ptr(9), // NovelKeys
			BrandID:          ptr(9), // NovelKeys
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2020-09-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(63.5)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(63.5)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(50), // Yellow
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Cream Clickies",
			ShortDescription: "Unique clicky switch with a POM housing.",
			LongDescription:  "The NovelKeys Cream Clickies switch combines the smooth POM housing of the Cream line with a clicky mechanism, providing a unique auditory and tactile experience that sets it apart from standard clicky switches. Ideal for users who love the thocky sound of the Cream switches but prefer a clicky feel.",
			ManufacturerID:   ptr(9), // NovelKeys
			BrandID:          ptr(9), // NovelKeys
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2020-12-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.2)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.2)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(23), // POM
				BaseHousingMaterialID: ptr(23), // POM
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(51), // White
				TopHousingColorID:     ptr(51), // White
				BottomHousingColorID:  ptr(51), // White
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				TactilityTypeID:       ptr(40), // Click Bar
				TactilityFeedbackID:   ptr(43), // Crisp
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Dry Black",
			ShortDescription: "Linear switch with a dry feel and heavier actuation.",
			LongDescription:  "The NovelKeys Dry Black switch offers a dry feel and heavier actuation force, providing a distinct linear experience that differs from typical smooth switches. Ideal for users who prefer a more solid and resistant keystroke, the Dry Black switch delivers both unique feel and reliable performance.",
			ManufacturerID:   ptr(9), // NovelKeys
			BrandID:          ptr(9), // NovelKeys
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2020-09-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(70)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(60)),
				BottomOutForce:        ptr(float32(70)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(false),
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
			Name:             "NK_ Blueberry",
			ShortDescription: "Tactile switch with a unique tactile event.",
			LongDescription:  "The NovelKeys NK_ Blueberry switch provides a distinct tactile event with a strong bump, offering a different tactile experience that stands out for those who appreciate a firm and pronounced tactile feel. With its unique characteristics, the Blueberry switch is perfect for users who enjoy a bolder typing experience.",
			ManufacturerID:   ptr(9), // NovelKeys
			BrandID:          ptr(9), // NovelKeys
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2020-11-01"),
			Available:        true,
			PricePoint:       3, // Expensive
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
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(41), // Sharp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "x Gateron Silent Black",
			ShortDescription: "Silent linear switch with a smooth feel and heavy actuation.",
			LongDescription:  "The NovelKeys x Gateron Silent Black switch offers a smooth linear action with a heavier actuation force, providing a silent and solid typing experience perfect for office environments or quiet settings. This switch is designed for users who prefer a more robust and quiet keystroke.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(9), // NovelKeys
			SwitchTypeID:     4,      // Silent Linear
			ReleaseDate:      parseDate("2019-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
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
			Name:             "x Gateron Silent Yellow",
			ShortDescription: "Silent linear switch with a light feel and smooth action.",
			LongDescription:  "The NovelKeys x Gateron Silent Yellow switch provides a smooth linear experience with a light actuation force, offering a silent keystroke ideal for those who appreciate a quiet and effortless typing session. Perfect for office environments, this switch combines performance and discretion.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(9), // NovelKeys
			SwitchTypeID:     4,      // Silent Linear
			ReleaseDate:      parseDate("2019-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(62)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(62)),
				TotalTravel:           ptr(float32(3.8)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(50), // Yellow
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Cream Speed",
			ShortDescription: "Speed linear switch with a POM housing and rapid actuation.",
			LongDescription:  "The NovelKeys Cream Speed switch offers rapid actuation and a smooth linear feel with a POM housing, designed for fast typists and gamers who demand quick response times and a seamless typing experience. The Speed switch is ideal for those who prioritize speed and precision in their keyboard setup.",
			ManufacturerID:   ptr(9), // NovelKeys
			BrandID:          ptr(9), // NovelKeys
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-03-01"),
			Available:        true,
			PricePoint:       3, // Expensive

			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(1.2)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(3.4)),
				PreTravel:             ptr(float32(1.2)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(23), // POM
				BaseHousingMaterialID: ptr(23), // POM
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(51), // White
				TopHousingColorID:     ptr(51), // White
				BottomHousingColorID:  ptr(51), // White
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "x Kailh Clickiez",
			ShortDescription: "Clicky switch with a unique mechanism and tactile feedback.",
			LongDescription:  "The NovelKeys x Kailh Clickiez switch combines a clicky mechanism with tactile feedback, providing a unique typing experience that offers both auditory and tactile satisfaction. This switch stands out for its innovative design, making it a top choice for enthusiasts who enjoy a pronounced click with each keystroke.",
			ManufacturerID:   ptr(3), // Kailh
			BrandID:          ptr(9), // NovelKeys
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2021-05-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.2)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(67)),
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
				TactilityFeedbackID:   ptr(41), // Sharp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "x Kailh Speed Navy",
			ShortDescription: "Clicky speed switch with rapid actuation and strong click.",
			LongDescription:  "The NovelKeys x Kailh Speed Navy switch offers rapid actuation with a strong clicky sound, designed for gamers and fast typists who require quick response times and a satisfying auditory experience. Its unique characteristics make it a standout choice for those who prioritize both speed and a pronounced click in their switches.",
			ManufacturerID:   ptr(3), // Kailh
			BrandID:          ptr(9), // NovelKeys
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2021-05-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(70)),
				ActuationPoint:        ptr(float32(1.4)),
				ActuationForce:        ptr(float32(60)),
				BottomOutForce:        ptr(float32(70)),
				TotalTravel:           ptr(float32(3.5)),
				PreTravel:             ptr(float32(1.4)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(51), // White
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(32), // Very High
				TactilityTypeID:       ptr(40), // Click Bar
				TactilityFeedbackID:   ptr(41), // Sharp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Eggnog",
			ShortDescription: "Linear switch with a smooth feel and unique colorway.",
			LongDescription:  "The NovelKeys Eggnog switch offers a smooth linear feel and a unique colorway inspired by the holiday treat. Providing a festive and enjoyable typing experience, the Eggnog switch is perfect for those who appreciate both performance and aesthetics in their keyboard switches.",
			ManufacturerID:   ptr(9), // NovelKeys
			BrandID:          ptr(9), // NovelKeys
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-12-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(62)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(62)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(69), // Beige
				TopHousingColorID:     ptr(67), // Coral
				BottomHousingColorID:  ptr(67), // Coral
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
	}

	err := processSwitches(tx, novelKey, admin)
	if err != nil {
		return err
	}

	return nil
}
