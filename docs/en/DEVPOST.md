# Devpost Submission Draft

## Project name

p7 — Privacy. Decoded.

## Track

Community Impact

## Inspiration and problem

People accept long privacy policies without understanding what data they give away, who may receive it, what the document leaves unclear, or what they can do next. Generic summaries shorten text but rarely support a personal decision. This is especially relevant for students and people without legal or technical training, and many existing tools do not begin with Spanish-speaking Latin American users.

## What it does

p7 transforms a policy into a personalized, verifiable privacy assessment. Users paste text, choose one of five privacy profiles, and receive:

- a declared data-flow map;
- a deterministic personal risk index;
- missing and ambiguous disclosure checks;
- findings linked to exact verified clauses;
- clearly labeled possible-impact scenarios;
- a prioritized consent action plan;
- complete Spanish and English experiences.

The normalized document receives a SHA-256 hash so the report identifies the exact analyzed version.

## How we built it

The web application uses Next.js, React, and TypeScript. A Go API segments policies and sends clause batches to local Ollama, where `qwen3:4b-instruct` returns schema-constrained semantic candidates. Go verifies every citation, evaluates disclosure controls, applies transparent profile weights, and derives the map and guidance. The MVP deliberately uses no account system, database, paid API, or API key.

## How we used AI

ChatGPT and Codex supported competitor research, problem framing, scope, architecture, coding, design, tests, and documentation. This assistance is disclosed in the repository and submission.

The executable calls `qwen3:4b-instruct` through local Ollama and requests structured semantic-classification candidates. It remains reproducible without paid API credits or secrets. Go verifies every proposed quote against its clause, rejects unsupported output, deduplicates results, and applies deterministic weights. We chose this design because a plausible answer is insufficient for privacy: every finding must trace back to the policy.

## Challenges

The hardest challenge was preventing the product from becoming another legal-text summarizer. We had to connect seven features while keeping the MVP usable, explain uncertainty without creating legal claims, and treat the analyzed document as untrusted input. We also had to preserve exact evidence while making the result understandable to non-specialists.

## What we learned

Trust requires separating what the document confirms, what the system interprets, and what could happen. In our first 40-clause synthetic evaluation, the model classified all 28 affirmative practices but also misclassified all 12 denials. Prompt revision reduced false positives to zero, while one affirmative retention clause remained missed. This taught us to measure mistakes explicitly and to treat a demo as evidence of functionality, not accuracy. The [personal reflection draft](REFLECTION.md) needs the author's review before submission.

## Accomplishments

- All seven proposed differentiators work in one flow.
- Every displayed citation is checked against its clause.
- The same policy produces different deterministic scores for different profiles.
- The interface and analysis output work in English and Spanish.
- Automated tests cover the seven pillars, API contract, security behavior, frontend language parity, production build, dependency audit, and documentation parity.
- A reproducible [40-clause synthetic evaluation](EVALUATION.md) reports 27 true positives, zero false positives, and one false negative after in-sample prompt tuning.
- A separate 20-excerpt check from published privacy sources found 17 of 19 labeled practices with no false positive labels in that run. The sample is small and does not measure whole-policy accuracy.

## What's next

We will expand the real-policy evaluation, review labels with another annotator, and conduct the prepared [five-person usability study](USER_STUDY.md). Future work includes policy-version comparison, safe URL and PDF ingestion, accessible report export, and verified legal sources for Latin American jurisdictions.

## Disclosure

p7 is an informational tool. It does not provide legal advice, confirm regulatory compliance, or observe an application's real network traffic.

