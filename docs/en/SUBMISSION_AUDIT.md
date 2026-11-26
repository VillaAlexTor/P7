# Submission readiness audit — November 26, 2026

## Official submission requirements

The [STEM2Connect Devpost page](https://stem2connect-online-hackathon.devpost.com/) requires a functional project, a **public GitHub repository** with source, run instructions and AI disclosure, a **demo video of at most six minutes** on public or unlisted YouTube/Vimeo, and a Devpost description covering the problem, solution, AI use, learning, and next steps. State one primary track, include all team members, and disclose pre-hackathon work if any. Participants must be students aged 13 or older; teams may contain up to six people. The deadline shown is **November 29, 2026, 11:59 p.m. EST** (November 30, 12:59 a.m. in Bolivia). The submitted demo is the prerecorded presentation; a second video is not required. Live Demo Day is optional.

## Verified in this workspace

- [VillaAlexTor/P7](https://github.com/VillaAlexTor/P7) is public, with `main` as its default branch. GitHub Actions reported success for the last published commit checked during this audit; newly edited files need another CI run after push.
- The complete `scripts/test-all.ps1` suite passed: Go tests, four frontend tests, TypeScript, production build, production dependency audit with **zero vulnerabilities**, bilingual project-document parity, and real-evaluation source metadata.
- A real local API request used Ollama `qwen3:4b-instruct` and returned four findings with four textually verified citations, seven disclosure checks, a five-node/four-edge map, four impact scenarios, and five actions. The English example scored **97** for Journalist and **80** for General when analyzed separately. The web returned HTTP 200 and the API health endpoint reported the correct model. CORS allowed the configured web origin.
- The checked-in documentation is paired in English and Spanish. A scan of tracked filenames and common credential patterns found no committed secrets; `.cache` and the video's virtual environment and generated media are ignored.
- The Devpost text, architecture, README, and demo script now describe the current seven category passes and the distinction between a textually matching citation and semantic proof.

## Still required before submission

- Finish and review the actual video, keep it at **6:00 or less**, show a completed live analysis, and upload to YouTube or Vimeo as public or unlisted. Verify the link while signed out. The video is being produced separately and was **not** reviewed in this audit.
- Complete the Devpost form, select **Community Impact** as the primary track, add the repository and video links, confirm each team member has joined, and review the personal reflection and previous-work disclosure for factual accuracy. Registration, student eligibility, team membership, and the Devpost submission cannot be verified from the repository.
- A fresh Docker Compose startup was **not** completed on this machine: Docker Hub timed out during the `ollama/ollama` image TLS handshake. `docker compose config -q` passed, and the host Ollama + Go API + production Next.js path was tested successfully. Ports `8080` and `3000` may conflict with other local applications; the smoke test used API `18080` and web `13000`. This is a setup constraint to check on the recording machine.

## Accuracy limit to state honestly

The repaired round-18 run on four localized full-policy pages recorded **14 TP, 2 FP, 6 FN**, or 87.5% precision and 70.0% recall under its frozen labels, with **0/4** exact page-level sets. A one-reviewer audit judged **13/16** retained citations substantively supportive. Dropbox's contacts label conflicts with round 15, and the repaired pass used the same pages as the initial pass. These figures do not establish general reliability or 90% accuracy. See the [bilingual case study](ROUND18_CASE_STUDY.md) and its Spanish counterpart.
