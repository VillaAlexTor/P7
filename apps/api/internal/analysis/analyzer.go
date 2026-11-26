package analysis

import (
	"context"
	"crypto/sha256"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"unicode"

	"github.com/VillaAlexTor/p7/apps/api/internal/domain"
)

type localized struct{ ES, EN string }

type SemanticCandidate struct {
	Category string `json:"category"`
	ClauseID string `json:"clauseId"`
	Quote    string `json:"quote"`
}

type SemanticClassifier interface {
	Classify(context.Context, []domain.Clause, string) ([]SemanticCandidate, error)
	Provider() string
	Model() string
}

func (value localized) text(locale string) string {
	if locale == "en" {
		return value.EN
	}
	return value.ES
}

type rule struct {
	category                  string
	title                     localized
	severity                  string
	interpretation            localized
	baseWeight                int
	dataLabel, impact, action localized
}

type gapCheck struct {
	control            localized
	present, ambiguous []string
	note               localized
}

var rules = []rule{
	{category: "location", title: localized{"Recopilación de ubicación", "Location collection"}, severity: "high", interpretation: localized{"El documento declara que puede recopilar datos de ubicación o movimiento.", "The document states that it may collect location or movement data."}, baseWeight: 18, dataLabel: localized{"Ubicación", "Location"}, impact: localized{"La ubicación frecuente puede permitir inferir rutinas, lugares habituales y patrones de desplazamiento.", "Frequent location data may reveal routines, regularly visited places, and movement patterns."}, action: localized{"Restringe la ubicación a «solo mientras usas la aplicación» o desactívala si no es necesaria.", "Restrict location access to “only while using the app,” or disable it when unnecessary."}},
	{category: "contacts", title: localized{"Acceso a contactos", "Contacts access"}, severity: "high", interpretation: localized{"El servicio contempla acceder o procesar información de la libreta de contactos.", "The service may access or process address-book information."}, baseWeight: 16, dataLabel: localized{"Contactos", "Contacts"}, impact: localized{"La libreta de contactos también contiene información de personas que quizá no usan el servicio.", "An address book also contains information about people who may not use the service."}, action: localized{"Deniega el permiso de contactos salvo que una función esencial lo requiera.", "Deny contacts access unless a core feature requires it."}},
	{category: "third-parties", title: localized{"Transferencia a terceros", "Third-party sharing"}, severity: "high", interpretation: localized{"La política permite compartir datos con terceros, socios o proveedores.", "The policy allows data sharing with third parties, partners, or service providers."}, baseWeight: 20, dataLabel: localized{"Datos personales", "Personal data"}, impact: localized{"Más destinatarios aumentan la superficie de exposición y dificultan seguir el recorrido de los datos.", "More recipients increase exposure and make the data trail harder to follow."}, action: localized{"Revisa las opciones para limitar publicidad, analítica y transferencias a socios.", "Review settings that limit advertising, analytics, and partner sharing."}},
	{category: "advertising", title: localized{"Publicidad personalizada", "Personalized advertising"}, severity: "medium", interpretation: localized{"El documento menciona perfiles, anuncios personalizados o seguimiento con fines publicitarios.", "The document mentions profiling, personalized ads, or advertising-related tracking."}, baseWeight: 12, dataLabel: localized{"Actividad", "Activity"}, impact: localized{"La actividad puede combinarse para inferir intereses y personalizar contenido o anuncios.", "Activity may be combined to infer interests and personalize content or advertising."}, action: localized{"Desactiva la personalización de anuncios y limita el seguimiento entre aplicaciones.", "Turn off ad personalization and limit cross-app tracking."}},
	{category: "ai-training", title: localized{"Uso para entrenamiento de IA", "AI training use"}, severity: "high", interpretation: localized{"El texto contempla usar contenido o datos para entrenar o mejorar modelos de inteligencia artificial.", "The text allows content or data to be used to train or improve AI models."}, baseWeight: 18, dataLabel: localized{"Contenido", "Content"}, impact: localized{"El contenido podría incorporarse a procesos de entrenamiento según los controles y condiciones del proveedor.", "Content could enter model-training processes depending on the provider's controls and terms."}, action: localized{"Busca una opción de exclusión del entrenamiento y evita enviar contenido sensible.", "Look for an AI-training opt-out and avoid submitting sensitive content."}},
	{category: "biometric", title: localized{"Datos biométricos", "Biometric data"}, severity: "critical", interpretation: localized{"La política menciona identificadores biométricos o rasgos físicos usados para reconocer a una persona.", "The policy mentions biometric identifiers or physical traits used to recognize a person."}, baseWeight: 24, dataLabel: localized{"Biometría", "Biometrics"}, impact: localized{"Los identificadores biométricos son difíciles de reemplazar si se exponen o usan fuera del contexto esperado.", "Biometric identifiers are difficult to replace if exposed or used outside the expected context."}, action: localized{"Evita habilitar identificación biométrica si existe un método alternativo adecuado.", "Avoid biometric identification when a suitable alternative is available."}},
	{category: "retention", title: localized{"Conservación prolongada", "Broad retention period"}, severity: "medium", interpretation: localized{"El documento permite conservar información durante periodos amplios o mientras sea necesaria.", "The document allows information to be retained for broad periods or as long as necessary."}, baseWeight: 10, dataLabel: localized{"Datos conservados", "Retained data"}, impact: localized{"Una conservación amplia mantiene los datos expuestos durante más tiempo y puede dificultar saber cuándo se eliminan.", "Broad retention keeps data exposed longer and may make deletion timing unclear."}, action: localized{"Solicita la eliminación de datos y confirma si las copias de respaldo siguen otro plazo.", "Request data deletion and confirm whether backup copies follow a different schedule."}},
}

var gapChecks = []gapCheck{
	{localized{"Datos recopilados", "Data collected"}, []string{"recopilamos", "collect", "información que recopil"}, nil, localized{"El documento debería enumerar las categorías de datos.", "The document should list the categories of data it collects."}},
	{localized{"Finalidad del tratamiento", "Purpose of processing"}, []string{"para proporcionar", "utilizamos", "we use", "purposes", "to provide"}, nil, localized{"El documento debería explicar para qué usa los datos.", "The document should explain why it uses the data."}},
	{localized{"Tiempo de conservación", "Retention period"}, []string{"conservaremos durante", "retention period", "retain for"}, []string{"mientras sea necesario", "as long as necessary"}, localized{"Un plazo concreto ayuda a entender durante cuánto tiempo permanecen los datos.", "A specific period helps people understand how long their data remains."}},
	{localized{"Identidad de terceros", "Identity of third parties"}, []string{"lista de proveedores", "third parties include", "proveedores son"}, []string{"terceros", "socios", "partners", "third parties"}, localized{"Nombrar o categorizar destinatarios reduce la ambigüedad.", "Naming or categorizing recipients reduces ambiguity."}},
	{localized{"Eliminación de cuenta", "Account deletion"}, []string{"eliminar tu cuenta", "delete your account", "account deletion"}, nil, localized{"El usuario debería saber cómo cerrar la cuenta y borrar sus datos.", "People should know how to close an account and delete their data."}},
	{localized{"Copias de respaldo", "Backup copies"}, []string{"copias de respaldo", "backups", "backup copies"}, nil, localized{"La política debería indicar qué ocurre con las copias tras una eliminación.", "The policy should explain what happens to backups after deletion."}},
	{localized{"Entrenamiento de IA", "AI training"}, []string{"entrenar modelos", "train our ai", "model training", "train artificial intelligence"}, nil, localized{"La política debería aclarar si el contenido se usa para entrenar modelos.", "The policy should clarify whether content is used to train models."}},
}

var profileMultipliers = map[string]map[string]float64{
	"general": {}, "student": {"advertising": 1.25, "ai-training": 1.2},
	"journalist": {"location": 1.65, "contacts": 1.35, "third-parties": 1.25},
	"parent":     {"location": 1.35, "contacts": 1.35, "biometric": 1.5},
	"activist":   {"location": 1.7, "contacts": 1.5, "third-parties": 1.35, "biometric": 1.35},
}

var instructionPatterns = []string{"ignore previous instructions", "ignore all instructions", "system prompt", "developer message", "ignora las instrucciones anteriores", "ignora todas las instrucciones", "mensaje del sistema"}

func AnalyzeWithClassifier(ctx context.Context, req domain.AnalyzeRequest, classifier SemanticClassifier) (domain.AnalysisResponse, error) {
	text := normalize(req.Text)
	locale := normalizeLocale(req.Locale, text)
	clauses := segment(text)
	profile := normalizeProfile(req.Profile)
	candidates, err := classifier.Classify(ctx, clauses, locale)
	if err != nil {
		return domain.AnalysisResponse{}, fmt.Errorf("semantic classification failed: %w", err)
	}
	findings, rejected := findSemantic(clauses, candidates, profile, locale)
	return buildResponse(req, text, locale, profile, clauses, findings, "p7-hybrid/0.3", true, classifier.Provider(), classifier.Model(), len(candidates), rejected), nil
}

func buildResponse(req domain.AnalyzeRequest, text, locale, profile string, clauses []domain.Clause, findings []domain.Finding, engine string, aiUsed bool, provider, model string, candidateCount, rejected int) domain.AnalysisResponse {
	gaps := detectGaps(text, locale)
	score := calculateScore(findings, gaps)
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(text)))
	nodes, edges := buildFlow(findings, locale)
	impacts, actions := buildGuidance(findings, gaps, locale)
	title := strings.TrimSpace(req.Title)
	if title == "" {
		title = localized{"Política sin título", "Untitled policy"}.text(locale)
	}
	verified := verifiedCount(findings)
	return domain.AnalysisResponse{
		Meta:  domain.AnalysisMeta{Title: title, Profile: profile, DocumentHash: hash, ClauseCount: len(clauses), Language: detectLanguage(text), Engine: engine, EvidenceChecked: verified, Locale: locale, AIUsed: aiUsed, AIProvider: provider, AIModel: model, AICandidates: candidateCount},
		Score: score, Level: scoreLevel(score), Summary: summary(score, len(findings), gaps, locale), Clauses: clauses,
		Findings: findings, Gaps: gaps, Nodes: nodes, Edges: edges, Impacts: impacts, Actions: actions,
		Safety: domain.SafetyReport{InputTreatedAsData: true, IgnoredInstructionPatterns: countInstructionPatterns(text), CitationsVerified: verified, CitationsRejected: rejected},
	}
}

func normalize(input string) string {
	input = strings.ReplaceAll(strings.ReplaceAll(input, "\r\n", "\n"), "\r", "\n")
	lines := strings.Split(input, "\n")
	for i := range lines {
		lines[i] = strings.Join(strings.Fields(lines[i]), " ")
	}
	return strings.TrimSpace(strings.Join(lines, "\n"))
}

func normalizeProfile(profile string) string {
	profile = strings.ToLower(strings.TrimSpace(profile))
	if _, ok := profileMultipliers[profile]; !ok {
		return "general"
	}
	return profile
}

func normalizeLocale(locale, text string) string {
	locale = strings.ToLower(strings.TrimSpace(locale))
	if strings.HasPrefix(locale, "en") {
		return "en"
	}
	if strings.HasPrefix(locale, "es") {
		return "es"
	}
	return detectLanguage(text)
}

func segment(text string) []domain.Clause {
	paragraphs := regexp.MustCompile(`\n\s*\n`).Split(text, -1)
	clauses := make([]domain.Clause, 0, len(paragraphs))
	for _, paragraph := range paragraphs {
		paragraph = strings.TrimSpace(paragraph)
		if paragraph == "" {
			continue
		}
		if len([]rune(paragraph)) > 700 {
			for _, part := range regexp.MustCompile(`(?m)(?:\.\s+|;\s+)`).Split(paragraph, -1) {
				part = strings.TrimSpace(part)
				if len([]rune(part)) >= 20 {
					clauses = append(clauses, domain.Clause{ID: fmt.Sprintf("clause-%03d", len(clauses)+1), Text: part})
				}
			}
		} else if len([]rune(paragraph)) >= 10 {
			clauses = append(clauses, domain.Clause{ID: fmt.Sprintf("clause-%03d", len(clauses)+1), Text: paragraph})
		}
	}
	if len(clauses) == 0 && text != "" {
		clauses = append(clauses, domain.Clause{ID: "clause-001", Text: text})
	}
	return clauses
}

func findSemantic(clauses []domain.Clause, candidates []SemanticCandidate, profile, locale string) ([]domain.Finding, int) {
	result := make([]domain.Finding, 0, len(candidates))
	seen := map[string]bool{}
	rejected := 0
	for _, proposed := range candidates {
		candidate, ok := ruleForCategory(strings.ToLower(strings.TrimSpace(proposed.Category)))
		evidence := domain.Evidence{ClauseID: strings.TrimSpace(proposed.ClauseID), Quote: strings.TrimSpace(proposed.Quote)}
		if !ok || seen[candidate.category] || !verifyEvidence(clauses, evidence) {
			rejected++
			continue
		}
		evidence.Verified = true
		multiplier := profileMultipliers[profile][candidate.category]
		if multiplier == 0 {
			multiplier = 1
		}
		result = append(result, domain.Finding{ID: fmt.Sprintf("finding-%03d", len(result)+1), Category: candidate.category, Title: candidate.title.text(locale), Severity: candidate.severity, Interpretation: candidate.interpretation.text(locale), Evidence: evidence, Weight: int(float64(candidate.baseWeight)*multiplier + 0.5)})
		seen[candidate.category] = true
	}
	sort.SliceStable(result, func(i, j int) bool { return result[i].Weight > result[j].Weight })
	return result, rejected
}

func ruleForCategory(category string) (rule, bool) {
	for _, candidate := range rules {
		if candidate.category == category {
			return candidate, true
		}
	}
	return rule{}, false
}

func verifyEvidence(clauses []domain.Clause, evidence domain.Evidence) bool {
	if strings.TrimSpace(evidence.Quote) == "" || strings.TrimSpace(evidence.ClauseID) == "" {
		return false
	}
	for _, clause := range clauses {
		if clause.ID == evidence.ClauseID {
			// PDF extraction can insert line breaks within a sentence. Preserve
			// every word and punctuation mark while allowing whitespace to vary.
			return strings.Contains(strings.Join(strings.Fields(clause.Text), " "), strings.Join(strings.Fields(evidence.Quote), " "))
		}
	}
	return false
}

func detectGaps(text, locale string) []domain.Gap {
	lower := strings.ToLower(text)
	result := make([]domain.Gap, 0, len(gapChecks))
	for _, check := range gapChecks {
		status := "missing"
		if containsAny(lower, check.present) {
			status = "specified"
		} else if containsAny(lower, check.ambiguous) {
			status = "ambiguous"
		}
		result = append(result, domain.Gap{Control: check.control.text(locale), Status: status, Note: check.note.text(locale)})
	}
	return result
}

func calculateScore(findings []domain.Finding, gaps []domain.Gap) int {
	total := 0
	for _, finding := range findings {
		total += finding.Weight
	}
	for _, gap := range gaps {
		if gap.Status == "missing" {
			total += 4
		}
		if gap.Status == "ambiguous" {
			total += 2
		}
	}
	if total > 100 {
		return 100
	}
	return total
}

func buildFlow(findings []domain.Finding, locale string) ([]domain.FlowNode, []domain.FlowEdge) {
	nodes := []domain.FlowNode{{ID: "you", Label: localized{"Tú", "You"}.text(locale), Kind: "person"}, {ID: "service", Label: localized{"Servicio", "Service"}.text(locale), Kind: "service"}}
	edges := []domain.FlowEdge{{From: "you", To: "service", DataLabel: localized{"Datos declarados", "Declared data"}.text(locale)}}
	if hasCategory(findings, "third-parties") {
		nodes = append(nodes, domain.FlowNode{ID: "third-parties", Label: localized{"Terceros", "Third parties"}.text(locale), Kind: "external"})
		edges = append(edges, domain.FlowEdge{From: "service", To: "third-parties", DataLabel: localized{"Datos compartidos", "Shared data"}.text(locale)})
	}
	if hasCategory(findings, "advertising") {
		nodes = append(nodes, domain.FlowNode{ID: "ads", Label: localized{"Publicidad", "Advertising"}.text(locale), Kind: "external"})
		edges = append(edges, domain.FlowEdge{From: "service", To: "ads", DataLabel: localized{"Actividad e intereses", "Activity and interests"}.text(locale)})
	}
	if hasCategory(findings, "ai-training") {
		nodes = append(nodes, domain.FlowNode{ID: "ai", Label: localized{"Modelos de IA", "AI models"}.text(locale), Kind: "processing"})
		edges = append(edges, domain.FlowEdge{From: "service", To: "ai", DataLabel: localized{"Contenido", "Content"}.text(locale)})
	}
	return nodes, edges
}

func buildGuidance(findings []domain.Finding, gaps []domain.Gap, locale string) ([]domain.Impact, []domain.Action) {
	impacts := make([]domain.Impact, 0)
	actions := make([]domain.Action, 0)
	for _, finding := range findings {
		for _, candidate := range rules {
			if candidate.category == finding.Category {
				impacts = append(impacts, domain.Impact{DataType: candidate.dataLabel.text(locale), Scenario: candidate.impact.text(locale), Confidence: "scenario"})
				actions = append(actions, domain.Action{ID: fmt.Sprintf("action-%03d", len(actions)+1), Priority: finding.Severity, Text: candidate.action.text(locale)})
				break
			}
		}
	}
	accountDeletion := localized{"Eliminación de cuenta", "Account deletion"}.text(locale)
	for _, gap := range gaps {
		if gap.Control == accountDeletion && gap.Status != "specified" {
			actions = append(actions, domain.Action{ID: fmt.Sprintf("action-%03d", len(actions)+1), Priority: "medium", Text: localized{"Antes de aceptar, localiza el procedimiento de eliminación o solicita esa información al proveedor.", "Before accepting, find the deletion procedure or request that information from the provider."}.text(locale)})
			break
		}
	}
	if len(actions) > 6 {
		actions = actions[:6]
	}
	if len(impacts) > 4 {
		impacts = impacts[:4]
	}
	return impacts, actions
}

func containsAny(text string, patterns []string) bool {
	for _, pattern := range patterns {
		if strings.Contains(text, pattern) {
			return true
		}
	}
	return false
}
func hasCategory(findings []domain.Finding, category string) bool {
	for _, finding := range findings {
		if finding.Category == category {
			return true
		}
	}
	return false
}
func verifiedCount(findings []domain.Finding) int {
	count := 0
	for _, finding := range findings {
		if finding.Evidence.Verified {
			count++
		}
	}
	return count
}
func countInstructionPatterns(text string) int {
	lower, count := strings.ToLower(text), 0
	for _, pattern := range instructionPatterns {
		if strings.Contains(lower, pattern) {
			count++
		}
	}
	return count
}
func scoreLevel(score int) string {
	if score >= 75 {
		return "critical"
	}
	if score >= 50 {
		return "high"
	}
	if score >= 25 {
		return "medium"
	}
	return "low"
}

func summary(score, findings int, gaps []domain.Gap, locale string) string {
	missing := 0
	for _, gap := range gaps {
		if gap.Status != "specified" {
			missing++
		}
	}
	if locale == "en" {
		return fmt.Sprintf("The analysis found %d relevant practices and %d missing or ambiguous controls. Your personalized index is %d/100.", findings, missing, score)
	}
	return fmt.Sprintf("El análisis encontró %d prácticas relevantes y %d controles ausentes o ambiguos. El índice personalizado es %d/100.", findings, missing, score)
}

func detectLanguage(text string) string {
	spanish, english := 0, 0
	for _, word := range strings.Fields(strings.ToLower(text)) {
		word = strings.TrimFunc(word, func(r rune) bool { return !unicode.IsLetter(r) })
		switch word {
		case "el", "la", "los", "las", "para", "datos", "información", "recopilamos", "podemos":
			spanish++
		case "the", "and", "data", "your", "information", "collect", "may":
			english++
		}
	}
	if english > spanish {
		return "en"
	}
	return "es"
}
