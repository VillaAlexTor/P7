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

The web application uses Next.js, React, and TypeScript. A Go API normalizes and segments policies, applies a bilingual rule catalog, verifies every citation against its source clause, evaluates seven disclosure controls, applies transparent profile weights, and derives the map and guidance. The MVP deliberately uses no account system or database.

## How we used AI

ChatGPT and Codex supported competitor research, problem framing, scope, architecture, coding, design, tests, and documentation. This assistance is disclosed in the repository and submission.

The current executable remains reproducible without paid API credits. AI can later provide semantic classification candidates, but it will not control the final score. The Go backend must still verify citations, deduplicate results, and apply deterministic weights. We chose this design because a plausible answer is insufficient for privacy: every finding must trace back to the policy.

## Challenges

The hardest challenge was preventing the product from becoming another legal-text summarizer. We had to connect seven features while keeping the MVP usable, explain uncertainty without creating legal claims, and treat the analyzed document as untrusted input. We also had to preserve exact evidence while making the result understandable to non-specialists.

## What we learned

Trust requires separating what the document confirms, what the system interprets, and what could happen. A useful score must expose its weights and change with a person's priorities. Reducing the scope to pasted text, five profiles, and an in-memory engine made it possible to ship a complete flow rather than many incomplete features. Next time, we would prepare the labeled evaluation set and begin user-comprehension testing earlier.

## Accomplishments

- All seven proposed differentiators work in one flow.
- Every displayed citation is checked against its clause.
- The same policy produces different deterministic scores for different profiles.
- The interface and analysis output work in English and Spanish.
- Automated tests cover the seven pillars, API contract, security behavior, frontend language parity, production build, dependency audit, and documentation parity.

## What's next

We will evaluate the engine with manually labeled clauses and short user studies. Future work includes an optional semantic classifier, policy-version comparison, safe URL and PDF ingestion, accessible report export, and verified legal sources for Latin American jurisdictions.

## Disclosure

p7 is an informational tool. It does not provide legal advice, confirm regulatory compliance, or observe an application's real network traffic.

