package seed

import (
	"switches/models"

	"gorm.io/gorm"
)

func seedDurock(tx *gorm.DB, admin models.User) error {
	durock := []models.Switch{
		{
			Name:             "T1",
			ShortDescription: "Tactile switch with a sharp bump and POM stem.",
			LongDescription:  "The Durock T1 is a highly regarded tactile switch, often praised for its sharp and distinct tactile bump that provides clear feedback with each keystroke. The POM stem ensures smooth operation and durability, while the high-quality materials minimize wobble, making it a popular choice among enthusiasts who value stability and precision in their typing experience. Its satisfying bump and overall performance make it a go-to option for both gaming and typing.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2019-04-01"),
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
				TopHousingColorID:     ptr(51), // White
				BottomHousingColorID:  ptr(51), // White
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(41), // Sharp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "L7",
			ShortDescription: "Linear switch with smooth operation and low noise.",
			LongDescription:  "The Durock L7 is a linear switch that exemplifies the smooth, consistent keystroke sought after by linear switch enthusiasts. It features minimal spring noise, ensuring a quiet typing experience without compromising on performance. Known for its smoothness and low noise levels, the L7 is ideal for users who prefer a distraction-free environment while typing or gaming. The refined design and exceptional consistency make it a standout choice for those who demand reliability.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-06-01"),
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
			Name:             "Koala",
			ShortDescription: "Tactile switch inspired by Holy Panda design.",
			LongDescription:  "The Durock Koala is a tactile switch inspired by the iconic Holy Panda design, featuring a medium to heavy tactile bump that delivers a smooth and rounded feel. This switch is known for its consistent tactile feedback and is highly favored among enthusiasts for its premium typing experience. The Koala's high-quality construction and smooth actuation make it a top choice for those who appreciate tactile switches with a well-defined bump.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2020-03-01"),
			Available:        true,
			PricePoint:       3, // Expensive
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
				TopHousingColorID:     ptr(58), // Gray
				BottomHousingColorID:  ptr(58), // Gray
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(42), // Rounded
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Sunflower",
			ShortDescription: "Linear switch with a bright yellow stem and smooth feel.",
			LongDescription:  "The Durock Sunflower is a linear switch that stands out with its vibrant yellow stem and incredibly smooth operation. It offers a quiet and fluid keystroke, making it an excellent choice for users who appreciate both aesthetics and performance. The Sunflower's bright color adds a touch of personality to any keyboard build, while its smooth action ensures a comfortable and enjoyable typing experience.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-09-01"),
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
			Name:             "Shrimp",
			ShortDescription: "Tactile switch with a unique pink housing and sharp bump.",
			LongDescription:  "The Durock Shrimp is a tactile switch that combines a distinctive pink housing with a sharp, precise tactile bump. This switch is a favorite among those who appreciate both style and substance, offering a delightful typing experience with a responsive tactile feel. The Shrimp's unique design and excellent performance make it a standout choice for users who want a tactile switch that delivers on both aesthetics and functionality.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2022-05-01"),
			Available:        true,
			PricePoint:       3, // Expensive
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
			Name:             "Dolphin",
			ShortDescription: "Linear switch with a smooth, consistent feel.",
			LongDescription:  "The Durock Dolphin is a linear switch that offers a smooth and consistent typing experience, characterized by its deep blue stem and exceptional build quality. Its refined design ensures minimal wobble and a reliable keystroke, making it an excellent option for users who prefer linear switches. The Dolphin's consistent performance and pleasing aesthetics make it a popular choice for both typing and gaming enthusiasts.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2020-11-01"),
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
			Name:             "POM",
			ShortDescription: "Linear switch with a POM housing for enhanced smoothness.",
			LongDescription:  "The Durock POM switch is a premium linear switch designed with a POM housing, known for its ability to enhance smoothness and reduce friction during keypresses. This switch is highly favored by enthusiasts who seek a top-tier linear experience with minimal scratchiness. The POM's superior material properties make it an excellent choice for those looking for a high-performance switch with a smooth and satisfying keystroke.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(63.5)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(63.5)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
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
			Name:             "Daybreak",
			ShortDescription: "Linear switch with a dawn-themed design and smooth actuation.",
			LongDescription:  "The Durock Daybreak switch features a visually striking dawn-themed design combined with smooth linear actuation. It offers a consistent and quiet typing experience, with minimal spring noise, making it ideal for users who prefer a peaceful and serene environment. The Daybreak's unique aesthetic and smooth performance make it a standout choice for those looking to add a touch of elegance to their keyboard.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2022-07-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(67)),
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
			Name:             "Owl",
			ShortDescription: "Tactile switch with a unique owl-themed housing and sharp bump.",
			LongDescription:  "The Durock Owl switch is a tactile switch that features an owl-themed housing and delivers a sharp, distinct tactile bump. Known for its precise feedback and unique aesthetic, the Owl switch is a favorite among tactile switch enthusiasts who appreciate both performance and visual appeal. This switch is perfect for users who want a well-defined tactile feel with a themed design that adds character to their keyboard.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2020-06-01"),
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
				TopHousingColorID:     ptr(58), // Gray
				BottomHousingColorID:  ptr(58), // Gray
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(41), // Sharp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Sea Turtle",
			ShortDescription: "Silent linear switch with a sea-themed design and smooth feel.",
			LongDescription:  "The Durock Sea Turtle is a silent linear switch with a sea-themed design that provides a smooth and quiet typing experience. It's perfect for those who require silent operation without sacrificing the smoothness of a linear switch. The Sea Turtle's aesthetic appeal and silent performance make it an excellent choice for users who need a tranquil workspace or who frequently type in shared environments.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     4,      // Silent Linear
			ReleaseDate:      parseDate("2021-10-01"),
			Available:        true,
			PricePoint:       3, // Expensive
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
				StemColorID:           ptr(62), // Teal
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(35), // Quiet
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Alpaca V2",
			ShortDescription: "Linear switch with a smooth feel and updated spring design.",
			LongDescription:  "The Durock Alpaca V2 is an improved version of the original Alpaca switch, featuring a smoother feel and an updated spring design. This linear switch is renowned for its consistent performance, minimal noise, and satisfying keystroke. The Alpaca V2's smooth actuation and vibrant color make it a popular choice among enthusiasts seeking both reliability and aesthetics in their typing experience.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
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
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(51), // White
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Marshmallow",
			ShortDescription: "Silent linear switch with a soft, marshmallow-like feel.",
			LongDescription:  "The Durock Marshmallow switch is a silent linear switch designed to provide a soft, marshmallow-like feel. With its quiet operation and gentle actuation, this switch is ideal for users who prefer a smooth, silent typing experience. The Marshmallow's unique texture and noise-dampening qualities make it a top choice for those seeking comfort and tranquility in their typing sessions.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     4,      // Silent Linear
			ReleaseDate:      parseDate("2022-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(40)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(3.8)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
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
			Name:             "Taro Ball",
			ShortDescription: "Tactile switch with a taro-themed design and smooth bump.",
			LongDescription:  "The Durock Taro Ball switch is a tactile switch that features a taro-themed design, delivering a smooth and satisfying tactile bump. This switch is well-loved for its unique aesthetic and reliable performance, making it a favorite among tactile switch enthusiasts. The Taro Ball combines visual appeal with a distinct tactile feel, offering a premium typing experience that stands out in any build.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2021-12-01"),
			Available:        true,
			PricePoint:       3, // Expensive
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
				StemColorID:           ptr(55), // Purple
				TopHousingColorID:     ptr(55), // Purple
				BottomHousingColorID:  ptr(55), // Purple
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(42), // Rounded
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Corgi",
			ShortDescription: "Linear switch with a corgi-themed design and buttery smoothness.",
			LongDescription:  "The Durock Corgi switch is a playful linear switch featuring a corgi-themed design and buttery smooth keystrokes. Known for its charming aesthetic and smooth operation, the Corgi switch is perfect for users who want to add a touch of fun to their keyboard while enjoying a premium typing experience. This switch is ideal for those who prioritize both performance and personality in their keyboard builds.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-04-01"),
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
				StemColorID:           ptr(56), // Orange
				TopHousingColorID:     ptr(51), // White
				BottomHousingColorID:  ptr(51), // White
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Shrimpy Click",
			ShortDescription: "Clicky switch with a vibrant design and sharp click.",
			LongDescription:  "The Durock Shrimpy Click is a vibrant clicky switch designed to provide a sharp, satisfying click with every keystroke. Featuring a unique and colorful design, this switch offers an engaging typing experience for those who enjoy auditory feedback and bold aesthetics. The Shrimpy Click is perfect for users who want a lively and responsive switch that delivers both performance and style.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2021-08-01"),
			Available:        true,
			PricePoint:       3, // Expensive

			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.2)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.2)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(57), // Pink
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				TactilityTypeID:       ptr(40), // Click Bar
				TactilityFeedbackID:   ptr(43), // Crisp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "White Lotus",
			ShortDescription: "Clicky switch with a floral aesthetic and distinct click.",
			LongDescription:  "The Durock White Lotus switch is a clicky switch that combines a floral aesthetic with a distinct, crisp click. This switch is designed for users who appreciate themed switches that offer both visual appeal and tactile feedback. The White Lotus switch's unique look and satisfying click make it a great choice for those looking to add a touch of elegance and performance to their keyboard.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2021-05-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.2)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.2)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(51), // White
				TopHousingColorID:     ptr(51), // White
				BottomHousingColorID:  ptr(51), // White
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				TactilityTypeID:       ptr(40), // Click Bar
				TactilityFeedbackID:   ptr(43), // Crisp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Thanos Click",
			ShortDescription: "Clicky switch with a deep, thocky click and unique theme.",
			LongDescription:  "The Durock Thanos Click switch offers a deep, thocky click sound paired with a powerful and unique theme. This switch is designed for enthusiasts who seek a distinctive auditory experience combined with a bold aesthetic. The Thanos Click switch is ideal for those who enjoy a commanding sound profile and a switch that stands out both visually and audibly.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2021-01-01"),
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
				StemColorID:           ptr(55), // Purple
				TopHousingColorID:     ptr(60), // Gold
				BottomHousingColorID:  ptr(55), // Purple
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				TactilityTypeID:       ptr(40), // Click Bar
				TactilityFeedbackID:   ptr(43), // Crisp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Peacock",
			ShortDescription: "Clicky switch with a colorful design and lively click.",
			LongDescription:  "The Durock Peacock switch is a clicky switch that boasts a colorful, peacock-themed design and a lively click, making it a favorite for those who want a vibrant aesthetic paired with satisfying tactile and auditory feedback. The Peacock switch is perfect for users who enjoy a bold and expressive keyboard that performs as well as it looks.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2022-04-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.2)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.2)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(54), // Green
				BottomHousingColorID:  ptr(54), // Green
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(31), // High
				TactilityTypeID:       ptr(40), // Click Bar
				TactilityFeedbackID:   ptr(43), // Crisp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Joyous Click",
			ShortDescription: "Clicky switch with a cheerful design and crisp click.",
			LongDescription:  "The Durock Joyous Click switch is known for its cheerful design and crisp, satisfying click. This switch offers a joyful typing experience, blending a playful aesthetic with responsive tactile feedback. Ideal for users who want to add a touch of happiness to their setup, the Joyous Click switch is a delightful addition to any keyboard that values both performance and fun.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2021-11-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(63)),
				ActuationPoint:        ptr(float32(2.2)),
				ActuationForce:        ptr(float32(48)),
				BottomOutForce:        ptr(float32(63)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.2)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(50), // Yellow
				TopHousingColorID:     ptr(56), // Orange
				BottomHousingColorID:  ptr(56), // Orange
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(30), // Medium
				TactilityTypeID:       ptr(40), // Click Bar
				TactilityFeedbackID:   ptr(43), // Crisp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Brightside Click",
			ShortDescription: "Clicky switch with a bright design and clear click sound.",
			LongDescription:  "The Durock Brightside Click switch features a bright and cheerful design, combined with a clear and crisp click sound. This switch is an excellent choice for users who appreciate a vibrant aesthetic along with satisfying tactile and auditory feedback. The Brightside Click is perfect for those looking to brighten up their keyboard with a switch that delivers both style and substance.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2022-02-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(62)),
				ActuationPoint:        ptr(float32(2.2)),
				ActuationForce:        ptr(float32(47)),
				BottomOutForce:        ptr(float32(62)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.2)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(50), // Yellow
				TopHousingColorID:     ptr(51), // White
				BottomHousingColorID:  ptr(51), // White
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(30), // Medium
				TactilityTypeID:       ptr(40), // Click Bar
				TactilityFeedbackID:   ptr(43), // Crisp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Firestorm",
			ShortDescription: "Clicky switch with a fiery design and explosive click.",
			LongDescription:  "The Durock Firestorm switch is a clicky switch that offers a dramatic, fiery design paired with an explosive click sound. This switch is perfect for users who want to make a bold statement with their keyboard, both visually and audibly. The Firestorm switch's powerful feedback and striking appearance make it an excellent choice for those seeking a dynamic and engaging typing experience.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     3,      // Clicky
			ReleaseDate:      parseDate("2021-03-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(68)),
				ActuationPoint:        ptr(float32(2.2)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(68)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.2)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(56), // Orange
				TopHousingColorID:     ptr(45), // Red
				BottomHousingColorID:  ptr(45), // Red
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(33), // Clicky
				SoundLevelID:          ptr(32), // Very High
				TactilityTypeID:       ptr(40), // Click Bar
				TactilityFeedbackID:   ptr(43), // Crisp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "T2",
			ShortDescription: "Enhanced tactile switch with smoother bump and less noise.",
			LongDescription:  "The Durock T2 is an enhanced version of the T1, offering a smoother tactile bump and reduced noise during actuation. Retaining the high-quality construction that Durock is known for, the T2 switch is a reliable choice for tactile switch fans who want a refined typing experience. This switch is perfect for users who appreciate precise tactile feedback with improved smoothness and reduced sound levels.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2020-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
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
				TopHousingColorID:     ptr(51), // White
				BottomHousingColorID:  ptr(51), // White
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(29), // Low
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(42), // Rounded
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Shrimp Tactile",
			ShortDescription: "Tactile switch with a unique pink housing and satisfying bump.",
			LongDescription:  "The Durock Shrimp Tactile switch is known for its unique pink housing and satisfying tactile bump. This switch offers both aesthetic appeal and excellent performance, making it a standout choice for users who value distinctive looks alongside great tactile feedback. The Shrimp Tactile switch is perfect for those who want a tactile typing experience with a touch of personality and style.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2022-05-01"),
			Available:        true,
			PricePoint:       3, // Expensive

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
			Name:             "Koala V2",
			ShortDescription: "Upgraded tactile switch with improved bump and smoothness.",
			LongDescription:  "The Durock Koala V2 is an upgraded version of the original Koala switch, offering a more refined tactile bump and enhanced smoothness. This switch is designed for enthusiasts who want a premium typing experience with reliable tactile feedback. The Koala V2's improvements in both feel and performance make it a top choice for those seeking an elevated tactile switch.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2021-06-01"),
			Available:        true,
			PricePoint:       3, // Expensive

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
				TopHousingColorID:     ptr(58), // Gray
				BottomHousingColorID:  ptr(58), // Gray
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(42), // Rounded
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "POM Tactile",
			ShortDescription: "Tactile switch with POM housing for a unique feel.",
			LongDescription:  "The Durock POM Tactile switch is designed with a POM housing, providing a unique feel and reduced friction for a smooth, satisfying tactile experience. This switch is perfect for those who seek a distinct tactile switch that offers both premium materials and excellent feedback. The POM Tactile switch's combination of smoothness and tactile response makes it a favorite among users looking for a high-quality switch with a different feel.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2021-03-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(63.5)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(63.5)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
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
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(42), // Rounded
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Daybreak Tactile",
			ShortDescription: "Tactile switch with a dawn-themed design and smooth bump.",
			LongDescription:  "The Durock Daybreak Tactile switch offers a dawn-themed design paired with a smooth tactile bump, providing an excellent balance between aesthetics and performance. This switch is ideal for users who want both style and substance in their typing experience. The Daybreak Tactile switch's unique look and reliable tactile feedback make it a great choice for those seeking a refined tactile switch with visual appeal.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2022-07-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(67)),
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
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(29), // Low
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(42), // Rounded
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Owl Tactile",
			ShortDescription: "Tactile switch with an owl-themed design and sharp bump.",
			LongDescription:  "The Durock Owl Tactile switch features an owl-themed design and a sharp tactile bump, providing precise feedback that is perfect for tactile switch enthusiasts who enjoy themed aesthetics. This switch is a great choice for users who want a well-defined tactile feel combined with a distinctive look, making it a unique addition to any keyboard.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2020-06-01"),
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
				TopHousingColorID:     ptr(58), // Gray
				BottomHousingColorID:  ptr(58), // Gray
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(41), // Sharp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Lilac Tactile",
			ShortDescription: "Tactile switch with a lilac color and smooth bump.",
			LongDescription:  "The Durock Lilac Tactile switch offers a soothing lilac color and a smooth tactile bump, making it a perfect choice for users who appreciate both aesthetics and performance. This switch is ideal for those who want a unique color-themed build without compromising on the quality of their tactile feedback. The Lilac Tactile switch's elegant design and reliable performance make it a standout in any keyboard.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2021-10-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(52)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(63), // Lavender
				BottomHousingColorID:  ptr(63), // Lavender
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(29), // Low
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(42), // Rounded
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Zilent V2",
			ShortDescription: "Silent tactile switch with a soft bump and minimal noise.",
			LongDescription:  "The Durock Zilent V2 is a silent tactile switch offering a soft bump and minimal noise, making it perfect for users who require a quiet typing experience without sacrificing tactile feedback. Known for its smooth actuation and quiet operation, the Zilent V2 is highly regarded among enthusiasts who value silence and performance in equal measure.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     5,      // Silent Tactile
			ReleaseDate:      parseDate("2022-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(53)),
				BottomOutForce:        ptr(float32(67)),
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
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(42), // Rounded
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Piano",
			ShortDescription: "Tactile switch with a musical theme and crisp bump.",
			LongDescription:  "The Durock Piano switch is a tactile switch that draws inspiration from the world of music, offering a crisp tactile bump and a unique aesthetic. This switch is designed for users who want a switch that not only performs well but also adds a touch of musical elegance to their setup. The Piano switch's combination of satisfying feedback and distinctive design makes it a popular choice among musicians and keyboard enthusiasts alike.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2021-05-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(52)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(51), // White
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(43), // Crisp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Lavender",
			ShortDescription: "Tactile switch with a lavender hue and smooth action.",
			LongDescription:  "The Durock Lavender switch offers a tactile bump with a calming lavender hue, providing a smooth typing experience with an appealing color scheme. This switch is perfect for those who want a stylish and satisfying tactile switch that stands out in both form and function. The Lavender switch's unique color and reliable performance make it a great choice for themed builds.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2022-06-01"),
			Available:        true,
			PricePoint:       3, // Expensive

			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(52)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(63), // Lavender
				BottomHousingColorID:  ptr(63), // Lavender
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(29), // Low
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(42), // Rounded
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Pudding",
			ShortDescription: "Tactile switch with a unique feel and playful design.",
			LongDescription:  "The Durock Pudding switch is a tactile switch that combines a unique feel with a playful design. This switch offers a satisfying tactile bump and is ideal for users who want a switch that stands out both in performance and appearance. The Pudding switch's fun and distinctive look, coupled with its reliable tactile feedback, make it a popular choice for those looking to add a touch of whimsy to their keyboard.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2021-08-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(63)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(63)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(51), // White
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(42), // Rounded
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Camel",
			ShortDescription: "Tactile switch with a sand-colored housing and smooth bump.",
			LongDescription:  "The Durock Camel switch is a tactile switch featuring a sand-colored housing and a smooth tactile bump. Inspired by the desert, this switch is designed for users who appreciate a unique, earthy aesthetic alongside a satisfying typing experience. The Camel switch's distinct design and reliable performance make it a great choice for those looking for a tactile switch with a natural, grounded feel.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2022-03-01"),
			Available:        true,
			PricePoint:       3, // Expensive
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
				TopHousingColorID:     ptr(72), // Tan
				BottomHousingColorID:  ptr(72), // Tan
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(42), // Rounded
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Azure",
			ShortDescription: "Tactile switch with an azure-themed design and crisp bump.",
			LongDescription:  "The Durock Azure switch offers a tactile bump with an azure-themed design, providing a crisp typing experience with a unique visual appeal. This switch is perfect for users who enjoy themed switches that deliver both style and performance. The Azure switch's combination of aesthetics and reliable tactile feedback makes it a standout choice for those looking to add a splash of color to their setup.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2021-12-01"),
			Available:        true,
			PricePoint:       2, // Average

			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(52)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(65), // Cyan
				BottomHousingColorID:  ptr(65), // Cyan
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(43), // Crisp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Coral Tactile",
			ShortDescription: "Tactile switch with a coral-themed design and strong bump.",
			LongDescription:  "The Durock Coral Tactile switch features a coral-themed design and a strong tactile bump, offering a distinct aesthetic and satisfying feedback. This switch is ideal for users who want a bold tactile switch with a unique theme, making it a perfect addition to any keyboard that values both performance and visual appeal.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     2,      // Tactile
			ReleaseDate:      parseDate("2022-05-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(55)),
				BottomOutForce:        ptr(float32(67)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(67), // Coral
				BottomHousingColorID:  ptr(67), // Coral
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(41), // Sharp
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "L1",
			ShortDescription: "Linear switch with smooth operation and minimal noise.",
			LongDescription:  "The Durock L1 is a linear switch known for its exceptionally smooth operation and minimal noise, making it an excellent choice for users seeking a quiet and fluid typing experience. This switch is favored for its consistency and low spring noise, providing a reliable performance that is ideal for both typing and gaming environments.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2020-05-01"),
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
			Name:             "Alpaca V1",
			ShortDescription: "Linear switch with a smooth feel and vibrant pink stem.",
			LongDescription:  "The Durock Alpaca V1 switch is a linear switch renowned for its smooth feel and vibrant pink stem. This switch offers a consistent typing experience with minimal spring noise, making it a favorite among enthusiasts seeking reliable performance. The Alpaca V1's iconic color and smooth action have made it a staple in the mechanical keyboard community.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2019-09-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(62)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(45)),
				BottomOutForce:        ptr(float32(62)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(false),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(57), // Pink
				TopHousingColorID:     ptr(51), // White
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Cobalt",
			ShortDescription: "Linear switch with a cobalt blue stem and smooth action.",
			LongDescription:  "The Durock Cobalt switch features a striking cobalt blue stem and provides a smooth linear action, making it a visually and functionally appealing choice for any keyboard build. This switch is designed for users who appreciate a deep color aesthetic combined with a consistent and reliable typing feel. The Cobalt switch's performance and design make it a top choice for those who want both style and substance in their linear switches.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-04-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(63)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(46)),
				BottomOutForce:        ptr(float32(63)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
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
			Name:             "Sunflower Linear",
			ShortDescription: "Linear switch with a bright yellow stem and buttery smoothness.",
			LongDescription:  "The Durock Sunflower Linear switch is known for its bright yellow stem and buttery smooth operation. It provides a quiet and fluid keystroke, ideal for those who want both a vibrant aesthetic and high performance. The Sunflower Linear switch's unique color and smooth feel make it an excellent choice for users looking for a switch that delivers both visual appeal and reliable performance.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-09-01"),
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
			Name:             "Ghost",
			ShortDescription: "Linear switch with a translucent housing and smooth keystroke.",
			LongDescription:  "The Durock Ghost switch features a translucent housing that highlights its smooth linear keystroke. This switch is designed for users who enjoy a clean aesthetic with the performance of a premium linear switch. The Ghost switch's combination of visual clarity and smooth action makes it a popular choice for those who want a sleek, minimalistic look without sacrificing quality.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2020-10-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(63)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(46)),
				BottomOutForce:        ptr(float32(63)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
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
			Name:             "Dolphin V2",
			ShortDescription: "Linear switch with improved smoothness and minimal spring noise.",
			LongDescription:  "The Durock Dolphin V2 switch is an upgraded version of the original, offering improved smoothness and minimal spring noise. Its consistent performance and pleasing aesthetics make it a popular choice for linear switch enthusiasts who demand reliability and quality. The Dolphin V2 switch is perfect for users seeking a refined linear switch that delivers on both feel and durability.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-03-01"),
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
				StemColorID:           ptr(47), // Blue
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Haze",
			ShortDescription: "Linear switch with a hazy design and buttery keystroke.",
			LongDescription:  "The Durock Haze switch is known for its hazy design and buttery keystroke, providing a smooth typing experience with a unique aesthetic. This switch is perfect for users who want a switch that combines style with performance, offering a consistent linear feel that is both comfortable and reliable. The Haze switch's misty appearance and smooth action make it a popular choice for enthusiasts looking to add a touch of mystery to their keyboard.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     1,      // Linear
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
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(58), // Gray
				TopHousingColorID:     ptr(53), // Smokey
				BottomHousingColorID:  ptr(53), // Smokey
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Mauve",
			ShortDescription: "Linear switch with a mauve-colored stem and smooth action.",
			LongDescription:  "The Durock Mauve switch offers a mauve-colored stem and a smooth linear action, making it a top choice for enthusiasts who prefer a unique color scheme and consistent performance. This switch's distinctive hue and smooth keystroke make it a standout option for users who value both aesthetics and reliable operation in their linear switches.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-08-01"),
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
				StemColorID:           ptr(64), // Magenta
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Midnight",
			ShortDescription: "Linear switch with a dark design and quiet operation.",
			LongDescription:  "The Durock Midnight switch is a linear switch characterized by its dark design and quiet operation, offering a smooth and subtle typing experience that is perfect for nighttime use. This switch is ideal for users who prefer a low-profile aesthetic and a quiet, reliable keystroke. The Midnight switch's combination of stealthy appearance and smooth action makes it an excellent choice for those who value discretion and performance.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2020-12-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(63)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(46)),
				BottomOutForce:        ptr(float32(63)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(53), // Smokey
				BottomHousingColorID:  ptr(53), // Smokey
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Lightwave",
			ShortDescription: "Linear switch with a light actuation force and smooth glide.",
			LongDescription:  "The Durock Lightwave switch provides a light actuation force and a smooth glide, making it ideal for users who prefer quick keystrokes with minimal resistance. This switch is perfect for those who want a responsive and effortless typing experience, with a smooth and consistent feel that enhances typing speed and comfort. The Lightwave switch's blend of light touch and smooth action makes it a favorite among users who value speed and precision.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2022-02-01"),
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
				StemColorID:           ptr(51), // White
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(28), // Very Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Euphoria",
			ShortDescription: "Linear switch with a euphoric feel and consistent action.",
			LongDescription:  "The Durock Euphoria switch is designed to provide a euphoric typing experience with a consistent linear action. This switch is perfect for those who want a smooth and satisfying keystroke with minimal noise, offering a blend of comfort and performance that enhances any typing session. The Euphoria switch's premium feel and reliable operation make it a top choice for enthusiasts seeking a high-quality linear switch.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-06-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(63)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(46)),
				BottomOutForce:        ptr(float32(63)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
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
			Name:             "Vanilla",
			ShortDescription: "Linear switch with a creamy smooth action and neutral design.",
			LongDescription:  "The Durock Vanilla switch offers a creamy smooth action and a neutral design, making it a versatile choice for any keyboard setup. This switch is known for its reliability and minimal spring noise, providing a consistent and quiet typing experience that is ideal for both work and play. The Vanilla switch's simplicity and smooth operation make it a dependable choice for those seeking a no-frills linear switch that delivers on performance.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-11-01"),
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
				StemColorID:           ptr(66), // Ivory
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Jet",
			ShortDescription: "Linear switch with a jet-black design and fast actuation.",
			LongDescription:  "The Durock Jet switch is known for its jet-black design and fast actuation, making it perfect for gamers who require quick and responsive keystrokes. This switch is designed for those who need performance and speed, offering a consistent and reliable feel that enhances gaming and typing efficiency. The Jet switch's sleek design and rapid response make it a top choice for competitive users who demand the best from their switches.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2022-04-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(58)),
				ActuationPoint:        ptr(float32(1.8)),
				ActuationForce:        ptr(float32(43)),
				BottomOutForce:        ptr(float32(58)),
				TotalTravel:           ptr(float32(3.8)),
				PreTravel:             ptr(float32(1.8)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(46), // Black
				BottomHousingColorID:  ptr(46), // Black
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Blossom",
			ShortDescription: "Linear switch with a floral design and smooth keystroke.",
			LongDescription:  "The Durock Blossom switch features a delicate floral design combined with a smooth linear keystroke. This switch is perfect for those who appreciate a gentle aesthetic alongside reliable performance. The Blossom switch's elegant appearance and buttery smooth action make it a popular choice for users looking to add a touch of grace to their keyboard while maintaining top-tier performance.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-07-01"),
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
			Name:             "Ink Black V2",
			ShortDescription: "Linear switch with a unique ink-colored housing and buttery smoothness.",
			LongDescription:  "The Durock Ink Black V2 switch offers a unique ink-colored housing and buttery smoothness, providing a premium typing experience with a distinctive look. This switch is ideal for users who want high performance and a unique design. The Ink Black V2's combination of style and substance makes it a top choice for those seeking a standout linear switch that delivers on both form and function.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2020-03-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(48)),
				BottomOutForce:        ptr(float32(65)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(53), // Smokey
				BottomHousingColorID:  ptr(53), // Smokey
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Cotton Candy",
			ShortDescription: "Linear switch with a sweet aesthetic and soft keystroke.",
			LongDescription:  "The Durock Cotton Candy switch is a linear switch known for its sweet aesthetic and soft keystroke, offering a playful look with smooth performance. This switch is perfect for users who want a switch that adds a fun element to their keyboard while providing a reliable typing experience. The Cotton Candy switch's vibrant colors and gentle actuation make it a delightful choice for those seeking a lighthearted and high-quality switch.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-02-01"),
			Available:        true,
			PricePoint:       2, // Average
			Details: models.Details{
				SpringForce:           ptr(float32(60)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(43)),
				BottomOutForce:        ptr(float32(60)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(57), // Pink
				TopHousingColorID:     ptr(47), // Blue
				BottomHousingColorID:  ptr(47), // Blue
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Silent Linear",
			ShortDescription: "Silent linear switch with smooth keystroke and noise reduction.",
			LongDescription:  "The Durock Silent Linear switch is designed for those who need a quiet typing experience without sacrificing performance. It offers a smooth linear keystroke with built-in dampening to reduce noise, making it ideal for office or shared spaces. The Silent Linear switch's blend of silence and smooth action makes it a reliable choice for environments where quiet operation is essential.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     4,      // Silent Linear
			ReleaseDate:      parseDate("2020-11-01"),
			Available:        true,
			PricePoint:       3, // Expensive
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
			Name:             "Silent T1",
			ShortDescription: "Silent tactile switch with a refined bump and minimal noise.",
			LongDescription:  "The Durock Silent T1 switch offers a tactile typing experience with a refined bump and minimal noise. This switch is perfect for users who want tactile feedback without the accompanying sound, making it suitable for quiet environments. The Silent T1's balance of tactile feel and noise reduction makes it a top choice for those seeking a silent yet satisfying switch.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     5,      // Silent Tactile
			ReleaseDate:      parseDate("2021-06-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(53)),
				BottomOutForce:        ptr(float32(67)),
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
				TactilityTypeID:       ptr(38), // Leaf Spring
				TactilityFeedbackID:   ptr(42), // Rounded
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Thocky",
			ShortDescription: "Specialty linear switch with a deep, thocky sound profile.",
			LongDescription:  "The Durock Thocky switch is a specialty linear switch known for its deep, thocky sound profile. It offers a unique auditory experience with smooth keystrokes, making it a favorite among enthusiasts who enjoy distinct sound feedback. The Thocky switch's combination of smooth action and rich sound makes it an excellent choice for users who prioritize both feel and sound in their typing experience.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive

			Details: models.Details{
				SpringForce:           ptr(float32(67)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(50)),
				BottomOutForce:        ptr(float32(67)),
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
				SoundTypeID:           ptr(34), // Thocky
				SoundLevelID:          ptr(30), // Medium
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "U4 Silent",
			ShortDescription: "Silent tactile switch with soft bump and ultra-quiet operation.",
			LongDescription:  "The Durock U4 Silent switch is designed to provide a soft tactile bump with ultra-quiet operation, making it ideal for environments where noise reduction is critical. This switch offers smooth, dampened keystrokes with reliable tactile feedback, ensuring a satisfying typing experience without the noise. The U4 Silent switch's balance of feel and silence makes it a top choice for users who need a quiet yet responsive switch.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     5,      // Silent Tactile
			ReleaseDate:      parseDate("2020-09-01"),
			Available:        true,
			PricePoint:       3, // Expensive

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
				StemColorID:           ptr(51), // White
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
			Name:             "Light Fog",
			ShortDescription: "Specialty switch with a translucent housing and mysterious feel.",
			LongDescription:  "The Durock Light Fog switch offers a specialty design with a translucent housing that creates a mysterious aesthetic. This switch provides a smooth keystroke and is perfect for those who appreciate both performance and style. The Light Fog switch's combination of unique appearance and reliable action makes it a standout choice for users who want a switch that delivers on both looks and feel.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-03-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(63)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(46)),
				BottomOutForce:        ptr(float32(63)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(26), // Polycarbonate
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(58), // Gray
				TopHousingColorID:     ptr(52), // Transparent
				BottomHousingColorID:  ptr(52), // Transparent
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Shadow",
			ShortDescription: "Specialty switch with a dark design and smooth performance.",
			LongDescription:  "The Durock Shadow switch is a specialty switch featuring a dark design and smooth performance. It offers a quiet typing experience and is ideal for those who want a switch with a stealthy appearance and reliable action. The Shadow switch's combination of aesthetics and smooth operation makes it a great choice for users who prefer a subtle, understated look without compromising on performance.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-08-01"),
			Available:        true,
			PricePoint:       2, // Average

			Details: models.Details{
				SpringForce:           ptr(float32(63)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(46)),
				BottomOutForce:        ptr(float32(63)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(46), // Black
				TopHousingColorID:     ptr(53), // Smokey
				BottomHousingColorID:  ptr(53), // Smokey
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(false),
			},
		},
		{
			Name:             "Marshmallow Silent",
			ShortDescription: "Silent linear switch with a soft, marshmallow-like feel.",
			LongDescription:  "The Durock Marshmallow Silent switch combines a soft, marshmallow-like feel with silent operation, offering a unique typing experience that's both quiet and satisfying. This switch is perfect for those who want a silent switch with a distinctive touch, providing smooth and noise-free keystrokes in any environment. The Marshmallow Silent switch's blend of comfort and silence makes it a popular choice for those seeking a premium, quiet typing experience.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     4,      // Silent Linear
			ReleaseDate:      parseDate("2022-01-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(55)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(40)),
				BottomOutForce:        ptr(float32(55)),
				TotalTravel:           ptr(float32(3.8)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
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
			Name:             "L4",
			ShortDescription: "Specialty linear switch with enhanced smoothness and reduced scratchiness.",
			LongDescription:  "The Durock L4 switch is a specialty linear switch offering enhanced smoothness and reduced scratchiness. Designed for users who seek a refined typing experience with consistent linear action, the L4 switch provides a reliable and satisfying keystroke that is ideal for both work and gaming. The L4's balance of smoothness and durability makes it a top choice for those looking to upgrade their linear switches.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2022-04-01"),
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
			Name:             "Daydream",
			ShortDescription: "Specialty switch with a dream-themed design and smooth keystroke.",
			LongDescription:  "The Durock Daydream switch offers a dream-themed design with a smooth keystroke, providing a unique aesthetic alongside reliable performance. This switch is perfect for those who want a switch that captures the imagination while delivering excellent typing feedback. The Daydream switch's combination of style and substance makes it a popular choice for users who want a switch that stands out in both form and function.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-06-01"),
			Available:        true,
			PricePoint:       2, // Average

			Details: models.Details{
				SpringForce:           ptr(float32(63)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(46)),
				BottomOutForce:        ptr(float32(63)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(63), // Lavender
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Aqua King",
			ShortDescription: "Linear switch with a clear housing and smooth, fluid action.",
			LongDescription:  "The Durock Aqua King switch features a clear housing that highlights its smooth, fluid linear action. This switch is designed for users who want a transparent aesthetic combined with top-tier linear performance. The Aqua King switch's crystal-clear appearance and silky-smooth keystroke make it a standout choice for those seeking both visual clarity and superior performance in their linear switches.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     1,      // Linear
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
				FactoryLubed:          ptr(true),
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
			Name:             "Ultra",
			ShortDescription: "Specialty switch with an ultra-smooth linear feel and minimal noise.",
			LongDescription:  "The Durock Ultra switch offers an ultra-smooth linear feel with minimal noise, providing a premium typing experience that's both quiet and satisfying. This switch is ideal for those who demand the highest level of performance from their switches, offering consistent keystrokes with a smooth and seamless action. The Ultra switch's blend of silence and smoothness makes it a top choice for users seeking a refined and high-performance linear switch.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2022-07-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(65)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(48)),
				BottomOutForce:        ptr(float32(65)),
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
			Name:             "Crystal Glacier",
			ShortDescription: "Specialty switch with a crystal-clear housing and icy smoothness.",
			LongDescription:  "The Durock Crystal Glacier switch is a specialty switch with a crystal-clear housing that offers icy smoothness in its linear action. This switch is designed for those who want a striking aesthetic combined with seamless performance. The Crystal Glacier switch's blend of visual appeal and smooth keystrokes makes it a popular choice for users who value both style and high performance in their linear switches.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-10-01"),
			Available:        true,
			PricePoint:       3, // Expensive
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
			Name:             "Emerald",
			ShortDescription: "Linear switch with an emerald hue and luxurious smoothness.",
			LongDescription:  "The Durock Emerald switch features a rich emerald hue and luxurious smoothness, providing a premium typing experience with a touch of elegance. This switch is perfect for users who want both performance and style in their keyboard. The Emerald switch's combination of deep color and refined keystroke makes it a standout choice for those who seek a high-quality linear switch with a luxurious feel.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2022-09-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(63)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(46)),
				BottomOutForce:        ptr(float32(63)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(54), // Green
				TopHousingColorID:     ptr(54), // Green
				BottomHousingColorID:  ptr(54), // Green
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
		{
			Name:             "Mystic",
			ShortDescription: "Specialty switch with a mystical design and smooth action.",
			LongDescription:  "The Durock Mystic switch offers a mystical design combined with smooth linear action, making it a favorite for users who want a switch with an enchanting appearance and reliable performance. The Mystic switch's unique aesthetic and consistent feel make it a great choice for those who appreciate both form and function in their keyboard switches.",
			ManufacturerID:   ptr(7), // Durock
			BrandID:          ptr(7), // Durock
			SwitchTypeID:     1,      // Linear
			ReleaseDate:      parseDate("2021-09-01"),
			Available:        true,
			PricePoint:       3, // Expensive
			Details: models.Details{
				SpringForce:           ptr(float32(63)),
				ActuationPoint:        ptr(float32(2.0)),
				ActuationForce:        ptr(float32(46)),
				BottomOutForce:        ptr(float32(63)),
				TotalTravel:           ptr(float32(4.0)),
				PreTravel:             ptr(float32(2.0)),
				FactoryLubed:          ptr(true),
				StemMaterialID:        ptr(23), // POM
				TopHousingMaterialID:  ptr(26), // Polycarbonate
				BaseHousingMaterialID: ptr(24), // Nylon
				SpringMaterialTypeID:  ptr(17), // Steel
				StemColorID:           ptr(55), // Purple
				TopHousingColorID:     ptr(49), // Clear
				BottomHousingColorID:  ptr(49), // Clear
				PinConfigurationID:    ptr(75), // 5-Pin
				SoundTypeID:           ptr(36), // Creamy
				SoundLevelID:          ptr(29), // Low
				HasShineThrough:       ptr(true),
			},
		},
	}

	err := processSwitches(tx, durock, admin)
	if err != nil {
		return err
	}

	return nil
}
