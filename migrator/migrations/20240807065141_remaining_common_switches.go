package migrations

import (
	"time"

	"github.com/Bparsons0904/deadigations"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	deadigations.RegisterMigration(deadigations.Migration{
		ID:          "20240807065141",
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

			kdfans := Producer{
				Name:    "KDFans",
				Alias:   "kdfans",
				SiteURL: "",
			}

			if err := tx.Create(&kdfans).Error; err != nil {
				return err
			}

			hmx := Producer{
				Name:    "HMX",
				Alias:   "hmx",
				SiteURL: "",
			}

			if err := tx.Create(&hmx).Error; err != nil {
				return err
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

			keychronSwitches := []Switch{
				{
					Name:             "Keychron K Pro Red",
					ShortDescription: "Linear switch with smooth keypresses.",
					LongDescription:  "Keychron K Pro Red is a linear switch that provides a smooth and consistent keypress, perfect for fast typists.",
					ManufacturerID:   ptr(12), // Keychron
					BrandID:          ptr(12), // Keychron
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2022-03-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Keychron K Pro Brown",
					ShortDescription: "Tactile switch for a responsive typing experience.",
					LongDescription:  "Keychron K Pro Brown offers a tactile bump, providing a responsive typing experience suitable for typing and gaming.",
					ManufacturerID:   ptr(12), // Keychron
					BrandID:          ptr(12), // Keychron
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2022-03-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Keychron K Pro Blue",
					ShortDescription: "Clicky switch with an audible click sound.",
					LongDescription:  "Keychron K Pro Blue is a clicky switch known for its audible click sound and tactile feedback, enhancing the typing experience.",
					ManufacturerID:   ptr(12), // Keychron
					BrandID:          ptr(12), // Keychron
					SwitchTypeID:     3,       // Clicky
					ReleaseDate:      parseDate("2022-03-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
			}
			err := processSwitches(keychronSwitches)
			if err != nil {
				return err
			}

			epomakerSwitches := []Switch{
				{
					Name:             "Epomaker Flamingo",
					ShortDescription: "Linear switch with smooth actuation.",
					LongDescription:  "The Epomaker Flamingo is a linear switch known for its smooth actuation and consistent feel, making it a favorite among enthusiasts.",
					ManufacturerID:   ptr(13), // Epomaker
					BrandID:          ptr(13), // Epomaker
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2021-11-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Epomaker Budgerigar",
					ShortDescription: "Tactile switch with a unique feel.",
					LongDescription:  "The Epomaker Budgerigar is a tactile switch offering a unique bump feel, ideal for those who prefer a pronounced tactile experience.",
					ManufacturerID:   ptr(13), // Epomaker
					BrandID:          ptr(13), // Epomaker
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2021-11-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Epomaker Ice Candy",
					ShortDescription: "Smooth linear switch with a distinct color.",
					LongDescription:  "Epomaker Ice Candy switches offer a smooth linear experience with a visually appealing design, perfect for colorful builds.",
					ManufacturerID:   ptr(13), // Epomaker
					BrandID:          ptr(13), // Epomaker
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2022-05-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Epomaker Blue Ice",
					ShortDescription: "Linear switch with smooth keypresses.",
					LongDescription:  "Epomaker Blue Ice is a linear switch that provides a smooth and consistent keypress, ideal for fast typists.",
					ManufacturerID:   ptr(13), // Epomaker
					BrandID:          ptr(13), // Epomaker
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2022-03-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
			}
			err = processSwitches(epomakerSwitches)
			if err != nil {
				return err
			}

			razerSwitches := []Switch{
				{
					Name:             "Razer Green",
					ShortDescription: "Clicky switch with tactile feedback.",
					LongDescription:  "Razer Green switches provide a satisfying click and tactile feedback, ideal for gamers and typists who enjoy a clicky experience.",
					ManufacturerID:   ptr(14), // Razer
					BrandID:          ptr(14), // Razer
					SwitchTypeID:     3,       // Clicky
					ReleaseDate:      parseDate("2014-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Razer Yellow",
					ShortDescription: "Smooth linear switch with fast actuation.",
					LongDescription:  "Razer Yellow is a linear switch designed for fast actuation and smooth keystrokes, making it ideal for competitive gaming.",
					ManufacturerID:   ptr(14), // Razer
					BrandID:          ptr(14), // Razer
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2016-10-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Razer Orange",
					ShortDescription: "Tactile switch with silent actuation.",
					LongDescription:  "Razer Orange switches offer tactile feedback without the noise, providing a silent yet responsive typing experience.",
					ManufacturerID:   ptr(14), // Razer
					BrandID:          ptr(14), // Razer
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2014-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
			}
			err = processSwitches(razerSwitches)
			if err != nil {
				return err
			}

			logitechSwitches := []Switch{
				{
					Name:             "Logitech GX Blue",
					ShortDescription: "Clicky switch with tactile feedback.",
					LongDescription:  "Logitech GX Blue switches provide a clicky feel with tactile feedback, ideal for users who enjoy audible and tactile typing.",
					ManufacturerID:   ptr(15), // Logitech
					BrandID:          ptr(15), // Logitech
					SwitchTypeID:     3,       // Clicky
					ReleaseDate:      parseDate("2019-05-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Logitech GX Brown",
					ShortDescription: "Tactile switch with a quieter operation.",
					LongDescription:  "Logitech GX Brown switches offer a tactile bump with quieter operation, perfect for both gaming and typing.",
					ManufacturerID:   ptr(15), // Logitech
					BrandID:          ptr(15), // Logitech
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2019-05-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Logitech GX Red",
					ShortDescription: "Linear switch for smooth keypresses.",
					LongDescription:  "Logitech GX Red is a linear switch designed for smooth keypresses and fast response times, suitable for gamers.",
					ManufacturerID:   ptr(15), // Logitech
					BrandID:          ptr(15), // Logitech
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2019-05-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
			}
			err = processSwitches(logitechSwitches)
			if err != nil {
				return err
			}
			roccatSwitches := []Switch{
				{
					Name:             "Roccat Titan Switch Tactile",
					ShortDescription: "Tactile switch with a crisp feel.",
					LongDescription:  "Roccat Titan Switch Tactile provides a crisp tactile feedback, enhancing typing precision and response.",
					ManufacturerID:   ptr(19), // Roccat
					BrandID:          ptr(19), // Roccat
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2018-07-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Roccat Titan Switch Linear",
					ShortDescription: "Smooth linear switch for fast response.",
					LongDescription:  "Roccat Titan Switch Linear offers smooth keystrokes and fast response, ideal for gamers seeking a seamless experience.",
					ManufacturerID:   ptr(19), // Roccat
					BrandID:          ptr(19), // Roccat
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2018-07-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
			}
			err = processSwitches(roccatSwitches)
			if err != nil {
				return err
			}

			coolerMasterSwitches := []Switch{
				{
					Name:             "Cooler Master Green",
					ShortDescription: "Clicky switch with a tactile feel.",
					LongDescription:  "Cooler Master Green switches offer a clicky feel with tactile feedback, perfect for users who enjoy audible typing.",
					ManufacturerID:   ptr(20), // Cooler Master
					BrandID:          ptr(20), // Cooler Master
					SwitchTypeID:     3,       // Clicky
					ReleaseDate:      parseDate("2016-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Cooler Master Brown",
					ShortDescription: "Tactile switch with a quieter operation.",
					LongDescription:  "Cooler Master Brown switches provide tactile feedback with quieter operation, ideal for both work and play.",
					ManufacturerID:   ptr(20), // Cooler Master
					BrandID:          ptr(20), // Cooler Master
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2016-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Cooler Master Red",
					ShortDescription: "Linear switch for fast keystrokes.",
					LongDescription:  "Cooler Master Red is a linear switch designed for fast keystrokes and smooth typing, suitable for gaming.",
					ManufacturerID:   ptr(20), // Cooler Master
					BrandID:          ptr(20), // Cooler Master
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2016-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Cooler Master Purple",
					ShortDescription: "Tactile switch with a pronounced bump.",
					LongDescription:  "Cooler Master Purple switches offer a tactile experience with a more pronounced bump than typical brown switches, suitable for typing enthusiasts.",
					ManufacturerID:   ptr(20),                 // Cooler Master
					BrandID:          ptr(20),                 // Cooler Master
					SwitchTypeID:     2,                       // Tactile
					ReleaseDate:      parseDate("2017-01-01"), // Approximate, please verify
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Cooler Master Blue",
					ShortDescription: "Clicky switch with tactile and audible feedback.",
					LongDescription:  "Cooler Master Blue switches provide both tactile and audible feedback, offering a satisfying click sound with each keystroke.",
					ManufacturerID:   ptr(20),                 // Cooler Master
					BrandID:          ptr(20),                 // Cooler Master
					SwitchTypeID:     3,                       // Clicky
					ReleaseDate:      parseDate("2017-01-01"), // Approximate, please verify
					Available:        true,
					PricePoint:       2, // Average
				},
			}
			err = processSwitches(coolerMasterSwitches)
			if err != nil {
				return err
			}

			jwkSwitches := []Switch{
				{
					Name:             "JWK Black",
					ShortDescription: "Smooth linear switch with deep sound.",
					LongDescription:  "JWK Black switches offer a smooth linear feel with a deep, satisfying sound, perfect for enthusiasts seeking a premium experience.",
					ManufacturerID:   ptr(22), // JWK
					BrandID:          ptr(22), // JWK
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2021-02-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "JWK Red",
					ShortDescription: "Smooth linear switch with light actuation.",
					LongDescription:  "JWK Red is a linear switch known for its smooth actuation and light feel, making it suitable for fast typists.",
					ManufacturerID:   ptr(22), // JWK
					BrandID:          ptr(22), // JWK
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2021-02-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "JWK Lavender",
					ShortDescription: "Linear switch with a unique color and smooth feel.",
					LongDescription:  "JWK Lavender switches are linear switches featuring a unique lavender color and known for their smooth keystroke.",
					ManufacturerID:   ptr(22),                 // JWK
					BrandID:          ptr(22),                 // JWK
					SwitchTypeID:     1,                       // Linear
					ReleaseDate:      parseDate("2021-03-01"), // Approximate
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "JWK Alpaca",
					ShortDescription: "Smooth linear switch with a light touch.",
					LongDescription:  "JWK Alpaca switches are known for their smooth linear feel and light touch, making them popular among enthusiasts for both typing and gaming.",
					ManufacturerID:   ptr(22),                 // JWK
					BrandID:          ptr(22),                 // JWK
					SwitchTypeID:     1,                       // Linear
					ReleaseDate:      parseDate("2020-01-01"), // Approximate
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "JWK Moss",
					ShortDescription: "Tactile switch with a rounded bump.",
					LongDescription:  "JWK Moss switches offer a tactile experience with a rounded bump, providing a satisfying feedback without being too sharp.",
					ManufacturerID:   ptr(22),                 // JWK
					BrandID:          ptr(22),                 // JWK
					SwitchTypeID:     2,                       // Tactile
					ReleaseDate:      parseDate("2020-06-01"), // Approximate
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "JWK Marshmallow",
					ShortDescription: "Smooth linear switch with a unique color.",
					LongDescription:  "JWK Marshmallow switches are linear switches known for their smooth feel and unique marshmallow-inspired coloration.",
					ManufacturerID:   ptr(22),                 // JWK
					BrandID:          ptr(22),                 // JWK
					SwitchTypeID:     1,                       // Linear
					ReleaseDate:      parseDate("2020-09-01"), // Approximate
					Available:        true,
					PricePoint:       3, // Expensive
				},
			}
			err = processSwitches(jwkSwitches)
			if err != nil {
				return err
			}

			tecseeSwitches := []Switch{
				{
					Name:             "Tecsee Carrot",
					ShortDescription: "Tactile switch with a unique feel.",
					LongDescription:  "Tecsee Carrot switches offer a unique tactile feel, making them a favorite among enthusiasts seeking something different.",
					ManufacturerID:   ptr(23), // Tecsee
					BrandID:          ptr(23), // Tecsee
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2020-11-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Tecsee Purple Panda",
					ShortDescription: "Tactile switch with a satisfying bump.",
					LongDescription:  "Tecsee Purple Panda provides a tactile experience with a satisfying bump, perfect for those who appreciate a pronounced tactile feedback.",
					ManufacturerID:   ptr(23), // Tecsee
					BrandID:          ptr(23), // Tecsee
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2020-11-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
			}
			err = processSwitches(tecseeSwitches)
			if err != nil {
				return err
			}

			everglideSwitches := []Switch{
				{
					Name:             "Everglide Aqua King",
					ShortDescription: "Smooth linear switch with a unique design.",
					LongDescription:  "Everglide Aqua King switches offer a smooth linear feel with a unique and visually appealing design, perfect for custom builds.",
					ManufacturerID:   ptr(24), // Everglide
					BrandID:          ptr(24), // Everglide
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2021-08-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Everglide Coral Red",
					ShortDescription: "Linear switch with a smooth feel.",
					LongDescription:  "Everglide Coral Red switches provide a smooth linear experience, offering consistent keypresses and a visually striking design.",
					ManufacturerID:   ptr(24), // Everglide
					BrandID:          ptr(24), // Everglide
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2021-08-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
			}
			err = processSwitches(everglideSwitches)
			if err != nil {
				return err
			}

			spStarSwitches := []Switch{
				{
					Name:             "SP-Star Purple",
					ShortDescription: "Tactile switch with a unique sound profile.",
					LongDescription:  "SP-Star Purple switches offer a unique tactile feel and sound profile, perfect for those who want a distinct typing experience.",
					ManufacturerID:   ptr(25), // SP-Star
					BrandID:          ptr(25), // SP-Star
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2021-09-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "SP-Star Yellow",
					ShortDescription: "Linear switch with smooth actuation.",
					LongDescription:  "SP-Star Yellow switches provide a smooth linear actuation, ideal for gamers and typists who prefer a seamless keystroke experience.",
					ManufacturerID:   ptr(25), // SP-Star
					BrandID:          ptr(25), // SP-Star
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2021-09-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
			}
			err = processSwitches(spStarSwitches)
			if err != nil {
				return err
			}

			gazzewSwitches := []Switch{
				{
					Name:             "Gazzew Boba U4",
					ShortDescription: "Silent tactile switch with a unique feel.",
					LongDescription:  "Gazzew Boba U4 switches offer silent tactile feedback, providing a unique and satisfying typing experience without the noise.",
					ManufacturerID:   ptr(26), // Gazzew
					BrandID:          ptr(26), // Gazzew
					SwitchTypeID:     5,       // Silent Tactile
					ReleaseDate:      parseDate("2019-11-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Gazzew Boba U4T",
					ShortDescription: "Tactile switch with a unique feel.",
					LongDescription:  "Gazzew Boba U4T switches provide tactile feedback with a unique bump, perfect for typists seeking a pronounced tactile experience.",
					ManufacturerID:   ptr(26), // Gazzew
					BrandID:          ptr(26), // Gazzew
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2019-11-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
			}
			err = processSwitches(gazzewSwitches)
			if err != nil {
				return err
			}

			hmxSwitches := []Switch{
				{
					Name:             "HMX Blue",
					ShortDescription: "Clicky switch with tactile feedback.",
					LongDescription:  "HMX Blue switches provide a clicky feel with tactile feedback, perfect for users who enjoy an audible typing experience.",
					ManufacturerID:   ptr(30), // HMX
					BrandID:          ptr(30), // HMX
					SwitchTypeID:     3,       // Clicky
					ReleaseDate:      parseDate("2018-07-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "HMX Red",
					ShortDescription: "Linear switch for smooth keypresses.",
					LongDescription:  "HMX Red is a linear switch offering smooth keypresses and fast response times, suitable for both gaming and typing.",
					ManufacturerID:   ptr(30), // HMX
					BrandID:          ptr(30), // HMX
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2018-07-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "HMX Brown",
					ShortDescription: "Tactile switch with a subtle bump.",
					LongDescription:  "HMX Brown switches offer a tactile bump with quieter operation, ideal for users who prefer a balance between tactile feedback and noise.",
					ManufacturerID:   ptr(30), // HMX
					BrandID:          ptr(30), // HMX
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2018-07-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "HMX Silent Red",
					ShortDescription: "Silent linear switch with smooth actuation.",
					LongDescription:  "HMX Silent Red switches provide a smooth linear experience without the noise, perfect for quiet environments.",
					ManufacturerID:   ptr(30), // HMX
					BrandID:          ptr(30), // HMX
					SwitchTypeID:     4,       // Silent Linear
					ReleaseDate:      parseDate("2019-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "HMX Black",
					ShortDescription: "Heavy linear switch with deep actuation.",
					LongDescription:  "HMX Black switches offer a heavy linear feel with deep actuation, suitable for typists who prefer a more substantial keystroke.",
					ManufacturerID:   ptr(30), // HMX
					BrandID:          ptr(30), // HMX
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2019-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "HMX Silent Black",
					ShortDescription: "Silent heavy linear switch.",
					LongDescription:  "HMX Silent Black switches provide a heavy linear feel with silent actuation, ideal for those who prefer a quiet typing experience with substantial feedback.",
					ManufacturerID:   ptr(30), // HMX
					BrandID:          ptr(30), // HMX
					SwitchTypeID:     4,       // Silent Linear
					ReleaseDate:      parseDate("2020-02-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "HMX Green",
					ShortDescription: "Clicky switch with strong tactile feedback.",
					LongDescription:  "HMX Green switches offer a strong clicky feel with pronounced tactile feedback, perfect for those who enjoy a loud typing experience.",
					ManufacturerID:   ptr(30), // HMX
					BrandID:          ptr(30), // HMX
					SwitchTypeID:     3,       // Clicky
					ReleaseDate:      parseDate("2019-11-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "HMX Clear",
					ShortDescription: "Tactile switch with medium resistance.",
					LongDescription:  "HMX Clear switches offer medium resistance with tactile feedback, ideal for those who prefer a balanced typing feel.",
					ManufacturerID:   ptr(30), // HMX
					BrandID:          ptr(30), // HMX
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2020-08-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "HMX Speed Silver",
					ShortDescription: "Fast linear switch for quick typing.",
					LongDescription:  "HMX Speed Silver switches are designed for fast actuation and quick typing, making them ideal for gamers seeking speed and efficiency.",
					ManufacturerID:   ptr(30), // HMX
					BrandID:          ptr(30), // HMX
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2021-03-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "HMX Pink",
					ShortDescription: "Linear switch with a soft touch.",
					LongDescription:  "HMX Pink switches offer a linear feel with a soft touch, suitable for those who prefer a gentle typing experience.",
					ManufacturerID:   ptr(30), // HMX
					BrandID:          ptr(30), // HMX
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2021-07-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "HMX Xinhai",
					ShortDescription: "Tactile switch with a soft bump.",
					LongDescription:  "HMX Xinhai switches provide a soft tactile bump with smooth actuation, offering a pleasant typing experience for both gaming and typing.",
					ManufacturerID:   ptr(30), // HMX
					BrandID:          ptr(30), // HMX
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2021-12-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "HMX Purple Dawn",
					ShortDescription: "Silent tactile switch with unique feedback.",
					LongDescription:  "HMX Purple Dawn switches offer silent tactile feedback with a unique actuation profile, perfect for those who enjoy a quiet yet tactile experience.",
					ManufacturerID:   ptr(30), // HMX
					BrandID:          ptr(30), // HMX
					SwitchTypeID:     5,       // Silent Tactile
					ReleaseDate:      parseDate("2022-03-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "HMX Cloud White",
					ShortDescription: "Linear switch with a soft feel.",
					LongDescription:  "HMX Cloud White switches provide a smooth linear feel with a soft actuation, offering a gentle and pleasant typing experience.",
					ManufacturerID:   ptr(30), // HMX
					BrandID:          ptr(30), // HMX
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2022-05-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "HMX Coral",
					ShortDescription: "Tactile switch with a colorful design.",
					LongDescription:  "HMX Coral switches offer a tactile experience with a visually striking design, making them a great choice for custom keyboard builds.",
					ManufacturerID:   ptr(30), // HMX
					BrandID:          ptr(30), // HMX
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2022-09-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "HMX Gold",
					ShortDescription: "Clicky switch with a premium feel.",
					LongDescription:  "HMX Gold switches provide a clicky feel with premium materials, offering a luxurious typing experience with audible feedback.",
					ManufacturerID:   ptr(30), // HMX
					BrandID:          ptr(30), // HMX
					SwitchTypeID:     3,       // Clicky
					ReleaseDate:      parseDate("2023-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Poro",
					ShortDescription: "Linear switch with a smooth and silent action.",
					LongDescription:  "Poro switches offer a unique linear feel with a smooth and silent action, providing a premium typing experience for those who prefer quieter keyboards.",
					ManufacturerID:   ptr(30), // HMX
					BrandID:          ptr(30), // HMX
					SwitchTypeID:     4,       // Silent Linear
					ReleaseDate:      parseDate("2023-03-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Hyacinth V2",
					ShortDescription: "Tactile switch with a medium bump.",
					LongDescription:  "Hyacinth V2 switches provide a medium tactile bump with a smooth actuation, making them ideal for typists who enjoy a balanced tactile feel.",
					ManufacturerID:   ptr(30), // HMX
					BrandID:          ptr(30), // HMX
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2023-06-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Hyacinth V2U",
					ShortDescription: "Enhanced tactile switch with a unique actuation profile.",
					LongDescription:  "Hyacinth V2U switches offer an enhanced tactile experience with a unique actuation profile, providing a distinctive and satisfying typing experience for enthusiasts.",
					ManufacturerID:   ptr(30), // HMX
					BrandID:          ptr(30), // HMX
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2023-09-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
			}

			err = processSwitches(hmxSwitches)
			if err != nil {
				return err
			}

			kbdfansSwitches := []Switch{
				{
					Name:             "KBDFans T1",
					ShortDescription: "Tactile switch with a distinct bump.",
					LongDescription:  "KBDFans T1 switches are known for their strong tactile bump and smooth actuation, offering a unique typing experience.",
					ManufacturerID:   ptr(29), // KBDFans
					BrandID:          ptr(29), // KBDFans
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2019-06-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "KBDFans D65",
					ShortDescription: "Linear switch with smooth keypress.",
					LongDescription:  "The KBDFans D65 offers a linear experience with smooth keypresses, ideal for gamers and fast typists seeking consistency.",
					ManufacturerID:   ptr(29), // KBDFans
					BrandID:          ptr(29), // KBDFans
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2021-04-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "KBDFans Kiwis",
					ShortDescription: "Tactile switch with a smooth feel.",
					LongDescription:  "KBDFans Kiwis are tactile switches known for their smooth feel and strong tactile bump, providing an enjoyable typing experience.",
					ManufacturerID:   ptr(29), // KBDFans
					BrandID:          ptr(29), // KBDFans
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2020-11-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "KBDFans Mint",
					ShortDescription: "Linear switch with a refreshing color.",
					LongDescription:  "KBDFans Mint switches provide a smooth linear feel with a refreshing mint color, perfect for custom keyboard builds.",
					ManufacturerID:   ptr(29), // KBDFans
					BrandID:          ptr(29), // KBDFans
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2021-08-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "KBDFans Crystal",
					ShortDescription: "Silent linear switch with a smooth feel.",
					LongDescription:  "KBDFans Crystal switches offer a silent linear feel, providing a quiet yet smooth typing experience ideal for office use.",
					ManufacturerID:   ptr(29), // KBDFans
					BrandID:          ptr(29), // KBDFans
					SwitchTypeID:     4,       // Silent Linear
					ReleaseDate:      parseDate("2021-09-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "KBDFans Roller",
					ShortDescription: "Innovative switch with rolling actuation.",
					LongDescription:  "KBDFans Roller switches offer a unique rolling actuation mechanism, providing a novel tactile feedback and sound profile for a different typing experience.",
					ManufacturerID:   ptr(29), // KBDFans
					BrandID:          ptr(29), // KBDFans
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2023-05-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
			}
			err = processSwitches(kbdfansSwitches)
			if err != nil {
				return err
			}

			kineticLabsSwitches := []Switch{
				{
					Name:             "Kinetic Labs Hippo",
					ShortDescription: "Tactile switch with a medium actuation force.",
					LongDescription:  "The Kinetic Labs Hippo switch offers a delightful tactile bump with a smooth downstroke and upstroke, featuring a medium actuation force of 67g. Ideal for both typing and gaming.",
					ManufacturerID:   ptr(17), // Kinetic Labs
					BrandID:          ptr(17), // Kinetic Labs
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2021-03-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kinetic Labs Salmon",
					ShortDescription: "Linear switch with a smooth, medium-weight feel.",
					LongDescription:  "Kinetic Labs Salmon switches are linear switches with a buttery-smooth feel, medium weighting at 63.5g, perfect for users who enjoy a lighter linear experience.",
					ManufacturerID:   ptr(17), // Kinetic Labs
					BrandID:          ptr(17), // Kinetic Labs
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2021-06-15"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kinetic Labs Snow Globe V2",
					ShortDescription: "Silent linear switch with a smooth and quiet operation.",
					LongDescription:  "The Kinetic Labs Snow Globe V2 is a silent linear switch known for its smoothness and quiet operation, featuring a 63.5g actuation force. Great for office and shared spaces.",
					ManufacturerID:   ptr(17), // Kinetic Labs
					BrandID:          ptr(17), // Kinetic Labs
					SwitchTypeID:     4,       // Silent Linear
					ReleaseDate:      parseDate("2022-01-20"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kinetic Labs Polar Panda",
					ShortDescription: "Tactile switch with a distinct and pronounced bump.",
					LongDescription:  "Kinetic Labs Polar Panda switches feature a pronounced tactile bump with a 67g actuation force, making them ideal for tactile switch enthusiasts who prefer a noticeable feedback.",
					ManufacturerID:   ptr(22), // JWK
					BrandID:          ptr(17), // Kinetic Labs
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2022-08-10"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kinetic Labs Dragon Fruit",
					ShortDescription: "Linear switch with a medium actuation force and unique aesthetic.",
					LongDescription:  "The Kinetic Labs Dragon Fruit switch provides a smooth linear experience with a 67g actuation force, complemented by a vibrant, unique aesthetic.",
					ManufacturerID:   ptr(17), // Kinetic Labs
					BrandID:          ptr(17), // Kinetic Labs
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2023-05-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kinetic Labs Blueberry",
					ShortDescription: "Tactile switch with a medium-weight feel and crisp feedback.",
					LongDescription:  "Kinetic Labs Blueberry switches offer crisp tactile feedback with a medium-weight feel of 55g, ideal for typists who enjoy tactile sensations without excessive resistance.",
					ManufacturerID:   ptr(17), // Kinetic Labs
					BrandID:          ptr(17), // Kinetic Labs
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2023-07-15"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Kinetic Labs Penguin",
					ShortDescription: "Silent linear switch with smooth and quiet keystrokes.",
					LongDescription:  "The Kinetic Labs Penguin switch offers a silent linear experience with smooth and quiet keystrokes, featuring a 62g actuation force. Ideal for users who prefer a noiseless typing environment.",
					ManufacturerID:   ptr(17), // Kinetic Labs
					BrandID:          ptr(17), // Kinetic Labs
					SwitchTypeID:     4,       // Silent Linear
					ReleaseDate:      parseDate("2023-02-10"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Kinetic Labs Gecko Silent Linear",
					ShortDescription: "Ultra-silent linear switch with soft actuation and muted sound.",
					LongDescription:  "Kinetic Labs Gecko Silent Linear switches are ultra-silent linear switches designed for a quiet typing experience. With a soft actuation force of 50g, they provide muted sound and smooth keystrokes.",
					ManufacturerID:   ptr(17), // Kinetic Labs
					BrandID:          ptr(17), // Kinetic Labs
					SwitchTypeID:     4,       // Silent Linear
					ReleaseDate:      parseDate("2023-05-25"),
					Available:        true,
					PricePoint:       2, // Average
				},
			}

			err = processSwitches(kineticLabsSwitches)
			if err != nil {
				return err
			}

			chosfoxSwitches := []Switch{
				{
					Name:             "Chosfox Kiwi",
					ShortDescription: "Tactile switch with a strong bump and satisfying feedback.",
					LongDescription:  "The Chosfox Kiwi switch offers a pronounced tactile bump with a 67g actuation force, providing satisfying feedback for typists who enjoy tactile sensations.",
					ManufacturerID:   ptr(22), // JWK
					BrandID:          ptr(18), // Chosfox
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2021-06-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Chosfox Mint",
					ShortDescription: "Linear switch with a smooth and creamy feel.",
					LongDescription:  "The Chosfox Mint switch provides a smooth linear experience with a 63.5g actuation force, perfect for users who enjoy a creamy and seamless keystroke.",
					ManufacturerID:   ptr(22), // JWK
					BrandID:          ptr(18), // Chosfox
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2022-02-15"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Chosfox Blueberry",
					ShortDescription: "Tactile switch with a medium-weight feel and crisp feedback.",
					LongDescription:  "Chosfox Blueberry switches offer a medium-weight tactile feel with crisp feedback, featuring a 65g actuation force. Ideal for typists who appreciate tactile sensations without excessive resistance.",
					ManufacturerID:   ptr(22), // JWK
					BrandID:          ptr(18), // Chosfox
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2022-07-10"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Chosfox Chameleon",
					ShortDescription: "Linear switch with a unique aesthetic and smooth operation.",
					LongDescription:  "The Chosfox Chameleon switch offers a unique aesthetic with smooth linear action, providing a 68g actuation force for users who prefer a consistent and fluid typing experience.",
					ManufacturerID:   ptr(22), // JWK
					BrandID:          ptr(18), // Chosfox
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2023-01-20"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Chosfox Marshmallow",
					ShortDescription: "Silent linear switch with a soft and quiet typing experience.",
					LongDescription:  "Chosfox Marshmallow switches provide a silent linear experience with soft and quiet keystrokes, featuring a 58g actuation force. Ideal for users who need a noiseless typing environment.",
					ManufacturerID:   ptr(22), // JWK
					BrandID:          ptr(18), // Chosfox
					SwitchTypeID:     4,       // Silent Linear
					ReleaseDate:      parseDate("2023-06-05"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Chosfox Watermelon",
					ShortDescription: "Tactile switch with a fruity feel and moderate resistance.",
					LongDescription:  "The Chosfox Watermelon switch offers a tactile experience with a fruity feel and moderate resistance, providing a 62g actuation force. Perfect for users who enjoy a unique tactile sensation.",
					ManufacturerID:   ptr(22), // JWK
					BrandID:          ptr(18), // Chosfox
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2023-09-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Chosfox Grape",
					ShortDescription: "Linear switch with a heavier actuation force and smooth travel.",
					LongDescription:  "Chosfox Grape switches offer a smooth linear action with a heavier actuation force of 70g, ideal for users who prefer more substantial resistance in their keystrokes.",
					ManufacturerID:   ptr(22), // JWK
					BrandID:          ptr(18), // Chosfox
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2024-01-10"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Chosfox Silent Forest",
					ShortDescription: "Silent tactile switch with a gentle bump and quiet operation.",
					LongDescription:  "The Chosfox Silent Forest switch offers a silent tactile experience with a gentle bump and quiet operation, providing a 62g actuation force. Perfect for users who want tactile feedback without noise.",
					ManufacturerID:   ptr(22), // JWK
					BrandID:          ptr(18), // Chosfox
					SwitchTypeID:     5,       // Silent Tactile
					ReleaseDate:      parseDate("2024-03-20"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Chosfox Banana Split",
					ShortDescription: "Linear switch with a smooth feel and playful aesthetic.",
					LongDescription:  "Chosfox Banana Split switches offer a smooth linear feel with a playful aesthetic, featuring a 62g actuation force. Ideal for users who enjoy a fun and seamless typing experience.",
					ManufacturerID:   ptr(22), // JWK
					BrandID:          ptr(18), // Chosfox
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2024-05-15"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Chosfox Sunset",
					ShortDescription: "Linear switch with a vibrant color scheme and consistent action.",
					LongDescription:  "The Chosfox Sunset switch offers a vibrant color scheme with consistent linear action, providing a 67g actuation force. Perfect for users who prefer a reliable and visually appealing typing experience.",
					ManufacturerID:   ptr(22), // JWK
					BrandID:          ptr(18), // Chosfox
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2024-07-10"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Chosfox White Fox",
					ShortDescription: "Tactile switch with a smooth bump and quiet operation.",
					LongDescription:  "Chosfox White Fox switches offer a smooth tactile bump and quiet operation, featuring a 58g actuation force. Ideal for users who enjoy a soft tactile feel with minimal noise.",
					ManufacturerID:   ptr(3),  // Kailh
					BrandID:          ptr(18), // Chosfox
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2023-01-15"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Chosfox Arctic Fox",
					ShortDescription: "Linear switch with a cool feel and light actuation force.",
					LongDescription:  "Chosfox Arctic Fox switches provide a cool linear feel with a light 48g actuation force. Perfect for users who prefer a gentle and smooth keystroke.",
					ManufacturerID:   ptr(3),  // Kailh
					BrandID:          ptr(18), // Chosfox
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2023-02-10"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Chosfox Voyager V2",
					ShortDescription: "Linear switch designed for smooth and consistent travel.",
					LongDescription:  "Chosfox Voyager V2 switches offer a smooth and consistent linear travel with a 65g actuation force, providing a balanced and fluid typing experience.",
					ManufacturerID:   ptr(22), // JWK
					BrandID:          ptr(18), // Chosfox
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2023-04-20"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Chosfox Hanami Dango",
					ShortDescription: "Linear switch with soft feel and pastel aesthetics.",
					LongDescription:  "Chosfox Hanami Dango switches offer a soft linear feel with pastel aesthetics, featuring a 50g actuation force. Perfect for users who appreciate a gentle touch and pleasing visuals.",
					ManufacturerID:   ptr(22), // JWK
					BrandID:          ptr(18), // Chosfox
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2023-05-15"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Chosfox Summer Lime Silent",
					ShortDescription: "Silent linear switch with a soft and quiet typing experience.",
					LongDescription:  "Chosfox Summer Lime Silent switches provide a silent linear experience with a soft and quiet 40g actuation force, ideal for users who need a noiseless typing environment.",
					ManufacturerID:   ptr(22), // JWK
					BrandID:          ptr(18), // Chosfox
					SwitchTypeID:     4,       // Silent Linear
					ReleaseDate:      parseDate("2023-07-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Chosfox Poison Gas",
					ShortDescription: "Tactile switch with strong feedback and vibrant color scheme.",
					LongDescription:  "Chosfox Poison Gas switches offer a strong tactile feedback with a vibrant color scheme, featuring a 67g actuation force. Ideal for users who enjoy pronounced tactile sensations.",
					ManufacturerID:   ptr(22), // JWK
					BrandID:          ptr(18), // Chosfox
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2023-08-10"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Chosfox DD Jingle Linear",
					ShortDescription: "Linear switch with a unique auditory profile and smooth action.",
					LongDescription:  "Chosfox DD Jingle Linear switches provide a unique auditory profile with smooth linear action, featuring a 55g actuation force. Perfect for users who enjoy distinct sounds and smooth keystrokes.",
					ManufacturerID:   ptr(22), // JWK
					BrandID:          ptr(18), // Chosfox
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2023-09-05"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
			}
			err = processSwitches(chosfoxSwitches)
			if err != nil {
				return err
			}

			wuqueStudioSwitches := []Switch{
				{
					Name:             "Wuque Studio Aurora",
					ShortDescription: "Linear switch with a smooth feel and vibrant color scheme.",
					LongDescription:  "The Wuque Studio Aurora switch provides a smooth linear experience with a 63.5g actuation force, offering a vibrant color scheme for an aesthetically pleasing keyboard setup.",
					ManufacturerID:   ptr(22), // JWK
					BrandID:          ptr(21), // Wuque Studio
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2022-04-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Wuque Studio Polar",
					ShortDescription: "Tactile switch with a strong bump and icy aesthetic.",
					LongDescription:  "The Wuque Studio Polar switch offers a pronounced tactile bump with a 67g actuation force, providing an icy aesthetic and satisfying feedback for tactile enthusiasts.",
					ManufacturerID:   ptr(22), // JWK
					BrandID:          ptr(21), // Wuque Studio
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2022-07-15"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Wuque Studio Iceberg",
					ShortDescription: "Silent linear switch with smooth and quiet operation.",
					LongDescription:  "Wuque Studio Iceberg switches provide a silent linear experience with smooth and quiet keystrokes, featuring a 60g actuation force. Perfect for users who need a noiseless typing environment.",
					ManufacturerID:   ptr(22), // JWK
					BrandID:          ptr(21), // Wuque Studio
					SwitchTypeID:     4,       // Silent Linear
					ReleaseDate:      parseDate("2023-01-20"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Wuque Studio Ocean Blue",
					ShortDescription: "Clicky switch with crisp feedback and oceanic feel.",
					LongDescription:  "The Wuque Studio Ocean Blue switch offers a crisp clicky experience with a 50g actuation force, providing a satisfying feedback for users who enjoy audible feedback.",
					ManufacturerID:   ptr(22), // JWK
					BrandID:          ptr(21), // Wuque Studio
					SwitchTypeID:     3,       // Clicky
					ReleaseDate:      parseDate("2023-05-10"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Wuque Studio Forest Green",
					ShortDescription: "Tactile switch with a natural feel and moderate resistance.",
					LongDescription:  "Wuque Studio Forest Green switches offer a tactile experience with a natural feel, providing moderate resistance with a 62g actuation force. Ideal for users who prefer tactile feedback with a nature-inspired design.",
					ManufacturerID:   ptr(22), // JWK
					BrandID:          ptr(21), // Wuque Studio
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2023-10-05"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Wuque Studio Jade Linear",
					ShortDescription: "Linear switch with a smooth and consistent action.",
					LongDescription:  "Wuque Studio Jade Linear switches provide a smooth and consistent linear action with a 45g actuation force. Ideal for users who prefer a gentle and fluid typing experience.",
					ManufacturerID:   ptr(21), // Wuque Studio
					BrandID:          ptr(21), // Wuque Studio
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2023-03-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Wuque Studio Heavy Tactile",
					ShortDescription: "Tactile switch with a strong bump and substantial resistance.",
					LongDescription:  "Wuque Studio Heavy Tactile switches offer a strong tactile bump with substantial resistance, featuring a 70g actuation force. Perfect for users who enjoy pronounced tactile feedback.",
					ManufacturerID:   ptr(21), // Wuque Studio
					BrandID:          ptr(21), // Wuque Studio
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2023-04-15"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Wuque Studio Aurora Linear",
					ShortDescription: "Smooth linear switch with vibrant aesthetics.",
					LongDescription:  "Wuque Studio Aurora Linear switches provide a smooth linear experience with vibrant aesthetics, featuring a 63.5g actuation force. Ideal for users who appreciate a visually appealing and fluid typing experience.",
					ManufacturerID:   ptr(21), // Wuque Studio
					BrandID:          ptr(21), // Wuque Studio
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2023-06-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Wuque Studio WS Silent",
					ShortDescription: "Ultra-silent linear switch with muted sound.",
					LongDescription:  "Wuque Studio WS Silent switches provide an ultra-silent linear experience with muted sound, featuring a 60g actuation force. Ideal for users who require a quiet typing environment.",
					ManufacturerID:   ptr(21), // Wuque Studio
					BrandID:          ptr(21), // Wuque Studio
					SwitchTypeID:     4,       // Silent Linear
					ReleaseDate:      parseDate("2023-07-10"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Wuque Studio POM+",
					ShortDescription: "Linear switch with POM material for smooth keystrokes.",
					LongDescription:  "Wuque Studio POM+ switches feature POM material for exceptionally smooth linear keystrokes, with a 55g actuation force. Ideal for users who enjoy a seamless typing experience.",
					ManufacturerID:   ptr(21), // Wuque Studio
					BrandID:          ptr(21), // Wuque Studio
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2023-09-15"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Wuque Studio Light Tactile",
					ShortDescription: "Tactile switch with a gentle bump and light resistance.",
					LongDescription:  "Wuque Studio Light Tactile switches provide a gentle tactile bump with light resistance, featuring a 50g actuation force. Perfect for users who prefer a subtle tactile feedback.",
					ManufacturerID:   ptr(21), // Wuque Studio
					BrandID:          ptr(21), // Wuque Studio
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2023-10-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Wuque Studio Onion Linear",
					ShortDescription: "Linear switch with smooth travel and a unique aesthetic.",
					LongDescription:  "Wuque Studio Onion Linear switches offer smooth travel with a unique aesthetic, featuring a 52g actuation force. Ideal for users who appreciate a distinctive look and feel.",
					ManufacturerID:   ptr(21), // Wuque Studio
					BrandID:          ptr(21), // Wuque Studio
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2023-11-10"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Wuque Studio Morandi Linear",
					ShortDescription: "Linear switch with subtle elegance and smooth action.",
					LongDescription:  "Wuque Studio Morandi Linear switches provide a subtle elegance with smooth linear action, featuring a 62g actuation force. Ideal for users who prefer a refined and fluid typing experience.",
					ManufacturerID:   ptr(21), // Wuque Studio
					BrandID:          ptr(21), // Wuque Studio
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2024-01-05"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Wuque Studio Quartz Linear",
					ShortDescription: "Linear switch with a clear, transparent aesthetic.",
					LongDescription:  "Wuque Studio Quartz Linear switches offer a clear, transparent aesthetic with smooth linear action, featuring a 55g actuation force. Perfect for users who enjoy a visually striking and seamless typing experience.",
					ManufacturerID:   ptr(21), // Wuque Studio
					BrandID:          ptr(21), // Wuque Studio
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2024-03-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Wuque Studio Standard Red",
					ShortDescription: "Linear switch with a light actuation force and smooth travel.",
					LongDescription:  "Wuque Studio Standard Red switches provide a smooth linear experience with a 45g actuation force. Ideal for users who prefer light and fluid keystrokes.",
					ManufacturerID:   ptr(21), // Wuque Studio
					BrandID:          ptr(21), // Wuque Studio
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2023-08-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Wuque Studio Standard Yellow",
					ShortDescription: "Clicky switch with crisp feedback and moderate resistance.",
					LongDescription:  "Wuque Studio Standard Yellow switches offer a satisfying clicky experience with a 55g actuation force. Perfect for users who enjoy audible feedback and moderate resistance.",
					ManufacturerID:   ptr(21), // Wuque Studio
					BrandID:          ptr(21), // Wuque Studio
					SwitchTypeID:     3,       // Clicky
					ReleaseDate:      parseDate("2023-08-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Wuque Studio Standard Brown",
					ShortDescription: "Tactile switch with a pronounced bump and firm feedback.",
					LongDescription:  "Wuque Studio Standard Brown switches provide a tactile experience with a pronounced bump and firm feedback, featuring a 65g actuation force. Ideal for users who enjoy a strong tactile feel.",
					ManufacturerID:   ptr(21), // Wuque Studio
					BrandID:          ptr(21), // Wuque Studio
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2023-08-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
			}
			err = processSwitches(wuqueStudioSwitches)
			if err != nil {
				return err
			}

			tbcatSwitches := []Switch{
				{
					Name:             "Tbcats Phoenix",
					ShortDescription: "Linear switch with a smooth glide and moderate actuation force.",
					LongDescription:  "The Tbcats Phoenix switch offers a linear feel with a smooth glide and a moderate actuation force of around 45g, ideal for fast typing and gaming.",
					ManufacturerID:   ptr(31), // Tbcats Studio
					BrandID:          ptr(31), // Tbcats Studio
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2022-09-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Tbcats Nebula",
					ShortDescription: "Tactile switch with a subtle bump and quiet operation.",
					LongDescription:  "The Tbcats Nebula switch is a tactile switch that offers a subtle bump at the actuation point, providing a quiet and pleasant typing experience with an actuation force of 55g.",
					ManufacturerID:   ptr(31), // Tbcats Studio
					BrandID:          ptr(31), // Tbcats Studio
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2023-02-15"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Tbcats Thunder",
					ShortDescription: "Clicky switch with a pronounced tactile bump and click sound.",
					LongDescription:  "The Tbcats Thunder switch is a clicky switch that delivers a pronounced tactile bump and an audible click sound, making it perfect for those who enjoy a classic mechanical keyboard feel.",
					ManufacturerID:   ptr(31), // Tbcats Studio
					BrandID:          ptr(31), // Tbcats Studio
					SwitchTypeID:     3,       // Clicky
					ReleaseDate:      parseDate("2023-05-10"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Tbcats Eclipse",
					ShortDescription: "Silent linear switch with a soft and smooth keystroke.",
					LongDescription:  "The Tbcats Eclipse switch is a silent linear switch designed for those who prefer a quiet typing experience without sacrificing the smoothness and softness of a linear action.",
					ManufacturerID:   ptr(31), // Tbcats Studio
					BrandID:          ptr(31), // Tbcats Studio
					SwitchTypeID:     4,       // Silent Linear
					ReleaseDate:      parseDate("2023-08-05"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Tbcats Cosmic",
					ShortDescription: "Silent tactile switch with a gentle tactile bump.",
					LongDescription:  "The Tbcats Cosmic switch offers a silent tactile experience with a gentle bump at the actuation point, ideal for those who enjoy tactile feedback without the noise.",
					ManufacturerID:   ptr(31), // Tbcats Studio
					BrandID:          ptr(31), // Tbcats Studio
					SwitchTypeID:     5,       // Silent Tactile
					ReleaseDate:      parseDate("2024-01-20"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Tbcats Cloud Water V2",
					ShortDescription: "Smooth linear switch with a soft and quiet keystroke.",
					LongDescription:  "The Tbcats Cloud Water V2 is a smooth linear switch that provides a soft and quiet keystroke, perfect for those who enjoy a seamless and silent typing experience.",
					ManufacturerID:   ptr(31), // Tbcats Studio
					BrandID:          ptr(31), // Tbcats Studio
					SwitchTypeID:     4,       // Silent Linear
					ReleaseDate:      parseDate("2023-04-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Tbcats Eclair Green",
					ShortDescription: "Tactile switch with a medium bump and satisfying feedback.",
					LongDescription:  "The Tbcats Eclair Green switch offers a tactile experience with a medium bump, providing satisfying feedback and responsiveness for both typing and gaming.",
					ManufacturerID:   ptr(31), // Tbcats Studio
					BrandID:          ptr(31), // Tbcats Studio
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2023-03-15"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Tbcats Blue Balloon VW",
					ShortDescription: "Clicky switch with a unique sound profile and sharp feedback.",
					LongDescription:  "The Tbcats Blue Balloon VW switch is known for its unique clicky sound profile and sharp feedback, offering an enjoyable typing experience with a distinctive sound.",
					ManufacturerID:   ptr(31), // Tbcats Studio
					BrandID:          ptr(31), // Tbcats Studio
					SwitchTypeID:     3,       // Clicky
					ReleaseDate:      parseDate("2024-06-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Tbcats Eclair Mix",
					ShortDescription: "Hybrid switch offering a blend of tactile and linear characteristics.",
					LongDescription:  "The Tbcats Eclair Mix switch combines tactile and linear characteristics, providing a unique typing experience with a balance of smoothness and feedback.",
					ManufacturerID:   ptr(31), // Tbcats Studio
					BrandID:          ptr(31), // Tbcats Studio
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2023-07-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Tbcats Eclair Orange",
					ShortDescription: "Tactile switch with a pronounced bump and strong feedback.",
					LongDescription:  "The Tbcats Eclair Orange switch offers a pronounced tactile bump and strong feedback, making it ideal for users who prefer noticeable tactile sensations in their typing.",
					ManufacturerID:   ptr(31), // Tbcats Studio
					BrandID:          ptr(31), // Tbcats Studio
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2023-11-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
			}
			err = processSwitches(tbcatSwitches)
			if err != nil {
				return err
			}

			c3EqualzTkcSwitches := []Switch{
				{
					Name:             "C³ Equalz x TKC Banana Split",
					ShortDescription: "Linear switch with a thocky sound profile and smooth action.",
					LongDescription:  "The C³ Equalz x TKC Banana Split switch is a linear switch known for its thocky sound profile and smooth action. It's part of the Macho Linear series, providing a satisfying typing experience for enthusiasts who love a deep thock sound.",
					ManufacturerID:   ptr(22), // JWK
					BrandID:          ptr(32), // Collaboration: C³ Equalz x TKC
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2021-07-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "C³ Equalz x TKC Tangerine 67 (Dark Green)",
					ShortDescription: "Smooth linear switch with a 67g actuation force.",
					LongDescription:  "The C³ Equalz x TKC Tangerine 67 switch features a dark green housing with a smooth linear action and a 67g actuation force. Known for its thocky sound and excellent build quality, this switch is perfect for those who enjoy a responsive linear feel.",
					ManufacturerID:   ptr(22), // JWK
					BrandID:          ptr(32), // Collaboration: C³ Equalz x TKC
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2020-05-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "C³ Equalz x TKC Tangerine 62 (Light Green)",
					ShortDescription: "Linear switch with a light 62g actuation force.",
					LongDescription:  "The C³ Equalz x TKC Tangerine 62 switch is a linear switch with a light green housing, offering a light 62g actuation force. It is pre-lubed for extra smoothness and delivers a pleasant typing experience with a satisfying sound profile.",
					ManufacturerID:   ptr(22), // JWK
					BrandID:          ptr(32), // Collaboration: C³ Equalz x TKC
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2020-05-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "C³ Equalz x TKC Kiwi",
					ShortDescription: "Tactile switch with a pronounced bump and 67g force.",
					LongDescription:  "The C³ Equalz x TKC Kiwi switch is a tactile switch featuring a pronounced bump and a 67g actuation force. This switch offers a strong tactile feedback, making it a favorite for users who enjoy a tactile typing experience.",
					ManufacturerID:   ptr(22), // JWK
					BrandID:          ptr(32), // Collaboration: C³ Equalz x TKC
					SwitchTypeID:     2,       // Tactile
					ReleaseDate:      parseDate("2021-02-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "C³ Equalz x TKC Red Smoothie",
					ShortDescription: "Linear fruit switch with a smooth action.",
					LongDescription:  "The C³ Equalz x TKC Red Smoothie switch is a linear fruit switch known for its smooth action and vibrant red color. It's designed to provide a delightful typing experience with a consistent feel.",
					ManufacturerID:   ptr(22), // JWK
					BrandID:          ptr(32), // Collaboration: C³ Equalz x TKC
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2023-01-15"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "C³ Equalz x TKC Blue Smoothie",
					ShortDescription: "Linear fruit switch with 67g actuation force.",
					LongDescription:  "The C³ Equalz x TKC Blue Smoothie switch offers a linear action with a 67g actuation force, designed as a part of the linear fruit switch series. This switch provides a smooth typing experience with a pleasant sound profile.",
					ManufacturerID:   ptr(22), // JWK
					BrandID:          ptr(32), // Collaboration: C³ Equalz x TKC
					SwitchTypeID:     1,       // Linear
					ReleaseDate:      parseDate("2023-01-15"),
					Available:        true,
					PricePoint:       2, // Average
				},
			}

			err = processSwitches(c3EqualzTkcSwitches)
			if err != nil {
				return err
			}

			cherryMX2ASwitches := []Switch{
				{
					Name:             "Cherry MX2A Speed Silver RGB Linear",
					ShortDescription: "Linear switch with fast actuation and RGB support.",
					LongDescription:  "The Cherry MX2A Speed Silver RGB Linear switch is designed for quick actuation with a light touch, making it ideal for gaming. It also features RGB support for customizable lighting effects.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2024-03-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Cherry MX2A Brown Tactile",
					ShortDescription: "Tactile switch with a gentle bump for feedback.",
					LongDescription:  "The Cherry MX2A Brown Tactile switch offers a gentle tactile bump, providing noticeable feedback for typing without the click sound, making it a great choice for both typing and gaming.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2024-03-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Cherry MX2A Blue Clicky",
					ShortDescription: "Clicky switch with an audible click for tactile feedback.",
					LongDescription:  "The Cherry MX2A Blue Clicky switch delivers a crisp click sound and tactile feedback, perfect for typists who appreciate auditory feedback while typing.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2024-03-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Cherry MX2A Blue Clicky RGB",
					ShortDescription: "Clicky switch with RGB lighting and tactile feedback.",
					LongDescription:  "The Cherry MX2A Blue Clicky RGB switch combines the satisfying click and tactile feedback of a clicky switch with RGB lighting, allowing for vibrant and customizable keyboard backlighting.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2024-03-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Cherry MX2A Black Linear",
					ShortDescription: "Linear switch with a heavy actuation force.",
					LongDescription:  "The Cherry MX2A Black Linear switch offers a smooth keystroke with a heavier actuation force, providing a firm and controlled typing experience suitable for heavy typists and gamers.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2024-03-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Cherry MX2A Silent Red / Pink RGB Linear",
					ShortDescription: "Silent linear switch with RGB support and a smooth feel.",
					LongDescription:  "The Cherry MX2A Silent Red / Pink RGB Linear switch features a quiet and smooth keystroke, enhanced with RGB lighting for colorful customization, ideal for quiet environments.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2024-03-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Cherry MX2A Black RGB Linear",
					ShortDescription: "Linear switch with RGB lighting and a heavy feel.",
					LongDescription:  "The Cherry MX2A Black RGB Linear switch offers a heavier actuation force combined with RGB lighting, providing a robust and visually appealing typing experience.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2024-03-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Cherry MX2A Speed Silver Linear",
					ShortDescription: "Fast linear switch for quick actuation.",
					LongDescription:  "The Cherry MX2A Speed Silver Linear switch is engineered for rapid actuation, making it perfect for gamers who need a quick and responsive keystroke.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2024-03-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Cherry MX2A Red RGB Linear",
					ShortDescription: "Linear switch with RGB support and a smooth keystroke.",
					LongDescription:  "The Cherry MX2A Red RGB Linear switch combines a smooth keystroke with RGB lighting, providing both visual appeal and a satisfying typing experience.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2024-03-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Cherry MX2A Silent Red / Pink Linear",
					ShortDescription: "Silent linear switch for a quiet typing experience.",
					LongDescription:  "The Cherry MX2A Silent Red / Pink Linear switch is designed for a quiet typing experience without sacrificing smoothness, making it perfect for office settings or late-night typing.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2024-03-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Cherry MX2A Red Linear",
					ShortDescription: "Smooth linear switch for a balanced typing feel.",
					LongDescription:  "The Cherry MX2A Red Linear switch offers a balanced and smooth typing experience, suitable for both gaming and typing due to its medium actuation force.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2024-03-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
			}

			err = processSwitches(cherryMX2ASwitches)
			if err != nil {
				return err
			}
			return nil // Replace with actual code
		},
		Rollback: func(tx *gorm.DB) error {
			// Your rollback logic goes here.
			return nil // Replace with actual code
		},
	})
}
