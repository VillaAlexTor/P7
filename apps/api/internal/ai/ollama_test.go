package ai

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/VillaAlexTor/p7/apps/api/internal/domain"
)

func TestOllamaUsesStructuredOutput(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		_ = json.NewEncoder(w).Encode(map[string]any{"message": map[string]string{"role": "assistant", "content": `{"findings":[{"category":"location","clauseId":"clause-001","quote":"We collect precise location."}]}`}})
	}))
	defer server.Close()

	client := NewOllama(server.URL, "test-model")
	findings, err := client.Classify(context.Background(), []domain.Clause{{ID: "clause-001", Text: "We collect precise location."}}, "en")
	if err != nil {
		t.Fatal(err)
	}
	if len(findings) != 1 || findings[0].Category != "location" || client.Model() != "test-model" {
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
