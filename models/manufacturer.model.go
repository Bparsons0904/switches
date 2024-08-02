package models

import "github.com/google/uuid"

type Manufacturer struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid__v7();primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(50);not null" json:"name"`
	SiteURL   string    `gorm:"type:varchar(50);not null" json:"siteURL"`
	CreatedAt string    `gorm:"autoCreateTime" json:"createdAt"`
	UpdateAt  string    `gorm:"autoUpdateTime" json:"updatedAt"`
}
