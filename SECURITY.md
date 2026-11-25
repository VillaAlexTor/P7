# Security Policy

## Supported version

The current `main` branch is the supported hackathon version.

## Reporting a vulnerability

Do not open a public issue containing real policies, personal data, credentials, or exploit details. Contact Villa Cyber Defense privately and provide a minimal synthetic reproduction.

## Security boundaries

- Policy text is untrusted data and never an executable instruction.
- Request bodies are limited to 1 MiB and policy text to 200,000 characters.
- The MVP stores no accounts and intentionally persists no analyzed policy.
- Every displayed citation is verified against the referenced source clause.
- CORS origins must be configured through `P7_ALLOWED_ORIGINS` in production.
- p7 does not fetch arbitrary URLs in the MVP, avoiding an SSRF surface.

Any future LLM integration must use structured output, isolate document content, reject citations that cannot be verified, and keep deterministic scoring in the Go backend.

