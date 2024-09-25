package models

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
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
	wg.Add(2)

	adminReviewRequired := false
	var toxicityScore, profanityScore, urlSafetyScore float64

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

	go func() {
		defer wg.Done()
		var err error
		urlSafetyScore, err = performURLSafetyCheck(ctx, rating.Review)
		if err != nil {
			log.Error().Err(err).Msg("Failed to perform URL safety check")
			adminReviewRequired = true
			return
		}
		if urlSafetyScore > 0.7 {
			adminReviewRequired = true
		}
	}()

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
			SafeURLScore:        urlSafetyScore,
			AdminReviewRequired: adminReviewRequired,
			Review:              rating.Review,
		}).Error; err != nil {
		log.Error().Err(err).Msg("Failed to update rating with quality check results")
	}
}

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

func performURLSafetyCheck(ctx context.Context, review string) (float64, error) {
	urls := extractURLs(review)
	if len(urls) == 0 {
		return 1.0, nil // All safe if no URLs
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

	addThreatsToDatabase(safeBrowsingResp.Matches)

	unsafeURLCount := len(safeBrowsingResp.Matches)
	safetyScore := 1.0 - (float64(unsafeURLCount) / float64(len(urls)))
	return safetyScore, nil
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

type ThreatEntryMetadata struct {
	Entries []Entry `json:"entries"`
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
