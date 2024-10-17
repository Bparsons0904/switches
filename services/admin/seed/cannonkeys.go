package seed

import (
	"switches/models"

	"gorm.io/gorm"
)

func seedCannonkeys(tx *gorm.DB, admin models.User) error {
	cannonKeys := []models.Switch{
		{
			Name:             "Lavenders",
			ShortDescription: "Smooth linear switch with a unique lavender color.",
			LongDescription:  "CannonKeys Lavenders are linear switches celebrated for their smooth, consistent feel and light actuation force. The lavender-colored housing and stem add a unique visual flair to any keyboard, making them a favorite among enthusiasts who prioritize both aesthetics and performance.",
			ManufacturerID:   ptr(22), // JWK
			BrandID:          ptr(28), // Cannon Keys
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2020-11-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(63), // Lavender
				TopHousingColorID:     ptr(63), // Lavender
				BottomHousingColorID:  ptr(63), // Lavender
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Anubis",
			ShortDescription: "Tactile switch with a medium actuation force.",
			LongDescription:  "The CannonKeys Anubis switch is a tactile switch that delivers a satisfying medium actuation force combined with a prominent tactile bump. Known for its smooth operation and deep sound profile, the Anubis switch is highly regarded among tactile switch enthusiasts who seek a rich typing experience.",
			ManufacturerID:   ptr(23), // Tecsee
			BrandID:          ptr(28), // Cannon Keys
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2021-08-01"),
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
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(41), // Sharp
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Dragonfruit",
			ShortDescription: "Tactile switch with a heavy bump and unique aesthetics.",
			LongDescription:  "Dragonfruit switches from CannonKeys offer a distinct tactile bump and a vibrant color inspired by the exotic dragon fruit. These tactile switches provide a delightful typing experience with a balance of responsiveness and feedback, making them a standout choice for those who appreciate unique aesthetics and performance.",
			ManufacturerID:   ptr(23), // Tecsee
			BrandID:          ptr(28), // Cannon Keys
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2021-05-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(53)),
				BottomOutForce:        ptr(float32(67)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(57), // Pink
				TopHousingColorID:     ptr(57), // Pink
				BottomHousingColorID:  ptr(57), // Pink
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(41), // Sharp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Banana Split",
			ShortDescription: "Linear switch with a colorful and smooth experience.",
			LongDescription:  "CannonKeys Banana Split switches are designed for those who enjoy a vibrant, playful aesthetic paired with ultra-smooth linear action. These switches offer a delightful typing experience with minimal resistance, making them a popular choice for both gaming and typing enthusiasts.",
			ManufacturerID:   ptr(23), // Tecsee
			BrandID:          ptr(28), // Cannon Keys
			SwitchTypeID:     1,       // Linear
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
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(50), // Yellow
				TopHousingColorID:     ptr(51), // White
				BottomHousingColorID:  ptr(57), // Pink
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Durock POM",
			ShortDescription: "Linear switch with a solid POM housing.",
			LongDescription:  "The CannonKeys Durock POM switch is a premium linear switch featuring a POM (polyoxymethylene) housing that ensures a smooth and consistent typing experience. Known for its durability and minimal scratchiness, this switch is ideal for heavy typists seeking a reliable and satisfying linear feel.",
			ManufacturerID:   ptr(7),  // Durock
			BrandID:          ptr(28), // Cannon Keys
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2020-06-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(63.5)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(63.5)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(23), // POM
				BaseHousingMaterialID: ptr(23), // POM
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Blueberry",
			ShortDescription: "Clicky switch with a tactile bump and crisp sound.",
			LongDescription:  "The CannonKeys Blueberry switch is a clicky switch designed to deliver a strong tactile bump accompanied by a crisp, satisfying sound. It offers a distinctive typing experience that combines tactile feedback with a sharp auditory click, making it a favorite for those who enjoy pronounced feedback.",
			ManufacturerID:   ptr(23), // Tecsee
			BrandID:          ptr(28), // Cannon Keys
			SwitchTypeID:     3,       // Clicky
			ReleaseDate:      parseDate("2020-10-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.2)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.2)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(47), // Blue
				BottomHousingColorID:  ptr(47), // Blue
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				TactilityTypeID:       ptr(40), // Click Bar
				TactilityFeedbackID:   ptr(43), // Crisp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Orange Pies",
			ShortDescription: "Clicky switch with a strong tactile feel and sharp sound.",
			LongDescription:  "CannonKeys Orange Pies switches offer a strong tactile bump with a sharp, clicky sound. These switches are perfect for typists who prefer a pronounced tactile feedback paired with a distinct auditory experience, making each keystroke satisfying and engaging.",
			ManufacturerID:   ptr(23), // Tecsee
			BrandID:          ptr(28), // Cannon Keys
			SwitchTypeID:     3,       // Clicky
			ReleaseDate:      parseDate("2021-04-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.2)),
				ActuationForce:        ptr(float32(57)),
				BottomOutForce:        ptr(float32(67)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.2)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(56), // Orange
				TopHousingColorID:     ptr(56), // Orange
				BottomHousingColorID:  ptr(56), // Orange
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(32), // Very High
				TactilityTypeID:       ptr(40), // Click Bar
				TactilityFeedbackID:   ptr(41), // Sharp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Matcha Latte",
			ShortDescription: "Linear switch with a creamy smoothness and light actuation.",
			LongDescription:  "CannonKeys Matcha Latte switches provide a smooth linear action with a light actuation force, making them ideal for typists who prefer a gentle, effortless keystroke. The matcha-colored housing adds a unique and aesthetically pleasing touch to any keyboard.",
			ManufacturerID:   ptr(23), // Tecsee
			BrandID:          ptr(28), // Cannon Keys
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2021-06-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(58)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(43)),
				BottomOutForce:        ptr(float32(58)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(51), // White
				TopHousingColorID:     ptr(54), // Green
				BottomHousingColorID:  ptr(54), // Green
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Taro Balls",
			ShortDescription: "Linear switch with a medium weight and quiet operation.",
			LongDescription:  "CannonKeys Taro Balls switches offer a medium actuation force with a quiet and smooth linear action. Inspired by the purple taro root, these switches provide a unique look and a satisfying typing experience, perfect for those who appreciate both performance and aesthetics.",
			ManufacturerID:   ptr(23), // Tecsee
			BrandID:          ptr(28), // Cannon Keys
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2021-10-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(63)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(48)),
				BottomOutForce:        ptr(float32(63)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(55), // Purple
				TopHousingColorID:     ptr(55), // Purple
				BottomHousingColorID:  ptr(55), // Purple
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Crystal",
			ShortDescription: "Linear switch with a transparent housing and smooth feel.",
			LongDescription:  "CannonKeys Crystal switches feature a transparent housing that beautifully showcases RGB lighting while delivering a smooth linear action. These switches are designed for users who seek both performance and visual appeal, offering a clean and consistent typing experience.",
			ManufacturerID:   ptr(23), // Tecsee
			BrandID:          ptr(28), // Cannon Keys
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2022-02-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(62)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(62)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
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
			Name:             "Silent Alpaca V2",
			ShortDescription: "Silent linear switch with a smooth feel and muted sound.",
			LongDescription:  "CannonKeys Silent Alpaca V2 switches are designed to provide a silent typing experience without compromising on smoothness. The switch's unique construction minimizes noise, making it ideal for environments where quiet operation is essential, while still offering the smooth feel Alpacas are known for.",
			ManufacturerID:   ptr(22), // JWK
			BrandID:          ptr(28), // Cannon Keys
			SwitchTypeID:     4,       // Silent Linear
			ReleaseDate:      parseDate("2020-12-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(62)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(62)),
				TotalTravel:           ptr(float32(3.8)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(51), // White
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Mint Chocolate",
			ShortDescription: "Specialty switch with a mint-themed color and smooth action.",
			LongDescription:  "The CannonKeys Mint Chocolate switch is a specialty linear switch that combines a mint-themed colorway with a smooth, consistent action. Its unique design and pleasant typing feel make it a popular choice for custom keyboard enthusiasts looking for both style and performance.",
			ManufacturerID:   ptr(23), // Tecsee
			BrandID:          ptr(28), // Cannon Keys
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2022-05-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(63)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(46)),
				BottomOutForce:        ptr(float32(63)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(70), // Mint
				BottomHousingColorID:  ptr(48), // Brown
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Snow Globe",
			ShortDescription: "Specialty linear switch with frosted housing and smooth action.",
			LongDescription:  "The CannonKeys Snow Globe switch is a specialty linear switch featuring a frosted housing that complements its smooth typing action. With its winter-themed design, this switch provides both a unique aesthetic and a satisfying, smooth keystroke, perfect for custom builds.",
			ManufacturerID:   ptr(22), // JWK
			BrandID:          ptr(28), // Cannon Keys
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2022-12-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(62)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(62)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(51), // White
				TopHousingColorID:     ptr(53), // Smokey
				BottomHousingColorID:  ptr(53), // Smokey
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Milkshake",
			ShortDescription: "Specialty linear switch with a milky smooth feel.",
			LongDescription:  "CannonKeys Milkshake switches deliver a milky smooth linear action with a focus on consistency and a satisfying keystroke. Ideal for those who prefer a straightforward and enjoyable typing experience, these switches combine smoothness with a unique, playful aesthetic.",
			ManufacturerID:   ptr(23), // Tecsee
			BrandID:          ptr(28), // Cannon Keys
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2021-08-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(43)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(51), // White
				TopHousingColorID:     ptr(51), // White
				BottomHousingColorID:  ptr(51), // White
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
	}

	err := processSwitches(tx, cannonKeys, admin)
	if err != nil {
		return err
	}

	return nil
}
