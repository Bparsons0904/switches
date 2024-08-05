package models

import "github.com/google/uuid"

type Producer struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v7();primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(50);not null"                       json:"name"`
	Alias     string    `gorm:"type:varchar(50);not null;uniqueIndex"           json:"alias"`
	SiteURL   string    `gorm:"type:varchar(50)"                                json:"siteURL,omitempty"`
	CreatedAt string    `gorm:"autoCreateTime"                                  json:"createdAt"`
	UpdateAt  string    `gorm:"autoUpdateTime"                                  json:"updatedAt"`
}
