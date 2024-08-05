package migrations

import (
	"time"

	"github.com/Bparsons0904/deadigations"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	deadigations.RegisterMigration(deadigations.Migration{
		ID:          "20240804164723",
		Description: "Switch to the producers table, drop the manufactures table. Update some indexes and other tables columns",
		Migrate: func(tx *gorm.DB) error {
			err := tx.Transaction(func(tx *gorm.DB) error {
				type Producer struct {
					ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v7();primaryKey" json:"id"`
					Name      string    `gorm:"type:varchar(50);not null"                       json:"name"`
					Alias     string    `gorm:"type:varchar(50);not null;uniqueIndex"           json:"alias"`
					SiteURL   string    `gorm:"type:varchar(50)"                                json:"siteURL,omitempty"`
					CreatedAt time.Time `gorm:"autoCreateTime"                                  json:"createdAt"`
					UpdateAt  time.Time `gorm:"autoUpdateTime"                                  json:"updatedAt"`
				}
				if err := tx.Migrator().CreateTable(&Producer{}); err != nil {
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

				if err := tx.Migrator().DropIndex(&Type{}, "idx_category_code"); err != nil {
					return err
				}

				if err := tx.Migrator().AutoMigrate(&Type{}); err != nil {
					return err
				}

				type Switch struct {
					ID               uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v7();primaryKey"                                  json:"id"`
					Name             string     `gorm:"type:varchar(50);not null;index:idx_name;uniqueIndex:idx_name_manufacturer_brand" json:"name"`
					ShortDescription string     `gorm:"type:varchar(255);not null"                                                       json:"shortDescription"`
					LongDescription  string     `gorm:"type:text;not null"                                                               json:"longDescription"`
					ManufacturerID   *uuid.UUID `gorm:"type:uuid;index;uniqueIndex:idx_name_manufacturer_brand"                          json:"manufacturerId,omitempty"`
					Manufacturer     *Producer  `gorm:"foreignKey:ManufacturerID"                                                        json:"manufacturer,omitempty"`
					BrandID          *uuid.UUID `gorm:"type:uuid;index;uniqueIndex:idx_name_manufacturer_brand"                          json:"brandId,omitempty"`
					Brand            *Producer  `gorm:"foreignKey:BrandID"                                                               json:"brand,omitempty"`
					SwitchTypeID     int        `gorm:"type:int;not null;index;"                                                         json:"switchTypeId"`
					SwitchType       *Type      `gorm:"foreignKey:SwitchTypeID"                                                          json:"switchType,omitempty"`
					ReleaseDate      *time.Time `gorm:"type:date"                                                                        json:"releaseDate,omitempty"`
					Available        bool       `gorm:"type:boolean;default:true"                                                        json:"available"`
					PricePoint       int        `gorm:"type:int;not null"                                                                json:"pricePoint"`
					CreatedAt        time.Time  `gorm:"autoCreateTime"                                                                   json:"createdAt"`
					UpdatedAt        time.Time  `gorm:"autoUpdateTime"                                                                   json:"updatedAt"`
				}

				if err := tx.Migrator().AutoMigrate(&Switch{}); err != nil {
					return err
				}

				if err := tx.Migrator().DropTable("manufacturers"); err != nil {
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
						Name:    "Varmilo",
						Alias:   "varmilo",
						SiteURL: "https://en.varmilo.com/",
					},
					{
						Name:    "Ducky",
						Alias:   "ducky",
						SiteURL: "https://www.duckychannel.com.tw/",
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
						Name:    "AKKO X Varmilo",
						Alias:   "akko_varmilo",
						SiteURL: "https://en.akkogear.com/collections/varmilo/",
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
				}

				if err := tx.Create(&producers).Error; err != nil {
					return err
				}

				return nil // Commit transaction
			})
			return err
		},
		Rollback: func(tx *gorm.DB) error {
			err := tx.Transaction(func(tx *gorm.DB) error {
				type Manufacturer struct {
					ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v7();primaryKey" json:"id"`
					Name      string    `gorm:"type:varchar(50);not null"                       json:"name"`
					SiteURL   string    `gorm:"type:varchar(50);not null"                       json:"siteURL"`
					CreatedAt string    `gorm:"autoCreateTime"                                  json:"createdAt"`
					UpdateAt  string    `gorm:"autoUpdateTime"                                  json:"updatedAt"`
				}
				if err := tx.Migrator().CreateTable(&Manufacturer{}); err != nil {
					return err
				}

				type Type struct {
					ID        int       `gorm:"primaryKey;autoIncrement"                                json:"id"`
					Name      string    `gorm:"type:varchar(50);not null"                               json:"name"`
					Code      string    `gorm:"type:varchar(50);not null;uniqueIndex:idx_category_code" json:"code"`
					Category  string    `gorm:"type:varchar(50);not null;uniqueIndex:idx_category_code" json:"category"`
					CreatedAt time.Time `gorm:"autoCreateTime"                                          json:"createdAt"`
					UpdatedAt time.Time `gorm:"autoUpdateTime"                                          json:"updatedAt"`
				}
				if err := tx.Migrator().AutoMigrate(&Type{}); err != nil {
					return err
				}

				type Switch struct {
					ID               uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v7();primaryKey" json:"id"`
					Name             string         `gorm:"type:varchar(50);not null"                       json:"name"`
					ShortDescription string         `gorm:"type:varchar(255);not null"                      json:"shortDescription"`
					LongDescription  string         `gorm:"type:text;not null"                              json:"longDescription"`
					ManufacturerID   *uuid.UUID     `gorm:"type:uuid"                                       json:"manufacturerId,omitempty"`
					Manufacturer     *Manufacturer  `gorm:"foreignKey:ManufacturerID"                       json:"manufacturer,omitempty"`
					SwitchTypeID     int            `gorm:"type:int;not null"                               json:"switchTypeId"`
					SwitchType       *Type          `gorm:"foreignKey:SwitchTypeID"                         json:"switchType,omitempty"`
					ReleaseDate      *time.Time     `gorm:"type:date"                                       json:"releaseDate,omitempty"`
					Available        bool           `gorm:"type:boolean;default:true"                       json:"available"`
					PricePoint       int            `gorm:"type:int;not null"                               json:"pricePoint"`
					CreatedAt        time.Time      `gorm:"autoCreateTime"                                  json:"createdAt"`
					UpdatedAt        time.Time      `gorm:"autoUpdateTime"                                  json:"updatedAt"`
					DeletedAt        gorm.DeletedAt `gorm:"index"                                           json:"deletedAt,omitempty"`
				}
				if err := tx.Migrator().AutoMigrate(&Switch{}); err != nil {
					return err
				}

				if err := tx.Migrator().DropTable("producers"); err != nil {
					return err
				}

				return nil // Commit rollback transaction
			})
			return err
		},
	})
}
