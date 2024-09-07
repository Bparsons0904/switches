package migrations

import (
	"time"

	"github.com/Bparsons0904/deadigations"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	deadigations.RegisterMigration(deadigations.Migration{
		ID:          "20240906221052",
		Description: "Add description of changes",
		Migrate: func(tx *gorm.DB) error {
			type User struct {
				ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v7();primaryKey" json:"id"`
			}
			type Producer struct {
				ID int `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
			}
			type Type struct {
				ID int `gorm:"primaryKey;autoIncrement" json:"id"`
			}
			type ImageLink struct {
				ID          uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v7();primaryKey" json:"id"`
				LinkAddress string         `gorm:"type:varchar(255);not null"                      json:"linkAddress"`
				AltText     string         `gorm:"type:varchar(255);default:''"                    json:"altText,omitempty"`
				Thumbnail   string         `gorm:"type:varchar(255);default:''"                    json:"thumbnail,omitempty"`
				FullSize    string         `gorm:"type:varchar(255);default:''"                    json:"fullSize,omitempty"`
				IsPrimary   bool           `gorm:"default:false"                                   json:"isPrimary"`
				OwnerID     uuid.UUID      `gorm:"type:uuid;indes:idx_type_id"                     json:"objectId"`
				OwnerType   string         `gorm:"type:varchar(50);index:idx_type_id"              json:"objectType"`
				CreatedAt   time.Time      `gorm:"autoCreateTime"                                  json:"createdAt"`
				UpdatedAt   time.Time      `gorm:"autoUpdateTime"                                  json:"updatedAt"`
				DeletedAt   gorm.DeletedAt `gorm:"index"                                           json:"deletedAt"`
				CreatedByID uuid.UUID      `gorm:"type:uuid"                                       json:"createdById"`
				CreatedBy   User           `gorm:"foreignKey:CreatedByID;references:ID"            json:"createdBy,omitempty"`
				UpdateByID  uuid.UUID      `gorm:"type:uuid"                                       json:"updateById"`
				UpdateBy    User           `gorm:"foreignKey:UpdateByID;references:ID"             json:"updateBy,omitempty"`
			}
			type Switch struct {
				ID               uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v7();primaryKey"                                          json:"id"`
				Name             string         `gorm:"type:varchar(50);not null;index:idx_name;uniqueIndex:idx_name_manufacturer_type"          json:"name"`
				ShortDescription string         `gorm:"type:varchar(255);not null"                                                               json:"shortDescription"`
				LongDescription  string         `gorm:"type:text;not null"                                                                       json:"longDescription"`
				ManufacturerID   *int           `gorm:"type:int;index;uniqueIndex:idx_name_manufacturer_type"                                    json:"manufacturerId,omitempty"`
				Manufacturer     *Producer      `gorm:"foreignKey:ManufacturerID"                                                                json:"manufacturer,omitempty"`
				BrandID          *int           `gorm:"type:int;index"                                                                           json:"brandId,omitempty"`
				Brand            *Producer      `gorm:"foreignKey:BrandID"                                                                       json:"brand,omitempty"`
				SwitchTypeID     int            `gorm:"type:int;not null;index;uniqueIndex:idx_name_manufacturer_type"                           json:"switchTypeId"`
				SwitchType       *Type          `gorm:"foreignKey:SwitchTypeID"                                                                  json:"switchType,omitempty"`
				ReleaseDate      *time.Time     `gorm:"type:date"                                                                                json:"releaseDate,omitempty"`
				Available        bool           `gorm:"type:boolean;default:true"                                                                json:"available"`
				PricePoint       int            `gorm:"type:int;not null"                                                                        json:"pricePoint"`
				SiteURL          string         `gorm:"type:varchar(255)"                                                                        json:"siteURL,omitempty"`
				CreatedAt        time.Time      `gorm:"autoCreateTime"                                                                           json:"createdAt"`
				UpdatedAt        time.Time      `gorm:"autoUpdateTime"                                                                           json:"updatedAt"`
				DeletedAt        gorm.DeletedAt `gorm:"index"                                                                                    json:"deletedAt"`
				ImageLinks       []*ImageLink   `gorm:"foreignKey:OwnerID;polymorphicValue:Switch;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"imageLinks,omitempty"`
			}

			err := tx.Migrator().AutoMigrate(&Switch{})
			if err != nil {
				return err
			}

			err = tx.Migrator().AutoMigrate(&ImageLink{})
			if err != nil {
				return err
			}

			return nil // Replace with actual code
		},
		Rollback: func(tx *gorm.DB) error {
			type Producer struct {
				ID int `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
			}
			type Type struct {
				ID int `gorm:"primaryKey;autoIncrement" json:"id"`
			}
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
			type Switch struct {
				ID               uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v7();primaryKey"                                          json:"id"`
				Name             string         `gorm:"type:varchar(50);not null;index:idx_name;uniqueIndex:idx_name_manufacturer_type"          json:"name"`
				ShortDescription string         `gorm:"type:varchar(255);not null"                                                               json:"shortDescription"`
				LongDescription  string         `gorm:"type:text;not null"                                                                       json:"longDescription"`
				ManufacturerID   *int           `gorm:"type:int;index;uniqueIndex:idx_name_manufacturer_type"                                    json:"manufacturerId,omitempty"`
				Manufacturer     *Producer      `gorm:"foreignKey:ManufacturerID"                                                                json:"manufacturer,omitempty"`
				BrandID          *int           `gorm:"type:int;index"                                                                           json:"brandId,omitempty"`
				Brand            *Producer      `gorm:"foreignKey:BrandID"                                                                       json:"brand,omitempty"`
				SwitchTypeID     int            `gorm:"type:int;not null;index;uniqueIndex:idx_name_manufacturer_type"                           json:"switchTypeId"`
				SwitchType       *Type          `gorm:"foreignKey:SwitchTypeID"                                                                  json:"switchType,omitempty"`
				ReleaseDate      *time.Time     `gorm:"type:date"                                                                                json:"releaseDate,omitempty"`
				Available        bool           `gorm:"type:boolean;default:true"                                                                json:"available"`
				PricePoint       int            `gorm:"type:int;not null"                                                                        json:"pricePoint"`
				SiteURL          string         `gorm:"type:varchar(255)"                                                                        json:"siteURL,omitempty"`
				CreatedAt        time.Time      `gorm:"autoCreateTime"                                                                           json:"createdAt"`
				UpdatedAt        time.Time      `gorm:"autoUpdateTime"                                                                           json:"updatedAt"`
				DeletedAt        gorm.DeletedAt `gorm:"index"                                                                                    json:"deletedAt"`
				ImageLinks       []*ImageLink   `gorm:"foreignKey:OwnerID;polymorphicValue:Switch;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"imageLinks,omitempty"`
			}

			err := tx.Migrator().AutoMigrate(&Switch{})
			if err != nil {
				return err
			}

			err = tx.Migrator().AutoMigrate(&ImageLink{})
			if err != nil {
				return err
			}

			return nil
		},
	})
}
