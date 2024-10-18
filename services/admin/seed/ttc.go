package seed

import (
	"switches/models"

	"gorm.io/gorm"
)

func seedTtc(tx *gorm.DB, admin models.User) error {
	ttc := []models.Switch{
		{
			Name:             "Gold Pink",
			ShortDescription: "Tactile switch with a light actuation force and smooth bump.",
			LongDescription:  "The TTC Gold Pink switch is designed with a light actuation force, making it ideal for fast typists and gamers alike. It offers a smooth, rounded tactile bump that provides satisfying feedback without being overly harsh, making it a versatile choice for various typing preferences.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2019-11-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(57), // Pink
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Bluish White",
			ShortDescription: "Clicky switch with a crisp tactile bump and audible click.",
			LongDescription:  "The TTC Bluish White switch combines a crisp tactile bump with an audible click, delivering a distinct typing experience that is perfect for users who enjoy both tactile feedback and sound. Its bright color and satisfying click make it a popular choice for mechanical keyboard enthusiasts.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2018-08-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.2)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(70)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.2)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(51), // White
				TopHousingColorID:     ptr(47), // Blue
				BottomHousingColorID:  ptr(47), // Blue
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Silent Red",
			ShortDescription: "Silent linear switch with a smooth keystroke and reduced noise.",
			LongDescription:  "The TTC Silent Red switch offers a near-silent linear typing experience, thanks to its built-in dampening technology. It is ideal for quiet environments where noise reduction is key, without compromising on smooth and responsive keystrokes.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     4,      // Silent Linear
			ReleaseDate:      parseDate("2020-05-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Gold Brown V3",
			ShortDescription: "Tactile switch with a distinct bump and gold-plated springs.",
			LongDescription:  "The TTC Gold Brown V3 switch features a distinct tactile bump with the added durability of gold-plated springs. This combination ensures a premium typing experience with enhanced longevity, making it suitable for users who seek reliability and performance.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2021-06-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(48), // Brown
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Wild",
			ShortDescription: "Linear switch with a smooth actuation and wild-themed aesthetics.",
			LongDescription:  "The TTC Wild switch offers a smooth and consistent linear actuation, coupled with vibrant wild-themed aesthetics that add a unique flair to any keyboard. It’s an excellent choice for users who value both performance and distinctive design.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2023-01-01"),
			Available:        true,
			PricePoint:       1, // Value
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
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
				StemColorID:           ptr(54), // Green
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Flame Red",
			ShortDescription: "Linear switch with a smooth keystroke and vibrant red color.",
			LongDescription:  "The TTC Flame Red switch is characterized by its smooth linear keystroke and eye-catching vibrant red color. It offers both aesthetic appeal and reliable performance, making it a great option for those looking to enhance their keyboard's visual and functional aspects.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2017-09-01"),
			Available:        true,
			PricePoint:       1, // Value
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Silent Brown",
			ShortDescription: "Silent tactile switch with a smooth bump and reduced noise.",
			LongDescription:  "The TTC Silent Brown switch offers a quiet tactile typing experience with a smooth bump and dampened noise. It’s perfect for users who need a tactile feel without the noise, making it suitable for shared or quiet environments.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     5,      // Silent Tactile
			ReleaseDate:      parseDate("2020-02-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(48), // Brown
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Frozen Silent",
			ShortDescription: "Silent linear switch with a cold-themed design and smooth keystroke.",
			LongDescription:  "The TTC Frozen Silent switch combines a smooth linear action with near-silent operation, all wrapped in a cold-themed design. Its unique aesthetics and quiet performance make it a favorite for custom builds in noise-sensitive environments.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     4,      // Silent Linear
			ReleaseDate:      parseDate("2021-12-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Matrix White",
			ShortDescription: "Tactile switch with a noticeable bump and bright white color.",
			LongDescription:  "The TTC Matrix White switch offers a pronounced tactile bump paired with a bright white color, making it a standout choice for those who appreciate both performance and visual appeal in their keyboard builds.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2018-04-01"),
			Available:        true,
			PricePoint:       1, // Value
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
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
			Name:             "Speed Silver",
			ShortDescription: "Linear switch with a short actuation distance for rapid key presses.",
			LongDescription:  "The TTC Speed Silver switch is engineered for speed, featuring a short actuation distance that makes it perfect for fast-paced gaming. Its quick responsiveness ensures that every key press is registered with minimal delay, giving gamers a competitive edge.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2019-02-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(1.5)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(3.5)),
				PreTravel:             ptr(float32(1.5)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(59), // Silver
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Blue",
			ShortDescription: "Classic clicky switch with a pronounced tactile bump and audible click.",
			LongDescription:  "The TTC Blue switch is a classic clicky switch known for its pronounced tactile bump and satisfying audible click. It is a staple in the mechanical keyboard community, offering a traditional typing experience that many users love.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2015-07-01"),
			Available:        true,
			PricePoint:       1, // Value
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.2)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(70)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.2)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Gold Blue",
			ShortDescription: "Enhanced clicky switch with gold-plated springs for improved durability.",
			LongDescription:  "The TTC Gold Blue switch offers an elevated clicky experience, featuring gold-plated springs that enhance both durability and tactile feedback. Its smooth, satisfying click makes it a top choice for users seeking a premium clicky switch.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2020-03-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.2)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(70)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.2)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Gold Blue V2",
			ShortDescription: "Updated version of the TTC Gold Blue with refined click mechanism.",
			LongDescription:  "The TTC Gold Blue V2 switch builds on the success of the original, featuring a refined click mechanism that enhances both tactile and auditory feedback. It’s perfect for those who appreciate a crisp, consistent click with every keystroke.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2021-05-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.2)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(70)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.2)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Diamond Pink",
			ShortDescription: "Clicky switch with a unique pink color and diamond-like sound.",
			LongDescription:  "The TTC Diamond Pink switch stands out with its vibrant pink color and sharp, diamond-like click sound. It combines aesthetic appeal with a distinct tactile and auditory experience, making it a perfect choice for those who want their keyboard to make a statement.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2019-09-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.2)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(70)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.2)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(57), // Pink
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Golden Green V3",
			ShortDescription: "Clicky switch with a pronounced tactile bump and enhanced click sound.",
			LongDescription:  "The TTC Golden Green V3 switch delivers a powerful tactile bump coupled with an enhanced, satisfying click sound. It’s designed for users who enjoy a strong, noticeable feedback with every keystroke, making it a popular choice among clicky switch enthusiasts.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2022-08-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.2)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(70)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.2)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(54), // Green
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(32), // Very High
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Hart Blue",
			ShortDescription: "Clicky switch with a unique heart-shaped stem design.",
			LongDescription:  "The TTC Hart Blue switch stands out with its unique heart-shaped stem, offering a distinctive clicky feel with both visual and tactile appeal. This switch is perfect for users who want a playful yet satisfying typing experience.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2020-11-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.2)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(70)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.2)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Speed Purple",
			ShortDescription: "Fast actuation clicky switch designed for gaming enthusiasts.",
			LongDescription:  "The TTC Speed Purple switch is engineered for gaming, offering a fast actuation clicky experience that enhances responsiveness in gameplay. Its quick keypress and satisfying click make it an ideal choice for gamers who demand precision and speed.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2019-04-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(1.8)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(3.8)),
				PreTravel:             ptr(float32(1.8)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(55), // Purple
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Watermelon Milkshake",
			ShortDescription: "Novelty clicky switch with a fruity aesthetic and satisfying click.",
			LongDescription:  "The TTC Watermelon Milkshake switch is a fun, novelty switch with a fruity aesthetic and a crisp, satisfying click. It's perfect for those who want to add a playful touch to their keyboard while enjoying a clicky typing experience.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2021-03-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.2)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(70)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.2)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(54), // Green
				BottomHousingColorID:  ptr(54), // Green
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Crystal Ice",
			ShortDescription: "Clicky switch with a transparent housing and icy click sound.",
			LongDescription:  "The TTC Crystal Ice switch combines a clear, transparent housing with a sharp, icy click sound, offering a visually stunning and audibly satisfying typing experience. It’s an excellent choice for those who prioritize both aesthetics and performance.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2018-06-01"),
			Available:        true,
			PricePoint:       1, // Value
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.2)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(70)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.2)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(49), // Clear
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Light Ring",
			ShortDescription: "Clicky switch with integrated LED support for enhanced lighting.",
			LongDescription:  "The TTC Light Ring switch is designed to enhance your keyboard’s lighting setup with integrated LED support, all while providing a satisfying clicky feel. It’s a great choice for users looking to combine visual flair with tactile feedback.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2019-07-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.2)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(70)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.2)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Brown",
			ShortDescription: "Classic tactile switch with a gentle bump and smooth action.",
			LongDescription:  "The TTC Brown switch is a staple in the mechanical keyboard world, known for its gentle tactile bump and smooth keystroke. It offers a balanced typing experience that’s suitable for both typing and gaming.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2015-03-01"),
			Available:        true,
			PricePoint:       1, // Value
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(48), // Brown
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Silent Brown V2",
			ShortDescription: "Silent tactile switch with a refined bump and reduced noise.",
			LongDescription:  "The TTC Silent Brown V2 switch offers an upgraded silent tactile experience, with a more refined bump and significantly reduced noise. It’s ideal for those who want the feel of a tactile switch without the accompanying sound.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     5,      // Silent Tactile
			ReleaseDate:      parseDate("2021-11-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(48), // Brown
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Gold Brown",
			ShortDescription: "Tactile switch with gold-plated springs for enhanced durability.",
			LongDescription:  "The TTC Gold Brown switch combines the tactile bump loved by many with the added durability of gold-plated springs. This switch offers a premium typing experience with increased longevity, making it a top choice for those seeking both quality and performance.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2019-06-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(48), // Brown
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Watermelon V2",
			ShortDescription: "Tactile switch with a fruity aesthetic and smooth bump.",
			LongDescription:  "The TTC Watermelon V2 switch offers a fun, tactile typing experience with a smooth bump and a fruity aesthetic. It’s perfect for those who want to add a playful touch to their keyboard without sacrificing performance.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2020-09-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(54), // Green
				BottomHousingColorID:  ptr(54), // Green
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Heart",
			ShortDescription: "Tactile switch with a unique heart-shaped stem and gentle bump.",
			LongDescription:  "The TTC Heart switch combines a unique heart-shaped stem design with a gentle tactile bump, offering both a distinctive aesthetic and a satisfying typing experience. It’s perfect for users looking to add a touch of personality to their keyboard.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2020-02-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(57), // Pink
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Light Salmon",
			ShortDescription: "Tactile switch with a light actuation force and smooth bump.",
			LongDescription:  "The TTC Light Salmon switch offers a light actuation force paired with a smooth tactile bump, making it an excellent choice for users who prefer a more effortless typing experience. Its unique salmon color adds a touch of personality to any keyboard.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2019-04-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(67), // Coral
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Honey V3",
			ShortDescription: "Tactile switch with a smooth bump and honey-themed design.",
			LongDescription:  "The TTC Honey V3 switch offers a smooth tactile bump with a unique honey-themed design. It’s perfect for users who appreciate both a satisfying typing feel and a visually appealing keyboard setup.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2022-07-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(50), // Yellow
				TopHousingColorID:     ptr(50), // Yellow
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Purple Sky",
			ShortDescription: "Tactile switch with a medium actuation force and sky-themed color.",
			LongDescription:  "The TTC Purple Sky switch offers a balanced tactile typing experience with a medium actuation force and a serene sky-themed color. It’s a great choice for users who want a unique aesthetic combined with a reliable typing feel.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2021-05-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(55), // Purple
				TopHousingColorID:     ptr(47), // Blue
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Purple Starlight",
			ShortDescription: "Tactile switch with a celestial aesthetic and distinct bump.",
			LongDescription:  "The TTC Purple Starlight switch offers a celestial aesthetic combined with a distinct tactile bump. It’s perfect for users who want a typing experience that’s out of this world, both in feel and appearance.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2019-11-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
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
				StemColorID:           ptr(55), // Purple
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Greystar",
			ShortDescription: "Tactile switch with a crisp bump and neutral color palette.",
			LongDescription:  "The TTC Greystar switch offers a crisp tactile bump paired with a neutral color palette, making it an ideal choice for professional setups. Its balanced typing feel and understated design make it a versatile option for any keyboard.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2018-06-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(58), // Gray
				TopHousingColorID:     ptr(58), // Gray
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Titan Heart",
			ShortDescription: "Tactile switch with a strong bump and heart-themed design.",
			LongDescription:  "The TTC Titan Heart switch offers a strong tactile bump and a heart-themed design, providing both a satisfying typing experience and a unique aesthetic appeal. It’s perfect for those who want a bold and reliable tactile switch.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2020-11-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(70)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(57), // Pink
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(31), // High
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "V3 Blue",
			ShortDescription: "Tactile switch with a refined bump and blue-themed design.",
			LongDescription:  "The TTC V3 Blue switch offers a refined tactile bump and a vibrant blue-themed design, providing a satisfying typing feel that’s both smooth and responsive. It’s an excellent choice for those who prefer a balance between aesthetics and performance.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2018-12-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(47), // Blue
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Pink Girl",
			ShortDescription: "Tactile switch with a playful pink color and gentle bump.",
			LongDescription:  "The TTC Pink Girl switch offers a playful pink color and a gentle tactile bump, making it perfect for those who want a lighthearted yet satisfying typing experience. Its unique aesthetic adds a fun element to any keyboard setup.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     2,      // Tactile
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
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(57), // Pink
				TopHousingColorID:     ptr(57), // Pink
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Leaf",
			ShortDescription: "Tactile switch with a unique leaf-shaped stem and smooth bump.",
			LongDescription:  "The TTC Leaf switch features a unique leaf-shaped stem that provides a smooth tactile bump, offering a typing experience that’s both distinctive and enjoyable. It’s a great choice for users who want a nature-inspired design in their keyboard.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
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
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(54), // Green
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Silent Red V3",
			ShortDescription: "Silent linear switch with enhanced smoothness and reduced noise.",
			LongDescription:  "The TTC Silent Red V3 switch offers a quieter linear typing experience with enhanced smoothness, making it ideal for office environments or quiet spaces where noise reduction is essential. Its improved design ensures a whisper-quiet operation.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     4,      // Silent Linear
			ReleaseDate:      parseDate("2022-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Silent Black",
			ShortDescription: "Silent linear switch with a heavier actuation force for quieter operation.",
			LongDescription:  "The TTC Silent Black switch offers a heavier actuation force combined with silent operation, providing a linear typing experience that’s both smooth and quiet. It’s perfect for users who need a silent keyboard without sacrificing feel.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     4,      // Silent Linear
			ReleaseDate:      parseDate("2021-03-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(60)),
				BottomOutForce:        ptr(float32(75)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Silent Sky",
			ShortDescription: "Silent tactile switch with a gentle bump and quiet operation.",
			LongDescription:  "The TTC Silent Sky switch combines a gentle tactile bump with a silent operation, providing a smooth typing experience that’s both satisfying and quiet. It’s ideal for shared or quiet environments where noise reduction is key.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     5,      // Silent Tactile
			ReleaseDate:      parseDate("2020-08-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Silver Blue",
			ShortDescription: "Specialty linear switch with fast actuation for gaming enthusiasts.",
			LongDescription:  "The TTC Silver Blue switch is a specialty linear switch designed for fast actuation and low travel distance, catering specifically to gamers who require quick responsiveness. Its performance-driven design makes it a top choice for competitive gaming.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2019-11-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(1.5)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(3.5)),
				PreTravel:             ptr(float32(1.5)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(59), // Silver
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Ace",
			ShortDescription: "Premium tactile switch with a unique actuation curve and smooth bump.",
			LongDescription:  "The TTC Ace switch offers a premium tactile experience with a unique actuation curve, providing a smooth bump that enhances the typing experience. Its refined design makes it a top choice for enthusiasts seeking both quality and performance.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2023-04-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(70)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(60), // Gold
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Gold Pink V2",
			ShortDescription: "Updated tactile switch with improved smoothness and actuation.",
			LongDescription:  "The TTC Gold Pink V2 switch is an updated version of the original Gold Pink, offering improved smoothness and a more refined tactile bump. It’s ideal for both gaming and typing, providing a premium feel with every keypress.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2022-06-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(57), // Pink
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Moonlight Silver",
			ShortDescription: "Linear switch with a moonlight-themed design and smooth action.",
			LongDescription:  "The TTC Moonlight Silver switch offers a smooth linear typing experience with a moonlight-themed design, combining aesthetic appeal with reliable performance. It’s an excellent choice for those who want a keyboard that looks as good as it feels.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2020-05-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(59), // Silver
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Sunset",
			ShortDescription: "Tactile switch with a sunset-themed design and distinctive bump.",
			LongDescription:  "The TTC Sunset switch combines a distinctive tactile bump with a sunset-themed design, offering a unique typing experience that blends visual appeal with satisfying tactile feedback. It’s perfect for users who want both style and substance in their keyboard.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2021-09-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(56), // Orange
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Wild Blue",
			ShortDescription: "Tactile switch with a distinctive wild-themed design and bump.",
			LongDescription:  "The TTC Wild Blue switch offers a bold tactile bump with a wild-themed design, perfect for users seeking a distinctive and satisfying typing experience. Its unique aesthetic makes it a great addition to any keyboard build.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2020-07-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
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
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Coral Sea",
			ShortDescription: "Linear switch with a coral-themed design and smooth keystroke.",
			LongDescription:  "The TTC Coral Sea switch provides a smooth linear typing experience with a coral-themed design, offering both aesthetic appeal and reliable performance. It’s an excellent choice for custom keyboard builds where both look and feel matter.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2022-08-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(67), // Coral
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Tiger Light",
			ShortDescription: "Linear switch with a tiger-themed design and light actuation force.",
			LongDescription:  "The TTC Tiger Light switch offers a smooth linear action with a light actuation force, all wrapped up in a striking tiger-themed design. It’s perfect for those who want a fast, responsive typing experience with a touch of wild flair.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2023-03-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
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
				StemColorID:           ptr(50), // Yellow
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Forest Green",
			ShortDescription: "Tactile switch with a forest-themed design and smooth bump.",
			LongDescription:  "The TTC Forest Green switch offers a smooth tactile bump combined with a nature-inspired forest-themed design. It’s perfect for users who appreciate both performance and aesthetics, bringing a touch of the outdoors to their keyboard.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2019-10-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(54), // Green
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Mars Orange",
			ShortDescription: "Tactile switch with a Mars-themed design and distinct tactile bump.",
			LongDescription:  "The TTC Mars Orange switch offers a tactile typing experience with a distinct bump and a bold Mars-themed design, perfect for users who want a keyboard that’s both functional and visually striking.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
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
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(56), // Orange
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Neptune Blue",
			ShortDescription: "Linear switch with a Neptune-themed design and smooth keystroke.",
			LongDescription:  "The TTC Neptune Blue switch provides a smooth linear typing experience with a Neptune-themed design, making it a great choice for users who want a keyboard that combines aquatic aesthetics with top-notch performance.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-05-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Cool Mint",
			ShortDescription: "Linear switch with a mint-themed design and cool aesthetic.",
			LongDescription:  "The TTC Cool Mint switch offers a refreshing linear typing experience with a mint-themed design, providing both smooth action and a cool aesthetic that’s perfect for custom keyboard builds.",
			ManufacturerID:   ptr(6), // TTC
			BrandID:          ptr(6), // TTC
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2022-10-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(70), // Mint
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
	}

	err := processSwitches(tx, ttc, admin)
	if err != nil {
		return err
	}

	return nil
}
