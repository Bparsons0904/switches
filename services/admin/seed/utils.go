package seed

import (
	"fmt"
	"switches/models"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func deleteTables(db *gorm.DB) error {
	tables := []string{
		"types",
		"producers",
		"image_links",
		"switches",
		"ratings",
		"details",
	}

	return db.Transaction(func(tx *gorm.DB) error {
		// Disable all triggers
		if err := tx.Exec("SET session_replication_role = 'replica';").Error; err != nil {
			return fmt.Errorf("failed to disable triggers: %w", err)
		}

		for _, table := range tables {
			log.Info().Msgf("Deleting %s", table)
			if err := tx.Exec("TRUNCATE TABLE " + table + " RESTART IDENTITY CASCADE").Error; err != nil {
				log.Error().Err(err).Msgf("Error truncating table %s", table)
				return err
			}
		}

		// Re-enable all triggers
		if err := tx.Exec("SET session_replication_role = 'origin';").Error; err != nil {
			return fmt.Errorf("failed to re-enable triggers: %w", err)
		}

		return nil
	})
}

func ptr[T any](value T) *T {
	return &value
}

func processSwitches(tx *gorm.DB, switches []models.Switch, admin models.User) error {
	for _, clickyClack := range switches {
		clickyClack.CreatedByID = admin.ID
		clickyClack.UpdatedByID = admin.ID
		for i := range clickyClack.ImageLinks {
			clickyClack.ImageLinks[i].CreatedByID = admin.ID
			clickyClack.ImageLinks[i].UpdatedByID = admin.ID
		}
		if err := tx.Create(&clickyClack).Error; err != nil {
			return err
		}

		// Sleep for 5 sec to ensure unique timestamps
		// time.Sleep(5 * time.Second)
		time.Sleep(25 * time.Millisecond)
	}

	return nil
}

func parseDate(date string) *time.Time {
	t, _ := time.Parse("2006-01-02", date)
	return &t
}
