package seed

import (
	"switches/models"

	"gorm.io/gorm"
)

func seedZealPC(tx *gorm.DB, admin models.User) error {
	zealpcSwitches := []models.Switch{
		{
			Name:             "Zealios V2",
			ShortDescription: "Tactile switch with a rounded bump and smooth action.",
			LongDescription:  "Zealios V2 switches offer a distinct tactile bump that is both smooth and satisfying. Designed for typists who prefer a clear but non-disruptive tactile feedback, these switches deliver precision and comfort with every keystroke, making them a top choice for enthusiasts seeking a refined typing experience.",
			ManufacturerID:   ptr(5), // ZealPC
			BrandID:          ptr(5), // ZealPC
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2018-11-15"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce: ptr(
					float32(65),
				), // Common weight, but Zealios V2 comes in various spring weights
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(55), // Purple
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(42), // Rounded
			},
		},
		{
			Name:             "Tealios V2",
			ShortDescription: "Linear switch with a unique teal color and smooth feel.",
			LongDescription:  "Tealios V2 switches are celebrated for their buttery-smooth keystrokes, making them a favorite among both gamers and typists. The combination of a unique teal housing and high-quality materials ensures a frictionless, consistent typing experience with fast actuation, perfect for high-performance applications.",
			ManufacturerID:   ptr(5), // ZealPC
			BrandID:          ptr(5), // ZealPC
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2018-09-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(67)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(62), // Teal
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(62), // Teal
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Healios",
			ShortDescription: "Silent linear switch for quiet typing.",
			LongDescription:  "Healios switches are designed for near-silent operation, making them ideal for quiet work environments or late-night typing. These switches feature a dampened linear action that eliminates noise without compromising on the smoothness and responsiveness that ZealPC is known for.",
			ManufacturerID:   ptr(5), // ZealPC
			BrandID:          ptr(5), // ZealPC
			SwitchTypeID:     4,      // Silent Linear
			ReleaseDate:      parseDate("2018-07-20"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(67)),
				BottomOutForcePoint:   ptr(float32(3.7)), // Slightly reduced due to silencing pads
				TotalTravel:           ptr(float32(3.7)), // Slightly reduced due to silencing pads
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
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
			Name:             "Zilents V2",
			ShortDescription: "Silent tactile switch with a smooth bump.",
			LongDescription:  "Zilents V2 switches provide a silent yet tactile typing experience, combining a smooth, rounded bump with noise-dampening features. Ideal for office use or shared spaces, these switches deliver the satisfying tactile feedback of Zealios with minimal sound.",
			ManufacturerID:   ptr(5), // ZealPC
			BrandID:          ptr(5), // ZealPC
			SwitchTypeID:     5,      // Silent Tactile
			ReleaseDate:      parseDate("2019-02-10"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce: ptr(
					float32(65),
				), // Common weight, but Zilents V2 comes in various spring weights
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				BottomOutForcePoint:   ptr(float32(3.7)), // Slightly reduced due to silencing pads
				TotalTravel:           ptr(float32(3.7)), // Slightly reduced due to silencing pads
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(63), // Lavender
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(42), // Rounded
			},
		},
		{
			Name:             "Rosélios",
			ShortDescription: "Smooth linear switch with a unique color.",
			LongDescription:  "Rosélios switches stand out for their smooth linear action and elegant rose-colored housing. Perfect for those who value both aesthetics and performance, these switches provide a luxurious typing experience with minimal resistance, making them a popular choice for custom builds.",
			ManufacturerID:   ptr(5), // ZealPC
			BrandID:          ptr(5), // ZealPC
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2020-05-18"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(67)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(57), // Pink
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(57), // Pink
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Zeal Clickiez",
			ShortDescription: "Unique clicky switch with customizable click sound.",
			LongDescription:  "Zeal Clickiez switches offer a customizable clicky experience, allowing users to adjust the click element to suit their preferences. These switches are perfect for those who enjoy a tactile clicky feedback, with the flexibility to tailor the sound and feel to their liking.",
			ManufacturerID:   ptr(5), // ZealPC
			BrandID:          ptr(5), // ZealPC
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2021-06-25"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(67)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				HasShineThrough:       ptr(true),
				TactilityTypeID:       ptr(40), // Click Bar
				TactilityFeedbackID:   ptr(43), // Crisp
			},
		},
		{
			Name:             "Zealios Clicky V2",
			ShortDescription: "Clicky switch with smooth operation and crisp sound.",
			LongDescription:  "Zealios Clicky V2 switches are engineered for those who love a satisfying click with every keystroke. With a crisp sound and smooth action, these switches provide a delightful clicky experience while maintaining ZealPC's renowned build quality.",
			ManufacturerID:   ptr(5), // ZealPC as manufacturer
			BrandID:          ptr(5), // ZealPC as brand
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2019-11-20"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(67)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				HasShineThrough:       ptr(true),
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(43), // Crisp
			},
		},
		{
			Name:             "Zealios V1",
			ShortDescription: "Original tactile switch with a smooth and satisfying bump.",
			LongDescription:  "Zealios V1 switches were the original tactile offering from ZealPC, providing a smooth tactile bump that is both satisfying and responsive. These switches are perfect for typists who appreciate a classic tactile feel, with a balance of smoothness and feedback. Although discontinued, they remain a favorite among keyboard enthusiasts.",
			ManufacturerID:   ptr(5), // ZealPC as manufacturer
			BrandID:          ptr(5), // ZealPC as brand
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2015-04-15"),
			Available:        false, // Discontinued
			PricePoint:       3,     // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(55), // Purple
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(42), // Rounded
			},
		},
		{
			Name:             "Zealios V2.2",
			ShortDescription: "Refined tactile switch with improved bump and smoother operation.",
			LongDescription:  "The Zealios V2.2 is an enhanced version of the popular Zealios V2, featuring a more pronounced tactile bump and an even smoother keystroke. This switch is ideal for users who demand precise tactile feedback and a high-quality typing experience, making it a top choice for both typing and gaming.",
			ManufacturerID:   ptr(5), // ZealPC as manufacturer
			BrandID:          ptr(5), // ZealPC as brand
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2022-08-05"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(55), // Purple
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(41), // Sharp
			},
		},
		{
			Name:             "Zilent V2",
			ShortDescription: "Silent tactile switch with a soft and quiet bump.",
			LongDescription:  "Zilent V2 switches are designed to offer a quiet yet tactile typing experience. With a soft bump and noise-dampening technology, these switches are perfect for those who require both tactile feedback and silent operation, making them ideal for office environments and shared spaces.",
			ManufacturerID:   ptr(5), // ZealPC as manufacturer
			BrandID:          ptr(5), // ZealPC as brand
			SwitchTypeID:     5,      // Silent Tactile
			ReleaseDate:      parseDate("2019-02-10"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				BottomOutForcePoint:   ptr(float32(3.7)),
				TotalTravel:           ptr(float32(3.7)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(63), // Lavender
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(42), // Rounded
			},
		},
		{
			Name:             "Zilent V2.2",
			ShortDescription: "Enhanced silent tactile switch with smoother and quieter operation.",
			LongDescription:  "Zilent V2.2 switches offer an upgraded silent tactile experience, with a softer and more refined bump than previous versions. Designed for users who need tactile feedback without the noise, these switches are perfect for quiet environments where precision and silence are essential.",
			ManufacturerID:   ptr(5), // ZealPC as manufacturer
			BrandID:          ptr(5), // ZealPC as brand
			SwitchTypeID:     5,      // Silent Tactile
			ReleaseDate:      parseDate("2022-12-10"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				BottomOutForcePoint:   ptr(float32(3.7)),
				TotalTravel:           ptr(float32(3.7)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(63), // Lavender
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(42), // Rounded
			},
		},
		{
			Name:             "Tealios V1",
			ShortDescription: "Smooth linear switch with a teal housing.",
			LongDescription:  "Tealios V1 switches were the original linear switch offering from ZealPC, known for their ultra-smooth operation and distinctive teal housing. These switches provide a fast and effortless typing experience, making them popular among both gamers and typists. Although discontinued, they are still highly regarded in the mechanical keyboard community.",
			ManufacturerID:   ptr(5), // ZealPC as manufacturer
			BrandID:          ptr(5), // ZealPC as brand
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2016-10-01"),
			Available:        false, // Discontinued
			PricePoint:       3,     // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(67)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(62), // Teal
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(62), // Teal
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Rosélios V2",
			ShortDescription: "Upgraded linear switch with a unique aesthetic.",
			LongDescription:  "Rosélios V2 switches are the upgraded version of the original Rosélios, offering enhanced smoothness and a striking rose-colored housing. These switches are designed for users who value a premium linear experience with a touch of elegance, making them a standout choice for custom keyboard builds.",
			ManufacturerID:   ptr(5), // ZealPC as manufacturer
			BrandID:          ptr(5), // ZealPC as brand
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-03-15"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(67)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(57), // Pink
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(57), // Pink
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Zealios V2.1 (Linear)",
			ShortDescription: "Linear variant of the popular Zealios V2 switch.",
			LongDescription:  "Zealios V2.1 Linear switches deliver the same high-quality build as the tactile version but with a smooth linear action. These switches are ideal for users who prefer a linear feel in their typing or gaming experience, offering a fast and responsive keystroke with minimal resistance.",
			ManufacturerID:   ptr(5), // ZealPC as manufacturer
			BrandID:          ptr(5), // ZealPC as brand
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2020-05-18"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(67)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
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
			Name:             "Silent Zealios V1",
			ShortDescription: "Original silent tactile switch with smooth feedback.",
			LongDescription:  "Silent Zealios V1 switches were the first silent tactile switches from ZealPC, offering a smooth feedback with noise-dampening technology. Designed for quiet typing environments, these switches provide a soft tactile bump without the accompanying sound, making them perfect for offices and shared spaces. Although discontinued, they remain a beloved choice for silent typing enthusiasts.",
			ManufacturerID:   ptr(5), // ZealPC as manufacturer
			BrandID:          ptr(5), // ZealPC as brand
			SwitchTypeID:     5,      // Silent Tactile
			ReleaseDate:      parseDate("2017-04-15"),
			Available:        false, // Discontinued
			PricePoint:       3,     // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				BottomOutForcePoint:   ptr(float32(3.7)),
				TotalTravel:           ptr(float32(3.7)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(55), // Purple
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(42), // Rounded
			},
		},
		{
			Name:             "Zilent V1",
			ShortDescription: "Quiet tactile switch with a soft and gentle bump.",
			LongDescription:  "Zilent V1 switches are known for their quiet operation and soft tactile bump, making them suitable for noise-sensitive environments. These switches offer a gentle tactile feel without the noise, perfect for offices and shared spaces where quiet typing is essential. Though discontinued, they remain a popular choice among those who prioritize silence in their typing experience.",
			ManufacturerID:   ptr(5), // ZealPC as manufacturer
			BrandID:          ptr(5), // ZealPC as brand
			SwitchTypeID:     5,      // Silent Tactile
			ReleaseDate:      parseDate("2018-02-10"),
			Available:        false, // Discontinued
			PricePoint:       3,     // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				BottomOutForcePoint:   ptr(float32(3.7)),
				TotalTravel:           ptr(float32(3.7)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(63), // Lavender
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(42), // Rounded
			},
		},
		{
			Name:             "Sakurios",
			ShortDescription: "Limited edition linear switch with a sakura pink color.",
			LongDescription:  "Sakurios switches are a special edition linear switch featuring a sakura pink color. Designed for a smooth and quiet typing experience, these switches are popular among custom keyboard builders and enthusiasts who value both aesthetics and performance. Their limited availability makes them a sought-after choice for unique keyboard builds.",
			ManufacturerID:   ptr(5), // ZealPC as manufacturer
			BrandID:          ptr(5), // ZealPC as brand
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2020-02-14"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(62)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(62)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(57), // Pink
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(57), // Pink
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Rosélios (Special Edition)",
			ShortDescription: "Special edition linear switch with unique aesthetics.",
			LongDescription:  "The Rosélios Special Edition switches offer a smooth linear experience with a unique rose-colored housing. Designed for users who want a premium feel and appearance, these limited edition switches are perfect for custom builds and enthusiasts looking for something distinctive.",
			ManufacturerID:   ptr(5), // ZealPC as manufacturer
			BrandID:          ptr(5), // ZealPC as brand
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-05-18"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(67)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(57), // Pink
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(57), // Pink
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Turquoise Tealios",
			ShortDescription: "Smooth linear switch with a unique turquoise color.",
			LongDescription:  "Turquoise Tealios are a variant of the popular Tealios linear switch, featuring a distinctive turquoise color. They offer the same ultra-smooth keystrokes and high performance as the original Tealios, with a fresh and vibrant aesthetic that stands out in custom keyboard builds.",
			ManufacturerID:   ptr(5), // ZealPC
			BrandID:          ptr(5), // ZealPC
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2020-09-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(67)),
				BottomOutForcePoint:   ptr(float32(4.0)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(61), // Turquoise
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(61), // Turquoise
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Aqua Zilents",
			ShortDescription: "Silent tactile switch with a unique aqua color.",
			LongDescription:  "Aqua Zilents are a color variant of the Zilent silent tactile switches. They provide the same quiet, tactile experience as regular Zilents but with a distinctive aqua-colored housing. These switches are perfect for those who desire a silent tactile feel with a unique visual appeal.",
			ManufacturerID:   ptr(5), // ZealPC
			BrandID:          ptr(5), // ZealPC
			SwitchTypeID:     5,      // Silent Tactile
			ReleaseDate:      parseDate("2020-09-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				BottomOutForcePoint:   ptr(float32(3.7)),
				TotalTravel:           ptr(float32(3.7)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          true,
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(18), // Gold-Plated Steel
				StemColorID:           ptr(61), // Turquoise
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(61), // Turquoise
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(42), // Rounded
			},
		},
	}

	err := processSwitches(tx, zealpcSwitches, admin)
	if err != nil {
		return err
	}

	return nil
}
