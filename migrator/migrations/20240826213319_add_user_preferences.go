package migrations

import (
	"time"

	"github.com/Bparsons0904/deadigations"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	deadigations.RegisterMigration(deadigations.Migration{
		ID:          "20240826213319",
		Description: "Add description of changes",
		Migrate: func(tx *gorm.DB) error {
			type Switch struct {
				ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v7();primaryKey" json:"id"`
			}

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
				OwnedSwitches     []*Switch `gorm:"many2many:user_owned_switches;"                  json:"ownedSwitches,omitempty"`
				LikedSwitches     []*Switch `gorm:"many2many:user_liked_switches;"                  json:"likedSwitches,omitempty"`
			}

			type UserOwnedSwitches struct {
				UserID   uuid.UUID `gorm:"type:uuid;index;not null" json:"user_id"`
				User     User      `gorm:"foreignKey:UserID"        json:"user"`
				SwitchID uuid.UUID `gorm:"type:uuid;not null"       json:"switch_id"`
				Switch   Switch    `gorm:"foreignKey:SwitchID"      json:"switch"`
			}
			if err := tx.Migrator().CreateTable(&UserOwnedSwitches{}); err != nil {
				return err
			}

			type UserLikedSwitches struct {
				UserID   uuid.UUID `gorm:"type:uuid;index;not null" json:"user_id"`
				User     User      `gorm:"foreignKey:UserID"        json:"user"`
				SwitchID uuid.UUID `gorm:"type:uuid;not null"       json:"switch_id"`
				Switch   Switch    `gorm:"foreignKey:SwitchID"      json:"switch"`
			}
			if err := tx.Migrator().CreateTable(&UserLikedSwitches{}); err != nil {
				return err
			}

			if err := tx.Migrator().AutoMigrate(&User{}); err != nil {
				return err
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
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

			if err := tx.AutoMigrate(&User{}); err != nil {
				return err
			}

			if err := tx.Migrator().DropTable("user_liked_switches"); err != nil {
				return err
			}
			if err := tx.Migrator().DropTable("user_owned_switches"); err != nil {
				return err
			}
			return nil
		},
	})
}
