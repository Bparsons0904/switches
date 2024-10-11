package seed

import (
	"switches/models"

	"gorm.io/gorm"
)

func seedCherry(tx *gorm.DB, admin models.User) error {
	cherry := []models.Switch{
		{
			Name:             "Cherry MX2A Silent Red",
			ShortDescription: "Silent linear switch with RGB support and a smooth feel.",
			LongDescription:  "The Cherry MX2A Silent Red Linear switch features a quiet and smooth keystroke, making it perfect for quiet environments such as offices or shared spaces. With an actuation force of 45 grams, this switch ensures a light touch while minimizing noise through built-in dampeners. The RGB lighting adds a customizable visual element, making it an excellent choice for those who want both a serene typing experience and a vibrant keyboard.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     4,      // Silent Linear
			ReleaseDate:      parseDate("2024-03-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			SiteURL:          "https://www.cherry-world.com/mx2a-silent-red",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx3a-l1na_l1nn_image.jpg",
					AltText:     "Cherry MX2A Silent Red switch",
					OwnerType:   "switches",
				},
				{
					LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx3a-l1nn_image.jpg",
					AltText:     "Cherry MX2A Silent Red switch",
					OwnerType:   "switches",
				},
				{
					LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx3a-l1na_image.jpg",
					AltText:     "Cherry MX2A Silent Red switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(60)),
				BottomOutForcePoint:   ptr(float32(4.0)),
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
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Cherry MX Silent Black",
			ShortDescription: "Silent linear switch with a heavier actuation force.",
			LongDescription:  "Cherry MX Silent Black switches are engineered to provide the quiet operation of Silent Reds with a higher actuation force of 60 grams. This switch is ideal for users who prefer a more deliberate keystroke, combining the smooth, linear action with noise-reducing technology. Silent Black switches are well-suited for environments that require both quiet and a firm keypress, balancing performance with sound reduction.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     4,      // Silent Linear
			ReleaseDate:      parseDate("2015-01-01"),
			Available:        false,
			PricePoint:       3,
			SiteURL:          "https://www.cherry-world.com/mx-silent-black.html",
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(80)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
			},
		},
		{
			Name:             "Cherry MX Silent Red",
			ShortDescription: "Silent linear switch with a light actuation force.",
			LongDescription:  "Cherry MX Silent Red switches are designed to offer the smooth, linear feel of MX Red switches while significantly reducing noise during actuation. These switches incorporate built-in dampeners that minimize the sound of both the key press and release, making them an excellent choice for quiet environments like offices or shared spaces. With an actuation force of 45 grams, they maintain a light, responsive feel, ideal for fast typing or gaming without the auditory distraction.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     4,      // Silent Linear
			ReleaseDate:      parseDate("2015-01-01"),
			Available:        false,
			PricePoint:       3,
			SiteURL:          "",
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(60)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
			},
		},
		{
			Name:             "Cherry MX2A Silent Black",
			ShortDescription: "Silent linear switch with a heavy actuation force.",
			LongDescription:  "The Cherry MX2A Silent Black switch offers a quiet, smooth linear typing experience with a heavier actuation force of 60 grams. Designed to minimize noise, this switch features built-in dampeners that significantly reduce the sound of both key press and release, making it ideal for quiet environments such as offices or shared spaces. The MX2A Silent Black is perfect for users who prefer a firmer keystroke while maintaining a low-noise profile, ensuring both performance and discretion in their typing setup.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     4,      // Silent Linear
			ReleaseDate:      parseDate("2024-03-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			SiteURL:          "https://www.cherry-world.com/mx-silent-black",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx3a-11na_11nn_image.jpg",
					AltText:     "Cherry MX2A Silent Black switch",
					OwnerType:   "switches",
				},
				{
					LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx-silent-black.jpg",
					AltText:     "Cherry MX2A Silent Black switch opened",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(80)),
				BottomOutForcePoint:   ptr(float32(4.0)),
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
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
			},
		},

		{
			Name:             "Cherry MX2A Brown Tactile",
			ShortDescription: "Tactile switch with a gentle bump for feedback.",
			LongDescription:  "The Cherry MX2A Brown Tactile switch provides a subtle tactile bump, making it suitable for users who prefer feedback without the accompanying click sound. With an actuation force of 55 grams, this switch offers a comfortable balance between typing and gaming, providing just enough feedback to avoid bottoming out while maintaining a smooth and controlled keypress.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2024-03-01"),
			Available:        true,
			PricePoint:       2, // Average
			SiteURL:          "https://www.cherry-world.com/mx2a-brown",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/9380e5bfba4d8931eee53843958a2c83/m/x/mx1a-g1na_g1nn_image.jpg",
					AltText:     "Cherry MX2A Brown switch",
					OwnerType:   "switches",
				},
				{
					LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-g1nn_explo.jpg",
					AltText:     "Cherry MX2A Brown switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(65)),
				BottomOutForcePoint:   ptr(float32(4.0)),
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
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(30), // Medium
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(42), // Rounded
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Cherry MX2A Red Linear",
			ShortDescription: "Smooth linear switch for a balanced typing feel.",
			LongDescription:  "The Cherry MX2A Red Linear switch offers a balanced and smooth typing experience, suitable for both gaming and typing due to its medium actuation force of 45 grams. This switch is designed for users who prefer a light, responsive keystroke with no tactile bump or audible click, allowing for fluid and uninterrupted keypresses.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2024-03-01"),
			Available:        true,
			PricePoint:       2, // Average
			SiteURL:          "https://www.cherry-world.com/mx2a-red",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-l1nn_image.jpg",
					AltText:     "Cherry MX2A Red switch",
					OwnerType:   "switches",
				},
				{
					LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-l1nn_explo.jpg",
					AltText:     "Cherry MX2A Red opened switch",
					OwnerType:   "switches",
				},
				{
					LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-l1na_image.jpg",
					AltText:     "Cherry MX2A Red switch",
					OwnerType:   "switches",
				},
				{
					LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-l1nn_explo.jpg",
					AltText:     "Cherry MX2A Red opened switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(60)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(23), // POM
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
			Name:             "Cherry MX2A Speed Silver Linear",
			ShortDescription: "Fast linear switch for quick actuation.",
			LongDescription:  "The Cherry MX2A Speed Silver Linear switch is designed for rapid keystrokes, making it ideal for gaming and other high-speed typing scenarios. With a short 1.2mm actuation distance and 45 grams of force, this switch enables quick responses, helping to improve performance in fast-paced environments. The MX2A Speed Silver is a top choice for competitive gamers seeking the fastest actuation possible.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2024-03-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			SiteURL:          "https://www.cherry-world.com/mx2a-speed-silver",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-51nn_image.jpg",
					AltText:     "Cherry MX2A Speed Silver switch",
					OwnerType:   "switches",
				},
				{
					LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-51nn_sideview4.jpg",
					AltText:     "Cherry MX2A Speed Silver switch",
					OwnerType:   "switches",
				},
				{
					LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-51na_image.jpg",
					AltText:     "Cherry MX2A Speed Silver switch",
					OwnerType:   "switches",
				},
				{
					LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-51na_sideview4.jpg",
					AltText:     "Cherry MX2A Speed Silver switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(1.2)),
				BottomOutForce:        ptr(float32(60)),
				BottomOutForcePoint:   ptr(float32(3.4)),
				TotalTravel:           ptr(float32(3.4)),
				PreTravel:             ptr(float32(1.2)),
				FactoryLubed:          false,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(59), // Silver
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(37), // Clacky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Cherry MX2A Black Linear",
			ShortDescription: "Linear switch with a heavy actuation force.",
			LongDescription:  "The Cherry MX2A Black Linear switch offers a smooth keystroke with a heavier actuation force of 60 grams. This makes it ideal for users who prefer a firm, controlled typing experience. The switch is well-suited for both heavy typists and gamers who require a more deliberate keypress. With its linear feel and no tactile bump, the MX2A Black ensures a consistent and smooth actuation with every keystroke.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2024-03-01"),
			Available:        true,
			PricePoint:       2, // Average
			SiteURL:          "https://www.cherry-world.com/mx2a-black",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-11nn_image.jpg",
					AltText:     "Cherry MX2A Black switch",
					OwnerType:   "switches",
				},
				{
					LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-11nn_explo.jpg",
					AltText:     "Cherry MX2A Black opened switch",
					OwnerType:   "switches",
				},
				{
					LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-11na_image.jpg",
					AltText:     "Cherry MX2A Black switch",
					OwnerType:   "switches",
				},
				{
					LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-11na_explo.jpg",
					AltText:     "Cherry MX2A Black opened switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				ActuationPoint:        ptr(float32(2.0)),
				BottomOutForce:        ptr(float32(80)),
				BottomOutForcePoint:   ptr(float32(4.0)),
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
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Cherry MX2A Clear",
			ShortDescription: "Tactile switch with a high actuation force.",
			LongDescription:  "The Cherry MX2A Clear switch is designed for users who prefer a strong tactile feedback without the click sound. With an actuation force of 65 grams, this switch offers a pronounced tactile bump that is ideal for heavy typists who enjoy a firm and responsive keystroke. The MX2A Clear is particularly suited for those who need a tactile switch that provides clear feedback for every keypress, making it a great choice for both typing and precise input tasks.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2024-03-01"),
			Available:        true,
			PricePoint:       3,                                       // Expensive
			SiteURL:          "https://www.cherry-world.com/mx-clear", // If an official URL is available, you can update it here
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-c1na_c1nn_image.jpg",
					AltText:     "Cherry MX2A Clear switch",
					OwnerType:   "switches",
				},
				{
					LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-c1nn_image.jpg",
					AltText:     "Cherry MX2A Clear switch",
					OwnerType:   "switches",
				},
			},
		},
		{
			Name:             "Cherry MX2A Grey",
			ShortDescription: "Tactile switch with an even higher actuation force.",
			LongDescription:  "The Cherry MX2A Grey switch is designed for users who prefer a very strong tactile feedback with an actuation force of 80 grams. This switch offers a firm tactile bump, providing clear and pronounced feedback with every keystroke. The MX2A Grey is perfect for users who require a high level of resistance in their switches, making it ideal for precise and deliberate key presses in both typing and heavy-duty applications.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2024-03-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			SiteURL:          "https://www.cherry-world.com/mx-grey-tactile",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-d1na_d1nn_image.jpg",
					AltText:     "Cherry MX2A Grey switch",
					OwnerType:   "switches",
				},
				{
					LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-d1nn_sideview4.jpg",
					AltText:     "Cherry MX2A Grey switch",
					OwnerType:   "switches",
				},
				{
					LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-d1na_top",
					AltText:     "Cherry MX2A Grey switch",
					OwnerType:   "switches",
				},
			},
		},
		{
			Name:             "Cherry MX Purple",
			ShortDescription: "Linear switch with a light actuation force and unique purple stem.",
			LongDescription:  "The Cherry MX Purple switch is a linear switch designed for users who prefer a smooth and consistent keystroke with a light actuation force of 45 grams. This switch features a unique purple stem, which adds a distinct aesthetic to any keyboard. The MX Purple is ideal for both gaming and typing, offering a balance between speed and precision while providing a visually appealing design. It is a popular choice for users looking to combine performance with style.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2024-03-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			SiteURL:          "https://www.cherry-world.com/mx-purple",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx2a-v1nw_imageview.jpg",
					AltText:     "Cherry MX Purple switch",
					OwnerType:   "switches",
				},
				{
					LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx2a-v1nw_explo.jpg",
					AltText:     "Cherry MX Purple switch",
					OwnerType:   "switches",
				},
			},
		},
		{
			Name:             "Cherry MX Clear-Top",
			ShortDescription: "Tactile switch with a transparent housing for enhanced RGB lighting.",
			LongDescription:  "The Cherry MX Clear-Top switch is a tactile switch that combines the renowned tactile feedback of the Cherry MX Clear with a transparent housing, making it perfect for keyboards with RGB lighting. With an actuation force of 65 grams, this switch offers a firm tactile bump that provides a satisfying typing experience. The transparent housing allows for vibrant RGB lighting, enhancing the visual appeal of your keyboard setup while maintaining the high performance and reliability of the MX Clear switch.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2024-03-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			SiteURL:          "https://www.cherry-world.com/mx-black-clear-top",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-61nw_image.jpg",
					AltText:     "Cherry MX Clear-Top switch",
					OwnerType:   "switches",
				},
				{
					LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-61nw_explo.jpg",
					AltText:     "Cherry MX Clear-Top switch",
					OwnerType:   "switches",
				},
			},
		},
		{
			Name:             "Cherry MX2A Blue Clicky",
			ShortDescription: "Clicky switch with an audible click for tactile feedback.",
			LongDescription:  "The Cherry MX2A Blue Clicky switch delivers the satisfying click sound and tactile feedback that typists love. With an actuation force of 50 grams, this switch provides a crisp click and a noticeable bump, ensuring each keystroke is both audible and precise. Ideal for those who enjoy the classic typing experience, the MX2A Blue is perfect for heavy typists or those who simply love the sound of their keyboard.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2024-03-01"),
			Available:        true,
			PricePoint:       2, // Average
			SiteURL:          "https://www.cherry-world.com/mx2a-blue",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-e1na_e1nn_image.jpg",
					AltText:     "Cherry MX2A switch",
					OwnerType:   "switches",
				},
				{
					LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-e1nn_explo.jpg",
					AltText:     "Cherry MX2A switch",
					OwnerType:   "switches",
				},
			},
		},
		{
			Name:             "Cherry MX Blue",
			ShortDescription: "Clicky switch with a tactile bump.",
			LongDescription:  "Cherry MX Blue switches are iconic for their audible click and tactile feedback, offering a distinct typing experience that is both responsive and satisfying. With an actuation force of 50 grams, these switches provide a noticeable bump before actuation, coupled with a sharp click sound that appeals to typists who appreciate clear auditory and physical feedback. While they are less common in gaming due to their pronounced click, MX Blues are a top choice for those who value precision and clarity in their keystrokes.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("1987-01-01"),
			Available:        false,
			PricePoint:       2,
			SiteURL:          "https://www.cherry-world.com/mx-blue.html",
		},
		{
			Name:             "Cherry MX Brown",
			ShortDescription: "Tactile switch with a moderate actuation force.",
			LongDescription:  "Cherry MX Brown switches strike a balance between typing and gaming with their tactile bump and moderate actuation force of 55 grams. These switches are quieter than their Blue counterparts, lacking the audible click but still providing tactile feedback that is gentle yet noticeable. This makes them versatile and well-suited for both work and play, appealing to users who want a middle ground between the stark tactility of MX Blue switches and the smoothness of MX Red switches.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("1994-01-01"),
			Available:        false,
			PricePoint:       2,
			SiteURL:          "https://www.cherry-world.com/mx-brown.html",
		},
		{
			Name:             "Cherry MX Black",
			ShortDescription: "Linear switch with a heavier actuation force.",
			LongDescription:  "Cherry MX Black switches offer a smooth linear actuation similar to MX Red switches but with a heavier actuation force of 60 grams. This increased resistance makes them ideal for users who prefer a more deliberate keystroke, reducing the chance of accidental presses. MX Black switches are often favored by those who type heavily or for gaming environments where controlled inputs are critical. The absence of tactile bumps or clicks also ensures a quieter, more focused typing experience.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("1984-01-01"),
			Available:        false,
			PricePoint:       2,
			SiteURL:          "https://www.cherry-world.com/mx-black.html",
		},
		{
			Name:             "Cherry MX Green",
			ShortDescription: "Clicky switch with a heavy actuation force.",
			LongDescription:  "Cherry MX Green switches are a variant of the MX Blue, designed for those who enjoy a clicky experience but prefer a heavier actuation force. Requiring 80 grams to actuate, these switches provide a pronounced tactile bump and a louder click, making them ideal for users who want a more deliberate, assertive keystroke. MX Green switches are less common in standard keyboards but are highly prized by enthusiasts who enjoy a stiffer, more resistant typing experience.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2011-01-01"),
			Available:        true,
			PricePoint:       3,
			SiteURL:          "https://www.cherry-world.com/mx-green",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-f1na_f1nn_image.jpg",
					AltText:     "Cherry MX Green switch",
					OwnerType:   "switches",
				},
				{
					LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-f1nn_sideview3.jpg",
					AltText:     "Cherry MX Green switch",
					OwnerType:   "switches",
				},

				{
					LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-f1na_image.jpg",
					AltText:     "Cherry MX Green switch",
					OwnerType:   "switches",
				},
			},
		},
		{
			Name:             "Cherry MX Clear",
			ShortDescription: "Tactile switch with a high actuation force.",
			LongDescription:  "Cherry MX Clear switches are designed for those who prefer a heavier tactile bump in their typing experience. With an actuation force of 65 grams, these switches provide a noticeable resistance, making each keystroke deliberate and satisfying. MX Clear switches are often chosen by users who want a tactile switch that offers more feedback than MX Browns but without the clicky noise associated with MX Blue switches.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("1994-01-01"),
			Available:        false,
			PricePoint:       3,
			SiteURL:          "https://www.cherry-world.com/mx-clear.html",
		},
		{
			Name:             "Cherry MX Grey",
			ShortDescription: "Tactile switch with a heavy actuation force.",
			LongDescription:  "Cherry MX Grey switches are an evolution of the MX Clear, featuring an even higher actuation force of 80 grams. This provides a pronounced tactile bump and substantial resistance, making them suitable for users who prefer a firm and assertive keystroke. These switches are less common in mass-market keyboards and are often found in specialized or enthusiast builds that prioritize a heavier, more tactile typing experience.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("1994-01-01"),
			Available:        false,
			PricePoint:       3,
			SiteURL:          "https://www.cherry-world.com/mx-grey.html",
		},
		{
			Name:             "Cherry MX Speed Silver",
			ShortDescription: "Linear switch with a shorter travel distance for faster actuation.",
			LongDescription:  "Cherry MX Speed Silver switches are optimized for rapid keystrokes with a shorter actuation distance and total travel. These linear switches require just 1.2mm to actuate and only 45 grams of force, making them perfect for competitive gaming where speed is crucial. The reduced travel distance allows for faster key presses, helping to shave milliseconds off each action in fast-paced environments. They are a go-to choice for gamers seeking performance and responsiveness.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2016-01-01"),
			Available:        false,
			PricePoint:       3,
			SiteURL:          "https://www.cherry-world.com/mx-speed-silver.html",
		},
		{
			Name:             "Cherry MX Nature White",
			ShortDescription: "Linear switch with a medium actuation force, between Red and Black.",
			LongDescription:  "Cherry MX Nature White switches are designed to offer a balanced typing experience, with an actuation force that sits comfortably between the lighter MX Red and the heavier MX Black switches. With a smooth, linear action and a medium force of 55 grams, these switches provide a well-rounded keystroke that is ideal for users who want a linear switch that isn’t too light or too heavy. They are often chosen by those who prefer a slightly more resistant keystroke without sacrificing speed or smoothness.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2016-01-01"),
			Available:        false,
			PricePoint:       2,
			SiteURL:          "https://www.cherry-world.com/mx-nature-white.html",
		},
		{
			Name:             "Cherry MX Violet",
			ShortDescription: "Tactile switch with a light actuation force, similar to Brown but lighter.",
			LongDescription:  "Cherry MX Violet switches offer a tactile bump similar to MX Browns but with a lighter actuation force, providing a gentle yet satisfying typing experience. Requiring only 45 grams of force to actuate, these switches are ideal for users who appreciate tactile feedback but prefer a softer, less resistant keypress. They are well-suited for both typing and gaming, offering a balance of feedback and ease of use.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2020-01-01"),
			Available:        false,
			PricePoint:       2,
			SiteURL:          "",
		},
		{
			Name:             "Cherry MX White",
			ShortDescription: "A softer clicky switch, sometimes called 'Milk'.",
			LongDescription:  "Cherry MX White switches, often referred to as 'Milk' switches due to their smoother click, are a quieter alternative to the more common MX Blue switches. They retain the clicky feedback but with a softer, less pronounced sound, making them a great choice for users who enjoy a clicky switch but want something less intrusive in terms of noise. These switches are less common and have become a favorite among those who want a unique auditory and tactile experience.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2011-01-01"),
			Available:        false,
			PricePoint:       2,
			SiteURL:          "",
		},
		{
			Name:             "Cherry MX Tactile Grey",
			ShortDescription: "A variant of the Grey switch with a slightly different force curve.",
			LongDescription:  "Cherry MX Tactile Grey switches are a specialized variant of the standard MX Grey switches, offering a unique force curve that provides a slightly different tactile experience. These switches are designed for users who appreciate a heavy, pronounced tactile bump and prefer a switch with a firm and deliberate keystroke. They are often used in niche keyboard builds where a distinct tactile feel is desired.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("1994-01-01"),
			Available:        false,
			PricePoint:       3,
			SiteURL:          "",
		},
		{
			Name:             "Cherry MX Super Black",
			ShortDescription: "An extremely heavy linear switch.",
			LongDescription:  "Cherry MX Super Black switches are designed for users who demand an extremely heavy actuation force. With a force requirement of 150 grams, these linear switches are among the stiffest in the Cherry MX lineup, offering unparalleled resistance to accidental key presses. Super Black switches are typically used in specialized keyboards or applications where a very deliberate keypress is necessary, providing a unique typing experience that is both demanding and rewarding.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("1984-01-01"),
			Available:        false,
			PricePoint:       3,
			SiteURL:          "",
		},
		{
			Name:             "Cherry MX Red",
			ShortDescription: "Linear switch with a light actuation force.",
			LongDescription:  "Cherry MX Red switches are renowned for their smooth, linear actuation and light operating force, making them a preferred choice for gamers and typists who enjoy a fast, uninterrupted keystroke. These switches require only 45 grams of force to actuate, ensuring a quick response without any tactile bump or audible click, allowing for a fluid typing experience. They are especially favored in gaming for their consistent key presses and rapid input, reducing finger fatigue during extended use.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("1984-01-01"),
			Available:        false,
			PricePoint:       2,
			SiteURL:          "https://www.cherry-world.com/mx-red.html",
		},
	}

	err := processSwitches(tx, cherry, admin)
	if err != nil {
		return err
	}

	return nil
}
