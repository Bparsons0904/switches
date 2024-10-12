package seed

import (
	"switches/models"

	"gorm.io/gorm"
)

func seedGateron(tx *gorm.DB, admin models.User) error {
	gateron := []models.Switch{
		{
			Name:             "Gateron G Pro Red",
			ShortDescription: "Linear switch with factory lubrication for smoother keystrokes.",
			LongDescription:  "Gateron G Pro Red switches are enhanced linear switches featuring factory lubrication for smoother keystrokes and improved performance. With an actuation force of 45 grams, these switches provide a light and responsive typing experience, making them ideal for fast typists and gamers. The G Pro Red switches are designed to offer minimal resistance and friction, ensuring a consistent and satisfying keystroke.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2020-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			SiteURL:          "https://www.gateron.co/products/gateron-switch-set",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-G-Pro-Red-Switch-Set_1000x.jpg?v=1657361719",
					AltText:     "Gateron G Pro Red switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				SpringTypeID:          ptr(6),  // Linear spring
				SpringMaterialTypeID:  ptr(17), // Steel (typical for most switches)
				TopHousingMaterialID:  ptr(26), // Polycarbonate (common for G Pro series)
				BaseHousingMaterialID: ptr(24), // Nylon (common for G Pro series)
				StemMaterialID:        ptr(23), // POM (typical for Gateron switches)
				HasShineThrough:       ptr(true),
				PreTravel:             ptr(float32(2.0)), // Typical for Gateron switches
				TotalTravel:           ptr(float32(4.0)), // Typical for Gateron switches
				ActuationPoint:        ptr(float32(2.0)), // Typical for Gateron switches
				BottomOutForce:        ptr(float32(60)),  // Estimated
				SoundLevelID:          ptr(29),           // Low
				SoundTypeID:           ptr(36),           // Creamy (typical for lubed linears)
				FactoryLubed:          true,
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(52), // Transparent (common for G Pro series)
				BottomHousingColorID:  ptr(52), // Transparent (common for G Pro series)
				PinConfigurationID:    ptr(75), // 5-Pin (common for modern Gateron switches)
			},
		},
		{
			Name:             "Gateron G Pro Yellow",
			ShortDescription: "Linear switch with a medium actuation force and improved smoothness.",
			LongDescription:  "Gateron G Pro Yellow switches offer a medium actuation force of 50 grams and are factory lubricated to provide a smooth and responsive typing experience. These linear switches are perfect for users who want a balanced feel that works well for both gaming and typing. The G Pro Yellow switches are known for their reliability and smoothness, making them a popular choice among enthusiasts.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2020-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			SiteURL:          "https://www.gateron.co/products/gateron-switch-set",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-G-Pro-Yellow-Switch-Set_1000x.jpg?v=1657361719",
					AltText:     "Gateron G Pro Yellow switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				SpringTypeID:          ptr(6),  // Linear spring
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(true),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(65)), // Estimated
				SoundLevelID:          ptr(30),          // Medium
				SoundTypeID:           ptr(36),          // Creamy
				FactoryLubed:          true,
				StemColorID:           ptr(50), // Yellow
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
			},
		},
		{
			Name:             "Gateron G Pro Black",
			ShortDescription: "Linear switch with a heavier actuation force for more resistance.",
			LongDescription:  "Gateron G Pro Black switches are linear switches designed for users who prefer a heavier actuation force of 60 grams. These switches provide a firm and controlled typing experience, ideal for those who enjoy more resistance in their keystrokes. Factory lubrication ensures smooth operation, making the G Pro Black switches a reliable option for both gaming and long typing sessions.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2020-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			SiteURL:          "https://www.gateron.co/products/gateron-switch-set",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-G-Pro-Black-Switch-Set_1000x.jpg?v=1657361719",
					AltText:     "Gateron G Pro Black switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				SpringTypeID:          ptr(6),  // Linear spring
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(true),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(75)), // Estimated
				SoundLevelID:          ptr(30),          // Medium
				SoundTypeID:           ptr(34),          // Thocky (due to heavier spring)
				FactoryLubed:          true,
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
			},
		},
		{
			Name:             "Gateron G Pro Brown",
			ShortDescription: "Tactile switch with factory lubrication and a gentle bump.",
			LongDescription:  "Gateron G Pro Brown switches feature a gentle tactile bump with factory lubrication, offering a smooth and quiet typing experience. With an actuation force of 55 grams, these tactile switches provide a satisfying feedback without the click sound, making them suitable for both gaming and typing. The G Pro Brown switches are designed to offer a balanced tactile feel, appealing to users who prefer moderate resistance in their typing.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2020-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			SiteURL:          "https://www.gateron.co/products/gateron-switch-set",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-G-Pro-Brown-Switch-Set_1000x.jpg?v=1657361719",
					AltText:     "Gateron G Pro Brown switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce: ptr(float32(55)),
				SpringTypeID: ptr(
					6,
				), // Linear spring (the tactility comes from the stem, not the spring)
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(true),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(70)), // Estimated
				SoundLevelID:          ptr(30),          // Medium
				SoundTypeID:           ptr(36),          // Creamy
				TactilityTypeID: ptr(
					38,
				), // Leaf Spring (typical for tactile switches)
				TactilityFeedbackID: ptr(
					42,
				), // Rounded (Gateron Browns are known for a subtle bump)
				FactoryLubed:         true,
				StemColorID:          ptr(48), // Brown
				TopHousingColorID:    ptr(52), // Transparent
				BottomHousingColorID: ptr(52), // Transparent
				PinConfigurationID:   ptr(75), // 5-Pin
			},
		},
		{
			Name:             "Gateron G Pro Green",
			ShortDescription: "Clicky switch with a strong tactile bump and audible click.",
			LongDescription:  "Gateron G Pro Green switches provide a strong tactile bump with an audible click, ideal for users who enjoy feedback and sound in their typing. With an actuation force of 80 grams, these clicky switches offer a pronounced and satisfying typing experience. Factory lubrication ensures smooth keystrokes, making the G Pro Green switches a reliable choice for enthusiasts who value both tactile feedback and auditory satisfaction.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2020-01-01"),
			Available:        true,
			PricePoint:       2,                                                    // Average
			SiteURL:          "https://www.gateron.co/products/gateron-switch-set", // Update if an official URL becomes available
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-G-Pro-Green-Switch-Set_1000x.jpg?v=1662948459",
					AltText:     "Gateron G Pro Green switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(80)),
				SpringTypeID:          ptr(6),  // Linear spring
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(26), // Polycarbonate (common for G Pro series)
				BaseHousingMaterialID: ptr(24), // Nylon (common for G Pro series)
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(true),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(95)), // Estimated
				SoundLevelID:          ptr(31),          // High
				SoundTypeID:           ptr(33),          // Clicky
				TactilityTypeID:       ptr(40),          // Click Bar (typical for clicky switches)
				TactilityFeedbackID:   ptr(41),          // Sharp (typical for clicky switches)
				FactoryLubed:          true,
				StemColorID:           ptr(54), // Green
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
			},
		},
		{
			Name:             "Gateron G Pro Blue",
			ShortDescription: "Clicky switch with an audible click and tactile feedback.",
			LongDescription:  "Gateron G Pro Blue switches are clicky switches designed for users who appreciate tactile feedback paired with an audible click. With an actuation force of 60 grams, these switches offer a satisfying typing experience with a distinct click sound. The G Pro Blue switches are ideal for users who prefer a more traditional mechanical keyboard feel with pronounced feedback and sound.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2020-01-01"),
			Available:        true,
			PricePoint:       2,                                                    // Average
			SiteURL:          "https://www.gateron.co/products/gateron-switch-set", // Update if an official URL becomes available
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-G-Pro-Blue-Switch-Set_1000x.jpg?v=1662948459",
					AltText:     "Gateron G Pro Blue switch",
					OwnerType:   "switches",
				},
			}, Details: models.Details{
				SpringForce:           ptr(float32(60)),
				SpringTypeID:          ptr(6),  // Linear spring
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(true),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(75)), // Estimated
				SoundLevelID:          ptr(31),          // High
				SoundTypeID:           ptr(33),          // Clicky
				TactilityTypeID:       ptr(40),          // Click Bar
				TactilityFeedbackID:   ptr(41),          // Sharp
				FactoryLubed:          true,
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
			},
		},
		{
			Name:             "Gateron Pro 2.0 Red",
			ShortDescription: "Linear switch with upgraded smoothness and factory lubrication.",
			LongDescription:  "Gateron Pro 2.0 Red switches are enhanced linear switches designed for a smooth typing experience. With an actuation force of 45 grams, these switches are factory lubricated to provide a more refined and responsive keystroke. The Pro 2.0 Red switches offer a light and effortless typing experience, making them ideal for fast-paced gaming and everyday typing tasks.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-01-01"),
			Available:        true,
			PricePoint:       2,                                                              // Average
			SiteURL:          "https://www.gateron.co/products/gateron-g-pro-2-0-switch-set", // Update if an official URL becomes available
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-G-Pro-2.0-Red-Switch-Set_1000x.jpg?v=1657505037",
					AltText:     "Gateron G Pro 2.0 Blue switch",
					OwnerType:   "switches",
				},
			}, Details: models.Details{
				SpringForce:           ptr(float32(45)),
				SpringTypeID:          ptr(6),  // Linear spring
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(true),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(60)), // Estimated
				SoundLevelID:          ptr(29),          // Low
				SoundTypeID:           ptr(36),          // Creamy
				FactoryLubed:          true,
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
			},
		},
		{
			Name:             "Gateron Pro 2.0 Yellow",
			ShortDescription: "Linear switch with a medium actuation force and improved smoothness.",
			LongDescription:  "Gateron Pro 2.0 Yellow switches feature an upgraded linear design with a medium actuation force of 50 grams. These switches are factory lubricated, offering a smooth and responsive typing experience. The Pro 2.0 Yellow switches are perfect for users who prefer a balanced feel that works well for both gaming and typing.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-01-01"),
			Available:        true,
			PricePoint:       2,                                                              // Average
			SiteURL:          "https://www.gateron.co/products/gateron-g-pro-2-0-switch-set", // Update if an official URL becomes available
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-G-Pro-2.0-Yellow-Switch-Set_1000x.jpg?v=1657505119",
					AltText:     "Gateron G Pro 2.0 Blue switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				SpringTypeID:          ptr(6),  // Linear spring
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(true),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(65)), // Estimated
				SoundLevelID:          ptr(30),          // Medium
				SoundTypeID:           ptr(36),          // Creamy
				FactoryLubed:          true,
				StemColorID:           ptr(50), // Yellow
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
			},
		},
		{
			Name:             "Gateron Pro 2.0 Black",
			ShortDescription: "Linear switch with a heavier actuation force and factory lubrication.",
			LongDescription:  "Gateron Pro 2.0 Black switches are designed for users who prefer a heavier actuation force of 60 grams. These switches provide a firm and controlled typing experience, with factory lubrication ensuring smooth operation. The Pro 2.0 Black switches are ideal for those who enjoy more resistance in their keystrokes, offering a reliable option for both gaming and long typing sessions.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-01-01"),
			Available:        true,
			PricePoint:       2,                                                              // Average
			SiteURL:          "https://www.gateron.co/products/gateron-g-pro-2-0-switch-set", // Update if an official URL becomes available
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-G-Pro-2.0-Black-Switch-Set_1000x.jpg?v=1657505119",
					AltText:     "Gateron G Pro 2.0 Blue switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				SpringTypeID:          ptr(6),  // Linear spring
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(true),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(75)), // Estimated
				SoundLevelID:          ptr(30),          // Medium
				SoundTypeID:           ptr(34),          // Thocky (due to heavier spring)
				FactoryLubed:          true,
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
			},
		},
		{
			Name:             "Gateron Pro 2.0 Brown",
			ShortDescription: "Tactile switch with a gentle bump and improved stability.",
			LongDescription:  "Gateron Pro 2.0 Brown switches offer a tactile typing experience with a gentle bump and an actuation force of 55 grams. These switches are factory lubricated, providing a smooth and quiet typing experience. The Pro 2.0 Brown switches are designed to offer balanced tactile feedback, making them suitable for both gaming and typing.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2021-01-01"),
			Available:        true,
			PricePoint:       2,                                                              // Average
			SiteURL:          "https://www.gateron.co/products/gateron-g-pro-2-0-switch-set", // Update if an official URL becomes available
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-G-Pro-2.0-Brown-Switch-Set_1000x.jpg?v=1657505119",
					AltText:     "Gateron G Pro 2.0 Blue switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				SpringTypeID:          ptr(6),  // Linear spring
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(true),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(70)), // Estimated
				SoundLevelID:          ptr(30),          // Medium
				SoundTypeID:           ptr(36),          // Creamy
				TactilityTypeID:       ptr(38),          // Leaf Spring
				TactilityFeedbackID:   ptr(42),          // Rounded
				FactoryLubed:          true,
				StemColorID:           ptr(48), // Brown
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
			},
		},
		{
			Name:             "Gateron Pro 2.0 Blue",
			ShortDescription: "Clicky switch with an audible click and refined feedback.",
			LongDescription:  "Gateron Pro 2.0 Blue switches are clicky switches designed for users who appreciate tactile feedback paired with an audible click. With an actuation force of 60 grams, these switches offer a satisfying typing experience with a distinct click sound. The Pro 2.0 Blue switches are ideal for users who prefer a more traditional mechanical keyboard feel with pronounced feedback and sound.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2021-01-01"),
			Available:        true,
			PricePoint:       2,                                                              // Average
			SiteURL:          "https://www.gateron.co/products/gateron-g-pro-2-0-switch-set", // Update if an official URL becomes available
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-G-Pro-2.0-Blue-Switch-Set_1000x.jpg?v=1657505119",
					AltText:     "Gateron G Pro 2.0 Blue switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				SpringTypeID:          ptr(6),  // Linear spring
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(true),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(75)), // Estimated
				SoundLevelID:          ptr(31),          // High
				SoundTypeID:           ptr(33),          // Clicky
				TactilityTypeID:       ptr(40),          // Click Bar
				TactilityFeedbackID:   ptr(41),          // Sharp
				FactoryLubed:          true,
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
			},
		},
		{
			Name:             "Gateron Pro 3.0 Red",
			ShortDescription: "Linear switch with enhanced smoothness and stability.",
			LongDescription:  "Gateron Pro 3.0 Red switches represent the latest in Gateron's linear switch technology, featuring enhanced smoothness and stability. With an actuation force of 45 grams, these switches are designed for users who demand a light and effortless typing experience. The Pro 3.0 Red switches come factory lubricated, ensuring consistent and smooth keystrokes with minimal friction, making them ideal for both gaming and typing.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2022-01-01"),
			Available:        true,
			PricePoint:       3,                                                              // Expensive
			SiteURL:          "https://www.gateron.co/products/gateron-g-pro-3-0-switch-set", // Update if an official URL becomes available
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.gateron.co/cdn/shop/files/Gateron-G-Pro-3.0-Red-Switch-Set_1000x.jpg?v=1683342544",
					AltText:     "Gateron G Pro 3.0 Red switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				SpringTypeID:          ptr(6),  // Linear spring
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(true),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(60)), // Estimated
				SoundLevelID:          ptr(29),          // Low
				SoundTypeID:           ptr(36),          // Creamy
				FactoryLubed:          true,
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
			},
		},
		{
			Name:             "Gateron Pro 3.0 Yellow",
			ShortDescription: "Linear switch with medium actuation force and advanced performance.",
			LongDescription:  "Gateron Pro 3.0 Yellow switches are linear switches designed for users seeking a medium actuation force of 50 grams. These switches offer advanced performance with factory lubrication for ultra-smooth keystrokes. The Pro 3.0 Yellow switches are ideal for users who need a balanced typing experience that is suitable for both gaming and long typing sessions.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2022-01-01"),
			Available:        true,
			PricePoint:       3,                                                              // Expensive
			SiteURL:          "https://www.gateron.co/products/gateron-g-pro-3-0-switch-set", // Update if an official URL becomes available
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.gateron.co/cdn/shop/files/Gateron-G-Pro-3.0-Yellow-Switch-Set_1000x.jpg?v=1702451537",
					AltText:     "Gateron G Pro 3.0 Yellow  switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				SpringTypeID:          ptr(6),  // Linear spring
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(true),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(65)), // Estimated
				SoundLevelID:          ptr(30),          // Medium
				SoundTypeID:           ptr(36),          // Creamy
				FactoryLubed:          true,
				StemColorID:           ptr(50), // Yellow
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
			},
		},
		{
			Name:             "Gateron Pro 3.0 Black",
			ShortDescription: "Linear switch with a heavier actuation force and enhanced durability.",
			LongDescription:  "Gateron Pro 3.0 Black switches are the pinnacle of Gateron's linear switch lineup, offering a heavier actuation force of 60 grams. These switches provide a firm and controlled typing experience, making them ideal for users who prefer more resistance in their keystrokes. Factory lubrication and enhanced durability make the Pro 3.0 Black switches a reliable choice for both gaming and extended typing sessions.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2022-01-01"),
			Available:        true,
			PricePoint:       3,                                                              // Expensive
			SiteURL:          "https://www.gateron.co/products/gateron-g-pro-3-0-switch-set", // Update if an official URL becomes available
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.gateron.co/cdn/shop/files/Gateron-G-Pro-3.0-Black-Switch-Set_1000x.jpg?v=1702451537",
					AltText:     "Gateron G Pro 3.0 Black switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				SpringTypeID:          ptr(6),  // Linear spring
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(true),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(75)), // Estimated
				SoundLevelID:          ptr(30),          // Medium
				SoundTypeID:           ptr(34),          // Thocky (due to heavier spring)
				FactoryLubed:          true,
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
			},
		},
		{
			Name:             "Gateron Pro 3.0 Brown",
			ShortDescription: "Tactile switch with a refined bump and improved smoothness.",
			LongDescription:  "Gateron Pro 3.0 Brown switches offer a refined tactile typing experience with a gentle bump and an actuation force of 55 grams. These switches are factory lubricated to provide a smooth and quiet operation, making them suitable for both gaming and typing. The Pro 3.0 Brown switches are designed to offer a balanced tactile feel with improved stability and consistency.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2022-01-01"),
			Available:        true,
			PricePoint:       3,                                                              // Expensive
			SiteURL:          "https://www.gateron.co/products/gateron-g-pro-3-0-switch-set", // Update if an official URL becomes available
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.gateron.co/cdn/shop/files/Gateron-G-Pro-3.0-Brown-Switch-Set_1000x.jpg?v=1702451537",
					AltText:     "Gateron G Pro 3.0 Brown switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				SpringTypeID:          ptr(6),  // Linear spring
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(true),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(70)), // Estimated
				SoundLevelID:          ptr(30),          // Medium
				SoundTypeID:           ptr(36),          // Creamy
				TactilityTypeID:       ptr(38),          // Leaf Spring
				TactilityFeedbackID:   ptr(42),          // Rounded
				FactoryLubed:          true,
				StemColorID:           ptr(48), // Brown
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
			},
		},
		{
			Name:             "Gateron Pro 3.0 Blue",
			ShortDescription: "Clicky switch with an audible click and premium feedback.",
			LongDescription:  "Gateron Pro 3.0 Blue switches are clicky switches designed for users who appreciate tactile feedback paired with an audible click. With an actuation force of 60 grams, these switches offer a satisfying typing experience with a distinct click sound. The Pro 3.0 Blue switches are ideal for users who prefer a more traditional mechanical keyboard feel with pronounced feedback and sound.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2022-01-01"),
			Available:        true,
			PricePoint:       3,                                                              // Expensive
			SiteURL:          "https://www.gateron.co/products/gateron-g-pro-3-0-switch-set", // Update if an official URL becomes available
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.gateron.co/cdn/shop/files/Gateron-G-Pro-3.0-Blue-Switch-Set_1000x.jpg?v=1702451537",
					AltText:     "Gateron G Pro 3.0 Blue switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				SpringTypeID:          ptr(6),  // Linear spring
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(true),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(75)), // Estimated
				SoundLevelID:          ptr(31),          // High
				SoundTypeID:           ptr(33),          // Clicky
				TactilityTypeID:       ptr(40),          // Click Bar
				TactilityFeedbackID:   ptr(41),          // Sharp
				FactoryLubed:          true,
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
			},
		},
		{
			Name:             "Gateron Ink V2 Red",
			ShortDescription: "Linear switch with improved smoothness and unique housing.",
			LongDescription:  "Gateron Ink V2 Red switches are designed for users who seek a smooth and responsive linear typing experience. With an actuation force of 45 grams, these switches offer light and effortless keystrokes, perfect for fast typists and gamers. The Ink V2 Red switches feature Gateron's distinctive Ink housing, which provides enhanced stability and durability, ensuring consistent performance over time.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			SiteURL:          "https://www.gateron.co/products/gateron-ink-switch",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-Ink-V2-Red-Switch-Set_1000x.jpg?v=1657360038",
					AltText:     "Gateron Ink V2 Red switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				SpringTypeID:          ptr(6),  // Linear spring
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(26), // Polycarbonate (Ink housing)
				BaseHousingMaterialID: ptr(26), // Polycarbonate (Ink housing)
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(false),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(60)), // Estimated
				SoundLevelID:          ptr(29),          // Low
				SoundTypeID:           ptr(34),          // Thocky
				FactoryLubed:          true,
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(53), // Smokey
				BottomHousingColorID:  ptr(53), // Smokey
				PinConfigurationID:    ptr(75), // 5-Pin
			},
		},
		{
			Name:             "Gateron Ink V2 Black",
			ShortDescription: "Linear switch with a heavier actuation force and improved design.",
			LongDescription:  "Gateron Ink V2 Black switches are premium linear switches known for their smooth actuation and distinct housing design. With a heavier actuation force of 60 grams, these switches provide a firm and controlled typing experience, making them ideal for users who prefer a stronger keystroke. The unique Ink housing not only enhances the switch's aesthetics but also contributes to its exceptional durability and stability.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			SiteURL:          "https://www.gateron.co/products/gateron-ink-switch",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-Ink-V2-Black-Switch-Set_1000x.jpg?v=1657360038",
					AltText:     "Gateron Ink V2 Black switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				SpringTypeID:          ptr(6),  // Linear spring
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(26), // Polycarbonate (Ink housing)
				BaseHousingMaterialID: ptr(26), // Polycarbonate (Ink housing)
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(false),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(75)), // Estimated
				SoundLevelID:          ptr(30),          // Medium
				SoundTypeID:           ptr(34),          // Thocky
				FactoryLubed:          true,
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(53), // Smokey
				BottomHousingColorID:  ptr(53), // Smokey
				PinConfigurationID:    ptr(75), // 5-Pin
			},
		},
		{
			Name:             "Gateron Ink V2 Yellow",
			ShortDescription: "Linear switch with mid-range actuation and unique housing.",
			LongDescription:  "Gateron Ink V2 Yellow switches provide a smooth linear feel with a mid-range actuation force of 50 grams. These switches are part of Gateron's premium Ink series, known for their exceptional build quality and unique housing design. The Ink V2 Yellow switches are ideal for users who want a responsive and consistent typing experience with just the right amount of resistance for both gaming and typing.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			SiteURL:          "https://www.gateron.co/products/gateron-ink-switch",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-Ink-V2-Yellow-Switch-Set_1000x.jpg?v=1657360038",
					AltText:     "Gateron Ink V2 Yellow switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				SpringTypeID:          ptr(6),  // Linear spring
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(26), // Polycarbonate (Ink housing)
				BaseHousingMaterialID: ptr(26), // Polycarbonate (Ink housing)
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(false),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(65)), // Estimated
				SoundLevelID:          ptr(30),          // Medium
				SoundTypeID:           ptr(34),          // Thocky
				FactoryLubed:          true,
				StemColorID:           ptr(50), // Yellow
				TopHousingColorID:     ptr(53), // Smokey
				BottomHousingColorID:  ptr(53), // Smokey
				PinConfigurationID:    ptr(75), // 5-Pin
			},
		},
		{
			Name:             "Gateron Ink V2 Blue",
			ShortDescription: "Clicky switch with unique housing for enhanced sound and feel.",
			LongDescription:  "Gateron Ink V2 Blue switches are clicky switches designed for users who appreciate tactile feedback paired with an audible click. With an actuation force of 55 grams, these switches offer a satisfying typing experience with a distinct click sound. The unique Ink housing adds to the switch's durability and enhances the overall feel, making the Ink V2 Blue a popular choice among enthusiasts who enjoy clicky switches with a premium build.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2021-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			SiteURL:          "https://www.gateron.co/products/gateron-ink-switch",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-Ink-V2-Blue-Switch-Set_1000x.jpg?v=1657360038",
					AltText:     "Gateron Ink V2 Blue switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				SpringTypeID:          ptr(6),  // Linear spring
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(26), // Polycarbonate (Ink housing)
				BaseHousingMaterialID: ptr(26), // Polycarbonate (Ink housing)
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(false),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(70)), // Estimated
				SoundLevelID:          ptr(31),          // High
				SoundTypeID:           ptr(33),          // Clicky
				TactilityTypeID:       ptr(40),          // Click Bar
				TactilityFeedbackID:   ptr(41),          // Sharp
				FactoryLubed:          true,
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(53), // Smokey
				BottomHousingColorID:  ptr(53), // Smokey
				PinConfigurationID:    ptr(75), // 5-Pin
			},
		},

		{
			Name:             "Gateron Oil King",
			ShortDescription: "Linear switch with smooth actuation and unique black housing.",
			LongDescription:  "Gateron Oil King switches are linear switches celebrated for their ultra-smooth keystrokes and unique black housing. With an actuation force of 55 grams, these switches are designed to offer a premium typing experience with minimal friction, making them ideal for both gaming and long typing sessions. The Oil King switches are part of Gateron's top-tier lineup, providing a balance of smoothness and durability that appeals to enthusiasts and professionals alike.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			SiteURL:          "https://www.gateron.co/products/gateron-oil-king-switch-set",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-Oil-King-Linear-Switch-Set_1000x.jpg?v=1657359635",
					AltText:     "Gateron Oil King switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				SpringTypeID:          ptr(6),  // Linear spring
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(24), // Nylon (unique black housing)
				BaseHousingMaterialID: ptr(24), // Nylon (unique black housing)
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(false),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(70)), // Estimated
				SoundLevelID:          ptr(30),          // Medium
				SoundTypeID:           ptr(34),          // Thocky
				FactoryLubed:          true,
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
			},
		},

		{
			Name:             "Gateron Baby Kangaroo",
			ShortDescription: "Linear switch with a unique spring design.",
			LongDescription:  "Gateron Baby Kangaroo switches offer a distinctive actuation feel, featuring a unique spring design for a novel linear experience.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2020-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			SiteURL:          "https://www.gateron.co/products/gateron-baby-kangaroo-tactile-switch-set",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.gateron.co/cdn/shop/files/Gateron-Baby-Kangaroo-2.0-Tactile-Switch-Set_1000x.jpg?v=1698722119",
					AltText:     "Gateron Baby Kangaroo switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce: ptr(
					float32(45),
				), // Estimated, as specific force wasn't provided
				SpringTypeID:          ptr(14), // Custom (unique spring design)
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(26), // Polycarbonate (typical for Gateron)
				BaseHousingMaterialID: ptr(24), // Nylon (typical for Gateron)
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(true),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(60)), // Estimated
				SoundLevelID:          ptr(30),          // Medium
				SoundTypeID:           ptr(36),          // Creamy
				FactoryLubed:          true,
				StemColorID: ptr(
					50,
				), // Yellow (assuming based on typical Gateron color scheme)
				TopHousingColorID:    ptr(52), // Transparent
				BottomHousingColorID: ptr(52), // Transparent
				PinConfigurationID:   ptr(75), // 5-Pin
			},
		},
		{
			Name:             "Gateron Cap V2 Brown",
			ShortDescription: "Upgraded tactile switch with a moderate actuation force.",
			LongDescription:  "Gateron Cap V2 Brown switches feature an upgraded tactile experience with enhanced sound dampening, offering a more refined tactile feel. With a moderate actuation force of 55 grams, these switches provide a satisfying bump with reduced noise, making them ideal for users who appreciate tactile feedback without the audible distraction. The Cap V2 series is designed for users seeking both performance and a quieter typing experience.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2021-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			SiteURL:          "https://www.gateron.co/products/gateron-cap-switch-set",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-Cap-V2-Switch-Set-Golden-Brown_1000x.jpg?v=1659520097",
					AltText:     "Gateron Cap V2 Brown switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				SpringTypeID:          ptr(6),  // Linear spring
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(true),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(70)), // Estimated
				SoundLevelID:          ptr(29),          // Low (due to enhanced sound dampening)
				SoundTypeID:           ptr(36),          // Creamy
				TactilityTypeID:       ptr(38),          // Leaf Spring
				TactilityFeedbackID:   ptr(42),          // Rounded
				FactoryLubed:          true,
				StemColorID:           ptr(48), // Brown
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(60), // Gold (based on the "Golden Brown" description)
				PinConfigurationID:    ptr(75), // 5-Pin
			},
		},
		{
			Name:             "Gateron Cap V2 Blue",
			ShortDescription: "Upgraded clicky switch with a tactile bump.",
			LongDescription:  "Gateron Cap V2 Blue switches feature an upgraded clicky experience with enhanced sound and tactile feedback, providing a satisfying typing feel.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2021-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			SiteURL:          "https://www.gateron.co/products/gateron-cap-switch-set",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-Cap-V2-Switch-Set-Blue_1000x.jpg?v=1659520097",
					AltText:     "Gateron Cap V2 Blue switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce: ptr(
					float32(60),
				), // Estimated based on typical Blue switch force
				SpringTypeID:          ptr(6),  // Linear spring
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(true),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(75)), // Estimated
				SoundLevelID:          ptr(31),          // High
				SoundTypeID:           ptr(33),          // Clicky
				TactilityTypeID:       ptr(40),          // Click Bar
				TactilityFeedbackID:   ptr(41),          // Sharp
				FactoryLubed:          true,
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
			},
		},
		{
			Name:             "Gateron Cap V2 Yellow",
			ShortDescription: "Upgraded linear switch with a smooth feel and factory lubrication.",
			LongDescription:  "Gateron Cap V2 Yellow switches provide an enhanced linear typing experience with factory lubrication for smoother keystrokes.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			SiteURL:          "https://www.gateron.co/products/gateron-cap-switch-set",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-Cap-V2-Switch-Set-Golden-Yellow_1000x.jpg?v=1657360751",
					AltText:     "Gateron Cap V2 Yellow switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce: ptr(
					float32(50),
				), // Estimated based on typical Yellow switch force
				SpringTypeID:          ptr(6),  // Linear spring
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(true),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(65)), // Estimated
				SoundLevelID:          ptr(29),          // Low
				SoundTypeID:           ptr(36),          // Creamy
				FactoryLubed:          true,
				StemColorID:           ptr(50), // Yellow
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(60), // Gold (based on "Golden Yellow" description)
				PinConfigurationID:    ptr(75), // 5-Pin
			},
		},
		{
			Name:             "Gateron Red",
			ShortDescription: "Linear switch with a light actuation force.",
			LongDescription:  "Gateron Red switches are favored for their smooth linear action and light actuation force of 45 grams, making them ideal for gaming and typing. These switches provide a consistent and uninterrupted keystroke, offering a responsive and fast typing experience. They are a popular choice among users who prefer a softer and quieter switch without tactile feedback or audible click, providing an overall smooth and fluid typing experience.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2014-01-01"),
			Available:        true,
			PricePoint:       1, // Value
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				SpringTypeID:          ptr(6),  // Linear spring
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(true),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(60)), // Estimated
				SoundLevelID:          ptr(29),          // Low
				SoundTypeID:           ptr(36),          // Creamy
				FactoryLubed:          false,
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin (typical for standard Gateron switches)
			},
		},
		{
			Name:             "Gateron Blue",
			ShortDescription: "Clicky switch with a tactile bump.",
			LongDescription:  "Gateron Blue switches are known for their clicky and tactile typing experience, offering an audible click sound and a distinct tactile bump with each key press. With an actuation force of 55 grams, these switches are popular among typists who appreciate auditory feedback and a satisfying tactile response. The crisp click sound makes them less ideal for quiet environments but highly favored for their satisfying and precise keystrokes.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2014-01-01"),
			Available:        true,
			PricePoint:       1, // Value
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				SpringTypeID:          ptr(6),  // Linear spring
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(true),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(70)), // Estimated
				SoundLevelID:          ptr(31),          // High
				SoundTypeID:           ptr(33),          // Clicky
				TactilityTypeID:       ptr(40),          // Click Bar
				TactilityFeedbackID:   ptr(41),          // Sharp
				FactoryLubed:          false,
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
			},
		},
		{
			Name:             "Gateron Brown",
			ShortDescription: "Tactile switch with a moderate actuation force.",
			LongDescription:  "Gateron Brown switches provide a balanced typing experience with a tactile bump and moderate actuation force of 55 grams. These switches are ideal for users who want tactile feedback without the loud click associated with Blue switches. Gateron Browns are versatile and well-suited for both typing and gaming, offering a satisfying yet quiet keystroke, making them a favorite among users who seek a middle ground between linear and clicky switches.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2014-01-01"),
			Available:        true,
			PricePoint:       1, // Value
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				SpringTypeID:          ptr(6),  // Linear spring
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(true),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(70)), // Estimated
				SoundLevelID:          ptr(30),          // Medium
				SoundTypeID:           ptr(36),          // Creamy
				TactilityTypeID:       ptr(38),          // Leaf Spring
				TactilityFeedbackID:   ptr(42),          // Rounded
				FactoryLubed:          false,
				StemColorID:           ptr(48), // Brown
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
			},
		},
		{
			Name:             "Gateron Black",
			ShortDescription: "Linear switch with a heavier actuation force.",
			LongDescription:  "Gateron Black switches offer a smooth linear action with a heavier actuation force of 60 grams, providing a firmer typing experience. These switches are ideal for users who prefer more resistance in their keystrokes, reducing the likelihood of accidental key presses. The Gateron Black switch is often chosen by gamers and heavy typists who want a consistent and controlled key press, making it a popular choice for those who appreciate a stiffer feel without tactile feedback.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2014-01-01"),
			Available:        true,
			PricePoint:       1, // Value
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				SpringTypeID:          ptr(6),  // Linear spring
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(true),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(75)), // Estimated
				SoundLevelID:          ptr(30),          // Medium
				SoundTypeID:           ptr(34),          // Thocky
				FactoryLubed:          false,
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
			},
		},
		{
			Name:             "Gateron Yellow",
			ShortDescription: "Linear switch with a medium actuation force.",
			LongDescription:  "Gateron Yellow switches are known for their smooth linear feel and medium actuation force of 50 grams, offering a balanced experience suitable for both gaming and typing. These switches provide a consistent keystroke with just the right amount of resistance, making them a favorite among users who want a switch that is not too light or too heavy. The Gateron Yellow switch is praised for its responsiveness and smoothness, making it a versatile choice for various applications.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2014-01-01"),
			Available:        true,
			PricePoint:       1, // Value
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				SpringTypeID:          ptr(6),  // Linear spring
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(true),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(65)), // Estimated
				SoundLevelID:          ptr(30),          // Medium
				SoundTypeID:           ptr(36),          // Creamy
				FactoryLubed:          false,
				StemColorID:           ptr(50), // Yellow
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
			},
		},
		{
			Name:             "Gateron Green",
			ShortDescription: "Clicky switch with a heavy actuation force.",
			LongDescription:  "Gateron Green switches are a heavier variant of the Gateron Blue, offering a more pronounced tactile bump and a louder click with each key press. With an actuation force of 80 grams, these switches provide a firm and satisfying clicky typing experience, ideal for users who enjoy strong tactile feedback and audible keystrokes. Gateron Greens are often chosen by enthusiasts who want a more deliberate and forceful typing experience.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2014-01-01"),
			Available:        true,
			PricePoint:       1, // Value
			Details: models.Details{
				SpringForce:           ptr(float32(80)),
				SpringTypeID:          ptr(6),  // Linear spring
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(true),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(95)), // Estimated
				SoundLevelID:          ptr(32),          // Very High
				SoundTypeID:           ptr(33),          // Clicky
				TactilityTypeID:       ptr(40),          // Click Bar
				TactilityFeedbackID:   ptr(41),          // Sharp
				FactoryLubed:          false,
				StemColorID:           ptr(54), // Green
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
			},
		},
		{
			Name:             "Gateron Clear",
			ShortDescription: "Linear switch with an extremely light actuation force.",
			LongDescription:  "Gateron Clear switches offer an ultra-light linear feel with an actuation force of only 35 grams, making them ideal for users who prefer minimal resistance during keystrokes. These switches are extremely smooth and require very little force to actuate, providing a quick and effortless typing experience. They are particularly well-suited for users who suffer from finger fatigue or those who prefer a feather-light typing touch.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2014-01-01"),
			Available:        true,
			PricePoint:       1, // Value
			Details: models.Details{
				SpringForce:           ptr(float32(35)),
				SpringTypeID:          ptr(6),  // Linear spring
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(true),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(50)), // Estimated
				SoundLevelID:          ptr(28),          // Very Low
				SoundTypeID:           ptr(36),          // Creamy
				FactoryLubed:          false,
				StemColorID:           ptr(49), // Clear
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
			},
		},
		{
			Name:             "Gateron Silent Red",
			ShortDescription: "Silent linear switch with a light actuation force.",
			LongDescription:  "Gateron Silent Red switches are designed to offer the smooth, linear feel of Gateron Reds while significantly reducing noise during actuation. With an actuation force of 45 grams, these switches incorporate dampeners that minimize the sound of both key press and release, making them ideal for quiet environments like offices or shared spaces. Gateron Silent Reds provide a responsive and fluid typing experience without the auditory distraction.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     4,      // Silent Linear
			ReleaseDate:      parseDate("2018-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				SpringTypeID:          ptr(6),  // Linear spring
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(true),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(3.8)), // Slightly shorter due to dampeners
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(60)), // Estimated
				SoundLevelID:          ptr(28),          // Very Low
				SoundTypeID:           ptr(35),          // Quiet
				FactoryLubed:          true,
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
			},
		},
		{
			Name:             "Gateron Silent Black",
			ShortDescription: "Silent linear switch with a heavier actuation force.",
			LongDescription:  "Gateron Silent Black switches offer a quiet and smooth linear typing experience with a heavier actuation force of 60 grams. These switches are ideal for users who prefer a more deliberate keystroke while maintaining a low-noise profile, thanks to the built-in dampeners that reduce the sound of both key press and release. Gateron Silent Blacks are perfect for those who require both performance and silence in their typing environments.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     4,      // Silent Linear
			ReleaseDate:      parseDate("2018-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				SpringTypeID:          ptr(6),  // Linear spring
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(true),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(3.8)), // Slightly shorter due to dampeners
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(75)), // Estimated
				SoundLevelID:          ptr(28),          // Very Low
				SoundTypeID:           ptr(35),          // Quiet
				FactoryLubed:          true,
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
			},
		},
		{
			Name:             "Gateron Pro Yellow",
			ShortDescription: "Enhanced linear switch with a medium actuation force.",
			LongDescription:  "Gateron Pro Yellow switches are an upgraded version of the standard Gateron Yellows, featuring factory lubrication for an even smoother keystroke. With a medium actuation force of 50 grams, these linear switches offer a balance between lightness and resistance, making them highly versatile for both gaming and typing. The factory lubrication reduces friction, providing a more refined and satisfying typing experience.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2020-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				SpringTypeID:          ptr(6),  // Linear spring
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(26), // Polycarbonate (typical for Pro series)
				BaseHousingMaterialID: ptr(24), // Nylon
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(true),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(65)), // Estimated
				SoundLevelID:          ptr(29),          // Low
				SoundTypeID:           ptr(36),          // Creamy
				FactoryLubed:          true,
				StemColorID:           ptr(50), // Yellow
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin (common for Pro series)
			},
		},
		{
			Name:             "Gateron Aliaz Silent",
			ShortDescription: "Silent tactile switch with a gentle tactile bump.",
			LongDescription:  "Gateron Aliaz Silent switches provided a quiet tactile experience with a smooth bump, popular among those who preferred stealthy typing.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(5), // ZealPC
			SwitchTypeID:     5,      // Silent Tactile
			ReleaseDate:      parseDate("2017-01-01"),
			Available:        false,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce: ptr(
					float32(60),
				), // Common force, but Aliaz came in multiple weights
				SpringTypeID:          ptr(6),  // Linear spring
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(true),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(3.8)), // Slightly shorter due to silent design
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(75)), // Estimated
				SoundLevelID:          ptr(28),          // Very Low
				SoundTypeID:           ptr(35),          // Quiet
				TactilityTypeID:       ptr(38),          // Leaf Spring
				TactilityFeedbackID:   ptr(42),          // Rounded
				FactoryLubed:          true,
				StemColorID:           ptr(52), // Transparent
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
			},
		},
		{
			Name:             "Gateron Orange",
			ShortDescription: "Tactile switch with a moderate actuation force and unique orange color.",
			LongDescription:  "Gateron Orange switches offer a noticeable tactile bump with a moderate actuation force, ideal for those seeking a distinct tactile feel with a pop of color.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2018-01-01"),
			Available:        true,
			PricePoint:       1, // Value
			Details: models.Details{
				SpringForce: ptr(
					float32(55),
				), // Estimated based on typical tactile switches
				SpringTypeID:          ptr(6),  // Linear spring
				SpringMaterialTypeID:  ptr(17), // Steel
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(24), // Nylon
				StemMaterialID:        ptr(23), // POM
				HasShineThrough:       ptr(true),
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(70)), // Estimated
				SoundLevelID:          ptr(30),          // Medium
				SoundTypeID:           ptr(36),          // Creamy
				TactilityTypeID:       ptr(38),          // Leaf Spring
				TactilityFeedbackID:   ptr(42),          // Rounded
				FactoryLubed:          false,
				StemColorID:           ptr(56), // Orange
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
			},
		},
		{
			Name:             "Gateron Brown Pro",
			ShortDescription: "Enhanced tactile switch with factory lubrication.",
			LongDescription:  "Gateron Brown Pro switches provide a smoother tactile experience compared to standard Gateron Browns, featuring factory lubrication for improved performance.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2020-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringTypeID:          ptr(6),             // Linear spring
				SpringForce:           ptr(float32(55.0)), // Typical for Brown switches
				TopHousingMaterialID:  ptr(26),            // Polycarbonate
				BaseHousingMaterialID: ptr(23),            // POM
				StemMaterialID:        ptr(23),            // POM
				StemColorID:           ptr(48),            // Brown
				TopHousingColorID: ptr(
					49,
				), // Clear (assuming it's similar to standard Browns)
				BottomHousingColorID: ptr(
					46,
				), // Black (assuming it's similar to standard Browns)
				PinConfigurationID:  ptr(75), // 5-pin
				FactoryLubed:        true,
				PreTravel:           ptr(float32(2.0)),
				TotalTravel:         ptr(float32(4.0)),
				ActuationForce:      ptr(float32(55.0)),
				BottomOutForce:      ptr(float32(60.0)),
				SoundLevelID:        ptr(30), // Medium
				SoundTypeID:         ptr(36), // Creamy (due to factory lubrication)
				TactilityTypeID:     ptr(38), // Leaf spring
				TactilityFeedbackID: ptr(42), // Rounded
			},
		},
		{
			Name:             "Gateron Purple",
			ShortDescription: "Tactile switch with a light tactile bump and unique purple housing.",
			LongDescription:  "Gateron Purple switches offer a light tactile bump with a distinct purple housing, providing a unique typing feel for those seeking a balance between tactile feedback and aesthetics.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2019-01-01"),
			Available:        true,
			PricePoint:       1, // Value
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(50.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(55), // Purple
				TopHousingColorID:     ptr(55), // Purple
				BottomHousingColorID:  ptr(55), // Purple
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          false,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(50.0)),
				BottomOutForce:        ptr(float32(60.0)),
				SoundLevelID:          ptr(30), // Medium
				SoundTypeID:           ptr(37), // Clacky
				TactilityTypeID:       ptr(38), // Leaf spring
				TactilityFeedbackID:   ptr(42), // Rounded
			},
		},
		{
			Name:             "Gateron Zealios V2",
			ShortDescription: "Premium tactile switch with a refined tactile bump.",
			LongDescription:  "A collaboration with ZealPC, Gateron Zealios V2 switches offer a highly refined tactile bump with premium components for a superior typing experience.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(5), // ZealPC
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2018-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringTypeID: ptr(6),
				SpringForce: ptr(
					float32(65.0),
				), // Comes in various weights, 65g is common
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(55), // Purple
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          true,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(65.0)),
				BottomOutForce:        ptr(float32(75.0)),
				SoundLevelID:          ptr(30), // Medium
				SoundTypeID:           ptr(34), // Thocky
				TactilityTypeID:       ptr(38), // Leaf spring
				TactilityFeedbackID:   ptr(41), // Sharp
			},
		},
		{
			Name:             "Gateron Kangaroo",
			ShortDescription: "Tactile switch with a strong tactile bump.",
			LongDescription:  "Gateron Kangaroo switches are known for their unique actuation profile and distinctive kangaroo-themed color, providing a strong tactile bump and responsive feedback.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2020-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(60.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(72), // Tan (approximating kangaroo color)
				TopHousingColorID:     ptr(72), // Tan
				BottomHousingColorID:  ptr(72), // Tan
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          false,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(3.5)),
				ActuationForce:        ptr(float32(60.0)),
				BottomOutForce:        ptr(float32(70.0)),
				SoundLevelID:          ptr(30), // Medium
				SoundTypeID:           ptr(36), // Creamy
				TactilityTypeID:       ptr(38), // Leaf spring
				TactilityFeedbackID:   ptr(43), // Crisp
			},
		},
		{
			Name:             "Gateron Yellow Tactile",
			ShortDescription: "Discontinued tactile switch with a unique yellow color.",
			LongDescription:  "Gateron Yellow Tactile switches offered a distinct tactile bump with a unique yellow color, but are no longer in production.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2015-01-01"),
			Available:        false,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(55.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(50), // Yellow
				TopHousingColorID:     ptr(50), // Yellow
				BottomHousingColorID:  ptr(50), // Yellow
				PinConfigurationID:    ptr(74), // 3-pin (assuming standard Gateron configuration)
				FactoryLubed:          false,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(55.0)),
				BottomOutForce:        ptr(float32(65.0)),
				SoundLevelID:          ptr(30), // Medium
				SoundTypeID:           ptr(37), // Clacky
				TactilityTypeID:       ptr(38), // Leaf spring
				TactilityFeedbackID:   ptr(42), // Rounded
			},
		},
		{
			Name:             "Gateron Sky",
			ShortDescription: "Discontinued tactile switch developed with Akko.",
			LongDescription:  "A tactile switch developed in collaboration with Akko, known for its sky-blue color and refined tactile bump, but now discontinued.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2017-01-01"),
			Available:        false,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(62.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(47), // Blue (closest to sky blue)
				TopHousingColorID:     ptr(47), // Blue
				BottomHousingColorID:  ptr(47), // Blue
				PinConfigurationID:    ptr(75), // 5-pin (assuming premium configuration)
				FactoryLubed:          false,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(62.0)),
				BottomOutForce:        ptr(float32(70.0)),
				SoundLevelID:          ptr(30), // Medium
				SoundTypeID:           ptr(36), // Creamy
				TactilityTypeID:       ptr(38), // Leaf spring
				TactilityFeedbackID:   ptr(43), // Crisp
			},
		},
		{
			Name:             "Gateron Silver",
			ShortDescription: "Linear switch with a shorter travel distance for faster actuation.",
			LongDescription:  "Gateron Silver switches offer a linear feel with a reduced travel distance, making them ideal for rapid keystrokes and gaming.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2019-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(45.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(59), // Silver
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          false,
				PreTravel:             ptr(float32(1.2)),
				TotalTravel:           ptr(float32(3.4)),
				ActuationForce:        ptr(float32(45.0)),
				BottomOutForce:        ptr(float32(55.0)),
				SoundLevelID:          ptr(29), // Low
				SoundTypeID:           ptr(37), // Clacky
				TactilityTypeID:       ptr(39), // Coil Spring (for linear switches)
				TactilityFeedbackID:   ptr(44), // Smooth
			},
		},
		{
			Name:             "Gateron Pro Red",
			ShortDescription: "Enhanced linear switch with factory lubrication.",
			LongDescription:  "Gateron Pro Red switches offer a smoother keystroke compared to standard Gateron Red, thanks to factory lubrication and improved components.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2020-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(45.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          true,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(45.0)),
				BottomOutForce:        ptr(float32(55.0)),
				SoundLevelID:          ptr(29), // Low
				SoundTypeID:           ptr(36), // Creamy
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(44), // Smooth
			},
		},
		{
			Name:             "Gateron Raccoon",
			ShortDescription: "Linear switch with a smooth actuation and unique design.",
			LongDescription:  "Gateron Raccoon switches are designed for a smooth linear typing experience, featuring a unique spring design and distinctive raccoon-themed aesthetics.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2020-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringTypeID:          ptr(7), // Progressive (assuming unique spring design)
				SpringForce:           ptr(float32(55.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(58), // Gray (approximating raccoon color)
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          true,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(55.0)),
				BottomOutForce:        ptr(float32(65.0)),
				SoundLevelID:          ptr(30), // Medium
				SoundTypeID:           ptr(36), // Creamy
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(44), // Smooth
			},
		},
		{
			Name:             "Gateron Milky Yellow",
			ShortDescription: "Budget-friendly linear switch with a smooth feel.",
			LongDescription:  "Gateron Cap Milky Yellow switches are known for their smooth linear feel and cost-effectiveness, featuring a milky housing that provides a distinct look and softer bottom-out feel. With an actuation force of 50 grams, these switches offer a balanced typing experience, making them ideal for both typing and gaming. The milky housing also contributes to a more muted sound profile, appealing to users who prefer a quieter switch.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2018-01-01"),
			Available:        true,
			PricePoint:       1, // Value
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(50.0)),
				TopHousingMaterialID:  ptr(24), // Nylon (milky housing)
				BaseHousingMaterialID: ptr(24), // Nylon (milky housing)
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(50), // Yellow
				TopHousingColorID:     ptr(51), // White (milky)
				BottomHousingColorID:  ptr(51), // White (milky)
				PinConfigurationID:    ptr(74), // 3-pin
				FactoryLubed:          false,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(50.0)),
				BottomOutForce:        ptr(float32(60.0)),
				SoundLevelID:          ptr(29), // Low
				SoundTypeID:           ptr(34), // Thocky
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(44), // Smooth
			},
		},
		{
			Name:             "Gateron Crystal Pink",
			ShortDescription: "Linear switch with a crystal-clear housing and pink stem.",
			LongDescription:  "Gateron Crystal Pink switches provide a smooth linear typing experience with a unique crystal-clear housing and pink stem for an aesthetic touch.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2020-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(45.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(57), // Pink
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          false,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(45.0)),
				BottomOutForce:        ptr(float32(55.0)),
				SoundLevelID:          ptr(31), // High (due to transparent housing)
				SoundTypeID:           ptr(37), // Clacky
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(44), // Smooth
			},
		},
		{
			Name:             "Gateron Ink Yellow",
			ShortDescription: "Linear switch with smooth actuation and mid-range force.",
			LongDescription:  "Gateron Box Ink Yellow switches combine the smooth linear feel of the Ink series with the added stability and durability of a box-style design. With a mid-range actuation force of 50 grams, these switches are perfect for users who want a responsive and consistent typing experience with added protection against dust and moisture. The Box Ink Yellow switches are ideal for both gaming and typing, offering the best of both worlds in performance and durability.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2020-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(50.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(50), // Yellow
				TopHousingColorID:     ptr(46), // Black (Ink housing)
				BottomHousingColorID:  ptr(46), // Black (Ink housing)
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          true,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(50.0)),
				BottomOutForce:        ptr(float32(60.0)),
				SoundLevelID:          ptr(29), // Low
				SoundTypeID:           ptr(34), // Thocky
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(44), // Smooth
			},
		},
		{
			Name:             "Gateron Silent Yellow",
			ShortDescription: "Silent linear switch with a smooth feel and reduced noise.",
			LongDescription:  "Gateron Silent Yellow switches provide a smooth linear typing experience with reduced noise, making them ideal for quiet environments and office settings.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     4,      // Silent Linear
			ReleaseDate:      parseDate("2018-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(50.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(50), // Yellow
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          true,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(3.8)), // Slightly shorter due to dampeners
				ActuationForce:        ptr(float32(50.0)),
				BottomOutForce:        ptr(float32(60.0)),
				SoundLevelID:          ptr(28), // Very Low
				SoundTypeID:           ptr(35), // Quiet
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(44), // Smooth
			},
		},
		{
			Name:             "Gateron Silent Brown",
			ShortDescription: "Silent tactile switch with a gentle bump and quiet operation.",
			LongDescription:  "Gateron Silent Brown switches offer a tactile experience with a gentle bump and quiet operation, making them ideal for users who prefer stealthy typing without sacrificing feedback.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     5,      // Silent Tactile
			ReleaseDate:      parseDate("2018-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(45.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(48), // Brown
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          true,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(3.8)), // Slightly shorter due to dampeners
				ActuationForce:        ptr(float32(45.0)),
				BottomOutForce:        ptr(float32(55.0)),
				SoundLevelID:          ptr(28), // Very Low
				SoundTypeID:           ptr(35), // Quiet
				TactilityTypeID:       ptr(38), // Leaf spring
				TactilityFeedbackID:   ptr(42), // Rounded
			},
		},
		{
			Name:             "Gateron Baby Kangaroo V2",
			ShortDescription: "Enhanced tactile switch with a strong bump and smoother operation.",
			LongDescription:  "The Gateron Kangaroo V2 switches are tactile switches known for their strong tactile bump and smooth operation. With an actuation force of 55 grams, these switches are perfect for users who enjoy a pronounced tactile feedback that is both responsive and satisfying. The Kangaroo V2 switches are designed for those who prefer a more noticeable bump during typing, making them ideal for heavy typists and those who value distinct feedback in every keystroke.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2022-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(55.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(72), // Tan (approximating kangaroo color)
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          true,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(3.5)),
				ActuationForce:        ptr(float32(55.0)),
				BottomOutForce:        ptr(float32(65.0)),
				SoundLevelID:          ptr(30), // Medium
				SoundTypeID:           ptr(34), // Thocky
				TactilityTypeID:       ptr(38), // Leaf spring
				TactilityFeedbackID:   ptr(41), // Sharp
			},
		},
		{
			Name:             "Gateron Cap Red",
			ShortDescription: "Linear switch with enhanced performance and sound dampening.",
			LongDescription:  "Gateron Cap Red switches are linear switches designed to offer smooth keystrokes with a light actuation force of 45 grams. These switches come factory-lubricated, ensuring a quiet and responsive typing experience. The Cap Red switches are part of Gateron's premium line, providing users with a reliable and high-performance switch that is perfect for gaming and fast typing.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(45.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          true,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(45.0)),
				BottomOutForce:        ptr(float32(55.0)),
				SoundLevelID:          ptr(29), // Low
				SoundTypeID:           ptr(36), // Creamy
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(44), // Smooth
			},
		},
		{
			Name:             "Gateron Ink Red",
			ShortDescription: "Premium linear switch with a smooth keystroke.",
			LongDescription:  "The Gateron Box Ink Red switches are linear switches with a unique box-style design that enhances stability and offers protection against dust and moisture. With an actuation force of 45 grams, these switches provide a smooth and responsive keystroke, making them ideal for users who prioritize durability and performance. The Box Ink Red switches are part of Gateron's premium Ink series, known for their high-quality construction and reliable operation.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2019-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(45.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(46), // Black (Ink housing)
				BottomHousingColorID:  ptr(46), // Black (Ink housing)
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          true,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(45.0)),
				BottomOutForce:        ptr(float32(55.0)),
				SoundLevelID:          ptr(29), // Low
				SoundTypeID:           ptr(34), // Thocky
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(44), // Smooth
			},
		},
		{
			Name:             "Gateron Ink Black",
			ShortDescription: "Premium linear switch with a heavier actuation force.",
			LongDescription:  "Gateron Box Ink Black switches are designed for users who prefer a heavier linear feel with the added benefits of a box-style design. With an actuation force of 60 grams, these switches provide a firm and consistent keystroke while offering enhanced protection against dust and moisture. The Box Ink Black switches are a great choice for those seeking a durable switch that can withstand heavy use while maintaining a smooth typing experience.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2019-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(60.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(46), // Black (Ink housing)
				BottomHousingColorID:  ptr(46), // Black (Ink housing)
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          true,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(60.0)),
				BottomOutForce:        ptr(float32(70.0)),
				SoundLevelID:          ptr(30), // Medium
				SoundTypeID:           ptr(34), // Thocky
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(44), // Smooth
			},
		},
		{
			Name:             "Gateron Ink Blue",
			ShortDescription: "Clicky switch with a distinct click and tactile feedback.",
			LongDescription:  "Gateron Ink Blue switches provide a clicky typing experience with a distinct click and tactile feedback, offering premium quality in the Ink series.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2020-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(60.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(46), // Black (Ink housing)
				BottomHousingColorID:  ptr(46), // Black (Ink housing)
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          false,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(60.0)),
				BottomOutForce:        ptr(float32(70.0)),
				SoundLevelID:          ptr(32), // Very High
				SoundTypeID:           ptr(33), // Clicky
				TactilityTypeID:       ptr(40), // Click Bar
				TactilityFeedbackID:   ptr(41), // Sharp
			},
		},
		{
			Name:             "Gateron Box Ink Red",
			ShortDescription: "Box-style linear switch with improved dust and moisture protection.",
			LongDescription:  "Gateron Box Ink Red switches offer a linear typing experience with enhanced dust and moisture protection, providing durability and smooth keystrokes.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(45.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(46), // Black (Ink housing)
				BottomHousingColorID:  ptr(46), // Black (Ink housing)
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          true,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(3.8)), // Slightly shorter due to box design
				ActuationForce:        ptr(float32(45.0)),
				BottomOutForce:        ptr(float32(55.0)),
				SoundLevelID:          ptr(29), // Low
				SoundTypeID:           ptr(34), // Thocky
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(44), // Smooth
			},
		},
		{
			Name:             "Gateron Box Ink Black",
			ShortDescription: "Box-style linear switch with a heavier actuation force.",
			LongDescription:  "Gateron Box Ink Black switches provide a heavier actuation force with box-style protection, offering smooth keystrokes and enhanced durability.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(60.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(46), // Black (Ink housing)
				BottomHousingColorID:  ptr(46), // Black (Ink housing)
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          true,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(3.8)), // Slightly shorter due to box design
				ActuationForce:        ptr(float32(60.0)),
				BottomOutForce:        ptr(float32(70.0)),
				SoundLevelID:          ptr(30), // Medium
				SoundTypeID:           ptr(34), // Thocky
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(44), // Smooth
			},
		},
		{
			Name:             "Gateron Box Ink Yellow",
			ShortDescription: "Box-style linear switch with smooth actuation and mid-range force.",
			LongDescription:  "Gateron Box Ink Yellow switches offer smooth actuation with mid-range force, featuring box-style protection for enhanced durability and performance.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(50.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(50), // Yellow
				TopHousingColorID:     ptr(46), // Black (Ink housing)
				BottomHousingColorID:  ptr(46), // Black (Ink housing)
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          true,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(3.8)), // Slightly shorter due to box design
				ActuationForce:        ptr(float32(50.0)),
				BottomOutForce:        ptr(float32(60.0)),
				SoundLevelID:          ptr(29), // Low
				SoundTypeID:           ptr(34), // Thocky
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(44), // Smooth
			},
		},
		{
			Name:             "Gateron KS-3 Red",
			ShortDescription: "Linear switch with a smooth feel and light actuation force.",
			LongDescription:  "Gateron KS-3 Red switches offer a smooth linear feel with light actuation, providing an affordable and reliable option for linear enthusiasts.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2015-01-01"),
			Available:        true,
			PricePoint:       1, // Value
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(45.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-pin
				FactoryLubed:          false,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(45.0)),
				BottomOutForce:        ptr(float32(55.0)),
				SoundLevelID:          ptr(30), // Medium
				SoundTypeID:           ptr(37), // Clacky
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(44), // Smooth
			},
		},
		{
			Name:             "Gateron KS-3 Brown",
			ShortDescription: "Tactile switch with a moderate actuation force and gentle bump.",
			LongDescription:  "Gateron KS-3 Brown switches provide a moderate actuation force with a gentle tactile bump, offering a balanced and affordable option for tactile enthusiasts.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2015-01-01"),
			Available:        true,
			PricePoint:       1, // Value

			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(55.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(48), // Brown
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-pin
				FactoryLubed:          false,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(55.0)),
				BottomOutForce:        ptr(float32(65.0)),
				SoundLevelID:          ptr(30), // Medium
				SoundTypeID:           ptr(37), // Clacky
				TactilityTypeID:       ptr(38), // Leaf spring
				TactilityFeedbackID:   ptr(42), // Rounded
			},
		},
		{
			Name:             "Gateron KS-3 Blue",
			ShortDescription: "Clicky switch with an audible click and tactile feedback.",
			LongDescription:  "Gateron KS-3 Blue switches offer an audible click and tactile feedback, providing a clicky typing experience at an affordable price point.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2015-01-01"),
			Available:        true,
			PricePoint:       1, // Value
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(60.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-pin
				FactoryLubed:          false,
				PreTravel:             ptr(float32(2.2)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(60.0)),
				BottomOutForce:        ptr(float32(70.0)),
				SoundLevelID:          ptr(32), // Very High
				SoundTypeID:           ptr(33), // Clicky
				TactilityTypeID:       ptr(38), // Leaf spring
				TactilityFeedbackID:   ptr(41), // Sharp
			},
		},
		{
			Name:             "Gateron KS-3 Yellow",
			ShortDescription: "Linear switch with a medium actuation force and smooth keystrokes.",
			LongDescription:  "Gateron KS-3 Yellow switches offer smooth keystrokes with a medium actuation force, providing a versatile and cost-effective option for linear enthusiasts.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2015-01-01"),
			Available:        true,
			PricePoint:       1, // Value
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(50.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(50), // Yellow
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-pin
				FactoryLubed:          false,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(50.0)),
				BottomOutForce:        ptr(float32(60.0)),
				SoundLevelID:          ptr(30), // Medium
				SoundTypeID:           ptr(37), // Clacky
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(44), // Smooth
			},
		},
		{
			Name:             "Gateron KS-3 Black",
			ShortDescription: "Linear switch with a heavier actuation force for added resistance.",
			LongDescription:  "Gateron KS-3 Black switches offer a heavier actuation force with added resistance, providing a firm linear typing experience at a value price.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2015-01-01"),
			Available:        true,
			PricePoint:       1, // Value
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(60.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-pin
				FactoryLubed:          false,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(60.0)),
				BottomOutForce:        ptr(float32(70.0)),
				SoundLevelID:          ptr(30), // Medium
				SoundTypeID:           ptr(37), // Clacky
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(44), // Smooth
			},
		},
		{
			Name:             "Gateron KS-8 Red",
			ShortDescription: "Linear switch with a light actuation force and smooth feel.",
			LongDescription:  "Gateron KS-8 Red switches offer a light actuation force with smooth keystrokes, providing an enjoyable typing experience.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2018-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(45.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          true,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(45.0)),
				BottomOutForce:        ptr(float32(55.0)),
				SoundLevelID:          ptr(29), // Low
				SoundTypeID:           ptr(36), // Creamy
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(44), // Smooth
			},
		},
		{
			Name:             "Gateron KS-8 Brown",
			ShortDescription: "Tactile switch with a gentle bump and moderate actuation force.",
			LongDescription:  "Gateron KS-8 Brown switches provide a gentle tactile bump with moderate force, offering a balanced typing experience.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2018-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(55.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(48), // Brown
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          true,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(55.0)),
				BottomOutForce:        ptr(float32(65.0)),
				SoundLevelID:          ptr(30), // Medium
				SoundTypeID:           ptr(36), // Creamy
				TactilityTypeID:       ptr(38), // Leaf spring
				TactilityFeedbackID:   ptr(42), // Rounded
			},
		},
		{
			Name:             "Gateron KS-8 Blue",
			ShortDescription: "Clicky switch with an audible click and tactile feedback.",
			LongDescription:  "Gateron KS-8 Blue switches offer an audible click with tactile feedback, ideal for users who enjoy clicky typing.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2018-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(60.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          false,
				PreTravel:             ptr(float32(2.2)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(60.0)),
				BottomOutForce:        ptr(float32(70.0)),
				SoundLevelID:          ptr(32), // Very High
				SoundTypeID:           ptr(33), // Clicky
				TactilityTypeID:       ptr(38), // Leaf spring
				TactilityFeedbackID:   ptr(41), // Sharp
			},
		},
		{
			Name:             "Gateron KS-8 Yellow",
			ShortDescription: "Linear switch with a medium actuation force and enhanced smoothness.",
			LongDescription:  "Gateron KS-8 Yellow switches offer a medium actuation force with enhanced smoothness, providing a satisfying typing experience.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2018-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(50.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(50), // Yellow
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          true,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(50.0)),
				BottomOutForce:        ptr(float32(60.0)),
				SoundLevelID:          ptr(29), // Low
				SoundTypeID:           ptr(36), // Creamy
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(44), // Smooth
			},
		},
		{
			Name:             "Gateron KS-8 Black",
			ShortDescription: "Linear switch with a heavier actuation force for added resistance.",
			LongDescription:  "Gateron KS-8 Black switches offer a heavier actuation force for users who prefer more resistance in their keystrokes, providing a firm typing experience.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2018-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(60.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          true,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(60.0)),
				BottomOutForce:        ptr(float32(70.0)),
				SoundLevelID:          ptr(30), // Medium
				SoundTypeID:           ptr(36), // Creamy
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(44), // Smooth
			},
		},
		{
			Name:             "Gateron Luciola",
			ShortDescription: "Linear switch with a glow-in-the-dark housing and smooth keystrokes.",
			LongDescription:  "The Gateron Luciola switch features a glow-in-the-dark housing with smooth keystrokes, offering a premium aesthetic and performance.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(45.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate (glow-in-the-dark)
				BaseHousingMaterialID: ptr(26), // Polycarbonate (glow-in-the-dark)
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(52), // Transparent
				TopHousingColorID:     ptr(52), // Transparent (glow-in-the-dark)
				BottomHousingColorID:  ptr(52), // Transparent (glow-in-the-dark)
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          true,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(45.0)),
				BottomOutForce:        ptr(float32(55.0)),
				SoundLevelID:          ptr(29), // Low
				SoundTypeID:           ptr(34), // Thocky
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(44), // Smooth
			},
		},
		{
			Name:             "Gateron North Pole V2 (43g)",
			ShortDescription: "Light linear switch with transparent housing, optimized for RGB lighting.",
			LongDescription:  "The Gateron North Pole V2 43g variant offers a light linear typing experience with a transparent housing designed for optimal RGB lighting. With a low actuation force of 43g, this switch provides effortless keystrokes, making it ideal for users who prefer a gentle touch or for extended typing sessions. The transparent housing not only enhances RGB effects but also contributes to a unique sound profile, balancing smoothness and acoustic feedback.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2022-01-01"),
			Available:        true,
			PricePoint:       3,  // Expensive
			SiteURL:          "", // Update if an official URL becomes available
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(43.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(52), // Transparent
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          true,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(43.0)),
				BottomOutForce:        ptr(float32(53.0)),
				SoundLevelID:          ptr(29), // Low
				SoundTypeID:           ptr(34), // Thocky
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(44), // Smooth
			},
		},
		{
			Name:             "Gateron North Pole V2 (55g)",
			ShortDescription: "Balanced linear switch with transparent housing, optimized for RGB lighting.",
			LongDescription:  "The Gateron North Pole V2 55g variant provides a balanced linear typing experience with a transparent housing designed for optimal RGB lighting. With a moderate actuation force of 55g, this switch offers a satisfying middle ground between light and heavy switches, suitable for both typing and gaming. The transparent housing not only enhances RGB effects but also contributes to a distinct sound profile, offering a blend of smoothness and auditory feedback.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2022-01-01"),
			Available:        true,
			PricePoint:       3,  // Expensive
			SiteURL:          "", // Update if an official URL becomes available
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(55.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(52), // Transparent
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          true,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(55.0)),
				BottomOutForce:        ptr(float32(65.0)),
				SoundLevelID:          ptr(30), // Medium
				SoundTypeID:           ptr(34), // Thocky
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(44), // Smooth
			},
		},
		{
			Name:             "Gateron North Pole V2 (62g)",
			ShortDescription: "Firm linear switch with transparent housing, optimized for RGB lighting.",
			LongDescription:  "The Gateron North Pole V2 62g variant delivers a firm linear typing experience with a transparent housing designed for optimal RGB lighting. With a higher actuation force of 62g, this switch provides a more substantial feel, ideal for users who prefer a stronger push or for applications requiring precise key presses. The transparent housing not only enhances RGB effects but also contributes to a unique sound profile, balancing firmness with acoustic feedback.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2022-01-01"),
			Available:        true,
			PricePoint:       3,  // Expensive
			SiteURL:          "", // Update if an official URL becomes available
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(62.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(52), // Transparent
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          true,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(62.0)),
				BottomOutForce:        ptr(float32(72.0)),
				SoundLevelID:          ptr(30), // Medium
				SoundTypeID:           ptr(34), // Thocky
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(44), // Smooth
			},
		},
		{
			Name:             "Gateron Phantom Red",
			ShortDescription: "Linear switch with a light actuation force and smooth keystrokes.",
			LongDescription:  "Gateron Phantom Red switches offer smooth keystrokes with a light actuation force, providing a satisfying typing experience.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2020-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(45.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(53), // Smokey
				BottomHousingColorID:  ptr(53), // Smokey
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          true,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(45.0)),
				BottomOutForce:        ptr(float32(55.0)),
				SoundLevelID:          ptr(29), // Low
				SoundTypeID:           ptr(36), // Creamy
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(44), // Smooth
			},
		},
		{
			Name:             "Gateron Phantom Brown",
			ShortDescription: "Tactile switch with a gentle bump and moderate actuation force.",
			LongDescription:  "Gateron Phantom Brown switches provide a gentle tactile bump with moderate actuation force, offering a balanced typing experience.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2020-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(55.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(48), // Brown
				TopHousingColorID:     ptr(53), // Smokey
				BottomHousingColorID:  ptr(53), // Smokey
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          true,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(55.0)),
				BottomOutForce:        ptr(float32(65.0)),
				SoundLevelID:          ptr(30), // Medium
				SoundTypeID:           ptr(36), // Creamy
				TactilityTypeID:       ptr(38), // Leaf spring
				TactilityFeedbackID:   ptr(42), // Rounded
			},
		},
		{
			Name:             "Gateron Phantom Blue",
			ShortDescription: "Clicky switch with an audible click and tactile feedback.",
			LongDescription:  "Gateron Phantom Blue switches offer a clicky typing experience with an audible click and tactile feedback, providing a satisfying typing experience.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2020-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(60.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(53), // Smokey
				BottomHousingColorID:  ptr(53), // Smokey
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          false,
				PreTravel:             ptr(float32(2.2)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(60.0)),
				BottomOutForce:        ptr(float32(70.0)),
				SoundLevelID:          ptr(32), // Very High
				SoundTypeID:           ptr(33), // Clicky
				TactilityTypeID:       ptr(38), // Leaf spring
				TactilityFeedbackID:   ptr(41), // Sharp
			},
		},
		{
			Name:             "Gateron Melodic",
			ShortDescription: "Tactile switch with a musical sound profile and smooth operation.",
			LongDescription:  "The Gateron Melodic switch is a tactile switch designed to provide a unique typing experience with a musical sound profile. With an actuation force of 55 grams, the Melodic switch offers a gentle tactile bump along with a distinctive sound that sets it apart from other switches. This switch is ideal for users who enjoy a tactile feedback with an added element of auditory satisfaction, making typing not only functional but also enjoyable.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2021-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(55.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(55), // Purple (assuming a unique color for this switch)
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          true,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(55.0)),
				BottomOutForce:        ptr(float32(65.0)),
				SoundLevelID:          ptr(30), // Medium
				SoundTypeID:           ptr(34), // Thocky
				TactilityTypeID:       ptr(38), // Leaf spring
				TactilityFeedbackID:   ptr(42), // Rounded
			},
		},
		{
			Name:             "Gateron KS-9 Red",
			ShortDescription: "Linear switch with improved smoothness and sound.",
			LongDescription:  "Gateron KS-9 Red switches are enhanced linear switches known for their improved smoothness and sound. With a light actuation force of 45g, these switches offer a buttery smooth typing experience with minimal resistance. The KS-9 series features upgraded materials and manufacturing processes, resulting in reduced wobble and a more consistent feel compared to earlier Gateron switches. Ideal for both gaming and typing, KS-9 Red switches provide a quiet and effortless keystroke.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2020-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(45.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          true,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(45.0)),
				BottomOutForce:        ptr(float32(55.0)),
				SoundLevelID:          ptr(29), // Low
				SoundTypeID:           ptr(36), // Creamy
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(44), // Smooth
			},
		},
		{
			Name:             "Gateron KS-9 Brown",
			ShortDescription: "Tactile switch with a gentle bump and enhanced feel.",
			LongDescription:  "Gateron KS-9 Brown switches offer a gentle tactile bump and enhanced feel. With an actuation force of 55g, these switches provide a noticeable but not overwhelming tactile feedback. The KS-9 series improvements result in a more refined tactile experience with reduced wobble and improved consistency. These switches strike a balance between typing comfort and tactile response, making them suitable for both work and play.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2020-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(55.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(48), // Brown
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          true,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(55.0)),
				BottomOutForce:        ptr(float32(65.0)),
				SoundLevelID:          ptr(30), // Medium
				SoundTypeID:           ptr(36), // Creamy
				TactilityTypeID:       ptr(38), // Leaf spring
				TactilityFeedbackID:   ptr(42), // Rounded
			},
		},
		{
			Name:             "Gateron KS-9 Blue",
			ShortDescription: "Clicky switch with an audible click and tactile feedback.",
			LongDescription:  "Gateron KS-9 Blue switches offer an audible click with tactile feedback. With an actuation force of 60g, these switches provide a distinct clicking sound and a sharp tactile bump. The KS-9 series improvements result in a more consistent click and reduced wobble. These switches are perfect for typists who enjoy auditory feedback and a crisp typing experience.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2020-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(60.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          false,
				PreTravel:             ptr(float32(2.2)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(60.0)),
				BottomOutForce:        ptr(float32(70.0)),
				SoundLevelID:          ptr(32), // Very High
				SoundTypeID:           ptr(33), // Clicky
				TactilityTypeID:       ptr(38), // Leaf spring
				TactilityFeedbackID:   ptr(41), // Sharp
			},
		},
		{
			Name:             "Gateron KS-9 Yellow",
			ShortDescription: "Linear switch with a medium actuation force and smooth keystrokes.",
			LongDescription:  "Gateron KS-9 Yellow switches offer a medium actuation force with smooth keystrokes. With an actuation force of 50g, these switches provide a balanced typing experience that's neither too light nor too heavy. The KS-9 series improvements result in enhanced smoothness and reduced wobble. These versatile switches are popular among both typists and gamers for their comfortable middle-ground characteristics.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2020-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(50.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(50), // Yellow
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          true,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(50.0)),
				BottomOutForce:        ptr(float32(60.0)),
				SoundLevelID:          ptr(29), // Low
				SoundTypeID:           ptr(36), // Creamy
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(44), // Smooth
			},
		},
		{
			Name:             "Gateron KS-9 Black",
			ShortDescription: "Linear switch with a heavier actuation force for added resistance.",
			LongDescription:  "Gateron KS-9 Black switches offer a heavier actuation force, providing a firm typing experience. With an actuation force of 60g, these switches are ideal for users who prefer more resistance in their keystrokes. The KS-9 series improvements result in enhanced smoothness and reduced wobble, even with the heavier spring. These switches are perfect for those who enjoy a more substantial feel and want to minimize accidental key presses.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2020-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(60.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          true,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(60.0)),
				BottomOutForce:        ptr(float32(70.0)),
				SoundLevelID:          ptr(30), // Medium
				SoundTypeID:           ptr(36), // Creamy
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(44), // Smooth
			},
		},
		{
			Name:             "Gateron KS-9 Green",
			ShortDescription: "Clicky switch with a strong tactile bump and audible click.",
			LongDescription:  "Gateron KS-9 Green switches provide a strong tactile bump with an audible click. With a higher actuation force of 65g, these switches offer a more pronounced tactile and auditory feedback compared to Blue switches. The KS-9 series improvements result in a more consistent click and reduced wobble. These switches are ideal for users who enjoy a firm keystroke with a satisfying click, perfect for typing-intensive tasks.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2020-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(65.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(54), // Green
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          false,
				PreTravel:             ptr(float32(2.2)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(65.0)),
				BottomOutForce:        ptr(float32(75.0)),
				SoundLevelID:          ptr(32), // Very High
				SoundTypeID:           ptr(33), // Clicky
				TactilityTypeID:       ptr(38), // Leaf spring
				TactilityFeedbackID:   ptr(41), // Sharp
			},
		},
		{
			Name:             "Gateron KS-9 Clear",
			ShortDescription: "Linear switch with a light actuation force and smooth feel.",
			LongDescription:  "Gateron KS-9 Clear switches offer a light actuation force and smooth feel, ideal for users who prefer a softer typing experience. With an actuation force of 35g, these switches provide an ultra-light keystroke, making them perfect for extended typing sessions or for those with sensitive fingers. The KS-9 series improvements result in enhanced smoothness and reduced wobble, even with the lighter spring. These switches offer a unique typing experience that's both effortless and responsive.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2020-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(35.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(49), // Clear
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          true,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(35.0)),
				BottomOutForce:        ptr(float32(45.0)),
				SoundLevelID:          ptr(28), // Very Low
				SoundTypeID:           ptr(36), // Creamy
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(44), // Smooth
			},
		},
		{
			Name:             "Gateron Milky Red",
			ShortDescription: "Linear switch with a milky housing for a softer feel.",
			LongDescription:  "Gateron Milky Red switches feature a milky housing that provides a softer feel and muted sound, offering a smooth linear experience. With an actuation force of 45 grams, these switches are designed for users who seek a light and responsive typing experience. The milky housing not only enhances the aesthetic appeal of the switch but also contributes to a quieter and softer bottom-out, making it suitable for a wide range of applications.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2017-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(45.0)),
				TopHousingMaterialID:  ptr(24), // Nylon (milky)
				BaseHousingMaterialID: ptr(24), // Nylon (milky)
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(51), // White (milky)
				BottomHousingColorID:  ptr(51), // White (milky)
				PinConfigurationID:    ptr(74), // 3-pin
				FactoryLubed:          false,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(45.0)),
				BottomOutForce:        ptr(float32(55.0)),
				SoundLevelID:          ptr(29), // Low
				SoundTypeID:           ptr(34), // Thocky
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(44), // Smooth
			},
		},
		{
			Name:             "Gateron Milky Brown",
			ShortDescription: "Tactile switch with a milky housing for a unique feel.",
			LongDescription:  "Gateron Milky Brown switches provide a gentle tactile bump with a milky housing, offering a unique typing experience with a softer bottom-out feel. With an actuation force of 55 grams, these switches are ideal for users who prefer tactile feedback without the loud click. The milky housing also helps to produce a more subdued sound, making these switches a great choice for quieter environments.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2017-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(55.0)),
				TopHousingMaterialID:  ptr(24), // Nylon (milky)
				BaseHousingMaterialID: ptr(24), // Nylon (milky)
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(48), // Brown
				TopHousingColorID:     ptr(51), // White (milky)
				BottomHousingColorID:  ptr(51), // White (milky)
				PinConfigurationID:    ptr(74), // 3-pin
				FactoryLubed:          false,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(55.0)),
				BottomOutForce:        ptr(float32(65.0)),
				SoundLevelID:          ptr(30), // Medium
				SoundTypeID:           ptr(34), // Thocky
				TactilityTypeID:       ptr(38), // Leaf spring
				TactilityFeedbackID:   ptr(42), // Rounded
			},
		},
		{
			Name:             "Gateron Milky Blue",
			ShortDescription: "Clicky switch with a milky housing for a softer sound.",
			LongDescription:  "Gateron Milky Blue switches offer a clicky typing experience with a softer sound profile, thanks to the milky housing design. With an actuation force of 60 grams, these switches deliver a pronounced tactile bump and an audible click, providing a satisfying and responsive typing experience. The milky housing not only adds to the switch's aesthetic appeal but also dampens the sound slightly, making it a more subtle clicky option.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2017-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(60.0)),
				TopHousingMaterialID:  ptr(24), // Nylon (milky)
				BaseHousingMaterialID: ptr(24), // Nylon (milky)
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(51), // White (milky)
				BottomHousingColorID:  ptr(51), // White (milky)
				PinConfigurationID:    ptr(74), // 3-pin
				FactoryLubed:          false,
				PreTravel:             ptr(float32(2.2)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(60.0)),
				BottomOutForce:        ptr(float32(70.0)),
				SoundLevelID:          ptr(31), // High (but softer than standard Blue)
				SoundTypeID:           ptr(33), // Clicky
				TactilityTypeID:       ptr(38), // Leaf spring
				TactilityFeedbackID:   ptr(41), // Sharp
			},
		},
		{
			Name:             "Gateron Milky Black",
			ShortDescription: "Linear switch with a heavier actuation force and milky housing.",
			LongDescription:  "Gateron Milky Black switches provide a firm linear feel with a milky housing, offering a unique typing experience with added resistance. With an actuation force of 60 grams, these switches are designed for users who prefer a heavier keystroke. The milky housing contributes to a softer bottom-out and a more muted sound profile, making these switches ideal for those who want a quieter, yet still robust, linear switch.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2017-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(60.0)),
				TopHousingMaterialID:  ptr(24), // Nylon (milky)
				BaseHousingMaterialID: ptr(24), // Nylon (milky)
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(51), // White (milky)
				BottomHousingColorID:  ptr(51), // White (milky)
				PinConfigurationID:    ptr(74), // 3-pin
				FactoryLubed:          false,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(60.0)),
				BottomOutForce:        ptr(float32(70.0)),
				SoundLevelID:          ptr(29), // Low
				SoundTypeID:           ptr(34), // Thocky
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(44), // Smooth
			},
		},
		{
			Name:             "Gateron Milky Green",
			ShortDescription: "Clicky switch with a strong tactile bump and milky housing.",
			LongDescription:  "Gateron Milky Green switches offer a strong tactile bump with a milky housing, providing a unique clicky typing experience with a softer sound. With an actuation force of 80 grams, these switches are designed for users who prefer a heavier clicky switch with pronounced feedback. The milky housing not only enhances the switch's visual appeal but also helps to reduce the sharpness of the click sound, making it a more refined option for clicky switch enthusiasts.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2017-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(80.0)),
				TopHousingMaterialID:  ptr(24), // Nylon (milky)
				BaseHousingMaterialID: ptr(24), // Nylon (milky)
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(54), // Green
				TopHousingColorID:     ptr(51), // White (milky)
				BottomHousingColorID:  ptr(51), // White (milky)
				PinConfigurationID:    ptr(74), // 3-pin
				FactoryLubed:          false,
				PreTravel:             ptr(float32(2.2)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(80.0)),
				BottomOutForce:        ptr(float32(90.0)),
				SoundLevelID:          ptr(31), // High (but softer than standard Green)
				SoundTypeID:           ptr(33), // Clicky
				TactilityTypeID:       ptr(38), // Leaf spring
				TactilityFeedbackID:   ptr(41), // Sharp
			},
		},
		{
			Name:             "Gateron Milky Clear",
			ShortDescription: "Linear switch with a light actuation force and milky housing.",
			LongDescription:  "Gateron Milky Clear switches offer a light actuation force and smooth keystrokes with a milky housing, providing a softer feel and quieter operation. With an actuation force of 35 grams, these switches are perfect for users who prefer a very light and effortless typing experience. The milky housing also contributes to a more subdued sound profile, making these switches ideal for those who value both performance and quietness in their typing setup.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2017-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(35.0)),
				TopHousingMaterialID:  ptr(24), // Nylon (milky)
				BaseHousingMaterialID: ptr(24), // Nylon (milky)
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(49), // Clear
				TopHousingColorID:     ptr(51), // White (milky)
				BottomHousingColorID:  ptr(51), // White (milky)
				PinConfigurationID:    ptr(74), // 3-pin
				FactoryLubed:          false,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(35.0)),
				BottomOutForce:        ptr(float32(45.0)),
				SoundLevelID:          ptr(28), // Very Low
				SoundTypeID:           ptr(34), // Thocky
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(44), // Smooth
			},
		},
		{
			Name:             "Gateron Box Red",
			ShortDescription: "Linear switch with box design for added stability.",
			LongDescription:  "Gateron Box Red switches offer a smooth linear experience with a box design that enhances stability and protection against dust. With an actuation force of 45 grams, these switches provide a light and responsive feel. The box design not only improves the switch's durability but also contributes to a more consistent feel and reduced wobble. Ideal for both typing and gaming, Box Red switches deliver a premium linear experience with added reliability.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2020-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(45.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          true,
				PreTravel:             ptr(float32(1.8)),
				TotalTravel:           ptr(float32(3.6)),
				ActuationForce:        ptr(float32(45.0)),
				BottomOutForce:        ptr(float32(55.0)),
				SoundLevelID:          ptr(29), // Low
				SoundTypeID:           ptr(36), // Creamy
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(44), // Smooth
			},
		},
		{
			Name:             "Gateron Box Brown",
			ShortDescription: "Tactile switch with box design for enhanced feel.",
			LongDescription:  "Gateron Box Brown switches offer a gentle tactile bump with a box design, providing a stable typing experience and protection against dust. With an actuation force of 55 grams, these switches deliver a noticeable but not overwhelming tactile feedback. The box design enhances the switch's durability and contributes to a more consistent feel with reduced wobble. Box Brown switches are perfect for users who enjoy tactile feedback in a premium package.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2020-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(55.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(48), // Brown
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          true,
				PreTravel:             ptr(float32(1.8)),
				TotalTravel:           ptr(float32(3.6)),
				ActuationForce:        ptr(float32(55.0)),
				BottomOutForce:        ptr(float32(65.0)),
				SoundLevelID:          ptr(30), // Medium
				SoundTypeID:           ptr(36), // Creamy
				TactilityTypeID:       ptr(38), // Leaf spring
				TactilityFeedbackID:   ptr(42), // Rounded
			},
		},
		{
			Name:             "Gateron Box Blue",
			ShortDescription: "Clicky switch with box design for a crisp click.",
			LongDescription:  "Gateron Box Blue switches provide a clicky typing experience with a box design, enhancing the crispness of each click and offering stability. With an actuation force of 60 grams, these switches deliver a pronounced tactile bump and an audible click. The box design not only improves durability but also contributes to a more consistent feel and reduced wobble. Box Blue switches are ideal for users who enjoy a satisfying clicky feedback in a premium package.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2020-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(60.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          false,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(3.8)),
				ActuationForce:        ptr(float32(60.0)),
				BottomOutForce:        ptr(float32(70.0)),
				SoundLevelID:          ptr(32), // Very High
				SoundTypeID:           ptr(33), // Clicky
				TactilityTypeID:       ptr(38), // Leaf spring
				TactilityFeedbackID:   ptr(41), // Sharp
			},
		},
		{
			Name:             "Gateron Box Black",
			ShortDescription: "Linear switch with a heavier feel and box design.",
			LongDescription:  "Gateron Box Black switches offer a heavier linear feel with a box design, providing stability and protection against dust and debris. With an actuation force of 60 grams, these switches deliver a firmer typing experience that's ideal for users who prefer more resistance. The box design enhances durability and contributes to a more consistent feel with reduced wobble. Box Black switches are perfect for those seeking a premium heavy linear switch with added reliability.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2020-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(60.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          true,
				PreTravel:             ptr(float32(1.8)),
				TotalTravel:           ptr(float32(3.6)),
				ActuationForce:        ptr(float32(60.0)),
				BottomOutForce:        ptr(float32(70.0)),
				SoundLevelID:          ptr(30), // Medium
				SoundTypeID:           ptr(36), // Creamy
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(44), // Smooth
			},
		},
		{
			Name:             "Gateron Silent Clear",
			ShortDescription: "Silent linear switch with a light actuation force.",
			LongDescription:  "Gateron Silent Clear switches offer a quiet linear typing experience with a light actuation force, ideal for stealthy typing environments. With an actuation force of 35 grams, these switches provide an ultra-light keystroke while maintaining near-silent operation. The Silent Clear switches feature special dampeners to reduce both bottom-out and upstroke noise, making them perfect for office environments or late-night typing sessions where noise reduction is crucial.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     4,      // Silent Linear
			ReleaseDate:      parseDate("2019-01-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(35.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(49), // Clear
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          true,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(3.8)), // Slightly shorter due to dampeners
				ActuationForce:        ptr(float32(35.0)),
				BottomOutForce:        ptr(float32(45.0)),
				SoundLevelID:          ptr(28), // Very Low
				SoundTypeID:           ptr(35), // Quiet
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(44), // Smooth
			},
		},
		{
			Name:             "Gateron Milu",
			ShortDescription: "Tactile switch designed for a unique typing experience.",
			LongDescription:  "Gateron Milu switches are designed to offer a unique tactile typing experience, providing a satisfying tactile bump and smooth keystrokes. With an estimated actuation force of 55 grams, these switches deliver a pronounced tactile feedback without being overly heavy. The Milu switches are engineered to provide a balance between tactile response and typing comfort, making them ideal for enthusiasts who appreciate a distinct tactile sensation without the audible click of clicky switches.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2022-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(55.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(55), // Purple (assuming a unique color)
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          true,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(55.0)),
				BottomOutForce:        ptr(float32(65.0)),
				SoundLevelID:          ptr(30), // Medium
				SoundTypeID:           ptr(34), // Thocky
				TactilityTypeID:       ptr(38), // Leaf spring
				TactilityFeedbackID:   ptr(41), // Sharp
			},
		},
		{
			Name:             "Gateron Hippo",
			ShortDescription: "Linear switch with a thocky sound profile.",
			LongDescription:  "Gateron Hippo switches are linear switches renowned for their smooth keystrokes and distinctive thocky sound profile. With an actuation force of 50 grams, these switches are designed for users who appreciate a smooth and consistent typing experience with a deep, satisfying sound. The Hippo switches feature a unique combination of materials and design that contribute to their characteristic sound. They are a perfect choice for enthusiasts seeking a switch that combines performance with a unique auditory experience, ideal for those who enjoy a rich, low-pitched typing sound.",
			ManufacturerID:   ptr(2), // Gateron
			BrandID:          ptr(2), // Gateron
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2022-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringTypeID:          ptr(6),
				SpringForce:           ptr(float32(50.0)),
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(23), // POM
				StemMaterialID:        ptr(23), // POM
				StemColorID:           ptr(55), // Purple (assuming a unique color)
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-pin
				FactoryLubed:          true,
				PreTravel:             ptr(float32(2.0)),
				TotalTravel:           ptr(float32(4.0)),
				ActuationForce:        ptr(float32(50.0)),
				BottomOutForce:        ptr(float32(60.0)),
				SoundLevelID:          ptr(30), // Medium
				SoundTypeID:           ptr(34), // Thocky
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(44), // Smooth
			},
		},
	}
	err := processSwitches(tx, gateron, admin)
	if err != nil {
		return err
	}

	return nil
}
