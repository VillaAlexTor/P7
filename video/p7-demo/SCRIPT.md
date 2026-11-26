# English Narration

## 01 — Problem

We accept privacy policies all the time, usually without reading them. Even a summary may not tell me what matters for my situation, what the document leaves unclear, or what I can actually do next. I built p7 to turn that difficult consent decision into evidence, possible impact, and concrete actions, in Spanish and English.

## 02 — Input

The user pastes a privacy policy; p7 never needs account access or browsing history. I load the fictional example, choose Journalist, and press Analyze policy. That profile gives more weight to location, contacts, and third-party sharing. The text is processed locally and is not intentionally stored.

## 03 — Result and map

This index is not invented directly by the language model. The model identifies candidate practices. The Go backend verifies the evidence and applies transparent profile weights. The SHA-256 value identifies the exact normalized document used for this report. The map shows the path declared in the policy. It does not claim to observe real network traffic.

## 04 — Evidence and gaps

Every confirmed finding returns to an exact clause. In this run, p7 reports location, third-party sharing, personalized advertising, and AI training—but not contacts. The stricter definition requires explicit device address-book access. If a quote is absent from its referenced clause, Go rejects it. Missing information is a question, not proof of a violation.

## 05 — Impact and action

The next section translates findings into possible scenarios without presenting them as facts. Location could reveal routines. Shared activity may shape advertising, and submitted content may enter AI improvement. The action plan makes those concerns practical: restrict location, review partner sharing and ad controls, and avoid sensitive content when AI training cannot be disabled.

## 06 — Personalization

The evidence stays the same, but the priority changes. I select General and press Analyze again. The same four findings remain visible, while the index moves from ninety-seven for Journalist to eighty for General. Those profile weights are transparent and reproducible; the model cannot alter them.

## 07 — AI design

The executable product uses qwen three, four billion instruct through local Ollama. It performs focused semantic classification and returns structured candidates. Go rejects unsupported categories or fabricated evidence, removes duplicates, and calculates the index. ChatGPT and Codex also helped with research, scope, architecture, code, design, tests, and documentation. I reviewed and tested that work, and I disclose both kinds of AI assistance.

## 08 — The seven pillars

Those capabilities are the seven pillars behind the name p7: data mapping, personalized risk, gap detection, verifiable evidence, impact simulation, smart consent, and privacy for everyone. The final pillar matters. p7 works in English and Spanish, using plain-language explanations designed to reach people beyond legal and technical audiences, including communities in Latin America.

## 09 — Learning and limitation

My background gave me more confidence in backend and cybersecurity than in explainable AI. Early on, the model treated twelve denial clauses as if those practices happened. The answer looked confident, but it was wrong. In the latest revised complete-policy evaluation, precision reached eighty-seven point five percent and recall seventy percent, yet none of the four localized pages had a perfect category set. I learned that a polished demo proves functionality, not reliability, and that even exact evidence still needs human interpretation.

## 10 — Close

No application code existed before this hackathon. p7 was created during the event using credited open-source tools and models. Next, I would add independent label review, user comprehension testing, policy-version comparison, and verified sources for Latin America. p7 does not replace legal advice. It gives people a clearer first reading before they accept. Evidence before confidence.
