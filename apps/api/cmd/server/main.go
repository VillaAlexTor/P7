package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/VillaAlexTor/p7/apps/api/internal/ai"
	"github.com/VillaAlexTor/p7/apps/api/internal/analysis"
	"github.com/VillaAlexTor/p7/apps/api/internal/domain"
)

func main() {
	port := strings.TrimSpace(os.Getenv("P7_API_PORT"))
	if port == "" {
		port = "8080"
	}

	server := &http.Server{Addr: ":" + port, Handler: newHandler(), ReadHeaderTimeout: 5 * time.Second, ReadTimeout: 15 * time.Second, WriteTimeout: 3 * time.Minute, IdleTimeout: 60 * time.Second}
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
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", health(classifier))
	mux.HandleFunc("POST /api/v1/analyze", analyze(classifier))
	return withMiddleware(mux)
}

func health(classifier analysis.SemanticClassifier) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, map[string]any{"status": "ok", "service": "p7-api", "ai": map[string]string{"provider": classifier.Provider(), "model": classifier.Model()}})
	}
}

func analyze(classifier analysis.SemanticClassifier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var request domain.AnalyzeRequest
		decoder := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1<<20))
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&request); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": requestError(r, "invalid")})
			return
		}
		request.Text = strings.TrimSpace(request.Text)
		if len([]rune(request.Text)) < 80 {
			writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": localizedError(request.Locale, "short")})
			return
		}
		if len([]rune(request.Text)) > 200000 {
			writeJSON(w, http.StatusRequestEntityTooLarge, map[string]string{"error": localizedError(request.Locale, "large")})
			return
		}
		result, err := analysis.AnalyzeWithClassifier(r.Context(), request, classifier)
		if err != nil {
			log.Printf("AI analysis failed: %v", err)
			writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": localizedError(request.Locale, "ai")})
			return
		}
		writeJSON(w, http.StatusOK, result)
	}
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
			return "The document exceeds the 200,000-character limit."
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
		return "El documento supera el límite de 200 000 caracteres."
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
