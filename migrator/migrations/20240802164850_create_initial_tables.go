package migrations

import (
	"time"

	"github.com/Bparsons0904/deadigations"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	deadigations.RegisterMigration(deadigations.Migration{
		ID:          "20240802164850",
		Description: "Create the initial tables",
		Migrate: func(tx *gorm.DB) error {
			err := tx.Transaction(func(tx *gorm.DB) error {
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
				if err := tx.Migrator().CreateTable(&ImageLink{}); err != nil {
					return err
				}

				type Type struct {
					ID        int       `gorm:"primaryKey;autoIncrement"              json:"id"`
					Name      string    `gorm:"type:varchar(50);not null"             json:"name"`
					Code      string    `gorm:"type:varchar(50);not null;uniqueIndex" json:"code"`
					Category  string    `gorm:"type:varchar(50);not null;index"       json:"category"`
					CreatedAt time.Time `gorm:"autoCreateTime"                        json:"createdAt"`
					UpdatedAt time.Time `gorm:"autoUpdateTime"                        json:"updatedAt"`
				}
				if err := tx.Migrator().CreateTable(&Type{}); err != nil {
					return err
				}

				type Details struct {
					ID                    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v7();primaryKey" json:"id"`
					SwitchID              uuid.UUID `gorm:"type:uuid;not null;index"                        json:"switchId"`
					Version               *int      `gorm:"type:int"                                        json:"version,omitempty"`
					SpringTypeID          *int      `gorm:"type:int"                                        json:"springTypeId,omitempty"`
					SpringForce           *float32  `gorm:"type:real"                                       json:"springForce,omitempty"`
					SpringMaterialTypeID  *int      `gorm:"type:int"                                        json:"springMaterialTypeId,omitempty"`
					SpringMaterialType    *Type     `gorm:"foreignKey:SpringMaterialTypeID"                 json:"springMaterialType,omitempty"`
					TopHousingMaterialID  *int      `gorm:"type:int"                                        json:"topHousingMaterialId,omitempty"`
					TopHousingMaterial    *Type     `gorm:"foreignKey:TopHousingMaterialID"                 json:"topHousingMaterial,omitempty"`
					BaseHousingMaterialID *int      `gorm:"type:int"                                        json:"baseHousingMaterialId,omitempty"`
					BaseHousingMaterial   *Type     `gorm:"foreignKey:BaseHousingMaterialID"                json:"baseHousingMaterial,omitempty"`
					StemMaterialID        *int      `gorm:"type:int"                                        json:"stemMaterialId,omitempty"`
					StemMaterial          *Type     `gorm:"foreignKey:StemMaterialID"                       json:"stemMaterial,omitempty"`
					HasShineThrough       *bool     `gorm:"type:boolean"                                    json:"hasShineThrough,omitempty"`
					PreTravel             *float32  `gorm:"type:real"                                       json:"preTravel,omitempty"`
					TotalTravel           *float32  `gorm:"type:real"                                       json:"totalTravel,omitempty"`
					InitialForce          *float32  `gorm:"type:real"                                       json:"initialForce,omitempty"`
					ActuationPoint        *float32  `gorm:"type:real"                                       json:"actuationPoint,omitempty"`
					ResetPoint            *float32  `gorm:"type:real"                                       json:"resetPoint,omitempty"`
					BottomOutForcePoint   *float32  `gorm:"type:real"                                       json:"bottomOutForcePoint,omitempty"`
					BottomOutForce        *float32  `gorm:"type:real"                                       json:"bottomOutForce,omitempty"`
					SoundLevelID          *int      `gorm:"type:int"                                        json:"soundLevelId,omitempty"`
					SoundLevel            *Type     `gorm:"foreignKey:SoundLevelID"                         json:"soundLevel,omitempty"`
					SoundTypeID           *int      `gorm:"type:int"                                        json:"soundTypeId,omitempty"`
					SoundType             *Type     `gorm:"foreignKey:SoundTypeID"                          json:"soundType,omitempty"`
					TactilityTypeID       *int      `gorm:"type:int"                                        json:"tactilityTypeId,omitempty"`
					TactilityType         *Type     `gorm:"foreignKey:TactilityTypeID"                      json:"tactilityType,omitempty"`
					BumpPosition          *float32  `gorm:"type:real"                                       json:"bumpPosition,omitempty"`
					BumpForce             *float32  `gorm:"type:real"                                       json:"bumpForce,omitempty"`
					TactilityFeedbackID   *int      `gorm:"type:int"                                        json:"tactilityFeedbackId,omitempty"`
					TactilityFeedback     *Type     `gorm:"foreignKey:TactilityFeedbackID"                  json:"tactilityFeedback,omitempty"`
					FactoryLubed          bool      `gorm:"type:boolean;default:false"                      json:"factoryLubed"`
					StemColorID           *int      `gorm:"type:int"                                        json:"stemColorId,omitempty"`
					StemColor             *Type     `gorm:"foreignKey:StemColorID"                          json:"stemColor,omitempty"`
					TopHousingColorID     *int      `gorm:"type:int"                                        json:"topHousingColorId,omitempty"`
					TopHousingColor       *Type     `gorm:"foreignKey:TopHousingColorID"                    json:"topHousingColor,omitempty"`
					BottomHousingColorID  *int      `gorm:"type:int"                                        json:"bottomHousingColorId,omitempty"`
					BottomHousingColor    *Type     `gorm:"foreignKey:BottomHousingColorID"                 json:"bottomHousingColor,omitempty"`
					PinConfigurationID    *int      `gorm:"type:int"                                        json:"pinConfigurationId,omitempty"`
					PinConfiguration      *Type     `gorm:"foreignKey:PinConfigurationID"                   json:"pinConfiguration,omitempty"`
					CreatedAt             time.Time `gorm:"type:timestamp;autoCreateTime"                   json:"createdAt"`
					UpdatedAt             time.Time `gorm:"type:timestamp;autoUpdateTime"                   json:"updatedAt"`
				}
				if err := tx.Migrator().CreateTable(&Details{}); err != nil {
					return err
				}

				type Producer struct {
					ID        int       `gorm:"type:int;primaryKey;autoIncrement"     json:"id"`
					Name      string    `gorm:"type:varchar(50);not null"             json:"name"`
					Alias     string    `gorm:"type:varchar(50);not null;uniqueIndex" json:"alias"`
					SiteURL   string    `gorm:"type:varchar(50)"                      json:"siteURL,omitempty"`
					CreatedAt time.Time `gorm:"autoCreateTime"                        json:"createdAt"`
					UpdateAt  time.Time `gorm:"autoUpdateTime"                        json:"updatedAt"`
				}
				if err := tx.Migrator().CreateTable(&Producer{}); err != nil {
					return err
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
					SiteURL          string     `gorm:"type:varchar(255)"                                                                json:"siteURL,omitempty"`
					CreatedAt        time.Time  `gorm:"autoCreateTime"                                                                   json:"createdAt"`
					UpdatedAt        time.Time  `gorm:"autoUpdateTime"                                                                   json:"updatedAt"`
				}
				if err := tx.Migrator().CreateTable(&Switch{}); err != nil {
					return err
				}

				types := []Type{
					// Switch Types
					{Name: "Linear", Code: "linear", Category: "switch_type"},
					{Name: "Tactile", Code: "tactile", Category: "switch_type"},
					{Name: "Clicky", Code: "clicky", Category: "switch_type"},
					{Name: "Silent Linear", Code: "silent_linear", Category: "switch_type"},
					{Name: "Silent Tactile", Code: "silent_tactile", Category: "switch_type"},

					// Spring Types
					{Name: "Linear", Code: "spring_linear", Category: "spring_type"},
					{Name: "Progressive", Code: "progressive", Category: "spring_type"},
					{Name: "Slow", Code: "slow", Category: "spring_type"},
					{Name: "Fast", Code: "fast", Category: "spring_type"},
					{Name: "Variable", Code: "variable", Category: "spring_type"},
					{Name: "Balanced", Code: "balanced", Category: "spring_type"},
					{Name: "Soft", Code: "soft", Category: "spring_type"},
					{Name: "Heavy", Code: "heavy", Category: "spring_type"},
					{Name: "Custom", Code: "custom", Category: "spring_type"},
					{Name: "Dual-Stage", Code: "dual_stage", Category: "spring_type"},
					{Name: "Triple-Stage", Code: "triple_stage", Category: "spring_type"},

					// Spring Materials
					{Name: "Steel", Code: "steel", Category: "spring_material"},
					{
						Name:     "Gold-Plated Steel",
						Code:     "gold_plated_steel",
						Category: "spring_material",
					},
					{Name: "Copper", Code: "copper", Category: "spring_material"},
					{Name: "Phosphor Bronze", Code: "phosphor_bronze", Category: "spring_material"},
					{Name: "Stainless Steel", Code: "stainless_steel", Category: "spring_material"},

					// Plastic Materials
					{Name: "ABS", Code: "abs", Category: "material"},
					{Name: "POM", Code: "pom", Category: "material"},
					{Name: "Nylon", Code: "nylon", Category: "material"},
					{Name: "UHMWPE", Code: "uhmwpe", Category: "material"},
					{Name: "Polycarbonate", Code: "pc", Category: "material"},
					{Name: "PBT", Code: "pbt", Category: "material"},

					// Sound Levels
					{Name: "Very Low", Code: "very_low", Category: "sound_level"},
					{Name: "Low", Code: "low", Category: "sound_level"},
					{Name: "Medium", Code: "medium", Category: "sound_level"},
					{Name: "High", Code: "high", Category: "sound_level"},
					{Name: "Very High", Code: "very_high", Category: "sound_level"},

					// Sound Types
					{Name: "Clicky", Code: "sound_clicky", Category: "sound_type"},
					{Name: "Thocky", Code: "sound_thocky", Category: "sound_type"},
					{Name: "Quiet", Code: "sound_quiet", Category: "sound_type"},
					{Name: "Creamy", Code: "sound_creamy", Category: "sound_type"},
					{Name: "Clacky", Code: "sound_clacky", Category: "sound_type"},

					// Tactility Types
					{Name: "Leaf Spring", Code: "leaf_spring", Category: "tactility_type"},
					{Name: "Coil Spring", Code: "coil_spring", Category: "tactility_type"},
					{Name: "Click Bar", Code: "click_bar", Category: "tactility_type"},

					// Tactility Feedback
					{Name: "Sharp", Code: "sharp", Category: "tactility_feedback"},
					{Name: "Rounded", Code: "rounded", Category: "tactility_feedback"},
					{Name: "Crisp", Code: "crisp", Category: "tactility_feedback"},
					{Name: "Smooth", Code: "smooth", Category: "tactility_feedback"},

					// Colors
					{Name: "Red", Code: "red", Category: "color"},
					{Name: "Black", Code: "black", Category: "color"},
					{Name: "Blue", Code: "blue", Category: "color"},
					{Name: "Brown", Code: "brown", Category: "color"},
					{Name: "Clear", Code: "clear", Category: "color"},
					{Name: "Yellow", Code: "yellow", Category: "color"},
					{Name: "White", Code: "white", Category: "color"},
					{Name: "Transparent", Code: "transparent", Category: "color"},
					{Name: "Smokey", Code: "smokey", Category: "color"},
					{Name: "Green", Code: "green", Category: "color"},
					{Name: "Purple", Code: "purple", Category: "color"},
					{Name: "Orange", Code: "orange", Category: "color"},
					{Name: "Pink", Code: "pink", Category: "color"},
					{Name: "Gray", Code: "gray", Category: "color"},
					{Name: "Silver", Code: "silver", Category: "color"},
					{Name: "Gold", Code: "gold", Category: "color"},
					{Name: "Turquoise", Code: "turquoise", Category: "color"},
					{Name: "Teal", Code: "teal", Category: "color"},
					{Name: "Lavender", Code: "lavender", Category: "color"},
					{Name: "Magenta", Code: "magenta", Category: "color"},
					{Name: "Cyan", Code: "cyan", Category: "color"},
					{Name: "Ivory", Code: "ivory", Category: "color"},
					{Name: "Coral", Code: "coral", Category: "color"},
					{Name: "Maroon", Code: "maroon", Category: "color"},
					{Name: "Beige", Code: "beige", Category: "color"},
					{Name: "Mint", Code: "mint", Category: "color"},
					{Name: "Peach", Code: "peach", Category: "color"},
					{Name: "Tan", Code: "tan", Category: "color"},
					{Name: "Khaki", Code: "khaki", Category: "color"},

					// Pin Configuration
					{Name: "3-Pin", Code: "3_pin", Category: "pin_configuration"},
					{Name: "5-Pin", Code: "5_pin", Category: "pin_configuration"},
				}

				if err := tx.Create(&types).Error; err != nil {
					return err
				}

				producers := []Producer{
					{
						Name:    "Cherry",
						Alias:   "cherry",
						SiteURL: "https://www.cherrymx.de/",
					},
					{
						Name:    "Gateron",
						Alias:   "gateron",
						SiteURL: "https://www.gateron.com/",
					},
					{
						Name:    "Kailh",
						Alias:   "kailh",
						SiteURL: "https://www.kailh.com/",
					},
					{
						Name:    "Outemu",
						Alias:   "outemu",
						SiteURL: "",
					},
					{
						Name:    "ZealPC",
						Alias:   "zealpc",
						SiteURL: "https://zealpc.net/",
					},
					{
						Name:    "TTC",
						Alias:   "ttc",
						SiteURL: "",
					},
					{
						Name:    "Durock",
						Alias:   "durock",
						SiteURL: "https://www.durock.us/",
					},
					{
						Name:    "Akko",
						Alias:   "akko",
						SiteURL: "https://en.akkogear.com/",
					},
					{
						Name:    "NovelKeys",
						Alias:   "novelkeys",
						SiteURL: "https://novelkeys.com/",
					},
					{
						Name:    "Drop",
						Alias:   "drop",
						SiteURL: "https://drop.com/",
					},
					{
						Name:    "Glorious",
						Alias:   "glorious",
						SiteURL: "https://www.gloriousgaming.com/",
					},
					{
						Name:    "Keychron",
						Alias:   "keychron",
						SiteURL: "https://www.keychron.com/",
					},
					{
						Name:    "Epomaker",
						Alias:   "epomaker",
						SiteURL: "https://epomaker.com/",
					},
					{
						Name:    "Razer",
						Alias:   "razer",
						SiteURL: "https://www.razer.com/",
					},
					{
						Name:    "Logitech",
						Alias:   "logitech",
						SiteURL: "https://www.logitechg.com/",
					},
					{
						Name:    "SteelSeries",
						Alias:   "steelseries",
						SiteURL: "https://steelseries.com/",
					},
					{
						Name:    "Kinetic Labs",
						Alias:   "kintic_labs",
						SiteURL: "https://kineticlabs.com/",
					},
					{
						Name:    "Chosfox",
						Alias:   "chosfox",
						SiteURL: "https://chosfox.com/",
					},
					{
						Name:    "Roccat",
						Alias:   "roccat",
						SiteURL: "https://en.roccat.com/",
					},
					{
						Name:    "Cooler Master",
						Alias:   "cooler_master",
						SiteURL: "https://www.coolermaster.com/",
					},
					{
						Name:    "Wuque Studio",
						Alias:   "wuque_studio",
						SiteURL: "https://shop.wuquestudio.com/",
					},
					{
						Name:    "JWK",
						Alias:   "jwk",
						SiteURL: "http://www.jwk-tech.com/",
					},
					{
						Name:    "Tecsee",
						Alias:   "tecsee",
						SiteURL: "http://www.tecsee.com/",
					},
					{
						Name:    "Everglide",
						Alias:   "everglide",
						SiteURL: "https://everglide.com/",
					},
					{
						Name:    "SP-Star",
						Alias:   "sp_star",
						SiteURL: "",
					},
					{
						Name:    "Gazzew",
						Alias:   "gazzew",
						SiteURL: "",
					},
					{
						Name:    "Keydous",
						Alias:   "keydous",
						SiteURL: "https://www.keydous.com/",
					},
					{
						Name:    "Cannon Keys",
						Alias:   "cannon_keys",
						SiteURL: "https://cannonkeys.com/",
					},
					{
						Name:    "KBDFans",
						Alias:   "kbd_fans",
						SiteURL: "https://kbdfans.com/",
					},
					{
						Name:    "HMX",
						Alias:   "hmx",
						SiteURL: "",
					},
					{
						Name:    "Tbcats Studio",
						Alias:   "tbcats",
						SiteURL: "",
					},
					{
						Name:    "C³ EQUALZ X TKC",
						Alias:   "c3xtkc",
						SiteURL: "https://c3equalz.com/",
					},
				}

				if err := tx.Create(&producers).Error; err != nil {
					return err
				}

				return nil
			})
			return err
		},
		Rollback: func(tx *gorm.DB) error {
			err := tx.Transaction(func(tx *gorm.DB) error {
				if err := tx.Migrator().DropTable("switches"); err != nil {
					return err
				}
				if err := tx.Migrator().DropTable("producer"); err != nil {
					return err
				}
				if err := tx.Migrator().DropTable("details"); err != nil {
					return err
				}
				if err := tx.Migrator().DropTable("types"); err != nil {
					return err
				}

				return nil
			})
			return err
		},
	})
}
