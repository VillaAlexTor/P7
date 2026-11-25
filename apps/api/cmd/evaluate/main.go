package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/VillaAlexTor/p7/apps/api/internal/ai"
	"github.com/VillaAlexTor/p7/apps/api/internal/domain"
)

type example struct {
	id, locale, text string
	expected         map[string]bool
}

type metric struct {
	TP int `json:"tp"`
	FP int `json:"fp"`
	FN int `json:"fn"`
}

type outcome struct {
	ID        string   `json:"id"`
	Expected  []string `json:"expected"`
	Predicted []string `json:"predicted"`
	Rejected  int      `json:"rejectedCitations"`
}

func main() {
	path := flag.String("data", "evaluation/clauses.tsv", "labeled TSV dataset")
	flag.Parse()
	examples, err := readExamples(*path)
	if err != nil {
		fatal(err)
	}
	client := ai.NewOllama(os.Getenv("P7_OLLAMA_URL"), os.Getenv("P7_OLLAMA_MODEL"))
	metrics := map[string]*metric{}
	for _, category := range []string{"location", "contacts", "third-parties", "advertising", "ai-training", "biometric", "retention"} {
		metrics[category] = &metric{}
	}
	results := make([]outcome, 0, len(examples))
	start := time.Now()
	for _, sample := range examples {
		clause := domain.Clause{ID: sample.id, Text: sample.text}
		ctx, cancel := context.WithTimeout(context.Background(), 150*time.Second)
		candidates, err := client.Classify(ctx, []domain.Clause{clause}, sample.locale)
		cancel()
		if err != nil {
			fatal(fmt.Errorf("%s: %w", sample.id, err))
		}
		predicted := map[string]bool{}
		rejected := 0
		for _, candidate := range candidates {
			if candidate.ClauseID != sample.id || candidate.Quote == "" || !strings.Contains(sample.text, candidate.Quote) || metrics[candidate.Category] == nil {
				rejected++
				continue
			}
			predicted[candidate.Category] = true
		}
		for category, counts := range metrics {
			switch {
			case predicted[category] && sample.expected[category]:
				counts.TP++
			case predicted[category]:
				counts.FP++
			case sample.expected[category]:
				counts.FN++
			}
		}
		results = append(results, outcome{ID: sample.id, Expected: keys(sample.expected), Predicted: keys(predicted), Rejected: rejected})
		fmt.Fprintf(os.Stderr, "%s expected=%v predicted=%v rejected=%d\n", sample.id, keys(sample.expected), keys(predicted), rejected)
	}
	total := metric{}
	for _, counts := range metrics {
		total.TP += counts.TP
		total.FP += counts.FP
		total.FN += counts.FN
	}
	report := struct {
		Model           string             `json:"model"`
		Cases           int                `json:"cases"`
		DurationSeconds int                `json:"durationSeconds"`
		Overall         metric             `json:"overall"`
		ByCategory      map[string]*metric `json:"byCategory"`
		Outcomes        []outcome          `json:"outcomes"`
	}{client.Model(), len(examples), int(time.Since(start).Seconds()), total, metrics, results}
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(report); err != nil {
		fatal(err)
	}
}

func readExamples(path string) ([]example, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.FieldsPerRecord = 4
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	if len(rows) < 2 || strings.Join(rows[0], ",") != "id,locale,text,expected" {
		return nil, fmt.Errorf("invalid dataset header")
	}
	result := make([]example, 0, len(rows)-1)
	for _, row := range rows[1:] {
		expected := map[string]bool{}
		for _, label := range strings.Split(row[3], ",") {
			if label = strings.TrimSpace(label); label != "" && label != "-" {
				expected[label] = true
			}
		}
		result = append(result, example{row[0], row[1], row[2], expected})
	}
	return result, nil
}

func keys(values map[string]bool) []string {
	result := make([]string, 0, len(values))
	for key := range values {
		result = append(result, key)
	}
	sort.Strings(result)
	return result
}

func fatal(err error) { fmt.Fprintln(os.Stderr, err); os.Exit(1) }
