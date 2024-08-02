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
				type Type struct {
					ID        int       `gorm:"primaryKey;autoIncrement"                                json:"id"`
					Name      string    `gorm:"type:varchar(50);not null"                               json:"name"`
					Code      string    `gorm:"type:varchar(50);not null;uniqueIndex:idx_category_code" json:"code"`
					Category  string    `gorm:"type:varchar(50);not null;uniqueIndex:idx_category_code" json:"category"`
					CreatedAt time.Time `gorm:"autoCreateTime"                                          json:"createdAt"`
					UpdatedAt time.Time `gorm:"autoUpdateTime"                                          json:"updatedAt"`
				}
				if err := tx.Migrator().CreateTable(&Type{}); err != nil {
					return err
				}

				type Details struct {
					SwitchID              uuid.UUID `gorm:"type:uuid;not null;primaryKey"    json:"switchId"`
					Version               *int      `gorm:"type:int;default:0"               json:"version,omitempty"`
					SpringTypeID          *int      `gorm:"type:int"                         json:"springTypeId,omitempty"`
					SpringForce           *float32  `gorm:"type:real"                        json:"springForce,omitempty"`
					SpringMaterialTypeID  *int      `gorm:"type:int"                         json:"springMaterialTypeId,omitempty"`
					SpringMaterialType    *Type     `gorm:"foreignKey:SpringMaterialTypeID"  json:"springMaterialType,omitempty"`
					TopHousingMaterialID  *int      `gorm:"type:int"                         json:"topHousingMaterialId,omitempty"`
					TopHousingMaterial    *Type     `gorm:"foreignKey:TopHousingMaterialID"  json:"topHousingMaterial,omitempty"`
					BaseHousingMaterialID *int      `gorm:"type:int"                         json:"baseHousingMaterialId,omitempty"`
					BaseHousingMaterial   *Type     `gorm:"foreignKey:BaseHousingMaterialID" json:"baseHousingMaterial,omitempty"`
					StemMaterialID        *int      `gorm:"type:int"                         json:"stemMaterialId,omitempty"`
					StemMaterial          *Type     `gorm:"foreignKey:StemMaterialID"        json:"stemMaterial,omitempty"`
					HasShineThrough       *bool     `gorm:"type:boolean"                     json:"hasShineThrough,omitempty"`
					PreTravel             *float32  `gorm:"type:real"                        json:"preTravel,omitempty"`
					TotalTravel           *float32  `gorm:"type:real"                        json:"totalTravel,omitempty"`
					InitialForce          *float32  `gorm:"type:real"                        json:"initialForce,omitempty"`
					ActuationPoint        *float32  `gorm:"type:real"                        json:"actuationPoint,omitempty"`
					ResetPoint            *float32  `gorm:"type:real"                        json:"resetPoint,omitempty"`
					BottomOutForcePoint   *float32  `gorm:"type:real"                        json:"bottomOutForcePoint,omitempty"`
					BottomOutForce        *float32  `gorm:"type:real"                        json:"bottomOutForce,omitempty"`
					SoundLevelID          *int      `gorm:"type:int"                         json:"soundLevelId,omitempty"`
					SoundLevel            *Type     `gorm:"foreignKey:SoundLevelID"          json:"soundLevel,omitempty"`
					SoundTypeID           *int      `gorm:"type:int"                         json:"soundTypeId,omitempty"`
					SoundType             *Type     `gorm:"foreignKey:SoundTypeID"           json:"soundType,omitempty"`
					TactilityTypeID       *int      `gorm:"type:int"                         json:"tactilityTypeId,omitempty"`
					TactilityType         *Type     `gorm:"foreignKey:TactilityTypeID"       json:"tactilityType,omitempty"`
					BumpPosition          *float32  `gorm:"type:real"                        json:"bumpPosition,omitempty"`
					BumpForce             *float32  `gorm:"type:real"                        json:"bumpForce,omitempty"`
					TactilityFeedbackID   *int      `gorm:"type:int"                         json:"tactilityFeedbackId,omitempty"`
					TactilityFeedback     *Type     `gorm:"foreignKey:TactilityFeedbackID"   json:"tactilityFeedback,omitempty"`
					FactoryLubed          bool      `gorm:"type:boolean;default:false"       json:"factoryLubed"`
					StemColorID           *int      `gorm:"type:int"                         json:"stemColorId,omitempty"`
					StemColor             *Type     `gorm:"foreignKey:StemColorID"           json:"stemColor,omitempty"`
					TopHousingColorID     *int      `gorm:"type:int"                         json:"topHousingColorId,omitempty"`
					TopHousingColor       *Type     `gorm:"foreignKey:TopHousingColorID"     json:"topHousingColor,omitempty"`
					BottomHousingColorID  *int      `gorm:"type:int"                         json:"bottomHousingColorId,omitempty"`
					BottomHousingColor    *Type     `gorm:"foreignKey:BottomHousingColorID"  json:"bottomHousingColor,omitempty"`
					PinConfigurationID    *int      `gorm:"type:int"                         json:"pinConfigurationId,omitempty"`
					PinConfiguration      *Type     `gorm:"foreignKey:PinConfigurationID"    json:"pinConfiguration,omitempty"`
					CreatedAt             time.Time `gorm:"type:timestamp;autoCreateTime"    json:"createdAt"`
					UpdatedAt             time.Time `gorm:"type:timestamp;autoUpdateTime"    json:"updatedAt"`
				}
				if err := tx.Migrator().CreateTable(&Details{}); err != nil {
					return err
				}

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
				if err := tx.Migrator().CreateTable(&Switch{}); err != nil {
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
				if err := tx.Migrator().DropTable("manufacturer"); err != nil {
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
