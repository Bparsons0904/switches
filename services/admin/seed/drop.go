package seed

import (
	"switches/models"

	"gorm.io/gorm"
)

func seedDrop(tx *gorm.DB, admin models.User) error {
	drop := []models.Switch{
		{
			Name:             "Halo True",
			ShortDescription: "Tactile switch with a smooth bump and minimal noise.",
			LongDescription:  "Drop Halo True switches offer a smooth and satisfying tactile bump with minimal noise, making them a great choice for quiet environments. The Halo True provides a unique typing experience that balances tactility with quiet operation.",
			ManufacturerID:   ptr(3),  // Kaihua
			BrandID:          ptr(10), // Drop
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2017-12-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(29), // Low
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(42), // Rounded
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Holy Panda",
			ShortDescription: "Tactile switch with a unique feel and excellent sound.",
			LongDescription:  "The Drop Holy Panda switch is legendary for its satisfying tactile bump and signature sound profile. Combining the Halo True stem and Invyr Panda housing, it delivers a unique typing feel that is cherished by keyboard enthusiasts.",
			ManufacturerID:   ptr(2),  // Gateron
			BrandID:          ptr(10), // Drop
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2018-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(53)),
				BottomOutForce:        ptr(float32(67)),
				TotalTravel:           ptr(float32(3.8)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
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
			Name:             "Drop Invyr Holy Panda",
			ShortDescription: "Tactile switch with the famous Holy Panda feel.",
			LongDescription:  "Invyr Holy Panda switches from Drop provide the classic Holy Panda typing experience with their renowned tactile bump and satisfying sound. These switches are a must-have for tactile switch enthusiasts.",
			ManufacturerID:   ptr(2),  // Gateron
			BrandID:          ptr(10), // Drop
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2020-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive

			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(53)),
				BottomOutForce:        ptr(float32(67)),
				TotalTravel:           ptr(float32(3.8)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
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
			Name:             "ALT Rock Black",
			ShortDescription: "Linear switch with a heavier feel and smooth action.",
			LongDescription:  "Drop ALT Rock Black switches provide a smooth linear action with a heavier actuation force, catering to those who prefer more resistance in their typing. These switches are favored by gamers and typists who enjoy a substantial feel.",
			ManufacturerID:   ptr(22), // JWK
			BrandID:          ptr(10), // Drop
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2021-09-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(70)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(70)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
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
			Name:             "Silver Speed",
			ShortDescription: "Linear switch designed for fast actuation and gaming.",
			LongDescription:  "The Drop Silver Speed switch is optimized for fast actuation, making it ideal for gaming and rapid typing. Its smooth linear action ensures quick and precise inputs, reducing lag for competitive gamers.",
			ManufacturerID:   ptr(3),  // Kailh
			BrandID:          ptr(10), // Drop
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2022-04-01"),
			Available:        true,
			PricePoint:       2, // Average

			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(1.1)),
				ActuationForce:        ptr(float32(40)),
				BottomOutForce:        ptr(float32(45)),
				TotalTravel:           ptr(float32(3.5)),
				PreTravel:             ptr(float32(1.1)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(59), // Silver
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Drop Kaihua Box White",
			ShortDescription: "Clicky switch with a light actuation force and box design.",
			LongDescription:  "Drop Kaihua Box White switches feature a box design that protects against dust and moisture while offering a light actuation force and a crisp, satisfying click. They are a popular choice among fans of clicky switches who value durability and performance.",
			ManufacturerID:   ptr(3),  // Kailh
			BrandID:          ptr(10), // Drop
			SwitchTypeID:     3,       // Clicky
			ReleaseDate:      parseDate("2018-06-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(1.8)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(50)),
				TotalTravel:           ptr(float32(3.6)),
				PreTravel:             ptr(float32(1.8)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(51), // White
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
			Name:             "Kaihua Speed Silver",
			ShortDescription: "Linear switch designed for fast actuation and gaming.",
			LongDescription:  "The Drop Kaihua Speed Silver switch offers rapid actuation with a smooth linear action, making it a top choice for gamers and fast typists. Its quick response time and seamless performance ensure a competitive edge in fast-paced scenarios.",
			ManufacturerID:   ptr(3),  // Kailh
			BrandID:          ptr(10), // Drop
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2019-03-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(1.1)),
				ActuationForce:        ptr(float32(40)),
				BottomOutForce:        ptr(float32(45)),
				TotalTravel:           ptr(float32(3.5)),
				PreTravel:             ptr(float32(1.1)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(59), // Silver
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Kaihua Jade",
			ShortDescription: "Clicky switch with a thicker click bar for a louder sound.",
			LongDescription:  "Drop Kaihua Jade switches are known for their thicker click bar, resulting in a louder and more satisfying click. These switches offer pronounced tactile feedback, making them perfect for users who enjoy a bold and assertive typing experience.",
			ManufacturerID:   ptr(3),  // Kailh
			BrandID:          ptr(10), // Drop
			SwitchTypeID:     3,       // Clicky
			ReleaseDate:      parseDate("2019-07-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(1.8)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(50)),
				TotalTravel:           ptr(float32(3.6)),
				PreTravel:             ptr(float32(1.8)),
				FactoryLubed:          false,
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
			Name:             "Halo Linear",
			ShortDescription: "Linear switch with a smooth feel and minimal resistance.",
			LongDescription:  "Drop Halo Linear switches are designed for a smooth and consistent typing experience with minimal resistance. Ideal for users who prefer a linear switch without tactile feedback, they offer a straightforward and satisfying keystroke.",
			ManufacturerID:   ptr(3),  // Kaihua
			BrandID:          ptr(10), // Drop
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2019-07-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
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
			Name:             "Silent Aliaz",
			ShortDescription: "Silent tactile switch with a soft bump and minimal noise.",
			LongDescription:  "Drop Silent Aliaz switches offer a quiet tactile experience with a soft bump and minimal noise. These switches are ideal for environments where silence is essential, providing a peaceful typing experience without compromising on feedback.",
			ManufacturerID:   ptr(3),  // Kaihua
			BrandID:          ptr(10), // Drop
			SwitchTypeID:     5,       // Silent Tactile
			ReleaseDate:      parseDate("2019-11-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(3.8)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
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
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(42), // Rounded
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Kaihua Silent Pink",
			ShortDescription: "Silent linear switch with a fast actuation and pink color.",
			LongDescription:  "The Drop Kaihua Silent Pink switch provides a fast-actuating silent linear experience with a vibrant pink color. These switches are perfect for users who value both aesthetics and quiet operation, making them ideal for office use or shared spaces.",
			ManufacturerID:   ptr(3),  // Kailh
			BrandID:          ptr(10), // Drop
			SwitchTypeID:     4,       // Silent Linear
			ReleaseDate:      parseDate("2020-02-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(1.8)),
				ActuationForce:        ptr(float32(35)),
				BottomOutForce:        ptr(float32(45)),
				TotalTravel:           ptr(float32(3.6)),
				PreTravel:             ptr(float32(1.8)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(57), // Pink
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Emerald Silent Tactile",
			ShortDescription: "Silent tactile switch with a gentle bump and emerald design.",
			LongDescription:  "The Drop Emerald Silent Tactile switch features a gentle tactile bump combined with a quiet operation, making it a perfect choice for those who prefer a subdued yet satisfying typing experience. The emerald-themed design adds a touch of elegance to any keyboard.",
			ManufacturerID:   ptr(3),  // Kailh
			BrandID:          ptr(10), // Drop
			SwitchTypeID:     5,       // Silent Tactile
			ReleaseDate:      parseDate("2021-07-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(3.8)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(54), // Green
				TopHousingColorID:     ptr(54), // Green
				BottomHousingColorID:  ptr(54), // Green
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(42), // Rounded
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Holy Boba",
			ShortDescription: "Specialty tactile switch with Holy Panda and Boba U4 elements.",
			LongDescription:  "Drop Holy Boba switches are a unique combination of Holy Panda and Boba U4 elements, resulting in a premium tactile experience with a smooth bump and satisfying feel. These switches are perfect for enthusiasts who value both innovative design and excellent feedback.",
			ManufacturerID:   ptr(26), // Gazzew
			BrandID:          ptr(10), // Drop
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2022-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(68)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(68)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
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
	}

	err := processSwitches(tx, drop, admin)
	if err != nil {
		return err
	}

	return nil
}
