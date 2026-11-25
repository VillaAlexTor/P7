# Demo Deployment Decision

For this hackathon submission, the reliable path is a public GitHub repository, reproducible local setup, and a recorded end-to-end demo. Ollama runs on the presenter's computer; the model file is downloaded separately and is not committed. This avoids promising a public web endpoint that cannot reach a private laptop's local Ollama process.

Run `ollama pull qwen3:4b-instruct`, start Ollama, run the Go API, and run the Next.js frontend as described in the root README. Validate once from a clean checkout before recording. The public repository and video satisfy the published submission requirements; a hosted URL is useful but not listed as mandatory in the official rules.

A later public deployment needs a continuously available machine with enough memory for the model, an HTTPS frontend and API, rate limits, request timeouts, monitoring, and a decision about how users' policy text is handled. Docker Compose alone does not provide those production guarantees. Do not expose the Ollama port publicly without a deliberate security design.
