package seed

import (
	"switches/models"

	"gorm.io/gorm"
)

func seedCherry(tx *gorm.DB, admin models.User) error {
	cherry := []models.Switch{
		{
			Name:             "MX Orange",
			ShortDescription: "A tactile switch designed for typing with a medium actuation force and distinctive tactile feedback.",
			LongDescription:  "The Cherry MX Orange is a tactile switch that offers a balanced typing experience with a medium actuation force of 55g. It features a pronounced tactile bump that provides clear feedback without being too aggressive, making it suitable for both typing and gaming. The switch is designed with durability in mind, rated for 100 million keystrokes, and offers a good middle ground between Cherry's Brown and Clear switches in terms of tactility and force requirements.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2023-06-01"),
			Available:        true,
			PricePoint:       2, // Average
			SiteURL:          "https://www.cherrymx.de/en/mx-original/mx-orange.html",
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(65)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(46), // Black
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(56), // Orange
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "MX Orange RGB",
			ShortDescription: "RGB-compatible variant of the Cherry MX Orange with transparent housing for enhanced lighting.",
			LongDescription:  "The Cherry MX Orange RGB is the illumination-optimized version of the standard MX Orange switch. It maintains the same tactile characteristics and 55g actuation force of the standard Orange, but features a fully transparent housing designed specifically for RGB lighting compatibility. The switch provides clear tactile feedback and is rated for 100 million keystrokes, making it ideal for users who want both the tactile typing experience and vibrant RGB lighting effects.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2023-06-01"),
			Available:        true,
			PricePoint:       2, // Average
			SiteURL:          "https://www.cherrymx.de/en/mx-original/mx-orange.html",
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(65)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
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
			Name:             "MX2A Silent Red",
			ShortDescription: "Silent linear switch optimized for quiet operation with a smooth keystroke.",
			LongDescription:  "The Cherry MX2A Silent Red Linear switch is engineered for near-silent operation, making it ideal for quiet environments such as offices or shared spaces. This standard version maintains the same smooth linear action and 45g actuation force as its RGB counterpart, with built-in dampeners to minimize noise during both downstroke and upstroke. The switch features Cherry's latest technology for enhanced stability and reliability, while its black bottom housing provides a solid foundation for the typing experience.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     4,      // Silent Linear
			ReleaseDate:      parseDate("2024-03-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			SiteURL:          "https://www.cherry.de/en-gb/product/mx2a-silent-red",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.cherry.de/fileadmin/_processed_/0/3/csm_MX3A-L1NN_image01_73f19afb35.jpg",
					AltText:     "Cherry MX2A Silent Red switch",
					OwnerType:   "switches",
				},
				{
					LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx3a-l1na_l1nn_image.jpg",
					AltText:     "Cherry MX2A Silent Red switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(30)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(100)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(1.9)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(46), // Black
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
			Name:             "MX2A Silent Red RGB",
			ShortDescription: "Silent linear switch with RGB support and a smooth feel.",
			LongDescription:  "The Cherry MX2A Silent Red Linear switch features a quiet and smooth keystroke, making it perfect for quiet environments such as offices or shared spaces. With an actuation force of 45 grams, this switch ensures a light touch while minimizing noise through built-in dampeners. The RGB lighting adds a customizable visual element, making it an excellent choice for those who want both a serene typing experience and a vibrant keyboard.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     4,      // Silent Linear
			ReleaseDate:      parseDate("2024-03-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			SiteURL:          "https://www.cherry.de/en-gb/product/mx2a-silent-red",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.cherry.de/fileadmin/_processed_/0/3/csm_MX3A-L1NN_image01_73f19afb35.jpg",
					AltText:     "Cherry MX2A Silent Red switch",
					OwnerType:   "switches",
				},
				{
					LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx3a-l1na_l1nn_image.jpg",
					AltText:     "Cherry MX2A Silent Red switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(30)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(100)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(1.9)),
				FactoryLubed:          ptr(true),
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
			Name:             "MX Silent Red",
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
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
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
			Name:             "MX2A Silent Black",
			ShortDescription: "Silent linear switch with a heavy actuation force.",
			LongDescription:  "The Cherry MX2A Silent Black switch offers a quiet, smooth linear typing experience with a heavier actuation force of 60 grams. Designed to minimize noise, this switch features built-in dampeners that significantly reduce the sound of both key press and release, making it ideal for quiet environments such as offices or shared spaces. The MX2A Silent Black is perfect for users who prefer a firmer keystroke while maintaining a low-noise profile, ensuring both performance and discretion in their typing setup.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     4,      // Silent Linear
			ReleaseDate:      parseDate("2024-03-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			SiteURL:          "https://www.cherry.de/en-gb/product/mx2a-silent-black",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.cherry.de/fileadmin/_processed_/f/2/csm_7e9e251eb4a5186d12922ec55dbc6c60_f4b29af1c1.jpg",
					AltText:     "Cherry MX2A Silent Black switch",
					OwnerType:   "switches",
				},
				{
					LinkAddress: "https://www.cherry.de/fileadmin/_processed_/b/f/csm_e9f168ea1d27e84b87044b8740db7c2f_9c51413c04.jpg",
					AltText:     "Cherry MX2A Silent Black switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(60)),
				BottomOutForce:        ptr(float32(110)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(53), // Smokey
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "MX2A Silent Black RGB",
			ShortDescription: "Silent linear switch with a heavy actuation force.",
			LongDescription:  "The Cherry MX2A Silent Black switch offers a quiet, smooth linear typing experience with a heavier actuation force of 60 grams. Designed to minimize noise, this switch features built-in dampeners that significantly reduce the sound of both key press and release, making it ideal for quiet environments such as offices or shared spaces. The MX2A Silent Black is perfect for users who prefer a firmer keystroke while maintaining a low-noise profile, ensuring both performance and discretion in their typing setup.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     4,      // Silent Linear
			ReleaseDate:      parseDate("2024-03-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			SiteURL:          "https://www.cherry.de/en-gb/product/mx2a-silent-black",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.cherry.de/fileadmin/_processed_/2/7/csm_7b82b0a558264e756601bcf93f1fbdb8_7f256d6e4d.jpg",
					AltText:     "Cherry MX2A Silent Black switch",
					OwnerType:   "switches",
				},
				{
					LinkAddress: "https://www.cherry.de/fileadmin/_processed_/d/d/csm_b3b239a6b9956bf1e1aa2ed23c5546c6_30fe10091d.jpg",
					AltText:     "Cherry MX2A Silent Black switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(60)),
				BottomOutForce:        ptr(float32(110)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(53), // Smokey
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "MX2A Brown Tactile",
			ShortDescription: "Tactile switch with a gentle bump for feedback.",
			LongDescription:  "The Cherry MX2A Brown Tactile switch provides a subtle tactile bump, making it suitable for users who prefer feedback without the accompanying click sound. With an actuation force of 55 grams, this switch offers a comfortable balance between typing and gaming, providing just enough feedback to avoid bottoming out while maintaining a smooth and controlled keypress.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2024-03-01"),
			Available:        true,
			PricePoint:       2, // Average
			SiteURL:          "https://www.cherry.de/en-gb/product/mx2a-brown",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.cherry.de/fileadmin/_processed_/6/0/csm_9f59f930928d931aa4bce0f18dc25293_469966436e.jpg",
					AltText:     "Cherry MX2A Brown switch",
					OwnerType:   "switches",
				},
				{
					LinkAddress: "https://www.cherry.de/fileadmin/_processed_/d/6/csm_8c13f97b48e60f9bd848299ed8c1bc2b_bf05fd8194.jpg",
					AltText:     "Cherry MX2A Brown switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(80)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(48), // Brown
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(30), // Medium
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(42), // Rounded
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "MX2A Brown Tactile RGB",
			ShortDescription: "Tactile switch with a gentle bump for feedback.",
			LongDescription:  "The Cherry MX2A Brown Tactile switch provides a subtle tactile bump, making it suitable for users who prefer feedback without the accompanying click sound. With an actuation force of 55 grams, this switch offers a comfortable balance between typing and gaming, providing just enough feedback to avoid bottoming out while maintaining a smooth and controlled keypress.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2024-03-01"),
			Available:        true,
			PricePoint:       2, // Average
			SiteURL:          "https://www.cherry.de/en-gb/product/mx2a-brown",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.cherry.de/fileadmin/_processed_/6/0/csm_9f59f930928d931aa4bce0f18dc25293_469966436e.jpg",
					AltText:     "Cherry MX2A Brown switch",
					OwnerType:   "switches",
				},
				{
					LinkAddress: "https://www.cherry.de/fileadmin/_processed_/d/6/csm_8c13f97b48e60f9bd848299ed8c1bc2b_bf05fd8194.jpg",
					AltText:     "Cherry MX2A Brown switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(80)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(false),
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
			Name:             "MX2A Red Clear",
			ShortDescription: "Smooth linear switch for a balanced typing feel.",
			LongDescription:  "The Cherry MX2A Red Linear switch offers a balanced and smooth typing experience, suitable for both gaming and typing due to its medium actuation force of 45 grams. This switch is designed for users who prefer a light, responsive keystroke with no tactile bump or audible click, allowing for fluid and uninterrupted keypresses.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2024-03-01"),
			Available:        true,
			PricePoint:       2, // Average
			SiteURL:          "https://www.cherry.de/en-gb/product/mx2a-red",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.cherry.de/fileadmin/_processed_/d/7/csm_67119a5593db9c38d993ef1367e8260a_62644f551c.jpg",
					AltText:     "Cherry MX2A Red switch",
					OwnerType:   "switches",
				},
				{
					LinkAddress: "https://www.cherry.de/fileadmin/_processed_/1/a/csm_9161abb21c25569d1d45ef726c97a3bb_bf9c643fa8.jpg",
					AltText:     "Cherry MX2A Red switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(100)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(false),
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
			Name:             "MX2A Red",
			ShortDescription: "Smooth linear switch for a balanced typing feel.",
			LongDescription:  "The Cherry MX2A Red Linear switch offers a balanced and smooth typing experience, suitable for both gaming and typing due to its medium actuation force of 45 grams. This switch is designed for users who prefer a light, responsive keystroke with no tactile bump or audible click, allowing for fluid and uninterrupted keypresses.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2024-03-01"),
			Available:        true,
			PricePoint:       2, // Average
			SiteURL:          "https://www.cherry.de/en-gb/product/mx2a-red",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.cherry.de/fileadmin/_processed_/8/6/csm_3894bf083b6e8e4b775f54706234054b_5f6fc8b606.jpg",
					AltText:     "Cherry MX2A Red switch",
					OwnerType:   "switches",
				},
				{
					LinkAddress: "https://www.cherry.de/fileadmin/_processed_/1/a/csm_9161abb21c25569d1d45ef726c97a3bb_bf9c643fa8.jpg",
					AltText:     "Cherry MX2A Red switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(100)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(45), // Red
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "MX2A Speed Silver Linear",
			ShortDescription: "Fast linear switch for quick actuation.",
			LongDescription:  "The Cherry MX2A Speed Silver Linear switch is designed for rapid keystrokes, making it ideal for gaming and other high-speed typing scenarios. With a short 1.2mm actuation distance and 45 grams of force, this switch enables quick responses, helping to improve performance in fast-paced environments. The MX2A Speed Silver is a top choice for competitive gamers seeking the fastest actuation possible.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2024-03-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			SiteURL:          "https://www.cherry.de/en-gb/product/mx2a-speed-silver",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.cherry.de/fileadmin/_processed_/0/6/csm_5d98f7ba6a035459017c6c08503cee27_1e8bebf097.jpg",
					AltText:     "Cherry MX2A Speed Silver switch",
					OwnerType:   "switches",
				},
				{
					LinkAddress: "https://www.cherry.de/fileadmin/_processed_/6/5/csm_af984c43257e2e95b048f211647c6ba7_c84a2e36f9.jpg",
					AltText:     "Cherry MX2A Speed Silver switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(1.2)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(100)),
				BottomOutForcePoint:   ptr(float32(3.4)),
				TotalTravel:           ptr(float32(3.4)),
				PreTravel:             ptr(float32(1.2)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(59), // Silver
				TopHousingColorID:     ptr(49), // Black
				BottomHousingColorID:  ptr(49), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(37), // Clacky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "MX2A Speed Silver Linear RGB",
			ShortDescription: "Fast linear switch for quick actuation.",
			LongDescription:  "The Cherry MX2A Speed Silver Linear switch is designed for rapid keystrokes, making it ideal for gaming and other high-speed typing scenarios. With a short 1.2mm actuation distance and 45 grams of force, this switch enables quick responses, helping to improve performance in fast-paced environments. The MX2A Speed Silver is a top choice for competitive gamers seeking the fastest actuation possible.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2024-03-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			SiteURL:          "https://www.cherry.de/en-gb/product/mx2a-speed-silver",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.cherry.de/fileadmin/_processed_/0/0/csm_c11c911aa334e2f6d2cfcc2412edaf4f_2d82033979.jpg",
					AltText:     "Cherry MX2A Speed Silver switch",
					OwnerType:   "switches",
				},
				{
					LinkAddress: "https://www.cherry.de/fileadmin/_processed_/0/6/csm_5d98f7ba6a035459017c6c08503cee27_1e8bebf097.jpg",
					AltText:     "Cherry MX2A Speed Silver switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(1.2)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(100)),
				BottomOutForcePoint:   ptr(float32(3.4)),
				TotalTravel:           ptr(float32(3.4)),
				PreTravel:             ptr(float32(1.2)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(59), // Silver
				TopHousingColorID:     ptr(46), // Clear
				BottomHousingColorID:  ptr(46), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(37), // Clacky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "MX2A Black Linear",
			ShortDescription: "Linear switch with a heavy actuation force.",
			LongDescription:  "The Cherry MX2A Black Linear switch offers a smooth keystroke with a heavier actuation force of 60 grams. This makes it ideal for users who prefer a firm, controlled typing experience. The switch is well-suited for both heavy typists and gamers who require a more deliberate keypress. With its linear feel and no tactile bump, the MX2A Black ensures a consistent and smooth actuation with every keystroke.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2024-03-01"),
			Available:        true,
			PricePoint:       2, // Average
			SiteURL:          "https://www.cherry.de/en-gb/product/mx2a-black",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.cherry.de/fileadmin/_processed_/9/b/csm_3564453e094e91c909de04c5c92b222f_9ea01c2323.jpg",
					AltText:     "Cherry MX2A Black switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(60)),
				BottomOutForce:        ptr(float32(80)),
				BottomOutForcePoint:   ptr(float32(4.0)),
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
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		// {
		// 	Name:             "MX2A Clear",
		// 	ShortDescription: "Tactile switch with a high actuation force.",
		// 	LongDescription:  "The Cherry MX2A Clear switch is designed for users who prefer a strong tactile feedback without the click sound. With an actuation force of 65 grams, this switch offers a pronounced tactile bump that is ideal for heavy typists who enjoy a firm and responsive keystroke. The MX2A Clear is particularly suited for those who need a tactile switch that provides clear feedback for every keypress, making it a great choice for both typing and precise input tasks.",
		// 	ManufacturerID:   ptr(1), // Cherry
		// 	BrandID:          ptr(1), // Cherry
		// 	SwitchTypeID:     2,      // Tactile
		// 	ReleaseDate:      parseDate("2024-03-01"),
		// 	Available:        true,
		// 	PricePoint:       3, // Expensive
		// 	SiteURL:          "https://www.cherry-world.com/mx-clear",
		// 	ImageLinks: []models.ImageLink{
		// 		{
		// 			LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-c1na_c1nn_image.jpg",
		// 			AltText:     "Cherry MX2A Clear switch",
		// 			OwnerType:   "switches",
		// 		},
		// 		{
		// 			LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-c1nn_image.jpg",
		// 			AltText:     "Cherry MX2A Clear switch",
		// 			OwnerType:   "switches",
		// 		},
		// 	},
		// 	Details: models.Details{
		// 		SpringForce:           ptr(float32(65)),
		// 		ActuationPoint:        ptr(float32(2.0)),
		// 		ActuationForce:        ptr(float32(65)),
		// 		BottomOutForce:        ptr(float32(80)),
		// 		BottomOutForcePoint:   ptr(float32(4.0)),
		// 		TotalTravel:           ptr(float32(4.0)),
		// 		PreTravel:             ptr(float32(2.0)),
		// 		FactoryLubed:          ptr(false),
		// 		StemMaterialID:        ptr(23), // POM
		// 		TopHousingMaterialID:  ptr(26), // Polycarbonate
		// 		BaseHousingMaterialID: ptr(24), // Nylon
		// 		SpringMaterialTypeID:  ptr(17), // Steel
		// 		StemColorID:           ptr(49), // Clear
		// 		TopHousingColorID:     ptr(49), // Clear
		// 		BottomHousingColorID:  ptr(49), // Clear
		// 		PinConfigurationID:    ptr(74), // 3-Pin
		// 		SoundTypeID:           ptr(34), // Thocky
		// 		SoundLevelID:          ptr(30), // Medium
		// 		TactilityTypeID:       ptr(38), // Leaf Spring
		// 		TactilityFeedbackID:   ptr(41), // Sharp
		// 		HasShineThrough:       ptr(true),
		// 	},
		// },
		{
			Name:             "MX2A Grey",
			ShortDescription: "Tactile switch with an even higher actuation force.",
			LongDescription:  "The Cherry MX2A Grey switch is designed for users who prefer a very strong tactile feedback with an actuation force of 80 grams. This switch offers a firm tactile bump, providing clear and pronounced feedback with every keystroke. The MX2A Grey is perfect for users who require a high level of resistance in their switches, making it ideal for precise and deliberate key presses in both typing and heavy-duty applications.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2024-03-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			SiteURL:          "https://www.cherry.de/en-gb/product/mx-grey-tactile",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.cherry.de/fileadmin/_processed_/c/a/csm_29255bb460c39a5bf72d622257d52095_dbcf419cc0.jpg",
					AltText:     "Cherry MX2A Grey switch",
					OwnerType:   "switches",
				},
				{
					LinkAddress: "https://www.cherry.de/fileadmin/_processed_/9/0/csm_997f810ecfa73be1e76878acddd37a35_8bbe0e2943.jpg",
					AltText:     "Cherry MX2A Grey switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(80)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(70)),
				BottomOutForce:        ptr(float32(95)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(58), // Gray
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(31), // High
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(41), // Sharp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "MX2A Grey RBG",
			ShortDescription: "Tactile switch with an even higher actuation force.",
			LongDescription:  "The Cherry MX2A Grey switch is designed for users who prefer a very strong tactile feedback with an actuation force of 80 grams. This switch offers a firm tactile bump, providing clear and pronounced feedback with every keystroke. The MX2A Grey is perfect for users who require a high level of resistance in their switches, making it ideal for precise and deliberate key presses in both typing and heavy-duty applications.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2024-03-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			SiteURL:          "https://www.cherry.de/en-gb/product/mx-grey-tactile",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.cherry.de/fileadmin/_processed_/9/0/csm_997f810ecfa73be1e76878acddd37a35_8bbe0e2943.jpg",
					AltText:     "Cherry MX2A Grey switch",
					OwnerType:   "switches",
				},
				{
					LinkAddress: "https://www.cherry.de/fileadmin/_processed_/a/b/csm_dbab30e2e4d1e1a651bd597c90e22a31_3e671cbe0f.jpg",
					AltText:     "Cherry MX2A Grey switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(70)),
				ActuationPoint:        ptr(float32(2.2)),
				ActuationForce:        ptr(float32(70)),
				BottomOutForce:        ptr(float32(95)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(58), // Gray
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(31), // High
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(41), // Sharp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "MX Purple",
			ShortDescription: "Tactile switch with medium actuation force and a sharp tactile bump.",
			LongDescription:  "The Cherry MX Purple is a tactile switch that offers a distinct typing experience with its medium actuation force and pronounced tactile feedback. With an actuation force of 45 grams and a sharp tactile bump, this switch provides clear feedback during keystrokes while maintaining a comfortable typing experience. The switch features a unique purple stem and is designed for users who prefer definitive tactile feedback without the noise of clicky switches. Its carefully tuned spring and tactile leaf mechanism create a balanced feel that's suitable for both typing and gaming applications.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2024-03-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			SiteURL:          "https://www.cherry.de/en-gb/product/mx-purple",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.cherry.de/fileadmin/_processed_/a/c/csm_df43d3d6cd48241313c771062130929a_1ea5387229.jpg",
					AltText:     "Cherry MX Purple switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				InitialForce:          ptr(float32(30)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(100)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(55), // Purple
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(41), // Sharp
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "MX Purple RGB",
			ShortDescription: "RGB-compatible tactile switch with translucent housing and sharp tactile feedback.",
			LongDescription:  "The Cherry MX Purple RGB is the illumination-optimized version of the standard MX Purple switch. It maintains the same tactile characteristics with a 45g actuation force and sharp tactile feedback, while featuring a fully transparent housing designed for optimal RGB lighting effects. The switch combines the satisfying tactile feel of the original Purple with enhanced lighting capabilities, making it perfect for users who want both a responsive typing experience and vibrant keyboard aesthetics. The transparent housing allows for uniform light distribution while maintaining the switch's distinctive tactile characteristics.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2024-03-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			SiteURL:          "https://www.cherry.de/en-gb/product/mx-purple",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.cherry.de/fileadmin/_processed_/a/c/csm_df43d3d6cd48241313c771062130929a_1ea5387229.jpg",
					AltText:     "Cherry MX Purple RGB switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				InitialForce:          ptr(float32(30)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(100)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(55), // Purple
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(41), // Sharp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "MX2A Blue Clicky",
			ShortDescription: "Clicky switch with an audible click for tactile feedback.",
			LongDescription:  "The Cherry MX2A Blue Clicky switch delivers the satisfying click sound and tactile feedback that typists love. With an actuation force of 50 grams, this switch provides a crisp click and a noticeable bump, ensuring each keystroke is both audible and precise. Ideal for those who enjoy the classic typing experience, the MX2A Blue is perfect for heavy typists or those who simply love the sound of their keyboard.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2024-03-01"),
			Available:        true,
			PricePoint:       2, // Average
			SiteURL:          "https://www.cherry.de/en-gb/product/mx2a-blue",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.cherry.de/fileadmin/_processed_/5/2/csm_b8dc73682ef3a1f673e846aa767c2884_de6d3192a7.jpg",
					AltText:     "Cherry MX2A switch",
					OwnerType:   "switches",
				},
				{
					LinkAddress: "https://www.cherry.de/fileadmin/_processed_/3/e/csm_44ce55f4cfd9d0b8fe831e51a6b1f055_489a99cb9c.jpg",
					AltText:     "Cherry MX2A switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.2)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(60)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.2)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(41), // Sharp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "MX2A Blue Clicky RBG",
			ShortDescription: "Clicky switch with an audible click for tactile feedback.",
			LongDescription:  "The Cherry MX2A Blue Clicky switch delivers the satisfying click sound and tactile feedback that typists love. With an actuation force of 50 grams, this switch provides a crisp click and a noticeable bump, ensuring each keystroke is both audible and precise. Ideal for those who enjoy the classic typing experience, the MX2A Blue is perfect for heavy typists or those who simply love the sound of their keyboard.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2024-03-01"),
			Available:        true,
			PricePoint:       2, // Average
			SiteURL:          "https://www.cherry.de/en-gb/product/mx2a-blue",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.cherry.de/fileadmin/_processed_/6/e/csm_6a03867208aa7bca0d4c3677f6fd1afc_148a750076.jpg",
					AltText:     "Cherry MX2A switch",
					OwnerType:   "switches",
				},
				{
					LinkAddress: "https://www.cherry.de/fileadmin/_processed_/3/e/csm_44ce55f4cfd9d0b8fe831e51a6b1f055_489a99cb9c.jpg",
					AltText:     "Cherry MX2A switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.2)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(60)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.2)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(41), // Sharp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "MX Blue",
			ShortDescription: "Clicky switch with a tactile bump.",
			LongDescription:  "Cherry MX Blue switches are iconic for their audible click and tactile feedback, offering a distinct typing experience that is both responsive and satisfying. With an actuation force of 50 grams, these switches provide a noticeable bump before actuation, coupled with a sharp click sound that appeals to typists who appreciate clear auditory and physical feedback. While they are less common in gaming due to their pronounced click, MX Blues are a top choice for those who value precision and clarity in their keystrokes.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("1987-01-01"),
			Available:        false,
			PricePoint:       2,
			Details: models.Details{
				SpringForce:           ptr(float32(50)),
				ActuationPoint:        ptr(float32(2.2)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(60)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.2)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(41), // Sharp
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "MX Brown",
			ShortDescription: "Tactile switch with a moderate actuation force.",
			LongDescription:  "Cherry MX Brown switches strike a balance between typing and gaming with their tactile bump and moderate actuation force of 55 grams. These switches are quieter than their Blue counterparts, lacking the audible click but still providing tactile feedback that is gentle yet noticeable. This makes them versatile and well-suited for both work and play, appealing to users who want a middle ground between the stark tactility of MX Blue switches and the smoothness of MX Red switches.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("1994-01-01"),
			Available:        false,
			PricePoint:       2,
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(65)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(48), // Brown
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(30), // Medium
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(42), // Rounded
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "MX Black",
			ShortDescription: "Linear switch with a heavier actuation force.",
			LongDescription:  "Cherry MX Black switches offer a smooth linear actuation similar to MX Red switches but with a heavier actuation force of 60 grams. This increased resistance makes them ideal for users who prefer a more deliberate keystroke, reducing the chance of accidental presses. MX Black switches are often favored by those who type heavily or for gaming environments where controlled inputs are critical. The absence of tactile bumps or clicks also ensures a quieter, more focused typing experience.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("1984-01-01"),
			Available:        false,
			PricePoint:       2,
			SiteURL:          "https://www.cherry-world.com/mx-black.html",
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(60)),
				BottomOutForce:        ptr(float32(80)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "MX Clear",
			ShortDescription: "Tactile switch with a high actuation force.",
			LongDescription:  "Cherry MX Clear switches are designed for those who prefer a heavier tactile bump in their typing experience. With an actuation force of 65 grams, these switches provide a noticeable resistance, making each keystroke deliberate and satisfying. MX Clear switches are often chosen by users who want a tactile switch that offers more feedback than MX Browns but without the clicky noise associated with MX Blue switches.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("1994-01-01"),
			Available:        false,
			PricePoint:       3,
			SiteURL:          "https://www.cherry.de/en-gb/product/mx-clear",
			ImageLinks: []models.ImageLink{
				{
					LinkAddress: "https://www.cherry.de/fileadmin/_processed_/f/0/csm_2cd530be38bda55e7660a06d9ace615d_08dce11b60.jpg",
					AltText:     "Cherry MX Clear switch",
					OwnerType:   "switches",
				},
			},
			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(65)),
				BottomOutForce:        ptr(float32(95)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(49), // Clear
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(41), // Sharp
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "MX Grey",
			ShortDescription: "Tactile switch with a heavy actuation force.",
			LongDescription:  "Cherry MX Grey switches are an evolution of the MX Clear, featuring an even higher actuation force of 80 grams. This provides a pronounced tactile bump and substantial resistance, making them suitable for users who prefer a firm and assertive keystroke. These switches are less common in mass-market keyboards and are often found in specialized or enthusiast builds that prioritize a heavier, more tactile typing experience.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("1994-01-01"),
			Available:        false,
			PricePoint:       3,
			Details: models.Details{
				SpringForce:           ptr(float32(80)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(80)),
				BottomOutForce:        ptr(float32(105)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(58), // Gray
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(31), // High
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(41), // Sharp
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "MX Speed Silver",
			ShortDescription: "Linear switch with a shorter travel distance for faster actuation.",
			LongDescription:  "Cherry MX Speed Silver switches are optimized for rapid keystrokes with a shorter actuation distance and total travel. These linear switches require just 1.2mm to actuate and only 45 grams of force, making them perfect for competitive gaming where speed is crucial. The reduced travel distance allows for faster key presses, helping to shave milliseconds off each action in fast-paced environments. They are a go-to choice for gamers seeking performance and responsiveness.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2016-01-01"),
			Available:        false,
			PricePoint:       3,
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(1.2)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(75)),
				BottomOutForcePoint:   ptr(float32(3.4)),
				TotalTravel:           ptr(float32(3.4)),
				PreTravel:             ptr(float32(1.2)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(59), // Silver
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "MX Nature White",
			ShortDescription: "Linear switch with a medium actuation force, between Red and Black.",
			LongDescription:  "Cherry MX Nature White switches are designed to offer a balanced typing experience, with an actuation force that sits comfortably between the lighter MX Red and the heavier MX Black switches. With a smooth, linear action and a medium force of 55 grams, these switches provide a well-rounded keystroke that is ideal for users who want a linear switch that isn't too light or too heavy. They are often chosen by those who prefer a slightly more resistant keystroke without sacrificing speed or smoothness.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2016-01-01"),
			Available:        false,
			PricePoint:       2,
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(75)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(51), // White
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "MX Violet",
			ShortDescription: "Tactile switch with a light actuation force, similar to Brown but lighter.",
			LongDescription:  "Cherry MX Violet switches offer a tactile bump similar to MX Browns but with a lighter actuation force, providing a gentle yet satisfying typing experience. Requiring only 45 grams of force to actuate, these switches are ideal for users who appreciate tactile feedback but prefer a softer, less resistant keypress. They are well-suited for both typing and gaming, offering a balance of feedback and ease of use.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2020-01-01"),
			Available:        false,
			PricePoint:       2,
			SiteURL:          "",
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(55), // Purple
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(42), // Rounded
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "MX White",
			ShortDescription: "A softer clicky switch, sometimes called 'Milk'.",
			LongDescription:  "Cherry MX White switches, often referred to as 'Milk' switches due to their smoother click, are a quieter alternative to the more common MX Blue switches. They retain the clicky feedback but with a softer, less pronounced sound, making them a great choice for users who enjoy a clicky switch but want something less intrusive in terms of noise. These switches are less common and have become a favorite among those who want a unique auditory and tactile experience.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2011-01-01"),
			Available:        false,
			PricePoint:       2,
			SiteURL:          "",
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.2)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(70)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.2)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(51), // White
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(30), // Medium
				TactilityTypeID:       ptr(39), // Coil Spring
				TactilityFeedbackID:   ptr(42), // Rounded
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "MX Tactile Grey",
			ShortDescription: "A variant of the Grey switch with a slightly different force curve.",
			LongDescription:  "Cherry MX Tactile Grey switches are a specialized variant of the standard MX Grey switches, offering a unique force curve that provides a slightly different tactile experience. These switches are designed for users who appreciate a heavy, pronounced tactile bump and prefer a switch with a firm and deliberate keystroke. They are often used in niche keyboard builds where a distinct tactile feel is desired.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("1994-01-01"),
			Available:        false,
			PricePoint:       3,
			SiteURL:          "",
			Details: models.Details{
				SpringForce:           ptr(float32(80)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(80)),
				BottomOutForce:        ptr(float32(110)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(58), // Gray
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(31), // High
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(41), // Sharp
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "MX Super Black",
			ShortDescription: "An extremely heavy linear switch.",
			LongDescription:  "Cherry MX Super Black switches are designed for users who demand an extremely heavy actuation force. With a force requirement of 150 grams, these linear switches are among the stiffest in the Cherry MX lineup, offering unparalleled resistance to accidental key presses. Super Black switches are typically used in specialized keyboards or applications where a very deliberate keypress is necessary, providing a unique typing experience that is both demanding and rewarding.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("1984-01-01"),
			Available:        false,
			PricePoint:       3,
			SiteURL:          "",
			Details: models.Details{
				SpringForce:           ptr(float32(150)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(150)),
				BottomOutForce:        ptr(float32(180)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(22), // ABS
				BaseHousingMaterialID: ptr(22), // ABS
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(74), // 3-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(32), // Very High
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "MX Red",
			ShortDescription: "Linear switch with a light actuation force.",
			LongDescription:  "Cherry MX Red switches are renowned for their smooth, linear actuation and light operating force, making them a preferred choice for gamers and typists who enjoy a fast, uninterrupted keystroke. These switches require only 45 grams of force to actuate, ensuring a quick response without any tactile bump or audible click, allowing for a fluid typing experience. They are especially favored in gaming for their consistent key presses and rapid input, reducing finger fatigue during extended use.",
			ManufacturerID:   ptr(1), // Cherry
			BrandID:          ptr(1), // Cherry
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("1984-01-01"),
			Available:        false,
			PricePoint:       2,
			SiteURL:          "https://www.cherry-world.com/mx-red.html",
			Details: models.Details{
				SpringForce:           ptr(float32(45)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(60)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
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

	err := processSwitches(tx, cherry, admin)
	if err != nil {
		return err
	}

	return nil
}
