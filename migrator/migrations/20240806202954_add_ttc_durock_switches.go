package migrations

import (
	"time"

	"github.com/Bparsons0904/deadigations"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	deadigations.RegisterMigration(deadigations.Migration{
		ID:          "20240806202954",
		Description: "Add description of changes",
		Migrate: func(tx *gorm.DB) error {
			type Type struct {
				ID        int       `gorm:"primaryKey;autoIncrement"              json:"id"`
				Name      string    `gorm:"type:varchar(50);not null"             json:"name"`
				Code      string    `gorm:"type:varchar(50);not null;uniqueIndex" json:"code"`
				Category  string    `gorm:"type:varchar(50);not null;index"       json:"category"`
				CreatedAt time.Time `gorm:"autoCreateTime"                        json:"createdAt"`
				UpdatedAt time.Time `gorm:"autoUpdateTime"                        json:"updatedAt"`
			}

			type Producer struct {
				ID        int       `gorm:"type:int;primaryKey;autoIncrement"     json:"id"`
				Name      string    `gorm:"type:varchar(50);not null"             json:"name"`
				Alias     string    `gorm:"type:varchar(50);not null;uniqueIndex" json:"alias"`
				SiteURL   string    `gorm:"type:varchar(50)"                      json:"siteURL,omitempty"`
				CreatedAt time.Time `gorm:"autoCreateTime"                        json:"createdAt"`
				UpdateAt  time.Time `gorm:"autoUpdateTime"                        json:"updatedAt"`
			}

			type Switch struct {
				ID               uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v7();primaryKey"                                  json:"id"`
				Name             string     `gorm:"type:varchar(50);not null;index:idx_name;uniqueIndex:idx_name_manufacturer_brand" json:"name"`
				ShortDescription string     `gorm:"type:varchar(255);not null"                                                       json:"shortDescription"`
				LongDescription  string     `gorm:"type:text;not null"                                                               json:"longDescription"`
				ManufacturerID   *int       `gorm:"type:int;index;uniqueIndex:idx_name_manufacturer_brand"                           json:"manufacturerId,omitempty"`
				Manufacturer     *Producer  `gorm:"foreignKey:ManufacturerID"                                                        json:"manufacturer,omitempty"`
				BrandID          *int       `gorm:"type:int;index;uniqueIndex:idx_name_manufacturer_brand"                           json:"brandId,omitempty"`
				Brand            *Producer  `gorm:"foreignKey:BrandID"                                                               json:"brand,omitempty"`
				SwitchTypeID     int        `gorm:"type:int;not null;index;"                                                         json:"switchTypeId"`
				SwitchType       *Type      `gorm:"foreignKey:SwitchTypeID"                                                          json:"switchType,omitempty"`
				ReleaseDate      *time.Time `gorm:"type:date"                                                                        json:"releaseDate,omitempty"`
				Available        bool       `gorm:"type:boolean;default:true"                                                        json:"available"`
				PricePoint       int        `gorm:"type:int;not null"                                                                json:"pricePoint"`
				CreatedAt        time.Time  `gorm:"autoCreateTime"                                                                   json:"createdAt"`
				UpdatedAt        time.Time  `gorm:"autoUpdateTime"                                                                   json:"updatedAt"`
			}

			ttcSwitches := []Switch{
				{
					Name:             "TTC Gold Pink",
					ShortDescription: "Tactile switch with a light actuation force and smooth bump.",
					LongDescription:  "The TTC Gold Pink is a tactile switch with a light actuation force and a smooth, rounded bump, making it suitable for both typing and gaming.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2019-11-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "TTC Bluish White",
					ShortDescription: "Clicky switch with a crisp tactile bump and audible click.",
					LongDescription:  "The TTC Bluish White switch is known for its crisp tactile bump and satisfying audible click, providing a distinct typing experience for those who enjoy clicky switches.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2018-08-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "TTC Silent Red",
					ShortDescription: "Silent linear switch with a smooth keystroke and reduced noise.",
					LongDescription:  "The TTC Silent Red switch offers a smooth linear action with dampened noise, making it ideal for quiet environments where noise reduction is essential.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2020-05-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "TTC Gold Brown V3",
					ShortDescription: "Tactile switch with a distinct bump and gold-plated springs.",
					LongDescription:  "TTC Gold Brown V3 switches feature a tactile bump with gold-plated springs, offering a premium typing experience with improved durability and performance.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "TTC Wild",
					ShortDescription: "Linear switch with a smooth actuation and wild-themed aesthetics.",
					LongDescription:  "The TTC Wild switch offers a smooth linear actuation, coupled with wild-themed aesthetics that add a unique flair to any keyboard build.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2023-01-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				{
					Name:             "TTC Flame Red",
					ShortDescription: "Linear switch with a smooth keystroke and vibrant red color.",
					LongDescription:  "The TTC Flame Red is a linear switch characterized by its smooth keystroke and vibrant red color, offering both aesthetic appeal and reliable performance.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2017-09-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				{
					Name:             "TTC Silent Brown",
					ShortDescription: "Silent tactile switch with a smooth bump and reduced noise.",
					LongDescription:  "The TTC Silent Brown switch provides a tactile bump with dampened noise, perfect for those who seek a tactile typing experience without the accompanying noise.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     5,      // Silent Tactile
					ReleaseDate:      parseDate("2020-02-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "TTC Frozen Silent",
					ShortDescription: "Silent linear switch with a cold-themed design and smooth keystroke.",
					LongDescription:  "The TTC Frozen Silent switch combines a silent linear action with a cold-themed design, offering a unique aesthetic and a quiet typing experience.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2021-12-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "TTC Matrix White",
					ShortDescription: "Tactile switch with a noticeable bump and bright white color.",
					LongDescription:  "The TTC Matrix White switch offers a tactile typing experience with a noticeable bump and a bright white color, making it a popular choice for custom builds.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2018-04-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				{
					Name:             "TTC Speed Silver",
					ShortDescription: "Linear switch with a short actuation distance for rapid key presses.",
					LongDescription:  "The TTC Speed Silver switch is designed for rapid key presses with a short actuation distance, perfect for gamers who require quick responsiveness.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2019-02-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "TTC Blue",
					ShortDescription: "Classic clicky switch with a pronounced tactile bump and audible click.",
					LongDescription:  "The TTC Blue switch is a staple in the mechanical keyboard community, offering a distinct tactile bump and audible click, ideal for typists who enjoy a classic clicky experience.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2015-07-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				{
					Name:             "TTC Gold Blue",
					ShortDescription: "Enhanced clicky switch with gold-plated springs for improved durability.",
					LongDescription:  "The TTC Gold Blue switch offers an enhanced clicky experience with gold-plated springs, providing increased durability and a satisfying audible click with a smooth tactile bump.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2020-03-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "TTC Gold Blue V2",
					ShortDescription: "Updated version of the TTC Gold Blue with refined click mechanism.",
					LongDescription:  "The TTC Gold Blue V2 switch is an updated version of the original Gold Blue, featuring a refined click mechanism for a more satisfying tactile and auditory experience.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2021-05-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "TTC Diamond Pink",
					ShortDescription: "Clicky switch with a unique pink color and diamond-like sound.",
					LongDescription:  "The TTC Diamond Pink switch stands out with its unique pink color and a click sound that resembles a diamond's brilliance, offering both aesthetic appeal and a distinct auditory feedback.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2019-09-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "TTC Golden Green V3",
					ShortDescription: "Clicky switch with a pronounced tactile bump and enhanced click sound.",
					LongDescription:  "The TTC Golden Green V3 switch offers a pronounced tactile bump with an enhanced click sound, designed for those who appreciate a more noticeable and satisfying clicky feel.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2022-08-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "TTC Hart Blue",
					ShortDescription: "Clicky switch with a unique heart-shaped stem design.",
					LongDescription:  "The TTC Hart Blue switch features a unique heart-shaped stem design, providing a distinct clicky experience with both visual and tactile appeal.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2020-11-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "TTC Speed Purple",
					ShortDescription: "Fast actuation clicky switch designed for gaming enthusiasts.",
					LongDescription:  "The TTC Speed Purple switch is a clicky switch designed for gaming enthusiasts, offering fast actuation and a satisfying click sound that enhances the gaming experience.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2019-04-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "TTC Watermelon Milkshake",
					ShortDescription: "Novelty clicky switch with a fruity aesthetic and satisfying click.",
					LongDescription:  "The TTC Watermelon Milkshake switch is a novelty clicky switch with a fruity aesthetic and a satisfying click, perfect for those who enjoy a unique and playful keyboard setup.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2021-03-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "TTC Crystal Ice",
					ShortDescription: "Clicky switch with a transparent housing and icy click sound.",
					LongDescription:  "The TTC Crystal Ice switch offers a clicky experience with a transparent housing and an icy click sound, combining aesthetics with performance for a unique typing experience.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2018-06-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				{
					Name:             "TTC Light Ring",
					ShortDescription: "Clicky switch with integrated LED support for enhanced lighting.",
					LongDescription:  "The TTC Light Ring switch is a clicky switch with integrated LED support, providing enhanced lighting effects that can elevate the aesthetics of any keyboard build.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2019-07-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "TTC Brown",
					ShortDescription: "Classic tactile switch with a gentle bump and smooth action.",
					LongDescription:  "The TTC Brown switch is a classic tactile switch known for its gentle bump and smooth keystroke, offering a balanced typing experience ideal for both typing and gaming.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2015-03-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				{
					Name:             "TTC Silent Brown V2",
					ShortDescription: "Silent tactile switch with a refined bump and reduced noise.",
					LongDescription:  "The TTC Silent Brown V2 offers a refined tactile bump with reduced noise, making it an excellent choice for those who seek a quiet tactile typing experience.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     5,      // Silent Tactile
					ReleaseDate:      parseDate("2021-11-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "TTC Gold Brown",
					ShortDescription: "Tactile switch with gold-plated springs for enhanced durability.",
					LongDescription:  "The TTC Gold Brown switch features gold-plated springs for enhanced durability and a satisfying tactile bump, offering a premium typing experience with improved longevity.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2019-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "TTC Watermelon V2",
					ShortDescription: "Tactile switch with a fruity aesthetic and smooth bump.",
					LongDescription:  "The TTC Watermelon V2 switch offers a tactile typing experience with a smooth bump and a fruity aesthetic, making it a playful addition to any keyboard setup.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2020-09-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "TTC Heart",
					ShortDescription: "Tactile switch with a unique heart-shaped stem and gentle bump.",
					LongDescription:  "The TTC Heart switch features a unique heart-shaped stem that provides a gentle tactile bump, offering both a distinct aesthetic and a satisfying typing experience.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2020-02-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "TTC Light Salmon",
					ShortDescription: "Tactile switch with a light actuation force and smooth bump.",
					LongDescription:  "The TTC Light Salmon switch offers a tactile typing experience with a light actuation force and smooth bump, making it suitable for users who prefer a gentle tactile feel.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2019-04-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "TTC Honey V3",
					ShortDescription: "Tactile switch with a smooth bump and honey-themed design.",
					LongDescription:  "The TTC Honey V3 switch offers a smooth tactile bump with a honey-themed design, providing a sweet typing experience with both visual and tactile appeal.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2022-07-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "TTC Purple Sky",
					ShortDescription: "Tactile switch with a medium actuation force and sky-themed color.",
					LongDescription:  "The TTC Purple Sky switch features a tactile bump with a medium actuation force and a sky-themed color, offering a unique aesthetic and a satisfying typing feel.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-05-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "TTC Purple Starlight",
					ShortDescription: "Tactile switch with a celestial aesthetic and distinct bump.",
					LongDescription:  "The TTC Purple Starlight switch offers a tactile typing experience with a distinct bump and a celestial aesthetic, perfect for users seeking a unique look and feel.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2019-11-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "TTC Greystar",
					ShortDescription: "Tactile switch with a crisp bump and neutral color palette.",
					LongDescription:  "The TTC Greystar switch features a tactile bump with a crisp feel and a neutral color palette, offering a balanced typing experience with a professional aesthetic.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2018-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "TTC Titan Heart",
					ShortDescription: "Tactile switch with a strong bump and heart-themed design.",
					LongDescription:  "The TTC Titan Heart switch offers a strong tactile bump with a heart-themed design, providing both aesthetic appeal and a satisfying typing experience.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2020-11-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "TTC V3 Blue",
					ShortDescription: "Tactile switch with a refined bump and blue-themed design.",
					LongDescription:  "The TTC V3 Blue switch features a refined tactile bump and a blue-themed design, offering a smooth typing experience with an appealing aesthetic.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2018-12-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "TTC Pink Girl",
					ShortDescription: "Tactile switch with a playful pink color and gentle bump.",
					LongDescription:  "The TTC Pink Girl switch offers a tactile typing experience with a gentle bump and a playful pink color, perfect for users seeking a fun and lighthearted aesthetic.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-07-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "TTC Leaf",
					ShortDescription: "Tactile switch with a unique leaf-shaped stem and smooth bump.",
					LongDescription:  "The TTC Leaf switch features a unique leaf-shaped stem that provides a smooth tactile bump, offering a distinct typing experience with a touch of nature-inspired design.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2019-05-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "TTC Silent Red V3",
					ShortDescription: "Silent linear switch with enhanced smoothness and reduced noise.",
					LongDescription:  "The TTC Silent Red V3 is an updated version offering enhanced smoothness and significantly reduced noise, making it ideal for quiet environments where a silent typing experience is essential.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2022-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "TTC Silent Black",
					ShortDescription: "Silent linear switch with a heavier actuation force for quieter operation.",
					LongDescription:  "The TTC Silent Black switch is designed with a heavier actuation force, providing a silent linear typing experience that is perfect for those who require a quiet and smooth keystroke.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2021-03-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "TTC Silent Sky",
					ShortDescription: "Silent tactile switch with a gentle bump and quiet operation.",
					LongDescription:  "The TTC Silent Sky switch combines a tactile bump with silent operation, offering a smooth typing experience that is both satisfying and quiet, suitable for office environments.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     5,      // Silent Tactile
					ReleaseDate:      parseDate("2020-08-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "TTC Silver Blue",
					ShortDescription: "Specialty linear switch with fast actuation for gaming enthusiasts.",
					LongDescription:  "The TTC Silver Blue switch is a specialty linear switch with fast actuation and low travel distance, specifically designed for gaming enthusiasts who require quick responsiveness.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2019-11-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "TTC Ace",
					ShortDescription: "Premium tactile switch with a unique actuation curve and smooth bump.",
					LongDescription:  "The TTC Ace switch offers a premium tactile experience with a unique actuation curve, providing a smooth bump that enhances the typing experience, making it a top choice for enthusiasts.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2023-04-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "TTC Gold Pink V2",
					ShortDescription: "Updated tactile switch with improved smoothness and actuation.",
					LongDescription:  "The TTC Gold Pink V2 switch is an updated version of the original Gold Pink, offering improved smoothness and a more refined tactile bump, ideal for both gaming and typing.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2022-06-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "TTC Moonlight Silver",
					ShortDescription: "Linear switch with a moonlight-themed design and smooth action.",
					LongDescription:  "The TTC Moonlight Silver switch offers a linear typing experience with a moonlight-themed design, providing smooth action and aesthetic appeal for keyboard enthusiasts.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-05-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "TTC Sunset",
					ShortDescription: "Tactile switch with a sunset-themed design and distinctive bump.",
					LongDescription:  "The TTC Sunset switch features a tactile bump with a sunset-themed design, offering a unique typing experience that combines visual appeal with satisfying tactile feedback.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-09-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "TTC Wild Blue",
					ShortDescription: "Tactile switch with a distinctive wild-themed design and bump.",
					LongDescription:  "The TTC Wild Blue switch offers a tactile typing experience with a distinctive wild-themed design and a pronounced bump, perfect for users seeking a bold and satisfying tactile feel.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2020-07-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "TTC Coral Sea",
					ShortDescription: "Linear switch with a coral-themed design and smooth keystroke.",
					LongDescription:  "The TTC Coral Sea switch offers a linear typing experience with a coral-themed design, providing a smooth keystroke and aesthetic appeal for custom keyboard builds.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2022-08-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "TTC Tiger Light",
					ShortDescription: "Linear switch with a tiger-themed design and light actuation force.",
					LongDescription:  "The TTC Tiger Light switch features a linear action with a tiger-themed design and light actuation force, providing a smooth typing experience with a unique aesthetic.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2023-03-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "TTC Forest Green",
					ShortDescription: "Tactile switch with a forest-themed design and smooth bump.",
					LongDescription:  "The TTC Forest Green switch offers a tactile typing experience with a smooth bump and a forest-themed design, perfect for those who appreciate nature-inspired aesthetics.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2019-10-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "TTC Mars Orange",
					ShortDescription: "Tactile switch with a Mars-themed design and distinct tactile bump.",
					LongDescription:  "The TTC Mars Orange switch features a distinct tactile bump and a Mars-themed design, offering a unique typing experience with both visual and tactile appeal.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2020-02-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "TTC Neptune Blue",
					ShortDescription: "Linear switch with a Neptune-themed design and smooth keystroke.",
					LongDescription:  "The TTC Neptune Blue switch offers a linear typing experience with a Neptune-themed design, providing a smooth keystroke and an aquatic aesthetic for custom keyboard builds.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-05-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "TTC Cool Mint",
					ShortDescription: "Linear switch with a mint-themed design and cool aesthetic.",
					LongDescription:  "The TTC Cool Mint switch provides a linear typing experience with a mint-themed design, offering a smooth keystroke and a refreshing aesthetic for keyboard enthusiasts.",
					ManufacturerID:   ptr(6), // TTC
					BrandID:          ptr(6), // TTC
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2022-10-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
			}

			processSwitches := func(switches []Switch) error {
				for _, s := range switches {
					if err := tx.Create(&s).Error; err != nil {
						return err
					}

					// Sleep for 5 ms to ensure unique timestamps
					time.Sleep(5 * time.Millisecond)
				}

				return nil
			}

			err := processSwitches(ttcSwitches)
			if err != nil {
				return err
			}

			durockSwitches := []Switch{
				{
					Name:             "Durock T1",
					ShortDescription: "Tactile switch with a sharp bump and POM stem.",
					LongDescription:  "The Durock T1 is a tactile switch known for its sharp tactile bump and POM stem, offering a satisfying typing experience with great stability and minimal wobble. It's a popular choice among enthusiasts for its smoothness and reliability.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2019-04-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Durock L7",
					ShortDescription: "Linear switch with smooth operation and low noise.",
					LongDescription:  "The Durock L7 is a linear switch that offers exceptionally smooth operation with minimal spring noise. It is highly regarded for its consistency and is ideal for users who prefer a quiet typing experience without tactile feedback.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Durock Koala",
					ShortDescription: "Tactile switch inspired by Holy Panda design.",
					LongDescription:  "The Durock Koala switch is inspired by the Holy Panda design, featuring a medium to heavy tactile bump with a smooth, rounded feel. It is popular for its unique feel and high-quality construction, making it a favorite for tactile switch enthusiasts.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2020-03-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Durock Sunflower",
					ShortDescription: "Linear switch with a bright yellow stem and smooth feel.",
					LongDescription:  "The Durock Sunflower is a linear switch known for its bright yellow stem and incredibly smooth feel. It's designed for those who appreciate a vibrant aesthetic combined with a quiet and fluid keystroke.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-09-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Durock Shrimp",
					ShortDescription: "Tactile switch with a unique pink housing and sharp bump.",
					LongDescription:  "The Durock Shrimp switch is a tactile switch featuring a unique pink housing and a sharp, noticeable bump. It's well-loved for its distinctive look and excellent performance, providing a delightful typing experience.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2022-05-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Durock Dolphin",
					ShortDescription: "Linear switch with a smooth, consistent feel.",
					LongDescription:  "The Durock Dolphin is a linear switch that provides a smooth and consistent typing experience. It's characterized by its blue stem and excellent build quality, making it a popular choice for those who prefer linear switches.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-11-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Durock POM",
					ShortDescription: "Linear switch with a POM housing for enhanced smoothness.",
					LongDescription:  "The Durock POM switch is a linear switch designed with a POM housing to enhance smoothness and reduce friction. It's favored by enthusiasts who seek a premium linear experience with minimal scratchiness.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Durock Daybreak",
					ShortDescription: "Linear switch with a dawn-themed design and smooth actuation.",
					LongDescription:  "The Durock Daybreak switch features a dawn-themed design with a smooth linear actuation. It offers a consistent typing experience with minimal spring noise, making it ideal for users who prefer a quiet linear switch.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2022-07-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Durock Owl",
					ShortDescription: "Tactile switch with a unique owl-themed housing and sharp bump.",
					LongDescription:  "The Durock Owl switch is a tactile switch featuring an owl-themed housing and a sharp tactile bump. It is praised for its precise feel and distinct aesthetic, making it a popular choice among tactile switch enthusiasts.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2020-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Durock Sea Turtle",
					ShortDescription: "Silent linear switch with a sea-themed design and smooth feel.",
					LongDescription:  "The Durock Sea Turtle is a silent linear switch with a sea-themed design, providing a smooth and quiet typing experience. It's perfect for those who require silent operation without sacrificing smoothness.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2021-10-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Durock Alpaca V2",
					ShortDescription: "Linear switch with a smooth feel and updated spring design.",
					LongDescription:  "The Durock Alpaca V2 is a linear switch known for its smooth feel and updated spring design, offering a consistent typing experience with minimal noise. It's a favorite among enthusiasts for its reliable performance.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-02-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Durock Marshmallow",
					ShortDescription: "Silent linear switch with a soft, marshmallow-like feel.",
					LongDescription:  "The Durock Marshmallow switch is a silent linear switch that provides a soft, marshmallow-like feel, offering a quiet and satisfying typing experience. It's a great choice for those who prefer silent operation with a unique feel.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2022-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Durock Taro Ball",
					ShortDescription: "Tactile switch with a taro-themed design and smooth bump.",
					LongDescription:  "The Durock Taro Ball switch is a tactile switch with a taro-themed design and a smooth tactile bump, offering a unique aesthetic and a satisfying typing experience. It's well-loved for its distinctive look and performance.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-12-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Durock Corgi",
					ShortDescription: "Linear switch with a corgi-themed design and buttery smoothness.",
					LongDescription:  "The Durock Corgi switch is a linear switch featuring a corgi-themed design and buttery smoothness. It's perfect for users who seek a fun aesthetic combined with a smooth typing experience.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-04-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Durock Shrimpy Click",
					ShortDescription: "Clicky switch with a vibrant design and sharp click.",
					LongDescription:  "The Durock Shrimpy Click is a vibrant clicky switch featuring a sharp click and a unique design. It offers a playful aesthetic combined with a satisfying clicky typing experience, perfect for those who enjoy bold colors and feedback.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2021-08-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Durock White Lotus",
					ShortDescription: "Clicky switch with a floral aesthetic and distinct click.",
					LongDescription:  "The Durock White Lotus switch features a floral aesthetic and a distinct click sound, offering a unique typing experience that blends visual appeal with tactile feedback. It is ideal for those who appreciate themed switches.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2021-05-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Durock Thanos Click",
					ShortDescription: "Clicky switch with a deep, thocky click and unique theme.",
					LongDescription:  "The Durock Thanos Click switch offers a deep, thocky click sound and a unique theme inspired by powerful design elements. It's designed for enthusiasts who seek a distinctive auditory experience with a powerful aesthetic.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Durock Peacock",
					ShortDescription: "Clicky switch with a colorful design and lively click.",
					LongDescription:  "The Durock Peacock switch is a clicky switch that boasts a colorful design and a lively click, making it a favorite for those who want a vibrant aesthetic combined with the tactile and auditory satisfaction of a clicky switch.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2022-04-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Durock Joyous Click",
					ShortDescription: "Clicky switch with a cheerful design and crisp click.",
					LongDescription:  "The Durock Joyous Click switch is known for its cheerful design and crisp click, offering a joyful typing experience that combines a fun aesthetic with satisfying clicky feedback. It's perfect for users who want to add a touch of joy to their setup.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2021-11-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Durock Brightside Click",
					ShortDescription: "Clicky switch with a bright design and clear click sound.",
					LongDescription:  "The Durock Brightside Click switch features a bright design with a clear click sound, making it an excellent choice for those who appreciate a vibrant aesthetic alongside satisfying tactile and auditory feedback.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2022-02-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Durock Firestorm",
					ShortDescription: "Clicky switch with a fiery design and explosive click.",
					LongDescription:  "The Durock Firestorm switch is a clicky switch that offers a fiery design and an explosive click sound. It's designed for those who want a dramatic typing experience with a bold auditory and visual statement.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2021-03-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Durock T2",
					ShortDescription: "Enhanced tactile switch with smoother bump and less noise.",
					LongDescription:  "The Durock T2 is an enhanced version of the T1, providing a smoother tactile bump and less noise during actuation. It retains the high-quality construction that Durock is known for, making it a reliable choice for tactile switch fans.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Durock Shrimp Tactile",
					ShortDescription: "Tactile switch with a unique pink housing and satisfying bump.",
					LongDescription:  "The Durock Shrimp Tactile switch features a unique pink housing and a satisfying tactile bump, providing both aesthetic appeal and excellent performance. It's designed for users who appreciate a distinctive look with great tactile feedback.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2022-05-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Durock Koala V2",
					ShortDescription: "Upgraded tactile switch with improved bump and smoothness.",
					LongDescription:  "The Durock Koala V2 is an upgraded version of the original Koala, offering an improved tactile bump and enhanced smoothness. It's designed for enthusiasts who want a premium typing experience with reliable tactile feedback.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-06-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Durock POM Tactile",
					ShortDescription: "Tactile switch with POM housing for a unique feel.",
					LongDescription:  "The Durock POM Tactile switch features a POM housing that provides a unique feel and reduced friction, offering a distinct tactile experience that's smooth and satisfying. It's perfect for those who seek something different in their tactile switches.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-03-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Durock Daybreak Tactile",
					ShortDescription: "Tactile switch with a dawn-themed design and smooth bump.",
					LongDescription:  "The Durock Daybreak Tactile switch offers a dawn-themed design with a smooth tactile bump, providing an excellent balance between aesthetics and performance. It's a great choice for those who want both style and substance.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2022-07-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Durock Owl Tactile",
					ShortDescription: "Tactile switch with an owl-themed design and sharp bump.",
					LongDescription:  "The Durock Owl Tactile switch features an owl-themed design and a sharp tactile bump, providing a precise feel that's perfect for tactile switch enthusiasts who enjoy themed aesthetics.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2020-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Durock Lilac Tactile",
					ShortDescription: "Tactile switch with a lilac color and smooth bump.",
					LongDescription:  "The Durock Lilac Tactile switch offers a lilac color and a smooth tactile bump, providing a unique aesthetic alongside a satisfying typing experience. It's ideal for those who appreciate color-themed builds.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-10-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Durock Zilent V2",
					ShortDescription: "Silent tactile switch with a soft bump and minimal noise.",
					LongDescription:  "The Durock Zilent V2 is a silent tactile switch offering a soft bump and minimal noise, perfect for users who require a quiet typing experience without sacrificing tactile feedback. It's highly regarded for its smoothness and quiet operation.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     5,      // Silent Tactile
					ReleaseDate:      parseDate("2022-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Durock Piano",
					ShortDescription: "Tactile switch with a musical theme and crisp bump.",
					LongDescription:  "The Durock Piano switch is a tactile switch with a musical theme and a crisp tactile bump. It's designed for users who want a switch with a unique aesthetic and satisfying feedback, making it a popular choice among musicians and enthusiasts alike.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-05-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Durock Lavender",
					ShortDescription: "Tactile switch with a lavender hue and smooth action.",
					LongDescription:  "The Durock Lavender switch offers a tactile bump with a lavender hue, providing a smooth typing experience with an appealing color scheme. It's perfect for those who want a stylish and satisfying tactile switch.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2022-06-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Durock Pudding",
					ShortDescription: "Tactile switch with a unique feel and playful design.",
					LongDescription:  "The Durock Pudding switch is a tactile switch that offers a unique feel with a playful design. It provides a satisfying tactile bump and is ideal for those who want a switch that stands out both in performance and appearance.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-08-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Durock Camel",
					ShortDescription: "Tactile switch with a sand-colored housing and smooth bump.",
					LongDescription:  "The Durock Camel switch is a tactile switch featuring a sand-colored housing and a smooth tactile bump. It's designed for users who appreciate a desert-inspired aesthetic alongside a satisfying typing experience.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2022-03-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Durock Azure",
					ShortDescription: "Tactile switch with an azure-themed design and crisp bump.",
					LongDescription:  "The Durock Azure switch offers a tactile bump with an azure-themed design, providing a crisp typing experience with a unique visual appeal. It's perfect for those who enjoy themed switches that deliver both style and performance.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-12-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Durock Coral Tactile",
					ShortDescription: "Tactile switch with a coral-themed design and strong bump.",
					LongDescription:  "The Durock Coral Tactile switch features a coral-themed design and a strong tactile bump, offering a distinct aesthetic and satisfying feedback. It's ideal for those who want a bold tactile switch with a unique theme.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2022-05-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Durock L1",
					ShortDescription: "Linear switch with smooth operation and minimal noise.",
					LongDescription:  "The Durock L1 is a linear switch known for its exceptionally smooth operation and minimal noise, making it an excellent choice for users seeking a quiet and fluid typing experience. It's favored for its consistency and low spring noise.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-05-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Durock Alpaca V1",
					ShortDescription: "Linear switch with a smooth feel and vibrant pink stem.",
					LongDescription:  "The Durock Alpaca V1 switch is a linear switch known for its smooth feel and vibrant pink stem. It offers a consistent typing experience with minimal spring noise, making it a favorite for enthusiasts seeking reliable performance.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2019-09-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Durock Cobalt",
					ShortDescription: "Linear switch with a cobalt blue stem and smooth action.",
					LongDescription:  "The Durock Cobalt switch features a cobalt blue stem and provides a smooth linear action. It's designed for users who appreciate a deep color aesthetic combined with a consistent typing feel.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-04-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Durock Sunflower Linear",
					ShortDescription: "Linear switch with a bright yellow stem and buttery smoothness.",
					LongDescription:  "The Durock Sunflower Linear switch is known for its bright yellow stem and buttery smoothness. It provides a quiet and fluid keystroke, ideal for those who want both a vibrant aesthetic and high performance.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-09-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Durock Ghost",
					ShortDescription: "Linear switch with a translucent housing and smooth keystroke.",
					LongDescription:  "The Durock Ghost switch features a translucent housing and offers a smooth linear keystroke. It's designed for users who enjoy a clean aesthetic with the performance of a premium linear switch.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-10-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Durock Dolphin V2",
					ShortDescription: "Linear switch with improved smoothness and minimal spring noise.",
					LongDescription:  "The Durock Dolphin V2 switch is an upgraded version that offers improved smoothness and minimal spring noise. Its consistent performance and pleasing aesthetics make it a popular choice for linear switch enthusiasts.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-03-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Durock Haze",
					ShortDescription: "Linear switch with a hazy design and buttery keystroke.",
					LongDescription:  "The Durock Haze switch is known for its hazy design and buttery keystroke, providing a smooth typing experience with a unique aesthetic. It's perfect for users who want a switch that combines style with performance.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2022-05-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Durock Mauve",
					ShortDescription: "Linear switch with a mauve-colored stem and smooth action.",
					LongDescription:  "The Durock Mauve switch offers a mauve-colored stem and a smooth linear action, making it a top choice for enthusiasts who prefer a unique color scheme and consistent performance.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-08-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Durock Midnight",
					ShortDescription: "Linear switch with a dark design and quiet operation.",
					LongDescription:  "The Durock Midnight switch is a linear switch characterized by its dark design and quiet operation, offering a smooth and subtle typing experience that is perfect for nighttime use.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-12-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Durock Lightwave",
					ShortDescription: "Linear switch with a light actuation force and smooth glide.",
					LongDescription:  "The Durock Lightwave switch provides a light actuation force and a smooth glide, making it ideal for users who prefer quick keystrokes with minimal resistance.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2022-02-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Durock Euphoria",
					ShortDescription: "Linear switch with a euphoric feel and consistent action.",
					LongDescription:  "The Durock Euphoria switch is designed to provide a euphoric typing experience with a consistent linear action. It's perfect for those who want a smooth and satisfying keystroke with minimal noise.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-06-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Durock Vanilla",
					ShortDescription: "Linear switch with a creamy smooth action and neutral design.",
					LongDescription:  "The Durock Vanilla switch offers a creamy smooth action and a neutral design, making it a versatile choice for any keyboard setup. It's known for its reliability and minimal spring noise.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-11-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Durock Jet",
					ShortDescription: "Linear switch with a jet-black design and fast actuation.",
					LongDescription:  "The Durock Jet switch is known for its jet-black design and fast actuation, making it perfect for gamers who require quick and responsive keystrokes. It's designed for those who need performance and speed.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2022-04-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Durock Blossom",
					ShortDescription: "Linear switch with a floral design and smooth keystroke.",
					LongDescription:  "The Durock Blossom switch features a floral design and provides a smooth linear keystroke. It's perfect for those who appreciate a delicate aesthetic with the performance of a premium linear switch.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-07-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Durock Ink Black V2",
					ShortDescription: "Linear switch with a unique ink-colored housing and buttery smoothness.",
					LongDescription:  "The Durock Ink Black V2 switch offers a unique ink-colored housing and buttery smoothness, providing a premium typing experience with a distinctive look. It's ideal for users who want high performance and a unique design.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-03-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Durock Cotton Candy",
					ShortDescription: "Linear switch with a sweet aesthetic and soft keystroke.",
					LongDescription:  "The Durock Cotton Candy switch is a linear switch known for its sweet aesthetic and soft keystroke, offering a playful look with smooth performance. It's perfect for users who want a switch that adds a fun element to their keyboard.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-02-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Durock Silent Linear",
					ShortDescription: "Silent linear switch with smooth keystroke and noise reduction.",
					LongDescription:  "The Durock Silent Linear switch is designed for those who need a quiet typing experience without sacrificing performance. It offers a smooth linear keystroke with built-in dampening to reduce noise, making it ideal for office or shared spaces.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2020-11-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Durock Silent T1",
					ShortDescription: "Silent tactile switch with a refined bump and minimal noise.",
					LongDescription:  "The Durock Silent T1 switch offers a tactile typing experience with a refined bump and minimal noise. It is perfect for users who want tactile feedback without the accompanying sound, suitable for quiet environments.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     5,      // Silent Tactile
					ReleaseDate:      parseDate("2021-06-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Durock Thocky",
					ShortDescription: "Specialty linear switch with a deep, thocky sound profile.",
					LongDescription:  "The Durock Thocky switch is a specialty linear switch known for its deep, thocky sound profile. It offers a unique auditory experience with smooth keystrokes, making it a favorite among enthusiasts who enjoy distinct sound feedback.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Durock U4 Silent",
					ShortDescription: "Silent tactile switch with soft bump and ultra-quiet operation.",
					LongDescription:  "The Durock U4 Silent switch is designed to provide a soft tactile bump with ultra-quiet operation, making it ideal for environments where noise reduction is critical. It offers smooth, dampened keystrokes with reliable tactile feedback.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     5,      // Silent Tactile
					ReleaseDate:      parseDate("2020-09-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Durock Light Fog",
					ShortDescription: "Specialty switch with a translucent housing and mysterious feel.",
					LongDescription:  "The Durock Light Fog switch offers a specialty design with a translucent housing that creates a mysterious aesthetic. It provides a smooth keystroke and is perfect for those who appreciate both performance and style.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-03-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Durock Shadow",
					ShortDescription: "Specialty switch with a dark design and smooth performance.",
					LongDescription:  "The Durock Shadow switch is a specialty switch featuring a dark design and smooth performance. It offers a quiet typing experience and is ideal for those who want a switch with a stealthy appearance and reliable action.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-08-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Durock Marshmallow Silent",
					ShortDescription: "Silent linear switch with a soft, marshmallow-like feel.",
					LongDescription:  "The Durock Marshmallow Silent switch combines a soft, marshmallow-like feel with silent operation, offering a unique typing experience that's both quiet and satisfying. It's perfect for those who want a silent switch with a distinctive touch.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2022-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Durock L4",
					ShortDescription: "Specialty linear switch with enhanced smoothness and reduced scratchiness.",
					LongDescription:  "The Durock L4 switch is a specialty linear switch offering enhanced smoothness and reduced scratchiness. It's designed for users who seek a refined typing experience with consistent linear action.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2022-04-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Durock Daydream",
					ShortDescription: "Specialty switch with a dream-themed design and smooth keystroke.",
					LongDescription:  "The Durock Daydream switch offers a dream-themed design with a smooth keystroke, providing a unique aesthetic alongside reliable performance. It's perfect for those who want a switch that captures the imagination while delivering excellent typing feedback.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Durock Aqua King",
					ShortDescription: "Linear switch with a clear housing and smooth, fluid action.",
					LongDescription:  "The Durock Aqua King switch features a clear housing that highlights its smooth, fluid linear action. It's designed for users who want a transparent aesthetic combined with top-tier linear performance.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-02-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Durock Ultra",
					ShortDescription: "Specialty switch with an ultra-smooth linear feel and minimal noise.",
					LongDescription:  "The Durock Ultra switch offers an ultra-smooth linear feel with minimal noise, providing a premium typing experience that's both quiet and satisfying. It's ideal for those who demand the highest level of performance from their switches.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2022-07-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Durock Crystal Glacier",
					ShortDescription: "Specialty switch with a crystal-clear housing and icy smoothness.",
					LongDescription:  "The Durock Crystal Glacier switch is a specialty switch with a crystal-clear housing and offers icy smoothness in its linear action. It's designed for those who want a striking aesthetic with seamless performance.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-10-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Durock Emerald",
					ShortDescription: "Linear switch with an emerald hue and luxurious smoothness.",
					LongDescription:  "The Durock Emerald switch features an emerald hue and luxurious smoothness, providing a premium typing experience with a touch of elegance. It's perfect for those who want both performance and style in their keyboard.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2022-09-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Durock Mystic",
					ShortDescription: "Specialty switch with a mystical design and smooth action.",
					LongDescription:  "The Durock Mystic switch offers a mystical design combined with smooth linear action, making it a favorite for users who want a switch with an enchanting appearance and reliable performance.",
					ManufacturerID:   ptr(7), // Durock
					BrandID:          ptr(7), // Durock
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-09-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
			}

			err = processSwitches(durockSwitches)
			if err != nil {
				return err
			}

			return nil // Replace with actual code
		},
		Rollback: func(tx *gorm.DB) error {
			var producersIDs []int
			type Producer struct {
				ID int `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
			}

			if err := tx.Model(&Producer{}).Where("alias IN (?)", []string{"ttc", "durock"}).Pluck("id", &producersIDs).Error; err != nil {
				return err
			}

			if err := tx.Exec("DELETE FROM switches WHERE brand_id IN (?)", producersIDs).Error; err != nil {
				return err
			}
			return nil // Replace with actual code
		},
	})
}
