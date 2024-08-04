package models

import "time"

type Type struct {
	ID        int       `gorm:"primaryKey;autoIncrement"              json:"id"`
	Name      string    `gorm:"type:varchar(50);not null"             json:"name"`
	Code      string    `gorm:"type:varchar(50);not null;uniqueIndex" json:"code"`
	Category  string    `gorm:"type:varchar(50);not null;index"       json:"category"`
	CreatedAt time.Time `gorm:"autoCreateTime"                        json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"                        json:"updatedAt"`
}
