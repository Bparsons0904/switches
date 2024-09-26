package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"switches/database"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type URLCheck struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v7();primaryKey" json:"id"`
	URL       string    `gorm:"type:text;uniqueIndex"                           json:"url"`
	IsSafe    bool      `gorm:"type:boolean"                                    json:"isSafe"`
	CheckedAt time.Time `gorm:"type:timestamp"                                  json:"checkedAt"`
}

type Client struct {
	ClientID      string `json:"clientId"`
	ClientVersion string `json:"clientVersion"`
}

type Threat struct {
	URL string `json:"url"`
}

type ThreatInfo struct {
	ThreatTypes      []string `json:"threatTypes"`
	PlatformTypes    []string `json:"platformTypes"`
	ThreatEntryTypes []string `json:"threatEntryTypes"`
	ThreatEntries    []Threat `json:"threatEntries"`
}

type SafeBrowsingRequest struct {
	Client     Client     `json:"client"`
	ThreatInfo ThreatInfo `json:"threatInfo"`
}

type Entry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ThreatEntryMetadata struct {
	Entries []Entry `json:"entries"`
}

type Match struct {
	ThreatType          string              `json:"threatType"`
	PlatformType        string              `json:"platformType"`
	ThreatEntryType     string              `json:"threatEntryType"`
	Threat              Threat              `json:"threat"`
	ThreatEntryMetadata ThreatEntryMetadata `json:"threatEntryMetadata"`
	CacheDuration       string              `json:"cacheDuration"`
}

type SafeBrowsingResponse struct {
	Matches []Match `json:"matches"`
}

func PerformURLSafetyCheck(ctx context.Context, review string) (float64, error) {
	urls := extractURLs(review)
	slog.Info("urls", "count", len(urls))
	if len(urls) == 0 {
		return 0.0, nil
	}

	apiKey := viper.GetString("GCP_API_KEY")
	apiURL := fmt.Sprintf(
		"https://safebrowsing.googleapis.com/v4/threatMatches:find?key=%s",
		apiKey,
	)

	req := &SafeBrowsingRequest{
		Client: Client{
			ClientID:      "waugze-dev",
			ClientVersion: "1.5.2",
		},
		ThreatInfo: ThreatInfo{
			ThreatTypes: []string{
				"MALWARE",
				"SOCIAL_ENGINEERING",
				"UNWANTED_SOFTWARE",
				"POTENTIALLY_HARMFUL_APPLICATION",
			},
			PlatformTypes: []string{
				"WINDOWS",
				"LINUX",
				"ANDROID",
				"OSX",
				"IOS",
				"ANY_PLATFORM",
			},
			ThreatEntryTypes: []string{"URL"},
			ThreatEntries:    make([]Threat, len(urls)),
		},
	}

	for i, u := range urls {
		req.ThreatInfo.ThreatEntries[i] = Threat{URL: u}
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		return 0.0, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(
		ctx,
		"POST",
		apiURL,
		strings.NewReader(string(jsonData)),
	)
	if err != nil {
		return 0.0, fmt.Errorf("failed to create request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return 0.0, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0.0, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	var safeBrowsingResp SafeBrowsingResponse
	if err := json.NewDecoder(resp.Body).Decode(&safeBrowsingResp); err != nil {
		return 0.0, fmt.Errorf("failed to decode response: %w", err)
	}

	if len(safeBrowsingResp.Matches) > 0 {
		addThreatsToDatabase(safeBrowsingResp.Matches)
		log.Info().Any("SafeBrowsingResponse Matches", safeBrowsingResp.Matches)
		return 1.0, nil
	}

	return 0.0, nil
}

func checkToIncludeURLCache(parsedURL *url.URL) bool {
	duration, err := database.GetJSONKeyDB[int64]("urlCache", parsedURL.String())
	if err != nil {
		return true
	}

	timeToExpire := time.Now().Add(time.Second * time.Duration(duration))
	return time.Now().After(timeToExpire)
}

func extractURLs(text string) []string {
	urlRegex := regexp.MustCompile(
		`(?i)\b((?:https?://|www\d{0,3}[.]|[a-z0-9.\-]+[.][a-z]{2,4}/)(?:[^\s()<>]+|\(([^\s()<>]+|(\([^\s()<>]+\)))*\))+(?:\(([^\s()<>]+|(\([^\s()<>]+\)))*\)|[^\s` + "`" + `!()\[\]{};:'\".,<>?«»""'']))`,
	)
	matches := urlRegex.FindAllString(text, -1)
	uniqueURLs := make(map[string]bool)
	for _, match := range matches {
		parsedURL, err := url.Parse(match)
		if err == nil {
			if checkToIncludeURLCache(parsedURL) {
				uniqueURLs[parsedURL.String()] = true
			}
		}
	}
	urls := make([]string, 0, len(uniqueURLs))
	for url := range uniqueURLs {
		urls = append(urls, url)
	}
	return urls
}

func addThreatsToDatabase(matches []Match) {
	uniqueThreats := make(map[string]Match)

	for _, match := range matches {
		threat := match.Threat.URL
		if existingMatch, exists := uniqueThreats[threat]; !exists ||
			match.CacheDuration > existingMatch.CacheDuration {
			uniqueThreats[threat] = match
		}
	}

	for threat, match := range uniqueThreats {
		cacheDuration := strings.TrimSuffix(match.CacheDuration, "s")
		cacheDurationSeconds, err := strconv.ParseFloat(cacheDuration, 64)
		if err != nil {
			log.Error().Err(err).Msg("Error parsing the duration")
			continue
		}

		cacheDurationMilliseconds := int64(cacheDurationSeconds * 1000)

		err = database.SetJSONKeyDB(
			"urlCache",
			threat,
			cacheDurationMilliseconds,
			time.Duration(cacheDurationSeconds)*time.Second)
		if err != nil {
			log.Error().Err(err).Msg("Error setting cache")
		}
	}
}
