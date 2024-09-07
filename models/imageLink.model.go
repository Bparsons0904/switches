package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

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

// Make a before create hook for the ImageLink model
func (imageLink *ImageLink) BeforeCreate(tx *gorm.DB) (err error) {
	if imageLink.LinkAddress == "" && imageLink.FullSize == "" {
		log.Error().Err(err).Msg("Link Address or Full Size is required")
		return errors.New("LinkAddress is required")
	}
	return nil
}
