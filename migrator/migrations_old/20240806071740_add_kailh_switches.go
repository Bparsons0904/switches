package migrations_old

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
					LongDescription:  "Kailh BOX White switches are popular clicky switches known for their crisp and precise click sound, offering a distinct tactile and audible feedback. With an actuation force of 50 grams, they are ideal for typists who enjoy a responsive and lively typing experience. The BOX design provides increased protection against dust and moisture, enhancing durability.",
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
					LongDescription:  "Kailh BOX Navy switches are designed for those who prefer a heavy click with pronounced tactile feedback. With a 75-gram actuation force, they offer a substantial key press and a loud, satisfying click, making them a favorite among users who seek both strong tactile and audible feedback. The BOX design enhances durability and protection against debris.",
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
					LongDescription:  "Kailh BOX Jade switches are tactile clicky switches known for their medium actuation force of 50 grams. They offer a delightful balance between tactile feedback and a sharp, satisfying click, making them a versatile choice for both typists and enthusiasts who appreciate tactile switches with a pronounced click.",
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
					LongDescription:  "Kailh Pro Light Green switches offer a light clicky feel with smooth actuation, providing a comfortable typing experience for long sessions. With a 50-gram actuation force, they are responsive without being too heavy, making them suitable for both casual and professional use.",
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
					LongDescription:  "Kailh Clicky Purple switches provide a unique clicky sound profile, combining a crisp auditory experience with a tactile bump. With an actuation force of 50 grams, they are an excellent choice for users who prioritize both sound and feel in their typing experience.",
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
					LongDescription:  "Kailh BOX Pink switches are light clicky switches designed for users who prefer a softer typing experience. With an actuation force of 45 grams, they offer a gentle click and tactile feedback, making them ideal for those who favor a light touch. The BOX design ensures durability and protection against dust and moisture.",
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
					LongDescription:  "Kailh Speed Bronze switches are optimized for speed, featuring a fast clicky actuation with a short travel distance. With an actuation force of 50 grams, they are perfect for gamers and typists who require quick key presses and tactile feedback. The short actuation point makes them one of the fastest clicky switches available.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(3), // Kailh
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2016-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kailh BOX Royal",
					ShortDescription: "Tactile switch with a royal feel.",
					LongDescription:  "Kailh BOX Royal switches are known for their strong tactile bump and unique 'royal' feel. With an actuation force of 75 grams, these switches offer a heavy and deliberate key press, making them a favorite among typists who prefer a substantial tactile feedback. The BOX design adds to their durability and reliability.",
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
					LongDescription:  "Kailh BOX Red switches offer a smooth linear feel with a light actuation force of 45 grams, making them ideal for fast-paced typing and gaming. The low actuation force provides minimal resistance, allowing for rapid key presses without tactile feedback. The BOX design enhances switch longevity and consistency.",
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
					LongDescription:  "Kailh BOX Black switches provide a linear feel with a heavier actuation force of 60 grams, making them suitable for users who prefer a more substantial key press. The increased force offers greater resistance, resulting in a deliberate and satisfying typing experience without tactile feedback. The BOX design ensures enhanced durability.",
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
					LongDescription:  "Kailh Speed Silver switches are designed for speed, featuring a fast linear actuation with a short travel distance of 1.1mm and a total travel of 3.5mm. With a light actuation force of 40 grams, these switches are perfect for competitive gaming and fast typists who need quick response times.",
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
					LongDescription:  "Kailh Pro Burgundy switches offer a smooth linear typing experience with a consistent actuation force of 50 grams. They are known for their quiet operation and buttery smooth keystrokes, making them ideal for both typing and gaming in quieter environments.",
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
					LongDescription:  "Kailh BOX Silent Pink switches are designed for those who need a quiet typing experience without sacrificing the smoothness of a linear switch. With an actuation force of 45 grams and built-in dampening, they are ideal for office settings or shared spaces where noise reduction is essential.",
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
					LongDescription:  "Kailh BOX Red Pro switches are an upgraded version of the original BOX Red, featuring an even smoother actuation with an actuation force of 45 grams. They are designed for gamers and typists who require precision and a consistent, light keystroke. The BOX design adds durability, making them a solid choice for extended use.",
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
					LongDescription:  "Kailh BOX Crystal Pink switches offer a smooth linear feel coupled with a striking transparent housing that enhances RGB lighting effects. With an actuation force of 45 grams, these switches provide a responsive and aesthetic typing experience, making them a favorite for custom builds.",
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
					LongDescription:  "Kailh Cream switches are highly regarded for their exceptionally smooth linear feel, featuring a 55-gram actuation force. Known for their creamy and consistent keystroke, they break in beautifully over time, making them a top choice for enthusiasts seeking a premium typing experience.",
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
					LongDescription:  "Kailh BOX Brown switches offer a tactile feel with a light bump and smooth actuation, making them a versatile choice for both typing and gaming. With an actuation force of 50 grams, they provide a satisfying tactile feedback without the noise of a clicky switch. The BOX design adds durability and protection against dust and moisture.",
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
					LongDescription:  "Kailh Pro Purple switches are designed with a shorter actuation distance and a medium actuation force of 50 grams. They provide a quick and tactile typing experience with a noticeable bump, making them ideal for fast typists and gamers who prefer a responsive switch with tactile feedback.",
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
					LongDescription:  "Kailh Speed Copper switches are engineered for speed and precision, featuring a short actuation distance and a light tactile bump. With an actuation force of 40 grams, these switches are perfect for users who need quick keypresses and appreciate the tactile feedback. Ideal for gaming and rapid typing.",
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
					LongDescription:  "Kailh Burnt Orange switches provide a tactile bump with a light actuation force of 50 grams, offering a smooth and comfortable typing experience. These switches are ideal for users who prefer a gentle tactile feedback without the fatigue of heavier switches, making them suitable for long typing sessions.",
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
					LongDescription:  "Kailh BOX Pale Blue switches feature a crisp and responsive tactile bump, with an actuation force of 60 grams. They provide a balanced typing experience suitable for both typists and gamers who enjoy tactile feedback without the noise of a clicky switch. The BOX design adds durability and protection against dust and moisture.",
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
					LongDescription:  "Kailh BOX Heavy Burnt Orange switches offer a strong tactile bump with a heavier actuation force of 70 grams, making them ideal for typists who prefer a more substantial key press. The pronounced bump provides satisfying feedback, while the BOX design enhances durability and longevity.",
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
					LongDescription:  "Kailh Polia switches provide a smooth tactile actuation with a distinct bump and an actuation force of 60 grams. They are known for their comfortable and consistent feedback, making them an excellent choice for typists who value both responsiveness and comfort during long typing sessions.",
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
					LongDescription:  "Kailh Silent Brown switches offer a tactile bump with silent operation, featuring an actuation force of 45 grams. These switches are designed for quiet environments, such as offices, where noise reduction is key while still providing tactile feedback for a satisfying typing experience.",
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
					LongDescription:  "Kailh BOX Dark Yellow switches feature a unique dark yellow stem and provide a heavy click with robust tactile feedback. With an actuation force of 65 grams, these switches cater to users who prefer a more substantial and deliberate typing experience, combined with the durability of the BOX design.",
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
					LongDescription:  "Kailh BOX Heavy Pale Blue switches offer a pronounced click and a robust feel, with an actuation force of 70 grams. These switches are perfect for typists who enjoy strong tactile and audible feedback, providing a satisfying and deliberate typing experience with the added protection of the BOX design.",
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
					LongDescription:  "Kailh BOX Heavy Pink switches provide a heavy actuation force of 70 grams with strong tactile feedback, making them ideal for those who prefer a pronounced typing experience. The distinctive pink stem and BOX design ensure both durability and a unique aesthetic.",
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
					LongDescription:  "Kailh Click Bar Jade switches feature a unique bar click mechanism that produces a crisp, sharp click, paired with tactile feedback. With an actuation force of 60 grams, these switches offer a distinctive sound profile and satisfying typing experience, appealing to users who enjoy both tactile and audible feedback.",
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
					LongDescription:  "Kailh BOX Red Silent switches offer a quiet linear typing experience with smooth actuation, featuring an actuation force of 45 grams. Built-in dampeners reduce noise, making these switches ideal for quiet environments. The BOX design ensures durability and consistent performance over time.",
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
					LongDescription:  "Kailh BOX Cream Pro switches are renowned for their creamy smooth feel, combined with enhanced durability for extended use. With an actuation force of 55 grams, these switches offer a premium linear typing experience that is both responsive and satisfying, ideal for enthusiasts and heavy typists.",
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
					LongDescription:  "Kailh BOX Red V2 switches deliver an updated linear experience with enhanced smoothness and consistency. With an actuation force of 45 grams, they are ideal for both gamers and typists who prefer a light, effortless keypress. The V2 improvements ensure a more refined keystroke and extended durability.",
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
					LongDescription:  "Kailh Smooth Linear switches are designed for an ultra-smooth typing experience with a light actuation force of 45 grams. These switches provide minimal friction and a seamless keypress, making them ideal for users seeking a responsive and fluid typing experience.",
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
					LongDescription:  "Kailh Chocolate switches feature a distinctive chocolate-colored housing and a smooth linear feel, with an actuation force of 50 grams. They offer a blend of aesthetics and performance, providing a reliable and enjoyable typing experience that stands out in custom builds.",
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
					LongDescription:  "Kailh Speed Pink switches offer a fast tactile experience with a short actuation distance and a light bump, featuring an actuation force of 40 grams. They are perfect for users who require quick response times and tactile feedback, making them ideal for gaming and fast-paced typing.",
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
					LongDescription:  "Kailh BOX Royal V2 switches offer an enhanced tactile experience with a more pronounced bump and improved feedback. With an actuation force of 70 grams, they cater to users who appreciate a strong tactile response and a substantial keypress. The V2 improvements bring refined actuation and increased durability.",
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
					LongDescription:  "Kailh BOX Silent Pale Blue switches combine a tactile bump with silent operation, featuring an actuation force of 45 grams. These switches are ideal for quiet environments where noise reduction is important, without compromising on the tactile feedback that typists enjoy.",
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
					LongDescription:  "Kailh Silent Heavy Orange switches offer a heavy actuation force of 70 grams combined with strong tactile feedback and silent operation. These switches are designed for users who require a robust typing experience in noise-sensitive environments, providing both feedback and quietness.",
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
					LongDescription:  "Kailh BOX Olive switches feature a unique olive stem and a smooth tactile bump, with an actuation force of 60 grams. They offer a satisfying blend of aesthetics and performance, making them a popular choice for custom keyboards where both look and feel matter.",
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
					LongDescription:  "The Kailh BOX White V2 switch provides an updated clicky experience with a crisp click and enhanced durability. With an actuation force of 50 grams, these switches are perfect for typists who enjoy clear, audible feedback, while the V2 improvements ensure longer-lasting performance.",
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
					LongDescription:  "Kailh BOX Brown V2 switches feature an enhanced tactile bump with smooth actuation and an actuation force of 55 grams. They provide a satisfying typing experience for users who enjoy tactile feedback without the noise of clicky switches, with the V2 improvements ensuring better consistency and durability.",
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
					LongDescription:  "The Kailh BOX Black V2 switch offers a linear feel with a heavier actuation force of 60 grams, providing a solid keypress for users who prefer more resistance. The V2 version brings improved smoothness and enhanced durability, making it a reliable choice for both gaming and typing.",
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
					LongDescription:  "The Kailh BOX Jade V2 switch offers a tactile clicky experience with an enhanced click and smoother keystroke. With an actuation force of 60 grams, these switches are perfect for typists who enjoy pronounced tactile feedback paired with a crisp click, while the V2 improvements add durability and consistency.",
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
					LongDescription:  "Kailh BOX Cream Milk switches feature a unique milky housing and deliver a smooth linear typing experience with an actuation force of 45 grams. These switches are designed for a gentle, quiet keypress, making them ideal for those who appreciate a smooth and consistent typing experience paired with a distinctive aesthetic.",
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
					LongDescription:  "Kailh BOX Cream Soda switches offer a bubbly smooth linear feel with an actuation force of 45 grams. They provide a delightful typing experience with subtle aesthetics, perfect for those seeking a light, gentle touch paired with distinctive visual appeal.",
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
					LongDescription:  "Kailh BOX Snow switches offer a gentle tactile bump with an actuation force of 50 grams, complemented by a frosty, winter-themed design. These switches are ideal for those who value both aesthetic and performance, providing a satisfying tactile feedback with a cool, unique appearance.",
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
					LongDescription:  "Kailh BOX Ice switches provide a crisp tactile bump with an actuation force of 55 grams, paired with an icy, winter-themed appearance. These switches are designed to offer a cool, refreshing aesthetic and a satisfying tactile feedback, making them a perfect choice for winter-themed builds.",
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
					LongDescription:  "Kailh BOX Glacier switches feature a pronounced tactile bump with an actuation force of 55 grams, coupled with cool, winter-themed aesthetics. These switches are perfect for users who enjoy strong tactile feedback and a unique, frosty design.",
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
					LongDescription:  "The Kailh Crystal Burgundy switch offers a smooth linear actuation with an actuation force of 45 grams and crystal-clear housing. These switches provide both a visually pleasing aesthetic and a responsive typing experience, making them ideal for builds where RGB lighting and performance are equally important.",
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
					LongDescription:  "Kailh Crystal Sapphire switches provide a smooth linear feel with an actuation force of 45 grams, enclosed in a sapphire-tinted housing. These switches combine visual appeal with a premium typing experience, making them an excellent choice for users who value both aesthetics and performance.",
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
					LongDescription:  "Kailh Crystal Ruby switches offer a premium linear experience with an actuation force of 45 grams, enhanced smoothness, and a striking ruby-red tint. These switches are perfect for users seeking a high-quality typing experience with distinctive aesthetics.",
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
					LongDescription:  "The Kailh Midnight Pro Silent Red switch delivers a silent linear experience with an actuation force of 45 grams. These switches are designed for smooth actuation with minimal noise, making them ideal for quiet environments or late-night typing sessions.",
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
					LongDescription:  "Kailh Midnight Pro Silent Brown switches offer a gentle tactile bump with an actuation force of 55 grams and quiet operation. These switches are ideal for noise-sensitive environments, providing a satisfying tactile feel without the noise of traditional tactile switches.",
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
					LongDescription:  "Kailh Midnight Pro Silent Black switches provide a silent linear experience with a heavier actuation force of 60 grams. These switches are designed for smooth, quiet operation, catering to users who prefer a more substantial feel without sacrificing silence.",
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
					LongDescription:  "The Kailh Clione Limacina switch offers an ultra-smooth linear feel with an actuation force of 45 grams, inspired by the delicate and ethereal beauty of sea angels. These switches are designed for users who value both a refined typing experience and unique, elegant aesthetics.",
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
					LongDescription:  "Kailh Clione Aqua switches feature a smooth linear actuation with an actuation force of 45 grams, complemented by a unique aquatic-themed design. These switches are perfect for users who appreciate a seamless typing experience combined with eye-catching aesthetics.",
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
					LongDescription:  "Kailh Clione Ocean switches provide a buttery smooth linear feel with an actuation force of 45 grams and an oceanic-themed design. These switches are designed to offer a premium typing experience with a touch of marine elegance.",
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
					LongDescription:  "The Kailh Super Speed Silver switch is engineered for speed, featuring an ultra-short actuation distance of just 1.1 mm and a total travel of 3.5 mm. With an actuation force of 40 grams, these switches are ideal for competitive gaming and fast typists who demand quick, responsive keystrokes.",
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
					LongDescription:  "Kailh Super Speed Copper switches combine rapid actuation with tactile feedback, featuring an actuation force of 45 grams and an actuation distance of 1.1 mm. These switches are ideal for users who need quick response times with the satisfying feel of a tactile bump.",
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
					LongDescription:  "Kailh Super Speed Gold switches offer a lightning-fast linear experience with an actuation force of 40 grams and an actuation distance of 1.1 mm. Enhanced with gold accents, these switches cater to gamers and typists who prioritize speed and aesthetics.",
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
					LongDescription:  "The Kailh Deep Sea Silent Pro switch delivers deep, satisfying tactile feedback with an actuation force of 55 grams and minimal noise. These switches are perfect for users who prefer a quiet typing experience with pronounced tactile feedback.",
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
					LongDescription:  "The Kailh Jellyfish Box Linear (X) switch offers smooth linear actuation with an actuation force of 45 grams, inspired by the fluid and graceful motion of jellyfish. These switches are perfect for users who seek a gentle, seamless typing experience.",
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
					LongDescription:  "Kailh Jellyfish Box Clicky (Y) switches provide a distinct clicky sound profile with an actuation force of 50 grams, capturing the essence of jellyfish in motion. These switches are ideal for users who enjoy both a unique design and pronounced auditory feedback.",
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
					LongDescription:  "The Kailh White Owl switch features a gentle tactile feedback with an actuation force of 50 grams, known for its aesthetic design and smooth operation. These switches are perfect for users who value both a pleasing visual and a balanced tactile experience.",
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
					LongDescription:  "The Kailh Red Bean Pudding Box switch offers a smooth linear experience with an actuation force of 45 grams and a unique pudding-like texture. These switches combine visual appeal with a soft, consistent keypress, making them perfect for users who seek both performance and distinctive aesthetics.",
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
					LongDescription:  "Kailh Coco Pink Box V2 switches offer a vibrant pink color and smooth linear operation with an actuation force of 45 grams. These switches combine style with performance, providing a seamless typing experience for those who enjoy a bold, colorful aesthetic.",
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
					LongDescription:  "The Kailh Box Cream switch provides a creamy smooth linear experience with an actuation force of 45 grams and high durability. These switches have become a favorite among enthusiasts for their performance and longevity, making them ideal for both typing and gaming.",
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
					LongDescription:  "Kailh Box Cream Pro switches offer an enhanced smooth linear experience with an actuation force of 45 grams and an extended lifespan. These switches are designed for discerning users who seek a premium typing experience with long-lasting performance.",
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
					LongDescription:  "The Kailh Box Summer Clicky switch provides a refreshing and lively typing experience with an actuation force of 55 grams. With its vibrant design and crisp click feedback, these switches are perfect for users who enjoy both strong auditory and tactile feedback.",
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
					LongDescription:  "Kailh Box Winter switches offer a tactile experience with an actuation force of 50 grams, paired with a winter-themed design. These switches provide a unique feel for those who appreciate themed aesthetics and satisfying tactile feedback.",
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
					LongDescription:  "The Kailh Box Brown Silent switch provides a quiet tactile experience with an actuation force of 55 grams and gentle feedback, making them perfect for noise-sensitive environments and typists who prefer a subtle tactile response.",
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
					LongDescription:  "Kailh Box Red Silent switches offer a smooth linear experience with an actuation force of 45 grams and minimal noise. These switches are designed for users who require a quiet typing environment without sacrificing performance.",
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
					LongDescription:  "The Kailh Box Navy V2 switch offers an updated heavy clicky experience with an actuation force of 70 grams, providing loud and tactile feedback. These switches are designed for users who enjoy pronounced sound and feel, making them ideal for those who seek an unmistakable click.",
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
					LongDescription:  "Kailh Box Pink V2 switches offer a light clicky experience with an actuation force of 55 grams and a gentle actuation force. These switches are perfect for users who enjoy vibrant aesthetics and satisfying auditory feedback.",
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
					LongDescription:  "The Kailh Box Heavy Black switch offers a heavy linear experience with an actuation force of 65 grams, catering to users who prefer a more substantial and deliberate typing feel. These switches are ideal for those who enjoy a pronounced keypress with smooth actuation.",
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
					LongDescription:  "The Kailh Box Heavy Yellow switch is designed for users who prefer a heavier key press with an actuation force of 65 grams. These switches offer a smooth linear actuation with minimal tactile feedback, making them ideal for those seeking a deliberate typing experience.",
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
					LongDescription:  "The Kailh Pro Brown switch offers a tactile experience with a medium bump, an actuation force of 50 grams, and a short actuation distance of 1.7 mm. These switches are ideal for fast typing and gaming, providing responsive feedback with a satisfying tactile feel.",
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
					LongDescription:  "Kailh Pro Silent Red switches offer a smooth linear experience with an actuation force of 45 grams and minimal noise. These switches are perfect for users who require a quiet typing environment, making them ideal for late-night use or noise-sensitive spaces.",
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
					LongDescription:  "The Kailh Crystal Royal switch provides a tactile experience with a pronounced bump, an actuation force of 55 grams, and a clear housing for maximum RGB visibility. These switches are ideal for users who value strong tactile feedback and a visually striking keyboard build.",
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
					LongDescription:  "Kailh Crystal Clear switches feature an ultra-clear housing for maximum RGB visibility, providing a smooth linear actuation with an actuation force of 45 grams. These switches are a perfect choice for users who prioritize both aesthetics and performance in their builds.",
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
					LongDescription:  "The Kailh Speed Purple switch provides a tactile experience with an actuation force of 45 grams, a short actuation distance of 1.1 mm, and a light tactile bump. These switches are perfect for fast typing and gaming that requires quick response and tactile feedback.",
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
					LongDescription:  "Kailh Super Speed Bronze switches offer a tactile experience with an actuation force of 50 grams, a short actuation distance of 1.1 mm, and a pronounced tactile bump. These switches provide a balance of speed and feedback, ideal for users who need quick response times without sacrificing tactile sensation.",
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
					LongDescription:  "The Kailh Speed Pro Heavy Green switch combines a heavier actuation force of 70 grams with a tactile bump, offering a substantial typing experience for those who prefer a more deliberate feel. These switches are designed for users who require both speed and strong tactile feedback.",
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
					LongDescription:  "Kailh Box Noble Yellow switches offer a smooth linear experience with a light actuation force of 45 grams, providing a comfortable typing experience for extended use. These switches are ideal for users who prefer a gentle keypress without sacrificing smoothness.",
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
					LongDescription:  "Kailh Box Mute Jade switches provide a tactile experience with an actuation force of 50 grams, similar to the Box Jade but with reduced noise. These switches are suitable for environments where quieter operation is preferred without compromising on tactile feedback.",
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
					LongDescription:  "The Kailh Canary switch offers a linear experience with an actuation force of 45 grams and a distinctive yellow color. These switches provide both aesthetic appeal and smooth operation, making them a great choice for users who want a unique yet functional switch.",
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
					LongDescription:  "Kailh Box Ancient Gray switches provide a heavy linear experience with an increased actuation force of 65 grams, catering to users who prefer a more substantial key press. These switches are ideal for those who enjoy a pronounced keypress with smooth actuation.",
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
					LongDescription:  "The Kailh Box Pro Green switch offers a clicky experience with an actuation force of 60 grams, providing a unique feel and feedback different from regular Box switches. These switches are designed for users who appreciate both tactile and audible feedback in their typing experience.",
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
					LongDescription:  "Kailh Sage switches offer a tactile experience with an actuation force of 55 grams and a unique bump, providing a distinctive feel for enthusiasts who appreciate novel tactile sensations. These switches are ideal for users who seek a balanced tactile response with a unique twist.",
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
					LongDescription:  "The Kailh Speed Gold switch is a clicky switch designed for fast actuation with an actuation force of 50 grams and an actuation distance of 1.1 mm. These switches are ideal for gamers and fast typists who enjoy audible feedback and quick keystrokes.",
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
					LongDescription:  "Kailh Box Glazed Green switches are a limited edition tactile switch with an actuation force of 55 grams, offering a unique aesthetic and tactile experience for keyboard enthusiasts. These switches are perfect for users who appreciate both distinctive design and satisfying tactile feedback.",
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
			var producersIDs []int
			type Producer struct {
				ID int `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
			}

			if err := tx.Model(&Producer{}).Where("alias IN (?)", []string{"kailh"}).Pluck("id", &producersIDs).Error; err != nil {
				return err
			}

			if err := tx.Exec("DELETE FROM switches WHERE brand_id IN (?)", producersIDs).Error; err != nil {
				return err
			}
			return nil // Replace with actual code
		},
	})
}
