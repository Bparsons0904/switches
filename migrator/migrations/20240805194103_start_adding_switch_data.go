package migrations

import (
	"log"
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
			type ImageLink struct {
				ID          uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v7();primaryKey" json:"id"`
				LinkAddress string         `gorm:"type:varchar(255);not null"                      json:"linkAddress"`
				AltText     string         `gorm:"type:varchar(255);default:''"                    json:"altText,omitempty"`
				OwnerID     uuid.UUID      `gorm:"type:uuid;indes:idx_type_id"                     json:"objectId"`
				OwnerType   string         `gorm:"type:varchar(50);index:idx_type_id"              json:"objectType"`
				CreatedAt   time.Time      `gorm:"autoCreateTime"                                  json:"createdAt"`
				UpdatedAt   time.Time      `gorm:"autoUpdateTime"                                  json:"updatedAt"`
				DeletedAt   gorm.DeletedAt `gorm:""                                                json:"deletedAt"`
			}

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
				ID               uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v7();primaryKey"                                          json:"id"`
				Name             string       `gorm:"type:varchar(50);not null;index:idx_name;uniqueIndex:idx_name_manufacturer_brand"         json:"name"`
				ShortDescription string       `gorm:"type:varchar(255);not null"                                                               json:"shortDescription"`
				LongDescription  string       `gorm:"type:text;not null"                                                                       json:"longDescription"`
				ManufacturerID   *int         `gorm:"type:int;index;uniqueIndex:idx_name_manufacturer_brand"                                   json:"manufacturerId,omitempty"`
				Manufacturer     *Producer    `gorm:"foreignKey:ManufacturerID"                                                                json:"manufacturer,omitempty"`
				BrandID          *int         `gorm:"type:int;index;uniqueIndex:idx_name_manufacturer_brand"                                   json:"brandId,omitempty"`
				Brand            *Producer    `gorm:"foreignKey:BrandID"                                                                       json:"brand,omitempty"`
				SwitchTypeID     int          `gorm:"type:int;not null;index;"                                                                 json:"switchTypeId"`
				SwitchType       *Type        `gorm:"foreignKey:SwitchTypeID"                                                                  json:"switchType,omitempty"`
				ReleaseDate      *time.Time   `gorm:"type:date"                                                                                json:"releaseDate,omitempty"`
				Available        bool         `gorm:"type:boolean;default:true"                                                                json:"available"`
				PricePoint       int          `gorm:"type:int;not null"                                                                        json:"pricePoint"`
				SiteURL          string       `gorm:"type:varchar(255)"                                                                        json:"siteURL,omitempty"`
				CreatedAt        time.Time    `gorm:"autoCreateTime"                                                                           json:"createdAt"`
				UpdatedAt        time.Time    `gorm:"autoUpdateTime"                                                                           json:"updatedAt"`
				ImageLinks       []*ImageLink `gorm:"foreignKey:OwnerID;polymorphicValue:Switch;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"imageLinks,omitempty"`
			}

			ptr := func(i int) *int {
				return &i
			}

			parseDate := func(date string) *time.Time {
				t, _ := time.Parse("2006-01-02", date)
				return &t
			}

			cherry := []Switch{
				{
					Name:             "Cherry MX2A Silent Red",
					ShortDescription: "Silent linear switch with RGB support and a smooth feel.",
					LongDescription:  "The Cherry MX2A Silent Red Linear switch features a quiet and smooth keystroke, making it perfect for quiet environments such as offices or shared spaces. With an actuation force of 45 grams, this switch ensures a light touch while minimizing noise through built-in dampeners. The RGB lighting adds a customizable visual element, making it an excellent choice for those who want both a serene typing experience and a vibrant keyboard.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2024-03-01"),
					Available:        true,
					PricePoint:       3, // Expensive
					SiteURL:          "https://www.cherry-world.com/mx2a-silent-red",
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx3a-l1na_l1nn_image.jpg",
							AltText:     "Cherry MX2A Silent Red switch",
						},
						{
							LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx3a-l1nn_image.jpg",
							AltText:     "Cherry MX2A Silent Red switch",
						},
						{
							LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx3a-l1na_image.jpg",
							AltText:     "Cherry MX2A Silent Red switch",
						},
					},
				},
				{
					Name:             "Cherry MX Silent Black",
					ShortDescription: "Silent linear switch with a heavier actuation force.",
					LongDescription:  "Cherry MX Silent Black switches are engineered to provide the quiet operation of Silent Reds with a higher actuation force of 60 grams. This switch is ideal for users who prefer a more deliberate keystroke, combining the smooth, linear action with noise-reducing technology. Silent Black switches are well-suited for environments that require both quiet and a firm keypress, balancing performance with sound reduction.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2015-01-01"),
					Available:        false,
					PricePoint:       3,
					SiteURL:          "https://www.cherry-world.com/mx-silent-black.html",
				},
				{
					Name:             "Cherry MX Silent Red",
					ShortDescription: "Silent linear switch with a light actuation force.",
					LongDescription:  "Cherry MX Silent Red switches are designed to offer the smooth, linear feel of MX Red switches while significantly reducing noise during actuation. These switches incorporate built-in dampeners that minimize the sound of both the key press and release, making them an excellent choice for quiet environments like offices or shared spaces. With an actuation force of 45 grams, they maintain a light, responsive feel, ideal for fast typing or gaming without the auditory distraction.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2015-01-01"),
					Available:        false,
					PricePoint:       3,
					SiteURL:          "",
				},

				{
					Name:             "Cherry MX2A Silent Black",
					ShortDescription: "Silent linear switch with a heavy actuation force.",
					LongDescription:  "The Cherry MX2A Silent Black switch offers a quiet, smooth linear typing experience with a heavier actuation force of 60 grams. Designed to minimize noise, this switch features built-in dampeners that significantly reduce the sound of both key press and release, making it ideal for quiet environments such as offices or shared spaces. The MX2A Silent Black is perfect for users who prefer a firmer keystroke while maintaining a low-noise profile, ensuring both performance and discretion in their typing setup.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2024-03-01"),
					Available:        true,
					PricePoint:       3, // Expensive
					SiteURL:          "https://www.cherry-world.com/mx-silent-black",
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx3a-11na_11nn_image.jpg",
							AltText:     "Cherry MX2A Silent Black switch",
						},
						{
							LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx-silent-black.jpg",
							AltText:     "Cherry MX2A Silent Black switch opened",
						},
					},
				},
				{
					Name:             "Cherry MX2A Brown Tactile",
					ShortDescription: "Tactile switch with a gentle bump for feedback.",
					LongDescription:  "The Cherry MX2A Brown Tactile switch provides a subtle tactile bump, making it suitable for users who prefer feedback without the accompanying click sound. With an actuation force of 55 grams, this switch offers a comfortable balance between typing and gaming, providing just enough feedback to avoid bottoming out while maintaining a smooth and controlled keypress.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2024-03-01"),
					Available:        true,
					PricePoint:       2, // Average
					SiteURL:          "https://www.cherry-world.com/mx2a-brown",
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/9380e5bfba4d8931eee53843958a2c83/m/x/mx1a-g1na_g1nn_image.jpg",
							AltText:     "Cherry MX2A Brown switch",
						},
						{
							LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-g1nn_explo.jpg",
							AltText:     "Cherry MX2A Brown switch",
						},
					},
				},
				{
					Name:             "Cherry MX2A Red Linear",
					ShortDescription: "Smooth linear switch for a balanced typing feel.",
					LongDescription:  "The Cherry MX2A Red Linear switch offers a balanced and smooth typing experience, suitable for both gaming and typing due to its medium actuation force of 45 grams. This switch is designed for users who prefer a light, responsive keystroke with no tactile bump or audible click, allowing for fluid and uninterrupted keypresses.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2024-03-01"),
					Available:        true,
					PricePoint:       2, // Average
					SiteURL:          "https://www.cherry-world.com/mx2a-red",
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-l1nn_image.jpg",
							AltText:     "Cherry MX2A Red switch",
						},
						{
							LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-l1nn_explo.jpg",
							AltText:     "Cherry MX2A Red opened switch",
						},
						{
							LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-l1na_image.jpg",
							AltText:     "Cherry MX2A Red switch",
						},
						{
							LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-l1nn_explo.jpg",
							AltText:     "Cherry MX2A Red opened switch",
						},
					},
				},
				{
					Name:             "Cherry MX2A Speed Silver Linear",
					ShortDescription: "Fast linear switch for quick actuation.",
					LongDescription:  "The Cherry MX2A Speed Silver Linear switch is designed for rapid keystrokes, making it ideal for gaming and other high-speed typing scenarios. With a short 1.2mm actuation distance and 45 grams of force, this switch enables quick responses, helping to improve performance in fast-paced environments. The MX2A Speed Silver is a top choice for competitive gamers seeking the fastest actuation possible.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2024-03-01"),
					Available:        true,
					PricePoint:       3, // Expensive
					SiteURL:          "https://www.cherry-world.com/mx2a-speed-silver",
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-51nn_image.jpg",
							AltText:     "Cherry MX2A Speed Silver switch",
						},
						{
							LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-51nn_sideview4.jpg",
							AltText:     "Cherry MX2A Speed Silver switch",
						},
						{
							LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-51na_image.jpg",
							AltText:     "Cherry MX2A Speed Silver switch",
						},
						{
							LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-51na_sideview4.jpg",
							AltText:     "Cherry MX2A Speed Silver switch",
						},
					},
				},
				{
					Name:             "Cherry MX2A Black Linear",
					ShortDescription: "Linear switch with a heavy actuation force.",
					LongDescription:  "The Cherry MX2A Black Linear switch offers a smooth keystroke with a heavier actuation force of 60 grams. This makes it ideal for users who prefer a firm, controlled typing experience. The switch is well-suited for both heavy typists and gamers who require a more deliberate keypress. With its linear feel and no tactile bump, the MX2A Black ensures a consistent and smooth actuation with every keystroke.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2024-03-01"),
					Available:        true,
					PricePoint:       2, // Average
					SiteURL:          "https://www.cherry-world.com/mx2a-black",
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-11nn_image.jpg",
							AltText:     "Cherry MX2A Black switch",
						},
						{
							LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-11nn_explo.jpg",
							AltText:     "Cherry MX2A Black opened switch",
						},

						{
							LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-11na_image.jpg",
							AltText:     "Cherry MX2A Black switch",
						},
						{
							LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-11na_explo.jpg",
							AltText:     "Cherry MX2A Black opened switch",
						},
					},
				},
				{
					Name:             "Cherry MX2A Clear",
					ShortDescription: "Tactile switch with a high actuation force.",
					LongDescription:  "The Cherry MX2A Clear switch is designed for users who prefer a strong tactile feedback without the click sound. With an actuation force of 65 grams, this switch offers a pronounced tactile bump that is ideal for heavy typists who enjoy a firm and responsive keystroke. The MX2A Clear is particularly suited for those who need a tactile switch that provides clear feedback for every keypress, making it a great choice for both typing and precise input tasks.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2024-03-01"),
					Available:        true,
					PricePoint:       3,                                       // Expensive
					SiteURL:          "https://www.cherry-world.com/mx-clear", // If an official URL is available, you can update it here
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-c1na_c1nn_image.jpg",
							AltText:     "Cherry MX2A Clear switch",
						},
						{
							LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-c1nn_image.jpg",
							AltText:     "Cherry MX2A Clear switch",
						},
					},
				},
				{
					Name:             "Cherry MX2A Grey",
					ShortDescription: "Tactile switch with an even higher actuation force.",
					LongDescription:  "The Cherry MX2A Grey switch is designed for users who prefer a very strong tactile feedback with an actuation force of 80 grams. This switch offers a firm tactile bump, providing clear and pronounced feedback with every keystroke. The MX2A Grey is perfect for users who require a high level of resistance in their switches, making it ideal for precise and deliberate key presses in both typing and heavy-duty applications.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2024-03-01"),
					Available:        true,
					PricePoint:       3, // Expensive
					SiteURL:          "https://www.cherry-world.com/mx-grey-tactile",
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-d1na_d1nn_image.jpg",
							AltText:     "Cherry MX2A Grey switch",
						},
						{
							LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-d1nn_sideview4.jpg",
							AltText:     "Cherry MX2A Grey switch",
						},
						{
							LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-d1na_top",
							AltText:     "Cherry MX2A Grey switch",
						},
					},
				},
				{
					Name:             "Cherry MX Purple",
					ShortDescription: "Linear switch with a light actuation force and unique purple stem.",
					LongDescription:  "The Cherry MX Purple switch is a linear switch designed for users who prefer a smooth and consistent keystroke with a light actuation force of 45 grams. This switch features a unique purple stem, which adds a distinct aesthetic to any keyboard. The MX Purple is ideal for both gaming and typing, offering a balance between speed and precision while providing a visually appealing design. It is a popular choice for users looking to combine performance with style.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2024-03-01"),
					Available:        true,
					PricePoint:       3, // Expensive
					SiteURL:          "https://www.cherry-world.com/mx-purple",
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx2a-v1nw_imageview.jpg",
							AltText:     "Cherry MX Purple switch",
						},
						{
							LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx2a-v1nw_explo.jpg",
							AltText:     "Cherry MX Purple switch",
						},
					},
				},
				{
					Name:             "Cherry MX Clear-Top",
					ShortDescription: "Tactile switch with a transparent housing for enhanced RGB lighting.",
					LongDescription:  "The Cherry MX Clear-Top switch is a tactile switch that combines the renowned tactile feedback of the Cherry MX Clear with a transparent housing, making it perfect for keyboards with RGB lighting. With an actuation force of 65 grams, this switch offers a firm tactile bump that provides a satisfying typing experience. The transparent housing allows for vibrant RGB lighting, enhancing the visual appeal of your keyboard setup while maintaining the high performance and reliability of the MX Clear switch.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2024-03-01"),
					Available:        true,
					PricePoint:       3, // Expensive
					SiteURL:          "https://www.cherry-world.com/mx-black-clear-top",
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-61nw_image.jpg",
							AltText:     "Cherry MX Clear-Top switch",
						},
						{
							LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-61nw_explo.jpg",
							AltText:     "Cherry MX Clear-Top switch",
						},
					},
				},
				{
					Name:             "Cherry MX2A Blue Clicky",
					ShortDescription: "Clicky switch with an audible click for tactile feedback.",
					LongDescription:  "The Cherry MX2A Blue Clicky switch delivers the satisfying click sound and tactile feedback that typists love. With an actuation force of 50 grams, this switch provides a crisp click and a noticeable bump, ensuring each keystroke is both audible and precise. Ideal for those who enjoy the classic typing experience, the MX2A Blue is perfect for heavy typists or those who simply love the sound of their keyboard.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2024-03-01"),
					Available:        true,
					PricePoint:       2, // Average
					SiteURL:          "https://www.cherry-world.com/mx2a-blue",
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-e1na_e1nn_image.jpg",
							AltText:     "Cherry MX2A switch",
						},
						{
							LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-e1nn_explo.jpg",
							AltText:     "Cherry MX2A switch",
						},
					},
				},
				{
					Name:             "Cherry MX Blue",
					ShortDescription: "Clicky switch with a tactile bump.",
					LongDescription:  "Cherry MX Blue switches are iconic for their audible click and tactile feedback, offering a distinct typing experience that is both responsive and satisfying. With an actuation force of 50 grams, these switches provide a noticeable bump before actuation, coupled with a sharp click sound that appeals to typists who appreciate clear auditory and physical feedback. While they are less common in gaming due to their pronounced click, MX Blues are a top choice for those who value precision and clarity in their keystrokes.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("1987-01-01"),
					Available:        false,
					PricePoint:       2,
					SiteURL:          "https://www.cherry-world.com/mx-blue.html",
				},
				{
					Name:             "Cherry MX Brown",
					ShortDescription: "Tactile switch with a moderate actuation force.",
					LongDescription:  "Cherry MX Brown switches strike a balance between typing and gaming with their tactile bump and moderate actuation force of 55 grams. These switches are quieter than their Blue counterparts, lacking the audible click but still providing tactile feedback that is gentle yet noticeable. This makes them versatile and well-suited for both work and play, appealing to users who want a middle ground between the stark tactility of MX Blue switches and the smoothness of MX Red switches.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("1994-01-01"),
					Available:        false,
					PricePoint:       2,
					SiteURL:          "https://www.cherry-world.com/mx-brown.html",
				},
				{
					Name:             "Cherry MX Black",
					ShortDescription: "Linear switch with a heavier actuation force.",
					LongDescription:  "Cherry MX Black switches offer a smooth linear actuation similar to MX Red switches but with a heavier actuation force of 60 grams. This increased resistance makes them ideal for users who prefer a more deliberate keystroke, reducing the chance of accidental presses. MX Black switches are often favored by those who type heavily or for gaming environments where controlled inputs are critical. The absence of tactile bumps or clicks also ensures a quieter, more focused typing experience.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("1984-01-01"),
					Available:        false,
					PricePoint:       2,
					SiteURL:          "https://www.cherry-world.com/mx-black.html",
				},
				{
					Name:             "Cherry MX Green",
					ShortDescription: "Clicky switch with a heavy actuation force.",
					LongDescription:  "Cherry MX Green switches are a variant of the MX Blue, designed for those who enjoy a clicky experience but prefer a heavier actuation force. Requiring 80 grams to actuate, these switches provide a pronounced tactile bump and a louder click, making them ideal for users who want a more deliberate, assertive keystroke. MX Green switches are less common in standard keyboards but are highly prized by enthusiasts who enjoy a stiffer, more resistant typing experience.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2011-01-01"),
					Available:        true,
					PricePoint:       3,
					SiteURL:          "https://www.cherry-world.com/mx-green",
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-f1na_f1nn_image.jpg",
							AltText:     "Cherry MX Green switch",
						},
						{
							LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-f1nn_sideview3.jpg",
							AltText:     "Cherry MX Green switch",
						},

						{
							LinkAddress: "https://www.cherry-world.com/media/catalog/product/cache/661c0bae3bb54b88fbb1b415a9d390cb/m/x/mx1a-f1na_image.jpg",
							AltText:     "Cherry MX Green switch",
						},
					},
				},
				{
					Name:             "Cherry MX Clear",
					ShortDescription: "Tactile switch with a high actuation force.",
					LongDescription:  "Cherry MX Clear switches are designed for those who prefer a heavier tactile bump in their typing experience. With an actuation force of 65 grams, these switches provide a noticeable resistance, making each keystroke deliberate and satisfying. MX Clear switches are often chosen by users who want a tactile switch that offers more feedback than MX Browns but without the clicky noise associated with MX Blue switches.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("1994-01-01"),
					Available:        false,
					PricePoint:       3,
					SiteURL:          "https://www.cherry-world.com/mx-clear.html",
				},
				{
					Name:             "Cherry MX Grey",
					ShortDescription: "Tactile switch with a heavy actuation force.",
					LongDescription:  "Cherry MX Grey switches are an evolution of the MX Clear, featuring an even higher actuation force of 80 grams. This provides a pronounced tactile bump and substantial resistance, making them suitable for users who prefer a firm and assertive keystroke. These switches are less common in mass-market keyboards and are often found in specialized or enthusiast builds that prioritize a heavier, more tactile typing experience.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("1994-01-01"),
					Available:        false,
					PricePoint:       3,
					SiteURL:          "https://www.cherry-world.com/mx-grey.html",
				},
				{
					Name:             "Cherry MX Speed Silver",
					ShortDescription: "Linear switch with a shorter travel distance for faster actuation.",
					LongDescription:  "Cherry MX Speed Silver switches are optimized for rapid keystrokes with a shorter actuation distance and total travel. These linear switches require just 1.2mm to actuate and only 45 grams of force, making them perfect for competitive gaming where speed is crucial. The reduced travel distance allows for faster key presses, helping to shave milliseconds off each action in fast-paced environments. They are a go-to choice for gamers seeking performance and responsiveness.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2016-01-01"),
					Available:        false,
					PricePoint:       3,
					SiteURL:          "https://www.cherry-world.com/mx-speed-silver.html",
				},
				{
					Name:             "Cherry MX Nature White",
					ShortDescription: "Linear switch with a medium actuation force, between Red and Black.",
					LongDescription:  "Cherry MX Nature White switches are designed to offer a balanced typing experience, with an actuation force that sits comfortably between the lighter MX Red and the heavier MX Black switches. With a smooth, linear action and a medium force of 55 grams, these switches provide a well-rounded keystroke that is ideal for users who want a linear switch that isn’t too light or too heavy. They are often chosen by those who prefer a slightly more resistant keystroke without sacrificing speed or smoothness.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2016-01-01"),
					Available:        false,
					PricePoint:       2,
					SiteURL:          "https://www.cherry-world.com/mx-nature-white.html",
				},
				{
					Name:             "Cherry MX Violet",
					ShortDescription: "Tactile switch with a light actuation force, similar to Brown but lighter.",
					LongDescription:  "Cherry MX Violet switches offer a tactile bump similar to MX Browns but with a lighter actuation force, providing a gentle yet satisfying typing experience. Requiring only 45 grams of force to actuate, these switches are ideal for users who appreciate tactile feedback but prefer a softer, less resistant keypress. They are well-suited for both typing and gaming, offering a balance of feedback and ease of use.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        false,
					PricePoint:       2,
					SiteURL:          "",
				},
				{
					Name:             "Cherry MX White",
					ShortDescription: "A softer clicky switch, sometimes called 'Milk'.",
					LongDescription:  "Cherry MX White switches, often referred to as 'Milk' switches due to their smoother click, are a quieter alternative to the more common MX Blue switches. They retain the clicky feedback but with a softer, less pronounced sound, making them a great choice for users who enjoy a clicky switch but want something less intrusive in terms of noise. These switches are less common and have become a favorite among those who want a unique auditory and tactile experience.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2011-01-01"),
					Available:        false,
					PricePoint:       2,
					SiteURL:          "",
				},
				{
					Name:             "Cherry MX Tactile Grey",
					ShortDescription: "A variant of the Grey switch with a slightly different force curve.",
					LongDescription:  "Cherry MX Tactile Grey switches are a specialized variant of the standard MX Grey switches, offering a unique force curve that provides a slightly different tactile experience. These switches are designed for users who appreciate a heavy, pronounced tactile bump and prefer a switch with a firm and deliberate keystroke. They are often used in niche keyboard builds where a distinct tactile feel is desired.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("1994-01-01"),
					Available:        false,
					PricePoint:       3,
					SiteURL:          "",
				},
				{
					Name:             "Cherry MX Super Black",
					ShortDescription: "An extremely heavy linear switch.",
					LongDescription:  "Cherry MX Super Black switches are designed for users who demand an extremely heavy actuation force. With a force requirement of 150 grams, these linear switches are among the stiffest in the Cherry MX lineup, offering unparalleled resistance to accidental key presses. Super Black switches are typically used in specialized keyboards or applications where a very deliberate keypress is necessary, providing a unique typing experience that is both demanding and rewarding.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("1984-01-01"),
					Available:        false,
					PricePoint:       3,
					SiteURL:          "",
				},
				{
					Name:             "Cherry MX Red",
					ShortDescription: "Linear switch with a light actuation force.",
					LongDescription:  "Cherry MX Red switches are renowned for their smooth, linear actuation and light operating force, making them a preferred choice for gamers and typists who enjoy a fast, uninterrupted keystroke. These switches require only 45 grams of force to actuate, ensuring a quick response without any tactile bump or audible click, allowing for a fluid typing experience. They are especially favored in gaming for their consistent key presses and rapid input, reducing finger fatigue during extended use.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("1984-01-01"),
					Available:        false,
					PricePoint:       2,
					SiteURL:          "https://www.cherry-world.com/mx-red.html",
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

			err := processSwitches(cherry)
			if err != nil {
				return err
			}

			gateron := []Switch{
				{
					Name:             "Gateron G Pro Red",
					ShortDescription: "Linear switch with factory lubrication for smoother keystrokes.",
					LongDescription:  "Gateron G Pro Red switches are enhanced linear switches featuring factory lubrication for smoother keystrokes and improved performance. With an actuation force of 45 grams, these switches provide a light and responsive typing experience, making them ideal for fast typists and gamers. The G Pro Red switches are designed to offer minimal resistance and friction, ensuring a consistent and satisfying keystroke.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       2,                                                    // Average
					SiteURL:          "https://www.gateron.co/products/gateron-switch-set", // Update if an official URL becomes available
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-G-Pro-Red-Switch-Set_1000x.jpg?v=1657361719",
							AltText:     "Gateron G Pro Red switch",
						},
					},
				},
				{
					Name:             "Gateron G Pro Yellow",
					ShortDescription: "Linear switch with a medium actuation force and improved smoothness.",
					LongDescription:  "Gateron G Pro Yellow switches offer a medium actuation force of 50 grams and are factory lubricated to provide a smooth and responsive typing experience. These linear switches are perfect for users who want a balanced feel that works well for both gaming and typing. The G Pro Yellow switches are known for their reliability and smoothness, making them a popular choice among enthusiasts.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       2,                                                    // Average
					SiteURL:          "https://www.gateron.co/products/gateron-switch-set", // Update if an official URL becomes available
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-G-Pro-Yellow-Switch-Set_1000x.jpg?v=1657361719",
							AltText:     "Gateron G Pro Red switch",
						},
					},
				},
				{
					Name:             "Gateron G Pro Black",
					ShortDescription: "Linear switch with a heavier actuation force for more resistance.",
					LongDescription:  "Gateron G Pro Black switches are linear switches designed for users who prefer a heavier actuation force of 60 grams. These switches provide a firm and controlled typing experience, ideal for those who enjoy more resistance in their keystrokes. Factory lubrication ensures smooth operation, making the G Pro Black switches a reliable option for both gaming and long typing sessions.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       2,                                                    // Average
					SiteURL:          "https://www.gateron.co/products/gateron-switch-set", // Update if an official URL becomes available
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-G-Pro-Black-Switch-Set_1000x.jpg?v=1657361719",
							AltText:     "Gateron G Pro Black switch",
						},
					},
				},
				{
					Name:             "Gateron G Pro Brown",
					ShortDescription: "Tactile switch with factory lubrication and a gentle bump.",
					LongDescription:  "Gateron G Pro Brown switches feature a gentle tactile bump with factory lubrication, offering a smooth and quiet typing experience. With an actuation force of 55 grams, these tactile switches provide a satisfying feedback without the click sound, making them suitable for both gaming and typing. The G Pro Brown switches are designed to offer a balanced tactile feel, appealing to users who prefer moderate resistance in their typing.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       2,                                                    // Average
					SiteURL:          "https://www.gateron.co/products/gateron-switch-set", // Update if an official URL becomes available
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-G-Pro-Brown-Switch-Set_1000x.jpg?v=1657361719",
							AltText:     "Gateron G Pro Brown switch",
						},
					},
				},
				{
					Name:             "Gateron G Pro Green",
					ShortDescription: "Clicky switch with a strong tactile bump and audible click.",
					LongDescription:  "Gateron G Pro Green switches provide a strong tactile bump with an audible click, ideal for users who enjoy feedback and sound in their typing. With an actuation force of 80 grams, these clicky switches offer a pronounced and satisfying typing experience. Factory lubrication ensures smooth keystrokes, making the G Pro Green switches a reliable choice for enthusiasts who value both tactile feedback and auditory satisfaction.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       2,                                                    // Average
					SiteURL:          "https://www.gateron.co/products/gateron-switch-set", // Update if an official URL becomes available
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-G-Pro-Green-Switch-Set_1000x.jpg?v=1662948459",
							AltText:     "Gateron G Pro Green switch",
						},
					},
				},
				{
					Name:             "Gateron G Pro Blue",
					ShortDescription: "Clicky switch with an audible click and tactile feedback.",
					LongDescription:  "Gateron G Pro Blue switches are clicky switches designed for users who appreciate tactile feedback paired with an audible click. With an actuation force of 60 grams, these switches offer a satisfying typing experience with a distinct click sound. The G Pro Blue switches are ideal for users who prefer a more traditional mechanical keyboard feel with pronounced feedback and sound.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       2,                                                    // Average
					SiteURL:          "https://www.gateron.co/products/gateron-switch-set", // Update if an official URL becomes available
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-G-Pro-Blue-Switch-Set_1000x.jpg?v=1662948459",
							AltText:     "Gateron G Pro Blue switch",
						},
					},
				},
				{
					Name:             "Gateron Pro 2.0 Red",
					ShortDescription: "Linear switch with upgraded smoothness and factory lubrication.",
					LongDescription:  "Gateron Pro 2.0 Red switches are enhanced linear switches designed for a smooth typing experience. With an actuation force of 45 grams, these switches are factory lubricated to provide a more refined and responsive keystroke. The Pro 2.0 Red switches offer a light and effortless typing experience, making them ideal for fast-paced gaming and everyday typing tasks.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       2,                                                              // Average
					SiteURL:          "https://www.gateron.co/products/gateron-g-pro-2-0-switch-set", // Update if an official URL becomes available
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-G-Pro-2.0-Red-Switch-Set_1000x.jpg?v=1657505037",
							AltText:     "Gateron G Pro 2.0 Blue switch",
						},
					},
				},
				{
					Name:             "Gateron Pro 2.0 Yellow",
					ShortDescription: "Linear switch with a medium actuation force and improved smoothness.",
					LongDescription:  "Gateron Pro 2.0 Yellow switches feature an upgraded linear design with a medium actuation force of 50 grams. These switches are factory lubricated, offering a smooth and responsive typing experience. The Pro 2.0 Yellow switches are perfect for users who prefer a balanced feel that works well for both gaming and typing.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       2,                                                              // Average
					SiteURL:          "https://www.gateron.co/products/gateron-g-pro-2-0-switch-set", // Update if an official URL becomes available
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-G-Pro-2.0-Yellow-Switch-Set_1000x.jpg?v=1657505119",
							AltText:     "Gateron G Pro 2.0 Blue switch",
						},
					},
				},
				{
					Name:             "Gateron Pro 2.0 Black",
					ShortDescription: "Linear switch with a heavier actuation force and factory lubrication.",
					LongDescription:  "Gateron Pro 2.0 Black switches are designed for users who prefer a heavier actuation force of 60 grams. These switches provide a firm and controlled typing experience, with factory lubrication ensuring smooth operation. The Pro 2.0 Black switches are ideal for those who enjoy more resistance in their keystrokes, offering a reliable option for both gaming and long typing sessions.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       2,                                                              // Average
					SiteURL:          "https://www.gateron.co/products/gateron-g-pro-2-0-switch-set", // Update if an official URL becomes available
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-G-Pro-2.0-Black-Switch-Set_1000x.jpg?v=1657505119",
							AltText:     "Gateron G Pro 2.0 Blue switch",
						},
					},
				},
				{
					Name:             "Gateron Pro 2.0 Brown",
					ShortDescription: "Tactile switch with a gentle bump and improved stability.",
					LongDescription:  "Gateron Pro 2.0 Brown switches offer a tactile typing experience with a gentle bump and an actuation force of 55 grams. These switches are factory lubricated, providing a smooth and quiet typing experience. The Pro 2.0 Brown switches are designed to offer balanced tactile feedback, making them suitable for both gaming and typing.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       2,                                                              // Average
					SiteURL:          "https://www.gateron.co/products/gateron-g-pro-2-0-switch-set", // Update if an official URL becomes available
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-G-Pro-2.0-Brown-Switch-Set_1000x.jpg?v=1657505119",
							AltText:     "Gateron G Pro 2.0 Blue switch",
						},
					},
				},
				{
					Name:             "Gateron Pro 2.0 Blue",
					ShortDescription: "Clicky switch with an audible click and refined feedback.",
					LongDescription:  "Gateron Pro 2.0 Blue switches are clicky switches designed for users who appreciate tactile feedback paired with an audible click. With an actuation force of 60 grams, these switches offer a satisfying typing experience with a distinct click sound. The Pro 2.0 Blue switches are ideal for users who prefer a more traditional mechanical keyboard feel with pronounced feedback and sound.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       2,                                                              // Average
					SiteURL:          "https://www.gateron.co/products/gateron-g-pro-2-0-switch-set", // Update if an official URL becomes available
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-G-Pro-2.0-Blue-Switch-Set_1000x.jpg?v=1657505119",
							AltText:     "Gateron G Pro 2.0 Blue switch",
						},
					},
				},
				{
					Name:             "Gateron Pro 3.0 Red",
					ShortDescription: "Linear switch with enhanced smoothness and stability.",
					LongDescription:  "Gateron Pro 3.0 Red switches represent the latest in Gateron's linear switch technology, featuring enhanced smoothness and stability. With an actuation force of 45 grams, these switches are designed for users who demand a light and effortless typing experience. The Pro 3.0 Red switches come factory lubricated, ensuring consistent and smooth keystrokes with minimal friction, making them ideal for both gaming and typing.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2022-01-01"),
					Available:        true,
					PricePoint:       3,                                                              // Expensive
					SiteURL:          "https://www.gateron.co/products/gateron-g-pro-3-0-switch-set", // Update if an official URL becomes available
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.gateron.co/cdn/shop/files/Gateron-G-Pro-3.0-Red-Switch-Set_1000x.jpg?v=1683342544",
							AltText:     "Gateron G Pro 3.0 Red switch",
						},
					},
				},
				{
					Name:             "Gateron Pro 3.0 Yellow",
					ShortDescription: "Linear switch with medium actuation force and advanced performance.",
					LongDescription:  "Gateron Pro 3.0 Yellow switches are linear switches designed for users seeking a medium actuation force of 50 grams. These switches offer advanced performance with factory lubrication for ultra-smooth keystrokes. The Pro 3.0 Yellow switches are ideal for users who need a balanced typing experience that is suitable for both gaming and long typing sessions.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2022-01-01"),
					Available:        true,
					PricePoint:       3,                                                              // Expensive
					SiteURL:          "https://www.gateron.co/products/gateron-g-pro-3-0-switch-set", // Update if an official URL becomes available
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.gateron.co/cdn/shop/files/Gateron-G-Pro-3.0-Yellow-Switch-Set_1000x.jpg?v=1702451537",
							AltText:     "Gateron G Pro 3.0 Yellow  switch",
						},
					},
				},
				{
					Name:             "Gateron Pro 3.0 Black",
					ShortDescription: "Linear switch with a heavier actuation force and enhanced durability.",
					LongDescription:  "Gateron Pro 3.0 Black switches are the pinnacle of Gateron's linear switch lineup, offering a heavier actuation force of 60 grams. These switches provide a firm and controlled typing experience, making them ideal for users who prefer more resistance in their keystrokes. Factory lubrication and enhanced durability make the Pro 3.0 Black switches a reliable choice for both gaming and extended typing sessions.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2022-01-01"),
					Available:        true,
					PricePoint:       3,                                                              // Expensive
					SiteURL:          "https://www.gateron.co/products/gateron-g-pro-3-0-switch-set", // Update if an official URL becomes available
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.gateron.co/cdn/shop/files/Gateron-G-Pro-3.0-Black-Switch-Set_1000x.jpg?v=1702451537",
							AltText:     "Gateron G Pro 3.0 Black switch",
						},
					},
				},
				{
					Name:             "Gateron Pro 3.0 Brown",
					ShortDescription: "Tactile switch with a refined bump and improved smoothness.",
					LongDescription:  "Gateron Pro 3.0 Brown switches offer a refined tactile typing experience with a gentle bump and an actuation force of 55 grams. These switches are factory lubricated to provide a smooth and quiet operation, making them suitable for both gaming and typing. The Pro 3.0 Brown switches are designed to offer a balanced tactile feel with improved stability and consistency.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2022-01-01"),
					Available:        true,
					PricePoint:       3,                                                              // Expensive
					SiteURL:          "https://www.gateron.co/products/gateron-g-pro-3-0-switch-set", // Update if an official URL becomes available
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.gateron.co/cdn/shop/files/Gateron-G-Pro-3.0-Brown-Switch-Set_1000x.jpg?v=1702451537",
							AltText:     "Gateron G Pro 3.0 Brown switch",
						},
					},
				},
				{
					Name:             "Gateron Pro 3.0 Blue",
					ShortDescription: "Clicky switch with an audible click and premium feedback.",
					LongDescription:  "Gateron Pro 3.0 Blue switches are clicky switches designed for users who appreciate tactile feedback paired with an audible click. With an actuation force of 60 grams, these switches offer a satisfying typing experience with a distinct click sound. The Pro 3.0 Blue switches are ideal for users who prefer a more traditional mechanical keyboard feel with pronounced feedback and sound.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2022-01-01"),
					Available:        true,
					PricePoint:       3,                                                              // Expensive
					SiteURL:          "https://www.gateron.co/products/gateron-g-pro-3-0-switch-set", // Update if an official URL becomes available
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.gateron.co/cdn/shop/files/Gateron-G-Pro-3.0-Blue-Switch-Set_1000x.jpg?v=1702451537",
							AltText:     "Gateron G Pro 3.0 Blue switch",
						},
					},
				},
				{
					Name:             "Gateron Ink V2 Red",
					ShortDescription: "Linear switch with improved smoothness and unique housing.",
					LongDescription:  "Gateron Ink V2 Red switches are designed for users who seek a smooth and responsive linear typing experience. With an actuation force of 45 grams, these switches offer light and effortless keystrokes, perfect for fast typists and gamers. The Ink V2 Red switches feature Gateron's distinctive Ink housing, which provides enhanced stability and durability, ensuring consistent performance over time.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
					SiteURL:          "https://www.gateron.co/products/gateron-ink-switch",
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-Ink-V2-Red-Switch-Set_1000x.jpg?v=1657360038",
							AltText:     "Gateron Ink V2 Red switch",
						},
					},
				},
				{
					Name:             "Gateron Ink V2 Black",
					ShortDescription: "Linear switch with a heavier actuation force and improved design.",
					LongDescription:  "Gateron Ink V2 Black switches are premium linear switches known for their smooth actuation and distinct housing design. With a heavier actuation force of 60 grams, these switches provide a firm and controlled typing experience, making them ideal for users who prefer a stronger keystroke. The unique Ink housing not only enhances the switch's aesthetics but also contributes to its exceptional durability and stability.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
					SiteURL:          "https://www.gateron.co/products/gateron-ink-switch",
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-Ink-V2-Black-Switch-Set_1000x.jpg?v=1657360038",
							AltText:     "Gateron Ink V2 Black switch",
						},
					},
				},
				{
					Name:             "Gateron Ink V2 Yellow",
					ShortDescription: "Linear switch with mid-range actuation and unique housing.",
					LongDescription:  "Gateron Ink V2 Yellow switches provide a smooth linear feel with a mid-range actuation force of 50 grams. These switches are part of Gateron's premium Ink series, known for their exceptional build quality and unique housing design. The Ink V2 Yellow switches are ideal for users who want a responsive and consistent typing experience with just the right amount of resistance for both gaming and typing.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
					SiteURL:          "https://www.gateron.co/products/gateron-ink-switch",
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-Ink-V2-Yellow-Switch-Set_1000x.jpg?v=1657360038",
							AltText:     "Gateron Ink V2 Yellow switch",
						},
					},
				},
				{
					Name:             "Gateron Ink V2 Blue",
					ShortDescription: "Clicky switch with unique housing for enhanced sound and feel.",
					LongDescription:  "Gateron Ink V2 Blue switches are clicky switches designed for users who appreciate tactile feedback paired with an audible click. With an actuation force of 55 grams, these switches offer a satisfying typing experience with a distinct click sound. The unique Ink housing adds to the switch's durability and enhances the overall feel, making the Ink V2 Blue a popular choice among enthusiasts who enjoy clicky switches with a premium build.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
					SiteURL:          "https://www.gateron.co/products/gateron-ink-switch",
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-Ink-V2-Blue-Switch-Set_1000x.jpg?v=1657360038",
							AltText:     "Gateron Ink V2 Blue switch",
						},
					},
				},

				{
					Name:             "Gateron Oil King",
					ShortDescription: "Linear switch with smooth actuation and unique black housing.",
					LongDescription:  "Gateron Oil King switches are linear switches celebrated for their ultra-smooth keystrokes and unique black housing. With an actuation force of 55 grams, these switches are designed to offer a premium typing experience with minimal friction, making them ideal for both gaming and long typing sessions. The Oil King switches are part of Gateron's top-tier lineup, providing a balance of smoothness and durability that appeals to enthusiasts and professionals alike.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
					SiteURL:          "https://www.gateron.co/products/gateron-oil-king-switch-set",
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-Oil-King-Linear-Switch-Set_1000x.jpg?v=1657359635",
							AltText:     "Gateron Oil King switch",
						},
					},
				},

				{
					Name:             "Gateron Baby Kangaroo",
					ShortDescription: "Linear switch with a unique spring design.",
					LongDescription:  "Gateron Baby Kangaroo switches offer a distinctive actuation feel, featuring a unique spring design for a novel linear experience.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       2, // Average
					SiteURL:          "https://www.gateron.co/products/gateron-baby-kangaroo-tactile-switch-set",
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.gateron.co/cdn/shop/files/Gateron-Baby-Kangaroo-2.0-Tactile-Switch-Set_1000x.jpg?v=1698722119",
							AltText:     "Gateron Baby Kangaroo switch",
						},
					},
				},
				{
					Name:             "Gateron Cap V2 Brown",
					ShortDescription: "Upgraded tactile switch with a moderate actuation force.",
					LongDescription:  "Gateron Cap V2 Brown switches feature an upgraded tactile experience with enhanced sound dampening, offering a more refined tactile feel. With a moderate actuation force of 55 grams, these switches provide a satisfying bump with reduced noise, making them ideal for users who appreciate tactile feedback without the audible distraction. The Cap V2 series is designed for users seeking both performance and a quieter typing experience.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
					SiteURL:          "https://www.gateron.co/products/gateron-cap-switch-set",
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-Cap-V2-Switch-Set-Golden-Brown_1000x.jpg?v=1659520097",
							AltText:     "Gateron Cap V2 Brown switch",
						},
					},
				},
				{
					Name:             "Gateron Cap V2 Blue",
					ShortDescription: "Upgraded clicky switch with a tactile bump.",
					LongDescription:  "Gateron Cap V2 Blue switches feature an upgraded clicky experience with enhanced sound and tactile feedback, providing a satisfying typing feel.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
					SiteURL:          "https://www.gateron.co/products/gateron-cap-switch-set",
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-Cap-V2-Switch-Set-Blue_1000x.jpg?v=1659520097",
							AltText:     "Gateron Cap V2 Blue switch",
						},
					},
				},
				{
					Name:             "Gateron Cap V2 Yellow",
					ShortDescription: "Upgraded linear switch with a smooth feel and factory lubrication.",
					LongDescription:  "Gateron Cap V2 Yellow switches provide an enhanced linear typing experience with factory lubrication for smoother keystrokes.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
					SiteURL:          "https://www.gateron.co/products/gateron-cap-switch-set",
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://www.gateron.co/cdn/shop/products/Gateron-Cap-V2-Switch-Set-Golden-Yellow_1000x.jpg?v=1657360751",
							AltText:     "Gateron Cap V2 Yellow switch",
						},
					},
				},
				// Marking where we left off
				{
					Name:             "Gateron Red",
					ShortDescription: "Linear switch with a light actuation force.",
					LongDescription:  "Gateron Red switches are favored for their smooth linear action and light actuation force of 45 grams, making them ideal for gaming and typing. These switches provide a consistent and uninterrupted keystroke, offering a responsive and fast typing experience. They are a popular choice among users who prefer a softer and quieter switch without tactile feedback or audible click, providing an overall smooth and fluid typing experience.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2014-01-01"),
					Available:        true,
					PricePoint:       1, // Value
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "",
							AltText:     "Gateron switch",
						},
						{
							LinkAddress: "",
							AltText:     "Gateron switch",
						},
					},
				},
				{
					Name:             "Gateron Blue",
					ShortDescription: "Clicky switch with a tactile bump.",
					LongDescription:  "Gateron Blue switches are known for their clicky and tactile typing experience, offering an audible click sound and a distinct tactile bump with each key press. With an actuation force of 55 grams, these switches are popular among typists who appreciate auditory feedback and a satisfying tactile response. The crisp click sound makes them less ideal for quiet environments but highly favored for their satisfying and precise keystrokes.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2014-01-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				{
					Name:             "Gateron Brown",
					ShortDescription: "Tactile switch with a moderate actuation force.",
					LongDescription:  "Gateron Brown switches provide a balanced typing experience with a tactile bump and moderate actuation force of 55 grams. These switches are ideal for users who want tactile feedback without the loud click associated with Blue switches. Gateron Browns are versatile and well-suited for both typing and gaming, offering a satisfying yet quiet keystroke, making them a favorite among users who seek a middle ground between linear and clicky switches.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2014-01-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				{
					Name:             "Gateron Black",
					ShortDescription: "Linear switch with a heavier actuation force.",
					LongDescription:  "Gateron Black switches offer a smooth linear action with a heavier actuation force of 60 grams, providing a firmer typing experience. These switches are ideal for users who prefer more resistance in their keystrokes, reducing the likelihood of accidental key presses. The Gateron Black switch is often chosen by gamers and heavy typists who want a consistent and controlled key press, making it a popular choice for those who appreciate a stiffer feel without tactile feedback.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2014-01-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				{
					Name:             "Gateron Yellow",
					ShortDescription: "Linear switch with a medium actuation force.",
					LongDescription:  "Gateron Yellow switches are known for their smooth linear feel and medium actuation force of 50 grams, offering a balanced experience suitable for both gaming and typing. These switches provide a consistent keystroke with just the right amount of resistance, making them a favorite among users who want a switch that is not too light or too heavy. The Gateron Yellow switch is praised for its responsiveness and smoothness, making it a versatile choice for various applications.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2014-01-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				{
					Name:             "Gateron Green",
					ShortDescription: "Clicky switch with a heavy actuation force.",
					LongDescription:  "Gateron Green switches are a heavier variant of the Gateron Blue, offering a more pronounced tactile bump and a louder click with each key press. With an actuation force of 80 grams, these switches provide a firm and satisfying clicky typing experience, ideal for users who enjoy strong tactile feedback and audible keystrokes. Gateron Greens are often chosen by enthusiasts who want a more deliberate and forceful typing experience.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2014-01-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				{
					Name:             "Gateron Clear",
					ShortDescription: "Linear switch with an extremely light actuation force.",
					LongDescription:  "Gateron Clear switches offer an ultra-light linear feel with an actuation force of only 35 grams, making them ideal for users who prefer minimal resistance during keystrokes. These switches are extremely smooth and require very little force to actuate, providing a quick and effortless typing experience. They are particularly well-suited for users who suffer from finger fatigue or those who prefer a feather-light typing touch.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2014-01-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				{
					Name:             "Gateron Silent Red",
					ShortDescription: "Silent linear switch with a light actuation force.",
					LongDescription:  "Gateron Silent Red switches are designed to offer the smooth, linear feel of Gateron Reds while significantly reducing noise during actuation. With an actuation force of 45 grams, these switches incorporate dampeners that minimize the sound of both key press and release, making them ideal for quiet environments like offices or shared spaces. Gateron Silent Reds provide a responsive and fluid typing experience without the auditory distraction.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2018-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron Silent Black",
					ShortDescription: "Silent linear switch with a heavier actuation force.",
					LongDescription:  "Gateron Silent Black switches offer a quiet and smooth linear typing experience with a heavier actuation force of 60 grams. These switches are ideal for users who prefer a more deliberate keystroke while maintaining a low-noise profile, thanks to the built-in dampeners that reduce the sound of both key press and release. Gateron Silent Blacks are perfect for those who require both performance and silence in their typing environments.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2018-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron Pro Yellow",
					ShortDescription: "Enhanced linear switch with a medium actuation force.",
					LongDescription:  "Gateron Pro Yellow switches are an upgraded version of the standard Gateron Yellows, featuring factory lubrication for an even smoother keystroke. With a medium actuation force of 50 grams, these linear switches offer a balance between lightness and resistance, making them highly versatile for both gaming and typing. The factory lubrication reduces friction, providing a more refined and satisfying typing experience.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				// Discontinued Gateron Switches
				{
					Name:             "Gateron Blue Bubblegum",
					ShortDescription: "Tactile switch with a unique actuation force.",
					LongDescription:  "Gateron Blue Bubblegum switches offered a unique tactile experience with a distinct color scheme, known for their smooth keystrokes.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2017-01-01"),
					Available:        false,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron Aliaz Silent",
					ShortDescription: "Silent tactile switch with a gentle tactile bump.",
					LongDescription:  "Gateron Aliaz Silent switches provided a quiet tactile experience with a smooth bump, popular among those who preferred stealthy typing.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(5), // ZealPC
					SwitchTypeID:     5,      // Silent Tactile
					ReleaseDate:      parseDate("2017-01-01"),
					Available:        false,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Gateron White",
					ShortDescription: "Clicky switch with a light actuation force.",
					LongDescription:  "Gateron White switches provide a tactile typing experience with a lighter actuation force compared to Gateron Blue, offering a crisp and responsive feel.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2015-01-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				{
					Name:             "Gateron Green Pro",
					ShortDescription: "Clicky switch with a high actuation force.",
					LongDescription:  "Gateron Green Pro switches offer a clicky feel with a higher actuation force, designed for users who prefer a more pronounced tactile feedback and click sound.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron Optical Blue",
					ShortDescription: "Clicky optical switch with a tactile bump.",
					LongDescription:  "Gateron Optical Blue switches use optical technology to provide a fast response with a clicky and tactile feel, perfect for gaming and fast typing.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2019-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron Optical Green",
					ShortDescription: "Clicky optical switch with a high actuation force.",
					LongDescription:  "Gateron Optical Green switches provide a clicky feel with a higher actuation force, utilizing optical technology for fast and accurate keystrokes.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2019-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron Green Yellow",
					ShortDescription: "Clicky switch with a unique color scheme.",
					LongDescription:  "Gateron Green Yellow switches offer a clicky typing experience with a unique green and yellow color scheme, providing both aesthetics and performance.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2017-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				// Discontinued Gateron Clicky Switches
				{
					Name:             "Gateron Pink",
					ShortDescription: "Clicky switch with a distinctive color.",
					LongDescription:  "Gateron Pink switches were known for their clicky feel and distinctive pink color, providing a unique typing experience.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2016-01-01"),
					Available:        false,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron Green Keydous",
					ShortDescription: "Custom clicky switch with enhanced features.",
					LongDescription:  "Gateron Green Keydous switches offered a customized clicky experience with enhanced features, providing a unique typing feel.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2018-01-01"),
					Available:        false,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Gateron Clicky Clear",
					ShortDescription: "Light clicky switch with a clear housing.",
					LongDescription:  "Gateron Clicky Clear switches offered a light clicky feel with a clear housing, making them both visually appealing and functional.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2016-01-01"),
					Available:        false,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron Orange",
					ShortDescription: "Tactile switch with a moderate actuation force and unique orange color.",
					LongDescription:  "Gateron Orange switches offer a noticeable tactile bump with a moderate actuation force, ideal for those seeking a distinct tactile feel with a pop of color.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2018-01-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				{
					Name:             "Gateron Brown Pro",
					ShortDescription: "Enhanced tactile switch with factory lubrication.",
					LongDescription:  "Gateron Brown Pro switches provide a smoother tactile experience compared to standard Gateron Browns, featuring factory lubrication for improved performance.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron Purple",
					ShortDescription: "Tactile switch with a light tactile bump and unique purple housing.",
					LongDescription:  "Gateron Purple switches offer a light tactile bump with a distinct purple housing, providing a unique typing feel for those seeking a balance between tactile feedback and aesthetics.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2019-01-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				{
					Name:             "Gateron Zealios V2",
					ShortDescription: "Premium tactile switch with a refined tactile bump.",
					LongDescription:  "A collaboration with ZealPC, Gateron Zealios V2 switches offer a highly refined tactile bump with premium components for a superior typing experience.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(5), // ZealPC
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2018-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Gateron Kangaroo",
					ShortDescription: "Tactile switch with a strong tactile bump.",
					LongDescription:  "Gateron Kangaroo switches are known for their unique actuation profile and distinctive kangaroo-themed color, providing a strong tactile bump and responsive feedback.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron Yellow Tactile",
					ShortDescription: "Discontinued tactile switch with a unique yellow color.",
					LongDescription:  "Gateron Yellow Tactile switches offered a distinct tactile bump with a unique yellow color, but are no longer in production.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2015-01-01"),
					Available:        false,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron Sky",
					ShortDescription: "Discontinued tactile switch developed with Akko.",
					LongDescription:  "A tactile switch developed in collaboration with Akko, known for its sky-blue color and refined tactile bump, but now discontinued.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2017-01-01"),
					Available:        false,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron Silver",
					ShortDescription: "Linear switch with a shorter travel distance for faster actuation.",
					LongDescription:  "Gateron Silver switches offer a linear feel with a reduced travel distance, making them ideal for rapid keystrokes and gaming.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2019-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron Pro Red",
					ShortDescription: "Enhanced linear switch with factory lubrication.",
					LongDescription:  "Gateron Pro Red switches offer a smoother keystroke compared to standard Gateron Red, thanks to factory lubrication and improved components.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron Raccoon",
					ShortDescription: "Linear switch with a smooth actuation and unique design.",
					LongDescription:  "Gateron Raccoon switches are designed for a smooth linear typing experience, featuring a unique spring design and distinctive raccoon-themed aesthetics.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron Green (Linear)",
					ShortDescription: "Linear switch with a medium actuation force.",
					LongDescription:  "Gateron Green Linear switches provide a balanced feel with a medium actuation force, suitable for both gaming and typing.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2019-01-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				{
					Name:             "Gateron Milky Yellow",
					ShortDescription: "Budget-friendly linear switch with a smooth feel.",
					LongDescription:  "Gateron Cap Milky Yellow switches are known for their smooth linear feel and cost-effectiveness, featuring a milky housing that provides a distinct look and softer bottom-out feel. With an actuation force of 50 grams, these switches offer a balanced typing experience, making them ideal for both typing and gaming. The milky housing also contributes to a more muted sound profile, appealing to users who prefer a quieter switch.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2018-01-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				{
					Name:             "Gateron Crystal Pink",
					ShortDescription: "Linear switch with a crystal-clear housing and pink stem.",
					LongDescription:  "Gateron Crystal Pink switches provide a smooth linear typing experience with a unique crystal-clear housing and pink stem for an aesthetic touch.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				// Discontinued Gateron Linear Switches
				{
					Name:             "Gateron Milk",
					ShortDescription: "Discontinued smooth linear switch with milky housing.",
					LongDescription:  "Gateron Milk switches were known for their smooth linear feel and unique milky white housing but are no longer in production.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2016-01-01"),
					Available:        false,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron X",
					ShortDescription: "Discontinued experimental linear switch.",
					LongDescription:  "Gateron X switches were available for a limited time, offering a unique linear actuation profile, but are now discontinued.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2017-01-01"),
					Available:        false,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Gateron Ink Yellow",
					ShortDescription: "Linear switch with smooth actuation and mid-range force.",
					LongDescription:  "Gateron Box Ink Yellow switches combine the smooth linear feel of the Ink series with the added stability and durability of a box-style design. With a mid-range actuation force of 50 grams, these switches are perfect for users who want a responsive and consistent typing experience with added protection against dust and moisture. The Box Ink Yellow switches are ideal for both gaming and typing, offering the best of both worlds in performance and durability.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Gateron Silent Yellow",
					ShortDescription: "Silent linear switch with a smooth feel and reduced noise.",
					LongDescription:  "Gateron Silent Yellow switches provide a smooth linear typing experience with reduced noise, making them ideal for quiet environments and office settings.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2018-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron Silent Brown",
					ShortDescription: "Silent tactile switch with a gentle bump and quiet operation.",
					LongDescription:  "Gateron Silent Brown switches offer a tactile experience with a gentle bump and quiet operation, making them ideal for users who prefer stealthy typing without sacrificing feedback.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     5,      // Silent Tactile
					ReleaseDate:      parseDate("2018-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron EF (Linear)",
					ShortDescription: "High-performance linear switch with smooth actuation and reduced wobble.",
					LongDescription:  "Gateron EF (Linear) switches are designed for high performance, featuring smooth actuation and reduced wobble for enhanced stability and precision.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Gateron EF (Tactile)",
					ShortDescription: "High-performance tactile switch with a defined bump and enhanced stability.",
					LongDescription:  "Gateron EF (Tactile) switches offer a defined tactile bump with enhanced stability, making them ideal for users seeking a premium tactile experience.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Gateron Baby Kangaroo V2",
					ShortDescription: "Enhanced tactile switch with a strong bump and smoother operation.",
					LongDescription:  "The Gateron Kangaroo V2 switches are tactile switches known for their strong tactile bump and smooth operation. With an actuation force of 55 grams, these switches are perfect for users who enjoy a pronounced tactile feedback that is both responsive and satisfying. The Kangaroo V2 switches are designed for those who prefer a more noticeable bump during typing, making them ideal for heavy typists and those who value distinct feedback in every keystroke.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2022-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron Cap Red",
					ShortDescription: "Linear switch with enhanced performance and sound dampening.",
					LongDescription:  "Gateron Cap Red switches are linear switches designed to offer smooth keystrokes with a light actuation force of 45 grams. These switches come factory-lubricated, ensuring a quiet and responsive typing experience. The Cap Red switches are part of Gateron's premium line, providing users with a reliable and high-performance switch that is perfect for gaming and fast typing.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Gateron CJ (Blue)",
					ShortDescription: "Linear switch with smooth actuation and unique blue housing.",
					LongDescription:  "Gateron CJ (Blue) switches provide a smooth linear typing experience with unique blue housing, offering a premium feel for enthusiasts.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Gateron CJ (Brown)",
					ShortDescription: "Tactile switch with a gentle bump and unique brown housing.",
					LongDescription:  "Gateron CJ (Brown) switches offer a distinct tactile experience with a gentle bump and unique brown housing, providing a premium feel for tactile enthusiasts.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Gateron Ink Red",
					ShortDescription: "Premium linear switch with a smooth keystroke.",
					LongDescription:  "The Gateron Box Ink Red switches are linear switches with a unique box-style design that enhances stability and offers protection against dust and moisture. With an actuation force of 45 grams, these switches provide a smooth and responsive keystroke, making them ideal for users who prioritize durability and performance. The Box Ink Red switches are part of Gateron's premium Ink series, known for their high-quality construction and reliable operation.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2019-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Gateron Ink Black",
					ShortDescription: "Premium linear switch with a heavier actuation force.",
					LongDescription:  "Gateron Box Ink Black switches are designed for users who prefer a heavier linear feel with the added benefits of a box-style design. With an actuation force of 60 grams, these switches provide a firm and consistent keystroke while offering enhanced protection against dust and moisture. The Box Ink Black switches are a great choice for those seeking a durable switch that can withstand heavy use while maintaining a smooth typing experience.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2019-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Gateron Ink Blue",
					ShortDescription: "Clicky switch with a distinct click and tactile feedback.",
					LongDescription:  "Gateron Ink Blue switches provide a clicky typing experience with a distinct click and tactile feedback, offering premium quality in the Ink series.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Gateron Box Ink Red",
					ShortDescription: "Box-style linear switch with improved dust and moisture protection.",
					LongDescription:  "Gateron Box Ink Red switches offer a linear typing experience with enhanced dust and moisture protection, providing durability and smooth keystrokes.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Gateron Box Ink Black",
					ShortDescription: "Box-style linear switch with a heavier actuation force.",
					LongDescription:  "Gateron Box Ink Black switches provide a heavier actuation force with box-style protection, offering smooth keystrokes and enhanced durability.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Gateron Box Ink Yellow",
					ShortDescription: "Box-style linear switch with smooth actuation and mid-range force.",
					LongDescription:  "Gateron Box Ink Yellow switches offer smooth actuation with mid-range force, featuring box-style protection for enhanced durability and performance.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Gateron KS-3 Red",
					ShortDescription: "Linear switch with a smooth feel and light actuation force.",
					LongDescription:  "Gateron KS-3 Red switches offer a smooth linear feel with light actuation, providing an affordable and reliable option for linear enthusiasts.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2015-01-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				{
					Name:             "Gateron KS-3 Brown",
					ShortDescription: "Tactile switch with a moderate actuation force and gentle bump.",
					LongDescription:  "Gateron KS-3 Brown switches provide a moderate actuation force with a gentle tactile bump, offering a balanced and affordable option for tactile enthusiasts.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2015-01-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				{
					Name:             "Gateron KS-3 Blue",
					ShortDescription: "Clicky switch with an audible click and tactile feedback.",
					LongDescription:  "Gateron KS-3 Blue switches offer an audible click and tactile feedback, providing a clicky typing experience at an affordable price point.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2015-01-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				{
					Name:             "Gateron KS-3 Yellow",
					ShortDescription: "Linear switch with a medium actuation force and smooth keystrokes.",
					LongDescription:  "Gateron KS-3 Yellow switches offer smooth keystrokes with a medium actuation force, providing a versatile and cost-effective option for linear enthusiasts.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2015-01-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				{
					Name:             "Gateron KS-3 Black",
					ShortDescription: "Linear switch with a heavier actuation force for added resistance.",
					LongDescription:  "Gateron KS-3 Black switches offer a heavier actuation force with added resistance, providing a firm linear typing experience at a value price.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2015-01-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				// More to include
				{
					Name:             "Gateron KS-8 Red",
					ShortDescription: "Linear switch with a light actuation force and smooth feel.",
					LongDescription:  "Gateron KS-8 Red switches offer a light actuation force with smooth keystrokes, providing an enjoyable typing experience.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2018-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron KS-8 Brown",
					ShortDescription: "Tactile switch with a gentle bump and moderate actuation force.",
					LongDescription:  "Gateron KS-8 Brown switches provide a gentle tactile bump with moderate force, offering a balanced typing experience.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2018-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron KS-8 Blue",
					ShortDescription: "Clicky switch with an audible click and tactile feedback.",
					LongDescription:  "Gateron KS-8 Blue switches offer an audible click with tactile feedback, ideal for users who enjoy clicky typing.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2018-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron KS-8 Yellow",
					ShortDescription: "Linear switch with a medium actuation force and enhanced smoothness.",
					LongDescription:  "Gateron KS-8 Yellow switches offer a medium actuation force with enhanced smoothness, providing a satisfying typing experience.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2018-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron KS-8 Black",
					ShortDescription: "Linear switch with a heavier actuation force for added resistance.",
					LongDescription:  "Gateron KS-8 Black switches offer a heavier actuation force for users who prefer more resistance in their keystrokes, providing a firm typing experience.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2018-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron Luciola",
					ShortDescription: "Linear switch with a glow-in-the-dark housing and smooth keystrokes.",
					LongDescription:  "The Gateron Luciola switch features a glow-in-the-dark housing with smooth keystrokes, offering a premium aesthetic and performance.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Gateron North Pole V2",
					ShortDescription: "Linear switches with transparent housing, optimized for RGB lighting.",
					LongDescription:  "The Gateron North Pole V2 series offers linear switches with different actuation forces, including 43g, 55g, and 62g variants. These switches feature a transparent housing, making them ideal for RGB lighting enthusiasts. Designed for smooth typing and enhanced resistance depending on the selected variant, the North Pole V2 series caters to users seeking a premium typing experience with visually striking aesthetics.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2022-01-01"),
					Available:        true,
					PricePoint:       3,  // Expensive
					SiteURL:          "", // Update if an official URL becomes available
				},
				{
					Name:             "Gateron Phantom Red",
					ShortDescription: "Linear switch with a light actuation force and smooth keystrokes.",
					LongDescription:  "Gateron Phantom Red switches offer smooth keystrokes with a light actuation force, providing a satisfying typing experience.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron Phantom Brown",
					ShortDescription: "Tactile switch with a gentle bump and moderate actuation force.",
					LongDescription:  "Gateron Phantom Brown switches provide a gentle tactile bump with moderate actuation force, offering a balanced typing experience.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron Phantom Blue",
					ShortDescription: "Clicky switch with an audible click and tactile feedback.",
					LongDescription:  "Gateron Phantom Blue switches offer a clicky typing experience with an audible click and tactile feedback, providing a satisfying typing experience.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron Aliaz Silent",
					ShortDescription: "Silent tactile switch with a gentle to strong bump and quiet operation.",
					LongDescription:  "Gateron Aliaz Silent switches are designed for users seeking a quiet tactile typing experience. Available in different spring weights such as 60g and 70g, these switches offer a range of tactile feedback from a gentle to a stronger bump. The Aliaz Silent switches are ideal for users who prefer a premium tactile feel with reduced noise, making them perfect for stealthy operation in quiet environments.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(5), // ZealPC
					SwitchTypeID:     5,      // Silent Tactile
					ReleaseDate:      parseDate("2018-01-01"),
					Available:        true,
					PricePoint:       3,  // Expensive
					SiteURL:          "", // Update if an official URL becomes available
				},
				{
					Name:             "Gateron Melodic",
					ShortDescription: "Tactile switch with a musical sound profile and smooth operation.",
					LongDescription:  "The Gateron Melodic switch is a tactile switch designed to provide a unique typing experience with a musical sound profile. With an actuation force of 55 grams, the Melodic switch offers a gentle tactile bump along with a distinctive sound that sets it apart from other switches. This switch is ideal for users who enjoy a tactile feedback with an added element of auditory satisfaction, making typing not only functional but also enjoyable.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Gateron KS-9 Red",
					ShortDescription: "Linear switch with improved smoothness and sound.",
					LongDescription:  "Gateron KS-9 Red switches are enhanced linear switches known for their improved smoothness and sound, offering an upgraded typing experience.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron KS-9 Brown",
					ShortDescription: "Tactile switch with a gentle bump and enhanced feel.",
					LongDescription:  "Gateron KS-9 Brown switches offer a gentle tactile bump and enhanced feel, providing a satisfying typing experience.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron KS-9 Blue",
					ShortDescription: "Clicky switch with an audible click and tactile feedback.",
					LongDescription:  "Gateron KS-9 Blue switches offer an audible click with tactile feedback, ideal for users who enjoy a clicky typing experience.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron KS-9 Yellow",
					ShortDescription: "Linear switch with a medium actuation force and smooth keystrokes.",
					LongDescription:  "Gateron KS-9 Yellow switches offer a medium actuation force with smooth keystrokes, providing a versatile typing experience.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron KS-9 Black",
					ShortDescription: "Linear switch with a heavier actuation force for added resistance.",
					LongDescription:  "Gateron KS-9 Black switches offer a heavier actuation force, providing a firm typing experience for users who prefer more resistance.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron KS-9 Green",
					ShortDescription: "Clicky switch with a strong tactile bump and audible click.",
					LongDescription:  "Gateron KS-9 Green switches provide a strong tactile bump with an audible click, offering a satisfying typing experience for clicky enthusiasts.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron KS-9 Clear",
					ShortDescription: "Linear switch with a light actuation force and smooth feel.",
					LongDescription:  "Gateron KS-9 Clear switches offer a light actuation force and smooth feel, ideal for users who prefer a softer typing experience.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				// Milky Series
				{
					Name:             "Gateron Milky Red",
					ShortDescription: "Linear switch with a milky housing for a softer feel.",
					LongDescription:  "Gateron Milky Red switches feature a milky housing that provides a softer feel and muted sound, offering a smooth linear experience. With an actuation force of 45 grams, these switches are designed for users who seek a light and responsive typing experience. The milky housing not only enhances the aesthetic appeal of the switch but also contributes to a quieter and softer bottom-out, making it suitable for a wide range of applications.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2017-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron Milky Brown",
					ShortDescription: "Tactile switch with a milky housing for a unique feel.",
					LongDescription:  "Gateron Milky Brown switches provide a gentle tactile bump with a milky housing, offering a unique typing experience with a softer bottom-out feel. With an actuation force of 55 grams, these switches are ideal for users who prefer tactile feedback without the loud click. The milky housing also helps to produce a more subdued sound, making these switches a great choice for quieter environments.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2017-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron Milky Blue",
					ShortDescription: "Clicky switch with a milky housing for a softer sound.",
					LongDescription:  "Gateron Milky Blue switches offer a clicky typing experience with a softer sound profile, thanks to the milky housing design. With an actuation force of 60 grams, these switches deliver a pronounced tactile bump and an audible click, providing a satisfying and responsive typing experience. The milky housing not only adds to the switch's aesthetic appeal but also dampens the sound slightly, making it a more subtle clicky option.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2017-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron Milky Black",
					ShortDescription: "Linear switch with a heavier actuation force and milky housing.",
					LongDescription:  "Gateron Milky Black switches provide a firm linear feel with a milky housing, offering a unique typing experience with added resistance. With an actuation force of 60 grams, these switches are designed for users who prefer a heavier keystroke. The milky housing contributes to a softer bottom-out and a more muted sound profile, making these switches ideal for those who want a quieter, yet still robust, linear switch.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2017-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron Milky Green",
					ShortDescription: "Clicky switch with a strong tactile bump and milky housing.",
					LongDescription:  "Gateron Milky Green switches offer a strong tactile bump with a milky housing, providing a unique clicky typing experience with a softer sound. With an actuation force of 80 grams, these switches are designed for users who prefer a heavier clicky switch with pronounced feedback. The milky housing not only enhances the switch's visual appeal but also helps to reduce the sharpness of the click sound, making it a more refined option for clicky switch enthusiasts.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2017-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron Milky Clear",
					ShortDescription: "Linear switch with a light actuation force and milky housing.",
					LongDescription:  "Gateron Milky Clear switches offer a light actuation force and smooth keystrokes with a milky housing, providing a softer feel and quieter operation. With an actuation force of 35 grams, these switches are perfect for users who prefer a very light and effortless typing experience. The milky housing also contributes to a more subdued sound profile, making these switches ideal for those who value both performance and quietness in their typing setup.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2017-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				// Box Series
				{
					Name:             "Gateron Box Red",
					ShortDescription: "Linear switch with box design for added stability.",
					LongDescription:  "Gateron Box Red switches offer a smooth linear experience with a box design that enhances stability and protection against dust.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Gateron Box Brown",
					ShortDescription: "Tactile switch with box design for enhanced feel.",
					LongDescription:  "Gateron Box Brown switches offer a gentle tactile bump with a box design, providing a stable typing experience and protection against dust.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Gateron Box Blue",
					ShortDescription: "Clicky switch with box design for a crisp click.",
					LongDescription:  "Gateron Box Blue switches provide a clicky typing experience with a box design, enhancing the crispness of each click and offering stability.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Gateron Box Black",
					ShortDescription: "Linear switch with a heavier feel and box design.",
					LongDescription:  "Gateron Box Black switches offer a heavier linear feel with a box design, providing stability and protection against dust and debris.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				// Silent Clear
				{
					Name:             "Gateron Silent Clear",
					ShortDescription: "Silent linear switch with a light actuation force.",
					LongDescription:  "Gateron Silent Clear switches offer a quiet linear typing experience with a light actuation force, ideal for stealthy typing environments.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2019-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				// Milu Series
				{
					Name:             "Gateron Milu",
					ShortDescription: "Tactile switch designed for a unique typing experience.",
					LongDescription:  "Gateron Milu switches are designed to offer a unique tactile typing experience, providing a satisfying tactile bump and smooth keystrokes.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2022-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				// Kangaroo Ink
				{
					Name:             "Gateron Kangaroo Ink",
					ShortDescription: "Tactile switch in the Ink series with a distinctive bump.",
					LongDescription:  "Gateron Kangaroo Ink switches are tactile switches within the premium Ink series, offering a distinctive bump and smooth actuation. With an actuation force of 60 grams, these switches deliver a strong tactile response, making them ideal for users who prefer a more assertive typing experience. The unique Ink housing not only provides a premium look and feel but also enhances the switch's overall durability and performance.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				// Hippo Switch
				{
					Name:             "Gateron Hippo",
					ShortDescription: "Linear switch with a thocky sound profile.",
					LongDescription:  "Gateron Hippo switches are linear switches renowned for their smooth keystrokes and distinctive thocky sound profile. With an actuation force of 50 grams, these switches are designed for users who appreciate a smooth and consistent typing experience with a deep, satisfying sound. The Hippo switches are a perfect choice for enthusiasts seeking a switch that combines performance with a unique auditory experience.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2022-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
			}

			err = processSwitches(gateron)
			if err != nil {
				return err
			}
			return nil // Replace with actual code
		},
		Rollback: func(tx *gorm.DB) error {
			if err := tx.Exec("DELETE FROM switches WHERE id NOT null").Error; err != nil {
				log.Printf("Failed to delete from Switch table: %v", err)
				return err
			}
			return nil // Replace with actual code
		},
	})
}
