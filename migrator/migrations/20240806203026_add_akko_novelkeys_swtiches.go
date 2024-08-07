package migrations

import (
	"time"

	"github.com/Bparsons0904/deadigations"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	deadigations.RegisterMigration(deadigations.Migration{
		ID:          "20240806203026",
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

			akkoSwitches := []Switch{
				{
					Name:             "Akko CS Jelly Pink",
					ShortDescription: "Linear switch with a smooth feel and moderate actuation force.",
					LongDescription:  "The Akko CS Jelly Pink switch is a linear switch that offers a smooth typing experience with a moderate actuation force. It's known for its durability and consistent performance, making it a popular choice among enthusiasts.",
					ManufacturerID:   ptr(8), // Akko
					BrandID:          ptr(8), // Akko
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-05-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				{
					Name:             "Akko CS Ocean Blue",
					ShortDescription: "Tactile switch with a light tactile bump and smooth operation.",
					LongDescription:  "The Akko CS Ocean Blue switch features a light tactile bump that offers a distinct typing experience. With its smooth operation and moderate actuation force, it's ideal for users who enjoy tactile feedback without the clicky sound.",
					ManufacturerID:   ptr(8), // Akko
					BrandID:          ptr(8), // Akko
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-05-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Akko CS Lavender Purple",
					ShortDescription: "Tactile switch with a pronounced bump and higher actuation force.",
					LongDescription:  "The Akko CS Lavender Purple switch offers a pronounced tactile bump with a higher actuation force, providing a satisfying typing experience for users who prefer a more tactile feel. It's known for its durability and consistent performance.",
					ManufacturerID:   ptr(8), // Akko
					BrandID:          ptr(8), // Akko
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-05-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Akko CS Rose Red",
					ShortDescription: "Linear switch with a light actuation force for fast typing.",
					LongDescription:  "The Akko CS Rose Red switch is a linear switch designed for fast typing with a light actuation force. Its smooth and quiet operation makes it a favorite among gamers and typists who prefer a more silent experience.",
					ManufacturerID:   ptr(8), // Akko
					BrandID:          ptr(8), // Akko
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-05-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				{
					Name:             "Akko CS Radiant Red",
					ShortDescription: "Linear switch with a heavier actuation force for deliberate keystrokes.",
					LongDescription:  "The Akko CS Radiant Red switch is a linear switch that provides a heavier actuation force, perfect for users who prefer a more deliberate typing experience. It's praised for its smoothness and durability.",
					ManufacturerID:   ptr(8), // Akko
					BrandID:          ptr(8), // Akko
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-05-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Akko CS Matcha Green",
					ShortDescription: "Tactile switch with a medium tactile bump and balanced actuation force.",
					LongDescription:  "The Akko CS Matcha Green switch is known for its medium tactile bump, offering a balanced actuation force that's ideal for both typing and gaming. Its unique green color adds a touch of style to any keyboard.",
					ManufacturerID:   ptr(8), // Akko
					BrandID:          ptr(8), // Akko
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-05-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Akko CS Crystal",
					ShortDescription: "Silent linear switch with smooth keystrokes and minimal noise.",
					LongDescription:  "The Akko CS Crystal switch is a silent linear switch that offers smooth keystrokes with minimal noise, making it perfect for office settings and quiet environments. Its transparent housing allows for RGB lighting to shine through beautifully.",
					ManufacturerID:   ptr(8), // Akko
					BrandID:          ptr(8), // Akko
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2021-06-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Akko CS Midnight Blue",
					ShortDescription: "Clicky switch with a deep, satisfying click and firm actuation force.",
					LongDescription:  "The Akko CS Midnight Blue switch provides a clicky typing experience with a deep and satisfying click sound. It offers a firm actuation force, making it suitable for users who prefer a more deliberate and tactile feel.",
					ManufacturerID:   ptr(8), // Akko
					BrandID:          ptr(8), // Akko
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2021-03-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Akko CS Blue",
					ShortDescription: "Traditional clicky switch with a sharp click sound and tactile bump.",
					LongDescription:  "The Akko CS Blue switch is a classic clicky switch that offers a sharp click sound and a noticeable tactile bump. It's a great option for users who enjoy a traditional clicky typing experience with a bit of resistance.",
					ManufacturerID:   ptr(8), // Akko
					BrandID:          ptr(8), // Akko
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2021-05-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				{
					Name:             "Akko CS Silver",
					ShortDescription: "Clicky switch with a bright sound and quick actuation.",
					LongDescription:  "The Akko CS Silver switch is a clicky switch with a bright and sharp sound. Its quick actuation makes it perfect for fast typists who enjoy a clicky feel without too much resistance.",
					ManufacturerID:   ptr(8), // Akko
					BrandID:          ptr(8), // Akko
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2021-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Akko CS Neon",
					ShortDescription: "Colorful clicky switch with a unique design and clear sound.",
					LongDescription:  "The Akko CS Neon switch offers a colorful and unique design along with a clear clicky sound. Its vibrant appearance makes it a fun addition to any keyboard setup, while its tactile bump provides a satisfying typing experience.",
					ManufacturerID:   ptr(8), // Akko
					BrandID:          ptr(8), // Akko
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2021-07-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Akko CS Sponge",
					ShortDescription: "Tactile switch with a gentle bump and soft feedback.",
					LongDescription:  "The Akko CS Sponge switch offers a gentle tactile bump with soft feedback, making it ideal for users who prefer a quieter and more subtle typing experience. Its unique design and feel make it a popular choice among enthusiasts.",
					ManufacturerID:   ptr(8), // Akko
					BrandID:          ptr(8), // Akko
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-04-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Akko CS Starfish",
					ShortDescription: "Tactile switch with a soft bump and light actuation force.",
					LongDescription:  "The Akko CS Starfish switch features a soft tactile bump with a light actuation force, offering a comfortable typing experience. Its smooth operation and unique design make it a favorite among those who appreciate a lighter tactile feel.",
					ManufacturerID:   ptr(8), // Akko
					BrandID:          ptr(8), // Akko
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-05-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Akko CS Orange",
					ShortDescription: "Tactile switch with a pronounced bump and medium actuation force.",
					LongDescription:  "The Akko CS Orange switch is known for its pronounced tactile bump and medium actuation force, providing a satisfying typing experience for those who prefer more noticeable feedback. It's praised for its durability and consistent performance.",
					ManufacturerID:   ptr(8), // Akko
					BrandID:          ptr(8), // Akko
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-06-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				{
					Name:             "Akko CS Green",
					ShortDescription: "Tactile switch with a crisp bump and responsive actuation.",
					LongDescription:  "The Akko CS Green switch offers a crisp tactile bump with responsive actuation, making it a great choice for users who enjoy a tactile typing experience with quick feedback. Its robust design ensures long-lasting performance.",
					ManufacturerID:   ptr(8), // Akko
					BrandID:          ptr(8), // Akko
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-07-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Akko CS Silent",
					ShortDescription: "Silent tactile switch with smooth operation and quiet feedback.",
					LongDescription:  "The Akko CS Silent switch is a silent tactile switch that provides smooth operation and quiet feedback, making it perfect for office settings and quiet environments. Its silent design ensures a pleasant typing experience without sacrificing tactility.",
					ManufacturerID:   ptr(8), // Akko
					BrandID:          ptr(8), // Akko
					SwitchTypeID:     5,      // Silent Tactile
					ReleaseDate:      parseDate("2021-08-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Akko CS Peach",
					ShortDescription: "Linear switch with a soft feel and smooth keystrokes.",
					LongDescription:  "The Akko CS Peach switch offers a soft linear feel with smooth keystrokes, making it an excellent choice for typists and gamers who prefer a quiet and fluid typing experience.",
					ManufacturerID:   ptr(8), // Akko
					BrandID:          ptr(8), // Akko
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-03-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Akko CS Cream Yellow",
					ShortDescription: "Linear switch with a buttery smooth action and light actuation force.",
					LongDescription:  "The Akko CS Cream Yellow switch is known for its buttery smooth action and light actuation force, providing a pleasant typing experience that's both fast and effortless. Its reliable performance makes it a favorite among linear enthusiasts.",
					ManufacturerID:   ptr(8), // Akko
					BrandID:          ptr(8), // Akko
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-05-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				{
					Name:             "Akko CS Mint Green",
					ShortDescription: "Linear switch with a crisp feel and moderate actuation force.",
					LongDescription:  "The Akko CS Mint Green switch offers a crisp linear feel with a moderate actuation force, perfect for users who enjoy a balanced typing experience. Its unique mint green color adds a touch of style to any keyboard.",
					ManufacturerID:   ptr(8), // Akko
					BrandID:          ptr(8), // Akko
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Akko CS Raspberry",
					ShortDescription: "Linear switch with a vibrant feel and responsive actuation.",
					LongDescription:  "The Akko CS Raspberry switch provides a vibrant linear feel with responsive actuation, ideal for fast typists and gamers. Its consistent performance and lively design make it a popular choice for those who appreciate a dynamic typing experience.",
					ManufacturerID:   ptr(8), // Akko
					BrandID:          ptr(8), // Akko
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-07-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Akko CS V2 Matcha Green",
					ShortDescription: "Linear switch with enhanced smoothness and stability.",
					LongDescription:  "The Akko CS V2 Matcha Green switch features enhanced smoothness and stability, offering a refined linear typing experience. Its improved design ensures long-lasting performance and a consistent feel throughout each keystroke.",
					ManufacturerID:   ptr(8), // Akko
					BrandID:          ptr(8), // Akko
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2022-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Akko CS Radiant Red Pro",
					ShortDescription: "Linear switch with a quick actuation and deep sound.",
					LongDescription:  "The Akko CS Radiant Red Pro switch is a linear switch that offers quick actuation with a deep sound, providing a satisfying typing experience. Its unique combination of speed and sound makes it a great choice for both gamers and typists.",
					ManufacturerID:   ptr(8), // Akko
					BrandID:          ptr(8), // Akko
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2022-02-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Akko CS Silent Pink",
					ShortDescription: "Silent linear switch with a smooth and quiet operation.",
					LongDescription:  "The Akko CS Silent Pink switch is designed for those who require a silent typing experience without sacrificing performance. Its smooth linear action and quiet operation make it ideal for office settings and environments where noise is a concern.",
					ManufacturerID:   ptr(8), // Akko
					BrandID:          ptr(8), // Akko
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2021-09-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Akko CS Silent Jelly Purple",
					ShortDescription: "Silent tactile switch with a subtle bump and quiet feedback.",
					LongDescription:  "The Akko CS Silent Jelly Purple switch offers a silent tactile experience with a subtle bump and quiet feedback, making it a perfect choice for users who enjoy tactility without the noise.",
					ManufacturerID:   ptr(8), // Akko
					BrandID:          ptr(8), // Akko
					SwitchTypeID:     5,      // Silent Tactile
					ReleaseDate:      parseDate("2021-10-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Akko CS Purple",
					ShortDescription: "Linear switch with a lightweight actuation and smooth feel.",
					LongDescription:  "The Akko CS Purple switch is a linear switch known for its lightweight actuation and smooth feel. It's a versatile switch that works well for both typing and gaming, offering a consistent and reliable performance.",
					ManufacturerID:   ptr(8), // Akko
					BrandID:          ptr(8), // Akko
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-04-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Akko V3 Cream Blue",
					ShortDescription: "Linear switch with enhanced smoothness and stability.",
					LongDescription:  "The Akko V3 Cream Blue switch is an advanced linear switch that offers enhanced smoothness and stability, providing a premium typing experience with refined performance.",
					ManufacturerID:   ptr(8), // Akko
					BrandID:          ptr(8), // Akko
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2022-03-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Akko CS Silent Forest Green",
					ShortDescription: "Silent linear switch with a medium actuation force.",
					LongDescription:  "The Akko CS Silent Forest Green switch combines a silent linear action with a medium actuation force, offering a balanced and quiet typing experience suitable for various applications.",
					ManufacturerID:   ptr(8), // Akko
					BrandID:          ptr(8), // Akko
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2021-11-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Akko CS Cherry",
					ShortDescription: "Linear switch with a clear housing for RGB lighting.",
					LongDescription:  "The Akko CS Cherry switch is a linear switch featuring a clear housing that allows for vibrant RGB lighting. Its smooth operation and aesthetic appeal make it a popular choice for customized keyboard builds.",
					ManufacturerID:   ptr(8), // Akko
					BrandID:          ptr(8), // Akko
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-12-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Akko Kingfisher",
					ShortDescription: "Tactile switch with a vibrant feel and moderate bump.",
					LongDescription:  "The Akko Kingfisher switch is a tactile switch known for its vibrant feel and moderate tactile bump, providing a satisfying typing experience with a unique colorway inspired by the Kingfisher bird.",
					ManufacturerID:   ptr(8), // Akko
					BrandID:          ptr(8), // Akko
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2020-02-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Akko Sakura Pink",
					ShortDescription: "Linear switch with a soft touch and low actuation force.",
					LongDescription:  "The Akko Sakura Pink switch offers a soft touch and low actuation force, making it ideal for typists who prefer a gentle and effortless typing experience. Its aesthetic design is inspired by cherry blossoms.",
					ManufacturerID:   ptr(8), // Akko
					BrandID:          ptr(8), // Akko
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-08-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Akko Monet Purple",
					ShortDescription: "Specialty switch with artistic design and smooth feel.",
					LongDescription:  "The Akko Monet Purple switch is a specialty switch designed with an artistic touch, offering a smooth linear feel that mimics the soft strokes of Monet's paintings. It's a perfect addition for those who appreciate aesthetics in their keyboard.",
					ManufacturerID:   ptr(8), // Akko
					BrandID:          ptr(8), // Akko
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2022-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Akko CS Piano",
					ShortDescription: "Linear switch with a smooth keystroke and balanced actuation force.",
					LongDescription:  "The Akko CS Piano switch offers a linear feel with a smooth keystroke and balanced actuation force, providing a typing experience reminiscent of playing a well-tuned piano. Its consistent performance makes it suitable for both typing and gaming.",
					ManufacturerID:   ptr(8),                  // Akko
					BrandID:          ptr(8),                  // Akko
					SwitchTypeID:     1,                       // Linear
					ReleaseDate:      parseDate("2022-06-01"), // Approximate date, please verify
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Akko CS V3 Cream Yellow",
					ShortDescription: "Enhanced linear switch with improved smoothness and reduced wobble.",
					LongDescription:  "The Akko CS V3 Cream Yellow is an upgraded version of the popular Cream Yellow switch, featuring improved smoothness and reduced stem wobble. It offers a refined typing experience with a light actuation force, making it ideal for extended typing or gaming sessions.",
					ManufacturerID:   ptr(8),                  // Akko
					BrandID:          ptr(8),                  // Akko
					SwitchTypeID:     1,                       // Linear
					ReleaseDate:      parseDate("2022-09-01"), // Approximate date, please verify
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Akko CS Vintage White",
					ShortDescription: "Linear switch with a classic feel and moderate actuation force.",
					LongDescription:  "The Akko CS Vintage White switch provides a linear feel with a classic touch, offering a moderate actuation force that caters to users who prefer a balanced typing experience. Its vintage-inspired design adds a touch of nostalgia to modern keyboards.",
					ManufacturerID:   ptr(8),                  // Akko
					BrandID:          ptr(8),                  // Akko
					SwitchTypeID:     1,                       // Linear
					ReleaseDate:      parseDate("2022-07-01"), // Approximate date, please verify
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Akko CS Maroon",
					ShortDescription: "Tactile switch with a deep, satisfying bump and rich color.",
					LongDescription:  "The Akko CS Maroon switch is a tactile switch that offers a deep and satisfying tactile bump. Its rich maroon color adds a touch of elegance to any keyboard, while its tactile feedback provides a gratifying typing experience for enthusiasts who enjoy pronounced feedback.",
					ManufacturerID:   ptr(8),                  // Akko
					BrandID:          ptr(8),                  // Akko
					SwitchTypeID:     2,                       // Tactile
					ReleaseDate:      parseDate("2022-08-01"), // Approximate date, please verify
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Akko CS Fly Wing",
					ShortDescription: "Ultra-light linear switch for fast typing and gaming.",
					LongDescription:  "The Akko CS Fly Wing switch is an ultra-light linear switch designed for fast typing and gaming. Its low actuation force and smooth travel make it perfect for users who prioritize speed and responsiveness in their keyboard switches.",
					ManufacturerID:   ptr(8),                  // Akko
					BrandID:          ptr(8),                  // Akko
					SwitchTypeID:     1,                       // Linear
					ReleaseDate:      parseDate("2022-10-01"), // Approximate date, please verify
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Akko CS Jet Black",
					ShortDescription: "Linear switch with a sleek design and smooth keystroke.",
					LongDescription:  "The Akko CS Jet Black switch is a linear switch that combines a sleek, all-black design with a smooth keystroke. It offers a premium typing experience with a moderate actuation force, suitable for users who appreciate both aesthetics and performance in their keyboard switches.",
					ManufacturerID:   ptr(8),                  // Akko
					BrandID:          ptr(8),                  // Akko
					SwitchTypeID:     1,                       // Linear
					ReleaseDate:      parseDate("2022-11-01"), // Approximate date, please verify
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

			err := processSwitches(akkoSwitches)
			if err != nil {
				return err
			}

			novelKeysSwitches := []Switch{
				{
					Name:             "NovelKeys Cream",
					ShortDescription: "Smooth linear switch with a POM housing.",
					LongDescription:  "The NovelKeys Cream switch is a highly regarded linear switch known for its smooth action and unique POM housing. Its buttery feel and thocky sound make it a favorite among keyboard enthusiasts.",
					ManufacturerID:   ptr(9), // NovelKeys
					BrandID:          ptr(9), // NovelKeys
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2019-05-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "NovelKeys Cream Tactile",
					ShortDescription: "Tactile switch with a smooth bump and POM housing.",
					LongDescription:  "The NovelKeys Cream Tactile switch combines the smoothness of the original Cream switch with a tactile bump, featuring a full POM housing for a unique tactile experience with a satisfying thocky sound.",
					ManufacturerID:   ptr(9), // NovelKeys
					BrandID:          ptr(9), // NovelKeys
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2020-07-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "NovelKeys x Kailh Box Jade",
					ShortDescription: "Clicky switch with a crisp and loud click.",
					LongDescription:  "The NovelKeys x Kailh Box Jade switch is a clicky switch renowned for its crisp and loud click, offering a satisfying typing experience for those who love a pronounced auditory feedback.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(9), // NovelKeys
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2018-04-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "NovelKeys x Kailh Box Royal",
					ShortDescription: "Tactile switch with a heavy bump and smooth action.",
					LongDescription:  "The NovelKeys x Kailh Box Royal switch is a tactile switch that features a heavy bump and smooth action. It's perfect for typists who prefer strong tactile feedback and a satisfying keystroke.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(9), // NovelKeys
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2018-04-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "NovelKeys Sherbet",
					ShortDescription: "Linear switch with a light actuation force and smooth keystrokes.",
					LongDescription:  "The NovelKeys Sherbet switch offers a light actuation force and smooth keystrokes, providing a gentle and enjoyable typing experience. Its unique color scheme adds a playful touch to any keyboard setup.",
					ManufacturerID:   ptr(9), // NovelKeys
					BrandID:          ptr(9), // NovelKeys
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-11-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "NovelKeys Silk Yellow",
					ShortDescription: "Silky smooth linear switch with a light actuation force.",
					LongDescription:  "The NovelKeys Silk Yellow switch is a linear switch that offers a silky smooth typing experience with a light actuation force, making it perfect for fast typists who enjoy a smooth and effortless keystroke.",
					ManufacturerID:   ptr(9), // NovelKeys
					BrandID:          ptr(9), // NovelKeys
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-02-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "NovelKeys Silk Black",
					ShortDescription: "Linear switch with a heavier actuation force and smooth feel.",
					LongDescription:  "The NovelKeys Silk Black switch offers a linear typing experience with a heavier actuation force, providing a smooth and satisfying keystroke. It's designed for users who prefer more resistance in their switches.",
					ManufacturerID:   ptr(9), // NovelKeys
					BrandID:          ptr(9), // NovelKeys
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-02-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "NovelKeys Silk Red",
					ShortDescription: "Smooth linear switch with a light actuation force.",
					LongDescription:  "The NovelKeys Silk Red switch is a smooth linear switch with a light actuation force, offering a fast and fluid typing experience. It's ideal for gamers and typists who enjoy quick key presses.",
					ManufacturerID:   ptr(9), // NovelKeys
					BrandID:          ptr(9), // NovelKeys
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-02-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "NovelKeys x Kailh Box Pale Blue",
					ShortDescription: "Clicky switch with a bright sound and crisp feel.",
					LongDescription:  "The NovelKeys x Kailh Box Pale Blue switch offers a bright and crisp clicky sound, providing a tactile and auditory experience that's satisfying for typists who enjoy a pronounced click.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(9), // NovelKeys
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2018-04-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "NovelKeys x Kailh Box Red",
					ShortDescription: "Linear switch with a light feel and smooth action.",
					LongDescription:  "The NovelKeys x Kailh Box Red switch is a linear switch known for its light feel and smooth action, making it a popular choice for both gamers and typists who prefer a fluid keystroke.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(9), // NovelKeys
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2018-04-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "NovelKeys x Kailh Box Black",
					ShortDescription: "Linear switch with a heavy actuation force and smooth feel.",
					LongDescription:  "The NovelKeys x Kailh Box Black switch offers a linear typing experience with a heavy actuation force, providing a satisfying keystroke for users who prefer more resistance in their switches.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(9), // NovelKeys
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2018-04-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "NovelKeys x Kailh Box Brown",
					ShortDescription: "Tactile switch with a soft bump and smooth action.",
					LongDescription:  "The NovelKeys x Kailh Box Brown switch is a tactile switch that offers a soft bump and smooth action, providing a balanced typing experience for those who enjoy subtle tactile feedback.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(9), // NovelKeys
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2018-04-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "NovelKeys x Kailh Box Pink",
					ShortDescription: "Clicky switch with a unique pink color and crisp feel.",
					LongDescription:  "The NovelKeys x Kailh Box Pink switch offers a unique aesthetic with its pink color and crisp clicky feel, providing both visual appeal and a satisfying auditory experience.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(9), // NovelKeys
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2018-04-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "NovelKeys x Kailh Box Navy",
					ShortDescription: "Heavy clicky switch with a deep, satisfying click.",
					LongDescription:  "The NovelKeys x Kailh Box Navy switch is a heavy clicky switch that offers a deep and satisfying click, perfect for typists who crave a pronounced tactile and auditory experience.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(9), // NovelKeys
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2018-04-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "NovelKeys x Kailh Box Heavy Pale Blue",
					ShortDescription: "Heavy clicky switch with a strong tactile bump and crisp click.",
					LongDescription:  "The NovelKeys x Kailh Box Heavy Pale Blue switch is designed for those who crave a strong tactile bump and crisp click sound, offering a satisfying and engaging typing experience with a heavier feel.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(9), // NovelKeys
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2018-04-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "NovelKeys Blueberry",
					ShortDescription: "Tactile switch with a firm bump and distinct feedback.",
					LongDescription:  "The NovelKeys Blueberry switch offers a firm tactile bump with distinct feedback, providing a unique typing experience with a pronounced and satisfying keystroke, ideal for those who enjoy a strong tactile sensation.",
					ManufacturerID:   ptr(9), // NovelKeys
					BrandID:          ptr(9), // NovelKeys
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2019-12-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "NovelKeys Box Royal",
					ShortDescription: "Tactile clicky switch with a heavy bump and satisfying click.",
					LongDescription:  "The NovelKeys Box Royal switch is a tactile clicky switch that offers a heavy bump and satisfying click, ideal for those who enjoy strong tactile feedback and a pronounced auditory experience.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(9), // NovelKeys
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2018-04-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "NovelKeys Sherbet Tactile",
					ShortDescription: "Tactile switch with a light bump and colorful design.",
					LongDescription:  "The NovelKeys Sherbet Tactile switch offers a light tactile bump and a colorful design, providing a playful typing experience with a unique and enjoyable feel, perfect for those who appreciate a softer tactile switch.",
					ManufacturerID:   ptr(9), // NovelKeys
					BrandID:          ptr(9), // NovelKeys
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2020-11-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "NovelKeys T1",
					ShortDescription: "Tactile switch with a pronounced bump and consistent feel.",
					LongDescription:  "The NovelKeys T1 switch is known for its pronounced tactile bump and consistent feel, offering a satisfying typing experience with a strong and distinctive tactile sensation, suitable for those who enjoy pronounced feedback.",
					ManufacturerID:   ptr(9), // NovelKeys
					BrandID:          ptr(9), // NovelKeys
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2019-10-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "NovelKeys x Kailh Cream",
					ShortDescription: "Linear switch with a unique feel and smooth action.",
					LongDescription:  "The NovelKeys x Kailh Cream switch is a collaboration with Kailh, offering a unique linear feel and smooth action with a POM housing for a thocky sound and consistent performance.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(9), // NovelKeys
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2019-05-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "NovelKeys Dry Yellow",
					ShortDescription: "Linear switch with a dry feel and quick actuation.",
					LongDescription:  "The NovelKeys Dry Yellow switch provides a unique linear feel with quick actuation, offering a dry and satisfying keystroke for those who prefer a different linear experience.",
					ManufacturerID:   ptr(9), // NovelKeys
					BrandID:          ptr(9), // NovelKeys
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-09-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "NovelKeys Cream Clickies",
					ShortDescription: "Unique clicky switch with a POM housing.",
					LongDescription:  "The NovelKeys Cream Clickies switch combines the smooth POM housing of the Cream line with a clicky mechanism, providing a unique auditory and tactile experience that sets it apart from standard clicky switches.",
					ManufacturerID:   ptr(9), // NovelKeys
					BrandID:          ptr(9), // NovelKeys
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2020-12-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "NovelKeys Dry Black",
					ShortDescription: "Linear switch with a dry feel and heavier actuation.",
					LongDescription:  "The NovelKeys Dry Black switch offers a dry feel and heavier actuation force, providing a distinct linear experience that differs from typical smooth switches, appealing to users who prefer a more solid and resistant keystroke.",
					ManufacturerID:   ptr(9), // NovelKeys
					BrandID:          ptr(9), // NovelKeys
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-09-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "NovelKeys NK_ Blueberry",
					ShortDescription: "Tactile switch with a unique tactile event.",
					LongDescription:  "The NovelKeys NK_ Blueberry switch provides a distinct tactile event with a strong bump, offering a different tactile experience that stands out for those who appreciate a firm and pronounced tactile feel.",
					ManufacturerID:   ptr(9), // NovelKeys
					BrandID:          ptr(9), // NovelKeys
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2020-11-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "NovelKeys x Gateron Silent Black",
					ShortDescription: "Silent linear switch with a smooth feel and heavy actuation.",
					LongDescription:  "The NovelKeys x Gateron Silent Black switch offers a smooth linear action with a heavier actuation force, providing a silent and solid typing experience perfect for office environments or quiet settings.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(9), // NovelKeys
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2019-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "NovelKeys x Gateron Silent Yellow",
					ShortDescription: "Silent linear switch with a light feel and smooth action.",
					LongDescription:  "The NovelKeys x Gateron Silent Yellow switch provides a smooth linear experience with a light actuation force, offering a silent keystroke ideal for those who appreciate a quiet and effortless typing session.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(9), // NovelKeys
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2019-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "NovelKeys Cream Speed",
					ShortDescription: "Speed linear switch with a POM housing and rapid actuation.",
					LongDescription:  "The NovelKeys Cream Speed switch offers rapid actuation and a smooth linear feel with a POM housing, designed for fast typists and gamers who demand quick response times and a seamless typing experience.",
					ManufacturerID:   ptr(9), // NovelKeys
					BrandID:          ptr(9), // NovelKeys
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-03-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "NovelKeys x Kailh Clickiez",
					ShortDescription: "Clicky switch with a unique mechanism and tactile feedback.",
					LongDescription:  "The NovelKeys x Kailh Clickiez switch combines a clicky mechanism with tactile feedback, providing a unique typing experience that offers both auditory and tactile satisfaction, making it a standout choice for enthusiasts.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(9), // NovelKeys
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2021-05-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "NovelKeys x Kailh Speed Navy",
					ShortDescription: "Clicky speed switch with rapid actuation and strong click.",
					LongDescription:  "The NovelKeys x Kailh Speed Navy switch offers rapid actuation with a strong clicky sound, designed for gamers and fast typists who require quick response times and a satisfying auditory experience.",
					ManufacturerID:   ptr(3), // Kailh
					BrandID:          ptr(9), // NovelKeys
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2021-05-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "NovelKeys Eggnog",
					ShortDescription: "Linear switch with a smooth feel and unique colorway.",
					LongDescription:  "The NovelKeys Eggnog switch offers a smooth linear feel and a unique colorway inspired by the holiday treat, providing a festive and enjoyable typing experience for those who appreciate both performance and aesthetics.",
					ManufacturerID:   ptr(9), // NovelKeys
					BrandID:          ptr(9), // NovelKeys
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-12-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
			}

			err = processSwitches(novelKeysSwitches)
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

			if err := tx.Model(&Producer{}).Where("alias IN (?)", []string{"akko", "novelkeys"}).Pluck("id", &producersIDs).Error; err != nil {
				return err
			}

			if err := tx.Exec("DELETE FROM switches WHERE brand_id IN (?)", producersIDs).Error; err != nil {
				return err
			}
			return nil // Replace with actual code
		},
	})
}
