package migrations_old

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

			ptr := func(i int) *int {
				return &i
			}

			parseDate := func(date string) *time.Time {
				t, _ := time.Parse("2006-01-02", date)
				return &t
			}

			outemus := []Switch{
				{
					Name:             "Outemu Blue",
					ShortDescription: "Clicky switch with a tactile bump and distinct sound.",
					LongDescription:  "Outemu Blue switches are known for their clicky sound and tactile bump, making them a popular choice for typists who enjoy audible feedback. With an actuation force of 50 grams, these switches provide a satisfying click and resistance, ideal for typists who value strong tactile and auditory feedback.",
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
					LongDescription:  "Outemu Ice Blue switches feature a transparent housing, making them perfect for RGB lighting setups in mechanical keyboards. These switches offer a clicky and tactile experience with an actuation force of 50 grams, delivering a satisfying typing feel while enhancing the visual aesthetics of your keyboard.",
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
					LongDescription:  "Outemu Purple switches are designed to deliver a louder click sound with a strong tactile bump. With an actuation force of 55 grams, these switches appeal to users who enjoy pronounced feedback, making them ideal for those who seek a more noticeable auditory experience during typing.",
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
					LongDescription:  "Outemu Black Dustproof switches offer a clicky typing experience combined with dustproof features, enhancing durability and consistent performance even in dusty environments. With an actuation force of 55 grams, these switches provide a tactile and audible feedback while ensuring longevity.",
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
					LongDescription:  "Outemu Green switches are heavy clicky switches that require more force to actuate, offering a strong tactile bump with an actuation force of 80 grams. Perfect for those who prefer a more resistant key press, these switches provide a satisfying and deliberate typing experience.",
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
					LongDescription:  "Outemu White switches offer a smooth clicky experience with a clear and sharp sound. With an actuation force of 50 grams, these switches provide an enjoyable typing experience for those who appreciate crisp auditory feedback combined with smooth keystrokes.",
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
					LongDescription:  "Outemu Sky Blue switches provide a balanced clicky experience with an actuation force of 55 grams, offering a satisfying tactile bump and moderate resistance for a comfortable typing experience. These switches are ideal for users who want a balance between tactile feedback and smooth operation.",
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
					LongDescription:  "Outemu Ocean Blue switches are designed for users who prefer a clicky switch with a lighter actuation force of 45 grams. These switches provide a quick response, making them suitable for fast typing and gaming, while still delivering satisfying tactile feedback.",
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
					LongDescription:  "Outemu Polar Blue switches provide enhanced tactile feedback with a clicky sound and an actuation force of 60 grams. These switches are ideal for users who value strong tactile sensations and a pronounced click, offering a robust and satisfying typing experience.",
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
					LongDescription:  "Outemu Brown switches offer a smooth tactile bump without a click sound, making them versatile for both typing and gaming enthusiasts. With an actuation force of 45 grams, these switches are perfect for users who prefer quieter switches with a subtle tactile response.",
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
					LongDescription:  "Outemu Ice Brown switches provide a tactile experience with a transparent housing, allowing for enhanced RGB lighting while maintaining a tactile feel. With an actuation force of 45 grams, these switches are ideal for users who seek both performance and aesthetics in their keyboard setups.",
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
					LongDescription:  "Outemu Purple Tactile switches offer a pronounced tactile bump with no click, providing a satisfying feedback with an actuation force of 50 grams. These switches are ideal for typing without the audible click sound, catering to users who prefer a quieter but responsive typing experience.",
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
					LongDescription:  "Outemu Orange switches offer a medium tactile experience with a lighter bump and an actuation force of 50 grams. These switches are ideal for users who prefer a softer touch with tactile feedback, providing a smooth and balanced typing feel.",
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
					LongDescription:  "Outemu Sky Tactile switches provide a satisfying tactile bump with a clear actuation point and an actuation force of 55 grams. These switches are perfect for users who want precise feedback during typing and gaming, offering a balanced and responsive experience.",
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
					LongDescription:  "Outemu Aqua switches are tactile switches that offer a smooth bump with a light actuation force of 45 grams. Ideal for users who want a tactile feel without too much resistance, these switches provide a comfortable and responsive typing experience.",
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
					LongDescription:  "Outemu Milk Brown switches provide a tactile typing experience with a muted sound profile and an actuation force of 50 grams. Perfect for quiet environments where tactile feedback is desired without noise, these switches offer a smooth and satisfying typing feel.",
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
					LongDescription:  "Outemu Jade switches are heavy tactile switches with an actuation force of 60 grams, offering a strong tactile bump and resistance. Ideal for those who enjoy pronounced tactile feedback, these switches provide a deliberate and satisfying typing experience.",
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
					LongDescription:  "Outemu Black switches are heavier linear switches with an actuation force of 60 grams, requiring more force to actuate. These switches offer a satisfying feel for those who prefer a more resistant key press with smooth and consistent keystrokes.",
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
					LongDescription:  "Outemu Ice Black switches offer a heavy linear feel with a transparent housing, allowing for enhanced RGB lighting effects. With an actuation force of 60 grams, these switches provide a robust typing experience with smooth and consistent performance.",
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
					LongDescription:  "Outemu Yellow switches provide a smooth linear experience with medium actuation force of 50 grams, suitable for both typing and gaming. These switches offer a balanced typing feel with no tactile bump, ideal for users who prefer a consistent and smooth keystroke.",
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
					LongDescription:  "Outemu Silver switches are speed linear switches designed for fast typing and gaming, featuring an actuation force of 45 grams and reduced travel distance for quicker actuation. These switches are perfect for users who prioritize speed and responsiveness in their keyboards.",
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
					LongDescription:  "Outemu Red Dustproof switches are linear switches with dustproof features, offering smooth actuation and an actuation force of 45 grams. These switches are ideal for users who want reliable performance in various environments, including those prone to dust.",
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
					LongDescription:  "Outemu Pink switches offer a linear typing experience with light actuation force of 45 grams, making them suitable for users who prefer minimal resistance. These switches provide a smooth and gentle typing feel, perfect for long typing sessions.",
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
					LongDescription:  "Outemu Peach switches provide a smooth linear experience with medium resistance and an actuation force of 50 grams. Perfect for typists who enjoy a balanced typing feel, these switches offer a consistent and satisfying keystroke.",
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
					LongDescription:  "Outemu Milk White switches are silent linear switches with noise-reducing features, offering a quiet typing experience for those who value silence. With an actuation force of 45 grams, these switches provide smooth and consistent keystrokes without the noise.",
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
					LongDescription:  "Outemu Cream switches offer a linear experience with a creamy, smooth feel and an actuation force of 50 grams. Ideal for users who enjoy a luxurious typing sensation, these switches provide a consistent and satisfying keystroke.",
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
					LongDescription:  "Outemu Purple Linear switches offer a smooth typing experience with no tactile bump and an actuation force of 45 grams. These switches provide a consistent feel for fast typing and gaming, ideal for users who prefer a seamless and uninterrupted keystroke.",
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
					LongDescription:  "Outemu Silent White switches offer a tactile typing experience with built-in noise-reducing features, perfect for quiet environments. With an actuation force of 45 grams, these switches provide a smooth tactile bump with minimal noise, catering to users who prefer a quieter typing experience.",
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
					LongDescription:  "Outemu Silent Red switches provide a smooth linear action with noise-dampening technology and an actuation force of 45 grams. These switches are ideal for environments where silence is essential, offering a quiet and responsive typing experience.",
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
					LongDescription:  "Outemu Silent Black switches are designed for those who prefer a heavier linear switch with an actuation force of 60 grams. These switches offer smooth operation with noise-dampening technology, providing resistance without the noise, perfect for quiet environments.",
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
					LongDescription:  "Outemu Ice Purple switches provide tactile feedback with a clear housing and an actuation force of 50 grams, allowing for optimal RGB lighting effects. These switches are perfect for users who desire both feedback and aesthetics, offering a balance between tactile sensation and visual appeal.",
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
					LongDescription:  "Outemu Polar Ice switches offer a unique smooth linear experience with an actuation force of 50 grams. Featuring a custom stem design that enhances stability and reduces wobble, these switches are ideal for precise typing, offering a consistent and stable keystroke.",
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
					LongDescription:  "Outemu Fire switches are designed for rapid actuation with an actuation force of 45 grams, providing a quick response suitable for gaming and fast typing. Ideal for users who prioritize speed and responsiveness, these switches offer a fast and precise typing experience.",
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
					LongDescription:  "Outemu Shadow switches offer a silent tactile experience with a subtle bump and an actuation force of 50 grams, making them perfect for users who enjoy tactile feedback without noise. These switches are ideal for quiet environments where a balance between feedback and silence is desired.",
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
					LongDescription:  "Outemu Crystal switches feature a fully transparent housing, designed to maximize RGB lighting while providing a consistent linear feel for typing and gaming. With an actuation force of 50 grams, these switches are perfect for users who want both performance and aesthetics in their keyboards.",
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
					LongDescription:  "Outemu Teal switches offer a tactile experience with a specially designed stem that provides a unique feedback and enhances overall stability. With an actuation force of 50 grams, these switches are perfect for users who appreciate distinctive tactile sensations and a consistent typing feel.",
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
					LongDescription:  "Outemu Moonlight switches provide a silent linear experience with a smooth keystroke and an actuation force of 45 grams. These switches offer a whisper-quiet operation, making them perfect for users who need silence while typing without sacrificing smoothness.",
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
					LongDescription:  "Outemu Red switches are light linear switches with an actuation force of 45 grams. These switches offer a smooth keystroke without tactile feedback, making them suitable for both typing and gaming where a seamless keystroke is desired.",
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
					LongDescription:  "Outemu Ice Light Purple switches offer a lighter tactile experience with an actuation force of 45 grams. These switches feature a transparent housing for enhanced RGB lighting effects, making them ideal for users who seek both a lighter touch and visual appeal.",
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
					LongDescription:  "Outemu Dustproof Red switches are linear switches with a dustproof design and an actuation force of 45 grams. These switches offer consistent performance and increased durability, making them ideal for various environments, including those prone to dust.",
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
					LongDescription:  "Zealios V2 switches offer a distinct tactile bump that is both smooth and satisfying. Designed for typists who prefer a clear but non-disruptive tactile feedback, these switches deliver precision and comfort with every keystroke, making them a top choice for enthusiasts seeking a refined typing experience.",
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
					LongDescription:  "Tealios V2 switches are celebrated for their buttery-smooth keystrokes, making them a favorite among both gamers and typists. The combination of a unique teal housing and high-quality materials ensures a frictionless, consistent typing experience with fast actuation, perfect for high-performance applications.",
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
					LongDescription:  "Healios switches are designed for near-silent operation, making them ideal for quiet work environments or late-night typing. These switches feature a dampened linear action that eliminates noise without compromising on the smoothness and responsiveness that ZealPC is known for.",
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
					LongDescription:  "Zilents V2 switches provide a silent yet tactile typing experience, combining a smooth, rounded bump with noise-dampening features. Ideal for office use or shared spaces, these switches deliver the satisfying tactile feedback of Zealios with minimal sound.",
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
					LongDescription:  "Rosélios switches stand out for their smooth linear action and elegant rose-colored housing. Perfect for those who value both aesthetics and performance, these switches provide a luxurious typing experience with minimal resistance, making them a popular choice for custom builds.",
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
					LongDescription:  "Zeal Clickiez switches offer a customizable clicky experience, allowing users to adjust the click element to suit their preferences. These switches are perfect for those who enjoy a tactile clicky feedback, with the flexibility to tailor the sound and feel to their liking.",
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
					LongDescription:  "Zealios Clicky V2 switches are engineered for those who love a satisfying click with every keystroke. With a crisp sound and smooth action, these switches provide a delightful clicky experience while maintaining ZealPC's renowned build quality.",
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
					LongDescription:  "Zealios V1 switches were the original tactile offering from ZealPC, providing a smooth tactile bump that is both satisfying and responsive. These switches are perfect for typists who appreciate a classic tactile feel, with a balance of smoothness and feedback. Although discontinued, they remain a favorite among keyboard enthusiasts.",
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
					LongDescription:  "The Zealios V2.2 is an enhanced version of the popular Zealios V2, featuring a more pronounced tactile bump and an even smoother keystroke. This switch is ideal for users who demand precise tactile feedback and a high-quality typing experience, making it a top choice for both typing and gaming.",
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
					LongDescription:  "Zilent V2 switches are designed to offer a quiet yet tactile typing experience. With a soft bump and noise-dampening technology, these switches are perfect for those who require both tactile feedback and silent operation, making them ideal for office environments and shared spaces.",
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
					LongDescription:  "Zilent V2.2 switches offer an upgraded silent tactile experience, with a softer and more refined bump than previous versions. Designed for users who need tactile feedback without the noise, these switches are perfect for quiet environments where precision and silence are essential.",
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
					LongDescription:  "Tealios V1 switches were the original linear switch offering from ZealPC, known for their ultra-smooth operation and distinctive teal housing. These switches provide a fast and effortless typing experience, making them popular among both gamers and typists. Although discontinued, they are still highly regarded in the mechanical keyboard community.",
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
					LongDescription:  "Rosélios V2 switches are the upgraded version of the original Rosélios, offering enhanced smoothness and a striking rose-colored housing. These switches are designed for users who value a premium linear experience with a touch of elegance, making them a standout choice for custom keyboard builds.",
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
					LongDescription:  "Zealios V2.1 Linear switches deliver the same high-quality build as the tactile version but with a smooth linear action. These switches are ideal for users who prefer a linear feel in their typing or gaming experience, offering a fast and responsive keystroke with minimal resistance.",
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
					LongDescription:  "Silent Zealios V1 switches were the first silent tactile switches from ZealPC, offering a smooth feedback with noise-dampening technology. Designed for quiet typing environments, these switches provide a soft tactile bump without the accompanying sound, making them perfect for offices and shared spaces. Although discontinued, they remain a beloved choice for silent typing enthusiasts.",
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
					LongDescription:  "Zilent V1 switches are known for their quiet operation and soft tactile bump, making them suitable for noise-sensitive environments. These switches offer a gentle tactile feel without the noise, perfect for offices and shared spaces where quiet typing is essential. Though discontinued, they remain a popular choice among those who prioritize silence in their typing experience.",
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
					LongDescription:  "Sakurios switches are a special edition linear switch featuring a sakura pink color. Designed for a smooth and quiet typing experience, these switches are popular among custom keyboard builders and enthusiasts who value both aesthetics and performance. Their limited availability makes them a sought-after choice for unique keyboard builds.",
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
					LongDescription:  "The Rosélios Special Edition switches offer a smooth linear experience with a unique rose-colored housing. Designed for users who want a premium feel and appearance, these limited edition switches are perfect for custom builds and enthusiasts looking for something distinctive.",
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
					LongDescription:  "Turquoise Tealios are a variant of the popular Tealios linear switch, featuring a distinctive turquoise color. They offer the same ultra-smooth keystrokes and high performance as the original Tealios, with a fresh and vibrant aesthetic that stands out in custom keyboard builds.",
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
					LongDescription:  "Aqua Zilents are a color variant of the Zilent silent tactile switches. They provide the same quiet, tactile experience as regular Zilents but with a distinctive aqua-colored housing. These switches are perfect for those who desire a silent tactile feel with a unique visual appeal.",
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

			if err := tx.Model(&Producer{}).Where("alias IN (?)", []string{"outemus", "zealps"}).Pluck("id", &producersIDs).Error; err != nil {
				return err
			}

			if err := tx.Exec("DELETE FROM switches WHERE brand_id IN (?)", producersIDs).Error; err != nil {
				return err
			}
			return nil // Replace with actual code
		},
	})
}
