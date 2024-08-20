package migrations

import (
	"time"

	"github.com/Bparsons0904/deadigations"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	deadigations.RegisterMigration(deadigations.Migration{
		ID:          "20240816075853",
		Description: "Add user table",
		Migrate: func(tx *gorm.DB) error {
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
			if err := tx.Migrator().CreateTable(&User{}); err != nil {
				return err
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			if err := tx.Migrator().DropTable("users"); err != nil {
				return err
			}

			return nil
		},
	})
}
