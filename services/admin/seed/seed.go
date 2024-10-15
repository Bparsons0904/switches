package seed

import (
	"switches/models"
	"switches/utils"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func SeedDatabase(db *gorm.DB) error {
	timer := utils.StartTimer("Seeding Database")
	defer timer.LogTotalTime()

	if err := deleteTables(db); err != nil {
		return err
	}

	timer.LogTime("Deleted tables")
	tx := db.Begin()

	if err := seedTypes(tx); err != nil {
		tx.Rollback()
		return err
	}
	timer.LogTime("Seed Types")

	if err := seedProducers(tx); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	timer.LogTime("Seed Producers")

	tx = db.Begin()
	var admin models.User
	if err := tx.First(&admin, "is_admin = true").Error; err != nil {
		log.Error().Err(err).Msg("Error finding admin user")
	}

	if err := seedCherry(tx, admin); err != nil {
		tx.Rollback()
		return err
	}

	if err := seedGateron(tx, admin); err != nil {
		tx.Rollback()
		return err
	}

	if err := seedKailh(tx, admin); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
