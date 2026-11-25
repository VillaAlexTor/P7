# AI Evaluation — Synthetic Clause Set

## Method

Forty hand-labeled, synthetic, single-clause cases were sent individually to the local `qwen3:4b-instruct` model through the same structured Ollama classifier used by p7. Twenty are English and twenty Spanish. Each language contains fourteen affirmative practices across seven categories and six negative controls, including denials and a finite deletion period. The evaluator accepts a prediction only when its category is supported and the proposed quote appears verbatim in the referenced clause.

Run from `apps/api` with Ollama running and the model installed:

```powershell
go run ./cmd/evaluate -data evaluation/clauses.tsv
```

The dataset and evaluator are committed. Output includes every case, per-category counts, and elapsed time. No user information or real company policy was used.

## Observed result, November 25, 2026

| Version | Correct positives | False positives | Missed positives | Precision | Recall |
| --- | ---: | ---: | ---: | ---: | ---: |
| Initial prompt | 28 | 12 | 0 | 70.0% | 100.0% |
| Explicit negation instruction | 22 | 0 | 6 | 100.0% | 78.6% |
| Current prompt with affirmative examples | 27 | 0 | 1 | 100.0% | 96.4% |

The current prompt recognized all twelve negative controls and 27 of 28 affirmative clauses. It missed E14: “We keep your information for as long as necessary for business purposes.” All 27 accepted predictions had a valid verbatim citation; no proposed citation was rejected in this run. The 40 individual model calls completed in about 13 seconds on the development machine.

These are **in-sample, synthetic results**. We used failures from this very set to revise the prompt, so these percentages are not an independent estimate of accuracy on real policies. Long documents, multiple practices in one clause, indirect wording, exceptions, and adversarial text require a separate held-out evaluation. The current classifier can still miss practices or misread scope, and exact citation matching alone does not establish that a finding is semantically true.

## Next evaluation

Freeze this prompt, collect a new labeled set from diverse public policies with permission to quote excerpts, have a second person review labels, and report per-language precision and recall with examples of disagreements. Keep this synthetic set as a regression check.
