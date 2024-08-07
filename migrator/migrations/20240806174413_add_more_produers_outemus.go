package migrations

import (
	"time"

	"github.com/Bparsons0904/deadigations"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	deadigations.RegisterMigration(deadigations.Migration{
		ID:          "20240806174413",
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

			outemus := []Switch{
				{
					Name:             "Outemu Blue",
					ShortDescription: "Clicky switch with a tactile bump and distinct sound.",
					LongDescription:  "Outemu Blue switches are known for their clicky sound and tactile bump, making them a popular choice for typists who enjoy audible feedback.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2015-01-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				{
					Name:             "Outemu Ice Blue",
					ShortDescription: "Transparent clicky switch for RGB setups.",
					LongDescription:  "Outemu Ice Blue switches feature a transparent housing, offering a clicky and tactile experience, ideal for RGB lighting setups in mechanical keyboards.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2017-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Outemu Purple",
					ShortDescription: "Clicky switch with a louder sound profile.",
					LongDescription:  "Outemu Purple switches are designed to deliver a louder click sound with a strong tactile bump, appealing to users who enjoy pronounced feedback.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2016-03-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				{
					Name:             "Outemu Black Dustproof",
					ShortDescription: "Clicky switch with dustproof features.",
					LongDescription:  "Outemu Black Dustproof switches are clicky switches that include dustproof features, offering durability and consistent performance even in dusty environments.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2018-05-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Outemu Green",
					ShortDescription: "Heavy clicky switch with strong tactile feedback.",
					LongDescription:  "Outemu Green switches are heavy clicky switches, requiring more force to actuate and offering a strong tactile bump, perfect for those who prefer a more resistant key press.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2017-09-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				{
					Name:             "Outemu White",
					ShortDescription: "Smooth clicky switch with a clear sound.",
					LongDescription:  "Outemu White switches offer a smooth clicky experience with a clear and sharp sound, providing an enjoyable typing experience for those who appreciate audible feedback.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2019-11-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Outemu Sky Blue",
					ShortDescription: "Clicky switch with a balanced feel.",
					LongDescription:  "Outemu Sky Blue switches provide a balanced clicky experience, offering a satisfying tactile bump and moderate resistance for a comfortable typing experience.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2020-03-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Outemu Ocean Blue",
					ShortDescription: "Clicky switch with a light actuation force.",
					LongDescription:  "Outemu Ocean Blue switches are designed for users who prefer a clicky switch with a lighter actuation force, making them suitable for fast typing and gaming.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2021-07-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Outemu Polar Blue",
					ShortDescription: "Clicky switch with enhanced tactile feedback.",
					LongDescription:  "Outemu Polar Blue switches provide enhanced tactile feedback with a clicky sound, offering a pronounced typing experience for users who value strong tactile sensations.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2022-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Outemu Brown",
					ShortDescription: "Tactile switch with a subtle bump and no click.",
					LongDescription:  "Outemu Brown switches offer a smooth tactile bump without a click sound, making them versatile for both typing and gaming enthusiasts who prefer quieter switches.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2015-01-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				{
					Name:             "Outemu Ice Brown",
					ShortDescription: "Tactile switch with transparent housing.",
					LongDescription:  "Outemu Ice Brown switches provide a tactile experience with a transparent housing, ideal for those looking to incorporate RGB lighting into their setups while maintaining a tactile feel.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2017-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Outemu Purple Tactile",
					ShortDescription: "Tactile switch with a distinct bump.",
					LongDescription:  "Outemu Purple Tactile switches offer a pronounced tactile bump with no click, providing a satisfying feedback ideal for typing without the audible click sound.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2018-03-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Outemu Orange",
					ShortDescription: "Medium tactile switch with a light tactile bump.",
					LongDescription:  "Outemu Orange switches offer a medium tactile experience with a lighter bump, making them ideal for users who prefer a softer touch with tactile feedback.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2019-05-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Outemu Sky Tactile",
					ShortDescription: "Tactile switch with a clear, satisfying bump.",
					LongDescription:  "Outemu Sky Tactile switches provide a satisfying tactile bump with a clear actuation point, perfect for users who want precise feedback during typing and gaming.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2020-02-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Outemu Aqua",
					ShortDescription: "Tactile switch with a smooth bump and light actuation.",
					LongDescription:  "Outemu Aqua switches are tactile switches that offer a smooth bump with a light actuation force, ideal for those who want a tactile feel without too much resistance.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Outemu Milk Brown",
					ShortDescription: "Tactile switch with a muted sound profile.",
					LongDescription:  "Outemu Milk Brown switches provide a tactile typing experience with a muted sound profile, perfect for quiet environments where tactile feedback is desired without noise.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2022-04-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Outemu Jade",
					ShortDescription: "Heavy tactile switch with a strong bump.",
					LongDescription:  "Outemu Jade switches are heavy tactile switches, offering a strong tactile bump and resistance, making them ideal for those who enjoy a pronounced tactile feedback.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2022-08-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Outemu Black",
					ShortDescription: "Heavy linear switch with no tactile bump.",
					LongDescription:  "Outemu Black switches are heavier linear switches, requiring more force to actuate, offering a satisfying feel for those who prefer a more resistant key press.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2016-01-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				{
					Name:             "Outemu Ice Black",
					ShortDescription: "Linear switch with a transparent housing.",
					LongDescription:  "Outemu Ice Black switches offer a heavy linear feel with a transparent housing, ideal for those who want a robust typing experience with RGB lighting.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2017-07-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Outemu Yellow",
					ShortDescription: "Medium-weight linear switch with smooth keystroke.",
					LongDescription:  "Outemu Yellow switches provide a smooth linear experience with medium actuation force, suitable for both typing and gaming.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2018-02-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Outemu Silver",
					ShortDescription: "Speed linear switch with reduced travel distance.",
					LongDescription:  "Outemu Silver switches are speed linear switches designed for fast typing and gaming, featuring reduced travel distance for quicker actuation.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2019-06-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Outemu Red Dustproof",
					ShortDescription: "Dustproof linear switch with smooth actuation.",
					LongDescription:  "Outemu Red Dustproof switches are linear switches with dustproof features, offering smooth actuation while enhancing durability in dusty environments.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-04-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Outemu Pink",
					ShortDescription: "Linear switch with light actuation force.",
					LongDescription:  "Outemu Pink switches offer a linear typing experience with light actuation force, making them suitable for users who prefer minimal resistance.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-11-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Outemu Peach",
					ShortDescription: "Smooth linear switch with medium resistance.",
					LongDescription:  "Outemu Peach switches provide a smooth linear experience with medium resistance, perfect for typists who enjoy a balanced typing feel.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-03-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Outemu Milk White",
					ShortDescription: "Silent linear switch with noise-reducing features.",
					LongDescription:  "Outemu Milk White switches are silent linear switches with noise-reducing features, offering a quiet typing experience for those who value silence.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2021-08-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Outemu Cream",
					ShortDescription: "Smooth linear switch with a creamy feel.",
					LongDescription:  "Outemu Cream switches offer a linear experience with a creamy, smooth feel, making them ideal for users who enjoy a luxurious typing sensation.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2022-02-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Outemu Purple Linear",
					ShortDescription: "Linear switch with a smooth keystroke and no bump.",
					LongDescription:  "Outemu Purple Linear switches offer a smooth typing experience with no tactile bump, providing a consistent feel for fast typing and gaming.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2022-09-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Outemu Silent White",
					ShortDescription: "Silent tactile switch with noise reduction.",
					LongDescription:  "Outemu Silent White switches offer a tactile typing experience with built-in noise-reducing features, perfect for quiet environments.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     5,      // Silent Tactile
					ReleaseDate:      parseDate("2019-10-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Outemu Silent Red",
					ShortDescription: "Silent linear switch with quiet operation.",
					LongDescription:  "Outemu Silent Red switches provide a smooth linear action with noise-dampening technology, ideal for environments where silence is essential.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2020-03-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Outemu Silent Black",
					ShortDescription: "Heavy silent linear switch with quiet performance.",
					LongDescription:  "Outemu Silent Black switches are designed for those who prefer a heavier linear switch with silent operation, offering resistance without the noise.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2020-08-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Outemu Ice Purple",
					ShortDescription: "Tactile switch with enhanced feedback and RGB support.",
					LongDescription:  "Outemu Ice Purple switches provide tactile feedback with a clear housing, allowing for optimal RGB lighting effects. Perfect for users who desire both feedback and aesthetics.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2018-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Outemu Polar Ice",
					ShortDescription: "Specialty switch with a smooth linear feel.",
					LongDescription:  "Outemu Polar Ice switches offer a unique smooth linear experience with a custom stem design that enhances stability and reduces wobble, ideal for precise typing.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-02-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Outemu Fire",
					ShortDescription: "Specialty switch with fast actuation.",
					LongDescription:  "Outemu Fire switches are designed for rapid actuation, providing a quick response suitable for gaming and fast typing. Ideal for users who prioritize speed.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-05-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Outemu Shadow",
					ShortDescription: "Silent tactile switch with subtle feedback.",
					LongDescription:  "Outemu Shadow switches offer a silent tactile experience with a subtle bump, making them perfect for users who enjoy tactile feedback without noise.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     5,      // Silent Tactile
					ReleaseDate:      parseDate("2021-09-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Outemu Crystal",
					ShortDescription: "Transparent switch for RGB enthusiasts.",
					LongDescription:  "Outemu Crystal switches feature a fully transparent housing, designed to maximize RGB lighting while providing a consistent linear feel for typing and gaming.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2022-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Outemu Teal",
					ShortDescription: "Tactile switch with a unique stem design.",
					LongDescription:  "Outemu Teal switches offer a tactile experience with a specially designed stem that provides unique feedback and enhances overall stability.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2022-06-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Outemu Moonlight",
					ShortDescription: "Silent linear switch with a smooth keystroke.",
					LongDescription:  "Outemu Moonlight switches provide a silent linear experience with a smooth keystroke, offering a whisper-quiet operation for users who need silence.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2022-11-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Outemu Red",
					ShortDescription: "Light linear switch for smooth typing and gaming.",
					LongDescription:  "Outemu Red switches are light linear switches that offer a smooth keystroke without tactile feedback, making them suitable for both typing and gaming.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2015-01-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				{
					Name:             "Outemu Ice Light Purple",
					ShortDescription: "Light tactile switch with transparent housing.",
					LongDescription:  "Outemu Ice Light Purple switches offer a lighter tactile experience compared to regular Purple switches, with a transparent housing for enhanced RGB lighting effects.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2018-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Outemu Dustproof Red",
					ShortDescription: "Linear switch with dustproof design.",
					LongDescription:  "Outemu Dustproof Red switches are linear switches with a dustproof design, offering consistent performance and increased durability in various environments.",
					ManufacturerID:   ptr(4), // Outemu
					BrandID:          ptr(4), // Outemu
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2018-09-01"),
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

			err := processSwitches(outemus)
			if err != nil {
				return err
			}

			zealpcSwitches := []Switch{
				{
					Name:             "Zealios V2",
					ShortDescription: "Tactile switch with a rounded bump and smooth action.",
					LongDescription:  "Zealios V2 switches offer a smooth tactile bump with minimal resistance. They are perfect for users who prefer a gentle tactile experience with a notable but non-disruptive feedback.",
					ManufacturerID:   ptr(5), // ZealPC
					BrandID:          ptr(5), // ZealPC
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2018-11-15"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Tealios V2",
					ShortDescription: "Linear switch with a unique teal color and smooth feel.",
					LongDescription:  "Tealios V2 switches are known for their exceptionally smooth keystrokes, achieved through meticulous engineering and quality materials. They are an excellent choice for gamers and typists who prioritize smoothness and speed.",
					ManufacturerID:   ptr(5), // ZealPC
					BrandID:          ptr(5), // ZealPC
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2018-09-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Healios",
					ShortDescription: "Silent linear switch for quiet typing.",
					LongDescription:  "Healios are designed for silent operation, providing a linear feel with built-in dampening to minimize noise. They are ideal for office environments or situations where quiet operation is a priority.",
					ManufacturerID:   ptr(5), // ZealPC
					BrandID:          ptr(5), // ZealPC
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2018-07-20"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Zilents V2",
					ShortDescription: "Silent tactile switch with a smooth bump.",
					LongDescription:  "Zilents V2 offer a silent tactile experience with a smooth, rounded bump, making them suitable for quiet typing enthusiasts who still desire tactile feedback.",
					ManufacturerID:   ptr(5), // ZealPC
					BrandID:          ptr(5), // ZealPC
					SwitchTypeID:     5,      // Silent Tactile
					ReleaseDate:      parseDate("2019-02-10"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Rosélios",
					ShortDescription: "Smooth linear switch with a unique color.",
					LongDescription:  "Rosélios are crafted for those who prefer an elegant, smooth linear experience with distinctive aesthetic appeal, making them a favorite among custom keyboard builders.",
					ManufacturerID:   ptr(5), // ZealPC
					BrandID:          ptr(5), // ZealPC
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-05-18"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Zeal Clickiez",
					ShortDescription: "Unique clicky switch with customizable click sound.",
					LongDescription:  "Zeal Clickiez provide a satisfying click with customizable click elements, allowing users to tailor the sound to their preference. Perfect for those who enjoy a tactile clicky experience.",
					ManufacturerID:   ptr(5), // ZealPC
					BrandID:          ptr(5), // ZealPC
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2021-06-25"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Zealios Clicky V2",
					ShortDescription: "Clicky switch with smooth operation and crisp sound.",
					LongDescription:  "Zealios Clicky V2 switches are designed for those who love a pronounced click with each keystroke. With a smooth action and crisp sound profile, they provide a delightful clicky experience while maintaining the high build quality ZealPC is known for.",
					ManufacturerID:   ptr(5), // ZealPC as manufacturer
					BrandID:          ptr(5), // ZealPC as brand
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2019-11-20"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Zealios V1",
					ShortDescription: "Original tactile switch with a smooth and satisfying bump.",
					LongDescription:  "The Zealios V1 is ZealPC's original tactile switch, offering a smooth tactile bump with moderate resistance. It is designed for typists who appreciate a classic tactile experience, combining smoothness and precision.",
					ManufacturerID:   ptr(5), // ZealPC as manufacturer
					BrandID:          ptr(5), // ZealPC as brand
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2015-04-15"),
					Available:        false, // Discontinued
					PricePoint:       3,     // Expensive
				},
				{
					Name:             "Zealios V2.2",
					ShortDescription: "Refined tactile switch with improved bump and smoother operation.",
					LongDescription:  "The Zealios V2.2 offers an upgraded tactile experience with a more pronounced bump and enhanced smoothness, ideal for those who desire a refined tactile feel. This version improves upon its predecessors with meticulous engineering and high-quality materials.",
					ManufacturerID:   ptr(5), // ZealPC as manufacturer
					BrandID:          ptr(5), // ZealPC as brand
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2022-08-05"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Zilent V2",
					ShortDescription: "Silent tactile switch with a soft and quiet bump.",
					LongDescription:  "Zilent V2 switches provide a quiet tactile experience, designed for users who prefer a tactile feedback without the accompanying noise. They are ideal for office environments or shared spaces where noise reduction is a priority.",
					ManufacturerID:   ptr(5), // ZealPC as manufacturer
					BrandID:          ptr(5), // ZealPC as brand
					SwitchTypeID:     5,      // Silent Tactile
					ReleaseDate:      parseDate("2019-02-10"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Zilent V2.2",
					ShortDescription: "Enhanced silent tactile switch with smoother and quieter operation.",
					LongDescription:  "Zilent V2.2 switches offer an enhanced silent tactile experience with a softer, more refined bump. They are perfect for typists who require both tactile feedback and noise reduction, featuring a quieter operation than previous models.",
					ManufacturerID:   ptr(5), // ZealPC as manufacturer
					BrandID:          ptr(5), // ZealPC as brand
					SwitchTypeID:     5,      // Silent Tactile
					ReleaseDate:      parseDate("2022-12-10"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Tealios V1",
					ShortDescription: "Smooth linear switch with a teal housing.",
					LongDescription:  "Tealios V1 switches are the original linear switches from ZealPC, known for their ultra-smooth operation and distinctive teal housing. They provide a fast and effortless typing experience, making them popular among gamers and fast typists.",
					ManufacturerID:   ptr(5), // ZealPC as manufacturer
					BrandID:          ptr(5), // ZealPC as brand
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2016-10-01"),
					Available:        false, // Discontinued
					PricePoint:       3,     // Expensive
				},
				{
					Name:             "Rosélios V2",
					ShortDescription: "Upgraded linear switch with a unique aesthetic.",
					LongDescription:  "Rosélios V2 switches are an upgraded version of the original Rosélios, offering improved smoothness and a striking rose-colored housing. They cater to users who want a premium linear experience with a touch of elegance.",
					ManufacturerID:   ptr(5), // ZealPC as manufacturer
					BrandID:          ptr(5), // ZealPC as brand
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-03-15"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Zealios V2.1 (Linear)",
					ShortDescription: "Linear variant of the popular Zealios V2 switch.",
					LongDescription:  "Zealios V2.1 Linear switches provide a smooth linear action, offering the same high-quality build as the tactile version but with a focus on smooth keystrokes. They are ideal for those who prefer a linear feel in their typing or gaming experience.",
					ManufacturerID:   ptr(5), // ZealPC as manufacturer
					BrandID:          ptr(5), // ZealPC as brand
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-05-18"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Silent Zealios V1",
					ShortDescription: "Original silent tactile switch with smooth feedback.",
					LongDescription:  "Silent Zealios V1 switches are designed to provide a silent tactile experience with the same smooth feedback as the original Zealios V1. They are ideal for users who prefer quiet typing with tactile feedback, offering a soft and satisfying bump without the noise.",
					ManufacturerID:   ptr(5), // ZealPC as manufacturer
					BrandID:          ptr(5), // ZealPC as brand
					SwitchTypeID:     5,      // Silent Tactile
					ReleaseDate:      parseDate("2017-04-15"),
					Available:        false, // Discontinued
					PricePoint:       3,     // Expensive
				},
				{
					Name:             "Zilent V1",
					ShortDescription: "Quiet tactile switch with a soft and gentle bump.",
					LongDescription:  "Zilent V1 switches are known for their quiet operation and soft tactile bump, making them suitable for noise-sensitive environments. They offer a gentle tactile feel without the noise, perfect for offices and shared spaces.",
					ManufacturerID:   ptr(5), // ZealPC as manufacturer
					BrandID:          ptr(5), // ZealPC as brand
					SwitchTypeID:     5,      // Silent Tactile
					ReleaseDate:      parseDate("2018-02-10"),
					Available:        false, // Discontinued
					PricePoint:       3,     // Expensive
				},
				{
					Name:             "Sakurios",
					ShortDescription: "Limited edition linear switch with a sakura pink color.",
					LongDescription:  "Sakurios switches are a special edition linear switch featuring a sakura pink color. They provide a smooth linear feel, making them popular among custom keyboard builders and enthusiasts who appreciate both aesthetics and performance.",
					ManufacturerID:   ptr(5), // ZealPC as manufacturer
					BrandID:          ptr(5), // ZealPC as brand
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-02-14"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Rosélios (Special Edition)",
					ShortDescription: "Special edition linear switch with unique aesthetics.",
					LongDescription:  "The Rosélios Special Edition switches offer a smooth linear experience with a unique rose-colored housing, catering to users who want a premium feel and appearance. This limited edition switch is perfect for custom builds and enthusiasts looking for something distinctive.",
					ManufacturerID:   ptr(5), // ZealPC as manufacturer
					BrandID:          ptr(5), // ZealPC as brand
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-05-18"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Turquoise Tealios",
					ShortDescription: "Smooth linear switch with a unique turquoise color.",
					LongDescription:  "Turquoise Tealios are a variant of the popular Tealios linear switch, featuring a distinctive turquoise coloration. They offer the same smooth linear experience as the original Tealios, with a different aesthetic appeal.",
					ManufacturerID:   ptr(5), // ZealPC
					BrandID:          ptr(5), // ZealPC
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-09-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Aqua Zilents",
					ShortDescription: "Silent tactile switch with a unique aqua color.",
					LongDescription:  "Aqua Zilents are a color variant of the Zilent silent tactile switches. They provide the same quiet tactile experience as regular Zilents, but with a distinctive aqua-colored housing, adding a unique aesthetic option for keyboard enthusiasts.",
					ManufacturerID:   ptr(5), // ZealPC
					BrandID:          ptr(5), // ZealPC
					SwitchTypeID:     5,      // Silent Tactile
					ReleaseDate:      parseDate("2020-09-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
			}

			err = processSwitches(zealpcSwitches)
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

			if err := tx.Model(&Producer{}).Where("alias IN (?)", []string{"cherry", "gateron"}).Pluck("id", &producersIDs).Error; err != nil {
				return err
			}

			if err := tx.Exec("DELETE FROM switches WHERE manufacturer_id IN (?)", producersIDs).Error; err != nil {
				return err
			}
			return nil // Replace with actual code
		},
	})
}

