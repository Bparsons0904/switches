package admin

import (
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func SeedDatabase(db *gorm.DB) error {
	log.Info().Msg("Seeding database")
	return nil
}
