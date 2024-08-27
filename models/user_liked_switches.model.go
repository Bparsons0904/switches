package models

import "github.com/google/uuid"

type UserLikedSwitches struct {
	UserID   uuid.UUID `gorm:"type:uuid;index;not null" json:"user_id"`
	User     User      `gorm:"foreignKey:UserID"        json:"user"`
	SwitchID uuid.UUID `gorm:"type:uuid;not null"       json:"switch_id"`
	Switch   Switch    `gorm:"foreignKey:SwitchID"      json:"switch"`
}
