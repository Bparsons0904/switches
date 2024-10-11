package migrations_old

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

			ptr := func(i int) *int {
				return &i
			}

			parseDate := func(date string) *time.Time {
				t, _ := time.Parse("2006-01-02", date)
				return &t
			}

			ttcSwitches := []Switch{
				{
					Name:             "TTC Gold Pink",
					ShortDescription: "Tactile switch with a light actuation force and smooth bump.",
					LongDescription:  "The TTC Gold Pink switch is designed with a light actuation force, making it ideal for fast typists and gamers alike. It offers a smooth, rounded tactile bump that provides satisfying feedback without being overly harsh, making it a versatile choice for various typing preferences.",
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
					LongDescription:  "The TTC Bluish White switch combines a crisp tactile bump with an audible click, delivering a distinct typing experience that is perfect for users who enjoy both tactile feedback and sound. Its bright color and satisfying click make it a popular choice for mechanical keyboard enthusiasts.",
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
					LongDescription:  "The TTC Silent Red switch offers a near-silent linear typing experience, thanks to its built-in dampening technology. It is ideal for quiet environments where noise reduction is key, without compromising on smooth and responsive keystrokes.",
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
					LongDescription:  "The TTC Gold Brown V3 switch features a distinct tactile bump with the added durability of gold-plated springs. This combination ensures a premium typing experience with enhanced longevity, making it suitable for users who seek reliability and performance.",
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
					LongDescription:  "The TTC Wild switch offers a smooth and consistent linear actuation, coupled with vibrant wild-themed aesthetics that add a unique flair to any keyboard. It’s an excellent choice for users who value both performance and distinctive design.",
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
					LongDescription:  "The TTC Flame Red switch is characterized by its smooth linear keystroke and eye-catching vibrant red color. It offers both aesthetic appeal and reliable performance, making it a great option for those looking to enhance their keyboard's visual and functional aspects.",
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
					LongDescription:  "The TTC Silent Brown switch offers a quiet tactile typing experience with a smooth bump and dampened noise. It’s perfect for users who need a tactile feel without the noise, making it suitable for shared or quiet environments.",
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
					LongDescription:  "The TTC Frozen Silent switch combines a smooth linear action with near-silent operation, all wrapped in a cold-themed design. Its unique aesthetics and quiet performance make it a favorite for custom builds in noise-sensitive environments.",
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
					LongDescription:  "The TTC Matrix White switch offers a pronounced tactile bump paired with a bright white color, making it a standout choice for those who appreciate both performance and visual appeal in their keyboard builds.",
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
					LongDescription:  "The TTC Speed Silver switch is engineered for speed, featuring a short actuation distance that makes it perfect for fast-paced gaming. Its quick responsiveness ensures that every key press is registered with minimal delay, giving gamers a competitive edge.",
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
					LongDescription:  "The TTC Blue switch is a classic clicky switch known for its pronounced tactile bump and satisfying audible click. It is a staple in the mechanical keyboard community, offering a traditional typing experience that many users love.",
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
					LongDescription:  "The TTC Gold Blue switch offers an elevated clicky experience, featuring gold-plated springs that enhance both durability and tactile feedback. Its smooth, satisfying click makes it a top choice for users seeking a premium clicky switch.",
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
					LongDescription:  "The TTC Gold Blue V2 switch builds on the success of the original, featuring a refined click mechanism that enhances both tactile and auditory feedback. It’s perfect for those who appreciate a crisp, consistent click with every keystroke.",
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
					LongDescription:  "The TTC Diamond Pink switch stands out with its vibrant pink color and sharp, diamond-like click sound. It combines aesthetic appeal with a distinct tactile and auditory experience, making it a perfect choice for those who want their keyboard to make a statement.",
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
					LongDescription:  "The TTC Golden Green V3 switch delivers a powerful tactile bump coupled with an enhanced, satisfying click sound. It’s designed for users who enjoy a strong, noticeable feedback with every keystroke, making it a popular choice among clicky switch enthusiasts.",
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
					LongDescription:  "The TTC Hart Blue switch stands out with its unique heart-shaped stem, offering a distinctive clicky feel with both visual and tactile appeal. This switch is perfect for users who want a playful yet satisfying typing experience.",
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
					LongDescription:  "The TTC Speed Purple switch is engineered for gaming, offering a fast actuation clicky experience that enhances responsiveness in gameplay. Its quick keypress and satisfying click make it an ideal choice for gamers who demand precision and speed.",
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
					LongDescription:  "The TTC Watermelon Milkshake switch is a fun, novelty switch with a fruity aesthetic and a crisp, satisfying click. It's perfect for those who want to add a playful touch to their keyboard while enjoying a clicky typing experience.",
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
					LongDescription:  "The TTC Crystal Ice switch combines a clear, transparent housing with a sharp, icy click sound, offering a visually stunning and audibly satisfying typing experience. It’s an excellent choice for those who prioritize both aesthetics and performance.",
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
					LongDescription:  "The TTC Light Ring switch is designed to enhance your keyboard’s lighting setup with integrated LED support, all while providing a satisfying clicky feel. It’s a great choice for users looking to combine visual flair with tactile feedback.",
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
					LongDescription:  "The TTC Brown switch is a staple in the mechanical keyboard world, known for its gentle tactile bump and smooth keystroke. It offers a balanced typing experience that’s suitable for both typing and gaming.",
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
					LongDescription:  "The TTC Silent Brown V2 switch offers an upgraded silent tactile experience, with a more refined bump and significantly reduced noise. It’s ideal for those who want the feel of a tactile switch without the accompanying sound.",
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
					LongDescription:  "The TTC Gold Brown switch combines the tactile bump loved by many with the added durability of gold-plated springs. This switch offers a premium typing experience with increased longevity, making it a top choice for those seeking both quality and performance.",
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
					LongDescription:  "The TTC Watermelon V2 switch offers a fun, tactile typing experience with a smooth bump and a fruity aesthetic. It’s perfect for those who want to add a playful touch to their keyboard without sacrificing performance.",
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
					LongDescription:  "The TTC Heart switch combines a unique heart-shaped stem design with a gentle tactile bump, offering both a distinctive aesthetic and a satisfying typing experience. It’s perfect for users looking to add a touch of personality to their keyboard.",
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
					LongDescription:  "The TTC Light Salmon switch offers a light actuation force paired with a smooth tactile bump, making it an excellent choice for users who prefer a more effortless typing experience. Its unique salmon color adds a touch of personality to any keyboard.",
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
					LongDescription:  "The TTC Honey V3 switch offers a smooth tactile bump with a unique honey-themed design. It’s perfect for users who appreciate both a satisfying typing feel and a visually appealing keyboard setup.",
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
					LongDescription:  "The TTC Purple Sky switch offers a balanced tactile typing experience with a medium actuation force and a serene sky-themed color. It’s a great choice for users who want a unique aesthetic combined with a reliable typing feel.",
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
					LongDescription:  "The TTC Purple Starlight switch offers a celestial aesthetic combined with a distinct tactile bump. It’s perfect for users who want a typing experience that’s out of this world, both in feel and appearance.",
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
					LongDescription:  "The TTC Greystar switch offers a crisp tactile bump paired with a neutral color palette, making it an ideal choice for professional setups. Its balanced typing feel and understated design make it a versatile option for any keyboard.",
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
					LongDescription:  "The TTC Titan Heart switch offers a strong tactile bump and a heart-themed design, providing both a satisfying typing experience and a unique aesthetic appeal. It’s perfect for those who want a bold and reliable tactile switch.",
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
					LongDescription:  "The TTC V3 Blue switch offers a refined tactile bump and a vibrant blue-themed design, providing a satisfying typing feel that’s both smooth and responsive. It’s an excellent choice for those who prefer a balance between aesthetics and performance.",
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
					LongDescription:  "The TTC Pink Girl switch offers a playful pink color and a gentle tactile bump, making it perfect for those who want a lighthearted yet satisfying typing experience. Its unique aesthetic adds a fun element to any keyboard setup.",
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
					LongDescription:  "The TTC Leaf switch features a unique leaf-shaped stem that provides a smooth tactile bump, offering a typing experience that’s both distinctive and enjoyable. It’s a great choice for users who want a nature-inspired design in their keyboard.",
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
					LongDescription:  "The TTC Silent Red V3 switch offers a quieter linear typing experience with enhanced smoothness, making it ideal for office environments or quiet spaces where noise reduction is essential. Its improved design ensures a whisper-quiet operation.",
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
					LongDescription:  "The TTC Silent Black switch offers a heavier actuation force combined with silent operation, providing a linear typing experience that’s both smooth and quiet. It’s perfect for users who need a silent keyboard without sacrificing feel.",
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
					LongDescription:  "The TTC Silent Sky switch combines a gentle tactile bump with a silent operation, providing a smooth typing experience that’s both satisfying and quiet. It’s ideal for shared or quiet environments where noise reduction is key.",
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
					LongDescription:  "The TTC Silver Blue switch is a specialty linear switch designed for fast actuation and low travel distance, catering specifically to gamers who require quick responsiveness. Its performance-driven design makes it a top choice for competitive gaming.",
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
					LongDescription:  "The TTC Ace switch offers a premium tactile experience with a unique actuation curve, providing a smooth bump that enhances the typing experience. Its refined design makes it a top choice for enthusiasts seeking both quality and performance.",
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
					LongDescription:  "The TTC Gold Pink V2 switch is an updated version of the original Gold Pink, offering improved smoothness and a more refined tactile bump. It’s ideal for both gaming and typing, providing a premium feel with every keypress.",
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
					LongDescription:  "The TTC Moonlight Silver switch offers a smooth linear typing experience with a moonlight-themed design, combining aesthetic appeal with reliable performance. It’s an excellent choice for those who want a keyboard that looks as good as it feels.",
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
					LongDescription:  "The TTC Sunset switch combines a distinctive tactile bump with a sunset-themed design, offering a unique typing experience that blends visual appeal with satisfying tactile feedback. It’s perfect for users who want both style and substance in their keyboard.",
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
					LongDescription:  "The TTC Wild Blue switch offers a bold tactile bump with a wild-themed design, perfect for users seeking a distinctive and satisfying typing experience. Its unique aesthetic makes it a great addition to any keyboard build.",
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
					LongDescription:  "The TTC Coral Sea switch provides a smooth linear typing experience with a coral-themed design, offering both aesthetic appeal and reliable performance. It’s an excellent choice for custom keyboard builds where both look and feel matter.",
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
					LongDescription:  "The TTC Tiger Light switch offers a smooth linear action with a light actuation force, all wrapped up in a striking tiger-themed design. It’s perfect for those who want a fast, responsive typing experience with a touch of wild flair.",
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
					LongDescription:  "The TTC Forest Green switch offers a smooth tactile bump combined with a nature-inspired forest-themed design. It’s perfect for users who appreciate both performance and aesthetics, bringing a touch of the outdoors to their keyboard.",
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
					LongDescription:  "The TTC Mars Orange switch offers a tactile typing experience with a distinct bump and a bold Mars-themed design, perfect for users who want a keyboard that’s both functional and visually striking.",
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
					LongDescription:  "The TTC Neptune Blue switch provides a smooth linear typing experience with a Neptune-themed design, making it a great choice for users who want a keyboard that combines aquatic aesthetics with top-notch performance.",
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
					LongDescription:  "The TTC Cool Mint switch offers a refreshing linear typing experience with a mint-themed design, providing both smooth action and a cool aesthetic that’s perfect for custom keyboard builds.",
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
					LongDescription:  "The Durock T1 is a highly regarded tactile switch, often praised for its sharp and distinct tactile bump that provides clear feedback with each keystroke. The POM stem ensures smooth operation and durability, while the high-quality materials minimize wobble, making it a popular choice among enthusiasts who value stability and precision in their typing experience. Its satisfying bump and overall performance make it a go-to option for both gaming and typing.",
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
					LongDescription:  "The Durock L7 is a linear switch that exemplifies the smooth, consistent keystroke sought after by linear switch enthusiasts. It features minimal spring noise, ensuring a quiet typing experience without compromising on performance. Known for its smoothness and low noise levels, the L7 is ideal for users who prefer a distraction-free environment while typing or gaming. The refined design and exceptional consistency make it a standout choice for those who demand reliability.",
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
					LongDescription:  "The Durock Koala is a tactile switch inspired by the iconic Holy Panda design, featuring a medium to heavy tactile bump that delivers a smooth and rounded feel. This switch is known for its consistent tactile feedback and is highly favored among enthusiasts for its premium typing experience. The Koala's high-quality construction and smooth actuation make it a top choice for those who appreciate tactile switches with a well-defined bump.",
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
					LongDescription:  "The Durock Sunflower is a linear switch that stands out with its vibrant yellow stem and incredibly smooth operation. It offers a quiet and fluid keystroke, making it an excellent choice for users who appreciate both aesthetics and performance. The Sunflower's bright color adds a touch of personality to any keyboard build, while its smooth action ensures a comfortable and enjoyable typing experience.",
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
					LongDescription:  "The Durock Shrimp is a tactile switch that combines a distinctive pink housing with a sharp, precise tactile bump. This switch is a favorite among those who appreciate both style and substance, offering a delightful typing experience with a responsive tactile feel. The Shrimp's unique design and excellent performance make it a standout choice for users who want a tactile switch that delivers on both aesthetics and functionality.",
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
					LongDescription:  "The Durock Dolphin is a linear switch that offers a smooth and consistent typing experience, characterized by its deep blue stem and exceptional build quality. Its refined design ensures minimal wobble and a reliable keystroke, making it an excellent option for users who prefer linear switches. The Dolphin's consistent performance and pleasing aesthetics make it a popular choice for both typing and gaming enthusiasts.",
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
					LongDescription:  "The Durock POM switch is a premium linear switch designed with a POM housing, known for its ability to enhance smoothness and reduce friction during keypresses. This switch is highly favored by enthusiasts who seek a top-tier linear experience with minimal scratchiness. The POM's superior material properties make it an excellent choice for those looking for a high-performance switch with a smooth and satisfying keystroke.",
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
					LongDescription:  "The Durock Daybreak switch features a visually striking dawn-themed design combined with smooth linear actuation. It offers a consistent and quiet typing experience, with minimal spring noise, making it ideal for users who prefer a peaceful and serene environment. The Daybreak's unique aesthetic and smooth performance make it a standout choice for those looking to add a touch of elegance to their keyboard.",
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
					LongDescription:  "The Durock Owl switch is a tactile switch that features an owl-themed housing and delivers a sharp, distinct tactile bump. Known for its precise feedback and unique aesthetic, the Owl switch is a favorite among tactile switch enthusiasts who appreciate both performance and visual appeal. This switch is perfect for users who want a well-defined tactile feel with a themed design that adds character to their keyboard.",
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
					LongDescription:  "The Durock Sea Turtle is a silent linear switch with a sea-themed design that provides a smooth and quiet typing experience. It's perfect for those who require silent operation without sacrificing the smoothness of a linear switch. The Sea Turtle's aesthetic appeal and silent performance make it an excellent choice for users who need a tranquil workspace or who frequently type in shared environments.",
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
					LongDescription:  "The Durock Alpaca V2 is an improved version of the original Alpaca switch, featuring a smoother feel and an updated spring design. This linear switch is renowned for its consistent performance, minimal noise, and satisfying keystroke. The Alpaca V2's smooth actuation and vibrant color make it a popular choice among enthusiasts seeking both reliability and aesthetics in their typing experience.",
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
					LongDescription:  "The Durock Marshmallow switch is a silent linear switch designed to provide a soft, marshmallow-like feel. With its quiet operation and gentle actuation, this switch is ideal for users who prefer a smooth, silent typing experience. The Marshmallow's unique texture and noise-dampening qualities make it a top choice for those seeking comfort and tranquility in their typing sessions.",
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
					LongDescription:  "The Durock Taro Ball switch is a tactile switch that features a taro-themed design, delivering a smooth and satisfying tactile bump. This switch is well-loved for its unique aesthetic and reliable performance, making it a favorite among tactile switch enthusiasts. The Taro Ball combines visual appeal with a distinct tactile feel, offering a premium typing experience that stands out in any build.",
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
					LongDescription:  "The Durock Corgi switch is a playful linear switch featuring a corgi-themed design and buttery smooth keystrokes. Known for its charming aesthetic and smooth operation, the Corgi switch is perfect for users who want to add a touch of fun to their keyboard while enjoying a premium typing experience. This switch is ideal for those who prioritize both performance and personality in their keyboard builds.",
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
					LongDescription:  "The Durock Shrimpy Click is a vibrant clicky switch designed to provide a sharp, satisfying click with every keystroke. Featuring a unique and colorful design, this switch offers an engaging typing experience for those who enjoy auditory feedback and bold aesthetics. The Shrimpy Click is perfect for users who want a lively and responsive switch that delivers both performance and style.",
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
					LongDescription:  "The Durock White Lotus switch is a clicky switch that combines a floral aesthetic with a distinct, crisp click. This switch is designed for users who appreciate themed switches that offer both visual appeal and tactile feedback. The White Lotus switch's unique look and satisfying click make it a great choice for those looking to add a touch of elegance and performance to their keyboard.",
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
					LongDescription:  "The Durock Thanos Click switch offers a deep, thocky click sound paired with a powerful and unique theme. This switch is designed for enthusiasts who seek a distinctive auditory experience combined with a bold aesthetic. The Thanos Click switch is ideal for those who enjoy a commanding sound profile and a switch that stands out both visually and audibly.",
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
					LongDescription:  "The Durock Peacock switch is a clicky switch that boasts a colorful, peacock-themed design and a lively click, making it a favorite for those who want a vibrant aesthetic paired with satisfying tactile and auditory feedback. The Peacock switch is perfect for users who enjoy a bold and expressive keyboard that performs as well as it looks.",
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
					LongDescription:  "The Durock Joyous Click switch is known for its cheerful design and crisp, satisfying click. This switch offers a joyful typing experience, blending a playful aesthetic with responsive tactile feedback. Ideal for users who want to add a touch of happiness to their setup, the Joyous Click switch is a delightful addition to any keyboard that values both performance and fun.",
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
					LongDescription:  "The Durock Brightside Click switch features a bright and cheerful design, combined with a clear and crisp click sound. This switch is an excellent choice for users who appreciate a vibrant aesthetic along with satisfying tactile and auditory feedback. The Brightside Click is perfect for those looking to brighten up their keyboard with a switch that delivers both style and substance.",
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
					LongDescription:  "The Durock Firestorm switch is a clicky switch that offers a dramatic, fiery design paired with an explosive click sound. This switch is perfect for users who want to make a bold statement with their keyboard, both visually and audibly. The Firestorm switch's powerful feedback and striking appearance make it an excellent choice for those seeking a dynamic and engaging typing experience.",
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
					LongDescription:  "The Durock T2 is an enhanced version of the T1, offering a smoother tactile bump and reduced noise during actuation. Retaining the high-quality construction that Durock is known for, the T2 switch is a reliable choice for tactile switch fans who want a refined typing experience. This switch is perfect for users who appreciate precise tactile feedback with improved smoothness and reduced sound levels.",
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
					LongDescription:  "The Durock Shrimp Tactile switch is known for its unique pink housing and satisfying tactile bump. This switch offers both aesthetic appeal and excellent performance, making it a standout choice for users who value distinctive looks alongside great tactile feedback. The Shrimp Tactile switch is perfect for those who want a tactile typing experience with a touch of personality and style.",
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
					LongDescription:  "The Durock Koala V2 is an upgraded version of the original Koala switch, offering a more refined tactile bump and enhanced smoothness. This switch is designed for enthusiasts who want a premium typing experience with reliable tactile feedback. The Koala V2's improvements in both feel and performance make it a top choice for those seeking an elevated tactile switch.",
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
					LongDescription:  "The Durock POM Tactile switch is designed with a POM housing, providing a unique feel and reduced friction for a smooth, satisfying tactile experience. This switch is perfect for those who seek a distinct tactile switch that offers both premium materials and excellent feedback. The POM Tactile switch's combination of smoothness and tactile response makes it a favorite among users looking for a high-quality switch with a different feel.",
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
					LongDescription:  "The Durock Daybreak Tactile switch offers a dawn-themed design paired with a smooth tactile bump, providing an excellent balance between aesthetics and performance. This switch is ideal for users who want both style and substance in their typing experience. The Daybreak Tactile switch's unique look and reliable tactile feedback make it a great choice for those seeking a refined tactile switch with visual appeal.",
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
					LongDescription:  "The Durock Owl Tactile switch features an owl-themed design and a sharp tactile bump, providing precise feedback that is perfect for tactile switch enthusiasts who enjoy themed aesthetics. This switch is a great choice for users who want a well-defined tactile feel combined with a distinctive look, making it a unique addition to any keyboard.",
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
					LongDescription:  "The Durock Lilac Tactile switch offers a soothing lilac color and a smooth tactile bump, making it a perfect choice for users who appreciate both aesthetics and performance. This switch is ideal for those who want a unique color-themed build without compromising on the quality of their tactile feedback. The Lilac Tactile switch's elegant design and reliable performance make it a standout in any keyboard.",
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
					LongDescription:  "The Durock Zilent V2 is a silent tactile switch offering a soft bump and minimal noise, making it perfect for users who require a quiet typing experience without sacrificing tactile feedback. Known for its smooth actuation and quiet operation, the Zilent V2 is highly regarded among enthusiasts who value silence and performance in equal measure.",
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
					LongDescription:  "The Durock Piano switch is a tactile switch that draws inspiration from the world of music, offering a crisp tactile bump and a unique aesthetic. This switch is designed for users who want a switch that not only performs well but also adds a touch of musical elegance to their setup. The Piano switch's combination of satisfying feedback and distinctive design makes it a popular choice among musicians and keyboard enthusiasts alike.",
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
					LongDescription:  "The Durock Lavender switch offers a tactile bump with a calming lavender hue, providing a smooth typing experience with an appealing color scheme. This switch is perfect for those who want a stylish and satisfying tactile switch that stands out in both form and function. The Lavender switch's unique color and reliable performance make it a great choice for themed builds.",
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
					LongDescription:  "The Durock Pudding switch is a tactile switch that combines a unique feel with a playful design. This switch offers a satisfying tactile bump and is ideal for users who want a switch that stands out both in performance and appearance. The Pudding switch's fun and distinctive look, coupled with its reliable tactile feedback, make it a popular choice for those looking to add a touch of whimsy to their keyboard.",
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
					LongDescription:  "The Durock Camel switch is a tactile switch featuring a sand-colored housing and a smooth tactile bump. Inspired by the desert, this switch is designed for users who appreciate a unique, earthy aesthetic alongside a satisfying typing experience. The Camel switch's distinct design and reliable performance make it a great choice for those looking for a tactile switch with a natural, grounded feel.",
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
					LongDescription:  "The Durock Azure switch offers a tactile bump with an azure-themed design, providing a crisp typing experience with a unique visual appeal. This switch is perfect for users who enjoy themed switches that deliver both style and performance. The Azure switch's combination of aesthetics and reliable tactile feedback makes it a standout choice for those looking to add a splash of color to their setup.",
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
					LongDescription:  "The Durock Coral Tactile switch features a coral-themed design and a strong tactile bump, offering a distinct aesthetic and satisfying feedback. This switch is ideal for users who want a bold tactile switch with a unique theme, making it a perfect addition to any keyboard that values both performance and visual appeal.",
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
					LongDescription:  "The Durock L1 is a linear switch known for its exceptionally smooth operation and minimal noise, making it an excellent choice for users seeking a quiet and fluid typing experience. This switch is favored for its consistency and low spring noise, providing a reliable performance that is ideal for both typing and gaming environments.",
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
					LongDescription:  "The Durock Alpaca V1 switch is a linear switch renowned for its smooth feel and vibrant pink stem. This switch offers a consistent typing experience with minimal spring noise, making it a favorite among enthusiasts seeking reliable performance. The Alpaca V1's iconic color and smooth action have made it a staple in the mechanical keyboard community.",
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
					LongDescription:  "The Durock Cobalt switch features a striking cobalt blue stem and provides a smooth linear action, making it a visually and functionally appealing choice for any keyboard build. This switch is designed for users who appreciate a deep color aesthetic combined with a consistent and reliable typing feel. The Cobalt switch's performance and design make it a top choice for those who want both style and substance in their linear switches.",
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
					LongDescription:  "The Durock Sunflower Linear switch is known for its bright yellow stem and buttery smooth operation. It provides a quiet and fluid keystroke, ideal for those who want both a vibrant aesthetic and high performance. The Sunflower Linear switch's unique color and smooth feel make it an excellent choice for users looking for a switch that delivers both visual appeal and reliable performance.",
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
					LongDescription:  "The Durock Ghost switch features a translucent housing that highlights its smooth linear keystroke. This switch is designed for users who enjoy a clean aesthetic with the performance of a premium linear switch. The Ghost switch's combination of visual clarity and smooth action makes it a popular choice for those who want a sleek, minimalistic look without sacrificing quality.",
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
					LongDescription:  "The Durock Dolphin V2 switch is an upgraded version of the original, offering improved smoothness and minimal spring noise. Its consistent performance and pleasing aesthetics make it a popular choice for linear switch enthusiasts who demand reliability and quality. The Dolphin V2 switch is perfect for users seeking a refined linear switch that delivers on both feel and durability.",
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
					LongDescription:  "The Durock Haze switch is known for its hazy design and buttery keystroke, providing a smooth typing experience with a unique aesthetic. This switch is perfect for users who want a switch that combines style with performance, offering a consistent linear feel that is both comfortable and reliable. The Haze switch's misty appearance and smooth action make it a popular choice for enthusiasts looking to add a touch of mystery to their keyboard.",
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
					LongDescription:  "The Durock Mauve switch offers a mauve-colored stem and a smooth linear action, making it a top choice for enthusiasts who prefer a unique color scheme and consistent performance. This switch's distinctive hue and smooth keystroke make it a standout option for users who value both aesthetics and reliable operation in their linear switches.",
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
					LongDescription:  "The Durock Midnight switch is a linear switch characterized by its dark design and quiet operation, offering a smooth and subtle typing experience that is perfect for nighttime use. This switch is ideal for users who prefer a low-profile aesthetic and a quiet, reliable keystroke. The Midnight switch's combination of stealthy appearance and smooth action makes it an excellent choice for those who value discretion and performance.",
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
					LongDescription:  "The Durock Lightwave switch provides a light actuation force and a smooth glide, making it ideal for users who prefer quick keystrokes with minimal resistance. This switch is perfect for those who want a responsive and effortless typing experience, with a smooth and consistent feel that enhances typing speed and comfort. The Lightwave switch's blend of light touch and smooth action makes it a favorite among users who value speed and precision.",
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
					LongDescription:  "The Durock Euphoria switch is designed to provide a euphoric typing experience with a consistent linear action. This switch is perfect for those who want a smooth and satisfying keystroke with minimal noise, offering a blend of comfort and performance that enhances any typing session. The Euphoria switch's premium feel and reliable operation make it a top choice for enthusiasts seeking a high-quality linear switch.",
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
					LongDescription:  "The Durock Vanilla switch offers a creamy smooth action and a neutral design, making it a versatile choice for any keyboard setup. This switch is known for its reliability and minimal spring noise, providing a consistent and quiet typing experience that is ideal for both work and play. The Vanilla switch's simplicity and smooth operation make it a dependable choice for those seeking a no-frills linear switch that delivers on performance.",
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
					LongDescription:  "The Durock Jet switch is known for its jet-black design and fast actuation, making it perfect for gamers who require quick and responsive keystrokes. This switch is designed for those who need performance and speed, offering a consistent and reliable feel that enhances gaming and typing efficiency. The Jet switch's sleek design and rapid response make it a top choice for competitive users who demand the best from their switches.",
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
					LongDescription:  "The Durock Blossom switch features a delicate floral design combined with a smooth linear keystroke. This switch is perfect for those who appreciate a gentle aesthetic alongside reliable performance. The Blossom switch's elegant appearance and buttery smooth action make it a popular choice for users looking to add a touch of grace to their keyboard while maintaining top-tier performance.",
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
					LongDescription:  "The Durock Ink Black V2 switch offers a unique ink-colored housing and buttery smoothness, providing a premium typing experience with a distinctive look. This switch is ideal for users who want high performance and a unique design. The Ink Black V2's combination of style and substance makes it a top choice for those seeking a standout linear switch that delivers on both form and function.",
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
					LongDescription:  "The Durock Cotton Candy switch is a linear switch known for its sweet aesthetic and soft keystroke, offering a playful look with smooth performance. This switch is perfect for users who want a switch that adds a fun element to their keyboard while providing a reliable typing experience. The Cotton Candy switch's vibrant colors and gentle actuation make it a delightful choice for those seeking a lighthearted and high-quality switch.",
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
					LongDescription:  "The Durock Silent Linear switch is designed for those who need a quiet typing experience without sacrificing performance. It offers a smooth linear keystroke with built-in dampening to reduce noise, making it ideal for office or shared spaces. The Silent Linear switch's blend of silence and smooth action makes it a reliable choice for environments where quiet operation is essential.",
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
					LongDescription:  "The Durock Silent T1 switch offers a tactile typing experience with a refined bump and minimal noise. This switch is perfect for users who want tactile feedback without the accompanying sound, making it suitable for quiet environments. The Silent T1's balance of tactile feel and noise reduction makes it a top choice for those seeking a silent yet satisfying switch.",
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
					LongDescription:  "The Durock Thocky switch is a specialty linear switch known for its deep, thocky sound profile. It offers a unique auditory experience with smooth keystrokes, making it a favorite among enthusiasts who enjoy distinct sound feedback. The Thocky switch's combination of smooth action and rich sound makes it an excellent choice for users who prioritize both feel and sound in their typing experience.",
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
					LongDescription:  "The Durock U4 Silent switch is designed to provide a soft tactile bump with ultra-quiet operation, making it ideal for environments where noise reduction is critical. This switch offers smooth, dampened keystrokes with reliable tactile feedback, ensuring a satisfying typing experience without the noise. The U4 Silent switch's balance of feel and silence makes it a top choice for users who need a quiet yet responsive switch.",
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
					LongDescription:  "The Durock Light Fog switch offers a specialty design with a translucent housing that creates a mysterious aesthetic. This switch provides a smooth keystroke and is perfect for those who appreciate both performance and style. The Light Fog switch's combination of unique appearance and reliable action makes it a standout choice for users who want a switch that delivers on both looks and feel.",
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
					LongDescription:  "The Durock Shadow switch is a specialty switch featuring a dark design and smooth performance. It offers a quiet typing experience and is ideal for those who want a switch with a stealthy appearance and reliable action. The Shadow switch's combination of aesthetics and smooth operation makes it a great choice for users who prefer a subtle, understated look without compromising on performance.",
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
					LongDescription:  "The Durock Marshmallow Silent switch combines a soft, marshmallow-like feel with silent operation, offering a unique typing experience that's both quiet and satisfying. This switch is perfect for those who want a silent switch with a distinctive touch, providing smooth and noise-free keystrokes in any environment. The Marshmallow Silent switch's blend of comfort and silence makes it a popular choice for those seeking a premium, quiet typing experience.",
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
					LongDescription:  "The Durock L4 switch is a specialty linear switch offering enhanced smoothness and reduced scratchiness. Designed for users who seek a refined typing experience with consistent linear action, the L4 switch provides a reliable and satisfying keystroke that is ideal for both work and gaming. The L4's balance of smoothness and durability makes it a top choice for those looking to upgrade their linear switches.",
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
					LongDescription:  "The Durock Daydream switch offers a dream-themed design with a smooth keystroke, providing a unique aesthetic alongside reliable performance. This switch is perfect for those who want a switch that captures the imagination while delivering excellent typing feedback. The Daydream switch's combination of style and substance makes it a popular choice for users who want a switch that stands out in both form and function.",
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
					LongDescription:  "The Durock Aqua King switch features a clear housing that highlights its smooth, fluid linear action. This switch is designed for users who want a transparent aesthetic combined with top-tier linear performance. The Aqua King switch's crystal-clear appearance and silky-smooth keystroke make it a standout choice for those seeking both visual clarity and superior performance in their linear switches.",
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
					LongDescription:  "The Durock Ultra switch offers an ultra-smooth linear feel with minimal noise, providing a premium typing experience that's both quiet and satisfying. This switch is ideal for those who demand the highest level of performance from their switches, offering consistent keystrokes with a smooth and seamless action. The Ultra switch's blend of silence and smoothness makes it a top choice for users seeking a refined and high-performance linear switch.",
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
					LongDescription:  "The Durock Crystal Glacier switch is a specialty switch with a crystal-clear housing that offers icy smoothness in its linear action. This switch is designed for those who want a striking aesthetic combined with seamless performance. The Crystal Glacier switch's blend of visual appeal and smooth keystrokes makes it a popular choice for users who value both style and high performance in their linear switches.",
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
					LongDescription:  "The Durock Emerald switch features a rich emerald hue and luxurious smoothness, providing a premium typing experience with a touch of elegance. This switch is perfect for users who want both performance and style in their keyboard. The Emerald switch's combination of deep color and refined keystroke makes it a standout choice for those who seek a high-quality linear switch with a luxurious feel.",
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
					LongDescription:  "The Durock Mystic switch offers a mystical design combined with smooth linear action, making it a favorite for users who want a switch with an enchanting appearance and reliable performance. The Mystic switch's unique aesthetic and consistent feel make it a great choice for those who appreciate both form and function in their keyboard switches.",
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
