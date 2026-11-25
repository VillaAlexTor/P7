package main

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/VillaAlexTor/p7/apps/api/internal/ai"
	"github.com/VillaAlexTor/p7/apps/api/internal/analysis"
	"github.com/VillaAlexTor/p7/apps/api/internal/domain"
)

type policy struct {
	ID, Title, Locale, URL, SHA256 string
	Characters                     int
	Expected                       []string
}

type manifest struct {
	Policies []policy `json:"policies"`
}

type counts struct {
	TP int `json:"tp"`
	FP int `json:"fp"`
	FN int `json:"fn"`
}

type finding struct {
	Category         string `json:"category"`
	ClauseID         string `json:"clauseId"`
	QuoteSHA256      string `json:"quoteSha256"`
	EvidenceVerified bool   `json:"evidenceVerified"`
}

type outcome struct {
	ID, URL, Locale, SourceSHA256 string
	Characters                    int
	Expected, Predicted           []string
	Findings                      []finding
	Meta                          domain.AnalysisMeta
	Score                         int
	Safety                        domain.SafetyReport
	DurationSeconds               int
}

func main() {
	manifestPath := flag.String("manifest", "evaluation/real_policies_round10.json", "frozen labels and source hashes")
	textDir := flag.String("text-dir", "../../.cache/real_policies", "local extracted policies")
	flag.Parse()
	data, err := os.ReadFile(*manifestPath)
	check(err)
	var input manifest
	check(json.Unmarshal(data, &input))
	client := ai.NewOllama(os.Getenv("P7_OLLAMA_URL"), os.Getenv("P7_OLLAMA_MODEL"))
	byCategory := map[string]*counts{}
	for _, category := range []string{"location", "contacts", "third-parties", "advertising", "ai-training", "biometric", "retention"} {
		byCategory[category] = &counts{}
	}
	results := make([]outcome, 0, len(input.Policies))
	started := time.Now()
	for _, item := range input.Policies {
		file, err := os.ReadFile(filepath.Join(*textDir, item.ID+".txt"))
		check(err)
		hash := fmt.Sprintf("%x", sha256.Sum256(file))
		if hash != item.SHA256 || len([]rune(string(file))) != item.Characters {
			check(fmt.Errorf("%s: source text differs from frozen manifest", item.ID))
		}
		fmt.Fprintln(os.Stderr, "analyzing", item.ID, item.Characters, "characters")
		oneStart := time.Now()
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Minute)
		response, err := analysis.AnalyzeWithClassifier(ctx, domain.AnalyzeRequest{Text: string(file), Title: item.Title, Locale: item.Locale, Profile: "general"}, client)
		cancel()
		check(err)
		prediction := map[string]bool{}
		compact := make([]finding, 0, len(response.Findings))
		for _, found := range response.Findings {
			prediction[found.Category] = true
			compact = append(compact, finding{found.Category, found.Evidence.ClauseID, fmt.Sprintf("%x", sha256.Sum256([]byte(found.Evidence.Quote))), found.Evidence.Verified})
		}
		expected := map[string]bool{}
		for _, category := range item.Expected {
			expected[category] = true
		}
		for category, tally := range byCategory {
			switch {
			case prediction[category] && expected[category]:
				tally.TP++
			case prediction[category]:
				tally.FP++
			case expected[category]:
				tally.FN++
			}
		}
		predicted := make([]string, 0, len(prediction))
		for category := range prediction {
			predicted = append(predicted, category)
		}
		sort.Strings(predicted)
		results = append(results, outcome{item.ID, item.URL, item.Locale, item.SHA256, item.Characters, item.Expected, predicted, compact, response.Meta, response.Score, response.Safety, int(time.Since(oneStart).Seconds())})
		fmt.Fprintln(os.Stderr, "finished", item.ID, "expected", item.Expected, "predicted", predicted)
	}
	total := counts{}
	for _, tally := range byCategory {
		total.TP += tally.TP
		total.FP += tally.FP
		total.FN += tally.FN
	}
	report := struct {
		Model, Method   string
		Documents       int
		DurationSeconds int
		Overall         counts
		ByCategory      map[string]*counts
		Outcomes        []outcome
	}{client.Model(), "full-document AnalyzeWithClassifier; one finding per category per document", len(results), int(time.Since(started).Seconds()), total, byCategory, results}
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	check(encoder.Encode(report))
}

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
