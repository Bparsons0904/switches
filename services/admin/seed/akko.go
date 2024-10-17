package seed

import (
	"switches/models"

	"gorm.io/gorm"
)

func seedAkko(tx *gorm.DB, admin models.User) error {
	akko := []models.Switch{
		{
			Name:             "CS Jelly Pink",
			ShortDescription: "Linear switch with a smooth feel and moderate actuation force.",
			LongDescription:  "The Akko CS Jelly Pink switch is a linear switch that delivers a smooth and consistent typing experience with a moderate actuation force. Known for its durability and reliability, the Jelly Pink is favored by both typists and gamers who seek a switch that performs well across various tasks. Its vibrant pink hue adds a playful aesthetic to any keyboard setup, making it a popular choice among those who appreciate both form and function.",
			ManufacturerID:   ptr(8), // Akko
			BrandID:          ptr(8), // Akko
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-05-01"),
			Available:        true,
			PricePoint:       1, // Value
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(1.9)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(1.9)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(57), // Pink
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "CS Ocean Blue",
			ShortDescription: "Tactile switch with a light tactile bump and smooth operation.",
			LongDescription:  "The Akko CS Ocean Blue switch features a light tactile bump that provides a gentle yet distinct typing experience. With its smooth operation and moderate actuation force, this switch is ideal for users who enjoy the feedback of a tactile switch without the clicky sound. The Ocean Blue’s serene color and reliable performance make it a great addition to any keyboard that values both aesthetics and a pleasant typing feel.",
			ManufacturerID:   ptr(8), // Akko
			BrandID:          ptr(8), // Akko
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2021-05-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(1.9)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(1.9)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "CS Lavender Purple",
			ShortDescription: "Tactile switch with a pronounced bump and higher actuation force.",
			LongDescription:  "The Akko CS Lavender Purple switch offers a pronounced tactile bump with a higher actuation force, delivering a satisfying typing experience for users who prefer a more substantial feel. Its durable construction ensures long-lasting performance, making it a popular choice for those who demand reliability from their switches. The Lavender Purple's rich color adds a bold visual element to any keyboard setup, making it both a functional and stylish choice.",
			ManufacturerID:   ptr(8), // Akko
			BrandID:          ptr(8), // Akko
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2021-05-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				ActuationPoint:        ptr(float32(1.9)),
				ActuationForce:        ptr(float32(60)),
				BottomOutForce:        ptr(float32(70)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(1.9)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
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
			Name:             "CS Rose Red",
			ShortDescription: "Linear switch with a light actuation force for fast typing.",
			LongDescription:  "The Akko CS Rose Red switch is designed for fast typing with its light actuation force and smooth linear action. Ideal for both gamers and typists who prefer a quiet experience, this switch offers quick and responsive keystrokes with minimal noise. The Rose Red's vibrant color and efficient performance make it a favorite for those who prioritize speed and subtlety in their typing.",
			ManufacturerID:   ptr(8), // Akko
			BrandID:          ptr(8), // Akko
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-05-01"),
			Available:        true,
			PricePoint:       1, // Value
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(1.9)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(1.9)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "CS Radiant Red",
			ShortDescription: "Linear switch with a heavier actuation force for deliberate keystrokes.",
			LongDescription:  "The Akko CS Radiant Red switch is a linear switch designed with a heavier actuation force, catering to users who prefer a more deliberate typing experience. Praised for its smoothness and durability, the Radiant Red is ideal for those who want a switch that provides a firm and controlled keystroke. Its bold red color and consistent performance make it a standout choice for users who appreciate both aesthetics and functionality.",
			ManufacturerID:   ptr(8), // Akko
			BrandID:          ptr(8), // Akko
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-05-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(63.5)),
				ActuationPoint:        ptr(float32(1.9)),
				ActuationForce:        ptr(float32(63.5)),
				BottomOutForce:        ptr(float32(73.5)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(1.9)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "CS Matcha Green",
			ShortDescription: "Tactile switch with a medium tactile bump and balanced actuation force.",
			LongDescription:  "The Akko CS Matcha Green switch is a tactile switch that offers a medium tactile bump and a balanced actuation force, making it suitable for both typing and gaming. Its distinctive green color adds a unique touch to any keyboard, while its smooth operation ensures a satisfying and consistent typing experience. The Matcha Green switch is perfect for users who appreciate a blend of style and performance.",
			ManufacturerID:   ptr(8), // Akko
			BrandID:          ptr(8), // Akko
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2021-05-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(1.9)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(1.9)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(54), // Green
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "CS Crystal",
			ShortDescription: "Silent linear switch with smooth keystrokes and minimal noise.",
			LongDescription:  "The Akko CS Crystal switch is a silent linear switch that provides smooth keystrokes with minimal noise, making it ideal for office settings and quiet environments. Its transparent housing allows for RGB lighting to shine through beautifully, adding a visually striking element to any keyboard. The Crystal switch combines silence, smoothness, and style, making it a top choice for users who want a quiet yet visually appealing switch.",
			ManufacturerID:   ptr(8), // Akko
			BrandID:          ptr(8), // Akko
			SwitchTypeID:     4,      // Silent Linear
			ReleaseDate:      parseDate("2021-06-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(1.9)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(1.9)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(52), // Transparent
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "CS Midnight Blue",
			ShortDescription: "Clicky switch with a deep, satisfying click and firm actuation force.",
			LongDescription:  "The Akko CS Midnight Blue switch is a clicky switch that offers a deep, satisfying click with a firm actuation force. This switch is perfect for users who enjoy a more deliberate and tactile typing experience. The Midnight Blue's rich color and distinct sound profile make it a favorite for those who prefer a traditional clicky feel with added style.",
			ManufacturerID:   ptr(8), // Akko
			BrandID:          ptr(8), // Akko
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2021-03-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(1.9)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(1.9)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
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
			Name:             "CS Blue",
			ShortDescription: "Traditional clicky switch with a sharp click sound and tactile bump.",
			LongDescription:  "The Akko CS Blue switch is a classic clicky switch that provides a sharp click sound and a noticeable tactile bump. It's an excellent option for users who enjoy a traditional clicky typing experience with a bit of resistance. The CS Blue switch's combination of crisp feedback and reliable performance makes it a popular choice for those who appreciate a vintage typing feel.",
			ManufacturerID:   ptr(8), // Akko
			BrandID:          ptr(8), // Akko
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2021-05-01"),
			Available:        true,
			PricePoint:       1, // Value
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				ActuationPoint:        ptr(float32(1.9)),
				ActuationForce:        ptr(float32(60)),
				BottomOutForce:        ptr(float32(70)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(1.9)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
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
			Name:             "CS Silver",
			ShortDescription: "Clicky switch with a bright sound and quick actuation.",
			LongDescription:  "The Akko CS Silver switch is a clicky switch known for its bright and sharp sound. Its quick actuation makes it ideal for fast typists who enjoy a clicky feel without too much resistance. The CS Silver switch's lively sound profile and responsive action make it a great choice for users who prioritize speed and tactile feedback.",
			ManufacturerID:   ptr(8), // Akko
			BrandID:          ptr(8), // Akko
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2021-06-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(1.5)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(3.5)),
				PreTravel:             ptr(float32(1.5)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(59), // Silver
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "CS Neon",
			ShortDescription: "Colorful clicky switch with a unique design and clear sound.",
			LongDescription:  "The Akko CS Neon switch is a colorful clicky switch that stands out with its vibrant design and clear clicky sound. Its unique appearance makes it a fun addition to any keyboard setup, while its tactile bump provides a satisfying and responsive typing experience. The Neon switch is perfect for users who want to combine a lively aesthetic with reliable performance.",
			ManufacturerID:   ptr(8), // Akko
			BrandID:          ptr(8), // Akko
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2021-07-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(1.9)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(1.9)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(52), // Transparent
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "CS Sponge",
			ShortDescription: "Tactile switch with a gentle bump and soft feedback.",
			LongDescription:  "The Akko CS Sponge switch offers a gentle tactile bump with soft feedback, making it ideal for users who prefer a quieter and more subtle typing experience. Its unique design and feel make it a popular choice among enthusiasts who appreciate a delicate balance between tactility and comfort. The Sponge switch is perfect for those who seek a smooth and understated typing experience.",
			ManufacturerID:   ptr(8), // Akko
			BrandID:          ptr(8), // Akko
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2021-04-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(69), // Beige
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "CS Starfish",
			ShortDescription: "Tactile switch with a soft bump and light actuation force.",
			LongDescription:  "The Akko CS Starfish switch features a soft tactile bump with a light actuation force, offering a comfortable typing experience. Its smooth operation and unique design make it a favorite among those who appreciate a lighter tactile feel. The Starfish switch's gentle feedback and playful appearance make it a standout choice for users who want both comfort and style in their typing experience.",
			ManufacturerID:   ptr(8), // Akko
			BrandID:          ptr(8), // Akko
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2021-05-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(40)),
				ActuationPoint:        ptr(float32(1.9)),
				ActuationForce:        ptr(float32(40)),
				BottomOutForce:        ptr(float32(50)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(1.9)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(56), // Orange
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "CS Orange",
			ShortDescription: "Tactile switch with a pronounced bump and medium actuation force.",
			LongDescription:  "The Akko CS Orange switch is known for its pronounced tactile bump and medium actuation force, providing a satisfying typing experience for those who prefer more noticeable feedback. It's praised for its durability and consistent performance, making it a reliable choice for users who want a tactile switch that offers both comfort and longevity.",
			ManufacturerID:   ptr(8), // Akko
			BrandID:          ptr(8), // Akko
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2021-06-01"),
			Available:        true,
			PricePoint:       1, // Value
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(56), // Orange
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "CS Green",
			ShortDescription: "Tactile switch with a crisp bump and responsive actuation.",
			LongDescription:  "The Akko CS Green switch offers a crisp tactile bump with responsive actuation, making it an excellent choice for users who enjoy a tactile typing experience with quick feedback. The switch is designed for robust performance, ensuring durability and consistency throughout its lifespan. With its vibrant green color, the CS Green switch not only provides a satisfying typing feel but also adds a fresh aesthetic to any keyboard setup, making it a favorite for enthusiasts who value both function and style.",
			ManufacturerID:   ptr(8), // Akko
			BrandID:          ptr(8), // Akko
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2021-07-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(1.9)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(1.9)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(54), // Green
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "CS Silent",
			ShortDescription: "Silent tactile switch with smooth operation and quiet feedback.",
			LongDescription:  "The Akko CS Silent switch is a silent tactile switch that combines smooth operation with quiet feedback, making it ideal for office settings and quiet environments. Its design ensures a pleasant typing experience without sacrificing tactility, allowing users to enjoy responsive keystrokes with minimal noise. The CS Silent switch is perfect for those who need a reliable and quiet switch that still delivers satisfying tactile feedback.",
			ManufacturerID:   ptr(8), // Akko
			BrandID:          ptr(8), // Akko
			SwitchTypeID:     5,      // Silent Tactile
			ReleaseDate:      parseDate("2021-08-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "CS Peach",
			ShortDescription: "Linear switch with a soft feel and smooth keystrokes.",
			LongDescription:  "The Akko CS Peach switch offers a soft linear feel combined with smooth keystrokes, making it an excellent choice for typists and gamers who prefer a quiet and fluid typing experience. Its gentle actuation and subtle design make it a versatile option for various keyboard setups, while its consistent performance ensures a pleasant typing experience across long sessions. The Peach switch's understated elegance and reliable action make it a standout choice for users seeking both comfort and quality.",
			ManufacturerID:   ptr(8), // Akko
			BrandID:          ptr(8), // Akko
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-03-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(1.9)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(1.9)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(71), // Peach
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "CS Cream Yellow",
			ShortDescription: "Linear switch with a buttery smooth action and light actuation force.",
			LongDescription:  "The Akko CS Cream Yellow switch is renowned for its buttery smooth action and light actuation force, offering a typing experience that is both fast and effortless. Ideal for extended typing sessions or fast-paced gaming, this switch delivers reliable performance with minimal resistance. Its soft yellow hue adds a subtle yet distinctive touch to any keyboard, making it a popular choice among linear switch enthusiasts.",
			ManufacturerID:   ptr(8), // Akko
			BrandID:          ptr(8), // Akko
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-05-01"),
			Available:        true,
			PricePoint:       1, // Value
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(1.9)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(1.9)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(50), // Yellow
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "CS Mint Green",
			ShortDescription: "Linear switch with a crisp feel and moderate actuation force.",
			LongDescription:  "The Akko CS Mint Green switch offers a crisp linear feel with a moderate actuation force, making it an excellent choice for users who enjoy a balanced typing experience. Its fresh mint green color adds a stylish touch to any keyboard setup, while its smooth and consistent keystrokes provide a satisfying feel with every press. The Mint Green switch's blend of visual appeal and reliable performance makes it a versatile option for both casual and dedicated users.",
			ManufacturerID:   ptr(8), // Akko
			BrandID:          ptr(8), // Akko
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-06-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(1.9)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(1.9)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(70), // Mint
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "CS Raspberry",
			ShortDescription: "Linear switch with a vibrant feel and responsive actuation.",
			LongDescription:  "The Akko CS Raspberry switch provides a vibrant linear feel with responsive actuation, making it an ideal choice for fast typists and gamers. Its consistent performance ensures reliable keystrokes, while its lively raspberry color adds a touch of personality to any keyboard. The Raspberry switch’s combination of quick response and eye-catching design makes it a favorite for users who value both speed and aesthetics.",
			ManufacturerID:   ptr(8), // Akko
			BrandID:          ptr(8), // Akko
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-07-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(1.9)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(1.9)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(57), // Pink
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "CS V2 Matcha Green",
			ShortDescription: "Linear switch with enhanced smoothness and stability.",
			LongDescription:  "The Akko CS V2 Matcha Green switch is an upgraded linear switch that offers enhanced smoothness and stability, providing a refined typing experience. This version improves upon the original Matcha Green by delivering even more consistent keystrokes with reduced wobble, making it ideal for both typing and gaming. The switch's signature green color and improved performance make it a must-have for enthusiasts who seek a premium linear switch.",
			ManufacturerID:   ptr(8), // Akko
			BrandID:          ptr(8), // Akko
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2022-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(1.9)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(1.9)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(54), // Green
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "CS Radiant Red Pro",
			ShortDescription: "Linear switch with a quick actuation and deep sound.",
			LongDescription:  "The Akko CS Radiant Red Pro switch is a linear switch that combines quick actuation with a deep, satisfying sound profile. Its fast response makes it perfect for both gamers and typists who require speed and precision in their keystrokes. The Radiant Red Pro's bold color and rich sound make it a standout choice for users who demand both performance and a unique auditory experience from their switches.",
			ManufacturerID:   ptr(8), // Akko
			BrandID:          ptr(8), // Akko
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2022-02-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(63.5)),
				ActuationPoint:        ptr(float32(1.9)),
				ActuationForce:        ptr(float32(63.5)),
				BottomOutForce:        ptr(float32(73.5)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(1.9)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "CS Silent Pink",
			ShortDescription: "Silent linear switch with a smooth and quiet operation.",
			LongDescription:  "The Akko CS Silent Pink switch is designed for those who require a silent typing experience without compromising on performance. Its smooth linear action and quiet operation make it ideal for office settings, shared spaces, and environments where noise is a concern. The Silent Pink switch's soft hue and near-silent keystrokes make it a perfect blend of form and function for users who prioritize discretion and quality.",
			ManufacturerID:   ptr(8), // Akko
			BrandID:          ptr(8), // Akko
			SwitchTypeID:     4,      // Silent Linear
			ReleaseDate:      parseDate("2021-09-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(1.9)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(3.5)),
				PreTravel:             ptr(float32(1.9)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
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
			Name:             "CS Silent Jelly Purple",
			ShortDescription: "Silent tactile switch with a subtle bump and quiet feedback.",
			LongDescription:  "The Akko CS Silent Jelly Purple switch offers a silent tactile experience with a subtle bump and quiet feedback. This switch is perfect for users who enjoy tactile typing without the accompanying noise, making it ideal for quiet environments. Its distinctive jelly-like color and smooth actuation provide both an aesthetic appeal and a satisfying typing feel, ensuring a seamless experience for users who need both tactility and silence.",
			ManufacturerID:   ptr(8), // Akko
			BrandID:          ptr(8), // Akko
			SwitchTypeID:     5,      // Silent Tactile
			ReleaseDate:      parseDate("2021-10-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(3.5)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(55), // Purple
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "CS Purple",
			ShortDescription: "Linear switch with a lightweight actuation and smooth feel.",
			LongDescription:  "The Akko CS Purple switch is a lightweight linear switch that offers a smooth and effortless typing experience. With its light actuation force, this switch is perfect for users who prefer quick and easy keystrokes, making it an excellent choice for both typing and gaming. The rich purple color adds a touch of elegance to any keyboard setup, while its consistent performance ensures that every keystroke is reliable and satisfying.",
			ManufacturerID:   ptr(8), // Akko
			BrandID:          ptr(8), // Akko
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2020-04-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(1.9)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(1.9)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(55), // Purple
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "V3 Cream Blue",
			ShortDescription: "Linear switch with enhanced smoothness and stability.",
			LongDescription:  "The Akko V3 Cream Blue switch is an advanced linear switch designed to offer enhanced smoothness and stability. This switch features a refined build that provides a premium typing experience, making it ideal for users who demand top-notch performance from their switches. The cool blue hue and consistent keystrokes make the Cream Blue switch a perfect choice for both typists and gamers who value quality and style.",
			ManufacturerID:   ptr(8), // Akko
			BrandID:          ptr(8), // Akko
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2022-03-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(1.9)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(1.9)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "CS Silent Forest Green",
			ShortDescription: "Silent linear switch with a medium actuation force.",
			LongDescription:  "The Akko CS Silent Forest Green switch combines a silent linear action with a medium actuation force, offering a balanced and quiet typing experience suitable for various applications. Its forest green color adds a natural, earthy tone to any keyboard, while its near-silent operation makes it perfect for environments where noise reduction is essential. This switch is a top pick for those who need both a smooth and silent typing experience.",
			ManufacturerID:   ptr(8), // Akko
			BrandID:          ptr(8), // Akko
			SwitchTypeID:     4,      // Silent Linear
			ReleaseDate:      parseDate("2021-11-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(1.9)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(3.5)),
				PreTravel:             ptr(float32(1.9)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(54), // Green
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "CS Cherry",
			ShortDescription: "Linear switch with a clear housing for RGB lighting.",
			LongDescription:  "The Akko CS Cherry switch is a linear switch featuring a clear housing that allows for vibrant RGB lighting, making it an excellent choice for users who want their keyboard to shine. Beyond its aesthetic appeal, the CS Cherry switch provides smooth keystrokes with a moderate actuation force, ensuring a comfortable typing experience. Whether for gaming or typing, this switch offers both style and performance.",
			ManufacturerID:   ptr(8), // Akko
			BrandID:          ptr(8), // Akko
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-12-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(1.9)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(1.9)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Kingfisher",
			ShortDescription: "Tactile switch with a vibrant feel and moderate bump.",
			LongDescription:  "The Akko Kingfisher switch is a tactile switch known for its vibrant feel and moderate tactile bump, providing a satisfying typing experience with each press. Inspired by the colorful Kingfisher bird, this switch adds a lively splash of color to any keyboard setup while delivering reliable tactile feedback. It's an excellent choice for those who appreciate a blend of aesthetics and performance in their typing experience.",
			ManufacturerID:   ptr(8), // Akko
			BrandID:          ptr(8), // Akko
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2020-02-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Sakura Pink",
			ShortDescription: "Linear switch with a soft touch and low actuation force.",
			LongDescription:  "The Akko Sakura Pink switch offers a soft linear touch with a low actuation force, making it ideal for users who prefer a gentle and effortless typing experience. Inspired by the delicate cherry blossoms, the Sakura Pink switch not only provides a smooth and quiet keystroke but also adds a beautiful and tranquil aesthetic to any keyboard. It's perfect for typists who value a light, responsive touch with a serene visual appeal.",
			ManufacturerID:   ptr(8), // Akko
			BrandID:          ptr(8), // Akko
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-08-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(35)),
				ActuationPoint:        ptr(float32(1.9)),
				ActuationForce:        ptr(float32(35)),
				BottomOutForce:        ptr(float32(45)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(1.9)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(57), // Pink
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Monet Purple",
			ShortDescription: "Specialty switch with artistic design and smooth feel.",
			LongDescription:  "The Akko Monet Purple switch is a specialty switch designed with an artistic touch, offering a smooth linear feel that mirrors the soft, flowing strokes of Monet's paintings. This switch is perfect for users who appreciate the combination of aesthetics and performance, delivering a premium typing experience with a unique visual appeal. Whether for typing or gaming, the Monet Purple switch adds an artistic flair to any keyboard.",
			ManufacturerID:   ptr(8), // Akko
			BrandID:          ptr(8), // Akko
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2022-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(1.9)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(1.9)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(55), // Purple
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "CS Piano",
			ShortDescription: "Linear switch with a smooth keystroke and balanced actuation force.",
			LongDescription:  "The Akko CS Piano switch offers a smooth linear feel with a balanced actuation force, providing a typing experience reminiscent of playing a finely tuned piano. Whether for typing or gaming, this switch delivers consistent and reliable performance, making it a versatile choice for various uses. Its elegant design and harmonious keystrokes make it an ideal switch for users who appreciate both form and function.",
			ManufacturerID:   ptr(8),                  // Akko
			BrandID:          ptr(8),                  // Akko
			SwitchTypeID:     1,                       // Linear
			ReleaseDate:      parseDate("2022-06-01"), // Approximate date, please verify
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(1.9)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(1.9)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
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
			Name:             "CS V3 Cream Yellow",
			ShortDescription: "Enhanced linear switch with improved smoothness and reduced wobble.",
			LongDescription:  "The Akko CS V3 Cream Yellow is an upgraded version of the popular Cream Yellow switch, featuring improved smoothness and reduced stem wobble. This switch offers a refined typing experience with a light actuation force, making it ideal for extended typing or gaming sessions. The V3 version enhances the already beloved qualities of the Cream Yellow switch, ensuring a premium and consistent feel with every keystroke.",
			ManufacturerID:   ptr(8),                  // Akko
			BrandID:          ptr(8),                  // Akko
			SwitchTypeID:     1,                       // Linear
			ReleaseDate:      parseDate("2022-09-01"), // Approximate date, please verify
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(1.9)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(1.9)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
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
			Name:             "CS Vintage White",
			ShortDescription: "Linear switch with a classic feel and moderate actuation force.",
			LongDescription:  "The Akko CS Vintage White switch offers a linear feel with a classic touch, providing a moderate actuation force that caters to users who prefer a balanced typing experience. Its vintage-inspired design adds a touch of nostalgia to modern keyboards, making it a versatile switch that blends the best of old and new. Whether for work or play, the Vintage White switch delivers a smooth and reliable performance with a timeless appeal.",
			ManufacturerID:   ptr(8),                  // Akko
			BrandID:          ptr(8),                  // Akko
			SwitchTypeID:     1,                       // Linear
			ReleaseDate:      parseDate("2022-07-01"), // Approximate date, please verify
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(1.9)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(1.9)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(51), // White
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "CS Maroon",
			ShortDescription: "Tactile switch with a deep, satisfying bump and rich color.",
			LongDescription:  "The Akko CS Maroon switch is a tactile switch that offers a deep and satisfying tactile bump. Its rich maroon color adds a touch of elegance to any keyboard, while its tactile feedback provides a gratifying typing experience for enthusiasts who enjoy pronounced feedback. The Maroon switch is perfect for those who value a strong tactile feel and a luxurious aesthetic in their typing experience.",
			ManufacturerID:   ptr(8),                  // Akko
			BrandID:          ptr(8),                  // Akko
			SwitchTypeID:     2,                       // Tactile
			ReleaseDate:      parseDate("2022-08-01"), // Approximate date, please verify
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(68), // Maroon
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "CS Fly Wing",
			ShortDescription: "Ultra-light linear switch for fast typing and gaming.",
			LongDescription:  "The Akko CS Fly Wing switch is an ultra-light linear switch designed for fast typing and gaming. Its low actuation force and smooth travel make it perfect for users who prioritize speed and responsiveness in their keyboard switches. The Fly Wing switch’s unique design and exceptional performance make it a standout choice for competitive gamers and speed typists alike.",
			ManufacturerID:   ptr(8),                  // Akko
			BrandID:          ptr(8),                  // Akko
			SwitchTypeID:     1,                       // Linear
			ReleaseDate:      parseDate("2022-10-01"), // Approximate date, please verify
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(35)),
				ActuationPoint:        ptr(float32(1.5)),
				ActuationForce:        ptr(float32(35)),
				BottomOutForce:        ptr(float32(45)),
				TotalTravel:           ptr(float32(3.5)),
				PreTravel:             ptr(float32(1.5)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(52), // Transparent
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "CS Jet Black",
			ShortDescription: "Linear switch with a sleek design and smooth keystroke.",
			LongDescription:  "The Akko CS Jet Black switch is a linear switch that combines a sleek, all-black design with a smooth keystroke. It offers a premium typing experience with a moderate actuation force, suitable for users who appreciate both aesthetics and performance in their keyboard switches. The Jet Black switch’s minimalist appearance and reliable performance make it a versatile option for various typing styles.",
			ManufacturerID:   ptr(8),                  // Akko
			BrandID:          ptr(8),                  // Akko
			SwitchTypeID:     1,                       // Linear
			ReleaseDate:      parseDate("2022-11-01"), // Approximate date, please verify
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(1.9)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(1.9)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
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
	}

	err := processSwitches(tx, akko, admin)
	if err != nil {
		return err
	}

	return nil
}
