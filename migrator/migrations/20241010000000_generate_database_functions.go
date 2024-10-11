package migrations

import (
	"github.com/Bparsons0904/deadigations"
	"gorm.io/gorm"
)

func init() {
	deadigations.RegisterMigration(deadigations.Migration{
		ID:          "20241010203752",
		Description: "Add required extensions",
		Migrate: func(tx *gorm.DB) error {
			if err := setupDB(tx); err != nil {
				return err
			}
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return nil
		},
	})
}

func setupDB(tx *gorm.DB) error {
	if err := tx.Raw("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
		return err
	}

	if err := tx.Raw("CREATE EXTENSION IF NOT EXISTS \"pg_uuidv7\";").Error; err != nil {
		return err
	}

	return nil
}

