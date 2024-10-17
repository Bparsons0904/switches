package seed

import (
	"switches/models"

	"gorm.io/gorm"
)

func seedCommonSwitches(tx *gorm.DB, admin models.User) error {
	commonSwitches := []models.Switch{
		{
			Name:             "K Pro Red",
			ShortDescription: "Linear switch with smooth keypresses.",
			LongDescription:  "The Keychron K Pro Red switch is a linear switch designed to deliver smooth and consistent keypresses, making it ideal for fast typists who require minimal resistance and swift actuation. Its quiet operation and reliable performance make it a versatile choice for both work and play.",
			ManufacturerID:   ptr(12), // Keychron
			BrandID:          ptr(12), // Keychron
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2022-03-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
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
			Name:             "K Pro Brown",
			ShortDescription: "Tactile switch for a responsive typing experience.",
			LongDescription:  "The Keychron K Pro Brown switch offers a tactile bump that provides a responsive and satisfying typing experience, making it well-suited for both typing and gaming. The balanced feedback and moderate actuation force offer a comfortable and efficient keystroke.",
			ManufacturerID:   ptr(12), // Keychron
			BrandID:          ptr(12), // Keychron
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2022-03-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(70)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(48), // Brown
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "K Pro Blue",
			ShortDescription: "Clicky switch with an audible click sound.",
			LongDescription:  "The Keychron K Pro Blue switch is known for its distinctive clicky sound and tactile feedback, making it a favorite for typists who enjoy a traditional mechanical keyboard feel. Its audible click and precise actuation deliver a satisfying and engaging typing experience.",
			ManufacturerID:   ptr(12), // Keychron
			BrandID:          ptr(12), // Keychron
			SwitchTypeID:     3,       // Clicky
			ReleaseDate:      parseDate("2022-03-01"),
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
			Name:             "Flamingo",
			ShortDescription: "Linear switch with smooth actuation.",
			LongDescription:  "The Epomaker Flamingo switch is a linear switch celebrated for its exceptionally smooth actuation and consistent keystroke. Its premium build and satisfying feel make it a favorite choice among mechanical keyboard enthusiasts looking for a high-quality linear experience.",
			ManufacturerID:   ptr(13), // Epomaker
			BrandID:          ptr(13), // Epomaker
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2021-11-01"),
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
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(57), // Pink
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Budgerigar",
			ShortDescription: "Tactile switch with a unique feel.",
			LongDescription:  "The Epomaker Budgerigar switch offers a distinct tactile bump that delivers a unique and pronounced feedback. Ideal for users who prefer a more tactile typing experience, this switch combines reliability with a satisfying keystroke.",
			ManufacturerID:   ptr(13), // Epomaker
			BrandID:          ptr(13), // Epomaker
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2021-11-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(67)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
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
			Name:             "Ice Candy",
			ShortDescription: "Smooth linear switch with a distinct color.",
			LongDescription:  "Epomaker Ice Candy switches provide a smooth linear typing experience complemented by a visually striking design. These switches are perfect for users who want both aesthetic appeal and consistent performance in their keyboard builds.",
			ManufacturerID:   ptr(13), // Epomaker
			BrandID:          ptr(13), // Epomaker
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2022-05-01"),
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
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
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
			Name:             "Blue Ice",
			ShortDescription: "Linear switch with smooth keypresses.",
			LongDescription:  "The Epomaker Blue Ice switch is a linear switch designed for fast typists, offering a smooth and consistent keypress. Its crisp actuation and unique color make it a popular choice for both gamers and typists who value performance and style.",
			ManufacturerID:   ptr(13), // Epomaker
			BrandID:          ptr(13), // Epomaker
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2022-03-01"),
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
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Green",
			ShortDescription: "Clicky switch with tactile feedback.",
			LongDescription:  "The Razer Green switch is a clicky switch that provides a distinct tactile bump along with an audible click, making it perfect for gamers and typists who enjoy a highly responsive and clicky typing experience. Its satisfying feedback is designed to enhance both gaming precision and typing accuracy.",
			ManufacturerID:   ptr(14), // Razer
			BrandID:          ptr(14), // Razer
			SwitchTypeID:     3,       // Clicky
			ReleaseDate:      parseDate("2014-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(1.9)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(70)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(1.9)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(54), // Green
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Yellow",
			ShortDescription: "Smooth linear switch with fast actuation.",
			LongDescription:  "The Razer Yellow switch is a linear switch engineered for ultra-fast actuation and smooth keypresses, making it ideal for competitive gaming. Its quiet operation and low actuation force allow for rapid key presses with minimal resistance, providing a seamless gaming experience.",
			ManufacturerID:   ptr(14), // Razer
			BrandID:          ptr(14), // Razer
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2016-10-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(1.2)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(3.5)),
				PreTravel:             ptr(float32(1.2)),
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
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Orange",
			ShortDescription: "Tactile switch with silent actuation.",
			LongDescription:  "The Razer Orange switch offers tactile feedback without the accompanying click noise, providing a silent yet responsive typing experience. This switch is ideal for users who prefer the feel of a tactile bump but require quieter operation, making it suitable for both work and gaming environments.",
			ManufacturerID:   ptr(14), // Razer
			BrandID:          ptr(14), // Razer
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2014-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(1.9)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(70)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(1.9)),
				FactoryLubed:          false,
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
			Name:             "GX Blue",
			ShortDescription: "Clicky switch with tactile feedback.",
			LongDescription:  "The Logitech GX Blue switch is a clicky switch that provides a distinctive tactile bump and audible click, delivering a satisfying typing experience for those who enjoy pronounced feedback. It's designed for users who appreciate both the sound and feel of a traditional mechanical keyboard.",
			ManufacturerID:   ptr(15), // Logitech
			BrandID:          ptr(15), // Logitech
			SwitchTypeID:     3,       // Clicky
			ReleaseDate:      parseDate("2019-05-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(3.7)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
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
			Name:             "GX Brown",
			ShortDescription: "Tactile switch with a quieter operation.",
			LongDescription:  "The Logitech GX Brown switch offers a tactile bump without the loud click, providing a more subtle yet responsive typing experience. Its quieter operation makes it suitable for both gaming and office use, striking a balance between feedback and noise control.",
			ManufacturerID:   ptr(15), // Logitech
			BrandID:          ptr(15), // Logitech
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2019-05-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(3.7)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(48), // Brown
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "GX Red",
			ShortDescription: "Linear switch for smooth keypresses.",
			LongDescription:  "The Logitech GX Red switch is a linear switch that delivers smooth and consistent keypresses, making it ideal for fast-paced gaming. Its low actuation force and quiet operation provide a seamless experience for gamers who prioritize speed and precision.",
			ManufacturerID:   ptr(15), // Logitech
			BrandID:          ptr(15), // Logitech
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2019-05-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(50)),
				TotalTravel:           ptr(float32(3.7)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
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
			Name:             "Titan Switch Tactile",
			ShortDescription: "Tactile switch with a crisp feel.",
			LongDescription:  "The Roccat Titan Switch Tactile is designed to deliver a crisp and responsive tactile feedback with each keystroke. This switch enhances typing precision and offers a satisfying bump, making it a great choice for both gaming and typing where accuracy is essential.",
			ManufacturerID:   ptr(19), // Roccat
			BrandID:          ptr(19), // Roccat
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2018-07-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(1.8)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(3.6)),
				PreTravel:             ptr(float32(1.8)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(48), // Brown
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Titan Switch Linear",
			ShortDescription: "Smooth linear switch for fast response.",
			LongDescription:  "The Roccat Titan Switch Linear offers smooth, consistent keystrokes with a fast response time, making it ideal for gamers who require seamless performance. Its low actuation force and quiet operation ensure a fluid typing experience, perfect for fast-paced gaming sessions.",
			ManufacturerID:   ptr(19), // Roccat
			BrandID:          ptr(19), // Roccat
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2018-07-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(1.8)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(3.6)),
				PreTravel:             ptr(float32(1.8)),
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
			Name:             "Green",
			ShortDescription: "Clicky switch with a tactile feel.",
			LongDescription:  "The Cooler Master Green switch is a clicky switch that provides a tactile bump with each keystroke. It offers an audible and tactile typing experience, making it a great choice for users who enjoy the traditional clicky feel and sound.",
			ManufacturerID:   ptr(20), // Cooler Master
			BrandID:          ptr(20), // Cooler Master
			SwitchTypeID:     3,       // Clicky
			ReleaseDate:      parseDate("2016-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(1.8)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(3.6)),
				PreTravel:             ptr(float32(1.8)),
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
			Name:             "Brown",
			ShortDescription: "Tactile switch with a quieter operation.",
			LongDescription:  "Cooler Master Brown switches provide a tactile bump with quieter operation, offering a balanced typing experience suitable for both work and play. These switches are ideal for users who prefer a tactile response without the loud click.",
			ManufacturerID:   ptr(20), // Cooler Master
			BrandID:          ptr(20), // Cooler Master
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2016-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(48), // Brown
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Red",
			ShortDescription: "Linear switch for fast keystrokes.",
			LongDescription:  "Cooler Master Red switches are linear switches designed for fast keystrokes and smooth typing, making them an excellent choice for gamers who require quick response times and a seamless typing experience.",
			ManufacturerID:   ptr(20), // Cooler Master
			BrandID:          ptr(20), // Cooler Master
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2016-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
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
			Name:             "Purple",
			ShortDescription: "Tactile switch with a pronounced bump.",
			LongDescription:  "Cooler Master Purple switches offer a tactile experience with a more pronounced bump compared to typical brown switches. These switches provide a stronger tactile feedback, making them well-suited for typing enthusiasts who prefer more noticeable keystrokes.",
			ManufacturerID:   ptr(20),                 // Cooler Master
			BrandID:          ptr(20),                 // Cooler Master
			SwitchTypeID:     2,                       // Tactile
			ReleaseDate:      parseDate("2017-01-01"), // Approximate, please verify
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
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
			Name:             "Blue",
			ShortDescription: "Clicky switch with tactile and audible feedback.",
			LongDescription:  "The Cooler Master Blue switch is known for providing both tactile and audible feedback with every keystroke. It delivers a satisfying click sound alongside a tactile bump, making it a favorite for those who enjoy a more interactive typing experience.",
			ManufacturerID:   ptr(20),                 // Cooler Master
			BrandID:          ptr(20),                 // Cooler Master
			SwitchTypeID:     3,                       // Clicky
			ReleaseDate:      parseDate("2017-01-01"), // Approximate, please verify
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				ActuationPoint:        ptr(float32(2.2)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.2)),
				FactoryLubed:          false,
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
			Name:             "Black",
			ShortDescription: "Smooth linear switch with deep sound.",
			LongDescription:  "JWK Black switches offer a premium linear experience with a smooth keystroke and deep, satisfying sound. These switches are ideal for enthusiasts who seek both performance and acoustics in their keyboard setup.",
			ManufacturerID:   ptr(22), // JWK
			BrandID:          ptr(22), // JWK
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2021-02-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(62)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(62)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Red",
			ShortDescription: "Smooth linear switch with light actuation.",
			LongDescription:  "JWK Red switches are linear switches known for their smooth actuation and light feel. These switches are perfect for fast typists who prefer a lighter keystroke for quick, responsive typing.",
			ManufacturerID:   ptr(22), // JWK
			BrandID:          ptr(22), // JWK
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2021-02-01"),
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
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Lavender",
			ShortDescription: "Linear switch with a unique color and smooth feel.",
			LongDescription:  "JWK Lavender switches offer a smooth linear typing experience with a distinctive lavender color. These switches are popular for their consistent feel and unique aesthetic appeal.",
			ManufacturerID:   ptr(22),                 // JWK
			BrandID:          ptr(22),                 // JWK
			SwitchTypeID:     1,                       // Linear
			ReleaseDate:      parseDate("2021-03-01"), // Approximate
			Available:        true,
			PricePoint:       3, // Expensive
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
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(63), // Lavender
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Alpaca",
			ShortDescription: "Smooth linear switch with a light touch.",
			LongDescription:  "JWK Alpaca switches are renowned for their smooth linear action and light touch. These switches are favored by both gamers and typists for their fluid keystrokes and consistent performance.",
			ManufacturerID:   ptr(22),                 // JWK
			BrandID:          ptr(22),                 // JWK
			SwitchTypeID:     1,                       // Linear
			ReleaseDate:      parseDate("2020-01-01"), // Approximate
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
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(51), // White
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Moss",
			ShortDescription: "Tactile switch with a rounded bump.",
			LongDescription:  "JWK Moss switches offer a tactile typing experience with a rounded bump that provides satisfying feedback without being overly sharp. These switches are great for users who prefer a more subtle tactile response.",
			ManufacturerID:   ptr(22),                 // JWK
			BrandID:          ptr(22),                 // JWK
			SwitchTypeID:     2,                       // Tactile
			ReleaseDate:      parseDate("2020-06-01"), // Approximate
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(67)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
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
			Name:             "Marshmallow",
			ShortDescription: "Smooth linear switch with a unique color.",
			LongDescription:  "JWK Marshmallow switches are known for their smooth linear feel and soft, marshmallow-inspired color. These switches are perfect for those who appreciate both performance and aesthetics in their keyboard setup.",
			ManufacturerID:   ptr(22),                 // JWK
			BrandID:          ptr(22),                 // JWK
			SwitchTypeID:     1,                       // Linear
			ReleaseDate:      parseDate("2020-09-01"), // Approximate
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
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(66), // Ivory
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Carrot",
			ShortDescription: "Tactile switch with a unique feel.",
			LongDescription:  "Tecsee Carrot switches offer a distinctive tactile experience with a crisp bump, providing a satisfying keystroke for those who enjoy pronounced tactile feedback.",
			ManufacturerID:   ptr(23), // Tecsee
			BrandID:          ptr(23), // Tecsee
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2020-11-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(63.5)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(63.5)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(56), // Orange
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Purple Panda",
			ShortDescription: "Tactile switch with a satisfying bump.",
			LongDescription:  "Tecsee Purple Panda switches provide a tactile experience with a satisfying and consistent bump, making them ideal for typists who appreciate a more pronounced feedback.",
			ManufacturerID:   ptr(23), // Tecsee
			BrandID:          ptr(23), // Tecsee
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2020-11-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(67)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(55), // Purple
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Aqua King",
			ShortDescription: "Smooth linear switch with a unique design.",
			LongDescription:  "Everglide Aqua King switches are known for their ultra-smooth linear action and unique transparent design, making them an excellent choice for both performance and aesthetics in custom keyboard builds.",
			ManufacturerID:   ptr(24), // Everglide
			BrandID:          ptr(24), // Everglide
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2021-08-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
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
			Name:             "Coral Red",
			ShortDescription: "Linear switch with a smooth feel.",
			LongDescription:  "Everglide Coral Red switches provide a smooth and consistent linear action with a striking red color, ideal for users seeking both performance and style.",
			ManufacturerID:   ptr(24), // Everglide
			BrandID:          ptr(24), // Everglide
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2021-08-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Purple",
			ShortDescription: "Tactile switch with a unique sound profile.",
			LongDescription:  "SP-Star Purple switches offer a distinct tactile feel combined with a unique sound profile, making them a favorite among enthusiasts looking for a different typing experience.",
			ManufacturerID:   ptr(25), // SP-Star
			BrandID:          ptr(25), // SP-Star
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2021-09-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(67)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(55), // Purple
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Yellow",
			ShortDescription: "Linear switch with smooth actuation.",
			LongDescription:  "SP-Star Yellow switches provide smooth linear actuation, making them an excellent choice for gamers and typists who prefer a seamless keystroke experience.",
			ManufacturerID:   ptr(25), // SP-Star
			BrandID:          ptr(25), // SP-Star
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2021-09-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(62)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(62)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
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
			Name:             "Boba U4",
			ShortDescription: "Silent tactile switch with a unique feel.",
			LongDescription:  "Gazzew Boba U4 switches are silent tactile switches that offer a smooth and quiet typing experience with a pronounced tactile bump, ideal for office use or shared spaces.",
			ManufacturerID:   ptr(26), // Gazzew
			BrandID:          ptr(26), // Gazzew
			SwitchTypeID:     5,       // Silent Tactile
			ReleaseDate:      parseDate("2019-11-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(68)),
				ActuationPoint:        ptr(float32(1.8)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(68)),
				TotalTravel:           ptr(float32(3.6)),
				PreTravel:             ptr(float32(1.8)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(51), // White
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Boba U4T",
			ShortDescription: "Tactile switch with a unique feel.",
			LongDescription:  "Gazzew Boba U4T switches provide a sharp and crisp tactile bump, offering a unique and satisfying typing experience for those who prefer tactile feedback with a distinct feel.",
			ManufacturerID:   ptr(26), // Gazzew
			BrandID:          ptr(26), // Gazzew
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2019-11-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(68)),
				ActuationPoint:        ptr(float32(1.8)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(68)),
				TotalTravel:           ptr(float32(3.6)),
				PreTravel:             ptr(float32(1.8)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(51), // White
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(31), // High
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Blue",
			ShortDescription: "Clicky switch with tactile feedback.",
			LongDescription:  "HMX Blue switches provide a clicky feel with tactile feedback, perfect for users who enjoy an audible typing experience.",
			ManufacturerID:   ptr(30), // HMX
			BrandID:          ptr(30), // HMX
			SwitchTypeID:     3,       // Clicky
			ReleaseDate:      parseDate("2018-07-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
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
			Name:             "Red",
			ShortDescription: "Linear switch for smooth keypresses.",
			LongDescription:  "HMX Red is a linear switch offering smooth keypresses and fast response times, suitable for both gaming and typing.",
			ManufacturerID:   ptr(30), // HMX
			BrandID:          ptr(30), // HMX
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2018-07-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
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
			Name:             "Brown",
			ShortDescription: "Tactile switch with a subtle bump.",
			LongDescription:  "HMX Brown switches offer a tactile bump with quieter operation, ideal for users who prefer a balance between tactile feedback and noise.",
			ManufacturerID:   ptr(30), // HMX
			BrandID:          ptr(30), // HMX
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2018-07-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(48), // Brown
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Silent Red",
			ShortDescription: "Silent linear switch with smooth actuation.",
			LongDescription:  "HMX Silent Red switches provide a smooth linear experience without the noise, perfect for quiet environments.",
			ManufacturerID:   ptr(30), // HMX
			BrandID:          ptr(30), // HMX
			SwitchTypeID:     4,       // Silent Linear
			ReleaseDate:      parseDate("2019-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(3.7)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Black",
			ShortDescription: "Heavy linear switch with deep actuation.",
			LongDescription:  "HMX Black switches offer a heavy linear feel with deep actuation, suitable for typists who prefer a more substantial keystroke.",
			ManufacturerID:   ptr(30), // HMX
			BrandID:          ptr(30), // HMX
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2019-06-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(70)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(60)),
				BottomOutForce:        ptr(float32(70)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Silent Black",
			ShortDescription: "Silent heavy linear switch.",
			LongDescription:  "HMX Silent Black switches provide a heavy linear feel with silent actuation, ideal for those who prefer a quiet typing experience with substantial feedback.",
			ManufacturerID:   ptr(30), // HMX
			BrandID:          ptr(30), // HMX
			SwitchTypeID:     4,       // Silent Linear
			ReleaseDate:      parseDate("2020-02-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(70)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(60)),
				BottomOutForce:        ptr(float32(70)),
				TotalTravel:           ptr(float32(3.7)),
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
			Name:             "Green",
			ShortDescription: "Clicky switch with strong tactile feedback.",
			LongDescription:  "HMX Green switches offer a strong clicky feel with pronounced tactile feedback, perfect for those who enjoy a loud typing experience.",
			ManufacturerID:   ptr(30), // HMX
			BrandID:          ptr(30), // HMX
			SwitchTypeID:     3,       // Clicky
			ReleaseDate:      parseDate("2019-11-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(80)),
				ActuationPoint:        ptr(float32(2.2)),
				ActuationForce:        ptr(float32(65)),
				BottomOutForce:        ptr(float32(80)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.2)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(54), // Green
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Clear",
			ShortDescription: "Tactile switch with medium resistance.",
			LongDescription:  "HMX Clear switches offer medium resistance with tactile feedback, ideal for those who prefer a balanced typing feel.",
			ManufacturerID:   ptr(30), // HMX
			BrandID:          ptr(30), // HMX
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2020-08-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(49), // Clear
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Speed Silver",
			ShortDescription: "Fast linear switch for quick typing.",
			LongDescription:  "HMX Speed Silver switches are designed for fast actuation and quick typing, making them ideal for gamers seeking speed and efficiency.",
			ManufacturerID:   ptr(30), // HMX
			BrandID:          ptr(30), // HMX
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2021-03-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(1.2)),
				ActuationForce:        ptr(float32(40)),
				BottomOutForce:        ptr(float32(45)),
				TotalTravel:           ptr(float32(3.4)),
				PreTravel:             ptr(float32(1.2)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
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
			Name:             "Pink",
			ShortDescription: "Linear switch with a soft touch.",
			LongDescription:  "HMX Pink switches offer a linear feel with a soft touch, suitable for those who prefer a gentle typing experience.",
			ManufacturerID:   ptr(30), // HMX
			BrandID:          ptr(30), // HMX
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2021-07-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(40)),
				ActuationPoint:        ptr(float32(1.9)),
				ActuationForce:        ptr(float32(35)),
				BottomOutForce:        ptr(float32(40)),
				TotalTravel:           ptr(float32(3.8)),
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
			Name:             "Xinhai",
			ShortDescription: "Tactile switch with a soft bump.",
			LongDescription:  "HMX Xinhai switches provide a soft tactile bump with smooth actuation, offering a pleasant typing experience for both gaming and typing.",
			ManufacturerID:   ptr(30), // HMX
			BrandID:          ptr(30), // HMX
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2021-12-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(1.8)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(3.6)),
				PreTravel:             ptr(float32(1.8)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(51), // White
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Purple Dawn",
			ShortDescription: "Silent tactile switch with unique feedback.",
			LongDescription:  "HMX Purple Dawn switches offer silent tactile feedback with a unique actuation profile, perfect for those who enjoy a quiet yet tactile experience.",
			ManufacturerID:   ptr(30), // HMX
			BrandID:          ptr(30), // HMX
			SwitchTypeID:     5,       // Silent Tactile
			ReleaseDate:      parseDate("2022-03-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				ActuationPoint:        ptr(float32(1.9)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(3.7)),
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
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Cloud White",
			ShortDescription: "Linear switch with a soft feel.",
			LongDescription:  "HMX Cloud White switches provide a smooth linear feel with a soft actuation, offering a gentle and pleasant typing experience.",
			ManufacturerID:   ptr(30), // HMX
			BrandID:          ptr(30), // HMX
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2022-05-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(1.9)),
				ActuationForce:        ptr(float32(40)),
				BottomOutForce:        ptr(float32(45)),
				TotalTravel:           ptr(float32(3.8)),
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
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Coral",
			ShortDescription: "Tactile switch with a colorful design.",
			LongDescription:  "HMX Coral switches offer a tactile experience with a visually striking design, making them a great choice for custom keyboard builds.",
			ManufacturerID:   ptr(30), // HMX
			BrandID:          ptr(30), // HMX
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2022-09-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(58)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(58)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(67), // Coral
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Gold",
			ShortDescription: "Clicky switch with a premium feel.",
			LongDescription:  "HMX Gold switches provide a clicky feel with premium materials, offering a luxurious typing experience with audible feedback.",
			ManufacturerID:   ptr(30), // HMX
			BrandID:          ptr(30), // HMX
			SwitchTypeID:     3,       // Clicky
			ReleaseDate:      parseDate("2023-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(60), // Gold
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Poro",
			ShortDescription: "Linear switch with a smooth and silent action.",
			LongDescription:  "Poro switches offer a unique linear feel with a smooth and silent action, providing a premium typing experience for those who prefer quieter keyboards.",
			ManufacturerID:   ptr(30), // HMX
			BrandID:          ptr(30), // HMX
			SwitchTypeID:     4,       // Silent Linear
			ReleaseDate:      parseDate("2023-03-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(1.8)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(50)),
				TotalTravel:           ptr(float32(3.6)),
				PreTravel:             ptr(float32(1.8)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(51), // White
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Hyacinth V2",
			ShortDescription: "Tactile switch with a medium bump.",
			LongDescription:  "Hyacinth V2 switches provide a medium tactile bump with a smooth actuation, making them ideal for typists who enjoy a balanced tactile feel.",
			ManufacturerID:   ptr(30), // HMX
			BrandID:          ptr(30), // HMX
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2023-06-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(62)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(62)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(57), // Pink
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Hyacinth V2U",
			ShortDescription: "Enhanced tactile switch with a unique actuation profile.",
			LongDescription:  "Hyacinth V2U switches offer an enhanced tactile experience with a unique actuation profile, providing a distinctive and satisfying typing experience for enthusiasts.",
			ManufacturerID:   ptr(30), // HMX
			BrandID:          ptr(30), // HMX
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2023-09-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(1.9)),
				ActuationForce:        ptr(float32(58)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(3.8)),
				PreTravel:             ptr(float32(1.9)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(63), // Lavender
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "KBDFans T1",
			ShortDescription: "Tactile switch with a distinct bump.",
			LongDescription:  "KBDFans T1 switches are known for their strong tactile bump and smooth actuation, offering a unique typing experience for those who prefer a pronounced tactile feedback.",
			ManufacturerID:   ptr(29), // KBDFans
			BrandID:          ptr(29), // KBDFans
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2019-06-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(67)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(48), // Brown
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "KBDFans D65",
			ShortDescription: "Linear switch with smooth keypress.",
			LongDescription:  "KBDFans D65 offers a linear experience with smooth keypresses, ideal for gamers and fast typists seeking consistency in their keystrokes.",
			ManufacturerID:   ptr(29), // KBDFans
			BrandID:          ptr(29), // KBDFans
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2021-04-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(62)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(62)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "KBDFans Kiwis",
			ShortDescription: "Tactile switch with a smooth feel.",
			LongDescription:  "KBDFans Kiwis are tactile switches known for their smooth feel and strong tactile bump, providing an enjoyable and responsive typing experience.",
			ManufacturerID:   ptr(29), // KBDFans
			BrandID:          ptr(29), // KBDFans
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2020-11-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(67)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
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
			Name:             "KBDFans Mint",
			ShortDescription: "Linear switch with a refreshing color.",
			LongDescription:  "KBDFans Mint switches provide a smooth linear feel with a refreshing mint color, perfect for custom keyboard builds and enthusiasts looking for a unique aesthetic.",
			ManufacturerID:   ptr(29), // KBDFans
			BrandID:          ptr(29), // KBDFans
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2021-08-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(62)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(62)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(70), // Mint
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "KBDFans Crystal",
			ShortDescription: "Silent linear switch with a smooth feel.",
			LongDescription:  "KBDFans Crystal switches offer a silent linear feel, providing a quiet yet smooth typing experience ideal for office use and shared spaces.",
			ManufacturerID:   ptr(29), // KBDFans
			BrandID:          ptr(29), // KBDFans
			SwitchTypeID:     4,       // Silent Linear
			ReleaseDate:      parseDate("2021-09-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(62)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(62)),
				TotalTravel:           ptr(float32(3.8)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
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
			Name:             "KBDFans Roller",
			ShortDescription: "Innovative switch with rolling actuation.",
			LongDescription:  "KBDFans Roller switches offer a unique rolling actuation mechanism, providing a novel tactile feedback and sound profile for a different typing experience.",
			ManufacturerID:   ptr(29), // KBDFans
			BrandID:          ptr(29), // KBDFans
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2023-05-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(51), // White
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Kinetic Labs Hippo",
			ShortDescription: "Tactile switch with a medium actuation force.",
			LongDescription:  "The Kinetic Labs Hippo switch delivers a delightful tactile bump combined with an exceptionally smooth downstroke and upstroke. With a medium actuation force of 67g, this switch is versatile for both typing and gaming, offering a satisfying and consistent tactile experience. Its durable construction and unique color scheme make it stand out in any build.",
			ManufacturerID:   ptr(17), // Kinetic Labs
			BrandID:          ptr(17), // Kinetic Labs
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2021-03-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(67)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(58), // Gray
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Kinetic Labs Salmon",
			ShortDescription: "Linear switch with a smooth, medium-weight feel.",
			LongDescription:  "Kinetic Labs Salmon switches are renowned for their buttery-smooth linear action and consistent keypresses. With a medium weighting of 63.5g, these switches offer a perfect balance between lightness and control, making them ideal for users who enjoy a gentle yet responsive typing experience. Their salmon-pink color adds a touch of elegance to any keyboard.",
			ManufacturerID:   ptr(17), // Kinetic Labs
			BrandID:          ptr(17), // Kinetic Labs
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2021-06-15"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(63.5)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(63.5)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(67), // Coral
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Kinetic Labs Snow Globe V2",
			ShortDescription: "Silent linear switch with a smooth and quiet operation.",
			LongDescription:  "The Kinetic Labs Snow Globe V2 is a silent linear switch that combines exceptional smoothness with near-silent operation, making it perfect for office environments or shared spaces. With a 63.5g actuation force, it delivers a satisfying and smooth keystroke while minimizing noise, allowing for distraction-free typing. The clear housing adds a frosty aesthetic that complements RGB lighting.",
			ManufacturerID:   ptr(17), // Kinetic Labs
			BrandID:          ptr(17), // Kinetic Labs
			SwitchTypeID:     4,       // Silent Linear
			ReleaseDate:      parseDate("2022-01-20"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(63.5)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(63.5)),
				TotalTravel:           ptr(float32(3.8)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
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
			Name:             "Kinetic Labs Polar Panda",
			ShortDescription: "Tactile switch with a distinct and pronounced bump.",
			LongDescription:  "Kinetic Labs Polar Panda switches are engineered for enthusiasts who crave a highly tactile typing experience. Featuring a pronounced bump with a 67g actuation force, these switches provide immediate feedback with each press, making them perfect for users who value tactile response. The unique white and black color scheme adds visual appeal to any keyboard.",
			ManufacturerID:   ptr(22), // JWK
			BrandID:          ptr(17), // Kinetic Labs
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2022-08-10"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(67)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(51), // White
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Kinetic Labs Dragon Fruit",
			ShortDescription: "Linear switch with a medium actuation force and unique aesthetic.",
			LongDescription:  "The Kinetic Labs Dragon Fruit switch offers a smooth linear experience with a medium actuation force of 67g. Its vibrant pink and green color scheme draws inspiration from the tropical fruit, making it a standout choice for custom builds. The switch’s consistent feel and solid performance make it suitable for both gaming and daily typing.",
			ManufacturerID:   ptr(17), // Kinetic Labs
			BrandID:          ptr(17), // Kinetic Labs
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2023-05-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(67)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(57), // Pink
				TopHousingColorID:     ptr(54), // Green
				BottomHousingColorID:  ptr(54), // Green
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Kinetic Labs Blueberry",
			ShortDescription: "Tactile switch with a medium-weight feel and crisp feedback.",
			LongDescription:  "Kinetic Labs Blueberry switches deliver crisp tactile feedback with a medium-weight actuation force of 55g. Ideal for typists who enjoy a sharp and responsive typing feel, these switches offer a satisfying tactile bump without being overly stiff. The deep blue housing adds a cool and refreshing aesthetic to any keyboard setup.",
			ManufacturerID:   ptr(17), // Kinetic Labs
			BrandID:          ptr(17), // Kinetic Labs
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2023-07-15"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
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
			Name:             "Kinetic Labs Penguin",
			ShortDescription: "Silent linear switch with smooth and quiet keystrokes.",
			LongDescription:  "The Kinetic Labs Penguin switch is designed for those who prioritize a quiet typing experience. With a 62g actuation force, it provides a smooth and silent keystroke, making it perfect for noise-sensitive environments. The switch features a sleek black and white design, reminiscent of a penguin, adding a playful yet professional touch to your keyboard.",
			ManufacturerID:   ptr(17), // Kinetic Labs
			BrandID:          ptr(17), // Kinetic Labs
			SwitchTypeID:     4,       // Silent Linear
			ReleaseDate:      parseDate("2023-02-10"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(62)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(62)),
				TotalTravel:           ptr(float32(3.8)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(51), // White
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Kinetic Labs Gecko Silent Linear",
			ShortDescription: "Ultra-silent linear switch with soft actuation and muted sound.",
			LongDescription:  "Kinetic Labs Gecko Silent Linear switches are ultra-silent switches designed for users who demand quietness. With a soft 50g actuation force, they offer an exceptionally smooth and silent typing experience, making them ideal for late-night work or shared living spaces. The muted green housing enhances the aesthetic appeal, adding a subtle touch of nature to your keyboard.",
			ManufacturerID:   ptr(17), // Kinetic Labs
			BrandID:          ptr(17), // Kinetic Labs
			SwitchTypeID:     4,       // Silent Linear
			ReleaseDate:      parseDate("2023-05-25"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(50)),
				TotalTravel:           ptr(float32(3.8)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(54), // Green
				TopHousingColorID:     ptr(54), // Green
				BottomHousingColorID:  ptr(54), // Green
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Chosfox Kiwi",
			ShortDescription: "Tactile switch with a strong bump and satisfying feedback.",
			LongDescription:  "The Chosfox Kiwi switch is known for its pronounced tactile bump, providing satisfying feedback for those who enjoy a strong, crisp actuation. With a 67g actuation force, it delivers a consistent and engaging typing experience, making it a favorite among tactile switch enthusiasts.",
			ManufacturerID:   ptr(22), // JWK
			BrandID:          ptr(18), // Chosfox
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2021-06-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(67)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
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
			Name:             "Chosfox Mint",
			ShortDescription: "Linear switch with a smooth and creamy feel.",
			LongDescription:  "The Chosfox Mint switch offers a smooth and creamy linear feel, perfect for users who enjoy a seamless keystroke. With a 63.5g actuation force, it provides a balanced typing experience, combining lightness with stability. Its cool minty color adds a refreshing aesthetic to any keyboard setup.",
			ManufacturerID:   ptr(22), // JWK
			BrandID:          ptr(18), // Chosfox
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2022-02-15"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(63.5)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(63.5)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(70), // Mint
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Chosfox Blueberry",
			ShortDescription: "Tactile switch with a medium-weight feel and crisp feedback.",
			LongDescription:  "Chosfox Blueberry switches provide a medium-weight tactile experience, delivering crisp feedback with every press. Featuring a 65g actuation force, they are ideal for typists who appreciate tactile sensations without excessive resistance. The deep blue color adds a visually pleasing touch to your build.",
			ManufacturerID:   ptr(22), // JWK
			BrandID:          ptr(18), // Chosfox
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2022-07-10"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
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
			Name:             "Chosfox Chameleon",
			ShortDescription: "Linear switch with a unique aesthetic and smooth operation.",
			LongDescription:  "The Chosfox Chameleon switch stands out with its unique aesthetic and smooth linear action. With a 68g actuation force, it provides a consistent and fluid typing experience, perfect for those who prefer a steady keystroke. Its vibrant, color-shifting design adds a dynamic visual element to any keyboard.",
			ManufacturerID:   ptr(22), // JWK
			BrandID:          ptr(18), // Chosfox
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2023-01-20"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(68)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(68)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
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
			Name:             "Chosfox Marshmallow",
			ShortDescription: "Silent linear switch with a soft and quiet typing experience.",
			LongDescription:  "Chosfox Marshmallow switches are designed for a silent and smooth linear typing experience. With a soft 58g actuation force, they are perfect for users who need a quiet environment without compromising on the feel. The soft pink color and smooth operation make them a delight for both eyes and fingers.",
			ManufacturerID:   ptr(22), // JWK
			BrandID:          ptr(18), // Chosfox
			SwitchTypeID:     4,       // Silent Linear
			ReleaseDate:      parseDate("2023-06-05"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(58)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(58)),
				TotalTravel:           ptr(float32(3.8)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(57), // Pink
				TopHousingColorID:     ptr(66), // Ivory
				BottomHousingColorID:  ptr(66), // Ivory
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Chosfox Watermelon",
			ShortDescription: "Tactile switch with a fruity feel and moderate resistance.",
			LongDescription:  "The Chosfox Watermelon switch offers a unique tactile experience with a fruity theme and moderate resistance. Featuring a 62g actuation force, it provides a satisfying bump with each keystroke, making it perfect for those who enjoy a playful yet responsive typing feel.",
			ManufacturerID:   ptr(22), // JWK
			BrandID:          ptr(18), // Chosfox
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2023-09-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(62)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(62)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
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
			Name:             "Chosfox Grape",
			ShortDescription: "Linear switch with a heavier actuation force and smooth travel.",
			LongDescription:  "Chosfox Grape switches offer a smooth linear action with a substantial 70g actuation force. These switches are ideal for users who prefer a heavier keystroke, providing a stable and controlled typing experience. The rich purple color adds a bold statement to any keyboard build.",
			ManufacturerID:   ptr(22), // JWK
			BrandID:          ptr(18), // Chosfox
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2024-01-10"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(70)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(60)),
				BottomOutForce:        ptr(float32(70)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(55), // Purple
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Chosfox Silent Forest",
			ShortDescription: "Silent tactile switch with a gentle bump and quiet operation.",
			LongDescription:  "The Chosfox Silent Forest switch is perfect for those who want tactile feedback without the noise. With a 62g actuation force, it offers a gentle bump and quiet operation, making it ideal for shared spaces or office environments. The green housing evokes a serene forest atmosphere, adding a calming touch to your keyboard.",
			ManufacturerID:   ptr(22), // JWK
			BrandID:          ptr(18), // Chosfox
			SwitchTypeID:     5,       // Silent Tactile
			ReleaseDate:      parseDate("2024-03-20"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(62)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(62)),
				TotalTravel:           ptr(float32(3.8)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(54), // Green
				TopHousingColorID:     ptr(54), // Green
				BottomHousingColorID:  ptr(54), // Green
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Chosfox Banana Split",
			ShortDescription: "Linear switch with a smooth feel and playful aesthetic.",
			LongDescription:  "Chosfox Banana Split switches offer a smooth and seamless linear experience, perfect for users who enjoy a fun and reliable keystroke. With a 62g actuation force, these switches combine a playful aesthetic with consistent performance, making them a popular choice for custom keyboard enthusiasts.",
			ManufacturerID:   ptr(22), // JWK
			BrandID:          ptr(18), // Chosfox
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2024-05-15"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(62)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(62)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
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
			Name:             "Chosfox Sunset",
			ShortDescription: "Linear switch with a vibrant color scheme and consistent action.",
			LongDescription:  "The Chosfox Sunset switch features a vibrant color scheme inspired by the hues of a sunset, paired with consistent linear action. With a 67g actuation force, it provides a reliable typing experience that is both visually appealing and satisfying to use.",
			ManufacturerID:   ptr(22), // JWK
			BrandID:          ptr(18), // Chosfox
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2024-07-10"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(67)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(56), // Orange
				TopHousingColorID:     ptr(45), // Red
				BottomHousingColorID:  ptr(50), // Yellow
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Chosfox White Fox",
			ShortDescription: "Tactile switch with a smooth bump and quiet operation.",
			LongDescription:  "Chosfox White Fox switches are designed to deliver a smooth tactile bump with minimal noise. Featuring a 58g actuation force, these switches are perfect for those who enjoy a soft tactile feel while maintaining a quiet typing environment. The clean white housing adds a sophisticated touch to any keyboard.",
			ManufacturerID:   ptr(3),  // Kailh
			BrandID:          ptr(18), // Chosfox
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2023-01-15"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(58)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(48)),
				BottomOutForce:        ptr(float32(58)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(51), // White
				TopHousingColorID:     ptr(51), // White
				BottomHousingColorID:  ptr(51), // White
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Chosfox Arctic Fox",
			ShortDescription: "Linear switch with a cool feel and light actuation force.",
			LongDescription:  "Chosfox Arctic Fox switches provide a cool and smooth linear feel with a light 48g actuation force. These switches are perfect for users who prefer a gentle keystroke, offering a crisp and clean typing experience. The frosty white color complements any winter-themed build.",
			ManufacturerID:   ptr(3),  // Kailh
			BrandID:          ptr(18), // Chosfox
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2023-02-10"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(48)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(40)),
				BottomOutForce:        ptr(float32(48)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(51), // White
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Chosfox Voyager V2",
			ShortDescription: "Linear switch designed for smooth and consistent travel.",
			LongDescription:  "Chosfox Voyager V2 switches are crafted for a smooth and consistent linear travel with a 65g actuation force. These switches provide a balanced and fluid typing experience, ideal for users who seek precision and reliability in their keystrokes.",
			ManufacturerID:   ptr(22), // JWK
			BrandID:          ptr(18), // Chosfox
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2023-04-20"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
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
			Name:             "Chosfox Hanami Dango",
			ShortDescription: "Linear switch with soft feel and pastel aesthetics.",
			LongDescription:  "Chosfox Hanami Dango switches are inspired by the delicate colors of Japanese sweets, offering a soft linear feel with a 50g actuation force. These switches are perfect for users who value a gentle touch and pleasing visuals, making them a delightful addition to any pastel-themed keyboard.",
			ManufacturerID:   ptr(22), // JWK
			BrandID:          ptr(18), // Chosfox
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2023-05-15"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(40)),
				BottomOutForce:        ptr(float32(50)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(57), // Pink
				TopHousingColorID:     ptr(54), // Green
				BottomHousingColorID:  ptr(51), // White
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Chosfox Summer Lime Silent",
			ShortDescription: "Silent linear switch with a soft and quiet typing experience.",
			LongDescription:  "Chosfox Summer Lime Silent switches provide a noiseless linear experience with a soft and quiet 40g actuation force. Ideal for users who need a silent typing environment, these switches offer a fresh lime-green aesthetic that adds a pop of color to any build.",
			ManufacturerID:   ptr(22), // JWK
			BrandID:          ptr(18), // Chosfox
			SwitchTypeID:     4,       // Silent Linear
			ReleaseDate:      parseDate("2023-07-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(40)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(35)),
				BottomOutForce:        ptr(float32(40)),
				TotalTravel:           ptr(float32(3.8)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(54), // Green
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Chosfox Poison Gas",
			ShortDescription: "Tactile switch with strong feedback and vibrant color scheme.",
			LongDescription:  "Chosfox Poison Gas switches offer strong tactile feedback with a vibrant, eye-catching color scheme. Featuring a 67g actuation force, these switches are perfect for users who enjoy pronounced tactile sensations combined with bold aesthetics.",
			ManufacturerID:   ptr(22), // JWK
			BrandID:          ptr(18), // Chosfox
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2023-08-10"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(67)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(54), // Green
				TopHousingColorID:     ptr(55), // Purple
				BottomHousingColorID:  ptr(55), // Purple
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Chosfox DD Jingle Linear",
			ShortDescription: "Linear switch with a unique auditory profile and smooth action.",
			LongDescription:  "Chosfox DD Jingle Linear switches combine a unique auditory profile with smooth linear action. With a 55g actuation force, they provide a distinct sound and feel, making them perfect for users who appreciate both performance and aural feedback in their keystrokes.",
			ManufacturerID:   ptr(22), // JWK
			BrandID:          ptr(18), // Chosfox
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2023-09-05"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(54), // Green
				BottomHousingColorID:  ptr(54), // Green
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Wuque Studio Aurora",
			ShortDescription: "Linear switch with a smooth feel and vibrant color scheme.",
			LongDescription:  "The Wuque Studio Aurora switch delivers a smooth linear typing experience with a 63.5g actuation force. Its vibrant color scheme adds visual appeal, making it a great choice for users who prioritize both aesthetics and performance in their keyboard setups.",
			ManufacturerID:   ptr(22), // JWK
			BrandID:          ptr(21), // Wuque Studio
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2022-04-01"),
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
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(55), // Purple
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Wuque Studio Polar",
			ShortDescription: "Tactile switch with a strong bump and icy aesthetic.",
			LongDescription:  "The Wuque Studio Polar switch offers a pronounced tactile bump with a 67g actuation force, delivering a satisfying and consistent feedback. The icy aesthetic complements its crisp feel, making it a top choice for tactile switch enthusiasts seeking both performance and style.",
			ManufacturerID:   ptr(22), // JWK
			BrandID:          ptr(21), // Wuque Studio
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2022-07-15"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(67)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Wuque Studio Iceberg",
			ShortDescription: "Silent linear switch with smooth and quiet operation.",
			LongDescription:  "Wuque Studio Iceberg switches provide an ultra-smooth and silent linear typing experience. With a 60g actuation force, these switches are perfect for those who need a noiseless environment without sacrificing the fluidity of their keystrokes.",
			ManufacturerID:   ptr(22), // JWK
			BrandID:          ptr(21), // Wuque Studio
			SwitchTypeID:     4,       // Silent Linear
			ReleaseDate:      parseDate("2023-01-20"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(3.8)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Wuque Studio Ocean Blue",
			ShortDescription: "Clicky switch with crisp feedback and oceanic feel.",
			LongDescription:  "The Wuque Studio Ocean Blue switch offers a crisp clicky experience with a 50g actuation force, providing tactile and audible feedback that mimics the refreshing vibe of ocean waves. Ideal for users who enjoy a lively, responsive typing feel.",
			ManufacturerID:   ptr(22), // JWK
			BrandID:          ptr(21), // Wuque Studio
			SwitchTypeID:     3,       // Clicky
			ReleaseDate:      parseDate("2023-05-10"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(50)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Wuque Studio Forest Green",
			ShortDescription: "Tactile switch with a natural feel and moderate resistance.",
			LongDescription:  "Wuque Studio Forest Green switches are designed to evoke the tranquility of nature, offering a moderate 62g actuation force with a natural tactile feel. These switches are perfect for users who appreciate a blend of tactile feedback and organic aesthetics.",
			ManufacturerID:   ptr(22), // JWK
			BrandID:          ptr(21), // Wuque Studio
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2023-10-05"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(62)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(62)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
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
			Name:             "Wuque Studio Jade Linear",
			ShortDescription: "Linear switch with a smooth and consistent action.",
			LongDescription:  "Wuque Studio Jade Linear switches provide a smooth, consistent linear action with a light 45g actuation force. These switches are ideal for users who prefer a gentle keystroke, making them perfect for both typing and gaming.",
			ManufacturerID:   ptr(21), // Wuque Studio
			BrandID:          ptr(21), // Wuque Studio
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2023-03-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(40)),
				BottomOutForce:        ptr(float32(45)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(62), // Teal
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Wuque Studio Heavy Tactile",
			ShortDescription: "Tactile switch with a strong bump and substantial resistance.",
			LongDescription:  "Wuque Studio Heavy Tactile switches are designed for users who crave a strong tactile bump with substantial resistance. With a 70g actuation force, these switches deliver a pronounced feedback that’s perfect for those who enjoy a robust tactile experience.",
			ManufacturerID:   ptr(21), // Wuque Studio
			BrandID:          ptr(21), // Wuque Studio
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2023-04-15"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(70)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(60)),
				BottomOutForce:        ptr(float32(70)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(31), // High
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Wuque Studio Aurora Linear",
			ShortDescription: "Smooth linear switch with vibrant aesthetics.",
			LongDescription:  "Wuque Studio Aurora Linear switches offer a vibrant aesthetic paired with smooth and consistent linear action. Featuring a 63.5g actuation force, these switches provide a visually appealing and fluid typing experience, perfect for custom keyboard builds.",
			ManufacturerID:   ptr(21), // Wuque Studio
			BrandID:          ptr(21), // Wuque Studio
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2023-06-01"),
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
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(55), // Purple
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Wuque Studio WS Silent",
			ShortDescription: "Ultra-silent linear switch with muted sound.",
			LongDescription:  "Wuque Studio WS Silent switches provide an ultra-silent linear experience, designed for those who need a quiet typing environment. With a 60g actuation force and muted sound, these switches are perfect for workspaces and late-night sessions.",
			ManufacturerID:   ptr(21), // Wuque Studio
			BrandID:          ptr(21), // Wuque Studio
			SwitchTypeID:     4,       // Silent Linear
			ReleaseDate:      parseDate("2023-07-10"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(3.8)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Wuque Studio POM+",
			ShortDescription: "Linear switch with POM material for smooth keystrokes.",
			LongDescription:  "Wuque Studio POM+ switches are crafted from high-quality POM material, delivering exceptionally smooth linear keystrokes. With a 55g actuation force, these switches offer a seamless typing experience, ideal for users seeking premium performance and durability.",
			ManufacturerID:   ptr(21), // Wuque Studio
			BrandID:          ptr(21), // Wuque Studio
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2023-09-15"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(23), // POM
				BaseHousingMaterialID: ptr(23), // POM
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(51), // White
				TopHousingColorID:     ptr(51), // White
				BottomHousingColorID:  ptr(51), // White
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Wuque Studio Light Tactile",
			ShortDescription: "Tactile switch with a gentle bump and light resistance.",
			LongDescription:  "Wuque Studio Light Tactile switches offer a subtle tactile feedback with light resistance, featuring a 50g actuation force. Perfect for users who prefer a less pronounced bump, these switches deliver a smooth and responsive typing experience.",
			ManufacturerID:   ptr(21), // Wuque Studio
			BrandID:          ptr(21), // Wuque Studio
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2023-10-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(50)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(51), // White
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Wuque Studio Onion Linear",
			ShortDescription: "Linear switch with smooth travel and a unique aesthetic.",
			LongDescription:  "Wuque Studio Onion Linear switches combine smooth travel with a distinctive aesthetic. With a 52g actuation force, these switches offer a light, fluid typing experience, making them perfect for users who value both performance and a unique visual appeal.",
			ManufacturerID:   ptr(21), // Wuque Studio
			BrandID:          ptr(21), // Wuque Studio
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2023-11-10"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(52)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(52)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(71), // Peach
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Wuque Studio Morandi Linear",
			ShortDescription: "Linear switch with subtle elegance and smooth action.",
			LongDescription:  "Wuque Studio Morandi Linear switches are characterized by their subtle elegance and smooth linear action. Featuring a 62g actuation force, these switches offer a refined typing experience, perfect for those who appreciate understated aesthetics and superior performance.",
			ManufacturerID:   ptr(21), // Wuque Studio
			BrandID:          ptr(21), // Wuque Studio
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2024-01-05"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(62)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(62)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(58), // Gray
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Wuque Studio Quartz Linear",
			ShortDescription: "Linear switch with a clear, transparent aesthetic.",
			LongDescription:  "Wuque Studio Quartz Linear switches combine a clear, transparent aesthetic with smooth linear action. Featuring a 55g actuation force, these switches are perfect for users who want a visually striking keyboard with a seamless typing experience.",
			ManufacturerID:   ptr(21), // Wuque Studio
			BrandID:          ptr(21), // Wuque Studio
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2024-03-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(26), // Polycarbonate
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
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
			Name:             "Wuque Studio Standard Red",
			ShortDescription: "Linear switch with a light actuation force and smooth travel.",
			LongDescription:  "Wuque Studio Standard Red switches offer a smooth linear action with a light 45g actuation force. Ideal for users who prefer quick, effortless keystrokes, these switches deliver a responsive and fluid typing experience.",
			ManufacturerID:   ptr(21), // Wuque Studio
			BrandID:          ptr(21), // Wuque Studio
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2023-08-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(40)),
				BottomOutForce:        ptr(float32(45)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
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
			Name:             "Wuque Studio Standard Yellow",
			ShortDescription: "Clicky switch with crisp feedback and moderate resistance.",
			LongDescription:  "Wuque Studio Standard Yellow switches deliver a satisfying clicky experience with moderate 55g resistance. Designed for users who enjoy both tactile and auditory feedback, these switches offer a balanced typing experience with crisp actuation.",
			ManufacturerID:   ptr(21), // Wuque Studio
			BrandID:          ptr(21), // Wuque Studio
			SwitchTypeID:     3,       // Clicky
			ReleaseDate:      parseDate("2023-08-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(50), // Yellow
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Wuque Studio Standard Brown",
			ShortDescription: "Tactile switch with a pronounced bump and firm feedback.",
			LongDescription:  "Wuque Studio Standard Brown switches offer a tactile experience with a pronounced bump and firm 65g feedback. Ideal for users who prefer a strong tactile response, these switches are perfect for those who enjoy a more substantial typing feel.",
			ManufacturerID:   ptr(21), // Wuque Studio
			BrandID:          ptr(21), // Wuque Studio
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2023-08-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(48), // Brown
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Tbcats Phoenix",
			ShortDescription: "Linear switch with a smooth glide and moderate actuation force.",
			LongDescription:  "The Tbcats Phoenix switch is a linear switch that excels in both typing and gaming scenarios. With a smooth glide and moderate actuation force of 45g, this switch provides a responsive and effortless keystroke. Its balanced feel makes it a versatile option for those who require precision and speed in their typing.",
			ManufacturerID:   ptr(31), // Tbcats Studio
			BrandID:          ptr(31), // Tbcats Studio
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2022-09-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(40)),
				BottomOutForce:        ptr(float32(45)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Tbcats Nebula",
			ShortDescription: "Tactile switch with a subtle bump and quiet operation.",
			LongDescription:  "The Tbcats Nebula switch offers a refined tactile experience with a subtle bump at the actuation point. With an actuation force of 55g, it strikes a balance between feedback and quiet operation, making it perfect for those who seek a comfortable and silent typing experience. Ideal for office environments and late-night sessions.",
			ManufacturerID:   ptr(31), // Tbcats Studio
			BrandID:          ptr(31), // Tbcats Studio
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2023-02-15"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(55), // Purple
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Tbcats Thunder",
			ShortDescription: "Clicky switch with a pronounced tactile bump and click sound.",
			LongDescription:  "The Tbcats Thunder switch delivers a classic mechanical keyboard feel with its pronounced tactile bump and audible click. With a robust 60g actuation force, this clicky switch is ideal for those who enjoy a distinct and satisfying sound with each keystroke, enhancing the overall typing experience.",
			ManufacturerID:   ptr(31), // Tbcats Studio
			BrandID:          ptr(31), // Tbcats Studio
			SwitchTypeID:     3,       // Clicky
			ReleaseDate:      parseDate("2023-05-10"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Tbcats Eclipse",
			ShortDescription: "Silent linear switch with a soft and smooth keystroke.",
			LongDescription:  "The Tbcats Eclipse switch is designed for those who demand a quiet and smooth typing experience. As a silent linear switch, it offers soft and fluid keystrokes with an actuation force of 50g. Perfect for shared spaces and noise-sensitive environments, the Eclipse ensures a seamless typing experience without compromising on performance.",
			ManufacturerID:   ptr(31), // Tbcats Studio
			BrandID:          ptr(31), // Tbcats Studio
			SwitchTypeID:     4,       // Silent Linear
			ReleaseDate:      parseDate("2023-08-05"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(50)),
				TotalTravel:           ptr(float32(3.8)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Tbcats Cosmic",
			ShortDescription: "Silent tactile switch with a gentle tactile bump.",
			LongDescription:  "The Tbcats Cosmic switch offers a unique blend of silence and tactility. Featuring a gentle bump and a 55g actuation force, this silent tactile switch provides the tactile feedback users love without the accompanying noise. It's ideal for those who prefer a quiet yet responsive typing experience.",
			ManufacturerID:   ptr(31), // Tbcats Studio
			BrandID:          ptr(31), // Tbcats Studio
			SwitchTypeID:     5,       // Silent Tactile
			ReleaseDate:      parseDate("2024-01-20"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(3.8)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(55), // Purple
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Tbcats Cloud Water V2",
			ShortDescription: "Smooth linear switch with a soft and quiet keystroke.",
			LongDescription:  "The Tbcats Cloud Water V2 switch is a refined linear switch known for its soft and quiet keystrokes. With an actuation force of 45g, it offers a smooth and effortless typing experience, making it an excellent choice for users who prioritize silence and fluidity in their keystrokes.",
			ManufacturerID:   ptr(31), // Tbcats Studio
			BrandID:          ptr(31), // Tbcats Studio
			SwitchTypeID:     4,       // Silent Linear
			ReleaseDate:      parseDate("2023-04-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(40)),
				BottomOutForce:        ptr(float32(45)),
				TotalTravel:           ptr(float32(3.8)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Tbcats Eclair Green",
			ShortDescription: "Tactile switch with a medium bump and satisfying feedback.",
			LongDescription:  "The Tbcats Eclair Green switch provides a tactile experience with a satisfying medium bump and a 55g actuation force. It offers just the right amount of resistance for both typing and gaming, making it a versatile choice for users who value tactile feedback with a smooth return.",
			ManufacturerID:   ptr(31), // Tbcats Studio
			BrandID:          ptr(31), // Tbcats Studio
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2023-03-15"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
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
			Name:             "Tbcats Blue Balloon VW",
			ShortDescription: "Linear switch with smooth actuation and distinct aesthetics.",
			LongDescription:  "The Tbcats Blue Balloon VW switch is a linear switch known for its smooth keystroke and unique colorway. With a moderate actuation force of 58g, this switch offers a fluid and consistent typing experience, making it ideal for both gaming and typing. Its distinct blue color and reliable performance make it a standout choice for users who prefer linear switches.",
			ManufacturerID:   ptr(31), // Tbcats Studio
			BrandID:          ptr(31), // Tbcats Studio
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2024-06-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(58)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(48)),
				BottomOutForce:        ptr(float32(58)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Tbcats Eclair Mix",
			ShortDescription: "Hybrid switch offering a blend of tactile and linear characteristics.",
			LongDescription:  "The Tbcats Eclair Mix switch combines the best of both worlds with a hybrid design that features both tactile and linear characteristics. With an actuation force of 52g, it offers smooth keystrokes with subtle feedback, providing a unique typing experience that caters to both tactile lovers and linear enthusiasts.",
			ManufacturerID:   ptr(31), // Tbcats Studio
			BrandID:          ptr(31), // Tbcats Studio
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2023-07-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(52)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(52)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(51), // White
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Tbcats Eclair Orange",
			ShortDescription: "Tactile switch with a pronounced bump and strong feedback.",
			LongDescription:  "The Tbcats Eclair Orange switch is perfect for users who crave a pronounced tactile bump and strong feedback. With an actuation force of 62g, this switch provides a bold and responsive typing experience, making it ideal for those who prefer a more assertive tactile feel.",
			ManufacturerID:   ptr(31), // Tbcats Studio
			BrandID:          ptr(31), // Tbcats Studio
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2023-11-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(62)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(52)),
				BottomOutForce:        ptr(float32(62)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(56), // Orange
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "C³ Equalz x TKC Banana Split",
			ShortDescription: "Linear switch with a thocky sound profile and smooth action.",
			LongDescription:  "The C³ Equalz x TKC Banana Split switch is a standout in the linear category, offering a deep, thocky sound that resonates with every keystroke. Known for its smooth action and consistent feel, this switch is part of the Macho Linear series and is perfect for enthusiasts who seek both performance and auditory satisfaction in their typing experience. Its unique design and reliable build make it a popular choice for custom keyboard builds.",
			ManufacturerID:   ptr(22), // JWK
			BrandID:          ptr(32), // Collaboration: C³ Equalz x TKC
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2021-07-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(62)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(62)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(50), // Yellow
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(51), // White
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "C³ Equalz x TKC Tangerine 67 (Dark Green)",
			ShortDescription: "Smooth linear switch with a 67g actuation force.",
			LongDescription:  "The C³ Equalz x TKC Tangerine 67 switch features a dark green housing and delivers a buttery-smooth linear action with a 67g actuation force. Its pre-lubed design ensures an ultra-smooth feel, making it a top pick for users who prioritize a thocky sound and responsive keypress. Ideal for both typing and gaming, this switch is a part of the renowned Tangerine series, known for its consistent performance and satisfying acoustics.",
			ManufacturerID:   ptr(22), // JWK
			BrandID:          ptr(32), // Collaboration: C³ Equalz x TKC
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2020-05-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(67)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(56), // Orange
				TopHousingColorID:     ptr(54), // Green
				BottomHousingColorID:  ptr(54), // Green
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "C³ Equalz x TKC Tangerine 62 (Light Green)",
			ShortDescription: "Linear switch with a light 62g actuation force.",
			LongDescription:  "The C³ Equalz x TKC Tangerine 62 switch is a favorite among those who prefer a lighter actuation. With a 62g force and light green housing, this linear switch offers a quick, responsive keystroke with minimal resistance. Pre-lubed for extra smoothness, it is designed for users who appreciate a fluid typing experience paired with a satisfying sound profile. The Tangerine series is celebrated for its excellent build quality and performance.",
			ManufacturerID:   ptr(22), // JWK
			BrandID:          ptr(32), // Collaboration: C³ Equalz x TKC
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2020-05-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(62)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(62)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(56), // Orange
				TopHousingColorID:     ptr(54), // Green
				BottomHousingColorID:  ptr(54), // Green
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "C³ Equalz x TKC Kiwi",
			ShortDescription: "Tactile switch with a pronounced bump and 67g force.",
			LongDescription:  "The C³ Equalz x TKC Kiwi switch is designed for tactile switch enthusiasts who crave a strong, pronounced bump with each keystroke. With a 67g actuation force, it offers firm feedback that is both satisfying and responsive, making it an excellent choice for typists who demand precision. Its distinctive green housing adds a touch of style to any keyboard build, while its performance keeps it at the top of the list for tactile switch lovers.",
			ManufacturerID:   ptr(22), // JWK
			BrandID:          ptr(32), // Collaboration: C³ Equalz x TKC
			SwitchTypeID:     2,       // Tactile
			ReleaseDate:      parseDate("2021-02-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(67)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(54), // Green
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(54), // Green
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "C³ Equalz x TKC Red Smoothie",
			ShortDescription: "Linear fruit switch with a smooth action.",
			LongDescription:  "The C³ Equalz x TKC Red Smoothie switch stands out with its vibrant red color and exceptionally smooth linear action. As part of the fruit-themed switch series, it is designed to deliver a consistent, fluid typing experience that feels just as delightful as its namesake. Whether for work or play, this switch offers a balance of aesthetics and performance that appeals to a wide range of users.",
			ManufacturerID:   ptr(22), // JWK
			BrandID:          ptr(32), // Collaboration: C³ Equalz x TKC
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2023-01-15"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(53)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(45), // Red
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "C³ Equalz x TKC Blue Smoothie",
			ShortDescription: "Linear fruit switch with 67g actuation force.",
			LongDescription:  "The C³ Equalz x TKC Blue Smoothie switch offers a smooth linear action paired with a 67g actuation force. Part of the beloved fruit-themed switch series, it combines a visually appealing blue hue with a satisfying keystroke, making it a solid choice for users who value both aesthetics and performance. This switch is ideal for those who enjoy a bit more resistance in their typing experience.",
			ManufacturerID:   ptr(22), // JWK
			BrandID:          ptr(32), // Collaboration: C³ Equalz x TKC
			SwitchTypeID:     1,       // Linear
			ReleaseDate:      parseDate("2023-01-15"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(67)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(47), // Blue
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
	}

	err := processSwitches(tx, commonSwitches, admin)
	if err != nil {
		return err
	}

	return nil
}
