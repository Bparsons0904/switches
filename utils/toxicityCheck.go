package utils

import (
	"context"
	"fmt"

	goaway "github.com/TwiN/go-away"
	"github.com/spf13/viper"
	commentanalyzer "google.golang.org/api/commentanalyzer/v1alpha1"
	"google.golang.org/api/option"
)

func PerformToxicityCheck(ctx context.Context, review string) (float64, float64, error) {
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

func FilterProfanity(review string) string {
	return goaway.Censor(review)
}

// type Comment struct {
// 	Text string `json:"text"`
// }
//
// type RequestBody struct {
// 	Comment             Comment                `json:"comment"`
// 	Languages           []string               `json:"languages"`
// 	RequestedAttributes map[string]interface{} `json:"requestedAttributes"`
// }
//
// type Score struct {
// 	Value float64 `json:"value"`
// 	Type  string  `json:"type"`
// }
//
// type SpanScore struct {
// 	Begin int   `json:"begin"`
// 	End   int   `json:"end"`
// 	Score Score `json:"score"`
// }
//
// type AttributeScore struct {
// 	SpanScores   []SpanScore `json:"spanScores"`
// 	SummaryScore Score       `json:"summaryScore"`
// }
//
// type ResponseBody struct {
// 	AttributeScores   map[string]AttributeScore `json:"attributeScores"`
// 	Languages         []string                  `json:"languages"`
// 	DetectedLanguages []string                  `json:"detectedLanguages"`
// }
