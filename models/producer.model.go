package models

import (
	"time"
)

type Producer struct {
	ID        int        `gorm:"type:int;primaryKey;autoIncrement"                            json:"id"`
	Name      string     `gorm:"type:varchar(50);not null"                                    json:"name"`
	Alias     string     `gorm:"type:varchar(50);not null;uniqueIndex"                        json:"alias"`
	SiteURL   string     `gorm:"type:varchar(50)"                                             json:"siteURL,omitempty"`
	Logo      *ImageLink `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"logo,omitempty"`
	CreatedAt time.Time  `gorm:"autoCreateTime"                                               json:"createdAt"`
	UpdateAt  time.Time  `gorm:"autoUpdateTime"                                               json:"updatedAt"`
}
