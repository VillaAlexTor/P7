package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/VillaAlexTor/p7/apps/api/internal/analysis"
	"github.com/VillaAlexTor/p7/apps/api/internal/domain"
)

type fakeClassifier struct{}

func (fakeClassifier) Provider() string { return "test" }
func (fakeClassifier) Model() string    { return "semantic-test" }
func (fakeClassifier) Classify(_ context.Context, clauses []domain.Clause, _ string) ([]analysis.SemanticCandidate, error) {
	for _, clause := range clauses {
		if strings.Contains(clause.Text, "precise location") {
			return []analysis.SemanticCandidate{{Category: "location", ClauseID: clause.ID, Quote: clause.Text}}, nil
		}
	}
	return []analysis.SemanticCandidate{}, nil
}

func testHandler() http.Handler { return newHandlerWithClassifier(fakeClassifier{}) }

func TestHealthAndSecurityHeaders(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	response := httptest.NewRecorder()
	testHandler().ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", response.Code)
	}
	if response.Header().Get("X-Content-Type-Options") != "nosniff" {
		t.Fatal("missing nosniff header")
	}
}

func TestAnalyzeContractInEnglish(t *testing.T) {
	payload := map[string]string{"title": "Demo", "profile": "journalist", "locale": "en", "text": "We collect precise location and contacts to provide the service. We may share personal data with third parties and service providers for personalized advertising. We retain information as long as necessary."}
	body, _ := json.Marshal(payload)
	request := httptest.NewRequest(http.MethodPost, "/api/v1/analyze", bytes.NewReader(body))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	testHandler().ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", response.Code, response.Body.String())
	}
	if !strings.Contains(response.Body.String(), `"locale":"en"`) || !strings.Contains(response.Body.String(), `"inputTreatedAsData":true`) {
		t.Fatalf("missing localized or safety contract: %s", response.Body.String())
	}
	if !strings.Contains(response.Body.String(), `"aiUsed":true`) || !strings.Contains(response.Body.String(), `"aiModel":"semantic-test"`) {
		t.Fatalf("missing AI execution metadata: %s", response.Body.String())
	}
}

func TestValidationAndCORS(t *testing.T) {
	t.Setenv("P7_ALLOWED_ORIGINS", "https://p7.example,http://localhost:3000")
	request := httptest.NewRequest(http.MethodPost, "/api/v1/analyze", strings.NewReader(`{"text":"short","locale":"en"}`))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Origin", "https://p7.example")
	response := httptest.NewRecorder()
	testHandler().ServeHTTP(response, request)
	if response.Code != http.StatusUnprocessableEntity {
		t.Fatalf("expected 422, got %d", response.Code)
	}
	if response.Header().Get("Access-Control-Allow-Origin") != "https://p7.example" {
		t.Fatal("configured origin was not allowed")
	}
	if !strings.Contains(response.Body.String(), "at least 80") {
		t.Fatalf("expected English validation error: %s", response.Body.String())
	}
}
