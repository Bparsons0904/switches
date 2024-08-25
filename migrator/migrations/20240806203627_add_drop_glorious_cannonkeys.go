package migrations

import (
	"time"

	"github.com/Bparsons0904/deadigations"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	deadigations.RegisterMigration(deadigations.Migration{
		ID:          "20240806203627",
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

			cannonKeysSwitches := []Switch{
				{
					Name:             "CannonKeys Lavenders",
					ShortDescription: "Smooth linear switch with a unique lavender color.",
					LongDescription:  "CannonKeys Lavenders are linear switches celebrated for their smooth, consistent feel and light actuation force. The lavender-colored housing and stem add a unique visual flair to any keyboard, making them a favorite among enthusiasts who prioritize both aesthetics and performance.",
					ManufacturerID:   ptr(22), // JWK
					BrandID:          ptr(28), // Cannon Keys
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2020-11-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "CannonKeys Anubis",
					ShortDescription: "Tactile switch with a medium actuation force.",
					LongDescription:  "The CannonKeys Anubis switch is a tactile switch that delivers a satisfying medium actuation force combined with a prominent tactile bump. Known for its smooth operation and deep sound profile, the Anubis switch is highly regarded among tactile switch enthusiasts who seek a rich typing experience.",
					ManufacturerID:   ptr(23), // Tecsee
					BrandID:          ptr(28), // Cannon Keys
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2021-08-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "CannonKeys Dragonfruit",
					ShortDescription: "Tactile switch with a heavy bump and unique aesthetics.",
					LongDescription:  "Dragonfruit switches from CannonKeys offer a distinct tactile bump and a vibrant color inspired by the exotic dragon fruit. These tactile switches provide a delightful typing experience with a balance of responsiveness and feedback, making them a standout choice for those who appreciate unique aesthetics and performance.",
					ManufacturerID:   ptr(23), // Tecsee
					BrandID:          ptr(28), // Cannon Keys
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2021-05-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "CannonKeys Banana Split",
					ShortDescription: "Linear switch with a colorful and smooth experience.",
					LongDescription:  "CannonKeys Banana Split switches are designed for those who enjoy a vibrant, playful aesthetic paired with ultra-smooth linear action. These switches offer a delightful typing experience with minimal resistance, making them a popular choice for both gaming and typing enthusiasts.",
					ManufacturerID:   ptr(23), // Tecsee
					BrandID:          ptr(28), // Cannon Keys
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2021-02-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "CannonKeys Durock POM",
					ShortDescription: "Linear switch with a solid POM housing.",
					LongDescription:  "The CannonKeys Durock POM switch is a premium linear switch featuring a POM (polyoxymethylene) housing that ensures a smooth and consistent typing experience. Known for its durability and minimal scratchiness, this switch is ideal for heavy typists seeking a reliable and satisfying linear feel.",
					ManufacturerID:   ptr(7),  // Durock
					BrandID:          ptr(28), // Cannon Keys
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2020-06-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "CannonKeys Blueberry",
					ShortDescription: "Clicky switch with a tactile bump and crisp sound.",
					LongDescription:  "The CannonKeys Blueberry switch is a clicky switch designed to deliver a strong tactile bump accompanied by a crisp, satisfying sound. It offers a distinctive typing experience that combines tactile feedback with a sharp auditory click, making it a favorite for those who enjoy pronounced feedback.",
					ManufacturerID:   ptr(23), // Tecsee
					BrandID:          ptr(28), // Cannon Keys
					SwitchTypeID:     3,       // Clicky
					ReleaseDate:      parseDate("2020-10-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "CannonKeys Orange Pies",
					ShortDescription: "Clicky switch with a strong tactile feel and sharp sound.",
					LongDescription:  "CannonKeys Orange Pies switches offer a strong tactile bump with a sharp, clicky sound. These switches are perfect for typists who prefer a pronounced tactile feedback paired with a distinct auditory experience, making each keystroke satisfying and engaging.",
					ManufacturerID:   ptr(23), // Tecsee
					BrandID:          ptr(28), // Cannon Keys
					SwitchTypeID:     3,       // Clicky
					ReleaseDate:      parseDate("2021-04-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "CannonKeys Matcha Latte",
					ShortDescription: "Linear switch with a creamy smoothness and light actuation.",
					LongDescription:  "CannonKeys Matcha Latte switches provide a smooth linear action with a light actuation force, making them ideal for typists who prefer a gentle, effortless keystroke. The matcha-colored housing adds a unique and aesthetically pleasing touch to any keyboard.",
					ManufacturerID:   ptr(23), // Tecsee
					BrandID:          ptr(28), // Cannon Keys
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2021-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "CannonKeys Taro Balls",
					ShortDescription: "Linear switch with a medium weight and quiet operation.",
					LongDescription:  "CannonKeys Taro Balls switches offer a medium actuation force with a quiet and smooth linear action. Inspired by the purple taro root, these switches provide a unique look and a satisfying typing experience, perfect for those who appreciate both performance and aesthetics.",
					ManufacturerID:   ptr(23), // Tecsee
					BrandID:          ptr(28), // Cannon Keys
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2021-10-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "CannonKeys Crystal",
					ShortDescription: "Linear switch with a transparent housing and smooth feel.",
					LongDescription:  "CannonKeys Crystal switches feature a transparent housing that beautifully showcases RGB lighting while delivering a smooth linear action. These switches are designed for users who seek both performance and visual appeal, offering a clean and consistent typing experience.",
					ManufacturerID:   ptr(23), // Tecsee
					BrandID:          ptr(28), // Cannon Keys
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2022-02-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "CannonKeys Silent Alpaca V2",
					ShortDescription: "Silent linear switch with a smooth feel and muted sound.",
					LongDescription:  "CannonKeys Silent Alpaca V2 switches are designed to provide a silent typing experience without compromising on smoothness. The switch's unique construction minimizes noise, making it ideal for environments where quiet operation is essential, while still offering the smooth feel Alpacas are known for.",
					ManufacturerID:   ptr(22), // JWK
					BrandID:          ptr(28), // Cannon Keys
					SwitchTypeID:     4,       // Silent Linear
					ReleaseDate:      parseDate("2020-12-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "CannonKeys Boba U4",
					ShortDescription: "Silent tactile switch with a soft bump and quiet operation.",
					LongDescription:  "CannonKeys Boba U4 switches are known for their silent operation and soft tactile bump, providing a satisfying yet quiet typing experience. These switches are perfect for shared workspaces or environments where noise reduction is a priority.",
					ManufacturerID:   ptr(26), // Gazzew
					BrandID:          ptr(28), // Cannon Keys
					SwitchTypeID:     5,       // Silent Tactile
					ReleaseDate:      parseDate("2021-03-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "CannonKeys Mint Chocolate",
					ShortDescription: "Specialty switch with a mint-themed color and smooth action.",
					LongDescription:  "The CannonKeys Mint Chocolate switch is a specialty linear switch that combines a mint-themed colorway with a smooth, consistent action. Its unique design and pleasant typing feel make it a popular choice for custom keyboard enthusiasts looking for both style and performance.",
					ManufacturerID:   ptr(23), // Tecsee
					BrandID:          ptr(28), // Cannon Keys
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2022-05-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "CannonKeys Snow Globe",
					ShortDescription: "Specialty linear switch with frosted housing and smooth action.",
					LongDescription:  "The CannonKeys Snow Globe switch is a specialty linear switch featuring a frosted housing that complements its smooth typing action. With its winter-themed design, this switch provides both a unique aesthetic and a satisfying, smooth keystroke, perfect for custom builds.",
					ManufacturerID:   ptr(22), // JWK
					BrandID:          ptr(28), // Cannon Keys
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2022-12-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "CannonKeys Milkshake",
					ShortDescription: "Specialty linear switch with a milky smooth feel.",
					LongDescription:  "CannonKeys Milkshake switches deliver a milky smooth linear action with a focus on consistency and a satisfying keystroke. Ideal for those who prefer a straightforward and enjoyable typing experience, these switches combine smoothness with a unique, playful aesthetic.",
					ManufacturerID:   ptr(23), // Tecsee
					BrandID:          ptr(28), // Cannon Keys
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2021-08-01"),
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

			err := processSwitches(cannonKeysSwitches)
			if err != nil {
				return err
			}

			dropSwitches := []Switch{
				{
					Name:             "Drop Halo True",
					ShortDescription: "Tactile switch with a smooth bump and minimal noise.",
					LongDescription:  "Drop Halo True switches offer a smooth and satisfying tactile bump with minimal noise, making them a great choice for quiet environments. The Halo True provides a unique typing experience that balances tactility with quiet operation.",
					ManufacturerID:   ptr(3),  // Kaihua
					BrandID:          ptr(10), // Drop
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2017-12-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Drop Holy Panda",
					ShortDescription: "Tactile switch with a unique feel and excellent sound.",
					LongDescription:  "The Drop Holy Panda switch is legendary for its satisfying tactile bump and signature sound profile. Combining the Halo True stem and Invyr Panda housing, it delivers a unique typing feel that is cherished by keyboard enthusiasts.",
					ManufacturerID:   ptr(2),  // Gateron
					BrandID:          ptr(10), // Drop
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2018-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Drop Invyr Holy Panda",
					ShortDescription: "Tactile switch with the famous Holy Panda feel.",
					LongDescription:  "Invyr Holy Panda switches from Drop provide the classic Holy Panda typing experience with their renowned tactile bump and satisfying sound. These switches are a must-have for tactile switch enthusiasts.",
					ManufacturerID:   ptr(2),  // Gateron
					BrandID:          ptr(10), // Drop
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Drop ALT Rock Black",
					ShortDescription: "Linear switch with a heavier feel and smooth action.",
					LongDescription:  "Drop ALT Rock Black switches provide a smooth linear action with a heavier actuation force, catering to those who prefer more resistance in their typing. These switches are favored by gamers and typists who enjoy a substantial feel.",
					ManufacturerID:   ptr(22), // JWK
					BrandID:          ptr(10), // Drop
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2021-09-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Drop Silver Speed",
					ShortDescription: "Linear switch designed for fast actuation and gaming.",
					LongDescription:  "The Drop Silver Speed switch is optimized for fast actuation, making it ideal for gaming and rapid typing. Its smooth linear action ensures quick and precise inputs, reducing lag for competitive gamers.",
					ManufacturerID:   ptr(3),  // Kailh
					BrandID:          ptr(10), // Drop
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2022-04-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Drop Kaihua Box White",
					ShortDescription: "Clicky switch with a light actuation force and box design.",
					LongDescription:  "Drop Kaihua Box White switches feature a box design that protects against dust and moisture while offering a light actuation force and a crisp, satisfying click. They are a popular choice among fans of clicky switches who value durability and performance.",
					ManufacturerID:   ptr(3),  // Kailh
					BrandID:          ptr(10), // Drop
					SwitchTypeID:     3,       // Clicky
					ReleaseDate:      parseDate("2018-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Drop Kaihua Speed Silver",
					ShortDescription: "Linear switch designed for fast actuation and gaming.",
					LongDescription:  "The Drop Kaihua Speed Silver switch offers rapid actuation with a smooth linear action, making it a top choice for gamers and fast typists. Its quick response time and seamless performance ensure a competitive edge in fast-paced scenarios.",
					ManufacturerID:   ptr(3),  // Kailh
					BrandID:          ptr(10), // Drop
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2019-03-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Drop Kaihua Jade",
					ShortDescription: "Clicky switch with a thicker click bar for a louder sound.",
					LongDescription:  "Drop Kaihua Jade switches are known for their thicker click bar, resulting in a louder and more satisfying click. These switches offer pronounced tactile feedback, making them perfect for users who enjoy a bold and assertive typing experience.",
					ManufacturerID:   ptr(3),  // Kailh
					BrandID:          ptr(10), // Drop
					SwitchTypeID:     3,       // Clicky
					ReleaseDate:      parseDate("2019-07-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Drop Halo Linear",
					ShortDescription: "Linear switch with a smooth feel and minimal resistance.",
					LongDescription:  "Drop Halo Linear switches are designed for a smooth and consistent typing experience with minimal resistance. Ideal for users who prefer a linear switch without tactile feedback, they offer a straightforward and satisfying keystroke.",
					ManufacturerID:   ptr(3),  // Kaihua
					BrandID:          ptr(10), // Drop
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2019-07-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Drop Silent Aliaz",
					ShortDescription: "Silent tactile switch with a soft bump and minimal noise.",
					LongDescription:  "Drop Silent Aliaz switches offer a quiet tactile experience with a soft bump and minimal noise. These switches are ideal for environments where silence is essential, providing a peaceful typing experience without compromising on feedback.",
					ManufacturerID:   ptr(3),  // Kaihua
					BrandID:          ptr(10), // Drop
					SwitchTypeID:     5,       // Silent Tactile
					ReleaseDate:      parseDate("2019-11-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Drop Kaihua Silent Pink",
					ShortDescription: "Silent linear switch with a fast actuation and pink color.",
					LongDescription:  "The Drop Kaihua Silent Pink switch provides a fast-actuating silent linear experience with a vibrant pink color. These switches are perfect for users who value both aesthetics and quiet operation, making them ideal for office use or shared spaces.",
					ManufacturerID:   ptr(3),  // Kailh
					BrandID:          ptr(10), // Drop
					SwitchTypeID:     4,       // Silent Linear
					ReleaseDate:      parseDate("2020-02-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Drop Emerald Silent Tactile",
					ShortDescription: "Silent tactile switch with a gentle bump and emerald design.",
					LongDescription:  "The Drop Emerald Silent Tactile switch features a gentle tactile bump combined with a quiet operation, making it a perfect choice for those who prefer a subdued yet satisfying typing experience. The emerald-themed design adds a touch of elegance to any keyboard.",
					ManufacturerID:   ptr(3),  // Kailh
					BrandID:          ptr(10), // Drop
					SwitchTypeID:     5,       // Silent Tactile
					ReleaseDate:      parseDate("2021-07-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Drop Holy Boba",
					ShortDescription: "Specialty tactile switch with Holy Panda and Boba U4 elements.",
					LongDescription:  "Drop Holy Boba switches are a unique combination of Holy Panda and Boba U4 elements, resulting in a premium tactile experience with a smooth bump and satisfying feel. These switches are perfect for enthusiasts who value both innovative design and excellent feedback.",
					ManufacturerID:   ptr(26), // Gazzew
					BrandID:          ptr(10), // Drop
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2022-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
			}
			err = processSwitches(dropSwitches)
			if err != nil {
				return err
			}

			gloriousSwitches := []Switch{
				{
					Name:             "Glorious Panda",
					ShortDescription: "Tactile switch with a crisp bump and smooth travel.",
					LongDescription:  "Glorious Panda switches are renowned for their crisp tactile bump and smooth travel, providing a premium typing experience with notable tactile feedback. These switches are favored by enthusiasts for their responsive and satisfying feel.",
					ManufacturerID:   ptr(11), // Glorious
					BrandID:          ptr(11), // Glorious
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2020-07-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Glorious Lynx",
					ShortDescription: "Linear switch with a smooth, fast feel for gaming.",
					LongDescription:  "The Glorious Lynx switch is designed for gaming enthusiasts who seek a smooth and fast linear feel. Its consistent actuation and quick response make it an ideal choice for competitive gaming and fast typing.",
					ManufacturerID:   ptr(11), // Glorious
					BrandID:          ptr(11), // Glorious
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2021-11-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Glorious Fox",
					ShortDescription: "Linear switch with a medium actuation force.",
					LongDescription:  "Glorious Fox switches offer a balanced linear experience with a medium actuation force, making them versatile for both gaming and everyday typing. Their smooth and consistent performance is ideal for a wide range of users.",
					ManufacturerID:   ptr(11), // Glorious
					BrandID:          ptr(11), // Glorious
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2022-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Glorious Ice",
					ShortDescription: "Linear switch with a smooth, icy feel and quiet operation.",
					LongDescription:  "The Glorious Ice switch offers a smooth linear action with quiet operation, designed to provide a cool, 'icy' feel. It's perfect for those who prefer a silent typing experience without compromising on smoothness.",
					ManufacturerID:   ptr(11), // Glorious
					BrandID:          ptr(11), // Glorious
					SwitchTypeID:     4,       // Silent Linear
					ReleaseDate:      parseDate("2023-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Glorious Holy Panda",
					ShortDescription: "Tactile switch with a distinct bump and premium feel.",
					LongDescription:  "The Glorious Holy Panda switch is a premium tactile switch that delivers a distinct bump with a satisfying feel. Loved by enthusiasts, it offers a top-tier typing experience with excellent feedback.",
					ManufacturerID:   ptr(11), // Glorious
					BrandID:          ptr(11), // Glorious
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2021-03-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Glorious Goat Clicky",
					ShortDescription: "Clicky switch with a medium actuation force and goat logo.",
					LongDescription:  "Glorious Goat Clicky switches feature a crisp click sound with a medium actuation force, delivering an engaging typing experience. The unique goat logo and sharp auditory feedback make them a standout choice for clicky switch fans.",
					ManufacturerID:   ptr(11), // Glorious
					BrandID:          ptr(11), // Glorious
					SwitchTypeID:     3,       // Clicky
					ReleaseDate:      parseDate("2022-02-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Glorious Fox Clicky",
					ShortDescription: "Clicky switch with a smooth action and audible feedback.",
					LongDescription:  "The Glorious Fox Clicky switch offers smooth action paired with satisfying clicky feedback. It is designed for users who enjoy a balance of tactile feel and audible satisfaction, making it perfect for typing and gaming alike.",
					ManufacturerID:   ptr(11), // Glorious
					BrandID:          ptr(11), // Glorious
					SwitchTypeID:     3,       // Clicky
					ReleaseDate:      parseDate("2021-12-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Glorious Lynx Clicky",
					ShortDescription: "Clicky switch with a light actuation force for fast typing.",
					LongDescription:  "Glorious Lynx Clicky switches provide a light actuation force, ideal for fast typing and gaming. The crisp clicky sound enhances the typing experience, delivering both auditory and tactile feedback.",
					ManufacturerID:   ptr(11), // Glorious
					BrandID:          ptr(11), // Glorious
					SwitchTypeID:     3,       // Clicky
					ReleaseDate:      parseDate("2022-05-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Glorious Panda Linear",
					ShortDescription: "Linear switch with a smooth and consistent feel.",
					LongDescription:  "Glorious Panda Linear switches provide a smooth and consistent linear feel, offering a premium typing experience suited for both gaming and everyday use. They are designed to deliver a high-quality, reliable keystroke.",
					ManufacturerID:   ptr(11), // Glorious
					BrandID:          ptr(11), // Glorious
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2021-09-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Glorious Fox Linear V2",
					ShortDescription: "Linear switch with a light actuation and refined design.",
					LongDescription:  "The Glorious Fox Linear V2 switch features a light actuation force and a refined design, delivering a seamless typing experience with enhanced smoothness and aesthetics. It's perfect for users seeking a responsive and elegant switch.",
					ManufacturerID:   ptr(11), // Glorious
					BrandID:          ptr(11), // Glorious
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2023-02-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Glorious Lynx Linear V2",
					ShortDescription: "Linear switch with improved smoothness and sound.",
					LongDescription:  "Glorious Lynx Linear V2 switches feature enhanced smoothness and sound, providing a refined typing experience for those who prefer a quieter and more fluid keystroke. These switches are a great upgrade for linear switch fans.",
					ManufacturerID:   ptr(11), // Glorious
					BrandID:          ptr(11), // Glorious
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2022-08-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Glorious Ice Linear V2",
					ShortDescription: "Linear switch with a quiet operation and icy aesthetics.",
					LongDescription:  "The Glorious Ice Linear V2 switch offers a quiet and smooth typing experience, featuring icy aesthetics that add a cool visual appeal to any keyboard. Its silent operation makes it ideal for office use or shared spaces.",
					ManufacturerID:   ptr(11), // Glorious
					BrandID:          ptr(11), // Glorious
					SwitchTypeID:     4,       // Silent Linear
					ReleaseDate:      parseDate("2023-05-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Glorious Panda Silent",
					ShortDescription: "Silent tactile switch with a quiet bump and smooth feel.",
					LongDescription:  "Glorious Panda Silent switches offer a quiet tactile bump with a smooth feel, providing a premium typing experience without the noise. These switches are perfect for those who enjoy tactile feedback without disturbing others.",
					ManufacturerID:   ptr(11), // Glorious
					BrandID:          ptr(11), // Glorious
					SwitchTypeID:     5,       // Silent Tactile
					ReleaseDate:      parseDate("2021-11-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Glorious Lynx Silent",
					ShortDescription: "Silent linear switch with a soft feel and quiet operation.",
					LongDescription:  "The Glorious Lynx Silent switch offers a soft linear action with silent operation, making it a great choice for offices or shared spaces. Its understated aesthetics and smooth performance make it a versatile switch for any build.",
					ManufacturerID:   ptr(11), // Glorious
					BrandID:          ptr(11), // Glorious
					SwitchTypeID:     4,       // Silent Linear
					ReleaseDate:      parseDate("2022-07-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Glorious Crystal Ice",
					ShortDescription: "Specialty linear switch with transparent housing and icy smoothness.",
					LongDescription:  "The Glorious Crystal Ice switch features a smooth linear action with a transparent housing, allowing RGB lighting to shine through. Its unique design and consistent feel make it a popular choice for custom builds that prioritize both aesthetics and performance.",
					ManufacturerID:   ptr(11), // Glorious
					BrandID:          ptr(11), // Glorious
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2023-04-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Glorious Aurora",
					ShortDescription: "Specialty tactile switch with RGB diffusion and enhanced tactility.",
					LongDescription:  "The Glorious Aurora switch offers enhanced tactile feedback combined with excellent RGB diffusion, providing both visual appeal and a satisfying typing experience. Its innovative design makes it a standout choice for enthusiasts who value both aesthetics and performance.",
					ManufacturerID:   ptr(11), // Glorious
					BrandID:          ptr(11), // Glorious
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2023-03-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Glorious Summer Sun",
					ShortDescription: "Specialty linear switch with a bright design and smooth feel.",
					LongDescription:  "The Glorious Summer Sun switch offers a bright and vibrant design paired with a smooth linear action, capturing the essence of summer in both aesthetics and typing experience. It's a popular choice for those who love cheerful keyboard setups.",
					ManufacturerID:   ptr(11), // Glorious
					BrandID:          ptr(11), // Glorious
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2023-07-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Glorious Bobcat",
					ShortDescription: "Tactile switch with a pronounced bump and medium actuation force.",
					LongDescription:  "The Glorious Bobcat switch offers a pronounced tactile bump with a medium actuation force, delivering satisfying feedback for those who prefer a more noticeable tactile response. This switch is ideal for typists seeking tactile feedback without the noise of a clicky switch.",
					ManufacturerID:   ptr(11), // Glorious
					BrandID:          ptr(11), // Glorious
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2023-09-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
			}
			err = processSwitches(gloriousSwitches)
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

			if err := tx.Model(&Producer{}).Where("alias IN (?)", []string{"cannonkeys", "drop", "glorious"}).Pluck("id", &producersIDs).Error; err != nil {
				return err
			}

			if err := tx.Exec("DELETE FROM switches WHERE brand_id IN (?)", producersIDs).Error; err != nil {
				return err
			}
			return nil // Replace with actual code
		},
	})
}
