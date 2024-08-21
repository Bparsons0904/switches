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

			keydous := Producer{
				Name:    "Cannon Keys",
				Alias:   "cannonkeys",
				SiteURL: "",
			}

			if err := tx.Create(&keydous).Error; err != nil {
				return err
			}

			cannonKeysSwitches := []Switch{
				{
					Name:             "CannonKeys Lavenders",
					ShortDescription: "Smooth linear switch with a unique lavender color.",
					LongDescription:  "CannonKeys Lavenders are linear switches known for their smooth feel and light actuation force. They feature a unique lavender-colored housing and stem, providing a visually distinct option for mechanical keyboard enthusiasts.",
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
					LongDescription:  "The CannonKeys Anubis switch offers a tactile typing experience with a medium actuation force and a satisfying bump. It is known for its smooth feel and rich sound profile, making it a favorite among tactile switch enthusiasts.",
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
					LongDescription:  "Dragonfruit switches by CannonKeys are tactile switches that offer a prominent tactile bump and a delightful typing experience. The switches are named for their vibrant, exotic color reminiscent of dragon fruit.",
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
					LongDescription:  "CannonKeys Banana Split switches are known for their ultra-smooth linear action and fun, vibrant colors. These switches have a pleasant feel and offer a delightful typing experience, making them a popular choice for custom keyboards.",
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
					LongDescription:  "The CannonKeys Durock POM switch is a linear switch that features a POM (polyoxymethylene) housing, offering a smooth, scratch-free experience. Its construction makes it highly durable and reliable for heavy typing sessions.",
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
					LongDescription:  "The CannonKeys Blueberry switch is a clicky switch known for its tactile bump and crisp sound. It offers a unique typing experience with a satisfying audible click.",
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
					LongDescription:  "Orange Pies by CannonKeys are clicky switches offering a strong tactile feel paired with a sharp sound. These switches are perfect for those who enjoy a more pronounced typing feedback.",
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
					LongDescription:  "CannonKeys Matcha Latte switches provide a creamy smooth linear action with a light actuation force. Their unique matcha color and satisfying smoothness make them a favorite among linear switch enthusiasts.",
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
					LongDescription:  "CannonKeys Taro Balls switches are linear switches with a medium actuation force and quiet operation. They offer a smooth typing experience with a unique purple color inspired by taro.",
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
					LongDescription:  "CannonKeys Crystal switches offer a smooth linear action with a transparent housing, allowing RGB lighting to shine through. They provide a clean and consistent typing experience with a touch of elegance.",
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
					LongDescription:  "CannonKeys Silent Alpaca V2 switches offer a silent linear experience with a smooth feel and muted sound, making them perfect for quiet typing environments. The switch features a unique housing designed to minimize noise while maintaining smoothness.",
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
					LongDescription:  "CannonKeys Boba U4 switches are silent tactile switches known for their soft bump and quiet operation. They provide a satisfying typing experience without the noise, making them ideal for shared spaces.",
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
					LongDescription:  "The CannonKeys Mint Chocolate switch is a specialty switch featuring a mint-themed color scheme and smooth linear action. Its unique aesthetics and pleasant typing feel make it a favorite among custom keyboard enthusiasts.",
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
					LongDescription:  "The CannonKeys Snow Globe switch offers a specialty linear action with a frosted housing, allowing for a unique aesthetic and smooth typing feel. Its winter-themed design makes it a favorite among custom keyboard builders.",
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
					LongDescription:  "CannonKeys Milkshake switches provide a milky smooth linear action, offering a unique tactile experience with a focus on smoothness and consistency. Ideal for those who prefer a straightforward and satisfying typing feel.",
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
					LongDescription:  "Drop Halo True switches are designed for a smooth tactile experience with a significant tactile bump. They offer minimal noise, making them suitable for quieter environments while still providing satisfying feedback.",
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
					LongDescription:  "The Drop Holy Panda switch is renowned for its distinctive tactile bump and satisfying sound profile. It combines components from the Halo True stem and Invyr Panda housing, resulting in a unique feel and excellent build quality.",
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
					LongDescription:  "Invyr Holy Panda switches by Drop offer the classic Holy Panda experience with their remarkable tactile bump and sound. These switches are a favorite among enthusiasts who enjoy tactile typing.",
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
					LongDescription:  "Drop ALT Rock Black switches are linear switches that offer a heavier actuation force for those who prefer more resistance. They deliver a smooth typing experience and are popular among gaming enthusiasts.",
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
					LongDescription:  "The Drop Silver Speed switch is engineered for rapid actuation, making it perfect for gaming and fast typing. Its smooth, linear action ensures a seamless experience, reducing input lag for gamers.",
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
					LongDescription:  "Drop Kaihua Box White switches are designed with a box structure to prevent dust and moisture. They offer a light actuation force and a satisfying click, making them popular among clicky switch enthusiasts.",
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
					LongDescription:  "The Drop Kaihua Speed Silver switch is a linear switch engineered for rapid actuation, offering a swift response time ideal for gaming. Its smooth action ensures seamless performance for fast-paced scenarios.",
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
					LongDescription:  "Drop Kaihua Jade switches feature a thicker click bar, resulting in a louder and more satisfying click sound. They are perfect for those who want a pronounced click with tactile feedback.",
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
					LongDescription:  "Drop Halo Linear switches provide a smooth linear feel with minimal resistance, making them ideal for those who prefer a straightforward typing experience without tactile bumps.",
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
					LongDescription:  "Drop Silent Aliaz switches are designed to offer a quiet tactile experience with a soft bump and minimal noise. They are perfect for typists who want tactile feedback without disturbing others.",
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
					LongDescription:  "The Drop Kaihua Silent Pink switch offers a silent linear action with a fast actuation force. Its vibrant pink color and quiet operation make it a standout choice for those who value aesthetics and silence.",
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
					LongDescription:  "The Drop Emerald Silent Tactile switch provides a gentle tactile bump with a quiet operation. Its emerald-themed design adds a touch of elegance to any keyboard setup while maintaining a peaceful typing experience.",
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
					LongDescription:  "Drop Holy Boba switches combine elements from both Holy Panda and Boba U4 switches, resulting in a unique tactile experience with a smooth bump and premium feel. Perfect for enthusiasts who appreciate tactile feedback and innovative design.",
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
					LongDescription:  "Glorious Panda switches are highly tactile switches known for their crisp bump and smooth travel. They are designed to offer a satisfying typing experience with a notable tactile feedback.",
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
					LongDescription:  "The Glorious Lynx switch is a linear switch designed for gaming enthusiasts seeking a smooth and fast feel. It provides a consistent actuation experience, ideal for quick response times in gaming scenarios.",
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
					LongDescription:  "Glorious Fox switches are linear switches that offer a medium actuation force, delivering a balanced typing experience suitable for both gaming and everyday typing.",
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
					LongDescription:  "The Glorious Ice switch offers a smooth linear action with quiet operation. It is designed to provide an 'icy' feel, making it ideal for those who prefer a more silent typing experience.",
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
					LongDescription:  "The Glorious Holy Panda switch combines the best of tactile feel with a distinct bump and premium build quality. It is a popular choice for enthusiasts seeking an elevated typing experience.",
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
					LongDescription:  "Glorious Goat Clicky switches provide a medium actuation force and a crisp click sound. The switch's distinctive goat logo and sound profile make it a unique addition to any keyboard setup.",
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
					LongDescription:  "The Glorious Fox Clicky switch is known for its smooth action and clicky feedback. It is designed to offer a balance between tactile feel and audible satisfaction, making it ideal for typists who enjoy a clicky experience.",
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
					LongDescription:  "Glorious Lynx Clicky switches provide a light actuation force, making them suitable for fast typing and gaming. The clicky sound enhances the typing experience, offering both auditory and tactile feedback.",
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
					LongDescription:  "Glorious Panda Linear switches offer a smooth and consistent feel, providing a premium typing experience for those who prefer linear switches. They are designed for both gaming and everyday use.",
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
					LongDescription:  "The Glorious Fox Linear V2 switch offers a light actuation force and refined design, providing a seamless typing experience with enhanced smoothness and aesthetics.",
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
					LongDescription:  "Glorious Lynx Linear V2 switches feature improved smoothness and sound, offering a refined typing experience for those who prefer a quieter and more fluid action.",
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
					LongDescription:  "The Glorious Ice Linear V2 switch offers quiet operation with icy aesthetics, providing a cool and smooth typing experience that is both silent and visually appealing.",
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
					LongDescription:  "Glorious Panda Silent switches offer a quiet tactile bump with a smooth feel, providing a premium typing experience without the noise. They are perfect for those who enjoy tactile feedback without disturbing others.",
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
					LongDescription:  "The Glorious Lynx Silent switch offers a silent linear action with a soft feel, providing a quiet typing experience perfect for offices and shared spaces. Its understated aesthetics and smooth operation make it a versatile choice.",
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
					LongDescription:  "The Glorious Crystal Ice switch offers an icy smooth linear action with a transparent housing, allowing RGB lighting to shine through. Its unique design and consistent feel make it a popular choice for custom builds.",
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
					LongDescription:  "The Glorious Aurora switch offers enhanced tactile feedback and RGB diffusion, providing both visual appeal and satisfying typing experience. Its innovative design makes it a standout choice for enthusiasts who value aesthetics and performance.",
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
					LongDescription:  "The Glorious Summer Sun switch offers a bright design with smooth linear action, capturing the essence of summer in both aesthetics and typing experience. It's a popular choice for those who love vibrant keyboard setups.",
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
					LongDescription:  "The Glorious Bobcat switch offers a tactile typing experience with a pronounced bump and medium actuation force. It provides a satisfying feedback for those who prefer a more noticeable tactile response without the noise of a clicky switch.",
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
