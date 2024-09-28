package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func CheckReviewRelavancy(review string) (float64, error) {
	data := map[string]interface{}{
		"review": review,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Error().Err(err).Msg("Error marshaling JSON")
		return 0.0, err
	}

	url := fmt.Sprintf("%s/review", viper.GetString("AI_SERVER"))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Error().Err(err).Msg("Error creating the request")
		return 0.0, err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("Error making the request")
		return 0.0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return 0.0, err
	}

	type Response struct {
		RelevanceScore float64 `json:"relevance_score"`
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error parsing JSON response:", err)
		return 0.0, err
	}

	return response.RelevanceScore, nil
}
