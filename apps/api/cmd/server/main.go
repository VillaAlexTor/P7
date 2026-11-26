package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/VillaAlexTor/p7/apps/api/internal/ai"
	"github.com/VillaAlexTor/p7/apps/api/internal/analysis"
	"github.com/VillaAlexTor/p7/apps/api/internal/domain"
)

const (
	maxDocumentCharacters = 125000
	analysisTimeout       = 5 * time.Minute
)

type readinessChecker interface {
	Health(context.Context) error
}

type requestGate struct {
	concurrent chan struct{}
	mu         sync.Mutex
	visitors   map[string]rateWindow
	limit      int
	window     time.Duration
}

type rateWindow struct {
	started time.Time
	count   int
}

func main() {
	port := strings.TrimSpace(os.Getenv("P7_API_PORT"))
	if port == "" {
		port = "8080"
	}

	server := &http.Server{Addr: ":" + port, Handler: newHandler(), ReadHeaderTimeout: 5 * time.Second, ReadTimeout: 15 * time.Second, WriteTimeout: 6 * time.Minute, IdleTimeout: 60 * time.Second}
	log.Printf("p7 API listening on http://localhost:%s", port)
	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}

func newHandler() http.Handler {
	classifier := ai.NewOllama(os.Getenv("P7_OLLAMA_URL"), os.Getenv("P7_OLLAMA_MODEL"))
	return newHandlerWithClassifier(classifier)
}

func newHandlerWithClassifier(classifier analysis.SemanticClassifier) http.Handler {
	gate := newRequestGate(envInt("P7_MAX_CONCURRENT_ANALYSES", 1), envInt("P7_ANALYSES_PER_MINUTE", 6))
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", health(classifier))
	mux.HandleFunc("POST /api/v1/analyze", analyze(classifier, gate))
	return withMiddleware(mux)
}

func health(classifier analysis.SemanticClassifier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if checker, ok := classifier.(readinessChecker); ok {
			ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
			defer cancel()
			if err := checker.Health(ctx); err != nil {
				writeJSON(w, http.StatusServiceUnavailable, map[string]any{"status": "unavailable", "service": "p7-api", "ai": map[string]string{"provider": classifier.Provider(), "model": classifier.Model()}})
				return
			}
		}
		writeJSON(w, http.StatusOK, map[string]any{"status": "ok", "service": "p7-api", "ai": map[string]string{"provider": classifier.Provider(), "model": classifier.Model()}})
	}
}

func analyze(classifier analysis.SemanticClassifier, gate *requestGate) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !gate.allow(clientIP(r)) || !gate.acquire() {
			writeJSON(w, http.StatusTooManyRequests, map[string]string{"error": localizedError(r.Header.Get("Accept-Language"), "busy")})
			return
		}
		defer gate.release()
		defer r.Body.Close()
		var request domain.AnalyzeRequest
		decoder := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1<<20))
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&request); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": requestError(r, "invalid")})
			return
		}
		if err := decoder.Decode(&struct{}{}); err != io.EOF {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": requestError(r, "invalid")})
			return
		}
		request.Text = strings.TrimSpace(request.Text)
		if len([]rune(request.Text)) < 80 {
			writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": localizedError(request.Locale, "short")})
			return
		}
		if len([]rune(request.Text)) > maxDocumentCharacters {
			writeJSON(w, http.StatusRequestEntityTooLarge, map[string]string{"error": localizedError(request.Locale, "large")})
			return
		}
		ctx, cancel := context.WithTimeout(r.Context(), analysisTimeout)
		defer cancel()
		result, err := analysis.AnalyzeWithClassifier(ctx, request, classifier)
		if err != nil {
			log.Printf("AI analysis failed: %v", err)
			writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": localizedError(request.Locale, "ai")})
			return
		}
		writeJSON(w, http.StatusOK, result)
	}
}

func newRequestGate(maxConcurrent, limit int) *requestGate {
	if maxConcurrent < 1 {
		maxConcurrent = 1
	}
	if limit < 1 {
		limit = 1
	}
	return &requestGate{concurrent: make(chan struct{}, maxConcurrent), visitors: make(map[string]rateWindow), limit: limit, window: time.Minute}
}

func (gate *requestGate) acquire() bool {
	select {
	case gate.concurrent <- struct{}{}:
		return true
	default:
		return false
	}
}

func (gate *requestGate) release() { <-gate.concurrent }

func (gate *requestGate) allow(ip string) bool {
	now := time.Now()
	gate.mu.Lock()
	defer gate.mu.Unlock()
	entry := gate.visitors[ip]
	if entry.started.IsZero() || now.Sub(entry.started) >= gate.window {
		gate.visitors[ip] = rateWindow{started: now, count: 1}
		return true
	}
	if entry.count >= gate.limit {
		return false
	}
	entry.count++
	gate.visitors[ip] = entry
	return true
}

func clientIP(r *http.Request) string {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil {
		return host
	}
	return r.RemoteAddr
}

func envInt(name string, fallback int) int {
	value, err := strconv.Atoi(strings.TrimSpace(os.Getenv(name)))
	if err != nil || value < 1 {
		return fallback
	}
	return value
}

func withMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if originAllowed(origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
		}
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Referrer-Policy", "no-referrer")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		started := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(started).Round(time.Millisecond))
	})
}

func originAllowed(origin string) bool {
	if origin == "" {
		return false
	}
	configured := strings.TrimSpace(os.Getenv("P7_ALLOWED_ORIGINS"))
	if configured == "" {
		configured = "http://localhost:3000"
	}
	for _, allowed := range strings.Split(configured, ",") {
		if strings.TrimSpace(allowed) == origin {
			return true
		}
	}
	return false
}

func requestError(r *http.Request, kind string) string {
	locale := r.Header.Get("Accept-Language")
	return localizedError(locale, kind)
}

func localizedError(locale, kind string) string {
	english := strings.HasPrefix(strings.ToLower(strings.TrimSpace(locale)), "en")
	if english {
		switch kind {
		case "short":
			return "Paste at least 80 characters from a policy to analyze it."
		case "large":
			return "The document exceeds the 125,000-character limit."
		case "busy":
			return "The local analyzer is busy. Wait a moment and try again."
		case "ai":
			return "The local AI is unavailable. Start Ollama and make sure the configured model is installed."
		default:
			return "The request is not valid JSON."
		}
	}
	switch kind {
	case "short":
		return "Pega al menos 80 caracteres de una política para analizarla."
	case "large":
		return "El documento supera el límite de 125 000 caracteres."
	case "busy":
		return "El analizador local está ocupado. Espera un momento y vuelve a intentarlo."
	case "ai":
		return "La IA local no está disponible. Inicia Ollama y confirma que el modelo configurado esté instalado."
	default:
		return "La solicitud no tiene un formato JSON válido."
	}
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}
