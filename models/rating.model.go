package models

import (
	"context"
	"fmt"
	"switches/database"
	"sync"
	"time"

	goaway "github.com/TwiN/go-away"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	commentanalyzer "google.golang.org/api/commentanalyzer/v1alpha1"
	"google.golang.org/api/option"

	"gorm.io/gorm"
)

type Rating struct {
	ID                  uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v7();primaryKey"        json:"id"`
	Rating              int            `gorm:"type:int;not null"                                      json:"rating"`
	Review              string         `gorm:"type:text"                                              json:"review"`
	OriginalReview      string         `gorm:"type:text"                                              json:"originalReview"`
	UserID              uuid.UUID      `gorm:"type:uuid;uniqueIndex:idx_user_switch"                  json:"userId"`
	User                User           `gorm:"foreignKey:UserID;references:ID"                        json:"user,omitempty"`
	SwitchID            uuid.UUID      `gorm:"type:uuid;uniqueIndex:idx_user_switch;index:idx_switch" json:"switchId"`
	Switch              Switch         `gorm:"foreignKey:SwitchID;references:ID"                      json:"switch,omitempty"`
	ToxicityScore       float64        `gorm:"type:float"                                             json:"toxicityScore"`
	ProfanityScore      float64        `gorm:"type:float"                                             json:"profanityScore"`
	SafeURLScore        float64        `gorm:"type:float"                                             json:"safeURLScore"`
	RelavanceScore      float64        `gorm:"type:float"                                             json:"relavanceScore"`
	AdminReviewRequired bool           `gorm:"type:bool;default:false"                                json:"adminReviewRequired"`
	AdminReviewNotes    string         `gorm:"type:text"                                              json:"adminReviewNotes"`
	AdminReviewedByID   *uuid.UUID     `gorm:"type:uuid"                                              json:"adminReviewedById"`
	AdminReviewedBy     *User          `gorm:"foreignKey:AdminReviewedByID;references:ID"             json:"adminReviewedBy,omitempty"`
	CreatedAt           time.Time      `gorm:"autoCreateTime"                                         json:"createdAt"`
	UpdatedAt           time.Time      `gorm:"autoUpdateTime"                                         json:"updatedAt"`
	DeleteAt            gorm.DeletedAt `gorm:"index"                                                  json:"deleteAt"`
}

func (r *Rating) BeforeCreate(tx *gorm.DB) (err error) {
	r.OriginalReview = r.Review
	return
}

func (r *Rating) AfterUpdate(tx *gorm.DB) (err error) {
	go RunQualityChecksAsync(*r)
	return
}

func RunQualityChecksAsync(rating Rating) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if rating.Review == "" {
		log.Warn().Msg("Review is empty, skipping quality checks")
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)

	adminReviewRequired := false
	var toxicityScore, profanityScore float64

	go func() {
		defer wg.Done()

		var err error
		toxicityScore, profanityScore, err = performToxicityCheck(ctx, rating.Review)
		if err != nil {
			log.Error().Err(err).Msg("Failed to perform toxicity check")
			adminReviewRequired = true
			return
		}
		if toxicityScore > 0.7 {
			adminReviewRequired = true
		}
	}()

	// safeURLScore := performURLSafetyCheck(ctx, rating.Review)
	// relevanceScore := performRelevanceCheck(ctx, rating.Review)

	wg.Wait()
	if profanityScore > 0.5 {
		rating.Review = filterProfanity(rating.Review)
	}

	if err := database.DB.
		Session(&gorm.Session{SkipHooks: true}).
		Model(&Rating{}).
		Where("id = ?", rating.ID).
		Updates(Rating{
			ToxicityScore:       toxicityScore,
			ProfanityScore:      profanityScore,
			AdminReviewRequired: adminReviewRequired,
			Review:              rating.Review,
		}).Error; err != nil {
		log.Error().Err(err).Msg("Failed to update rating with quality check results")
	}
}

type Comment struct {
	Text string `json:"text"`
}

type RequestBody struct {
	Comment             Comment                `json:"comment"`
	Languages           []string               `json:"languages"`
	RequestedAttributes map[string]interface{} `json:"requestedAttributes"`
}

type Score struct {
	Value float64 `json:"value"`
	Type  string  `json:"type"`
}

type SpanScore struct {
	Begin int   `json:"begin"`
	End   int   `json:"end"`
	Score Score `json:"score"`
}

type AttributeScore struct {
	SpanScores   []SpanScore `json:"spanScores"`
	SummaryScore Score       `json:"summaryScore"`
}

type ResponseBody struct {
	AttributeScores   map[string]AttributeScore `json:"attributeScores"`
	Languages         []string                  `json:"languages"`
	DetectedLanguages []string                  `json:"detectedLanguages"`
}

func filterProfanity(review string) string {
	return goaway.Censor(review)
}

func performToxicityCheck(ctx context.Context, review string) (float64, float64, error) {
	log.Info().Msg("Performing toxicity analysis")
	apiKey := viper.GetString("GCP_API_KEY")
	service, err := commentanalyzer.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return 0, 0, fmt.Errorf("failed to create Comment Analyzer service: %w", err)
	}

	analyzeRequest := &commentanalyzer.AnalyzeCommentRequest{
		Comment: &commentanalyzer.TextEntry{
			Text: review,
		},
		RequestedAttributes: map[string]commentanalyzer.AttributeParameters{
			"TOXICITY":  {},
			"PROFANITY": {},
		},
		Languages: []string{"en"},
	}

	response, err := service.Comments.Analyze(analyzeRequest).Context(ctx).Do()
	if err != nil {
		return 0, 0, fmt.Errorf("failed to analyze comment: %w", err)
	}

	toxicityScore := response.AttributeScores["TOXICITY"].SummaryScore.Value
	profanityScore := response.AttributeScores["PROFANITY"].SummaryScore.Value

	return toxicityScore, profanityScore, nil
}

// this just might be the smoothest fucking switch I have ever owned. I love it. It's great and very happy I bought these bitches.
