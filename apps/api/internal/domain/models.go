package domain

type AnalyzeRequest struct {
	Text    string `json:"text"`
	Profile string `json:"profile"`
	Title   string `json:"title,omitempty"`
	Locale  string `json:"locale,omitempty"`
}

type Clause struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type Evidence struct {
	ClauseID string `json:"clauseId"`
	Quote    string `json:"quote"`
	Verified bool   `json:"verified"`
}

type Finding struct {
	ID             string   `json:"id"`
	Category       string   `json:"category"`
	Title          string   `json:"title"`
	Severity       string   `json:"severity"`
	Interpretation string   `json:"interpretation"`
	Evidence       Evidence `json:"evidence"`
	Weight         int      `json:"weight"`
}

type Gap struct {
	Control string `json:"control"`
	Status  string `json:"status"`
	Note    string `json:"note"`
}

type FlowNode struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Kind  string `json:"kind"`
}

type FlowEdge struct {
	From      string `json:"from"`
	To        string `json:"to"`
	DataLabel string `json:"dataLabel"`
}

type Impact struct {
	DataType   string `json:"dataType"`
	Scenario   string `json:"scenario"`
	Confidence string `json:"confidence"`
}

type Action struct {
	ID       string `json:"id"`
	Priority string `json:"priority"`
	Text     string `json:"text"`
}

type AnalysisMeta struct {
	Title           string `json:"title"`
	Profile         string `json:"profile"`
	DocumentHash    string `json:"documentHash"`
	ClauseCount     int    `json:"clauseCount"`
	Language        string `json:"language"`
	Engine          string `json:"engine"`
	EvidenceChecked int    `json:"evidenceChecked"`
	Locale          string `json:"locale"`
	AIUsed          bool   `json:"aiUsed"`
	AIProvider      string `json:"aiProvider"`
	AIModel         string `json:"aiModel"`
	AICandidates    int    `json:"aiCandidates"`
}

type SafetyReport struct {
	InputTreatedAsData         bool `json:"inputTreatedAsData"`
	IgnoredInstructionPatterns int  `json:"ignoredInstructionPatterns"`
	CitationsVerified          int  `json:"citationsVerified"`
	CitationsRejected          int  `json:"citationsRejected"`
}

type AnalysisResponse struct {
	Meta     AnalysisMeta `json:"meta"`
	Score    int          `json:"score"`
	Level    string       `json:"level"`
	Summary  string       `json:"summary"`
	Clauses  []Clause     `json:"clauses"`
	Findings []Finding    `json:"findings"`
	Gaps     []Gap        `json:"gaps"`
	Nodes    []FlowNode   `json:"nodes"`
	Edges    []FlowEdge   `json:"edges"`
	Impacts  []Impact     `json:"impacts"`
	Actions  []Action     `json:"actions"`
	Safety   SafetyReport `json:"safety"`
}
