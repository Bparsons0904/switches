package models

import (
	"time"

	"github.com/google/uuid"
)

type Rating struct {
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v7();primaryKey"        json:"id"`
	Rating         int       `gorm:"type:int;not null"                                      json:"rating"`
	Review         string    `gorm:"type:text"                                              json:"review"`
	UserID         uuid.UUID `gorm:"type:uuid;uniqueIndex:idx_user_switch"                  json:"userId"`
	User           User      `gorm:"foreignKey:UserID;references:ID"                        json:"user,omitempty"`
	SwitchID       uuid.UUID `gorm:"type:uuid;uniqueIndex:idx_user_switch;index:idx_switch" json:"switchId"`
	Switch         Switch    `gorm:"foreignKey:SwitchID;references:ID"                      json:"switch,omitempty"`
	ToxicityScore  float64   `gorm:"type:float"                                             json:"toxicityScore"`
	SafeURLScore   float64   `gorm:"type:float"                                             json:"safeRRLScore"`
	RelavanceScore float64   `gorm:"type:float"                                             json:"relavanceScore"`
	CreatedAt      time.Time `gorm:"autoCreateTime"                                         json:"createdAt"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"                                         json:"updatedAt"`
}
