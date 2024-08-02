package migrations

import (
	"github.com/Bparsons0904/deadigations"
	"gorm.io/gorm"
)

type Type struct {
	ID        int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string `gorm:"type:varchar(50);not null" json:"name"`
	Code      string `gorm:"type:varchar(50);not null" json:"code"`
	Category  string `gorm:"index:idx_category;type:varchar(50);not null" json:"category"`
	CreatedAt string `gorm:"autoCreateTime" json:"createdAt"`
	UpdateAt  string `gorm:"autoUpdateTime" json:"updatedAt"`
}

func init() {
	deadigations.RegisterMigration(deadigations.Migration{
		ID:          "20240802130038",
		Description: "Adds the intial types table",
		Migrate: func(tx *gorm.DB) error {
			return tx.Migrator().CreateTable(&Type{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&Type{})
		},
	})
}
