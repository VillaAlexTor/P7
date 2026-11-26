# Demo Script

Target duration: 5 minutes 25 seconds. Hackathon limit: 6 minutes.

Use this as a natural guide rather than reading it word for word. Pause after starting the analysis if the local model needs extra time; that pause can be shortened during editing.

**Camera and narration:** You do not need to show your face. Record the application screen and narrate the demo with your own voice. Voice narration is not stated as a formal eligibility requirement, but it is strongly recommended because it makes your problem, decisions, learning, and next steps clear to the judges. A small face-camera window is optional; use it only if it makes you more comfortable. If you cannot record voice, use complete, readable English captions and on-screen explanations instead.

## 0:00–0:35 — Problem and impact

**On screen:** p7 home page in Spanish, then switch to English and keep English selected for the example analysis and the 97-to-80 comparison below.

**Say:**

“We accept privacy policies all the time, usually without reading them. Even when a tool summarizes one, it may not tell me what matters for my situation, what the document leaves unclear, or what I can actually do next. I built p7 to turn that difficult consent decision into evidence, possible impact, and concrete actions, with a complete experience in Spanish and English.”

## 0:35–1:10 — Input and profile

**On screen:** Load the example, select **Journalist**, and start the analysis.

**Say:**

“The user pastes a privacy policy instead of giving p7 access to an account or browsing history. I’ll use the included fictional example and choose the Journalist profile. This profile gives more weight to location, contacts, and third-party sharing. The policy is processed locally and is not intentionally stored by the application.”

## 1:10–1:55 — Personalized result and data map

**On screen:** Show the score, profile, clause count, SHA-256 identity, and data-flow map.

**Say:**

“This is not a score invented directly by the language model. The model identifies candidate practices; the Go backend verifies the evidence and applies transparent profile weights. The SHA-256 value identifies the exact normalized document used for this report. The map shows the path declared in the policy—from the person to the service and possible destinations. It does not claim to observe real network traffic.”

## 1:55–2:45 — Evidence and disclosure gaps

**On screen:** Open two findings and then the gaps section.

**Say:**

"Every confirmed finding returns to a cited clause. In this example, p7 identifies precise location and possible use of submitted content for AI training. The backend rejects quotes that do not match the referenced clause, allowing for PDF line breaks. A matching quote still needs human interpretation. p7 also separates specified, ambiguous, and missing disclosures; a missing disclosure is a reason to ask a question, not proof of a legal violation."

## 2:45–3:30 — Possible impact and actions

**On screen:** Show one impact scenario and mark two actions complete.

**Say:**

"The next section translates findings into possible scenarios without presenting them as facts. For example, precise location could reveal routines, while information shared with partners could travel beyond the original service. The action plan turns those concerns into practical steps, such as restricting location and reviewing sharing settings."

## 3:30–3:55 — Personalization

**On screen:** Return to the same example, change from **Journalist** to **General**, click **Analyze policy** again, and point out the changed score. In the verified English demo, it changed from 97 to 80.

**Say:**

"After I analyze the same policy under another profile, its priorities change. In this run the index changes from 97 to 80. A journalist, student, or parent may face different consequences. The profile weights are reproducible and are not changed by the model."

## 3:55–4:35 — Meaningful and disclosed AI use

**On screen:** Keep the report visible and point to the model metadata.

**Say:**

“The executable product uses `qwen3:4b-instruct` through local Ollama. It performs focused semantic-classification passes and returns structured candidates. Go checks the supported categories and whether cited text appears in the policy, deduplicates findings, and calculates the score. A matching quote still needs human review for meaning. ChatGPT and Codex also helped me with research, scope, architecture, coding, design, tests, and documentation. I reviewed and tested that work, and I disclose both kinds of AI assistance in the repository and in this submission.”

## 4:35–5:10 — Learning and an honest limitation

**On screen:** Briefly show the evaluation table in the README or Devpost gallery.

**Say:**

"My background gave me more confidence in backend and cybersecurity than in explainable AI. The first evaluation showed that the model had treated twelve denials as if the practices happened. Later, a PDF test exposed a citation-verification bug that rejected every Airbnb candidate. Fixing that bug raised recall from 30% to 70% on the same four localized pages, but none had a perfect category set and three citations still lacked semantic support. I learned that a working demo proves functionality, not general accuracy."

## 5:10–5:25 — Previous work and next steps

**On screen:** End on the action plan or project title.

**Say:**

“No application code existed before this hackathon; p7 was created during the event using credited open-source tools and models. Next I would add independent label review, user comprehension testing, policy-version comparison, and verified sources for Latin America. p7 does not replace legal advice. It gives people a clearer first reading before they accept.”

## Recording checklist

- Record the application screen; showing your face or webcam is optional.
- Prefer your own voice for the narration. Speak naturally; you do not need to memorize the script.
- Keep the final video at 6:00 or less.
- Show a real completed analysis, not only screenshots or slides.
- Keep model loading or waiting time short in the final edit.
- Add English captions and verify names such as Ollama, SHA-256, and `qwen3:4b-instruct`.
- Make the repository and video links viewable while signed out.
- Do not claim that user testing has occurred until sessions are actually completed.

## Criterion coverage

- Problem and impact: 0:00–0:35.
- Functionality: 0:35–3:55.
- Use of AI: 3:55–4:35.
- Learning and growth: 4:35–5:10.
- Presentation, prior work, and next steps: 5:10–5:25.
