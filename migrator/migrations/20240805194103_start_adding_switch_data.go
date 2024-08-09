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

			cherry := []Switch{
				{
					Name:             "Cherry MX Silent Red",
					ShortDescription: "Silent linear switch with a light actuation force.",
					LongDescription:  "Cherry MX Silent Red switches are designed to offer the smooth, linear feel of MX Red switches while significantly reducing noise during actuation. These switches incorporate built-in dampeners that minimize the sound of both the key press and release, making them an excellent choice for quiet environments like offices or shared spaces. With an actuation force of 45 grams, they maintain a light, responsive feel, ideal for fast typing or gaming without the auditory distraction.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2015-01-01"),
					Available:        true,
					PricePoint:       3,
					SiteURL:          "https://www.cherry-world.com/mx-silent-red.html",
					ImageLinks: []*ImageLink{
						{
							LinkAddress: "https://i.imgur.com/8rVjv2b.jpg",
							AltText:     "Cherry MX Silent Red switch",
						},
					},
				},

				{
					Name:             "Cherry MX Red",
					ShortDescription: "Linear switch with a light actuation force.",
					LongDescription:  "Cherry MX Red switches are renowned for their smooth, linear actuation and light operating force, making them a preferred choice for gamers and typists who enjoy a fast, uninterrupted keystroke. These switches require only 45 grams of force to actuate, ensuring a quick response without any tactile bump or audible click, allowing for a fluid typing experience. They are especially favored in gaming for their consistent key presses and rapid input, reducing finger fatigue during extended use.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("1984-01-01"),
					Available:        true,
					PricePoint:       2,
					SiteURL:          "https://www.cherry-world.com/mx-red.html",
				},
				{
					Name:             "Cherry MX Blue",
					ShortDescription: "Clicky switch with a tactile bump.",
					LongDescription:  "Cherry MX Blue switches are iconic for their audible click and tactile feedback, offering a distinct typing experience that is both responsive and satisfying. With an actuation force of 50 grams, these switches provide a noticeable bump before actuation, coupled with a sharp click sound that appeals to typists who appreciate clear auditory and physical feedback. While they are less common in gaming due to their pronounced click, MX Blues are a top choice for those who value precision and clarity in their keystrokes.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("1987-01-01"),
					Available:        true,
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
					Available:        true,
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
					Available:        true,
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
					SiteURL:          "https://www.cherry-world.com/mx-green.html",
				},
				{
					Name:             "Cherry MX Silent Black",
					ShortDescription: "Silent linear switch with a heavier actuation force.",
					LongDescription:  "Cherry MX Silent Black switches are engineered to provide the quiet operation of Silent Reds with a higher actuation force of 60 grams. This switch is ideal for users who prefer a more deliberate keystroke, combining the smooth, linear action with noise-reducing technology. Silent Black switches are well-suited for environments that require both quiet and a firm keypress, balancing performance with sound reduction.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     4,      // Silent Linear
					ReleaseDate:      parseDate("2015-01-01"),
					Available:        true,
					PricePoint:       3,
					SiteURL:          "https://www.cherry-world.com/mx-silent-black.html",
				},
				{
					Name:             "Cherry MX Clear",
					ShortDescription: "Tactile switch with a high actuation force.",
					LongDescription:  "Cherry MX Clear switches are designed for those who prefer a heavier tactile bump in their typing experience. With an actuation force of 65 grams, these switches provide a noticeable resistance, making each keystroke deliberate and satisfying. MX Clear switches are often chosen by users who want a tactile switch that offers more feedback than MX Browns but without the clicky noise associated with MX Blue switches.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("1994-01-01"),
					Available:        true,
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
					Available:        true,
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
					Available:        true,
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
					Available:        true,
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
					Available:        true,
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
					Name:             "Cherry MX2A Speed Silver RGB Linear",
					ShortDescription: "Linear switch with fast actuation and RGB support.",
					LongDescription:  "The Cherry MX2A Speed Silver RGB Linear switch is optimized for rapid actuation, making it ideal for gaming. With a short 1.2mm actuation distance and an actuation force of 45 grams, this switch allows for quick, responsive keystrokes. It also features RGB compatibility, enabling customizable lighting effects to enhance the aesthetic appeal of your keyboard. These switches are perfect for gamers who want both performance and style in their setup.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2024-03-01"),
					Available:        true,
					PricePoint:       3, // Expensive
					SiteURL:          "",
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
					SiteURL:          "",
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
					SiteURL:          "",
				},
				{
					Name:             "Cherry MX2A Blue Clicky RGB",
					ShortDescription: "Clicky switch with RGB lighting and tactile feedback.",
					LongDescription:  "The Cherry MX2A Blue Clicky RGB switch combines the crisp click and tactile bump of a traditional clicky switch with RGB lighting capabilities. With an actuation force of 50 grams, this switch is ideal for those who want both a satisfying typing experience and a vibrant, customizable keyboard backlight. It's a great choice for typists who enjoy both the sound and visual appeal of a well-lit keyboard.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2024-03-01"),
					Available:        true,
					PricePoint:       2, // Average
					SiteURL:          "",
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
					SiteURL:          "",
				},
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
				},
				{
					Name:             "Cherry MX2A Black RGB Linear",
					ShortDescription: "Linear switch with RGB lighting and a heavy feel.",
					LongDescription:  "The Cherry MX2A Black RGB Linear switch offers a heavier actuation force of 60 grams, providing a robust and controlled typing experience. The switch also features RGB lighting, allowing for vibrant, customizable backlighting to match your keyboard setup. This combination makes it ideal for users who prefer a solid, linear keystroke with the added benefit of visual flair.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2024-03-01"),
					Available:        true,
					PricePoint:       2, // Average
					SiteURL:          "",
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
					SiteURL:          "",
				},
				{
					Name:             "Cherry MX2A Red RGB Linear",
					ShortDescription: "Linear switch with RGB support and a smooth keystroke.",
					LongDescription:  "The Cherry MX2A Red RGB Linear switch combines a smooth keystroke with RGB lighting, providing both visual appeal and a satisfying typing experience. With an actuation force of 45 grams, this switch is light and responsive, making it suitable for both gaming and typing. The added RGB support enhances the keyboard's aesthetics, making it a great choice for users who value both performance and style.",
					ManufacturerID:   ptr(1), // Cherry
					BrandID:          ptr(1), // Cherry
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2024-03-01"),
					Available:        true,
					PricePoint:       2, // Average
					SiteURL:          "",
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
					SiteURL:          "",
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
					Name:             "Gateron Red",
					ShortDescription: "Linear switch with a light actuation force.",
					LongDescription:  "Gateron Red switches are known for their smooth linear action and light actuation force, making them ideal for gaming and typing.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2014-01-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				{
					Name:             "Gateron Blue",
					ShortDescription: "Clicky switch with a tactile bump.",
					LongDescription:  "Gateron Blue switches offer a clicky and tactile typing experience, popular among typists who enjoy auditory feedback.",
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
					LongDescription:  "Gateron Brown switches provide a tactile bump without a click, offering a balanced feel suitable for both gaming and typing.",
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
					LongDescription:  "Gateron Black switches are designed for those who prefer a stiffer linear feel, providing smooth keystrokes with a higher actuation force.",
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
					LongDescription:  "Gateron Yellow switches provide a smooth linear feel with a medium actuation force, offering a balanced experience for various applications.",
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
					LongDescription:  "Gateron Green switches are similar to Gateron Blue but with a higher actuation force, providing a more pronounced tactile and clicky feedback.",
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
					LongDescription:  "Gateron Clear switches offer an ultra-light linear feel, making them ideal for users who prefer minimal resistance during keystrokes.",
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
					LongDescription:  "Gateron Silent Red switches offer a smooth linear feel with dampened noise, perfect for quiet environments and stealthy gaming.",
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
					LongDescription:  "Gateron Silent Black switches provide a quiet, smooth linear feel with a higher actuation force, suitable for users who prefer more resistance.",
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
					LongDescription:  "Gateron Pro Yellow switches offer an enhanced linear feel with factory lubrication, providing a smoother experience compared to standard Gateron Yellows.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron Cap V2 Brown",
					ShortDescription: "Upgraded tactile switch with a moderate actuation force.",
					LongDescription:  "Gateron Cap V2 Brown switches feature an upgraded tactile experience with enhanced sound dampening, providing a more refined tactile feel.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
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
					Name:             "Gateron Cap V2 Blue",
					ShortDescription: "Upgraded clicky switch with a tactile bump.",
					LongDescription:  "Gateron Cap V2 Blue switches feature an upgraded clicky experience with enhanced sound and tactile feedback, providing a satisfying typing feel.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
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
					Name:             "Gateron Cap V2 Brown Pro",
					ShortDescription: "Upgraded tactile switch with a refined bump.",
					LongDescription:  "Gateron Cap V2 Brown Pro switches feature a refined tactile bump and sound dampening, offering a quieter typing experience with enhanced feedback.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
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
					Name:             "Gateron Cap V2 Yellow",
					ShortDescription: "Upgraded linear switch with a smooth feel and factory lubrication.",
					LongDescription:  "Gateron Cap V2 Yellow switches provide an enhanced linear typing experience with factory lubrication for smoother keystrokes.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
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
					Name:             "Gateron Baby Kangaroo",
					ShortDescription: "Linear switch with a unique spring design.",
					LongDescription:  "Gateron Baby Kangaroo switches offer a distinctive actuation feel, featuring a unique spring design for a novel linear experience.",
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
					LongDescription:  "Gateron Milky Yellow switches are popular for their smooth feel and budget-friendly price, featuring a milky housing for a unique aesthetic.",
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
					LongDescription:  "Gateron Ink Yellow switches are designed for a smooth typing experience with a mid-range actuation force, offering a premium build quality in the Ink series.",
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
					LongDescription:  "The Gateron Baby Kangaroo V2 offers an enhanced tactile experience with a stronger bump and smoother operation, perfect for users seeking improved feedback.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2022-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron Cap Brown V2",
					ShortDescription: "Upgraded tactile switch with sound dampening and enhanced feel.",
					LongDescription:  "Gateron Cap Brown V2 switches offer an upgraded tactile experience with sound dampening and an enhanced feel, providing a quieter and smoother typing experience.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Gateron Cap Yellow V2",
					ShortDescription: "Enhanced linear switch with improved smoothness and factory lubrication.",
					LongDescription:  "Gateron Cap Yellow V2 switches provide an enhanced linear feel with improved smoothness and factory lubrication, offering a superior typing experience.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Gateron Cap Milky Yellow",
					ShortDescription: "Linear switch with a milky housing for a smooth and cost-effective option.",
					LongDescription:  "Gateron Cap Milky Yellow switches are known for their smooth feel and cost-effectiveness, featuring a milky housing for a distinct look and feel.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2019-01-01"),
					Available:        true,
					PricePoint:       1, // Value
				},
				{
					Name:             "Gateron Cap Red",
					ShortDescription: "Linear switch with enhanced performance and sound dampening.",
					LongDescription:  "Gateron Cap Red switches offer enhanced performance and sound dampening, providing smooth keystrokes with a quiet operation, ideal for fast typing.",
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
					Name:             "Gateron G Pro Red",
					ShortDescription: "Linear switch with factory lubrication for smoother keystrokes.",
					LongDescription:  "Gateron G Pro Red switches are enhanced linear switches featuring factory lubrication for smoother keystrokes and improved performance.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron G Pro Yellow",
					ShortDescription: "Linear switch with a medium actuation force and improved smoothness.",
					LongDescription:  "Gateron G Pro Yellow switches offer a medium actuation force with factory lubrication, providing a smooth and responsive typing experience.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron G Pro Brown",
					ShortDescription: "Tactile switch with factory lubrication and a gentle bump.",
					LongDescription:  "Gateron G Pro Brown switches feature a gentle tactile bump with factory lubrication, offering a smooth and quiet typing experience.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     2,      // Tactile
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron G Pro Black",
					ShortDescription: "Linear switch with a heavier actuation force for more resistance.",
					LongDescription:  "Gateron G Pro Black switches offer a heavier actuation force for users who prefer more resistance in their keystrokes, providing a firm typing experience.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron G Pro Green",
					ShortDescription: "Clicky switch with a strong tactile bump and audible click.",
					LongDescription:  "Gateron G Pro Green switches provide a strong tactile bump with an audible click, ideal for users who enjoy feedback and sound in their typing.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2020-01-01"),
					Available:        true,
					PricePoint:       2, // Average
				},
				{
					Name:             "Gateron Ink Red",
					ShortDescription: "Premium linear switch with a smooth keystroke.",
					LongDescription:  "Gateron Ink Red switches offer a premium linear typing experience with a smooth keystroke and unique housing material for enhanced performance.",
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
					LongDescription:  "Gateron Ink Black switches provide a premium linear experience with a heavier actuation force, offering smooth keystrokes and a distinctive design.",
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
					Name:             "Gateron North Pole V2 43g",
					ShortDescription: "Linear switch with a light actuation force and transparent housing.",
					LongDescription:  "Gateron North Pole V2 43g switches offer a light actuation force with a transparent housing, ideal for RGB lighting and smooth typing.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2022-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Gateron North Pole V2 55g",
					ShortDescription: "Linear switch with a medium actuation force and transparent housing.",
					LongDescription:  "Gateron North Pole V2 55g switches offer a medium actuation force with a transparent housing, perfect for RGB lighting and a balanced typing experience.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2022-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Gateron North Pole V2 62g",
					ShortDescription: "Linear switch with a heavier actuation force and transparent housing.",
					LongDescription:  "Gateron North Pole V2 62g switches offer a heavier actuation force with a transparent housing, ideal for RGB enthusiasts and users seeking more resistance.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2022-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Gateron Oil King",
					ShortDescription: "Linear switch with smooth actuation and unique black housing.",
					LongDescription:  "Gateron Oil King switches are known for their smooth linear actuation and unique black housing, offering a premium feel and performance.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
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
					Name:             "Gateron Aliaz Silent 60g",
					ShortDescription: "Silent tactile switch with a gentle bump and quiet operation.",
					LongDescription:  "Gateron Aliaz Silent 60g switches offer a quiet tactile typing experience with a gentle bump, ideal for users seeking stealthy operation.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(5), // ZealPC
					SwitchTypeID:     5,      // Silent Tactile
					ReleaseDate:      parseDate("2018-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Gateron Aliaz Silent 70g",
					ShortDescription: "Silent tactile switch with a stronger bump and quiet operation.",
					LongDescription:  "Gateron Aliaz Silent 70g switches offer a stronger tactile bump with quiet operation, providing a premium tactile experience with reduced noise.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(5), // ZealPC
					SwitchTypeID:     5,      // Silent Tactile
					ReleaseDate:      parseDate("2018-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Gateron Melodic",
					ShortDescription: "Tactile switch with a musical sound profile and smooth operation.",
					LongDescription:  "The Gateron Melodic switch offers a tactile typing experience with a unique musical sound profile, providing a distinctive typing experience.",
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
					LongDescription:  "Gateron Milky Red switches feature a milky housing that provides a softer feel and muted sound, offering a smooth linear experience.",
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
					LongDescription:  "Gateron Milky Brown switches provide a gentle tactile bump with a milky housing, offering a unique typing experience with a softer bottom-out feel.",
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
					LongDescription:  "Gateron Milky Blue switches offer a clicky typing experience with a softer sound profile, thanks to the milky housing design.",
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
					LongDescription:  "Gateron Milky Black switches provide a firm linear feel with a milky housing, offering a unique typing experience with added resistance.",
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
					LongDescription:  "Gateron Milky Green switches offer a strong tactile bump with a milky housing, providing a unique clicky typing experience with a softer sound.",
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
					LongDescription:  "Gateron Milky Clear switches offer a light actuation force and smooth keystrokes with a milky housing, providing a softer feel.",
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
				// Ink V2 Series
				{
					Name:             "Gateron Ink V2 Red",
					ShortDescription: "Linear switch with improved smoothness and unique housing.",
					LongDescription:  "Gateron Ink V2 Red switches feature improved smoothness and unique housing, offering a premium linear typing experience.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Gateron Ink V2 Black",
					ShortDescription: "Linear switch with a heavier actuation force and improved design.",
					LongDescription:  "Gateron Ink V2 Black switches offer a premium linear experience with a heavier actuation force, featuring a unique and improved housing design.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Gateron Ink V2 Yellow",
					ShortDescription: "Linear switch with mid-range actuation and unique housing.",
					LongDescription:  "Gateron Ink V2 Yellow switches provide a smooth linear experience with a mid-range actuation force and unique housing design.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Gateron Ink V2 Blue",
					ShortDescription: "Clicky switch with unique housing for enhanced sound and feel.",
					LongDescription:  "Gateron Ink V2 Blue switches offer a clicky typing experience with enhanced sound and feel, thanks to their unique housing design.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				// CAP V2 Series
				{
					Name:             "Gateron CAP V2 Silver",
					ShortDescription: "Linear switch with smooth actuation and factory lubrication.",
					LongDescription:  "Gateron CAP V2 Silver switches feature smooth actuation with factory lubrication, offering a premium linear typing experience.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     1,      // Linear
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
				},
				{
					Name:             "Gateron CAP V2 Green",
					ShortDescription: "Clicky switch with a strong tactile bump and factory lubrication.",
					LongDescription:  "Gateron CAP V2 Green switches offer a strong tactile bump with factory lubrication, providing a premium clicky typing experience.",
					ManufacturerID:   ptr(2), // Gateron
					BrandID:          ptr(2), // Gateron
					SwitchTypeID:     3,      // Clicky
					ReleaseDate:      parseDate("2021-01-01"),
					Available:        true,
					PricePoint:       3, // Expensive
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
					LongDescription:  "Gateron Kangaroo Ink switches offer a tactile typing experience with a distinctive bump, known for their smooth actuation and unique design.",
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
					LongDescription:  "Gateron Hippo switches are known for their smooth linear actuation and thocky sound profile, offering a unique typing experience.",
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

func ptr(i int) *int {
	return &i
}

func parseDate(date string) *time.Time {
	t, _ := time.Parse("2006-01-02", date)
	return &t
}
