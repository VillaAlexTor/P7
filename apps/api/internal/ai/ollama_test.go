package ai

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/VillaAlexTor/p7/apps/api/internal/domain"
)

type transportFunc func(*http.Request) (*http.Response, error)

func (fn transportFunc) RoundTrip(request *http.Request) (*http.Response, error) { return fn(request) }

func TestOllamaUsesStructuredOutput(t *testing.T) {
	client := NewOllama("http://127.0.0.1:11434", "test-model")
	requests := 0
	client.client.Transport = transportFunc(func(r *http.Request) (*http.Response, error) {
		requests++
		if r.URL.Path != "/api/chat" {
			t.Fatalf("unexpected path %q", r.URL.Path)
		}
		var payload map[string]any
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			t.Fatal(err)
		}
		if payload["format"] == nil || payload["stream"] != false {
			t.Fatalf("expected non-streaming structured request: %#v", payload)
		}
		content := `{"message":{"role":"assistant","content":"{\"findings\":[]}"}}`
		if requests == 1 {
			content = `{"message":{"role":"assistant","content":"{\"findings\":[{\"category\":\"location\",\"clauseId\":\"clause-001\",\"quote\":\"We collect precise location.\"}]}"}}`
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(content)), Header: make(http.Header)}, nil
	})

	findings, err := client.Classify(context.Background(), []domain.Clause{{ID: "clause-001", Text: "We collect precise location."}}, "en")
	if err != nil {
		t.Fatal(err)
	}
	if len(findings) != 1 || findings[0].Category != "location" || client.Model() != "test-model" || requests != 5 {
		t.Fatalf("unexpected result: %#v", findings)
	}
}

func TestBatchClausesPreservesAllClauses(t *testing.T) {
	clauses := []domain.Clause{{ID: "1", Text: "12345"}, {ID: "2", Text: "67890"}, {ID: "3", Text: "abcde"}}
	batches := batchClauses(clauses, 8)
	if len(batches) != 3 {
		t.Fatalf("expected three batches, got %#v", batches)
	}
}
