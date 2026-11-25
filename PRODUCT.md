# Product

<!-- impeccable:product-schema 1 -->

## Platform

web

## Stack

Delegated from the existing project plan: Next.js with TypeScript for the web application, Go for the analysis API, Ollama with `qwen3:4b-instruct` for local semantic classification, and Docker Compose for orchestration. The MVP requires a real LLM but no paid AI API or API key.

## Users

The primary user is a student or everyday internet user in Latin America who needs to understand a privacy policy before accepting it. The product must work completely in English and Spanish and remain understandable without legal, privacy, or technical training.

## Product Purpose

p7 turns dense privacy policies and terms into personalized, evidence-backed privacy assessments. A successful session tells the user what data the service says it collects, where that data may go, what important information is missing or ambiguous, what consequences may follow, and what the user can do next.

## Positioning

p7 combines seven connected mechanisms instead of acting as a generic document summarizer: data-flow mapping, profile-weighted risk, missing-information detection, verifiable evidence, impact simulation, actionable consent guidance, and a Latin American accessibility focus. A local LLM classifies clauses semantically; deterministic rules verify citations and calculate the final score.

## Operating Context

For the hackathon MVP, a user pastes policy text, selects a privacy profile, and receives one analysis workspace. The demonstration must work locally without accounts, a database, or a paid API. The public submission needs source code, setup instructions, AI disclosure, a working demo flow, and material for a video no longer than six minutes.

## Capabilities and Constraints

- Accept pasted Spanish or English policy text and return a report in the selected locale. URL and PDF ingestion remain optional follow-up work.
- Segment the source into stable clauses and hash the complete normalized document with SHA-256.
- Detect declared data categories, purposes, recipients, retention, user rights, security statements, and AI-training language.
- Flag absent or ambiguous disclosures without presenting an omission as proof of illegality.
- Attach every confirmed finding to exact source text and validate that the citation exists.
- Calculate risk through reproducible rules weighted by a selected privacy profile; an LLM must not directly set the final score.
- Clearly separate confirmed evidence, interpretation, and hypothetical impact.
- Provide a visual data-flow map and a prioritized action plan.
- Treat document content as untrusted input and never execute instructions found inside it.
- No legal advice claims. Results explain declared practices and informational gaps.

## Brand Commitments

- Product name: `p7`.
- English line: `Privacy. Decoded.`
- Spanish line: `Privacidad, descifrada.`
- Seven pillars are a core part of the name and product story.
- Developer identity: Villa Cyber Defense.
- Voice: direct, calm, specific, and accessible. Avoid alarmism and generic AI marketing language.

## Evidence on Hand

The source ChatGPT conversation defines the seven differentiators, proposed architecture, MVP scope, hackathon requirements, and illustrative examples. No verified customer testimonials, performance benchmarks, legal certification, trademark clearance, or production deployment claims exist; future work must not fabricate them.

## Product Principles

- Evidence before interpretation.
- Reproducible scores over opaque model judgment.
- Explain consequences in ordinary language.
- Show uncertainty and omission honestly.
- Make privacy understanding accessible to Latin American communities.

## Accessibility & Inclusion

The MVP provides complete English and Spanish interfaces and analysis output. The interface must meet WCAG 2.2 AA fundamentals, support keyboard use and reduced motion, use plain language, and never rely on color alone to communicate risk.

## Open Decisions

- Production hosting remains undecided; the MVP AI provider is Ollama with `qwen3:4b-instruct`.
- Domain and trademark availability for `p7` have not been verified.
- The visual direction recorded during this build is an implementation assumption and remains revisable.
