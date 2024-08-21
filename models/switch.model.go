package models

import (
	"time"

	"github.com/google/uuid"
)

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
