# p7 — Privacy. Decoded.

<p align="center"><img src="docs/assets/devpost-thumbnail-p7.png" width="360" alt="p7 — Privacy. Decoded." /></p>

**A bilingual, evidence-backed privacy policy reader powered by a local LLM.**

**English** · [Español](README.es.md) · [Documentation](docs/README.md)

**STEM2Connect hackathon MVP.** Privacy policies are long and hard to compare. p7 turns pasted policy text into a structured, personal reading: declared practices, exact evidence, disclosure gaps, possible impacts, and concrete actions. It is a working local application designed for human review.

## Quick start and demo

With Docker Desktop running and Docker Hub reachable, use `iniciar.cmd` on Windows or run `docker compose up --build --wait` in a terminal. The first launch downloads container images and the local model and may take several minutes. Open **http://localhost:3000**, choose a language and profile, click **Load example**, then **Analyze policy**. Review the cited findings, data map, gaps, impact scenarios, and actions. The included example is fictional; you can also paste your own policy text. The app does not fetch arbitrary URLs. If image download fails, use the local setup below.

Stop the stack with `docker compose down`.

## Seven implemented differentiators

1. **Data Mapping:** represents the path declared in the policy.
2. **Personal Risk:** weights findings for five privacy profiles.
3. **Gap Detection:** classifies disclosures as specified, ambiguous, or missing.
4. **Verifiable Evidence:** links every finding to an exact clause and SHA-256 document identity.
5. **Impact Simulation:** explains possible consequences without presenting them as facts.
6. **Smart Consent:** creates concrete, prioritized actions.
7. **Privacy for Everyone:** complete interface, examples, and reports in English and Spanish.

## Measured AI evaluation

Cases, official source URLs, and **raw model output** are stored in [apps/api/evaluation](apps/api/evaluation). The [English history](docs/en/EVALUATION.md) and [Spanish history](docs/es/EVALUATION.md) explain every round and annotation caveats.

| Frozen classifier, new official excerpts | Precision | Recall | Exact cases |
| --- | ---: | ---: | ---: |
| Round 7 · 40 cases | 73.7% | 70.0% | 30/40 |
| Round 8 · 40 cases | 96.0% | 85.7% | 35/40 |
| **Round 9 · 40 cases** | **100.0%** | **75.0%** | **30/40** |

Round 9 used 40 selected short excerpts from official pages. Those results do **not** measure full-document performance. The prompt and model weights stayed unchanged between rounds.

[Cases A](apps/api/evaluation/real_round9a.tsv) · [Sources A](apps/api/evaluation/real_round9a_sources.tsv) · [Raw output A](apps/api/evaluation/real_round9a_result.json) · [Cases B](apps/api/evaluation/real_round9b.tsv) · [Sources B](apps/api/evaluation/real_round9b_sources.tsv) · [Raw output B](apps/api/evaluation/real_round9b_result.json)

**Complete-policy checks:** We also tested full official policies in English and Spanish. These document-level results are closer to actual use than selected excerpts; prompts and evidence verification changed between some later rounds:

| Official complete policies | Precision | Recall | Exact pages |
| --- | ---: | ---: | ---: |
| Round 10 · Mozilla Accounts, Proton | 54.5% | 60.0% | 1/4 |
| Round 11 · Apple, Spotify | 90.5% | 95.0% | 2/4 |
| **Round 12 · Google, WhatsApp** | **72.7%** | **100.0%** | **1/4** |
| **Round 13 · OpenAI, TikTok EEA** | **100.0%** | **95.8%** | **3/4** |
| Round 14 · LinkedIn, Zoom | 91.3% | 95.5% | 1/4 |
| Round 15 · Dropbox, GitHub | 85.0% | 94.4% | 1/4 |
| Round 16 · Pinterest, Slack | 81.0% | 85.0% | 1/4 |
| Round 17 · IKEA Spain, after prompt revision | 87.5% | 87.5% | 1/2 |
| **Round 18 · Dropbox, Airbnb, after PDF evidence fix** | **87.5%** | **70.0%** | **0/4** |

Round 12 used pages up to 120,506 characters. It found all 16 reference categories but added six unsupported findings; only WhatsApp Spanish was exact. The four analyses took **10 minutes 46 seconds** locally. The Google AI-training label is contestable and is examined separately in the [round 12 case study](docs/en/ROUND12_CASE_STUDY.md).

Round 18 revisited Dropbox with a stricter contact label and added Airbnb in two languages. A PDF line-wrap bug initially caused all Airbnb candidates to be rejected. After fixing textual citation verification, the same four pages yielded **14 true positives, 2 false positives, and 6 misses** under the frozen round-18 labels; none had a perfect category set. A separate human audit found **13 of 16 retained citations** semantically supportive. The repaired run is an **in-sample check**, not an independent holdout. The [bilingual round-18 study](docs/en/ROUND18_CASE_STUDY.md) and [raw outputs](apps/api/evaluation/real_policies_round18_result.json) explain the errors and the disputed Dropbox contact label, which was labeled positive in round 15. Selected policies, correlated translations, and one reviewer cannot establish general reliability. Human review remains necessary.

## Stack

- Frontend: Next.js 16, React 19, and TypeScript.
- Backend: Go using the standard HTTP library.
- AI: Ollama with `qwen3:4b-instruct` structured semantic classification.
- Tests: Go `testing`, `httptest`, and Vitest.
- Operations: Docker Compose and GitHub Actions.

## Run locally

Install Ollama, download `qwen3:4b-instruct` once, and keep Ollama running on `http://127.0.0.1:11434`:

```powershell
ollama pull qwen3:4b-instruct
```

Terminal 1:

```powershell
cd apps/api
$env:P7_API_PORT = "8080"
$env:P7_OLLAMA_URL = "http://127.0.0.1:11434"
$env:P7_OLLAMA_MODEL = "qwen3:4b-instruct"
go run ./cmd/server
```

Terminal 2:

```powershell
cd apps/web
npm install
$env:NEXT_PUBLIC_API_URL = "http://localhost:8080"
npm run dev
```

Open `http://localhost:3000`, load the example, and run the analysis. The ES/EN control switches the complete experience.

If ports 8080 or 3000 are occupied, set `P7_API_PORT`, `P7_ALLOWED_ORIGINS`, and `NEXT_PUBLIC_API_URL` to matching alternatives, then start Next.js with `npm run dev -- -p <web-port>`. For example, API `18080` and web `13000` were used in the final local smoke test.

## Docker

With Docker Desktop running, execute `iniciar.cmd` from the project root (double-click it or run it in a terminal). It builds the web app and API, starts Ollama, downloads the local model on first use, waits until every service is healthy, and opens `http://localhost:3000`.

You can also start it manually:

```powershell
docker compose up --build --wait
```

The model is stored in the `ollama-data` volume, so it is not downloaded on every start. Run `docker compose down` to stop the services, or `docker compose down --volumes` to also remove the downloaded model.

## Automated verification

```powershell
./scripts/test-all.ps1
```

The command runs Go tests, frontend tests, TypeScript checks, the production build, dependency audit, and bilingual documentation parity. GitHub Actions repeats these checks on every push and pull request.

## API

`POST /api/v1/analyze`

```json
{
  "title": "SocialApp Demo",
  "profile": "journalist",
  "locale": "en",
  "text": "We collect your location..."
}
```

Profiles: `general`, `student`, `journalist`, `parent`, and `activist`. Locales: `en` and `es`.

## Security and limits

- Maximum request body: 1 MiB; maximum document: 200,000 characters.
- Policy content is untrusted data; instruction-like content is ignored.
- Every citation is checked for textual presence in its referenced clause, allowing whitespace differences from PDF line wraps; this is not semantic validation.
- The MVP intentionally persists no policy and has no user accounts.
- Production CORS is configured with `P7_ALLOWED_ORIGINS`.
- Results are informational and not legal advice.

## AI use

The executable product calls the local `qwen3:4b-instruct` model through Ollama in seven category-focused passes per clause batch. The model classifies clauses with a constrained JSON schema and proposes citations. Go rejects unknown categories and quotes that do not textually match the referenced clause after whitespace normalization, then applies profile weights and calculates the score. A matching quote can still support an incorrect interpretation. When Ollama is unavailable, analysis returns a service error instead of silently substituting a keyword classifier. ChatGPT and Codex also supported research, scope, architecture, coding, design, tests, and documentation. No API key or secret is stored here.

## Documentation

English is the official submission language and every official document has a Spanish counterpart. See the [bilingual index](docs/README.md), [submission readiness audit](docs/en/SUBMISSION_AUDIT.md), [AI evaluation](docs/en/EVALUATION.md), [five-person study protocol](docs/en/USER_STUDY.md), and [demo deployment decision](docs/en/DEPLOYMENT.md).

The [presentation gallery](docs/assets/devpost-gallery/ready-1500x1000/) contains eight product screenshots; a [ready-to-upload ZIP](docs/assets/devpost-gallery/p7-devpost-gallery-1500x1000.zip) is also available.

## License

MIT. See [LICENSE](LICENSE).
