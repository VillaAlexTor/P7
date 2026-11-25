export type Severity = "low" | "medium" | "high" | "critical";
export type GapStatus = "specified" | "ambiguous" | "missing";

export interface AnalysisResponse {
  meta: {
    title: string;
    profile: string;
    documentHash: string;
    clauseCount: number;
    language: string;
    engine: string;
    evidenceChecked: number;
    locale: string;
    aiUsed: boolean;
    aiProvider: string;
    aiModel: string;
    aiCandidates: number;
  };
  score: number;
  level: Severity;
  summary: string;
  clauses: Array<{ id: string; text: string }>;
  findings: Array<{
    id: string;
    category: string;
    title: string;
    severity: Severity;
    interpretation: string;
    evidence: { clauseId: string; quote: string; verified: boolean };
    weight: number;
  }>;
  gaps: Array<{ control: string; status: GapStatus; note: string }>;
  nodes: Array<{ id: string; label: string; kind: string }>;
  edges: Array<{ from: string; to: string; dataLabel: string }>;
  impacts: Array<{ dataType: string; scenario: string; confidence: string }>;
  actions: Array<{ id: string; priority: Severity; text: string }>;
  safety: {
    inputTreatedAsData: boolean;
    ignoredInstructionPatterns: number;
    citationsVerified: number;
    citationsRejected: number;
  };
}

