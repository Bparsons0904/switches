package migrations

import (
	"time"

	"github.com/Bparsons0904/deadigations"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	deadigations.RegisterMigration(deadigations.Migration{
		ID:          "20241010081225",
		Description: "Add tables",
		Migrate: func(tx *gorm.DB) error {
			err := tx.Transaction(func(tx *gorm.DB) error {
				type User struct {
					ID                uuid.UUID `gorm:"type:uuid;default:uuid_generate_v7();primaryKey" json:"id"`
					Sub               int       `gorm:"type:int;index"                                  json:"sub"`
					Email             string    `gorm:"type:varchar(255);unique;not null"               json:"email"`
					EmailVerified     bool      `gorm:"type:boolean;default:false"                      json:"emailVerified"`
					GivenName         string    `gorm:"type:varchar(255)"                               json:"givenName"`
					FamilyName        string    `gorm:"type:varchar(255)"                               json:"familyName"`
					Name              string    `gorm:"type:varchar(255)"                               json:"name"`
					PreferredUsername string    `gorm:"type:varchar(255)"                               json:"preferredUsername"`
					IsAdmin           bool      `gorm:"type:boolean;default:false"                      json:"isAdmin"`
					CreatedAt         time.Time `gorm:"autoCreateTime"                                  json:"createdAt"`
					UpdatedAt         time.Time `gorm:"autoUpdateTime"                                  json:"updatedAt"`
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

				type Producer struct {
					ID        int        `gorm:"type:int;primaryKey;autoIncrement"                            json:"id"`
					Name      string     `gorm:"type:varchar(50);not null"                                    json:"name"`
					Alias     string     `gorm:"type:varchar(50);not null;uniqueIndex"                        json:"alias"`
					SiteURL   string     `gorm:"type:varchar(50)"                                             json:"siteURL,omitempty"`
					Logo      *ImageLink `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"logo,omitempty"`
					CreatedAt time.Time  `gorm:"autoCreateTime"                                               json:"createdAt"`
					UpdateAt  time.Time  `gorm:"autoUpdateTime"                                               json:"updatedAt"`
				}

				type Type struct {
					ID        int       `gorm:"primaryKey;autoIncrement"              json:"id"`
					Name      string    `gorm:"type:varchar(50);not null"             json:"name"`
					Code      string    `gorm:"type:varchar(50);not null;uniqueIndex" json:"code"`
					Category  string    `gorm:"type:varchar(50);not null;index"       json:"category"`
					CreatedAt time.Time `gorm:"autoCreateTime"                        json:"createdAt"`
					UpdatedAt time.Time `gorm:"autoUpdateTime"                        json:"updatedAt"`
				}

				type Switch struct {
					ID               uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v7();primaryKey"                                         json:"id"`
					Name             string         `gorm:"type:varchar(50);not null;"                                                              json:"name"`
					ShortDescription string         `gorm:"type:varchar(255);not null"                                                              json:"shortDescription"`
					LongDescription  string         `gorm:"type:text;not null"                                                                      json:"longDescription"`
					ManufacturerID   *int           `gorm:"type:int;index;"                                                                         json:"manufacturerId,omitempty"`
					Manufacturer     *Producer      `gorm:"foreignKey:ManufacturerID"                                                               json:"manufacturer,omitempty"`
					BrandID          *int           `gorm:"type:int;index"                                                                          json:"brandId,omitempty"`
					Brand            *Producer      `gorm:"foreignKey:BrandID"                                                                      json:"brand,omitempty"`
					SwitchTypeID     int            `gorm:"type:int;not null;index"                                                                 json:"switchTypeId"`
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
					AverageRating    float64        `gorm:"type:float;default:0.0"                                                                  json:"averageRating,omitempty"`
					RatingsCount     int            `gorm:"type:int;default:0"                                                                      json:"ratingsCount,omitempty"`
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
					ActuationForce        *float32  `gorm:"type:real"                                       json:"actuationForce,omitempty"`
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
				type Rating struct {
					ID                  uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v7();primaryKey"                               json:"id"`
					Rating              int            `gorm:"type:int;not null"                                                             json:"rating"`
					Review              string         `gorm:"type:text"                                                                     json:"review"`
					OriginalReview      string         `gorm:"type:text"                                                                     json:"originalReview"`
					UserID              uuid.UUID      `gorm:"type:uuid;uniqueIndex:idx_user_switch"                                         json:"userId"`
					User                User           `gorm:"foreignKey:UserID;references:ID"                                               json:"user,omitempty"`
					SwitchID            uuid.UUID      `gorm:"type:uuid;uniqueIndex:idx_user_switch;index:idx_switch;index:idx_switch_admin" json:"switchId"`
					Switch              Switch         `gorm:"foreignKey:SwitchID;references:ID"                                             json:"switch,omitempty"`
					ToxicityScore       float64        `gorm:"type:float"                                                                    json:"toxicityScore"`
					ProfanityScore      float64        `gorm:"type:float"                                                                    json:"profanityScore"`
					SafeURLScore        float64        `gorm:"type:float"                                                                    json:"safeURLScore"`
					RelavanceScore      float64        `gorm:"type:float"                                                                    json:"relavanceScore"`
					AdminReviewRequired bool           `gorm:"type:bool;default:false;index:idx_switch_admin"                                json:"adminReviewRequired"`
					AdminReviewNotes    string         `gorm:"type:text"                                                                     json:"adminReviewNotes"`
					AdminReviewedByID   *uuid.UUID     `gorm:"type:uuid"                                                                     json:"adminReviewedById"`
					AdminReviewedBy     *User          `gorm:"foreignKey:AdminReviewedByID;references:ID"                                    json:"adminReviewedBy,omitempty"`
					CreatedAt           time.Time      `gorm:"autoCreateTime"                                                                json:"createdAt"`
					UpdatedAt           time.Time      `gorm:"autoUpdateTime"                                                                json:"updatedAt"`
					DeleteAt            gorm.DeletedAt `gorm:"index"                                                                         json:"deleteAt"`
				}

				type UserLikedSwitches struct {
					UserID   uuid.UUID `gorm:"type:uuid;index;not null" json:"user_id"`
					User     User      `gorm:"foreignKey:UserID"        json:"user"`
					SwitchID uuid.UUID `gorm:"type:uuid;not null"       json:"switch_id"`
					Switch   Switch    `gorm:"foreignKey:SwitchID"      json:"switch"`
				}

				type UserOwnedSwitches struct {
					UserID   uuid.UUID `gorm:"type:uuid;index;not null" json:"user_id"`
					User     User      `gorm:"foreignKey:UserID"        json:"user"`
					SwitchID uuid.UUID `gorm:"type:uuid;not null"       json:"switch_id"`
					Switch   Switch    `gorm:"foreignKey:SwitchID"      json:"switch"`
				}

				if err := tx.Migrator().CreateTable(&User{}); err != nil {
					return err
				}

				if err := tx.Migrator().CreateTable(&ImageLink{}); err != nil {
					return err
				}

				if err := tx.Migrator().CreateTable(&Type{}); err != nil {
					return err
				}

				if err := tx.Migrator().CreateTable(&Details{}); err != nil {
					return err
				}

				if err := tx.Migrator().CreateTable(&Producer{}); err != nil {
					return err
				}

				if err := tx.Migrator().CreateTable(&Switch{}); err != nil {
					return err
				}

				if err := tx.Migrator().CreateTable(&Rating{}); err != nil {
					return err
				}

				if err := tx.Migrator().CreateTable(&UserOwnedSwitches{}); err != nil {
					return err
				}

				if err := tx.Migrator().CreateTable(&UserLikedSwitches{}); err != nil {
					return err
				}

				if err := finalizeUser(tx); err != nil {
					return err
				}

				if err := finalizeSwitch(tx); err != nil {
					return err
				}
				return nil
			})

			return err
		},
		Rollback: func(tx *gorm.DB) error {
			err := tx.Transaction(func(tx *gorm.DB) error {
				if err := tx.Migrator().DropTable("ratings"); err != nil {
					return err
				}
				if err := tx.Migrator().DropTable("switches"); err != nil {
					return err
				}
				if err := tx.Migrator().DropTable("producers"); err != nil {
					return err
				}
				if err := tx.Migrator().DropTable("details"); err != nil {
					return err
				}
				if err := tx.Migrator().DropTable("types"); err != nil {
					return err
				}
				if err := tx.Migrator().DropTable("image_links"); err != nil {
					return err
				}
				if err := tx.Migrator().DropTable("user_owned_switches"); err != nil {
					return err
				}
				if err := tx.Migrator().DropTable("user_liked_switches"); err != nil {
					return err
				}
				if err := tx.Migrator().DropTable("users"); err != nil {
					return err
				}

				return nil
			})
			return err
		},
	})
}

func finalizeSwitch(tx *gorm.DB) error {
	type Producer struct {
		ID        int       `gorm:"type:int;primaryKey;autoIncrement"     json:"id"`
		Name      string    `gorm:"type:varchar(50);not null"             json:"name"`
		Alias     string    `gorm:"type:varchar(50);not null;uniqueIndex" json:"alias"`
		SiteURL   string    `gorm:"type:varchar(50)"                      json:"siteURL,omitempty"`
		CreatedAt time.Time `gorm:"autoCreateTime"                        json:"createdAt"`
		UpdateAt  time.Time `gorm:"autoUpdateTime"                        json:"updatedAt"`
	}

	type Type struct {
		ID        int       `gorm:"primaryKey;autoIncrement"              json:"id"`
		Name      string    `gorm:"type:varchar(50);not null"             json:"name"`
		Code      string    `gorm:"type:varchar(50);not null;uniqueIndex" json:"code"`
		Category  string    `gorm:"type:varchar(50);not null;index"       json:"category"`
		CreatedAt time.Time `gorm:"autoCreateTime"                        json:"createdAt"`
		UpdatedAt time.Time `gorm:"autoUpdateTime"                        json:"updatedAt"`
	}

	type User struct {
		ID                uuid.UUID `gorm:"type:uuid;default:uuid_generate_v7();primaryKey" json:"id"`
		Sub               int       `gorm:"type:int;index"                                  json:"sub"`
		Email             string    `gorm:"type:varchar(255);unique;not null"               json:"email"`
		EmailVerified     bool      `gorm:"type:boolean;default:false"                      json:"emailVerified"`
		GivenName         string    `gorm:"type:varchar(255)"                               json:"givenName"`
		FamilyName        string    `gorm:"type:varchar(255)"                               json:"familyName"`
		Name              string    `gorm:"type:varchar(255)"                               json:"name"`
		PreferredUsername string    `gorm:"type:varchar(255)"                               json:"preferredUsername"`
		IsAdmin           bool      `gorm:"type:boolean;default:false"                      json:"isAdmin"`
		CreatedAt         time.Time `gorm:"autoCreateTime"                                  json:"createdAt"`
		UpdatedAt         time.Time `gorm:"autoUpdateTime"                                  json:"updatedAt"`
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
		ActuationForce        *float32  `gorm:"type:real"                                       json:"actuationForce,omitempty"`
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
	type Rating struct {
		ID                  uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v7();primaryKey"                               json:"id"`
		Rating              int            `gorm:"type:int;not null"                                                             json:"rating"`
		Review              string         `gorm:"type:text"                                                                     json:"review"`
		OriginalReview      string         `gorm:"type:text"                                                                     json:"originalReview"`
		UserID              uuid.UUID      `gorm:"type:uuid;uniqueIndex:idx_user_switch"                                         json:"userId"`
		User                User           `gorm:"foreignKey:UserID;references:ID"                                               json:"user,omitempty"`
		SwitchID            uuid.UUID      `gorm:"type:uuid;uniqueIndex:idx_user_switch;index:idx_switch;index:idx_switch_admin" json:"switchId"`
		ToxicityScore       float64        `gorm:"type:float"                                                                    json:"toxicityScore"`
		ProfanityScore      float64        `gorm:"type:float"                                                                    json:"profanityScore"`
		SafeURLScore        float64        `gorm:"type:float"                                                                    json:"safeURLScore"`
		RelavanceScore      float64        `gorm:"type:float"                                                                    json:"relavanceScore"`
		AdminReviewRequired bool           `gorm:"type:bool;default:false;index:idx_switch_admin"                                json:"adminReviewRequired"`
		AdminReviewNotes    string         `gorm:"type:text"                                                                     json:"adminReviewNotes"`
		AdminReviewedByID   *uuid.UUID     `gorm:"type:uuid"                                                                     json:"adminReviewedById"`
		AdminReviewedBy     *User          `gorm:"foreignKey:AdminReviewedByID;references:ID"                                    json:"adminReviewedBy,omitempty"`
		CreatedAt           time.Time      `gorm:"autoCreateTime"                                                                json:"createdAt"`
		UpdatedAt           time.Time      `gorm:"autoUpdateTime"                                                                json:"updatedAt"`
		DeleteAt            gorm.DeletedAt `gorm:"index"                                                                         json:"deleteAt"`
	}

	type Switch struct {
		ID               uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v7();primaryKey"                                         json:"id"`
		Name             string         `gorm:"type:varchar(50);not null;"                                                              json:"name"`
		ShortDescription string         `gorm:"type:varchar(255);not null"                                                              json:"shortDescription"`
		LongDescription  string         `gorm:"type:text;not null"                                                                      json:"longDescription"`
		ManufacturerID   *int           `gorm:"type:int;index;"                                                                         json:"manufacturerId,omitempty"`
		Manufacturer     *Producer      `gorm:"foreignKey:ManufacturerID"                                                               json:"manufacturer,omitempty"`
		BrandID          *int           `gorm:"type:int;index"                                                                          json:"brandId,omitempty"`
		Brand            *Producer      `gorm:"foreignKey:BrandID"                                                                      json:"brand,omitempty"`
		SwitchTypeID     int            `gorm:"type:int;not null;index"                                                                 json:"switchTypeId"`
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
		Ratings          []Rating       `gorm:"foreignKey:SwitchID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"         json:"ratings,omitempty"`
		AverageRating    float64        `gorm:"type:float;default:0.0"                                                                  json:"averageRating,omitempty"`
		RatingsCount     int            `gorm:"type:int;default:0"                                                                      json:"ratingsCount,omitempty"`
		UserRating       *Rating        `gorm:"-"                                                                                       json:"userRating,omitempty"`
		Details          Details        `gorm:"foreignKey:SwitchID;references:ID"                                                       json:"details,omitempty"`
	}

	return tx.AutoMigrate(Switch{})
}

func finalizeUser(tx *gorm.DB) error { // {{{
	type Switch struct {
		ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v7();primaryKey" json:"id"`
	}

	type User struct {
		ID                uuid.UUID `gorm:"type:uuid;default:uuid_generate_v7();primaryKey" json:"id"`
		Sub               int       `gorm:"type:int;index"                                  json:"sub"`
		Email             string    `gorm:"type:varchar(255);unique;not null"               json:"email"`
		EmailVerified     bool      `gorm:"type:boolean;default:false"                      json:"emailVerified"`
		GivenName         string    `gorm:"type:varchar(255)"                               json:"givenName"`
		FamilyName        string    `gorm:"type:varchar(255)"                               json:"familyName"`
		Name              string    `gorm:"type:varchar(255)"                               json:"name"`
		PreferredUsername string    `gorm:"type:varchar(255)"                               json:"preferredUsername"`
		IsAdmin           bool      `gorm:"type:boolean;default:false"                      json:"isAdmin"`
		CreatedAt         time.Time `gorm:"autoCreateTime"                                  json:"createdAt"`
		UpdatedAt         time.Time `gorm:"autoUpdateTime"                                  json:"updatedAt"`
		OwnedSwitches     []*Switch `gorm:"many2many:user_owned_switches;"                  json:"ownedSwitches,omitempty"`
		LikedSwitches     []*Switch `gorm:"many2many:user_liked_switches;"                  json:"likedSwitches,omitempty"`
	}

	return tx.AutoMigrate(User{})
} // }}}
