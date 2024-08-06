package migrations

import (
	"time"

	"github.com/Bparsons0904/deadigations"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	deadigations.RegisterMigration(deadigations.Migration{
		ID:          "20240806071740",
		Description: "Add initial data for kailh switches",
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

			ptr := func(i int) *int {
				return &i
			}

			parseDate := func(date string) *time.Time {
				t, _ := time.Parse("2006-01-02", date)
				return &t
			}

			processSwitches := func(switches []Switch) error {
				for _, s := range switches {
					if err := tx.Create(&s).Error; err != nil {
						return err
					}

					time.Sleep(5 * time.Millisecond)
				}
				return nil
			}

			kailhSwitches := []Switch{
				{
					Name:             "Kailh BOX White",
					ShortDescription: "Clicky switch with a crisp and precise click sound.",
					LongDescription:  "Kailh BOX White switches feature a clicky design with a crisp and precise click sound, making them ideal for typists who enjoy audible feedback.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2017-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh BOX Navy",
					ShortDescription: "Heavy clicky switch with a loud and tactile click.",
					LongDescription:  "Kailh BOX Navy switches are designed for those who prefer a heavy click with a loud and tactile feedback, providing a satisfying typing experience.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2017-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh BOX Jade",
					ShortDescription: "Tactile clicky switch with a medium actuation force.",
					LongDescription:  "Kailh BOX Jade switches are tactile clicky switches with a medium actuation force, offering a balance between tactile feedback and audible click.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2018-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh Pro Light Green",
					ShortDescription: "Light clicky switch with smooth actuation.",
					LongDescription:  "Kailh Pro Light Green switches offer a light clicky feel with smooth actuation, providing a comfortable typing experience for long sessions.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2018-06-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				{
					Name:             "Kailh Clicky Purple",
					ShortDescription: "Clicky switch with a distinct sound profile.",
					LongDescription:  "Kailh Clicky Purple switches provide a unique clicky sound profile, offering a satisfying auditory experience for typists.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2019-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh BOX Pink",
					ShortDescription: "Light clicky switch with a gentle actuation force.",
					LongDescription:  "Kailh BOX Pink switches offer a light clicky feel with a gentle actuation force, perfect for those who prefer a lighter touch.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh Speed Bronze",
					ShortDescription: "Fast clicky switch with rapid actuation.",
					LongDescription:  "Kailh Speed Bronze switches are designed for speed with a fast clicky actuation, making them ideal for gamers who require quick response times.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2016-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh BOX Royal",
					ShortDescription: "Tactile tactile switch with a royal feel.",
					LongDescription:  "Kailh BOX Royal switches offer a tactile clicky experience with a regal feel, providing both tactile feedback and an audible click.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2017-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh BOX Red",
					ShortDescription: "Smooth linear switch with a light actuation force.",
					LongDescription:  "Kailh BOX Red switches offer a smooth linear feel with a light actuation force, making them ideal for gamers who prefer rapid key presses without tactile feedback.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2017-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh BOX Black",
					ShortDescription: "Linear switch with a heavier actuation force.",
					LongDescription:  "Kailh BOX Black switches provide a linear feel with a heavier actuation force, suitable for those who prefer a more substantial key press.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2017-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh Speed Silver",
					ShortDescription: "Fast linear switch with a short actuation distance.",
					LongDescription:  "Kailh Speed Silver switches are designed for speed, featuring a fast linear actuation with a short travel distance, perfect for competitive gaming.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2016-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh Pro Burgundy",
					ShortDescription: "Linear switch with a smooth and consistent feel.",
					LongDescription:  "Kailh Pro Burgundy switches offer a linear typing experience with a smooth and consistent feel, ideal for both typing and gaming.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2018-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh BOX Silent Pink",
					ShortDescription: "Silent linear switch with a quiet operation.",
					LongDescription:  "Kailh BOX Silent Pink switches provide a linear feel with silent operation, making them suitable for quiet environments like offices.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2019-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh BOX Red Pro",
					ShortDescription: "Enhanced linear switch with smooth actuation.",
					LongDescription:  "Kailh BOX Red Pro switches offer an enhanced linear feel with smooth actuation, providing an upgraded experience for gamers and typists alike.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh BOX Crystal Pink",
					ShortDescription: "Linear switch with a unique transparent housing.",
					LongDescription:  "Kailh BOX Crystal Pink switches feature a linear feel with a unique transparent housing, providing a smooth typing experience with a distinctive aesthetic.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh Cream",
					ShortDescription: "Smooth linear switch with a creamy feel.",
					LongDescription:  "Kailh Cream switches are known for their smooth linear feel with a creamy actuation, providing a premium typing experience.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2018-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh BOX Brown",
					ShortDescription: "Tactile switch with a light bump and smooth actuation.",
					LongDescription:  "Kailh BOX Brown switches offer a tactile feel with a light bump and smooth actuation, providing a balanced experience for both typing and gaming.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2017-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh Pro Purple",
					ShortDescription: "Tactile switch with a medium actuation force and short travel.",
					LongDescription:  "Kailh Pro Purple switches feature a tactile bump with a medium actuation force and shorter travel distance, making them ideal for fast typists.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2018-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh Speed Copper",
					ShortDescription: "Fast tactile switch with a light bump.",
					LongDescription:  "Kailh Speed Copper switches feature a fast actuation with a tactile bump, making them ideal for gaming with quick response times.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2016-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh Burnt Orange",
					ShortDescription: "Tactile switch with a light actuation force.",
					LongDescription:  "Kailh Burnt Orange switches offer a tactile bump with a light actuation force, providing a smooth and comfortable typing experience.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2017-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh BOX Pale Blue",
					ShortDescription: "Tactile switch with a medium bump and crisp feedback.",
					LongDescription:  "Kailh BOX Pale Blue switches feature a tactile bump with crisp feedback, providing a balanced tactile experience for both typing and gaming.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2017-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh BOX Heavy Burnt Orange",
					ShortDescription: "Heavy tactile switch with a pronounced bump.",
					LongDescription:  "Kailh BOX Heavy Burnt Orange switches offer a strong tactile bump with a heavier actuation force, providing a robust typing experience.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2019-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh Polia",
					ShortDescription: "Tactile switch with a smooth actuation and distinct bump.",
					LongDescription:  "Kailh Polia switches offer a smooth tactile actuation with a distinct bump, providing a satisfying feedback for typists.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh Silent Brown",
					ShortDescription: "Silent tactile switch with a gentle bump.",
					LongDescription:  "Kailh Silent Brown switches offer a tactile bump with silent operation, making them suitable for quiet environments while maintaining tactile feedback.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     5,      // Silent Tactile
					ReleaseDate:      parseDate("2018-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh BOX Dark Yellow",
					ShortDescription: "Clicky switch with a unique dark yellow stem and a heavy click.",
					LongDescription:  "Kailh BOX Dark Yellow switches are designed with a unique dark yellow stem, providing a heavy click and robust tactile feedback, ideal for users who prefer a more substantial key press.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2018-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh BOX Heavy Pale Blue",
					ShortDescription: "Heavy clicky switch with a pronounced click and robust feel.",
					LongDescription:  "Kailh BOX Heavy Pale Blue switches offer a pronounced click and robust feel, catering to typists who enjoy a strong tactile and audible response.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2018-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh BOX Heavy Pink",
					ShortDescription: "Clicky switch with a heavy actuation force and strong tactile feedback.",
					LongDescription:  "Kailh BOX Heavy Pink switches provide a heavy actuation force with strong tactile feedback, making them a favorite among those who prefer a pronounced typing experience.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2018-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh Click Bar Jade",
					ShortDescription: "Clicky switch with a bar click mechanism and tactile feedback.",
					LongDescription:  "Kailh Click Bar Jade switches feature a unique bar click mechanism that delivers satisfying tactile feedback, offering a distinct sound profile and typing experience.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2019-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh BOX Red Silent",
					ShortDescription: "Silent linear switch with smooth actuation and dampened sound.",
					LongDescription:  "Kailh BOX Red Silent switches offer a silent linear experience with smooth actuation, perfect for quiet environments and those who prefer minimal noise.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2019-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh BOX Cream Pro",
					ShortDescription: "Linear switch with a creamy smooth feel and extended durability.",
					LongDescription:  "Kailh BOX Cream Pro switches are renowned for their creamy smooth feel and extended durability, offering a premium typing experience for enthusiasts.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-06-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh BOX Red V2",
					ShortDescription: "Updated linear switch with improved smoothness and consistency.",
					LongDescription:  "Kailh BOX Red V2 switches provide an updated linear experience with enhanced smoothness and consistency, suitable for gamers and typists alike.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh Smooth Linear",
					ShortDescription: "Ultra-smooth linear switch with a light actuation force.",
					LongDescription:  "Kailh Smooth Linear switches deliver an ultra-smooth feel with a light actuation force, perfect for those seeking a seamless typing experience.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh Chocolate",
					ShortDescription: "Linear switch with a unique chocolate-colored housing and smooth feel.",
					LongDescription:  "Kailh Chocolate switches offer a distinctive chocolate-colored housing and a smooth linear feel, providing both aesthetics and performance.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh Speed Pink",
					ShortDescription: "Fast tactile switch with a light bump and rapid actuation.",
					LongDescription:  "Kailh Speed Pink switches feature a fast tactile actuation with a light bump, designed for users who require quick response times.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh BOX Royal V2",
					ShortDescription: "Updated tactile switch with improved bump and feedback.",
					LongDescription:  "Kailh BOX Royal V2 switches offer an enhanced tactile experience with improved bump and feedback, catering to typists who enjoy pronounced tactile response.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh BOX Silent Pale Blue",
					ShortDescription: "Silent tactile switch with gentle bump and quiet operation.",
					LongDescription:  "Kailh BOX Silent Pale Blue switches provide a tactile bump with silent operation, making them suitable for quiet environments while maintaining tactile feedback.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     5,      // Silent Tactile
					ReleaseDate:      parseDate("2019-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh Silent Heavy Orange",
					ShortDescription: "Silent tactile switch with a heavy actuation force and strong feedback.",
					LongDescription:  "Kailh Silent Heavy Orange switches offer a heavy actuation force with strong tactile feedback and silent operation, ideal for users who require a robust yet quiet typing experience.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     5,      // Silent Tactile
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh BOX Olive",
					ShortDescription: "Tactile switch with a unique olive stem and smooth bump.",
					LongDescription:  "Kailh BOX Olive switches feature a unique olive stem and smooth tactile bump, offering a blend of aesthetics and performance for discerning typists.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh BOX White V2",
					ShortDescription: "Enhanced clicky switch with a crisp click and better durability.",
					LongDescription:  "The Kailh BOX White V2 switch offers an improved clicky experience with a crisp click and enhanced durability, suitable for typists who enjoy audible feedback.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh BOX Brown V2",
					ShortDescription: "Improved tactile switch with a pronounced bump and smooth actuation.",
					LongDescription:  "Kailh BOX Brown V2 switches feature an enhanced tactile bump with smooth actuation, providing a satisfying typing experience for users who enjoy tactile feedback.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh BOX Black V2",
					ShortDescription: "Updated linear switch with a heavier actuation force and improved durability.",
					LongDescription:  "The Kailh BOX Black V2 switch offers a linear feel with a heavier actuation force, providing improved durability and smoothness, perfect for users who prefer a substantial key press.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh BOX Jade V2",
					ShortDescription: "Tactile clicky switch with enhanced click and smoothness.",
					LongDescription:  "The Kailh BOX Jade V2 switch offers a tactile clicky experience with an enhanced click and smoothness, ideal for typists who enjoy tactile feedback with audible sound.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh BOX Cream Milk",
					ShortDescription: "Smooth linear switch with a unique milky housing and gentle feel.",
					LongDescription:  "Kailh BOX Cream Milk switches provide a smooth linear feel with a unique milky housing, offering a gentle typing experience with a distinctive aesthetic.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-06-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh BOX Cream Soda",
					ShortDescription: "Linear switch with a bubbly smoothness and subtle aesthetics.",
					LongDescription:  "Kailh BOX Cream Soda switches offer a bubbly smoothness and subtle aesthetics, providing a delightful typing experience with a gentle linear feel.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-06-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh BOX Snow",
					ShortDescription: "Tactile switch with a gentle bump and frosty design.",
					LongDescription:  "Kailh BOX Snow switches feature a gentle tactile bump with a frosty design, offering a unique winter-themed aesthetic and satisfying tactile feedback.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-12-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh BOX Ice",
					ShortDescription: "Tactile switch with a crisp bump and icy appearance.",
					LongDescription:  "The Kailh BOX Ice switch offers a crisp tactile bump and icy appearance, providing a cool winter-themed design with satisfying tactile feedback.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-12-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh BOX Glacier",
					ShortDescription: "Tactile switch with a pronounced bump and cool aesthetics.",
					LongDescription:  "Kailh BOX Glacier switches offer a pronounced tactile bump and cool aesthetics, providing a satisfying typing experience with a unique winter-themed design.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-12-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh Crystal Burgundy",
					ShortDescription: "Linear switch with smooth actuation and crystal-clear housing.",
					LongDescription:  "The Kailh Crystal Burgundy switch offers a smooth linear actuation with crystal-clear housing, providing an aesthetically pleasing and responsive typing experience.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh Crystal Sapphire",
					ShortDescription: "Linear switch with a sapphire-tinted housing and smooth feel.",
					LongDescription:  "Kailh Crystal Sapphire switches provide a smooth linear feel with a sapphire-tinted housing, offering both performance and visual appeal.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh Crystal Ruby",
					ShortDescription: "Linear switch with a ruby-red tint and enhanced smoothness.",
					LongDescription:  "Kailh Crystal Ruby switches offer enhanced smoothness and a ruby-red tint, delivering a premium linear experience with distinctive aesthetics.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh Midnight Pro Silent Red",
					ShortDescription: "Silent linear switch with smooth actuation and minimal noise.",
					LongDescription:  "The Kailh Midnight Pro Silent Red switch offers a silent linear experience with smooth actuation, perfect for quiet environments and those who prefer minimal noise.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2022-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh Midnight Pro Silent Brown",
					ShortDescription: "Silent tactile switch with gentle feedback and quiet operation.",
					LongDescription:  "Kailh Midnight Pro Silent Brown switches provide a gentle tactile feedback with quiet operation, offering a satisfying typing experience in noise-sensitive settings.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     5,      // Silent Tactile
					ReleaseDate:      parseDate("2022-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh Midnight Pro Silent Black",
					ShortDescription: "Silent linear switch with a heavier actuation force and quiet operation.",
					LongDescription:  "Kailh Midnight Pro Silent Black switches offer a silent linear experience with a heavier actuation force, providing smooth operation with minimal noise.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2022-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh Clione Limacina",
					ShortDescription: "Ultra-smooth linear switch inspired by sea angel aesthetics.",
					LongDescription:  "The Kailh Clione Limacina switch delivers an ultra-smooth linear feel inspired by the delicate beauty of sea angels, offering a unique and refined typing experience.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2023-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh Clione Aqua",
					ShortDescription: "Linear switch with a unique aquatic design and smooth actuation.",
					LongDescription:  "Kailh Clione Aqua switches offer a linear experience with a unique aquatic design, providing smooth actuation and an eye-catching appearance.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2023-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh Clione Ocean",
					ShortDescription: "Linear switch with oceanic themes and buttery smooth feel.",
					LongDescription:  "Kailh Clione Ocean switches feature oceanic themes and a buttery smooth feel, offering a premium linear experience with a touch of marine elegance.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2023-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh Super Speed Silver",
					ShortDescription: "Fast linear switch with ultra-short actuation distance.",
					LongDescription:  "The Kailh Super Speed Silver switch is designed for speed, featuring an ultra-short actuation distance and rapid response, ideal for competitive gaming and fast typists.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2023-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh Super Speed Copper",
					ShortDescription: "Fast switch with rapid actuation and tactile feedback.",
					LongDescription:  "Kailh Super Speed Copper switches provide rapid actuation with tactile feedback, making them perfect for users who require quick response times and enjoy audible feedback.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2023-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh Super Speed Gold",
					ShortDescription: "Linear switch with gold accents and lightning-fast response.",
					LongDescription:  "Kailh Super Speed Gold switches offer a linear experience with gold accents and a lightning-fast response, catering to gamers and typists who demand speed.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2023-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh Deep Sea Silent Pro",
					ShortDescription: "Silent tactile switch with deep tactile feedback and minimal noise.",
					LongDescription:  "The Kailh Deep Sea Silent Pro switch offers silent tactile feedback with minimal noise, providing a deep and satisfying typing experience for those who prefer a quiet environment.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     5,      // Silent Tactile
					ReleaseDate:      parseDate("2022-12-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh Jellyfish Box Linear (X)",
					ShortDescription: "Linear switch with smooth actuation inspired by jellyfish.",
					LongDescription:  "The Kailh Jellyfish Box Linear (X) switch offers smooth actuation with an ethereal design, inspired by the fluid motion of jellyfish.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2023-03-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh Jellyfish Box Clicky (Y)",
					ShortDescription: "Clicky switch with a distinct sound profile.",
					LongDescription:  "Kailh Jellyfish Box Clicky (Y) switches provide a clicky experience with a distinct sound profile, capturing the essence of jellyfish in motion.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2023-03-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh White Owl",
					ShortDescription: "Tactile switch with gentle feedback and aesthetic design.",
					LongDescription:  "The Kailh White Owl switch offers a tactile experience with gentle feedback, known for its aesthetic design and smooth operation.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2023-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh Red Bean Pudding Box",
					ShortDescription: "Linear switch with a pudding-like texture and smooth feel.",
					LongDescription:  "The Kailh Red Bean Pudding Box switch offers a smooth linear experience with a pudding-like texture, providing a unique feel and visual appeal.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2023-05-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh Coco Pink Box V2",
					ShortDescription: "Linear switch with vibrant pink color and smooth operation.",
					LongDescription:  "Kailh Coco Pink Box V2 switches offer a vibrant pink color and smooth linear operation, combining style with performance.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2023-02-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh Box Cream",
					ShortDescription: "Linear switch with creamy smooth feel and high durability.",
					LongDescription:  "The Kailh Box Cream switch offers a creamy smooth linear experience with high durability, making it a favorite among enthusiasts for its performance and longevity.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh Box Cream Pro",
					ShortDescription: "Pro version with enhanced smoothness and extended lifespan.",
					LongDescription:  "Kailh Box Cream Pro switches deliver an enhanced smooth linear experience with an extended lifespan, providing a premium typing experience for discerning users.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-06-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh Box Summer Clicky",
					ShortDescription: "Clicky switch designed for a refreshing and lively typing experience.",
					LongDescription:  "The Kailh Box Summer Clicky switch provides a refreshing and lively typing experience with vibrant design and crisp click feedback.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2022-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh Box Winter",
					ShortDescription: "Tactile switch with a winter-themed design and unique feel.",
					LongDescription:  "Kailh Box Winter switches offer a tactile experience with a winter-themed design, providing a unique feel for those who appreciate themed aesthetics.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-12-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh Box Brown Silent",
					ShortDescription: "Silent tactile switch with gentle feedback, designed for quiet typing environments.",
					LongDescription:  "The Kailh Box Brown Silent switch provides a quiet tactile experience with gentle feedback, perfect for noise-sensitive environments and typists who prefer subtle tactile response.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     5,      // Silent Tactile
					ReleaseDate:      parseDate("2019-03-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh Box Red Silent",
					ShortDescription: "Silent linear switch with smooth actuation and minimal noise.",
					LongDescription:  "Kailh Box Red Silent switches offer a smooth linear experience with minimal noise, providing a quiet typing environment without sacrificing performance.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2019-03-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh Box Navy V2",
					ShortDescription: "Updated heavy clicky switch with loud and tactile feedback.",
					LongDescription:  "The Kailh Box Navy V2 switch offers an updated heavy clicky experience with loud and tactile feedback, designed for users who enjoy pronounced sound and feel.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2021-06-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh Box Pink V2",
					ShortDescription: "Light clicky switch with gentle actuation and vibrant aesthetics.",
					LongDescription:  "Kailh Box Pink V2 switches provide a light clicky experience with gentle actuation force, offering vibrant aesthetics and satisfying feedback.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2021-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh Box Heavy Black",
					ShortDescription: "Heavy linear switch with increased actuation force for a substantial feel.",
					LongDescription:  "The Kailh Box Heavy Black switch offers a heavy linear experience with increased actuation force, catering to users who prefer a more substantial and deliberate typing feel.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2018-06-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh Box Heavy Yellow",
					ShortDescription: "Heavy linear switch with a substantial actuation force and smooth operation.",
					LongDescription:  "The Kailh Box Heavy Yellow switch is designed for users who prefer a heavier key press. It offers a smooth linear actuation with minimal tactile feedback, making it ideal for those seeking a deliberate typing experience.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2018-06-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh Pro Brown",
					ShortDescription: "Tactile switch with a medium bump and short actuation.",
					LongDescription:  "The Kailh Pro Brown switch offers a tactile experience with a medium bump and a short actuation distance, ideal for fast typing and gaming.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2018-10-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh Pro Silent Red",
					ShortDescription: "Silent linear switch with smooth actuation and minimal noise.",
					LongDescription:  "Kailh Pro Silent Red switches offer a smooth linear experience with minimal noise, making them perfect for quiet environments without sacrificing performance.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2019-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh Crystal Royal",
					ShortDescription: "Tactile switch with a pronounced bump and clear housing for RGB.",
					LongDescription:  "The Kailh Crystal Royal switch provides a tactile experience with a pronounced bump and clear housing, perfect for showcasing RGB lighting and offering a satisfying tactile feedback.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-06-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh Crystal Clear",
					ShortDescription: "Linear switch with ultra-clear housing for maximum RGB visibility.",
					LongDescription:  "Kailh Crystal Clear switches feature ultra-clear housing for maximum RGB visibility and provide a smooth linear actuation, making them a perfect choice for aesthetics and performance.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-06-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh Speed Purple",
					ShortDescription: "Tactile switch with a short actuation distance and light tactile bump.",
					LongDescription:  "The Kailh Speed Purple switch provides a tactile experience with a short actuation distance and a light tactile bump, perfect for fast typing and gaming that requires tactile feedback.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2017-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh Super Speed Bronze",
					ShortDescription: "Tactile switch with a short actuation distance and tactile bump.",
					LongDescription:  "Kailh Super Speed Bronze switches offer a tactile experience with a short actuation distance and a pronounced tactile bump, providing a balance of speed and feedback.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2023-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh Speed Pro Heavy Green",
					ShortDescription: "Tactile switch with a heavier actuation force and tactile bump.",
					LongDescription:  "The Kailh Speed Pro Heavy Green switch combines a heavier actuation force with a tactile bump, offering a substantial typing experience for those who prefer a more deliberate feel.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2018-10-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh Box Noble Yellow",
					ShortDescription: "Smooth linear switch with a light actuation force.",
					LongDescription:  "Kailh Box Noble Yellow switches offer a smooth linear experience with a light actuation force, providing a comfortable typing experience for extended use.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2019-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh Box Mute Jade",
					ShortDescription: "Quieter version of the Box Jade with tactile feedback.",
					LongDescription:  "Kailh Box Mute Jade switches provide a tactile experience similar to the Box Jade but with reduced noise, suitable for environments where quieter operation is preferred.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh Canary",
					ShortDescription: "Linear switch with a unique color and smooth operation.",
					LongDescription:  "The Kailh Canary switch offers a linear experience with a distinctive yellow color and smooth operation, providing both aesthetic appeal and performance.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh Box Ancient Gray",
					ShortDescription: "Heavy linear switch with increased actuation force.",
					LongDescription:  "Kailh Box Ancient Gray switches provide a heavy linear experience with increased actuation force, catering to users who prefer a more substantial key press.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2019-06-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh Box Pro Green",
					ShortDescription: "Clicky switch with a unique feel and feedback.",
					LongDescription:  "The Kailh Box Pro Green switch offers a clicky experience with a unique feel, different from regular Box switches, providing satisfying tactile and audible feedback.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2019-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh Sage",
					ShortDescription: "Tactile switch with a unique bump and feedback.",
					LongDescription:  "Kailh Sage switches offer a tactile experience with a unique bump and feedback, providing a distinctive feel for enthusiasts who appreciate novel tactile sensations.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kailh Speed Gold",
					ShortDescription: "Fast clicky switch designed for quick actuation.",
					LongDescription:  "The Kailh Speed Gold switch is a clicky switch designed for fast actuation, ideal for gamers and fast typists who enjoy audible feedback.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2017-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh Box Glazed Green",
					ShortDescription: "Limited edition tactile switch with unique aesthetics.",
					LongDescription:  "Kailh Box Glazed Green switches are a limited edition tactile switch, offering a unique aesthetic and tactile experience for keyboard enthusiasts.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
			}

			err := processSwitches(kailhSwitches)
			if err != nil {
				return err
			}

			return nil // Replace with actual code
		},
		Rollback: func(tx *gorm.DB) error {
			return nil // Replace with actual code
		},
	})
}

