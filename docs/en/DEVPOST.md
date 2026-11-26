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
- findings linked to cited policy clauses, with textual quote verification;
- clearly labeled possible-impact scenarios;
- a prioritized consent action plan;
- complete Spanish and English experiences.

The normalized document receives a SHA-256 hash so the report identifies the exact analyzed version.

## How we built it

The web application uses Next.js, React, and TypeScript. A Go API segments policies and sends clause batches to local Ollama, where `qwen3:4b-instruct` returns schema-constrained semantic candidates. Go checks each citation for textual presence in its clause, evaluates disclosure controls, applies transparent profile weights, and derives the map and guidance. This check does not prove the quote's interpretation. The MVP deliberately uses no account system, database, paid API, or API key.

## How we used AI

ChatGPT and Codex supported competitor research, problem framing, scope, architecture, coding, design, tests, and documentation. This assistance is disclosed in the repository and submission.

The demo video uses locally synthesized English narration via Kokoro and was assembled with HyperFrames. The voice is synthetic rather than a recording of the author's speech.

The executable calls `qwen3:4b-instruct` through local Ollama and requests structured semantic-classification candidates. It remains reproducible without paid API credits or secrets. Go checks proposed quotes against their cited clauses, allowing whitespace differences from PDF extraction, rejects unknown categories, deduplicates results, and applies deterministic weights. A separate one-reviewer citation audit shows that a matching quote can still support the wrong interpretation.

## Challenges

The hardest challenge was preventing the product from becoming another legal-text summarizer. We had to connect seven features while keeping the MVP usable, explain uncertainty without creating legal claims, and treat the analyzed document as untrusted input. We also had to preserve exact evidence while making the result understandable to non-specialists.

## What we learned

I came into p7 with more confidence in backend development and cybersecurity than in explainable AI or product design. My first instinct was to focus on making the analysis run. The first evaluation corrected me: the model confidently treated all twelve denial clauses as if the practices actually occurred. That was an uncomfortable but useful mistake. I revised the prompt, added negative controls, kept exact evidence checks, and stopped treating a convincing output as a correct one.

The later real-source rounds were just as important. Results varied sharply between selected excerpts and complete policies. In round 18, PDF line wraps caused our literal citation gate to reject every Airbnb candidate. Fixing whitespace handling raised recall on those same four localized pages from 30% to 70%, yet none had a perfect category set and 3 of 16 retained citations lacked semantic support. I learned that one strong result is not a reliability guarantee, a working demo is not an accuracy study, and a matching citation can still support an incorrect interpretation. If I started again, I would test complete documents and user comprehension much earlier.

## Accomplishments

- All seven proposed differentiators work in one flow.
- Every displayed citation is checked against its clause.
- The same policy produces different deterministic scores for different profiles.
- The interface and analysis output work in English and Spanish.
- Automated tests cover the seven pillars, API contract, security behavior, frontend language parity, production build, dependency audit, and documentation parity.
- A reproducible [40-clause synthetic evaluation](EVALUATION.md) reports 27 true positives, zero false positives, and one false negative after in-sample prompt tuning.
- Nine full-policy rounds have frozen source records and raw model outputs. The latest revised pass on Dropbox and Airbnb found 14 of 20 reference categories, added 2 under its disputed Dropbox contact label, and missed 6: 87.5% precision, 70.0% recall, and 0 of 4 exact localized pages. It was an in-sample repair check, not a fresh independent holdout. A one-reviewer audit found 13 of 16 retained citations substantively relevant; the [round-18 case study](ROUND18_CASE_STUDY.md) records the caveats.

## What's next

We will review the labels with another annotator, evaluate more complete policies from independent providers, and conduct the prepared [five-person usability study](USER_STUDY.md). Future work includes policy-version comparison, safe URL and PDF ingestion, accessible report export, and verified legal sources for Latin American jurisdictions.

## Previous work disclosure

No application code existed before this hackathon. The concept, implementation, evaluation, documentation, and presentation materials for p7 were created during the event. The project uses credited open-source frameworks, libraries, the Ollama runtime, and the `qwen3:4b-instruct` model.

## Disclosure

p7 is an informational tool. It does not provide legal advice, confirm regulatory compliance, or observe an application's real network traffic.

