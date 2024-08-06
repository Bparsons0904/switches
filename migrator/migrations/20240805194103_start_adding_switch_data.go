package migrations

import (
	"time"

	"github.com/Bparsons0904/deadigations"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	deadigations.RegisterMigration(deadigations.Migration{
		ID:          "20240805194103",
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

			switches := []Switch{
				{
					Name:             "Cherry MX Red",
					ShortDescription: "Linear switch with a light actuation force.",
					LongDescription:  "Cherry MX Red switches are known for their smooth and consistent feel with no tactile bump, making them ideal for gaming and fast typing.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("1984-01-01"),
					Available:        true,
					PricePoint:       2,
				},
				{
					Name:             "Cherry MX Blue",
					ShortDescription: "Clicky switch with a tactile bump.",
					LongDescription:  "Cherry MX Blue switches provide a clicky and tactile feedback, making them popular among typists who enjoy an audible click.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("1987-01-01"),
					Available:        true,
					PricePoint:       2,
				},
				{
					Name:             "Cherry MX Brown",
					ShortDescription: "Tactile switch with a moderate actuation force.",
					LongDescription:  "Cherry MX Brown switches are versatile, offering a tactile bump without an audible click, suitable for both typing and gaming.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("1994-01-01"),
					Available:        true,
					PricePoint:       2,
				},
				{
					Name:             "Cherry MX Black",
					ShortDescription: "Linear switch with a heavier actuation force.",
					LongDescription:  "Cherry MX Black switches are known for their smooth linear feel with a higher actuation force, favored by users who prefer a stiffer switch.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("1984-01-01"),
					Available:        true,
					PricePoint:       2,
				},
				{
					Name:             "Cherry MX Green",
					ShortDescription: "Clicky switch with a heavy actuation force.",
					LongDescription:  "Cherry MX Green switches are similar to MX Blue but with a higher actuation force, providing a clicky and tactile experience with more resistance.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2011-01-01"),
					Available:        true,
					PricePoint:       3,
				},
				{
					Name:             "Cherry MX Silent Red",
					ShortDescription: "Silent linear switch with a light actuation force.",
					LongDescription:  "Cherry MX Silent Red switches offer a smooth linear feel with dampened noise, ideal for quiet environments and gaming.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2015-01-01"),
					Available:        true,
					PricePoint:       3,
				},
				{
					Name:             "Cherry MX Silent Black",
					ShortDescription: "Silent linear switch with a heavier actuation force.",
					LongDescription:  "Cherry MX Silent Black switches provide a quiet, smooth linear feel with more resistance, suitable for those who prefer a heavier switch.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2015-01-01"),
					Available:        true,
					PricePoint:       3,
				},
				{
					Name:             "Cherry MX Clear",
					ShortDescription: "Tactile switch with a high actuation force.",
					LongDescription:  "Cherry MX Clear switches provide a tactile bump with a higher actuation force, ideal for those who prefer a noticeable resistance.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("1994-01-01"),
					Available:        true,
					PricePoint:       3,
				},
				{
					Name:             "Cherry MX Grey",
					ShortDescription: "Tactile switch with a heavy actuation force.",
					LongDescription:  "Cherry MX Grey switches are similar to MX Clear but with an even higher actuation force, providing a pronounced tactile bump.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("1994-01-01"),
					Available:        true,
					PricePoint:       3,
				},
				// New Cherry Switches
				{
					Name:             "Cherry MX Speed Silver",
					ShortDescription: "Linear switch with a shorter travel distance for faster actuation.",
					LongDescription:  "Cherry MX Speed Silver switches are designed for rapid keystrokes, offering a linear feel with a reduced travel distance and actuation point for quick responsiveness.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2016-01-01"),
					Available:        true,
					PricePoint:       3,
				},
				{
					Name:             "Cherry MX Low Profile Red",
					ShortDescription: "Slimmer version of the Red switch for low-profile keyboards.",
					LongDescription:  "Cherry MX Low Profile Red switches maintain the linear feel of the original Red switches but in a slimmer form factor, making them suitable for low-profile keyboards.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2018-01-01"),
					Available:        true,
					PricePoint:       3,
				},
				{
					Name:             "Cherry MX Low Profile Speed",
					ShortDescription: "Low-profile version of the Speed Silver switch.",
					LongDescription:  "Cherry MX Low Profile Speed switches combine the rapid actuation of the Speed Silver with a low-profile design, offering both quick responsiveness and a sleek form factor.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2018-01-01"),
					Available:        true,
					PricePoint:       3,
				},
				{
					Name:             "Cherry MX Nature White",
					ShortDescription: "Linear switch with a medium actuation force, between Red and Black.",
					LongDescription:  "Cherry MX Nature White switches offer a linear feel with an actuation force that sits between the lighter Red and the heavier Black switches, providing a balanced typing experience.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2016-01-01"),
					Available:        true,
					PricePoint:       2,
				},
				{
					Name:             "Cherry MX Violet",
					ShortDescription: "Tactile switch with a light actuation force, similar to Brown but lighter.",
					LongDescription:  "Cherry MX Violet switches provide a light tactile bump, similar to the MX Brown switches, but with a reduced actuation force for a gentler typing experience.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       2,
				},
				// Discontinued Cherry Switches
				{
					Name:             "Cherry MX White",
					ShortDescription: "A softer clicky switch, sometimes called 'Milk'.",
					LongDescription:  "Cherry MX White switches, also known as 'Milk' switches, offer a softer click compared to the MX Blue, providing a quieter yet still clicky experience.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2011-01-01"),
					Available:        false,
					PricePoint:       2,
				},
				{
					Name:             "Cherry MX Tactile Grey",
					ShortDescription: "A variant of the Grey switch with a slightly different force curve.",
					LongDescription:  "Cherry MX Tactile Grey switches provide a unique tactile feel with a different force curve compared to the standard Grey switches, offering an alternative tactile experience.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("1994-01-01"),
					Available:        false,
					PricePoint:       3,
				},
				{
					Name:             "Cherry MX Super Black",
					ShortDescription: "An extremely heavy linear switch.",
					LongDescription:  "Cherry MX Super Black switches are designed for users who prefer an extremely heavy linear feel, providing a significant resistance to prevent accidental key presses.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("1984-01-01"),
					Available:        false,
					PricePoint:       3,
				},
			}

			for _, s := range switches {
				if err := tx.Create(&s).Error; err != nil {
					return err
				}

				// Sleep for 5 ms to ensure unique timestamps
				time.Sleep(5 * time.Millisecond)
			}

			return nil // Replace with actual code
		},
		Rollback: func(tx *gorm.DB) error {
			// Your rollback logic goes here.
			return nil // Replace with actual code
		},
	})
}

func ptr(i int) *int {
	return &i
}

func parseDate(date string) *time.Time {
	t, _ := time.Parse("2006-01-02", date)
	return &t
}
