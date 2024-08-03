package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

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
