package migrations

import (
	"time"

	"github.com/Bparsons0904/deadigations"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	deadigations.RegisterMigration(deadigations.Migration{
		ID:          "20240911081632",
		Description: "Add description of changes",
		Migrate: func(tx *gorm.DB) error {
			err := addRating(tx)
			if err != nil {
				return err
			}

			err = updatedSwitch(tx)
			return err
		},
		Rollback: func(tx *gorm.DB) error {
			err := tx.Migrator().DropTable("ratings")
			if err != nil {
				return err
			}

			type User struct {
				ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v7();primaryKey" json:"id"`
			}
			type Type struct {
				ID int `gorm:"primaryKey;autoIncrement" json:"id"`
			}
			type Producer struct {
				ID int `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
			}
			type ImageLink struct {
				ID          uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v7();primaryKey" json:"id"`
				LinkAddress string         `gorm:"type:varchar(255);default:''"                    json:"linkAddress"`
				AltText     string         `gorm:"type:varchar(255);default:''"                    json:"altText,omitempty"`
				Thumbnail   string         `gorm:"type:varchar(255);default:''"                    json:"thumbnail,omitempty"`
				FullSize    string         `gorm:"type:varchar(255);default:''"                    json:"fullSize,omitempty"`
				IsPrimary   bool           `gorm:"default:false"                                   json:"isPrimary"`
				OwnerID     uuid.UUID      `gorm:"type:uuid;indes:idx_type_id"                     json:"objectId"`
				OwnerType   string         `gorm:"type:varchar(50);index:idx_type_id"              json:"objectType"`
				CreatedAt   time.Time      `gorm:"autoCreateTime"                                  json:"createdAt"`
				UpdatedAt   time.Time      `gorm:"autoUpdateTime"                                  json:"updatedAt"`
				DeletedAt   gorm.DeletedAt `gorm:""                                                json:"deletedAt"`
				CreatedByID uuid.UUID      `gorm:"type:uuid"                                       json:"createdById"`
				CreatedBy   User           `gorm:"foreignKey:CreatedByID;references:ID"            json:"createdBy,omitempty"`
				UpdatedByID uuid.UUID      `gorm:"type:uuid"                                       json:"updatedById"`
				UpdatedBy   User           `gorm:"foreignKey:UpdatedByID;references:ID"            json:"updatedBy,omitempty"`
			}
			type Switch struct {
				ID               uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v7();primaryKey"                                         json:"id"`
				Name             string         `gorm:"type:varchar(50);not null;index:idx_name;uniqueIndex:idx_name_manufacturer_type"         json:"name"`
				ShortDescription string         `gorm:"type:varchar(255);not null"                                                              json:"shortDescription"`
				LongDescription  string         `gorm:"type:text;not null"                                                                      json:"longDescription"`
				ManufacturerID   *int           `gorm:"type:int;index;uniqueIndex:idx_name_manufacturer_type"                                   json:"manufacturerId,omitempty"`
				Manufacturer     *Producer      `gorm:"foreignKey:ManufacturerID"                                                               json:"manufacturer,omitempty"`
				BrandID          *int           `gorm:"type:int;index"                                                                          json:"brandId,omitempty"`
				Brand            *Producer      `gorm:"foreignKey:BrandID"                                                                      json:"brand,omitempty"`
				SwitchTypeID     int            `gorm:"type:int;not null;index;uniqueIndex:idx_name_manufacturer_type"                          json:"switchTypeId"`
				SwitchType       *Type          `gorm:"foreignKey:SwitchTypeID"                                                                 json:"switchType,omitempty"`
				ReleaseDate      *time.Time     `gorm:"type:date"                                                                               json:"releaseDate,omitempty"`
				Available        bool           `gorm:"type:boolean;default:true"                                                               json:"available"`
				PricePoint       int            `gorm:"type:int;not null"                                                                       json:"pricePoint"`
				SiteURL          string         `gorm:"type:varchar(255)"                                                                       json:"siteURL,omitempty"`
				CreatedAt        time.Time      `gorm:"autoCreateTime"                                                                          json:"createdAt"`
				UpdatedAt        time.Time      `gorm:"autoUpdateTime"                                                                          json:"updatedAt"`
				DeletedAt        gorm.DeletedAt `gorm:"index"                                                                                   json:"deletedAt"`
				CreatedByID      uuid.UUID      `gorm:"type:uuid"                                                                               json:"createdById"`
				CreatedBy        User           `gorm:"foreignKey:CreatedByID;references:ID"                                                    json:"createdBy,omitempty"`
				UpdatedByID      uuid.UUID      `gorm:"type:uuid"                                                                               json:"updatedById"`
				UpdatedBy        User           `gorm:"foreignKey:UpdatedByID;references:ID"                                                    json:"updatedBy,omitempty"`
				ImageLinks       []ImageLink    `gorm:"polymorphic:Owner;polymorphicValue:switch;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"imageLinks,omitempty"`
			}

			err = tx.Migrator().DropColumn(&Switch{}, "average_rating")
			if err != nil {
				return err
			}
			err = tx.Migrator().DropColumn(&Switch{}, "ratings_count")
			if err != nil {
				return err
			}

			err = tx.Migrator().AutoMigrate(&Switch{})
			return err
		},
	})
}

func addRating(tx *gorm.DB) error {
	type User struct {
		ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v7();primaryKey" json:"id"`
	}
	type Switch struct {
		ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v7();primaryKey" json:"id"`
	}
	type Rating struct {
		ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v7();primaryKey"        json:"id"`
		Rating      int       `gorm:"type:int;not null"                                      json:"rating"`
		Comment     string    `gorm:"type:text"                                              json:"comment"`
		UserID      uuid.UUID `gorm:"type:uuid;uniqueIndex:idx_user_switch"                  json:"userId"`
		User        User      `gorm:"foreignKey:UserID;references:ID"                        json:"user,omitempty"`
		SwitchID    uuid.UUID `gorm:"type:uuid;uniqueIndex:idx_user_switch;index:idx_switch" json:"switchId"`
		Switch      Switch    `gorm:"foreignKey:SwitchID;references:ID"                      json:"switch,omitempty"`
		CreatedAt   time.Time `gorm:"autoCreateTime"                                         json:"createdAt"`
		UpdatedAt   time.Time `gorm:"autoUpdateTime"                                         json:"updatedAt"`
		CreatedByID uuid.UUID `gorm:"type:uuid"                                              json:"createdById"`
		CreatedBy   User      `gorm:"foreignKey:CreatedByID;references:ID"                   json:"createdBy,omitempty"`
		UpdatedByID uuid.UUID `gorm:"type:uuid"                                              json:"updatedById"`
		UpdatedBy   User      `gorm:"foreignKey:UpdatedByID;references:ID"                   json:"updatedBy,omitempty"`
	}

	err := tx.Migrator().CreateTable(&Rating{})

	return err
}

func updatedSwitch(tx *gorm.DB) error {
	type User struct {
		ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v7();primaryKey" json:"id"`
	}
	type Rating struct {
		ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v7();primaryKey" json:"id"`
	}
	type Type struct {
		ID int `gorm:"primaryKey;autoIncrement" json:"id"`
	}
	type Producer struct {
		ID int `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
	}
	type ImageLink struct {
		ID          uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v7();primaryKey" json:"id"`
		LinkAddress string         `gorm:"type:varchar(255);default:''"                    json:"linkAddress"`
		AltText     string         `gorm:"type:varchar(255);default:''"                    json:"altText,omitempty"`
		Thumbnail   string         `gorm:"type:varchar(255);default:''"                    json:"thumbnail,omitempty"`
		FullSize    string         `gorm:"type:varchar(255);default:''"                    json:"fullSize,omitempty"`
		IsPrimary   bool           `gorm:"default:false"                                   json:"isPrimary"`
		OwnerID     uuid.UUID      `gorm:"type:uuid;indes:idx_type_id"                     json:"objectId"`
		OwnerType   string         `gorm:"type:varchar(50);index:idx_type_id"              json:"objectType"`
		CreatedAt   time.Time      `gorm:"autoCreateTime"                                  json:"createdAt"`
		UpdatedAt   time.Time      `gorm:"autoUpdateTime"                                  json:"updatedAt"`
		DeletedAt   gorm.DeletedAt `gorm:""                                                json:"deletedAt"`
		CreatedByID uuid.UUID      `gorm:"type:uuid"                                       json:"createdById"`
		CreatedBy   User           `gorm:"foreignKey:CreatedByID;references:ID"            json:"createdBy,omitempty"`
		UpdatedByID uuid.UUID      `gorm:"type:uuid"                                       json:"updatedById"`
		UpdatedBy   User           `gorm:"foreignKey:UpdatedByID;references:ID"            json:"updatedBy,omitempty"`
	}
	type Switch struct {
		ID               uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v7();primaryKey"                                         json:"id"`
		Name             string         `gorm:"type:varchar(50);not null;index:idx_name;uniqueIndex:idx_name_manufacturer_type"         json:"name"`
		ShortDescription string         `gorm:"type:varchar(255);not null"                                                              json:"shortDescription"`
		LongDescription  string         `gorm:"type:text;not null"                                                                      json:"longDescription"`
		ManufacturerID   *int           `gorm:"type:int;index;uniqueIndex:idx_name_manufacturer_type"                                   json:"manufacturerId,omitempty"`
		Manufacturer     *Producer      `gorm:"foreignKey:ManufacturerID"                                                               json:"manufacturer,omitempty"`
		BrandID          *int           `gorm:"type:int;index"                                                                          json:"brandId,omitempty"`
		Brand            *Producer      `gorm:"foreignKey:BrandID"                                                                      json:"brand,omitempty"`
		SwitchTypeID     int            `gorm:"type:int;not null;index;uniqueIndex:idx_name_manufacturer_type"                          json:"switchTypeId"`
		SwitchType       *Type          `gorm:"foreignKey:SwitchTypeID"                                                                 json:"switchType,omitempty"`
		ReleaseDate      *time.Time     `gorm:"type:date"                                                                               json:"releaseDate,omitempty"`
		Available        bool           `gorm:"type:boolean;default:true"                                                               json:"available"`
		PricePoint       int            `gorm:"type:int;not null"                                                                       json:"pricePoint"`
		SiteURL          string         `gorm:"type:varchar(255)"                                                                       json:"siteURL,omitempty"`
		CreatedAt        time.Time      `gorm:"autoCreateTime"                                                                          json:"createdAt"`
		UpdatedAt        time.Time      `gorm:"autoUpdateTime"                                                                          json:"updatedAt"`
		DeletedAt        gorm.DeletedAt `gorm:"index"                                                                                   json:"deletedAt"`
		CreatedByID      uuid.UUID      `gorm:"type:uuid"                                                                               json:"createdById"`
		CreatedBy        User           `gorm:"foreignKey:CreatedByID;references:ID"                                                    json:"createdBy,omitempty"`
		UpdatedByID      uuid.UUID      `gorm:"type:uuid"                                                                               json:"updatedById"`
		UpdatedBy        User           `gorm:"foreignKey:UpdatedByID;references:ID"                                                    json:"updatedBy,omitempty"`
		ImageLinks       []ImageLink    `gorm:"polymorphic:Owner;polymorphicValue:switch;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"imageLinks,omitempty"`
		Ratings          []Rating       `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"                             json:"ratings,omitempty"`
		AverageRating    float64        `gorm:"type:decimal(3,2);default:0"                                                             json:"averageRating"`
		RatingsCount     int            `gorm:"type:int;default:0"                                                                      json:"ratingsCount"`
	}

	err := tx.Migrator().AutoMigrate(&ImageLink{}, &Switch{})
	return err
}
