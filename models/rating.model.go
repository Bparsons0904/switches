package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Rating struct {
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v7();primaryKey"        json:"id"`
	Rating         int       `gorm:"type:int;not null"                                      json:"rating"`
	Review         string    `gorm:"type:text"                                              json:"review"`
	UserID         uuid.UUID `gorm:"type:uuid;uniqueIndex:idx_user_switch"                  json:"userId"`
	User           User      `gorm:"foreignKey:UserID;references:ID"                        json:"user,omitempty"`
	SwitchID       uuid.UUID `gorm:"type:uuid;uniqueIndex:idx_user_switch;index:idx_switch" json:"switchId"`
	Switch         Switch    `gorm:"foreignKey:SwitchID;references:ID"                      json:"switch,omitempty"`
	ToxicityScore  float64   `gorm:"type:float"                                             json:"toxicityScore"`
	SafeURLScore   float64   `gorm:"type:float"                                             json:"safeURLScore"`
	RelavanceScore float64   `gorm:"type:float"                                             json:"relavanceScore"`
	CreatedAt      time.Time `gorm:"autoCreateTime"                                         json:"createdAt"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"                                         json:"updatedAt"`
}

func (r *Rating) AfterSave(tx *gorm.DB) (err error) {
	go ScoreToxicity(*r)
	// Check for any URLs that might not be safe
	// Check that the review is relavent
	return
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

func ScoreToxicity(rating Rating) {
	url := fmt.Sprintf(
		"https://commentanalyzer.googleapis.com/v1alpha1/comments:analyze?key=%s",
		viper.GetString("GCP_API_KEY"),
	)

	// Prepare the request body
	requestBody := RequestBody{
		Comment: Comment{
			Text: rating.Review,
		},
		Languages: []string{"en"},
		RequestedAttributes: map[string]interface{}{
			"TOXICITY": map[string]interface{}{},
		},
	}

	// Marshal the request body to JSON
	body, err := json.Marshal(requestBody)
	if err != nil {
		log.Error().Err(err).Msg("Error marshaling request body")
		return
	}

	log.Info().Msgf("Request body: %s", string(body))

	// Create a new request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		log.Error().Err(err).Msg("Error creating request")
		return
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Create an HTTP client
	client := &http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("Error sending request")
		return
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Msg("Error reading response")
		return
	}

	// Handle the response
	log.Info().Msgf("Response Status: %s", resp.Status)
	log.Info().Msgf("Response Body: %s", string(respBody))

	var responseBody ResponseBody
	err = json.Unmarshal(respBody, &responseBody)
	if err != nil {
		log.Error().Err(err).Msg("Error unmarshaling response")
		return
	}

	log.Info().Msgf("Response Body: %v", responseBody.AttributeScores)

	if toxicityScore, ok := responseBody.AttributeScores["TOXICITY"]; ok {
		log.Info().Msgf("Toxicity score: %f", toxicityScore.SummaryScore.Value)
	} else {
		log.Info().Msg("Toxicity score not found in response")
	}

	log.Info().Any("rating", rating.ID).Msg("Toxicity score calculated")
}

// This is a switch that I really hate. It's really scratchy and not smooth at all. Then it just sounds horrible. I wouldn't recommend this switch at all.
