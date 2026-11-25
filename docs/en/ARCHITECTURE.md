# Architecture and Trust Model

## System shape

```text
Browser (Next.js, TypeScript)
        │ POST /api/v1/analyze
        ▼
Go HTTP API
  ├─ input limits and locale selection
  ├─ normalization and clause segmentation
  ├─ Ollama structured semantic classification
  ├─ exact evidence verification
  ├─ deterministic profile weighting
  └─ flow map, gaps, scenarios, and actions
```

The MVP has no account system or database. Policy text is processed in memory and is not intentionally persisted by the application.

## Analysis flow

1. Normalize line endings and spacing without rewriting meaning.
2. Split the document into stable clause identifiers.
3. Send each clause batch to local `qwen3:4b-instruct` in five focused category passes, each with a constrained JSON schema.
4. Receive semantic categories, clause identifiers, and verbatim proposed quotes.
5. Reject any unsupported category or quote not contained exactly in its referenced clause.
6. Classify disclosure controls as specified, ambiguous, or missing.
7. Apply transparent profile multipliers and cap the index at 100.
8. Derive a declared data-flow map, possible-impact scenarios, and actions.
9. Return a SHA-256 hash for the normalized document version.

## Seven differentiators in code

| Differentiator | Implementation |
| --- | --- |
| Data Mapping | `buildFlow` produces typed nodes and edges from verified findings. |
| Risk Engine | `calculateScore` uses deterministic weights and profile multipliers. |
| Gap Detection | `detectGaps` evaluates seven disclosure controls. |
| Evidence Engine | `verifyEvidence` rejects empty, unknown, or fabricated citations. |
| Impact Simulation | `buildGuidance` emits scenarios explicitly labeled as hypothetical. |
| Smart Consent | The same guidance layer creates prioritized, interactive actions. |
| Privacy for Everyone | UI, samples, prompts, errors, and reports support Spanish and English. |

## Trust boundaries

The document is untrusted data. p7 never executes instructions found inside it. Known instruction-like patterns are counted in the safety report, but they do not alter analysis behavior. Request bodies are limited to 1 MiB and documents to 200,000 characters. CORS origins are configured through `P7_ALLOWED_ORIGINS`.

Ollama runs `qwen3:4b-instruct` locally. The five passes cover location/contacts/biometrics, third parties, advertising, AI training, and retention. Each pass returns candidates constrained by a category-specific JSON schema: category, clause ID, and verbatim quote. The Go backend rejects unsupported categories and quotes that do not occur in the referenced clause, deduplicates findings, applies profile weights, and calculates the final score. Five model calls per batch increase latency and must be measured on full policies. If Ollama is unavailable, the analysis endpoint returns a service error instead of silently replacing AI with keyword classification.

## Deliberate scope decisions

- No authentication reduces attack surface.
- No database avoids retaining potentially sensitive policies.
- No URL fetching avoids SSRF and dynamic-content complexity in the MVP.
- No legal conclusion: results explain declared practices and disclosure gaps.
- No claim of network observation: the map reflects policy text only.

