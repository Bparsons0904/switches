package models

import (
	"context"
	"switches/database"
	"switches/utils"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Rating struct {
	ID                  uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v7();primaryKey"                               json:"id"`
	Rating              int            `gorm:"type:int;not null"                                                             json:"rating"`
	Review              string         `gorm:"type:text"                                                                     json:"review"`
	OriginalReview      string         `gorm:"type:text"                                                                     json:"originalReview"`
	UserID              uuid.UUID      `gorm:"type:uuid;uniqueIndex:idx_user_switch"                                         json:"userId"`
	User                User           `gorm:"foreignKey:UserID;references:ID"                                               json:"user,omitempty"`
	SwitchID            uuid.UUID      `gorm:"type:uuid;uniqueIndex:idx_user_switch;index:idx_switch;index:idx_switch_admin" json:"switchId"`
	Switch              Switch         `gorm:"foreignKey:SwitchID;references:ID"                                             json:"switch,omitempty"`
	ToxicityScore       float64        `gorm:"type:float"                                                                    json:"toxicityScore"`
	ProfanityScore      float64        `gorm:"type:float"                                                                    json:"profanityScore"`
	SafeURLScore        float64        `gorm:"type:float"                                                                    json:"safeURLScore"`
	RelavanceScore      float64        `gorm:"type:float"                                                                    json:"relavanceScore"`
	AdminReviewRequired bool           `gorm:"type:bool;default:false;index:idx_switch_admin"                                json:"adminReviewRequired"`
	AdminReviewNotes    string         `gorm:"type:text"                                                                     json:"adminReviewNotes"`
	AdminReviewedByID   *uuid.UUID     `gorm:"type:uuid"                                                                     json:"adminReviewedById"`
	AdminReviewedBy     *User          `gorm:"foreignKey:AdminReviewedByID;references:ID"                                    json:"adminReviewedBy,omitempty"`
	CreatedAt           time.Time      `gorm:"autoCreateTime"                                                                json:"createdAt"`
	UpdatedAt           time.Time      `gorm:"autoUpdateTime"                                                                json:"updatedAt"`
	DeleteAt            gorm.DeletedAt `gorm:"index"                                                                         json:"deleteAt"`
}

func (r *Rating) AfterUpdate(tx *gorm.DB) (err error) {
	if r.Review != "" {
		go runQualityChecksAsync(*r)
		if r.OriginalReview == "" {
			r.OriginalReview = r.Review
		}
	}

	return
}

func runQualityChecksAsync(rating Rating) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if rating.Review == "" {
		log.Warn().Msg("Review is empty, skipping quality checks")
		return
	}

	var wg sync.WaitGroup
	wg.Add(3)

	adminReviewRequired := false
	var toxicityScore, profanityScore, urlSafetyScore, relavanceScore float64

	go func() {
		defer wg.Done()

		var err error
		toxicityScore, profanityScore, err = utils.PerformToxicityCheck(ctx, rating.Review)
		if err != nil {
			log.Error().Err(err).Msg("Failed to perform toxicity check")
			adminReviewRequired = true
			return
		}
		if toxicityScore > 0.7 {
			adminReviewRequired = true
		}
	}()

	go func() {
		defer wg.Done()
		var err error
		urlSafetyScore, err = utils.PerformURLSafetyCheck(ctx, rating.Review)
		if err != nil {
			log.Error().Err(err).Msg("Failed to perform URL safety check")
			adminReviewRequired = true
			return
		}
		if urlSafetyScore != 0.0 {
			adminReviewRequired = true
		}
	}()

	go func() {
		defer wg.Done()
		var err error
		relavanceScore, err = utils.CheckReviewRelavancy(rating.Review)
		log.Info().Any("Relavancy Score", relavanceScore).Msg("Relavancy Score")
		if err != nil {
			log.Error().Err(err).Msg("Failed to perform Relavancy check")
			adminReviewRequired = true
			return
		}

		if relavanceScore < 3.0 {
			adminReviewRequired = true
		}
	}()

	wg.Wait()
	if profanityScore > 0.5 {
		rating.Review = utils.FilterProfanity(rating.Review)
	}

	if err := database.DB.
		Session(&gorm.Session{SkipHooks: true}).
		Model(&Rating{}).
		Where("id = ?", rating.ID).
		Updates(map[string]interface{}{
			"toxicity_score":        toxicityScore,
			"profanity_score":       profanityScore,
			"safe_url_score":        urlSafetyScore,
			"relavance_score":       relavanceScore,
			"admin_review_required": adminReviewRequired,
			"review":                rating.Review,
		}).Error; err != nil {
		log.Error().Err(err).Msg("Failed to update rating with quality check results")
		return
	}

	if adminReviewRequired {
		_ = UpdateSwitchRating(rating.SwitchID, database.DB)
	}
}
