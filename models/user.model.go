package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                uuid.UUID `gorm:"type:uuid;default:uuid_generate_v7();primaryKey" json:"id"`
	Sub               int       `gorm:"type:int;index"                                  json:"sub"`
	Email             string    `gorm:"type:varchar(255);unique;not null"               json:"email"`
	EmailVerified     bool      `gorm:"type:boolean;default:false"                      json:"emailVerified"`
	GivenName         string    `gorm:"type:varchar(255)"                               json:"givenName"`
	FamilyName        string    `gorm:"type:varchar(255)"                               json:"familyName"`
	Name              string    `gorm:"type:varchar(255)"                               json:"name"`
	PreferredUsername string    `gorm:"type:varchar(255)"                               json:"preferredUsername"`
	IsAdmin           bool      `gorm:"type:boolean;default:false"                      json:"isAdmin"`
	CreatedAt         time.Time `gorm:"autoCreateTime"                                  json:"createdAt"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime"                                  json:"updatedAt"`
}
