package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

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
