# p7

**Privacy. Decoded.**

**English** · [Español](README.es.md)

p7 turns privacy policies into personal, explainable, evidence-backed assessments. The MVP works without accounts, a database, or a paid AI API.

## Seven implemented differentiators

1. **Data Mapping:** represents the path declared in the policy.
2. **Personal Risk:** weights findings for five privacy profiles.
3. **Gap Detection:** classifies disclosures as specified, ambiguous, or missing.
4. **Verifiable Evidence:** links every finding to an exact clause and SHA-256 document identity.
5. **Impact Simulation:** explains possible consequences without presenting them as facts.
6. **Smart Consent:** creates concrete, prioritized actions.
7. **Privacy for Everyone:** complete interface, examples, and reports in English and Spanish.

## Stack

- Frontend: Next.js 16, React 19, and TypeScript.
- Backend: Go using the standard HTTP library.
- Tests: Go `testing`, `httptest`, and Vitest.
- Operations: Docker Compose and GitHub Actions.

## Run locally

Terminal 1:

```powershell
cd apps/api
$env:P7_API_PORT = "8080"
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

## Docker

```powershell
docker compose up --build
```

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
- Every citation is verified against its referenced clause.
- The MVP intentionally persists no policy and has no user accounts.
- Production CORS is configured with `P7_ALLOWED_ORIGINS`.
- Results are informational and not legal advice.

## AI use

ChatGPT and Codex supported research, scope, architecture, coding, design, tests, and documentation. The current product uses a deterministic local engine and does not depend on external credits. A future semantic classifier will remain optional and will not control the final score.

## Documentation

English is the official submission language and every official document has a Spanish counterpart. See the [bilingual index](docs/README.md), [architecture](docs/en/ARCHITECTURE.md), [judging strategy](docs/en/JUDGING.md), and [demo script](docs/en/DEMO_SCRIPT.md).

## License

MIT. See [LICENSE](LICENSE).
