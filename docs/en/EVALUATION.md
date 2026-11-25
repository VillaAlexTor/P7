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

## Second evaluation: 40 new clauses

We froze the current prompt and created `evaluation/heldout.tsv` after the prompt revisions above. This second synthetic set contains 20 English and 20 Spanish clauses: 28 affirmative cases, two of which have two labels, plus 12 negative controls. The clauses test conditional consent, indirect wording, outside processors, backups, and mixed practices. We did not tune the prompt against this set before recording the result.

Run from `apps/api`:

```powershell
go run ./cmd/evaluate -data evaluation/heldout.tsv
```

| Set | Correct labels | False positive labels | Missed labels | Precision | Recall |
| --- | ---: | ---: | ---: | ---: | ---: |
| First set after tuning (in-sample) | 27 | 0 | 1 | 100.0% | 96.4% |
| Second set (held out from prompt tuning) | 19 | 0 | 11 | 100.0% | 63.3% |

All 12 negative controls were classified without a finding. No proposed citation was rejected. The 40 calls took about 11 seconds on the development machine. Category results on the second set:

| Category | Correct | Missed |
| --- | ---: | ---: |
| Location | 4 | 0 |
| Contacts | 1 | 3 |
| Third parties | 1 | 5 |
| Advertising | 4 | 0 |
| AI training | 4 | 0 |
| Biometric | 4 | 0 |
| Retention | 1 | 3 |

For example, H02 (conditional contact import), H03 (fraud contractor), and P07 (indefinite storage of inactive account records) were missed. H04 and P04 received the advertising label but missed the simultaneous third-party disclosure. These are material gaps for a tool that helps people interpret privacy terms. The zero false positive count applies only to these 40 synthetic clauses, and the labels were assigned by one reviewer. This is an independent check of the prompt on new synthetic cases, not an estimate of performance on real policies or full documents.

## Third evaluation: 40 more clauses

We kept the same `qwen3:4b-instruct` model and prompt. The third synthetic set, `evaluation/heldout_2.tsv`, was written before running the classifier and contains another 20 English and 20 Spanish clauses. Each language has 14 affirmative cases, including two cases with two labels, and six negative controls. It tests conditional access, subcontractors, combined practices, future hypotheticals, discontinued features, and fixed deletion periods. The labels were assigned by one reviewer.

Run from `apps/api`:

```powershell
go run ./cmd/evaluate -data evaluation/heldout_2.tsv
```

| Set | Correct labels | False positive labels | Missed labels | Precision | Recall |
| --- | ---: | ---: | ---: | ---: | ---: |
| First, after tuning on the same set | 27 | 0 | 1 | 100.0% | 96.4% |
| Second, new synthetic set | 19 | 0 | 11 | 100.0% | 63.3% |
| Third, new synthetic set | 25 | 0 | 7 | 100.0% | 78.1% |

The 12 negative controls produced no findings, and no citation was rejected. The 40 calls took about 18 seconds on the development machine. The English cases had 12 correct and four missed labels; the Spanish cases had 13 correct and three missed labels.

| Category | Correct | Missed |
| --- | ---: | ---: |
| Location | 3 | 1 |
| Contacts | 4 | 0 |
| Third parties | 5 | 1 |
| Advertising | 4 | 2 |
| AI training | 4 | 0 |
| Biometric | 4 | 0 |
| Retention | 1 | 3 |

T07, T14, and S07 were missed retention practices. T04 and S04 each lost one of two expected labels, as did T08 and S08. The third set confirms that broad retention language and clauses with two simultaneous practices remain weak spots. These figures describe only this small synthetic sample; they do not validate real policy analysis or whole-document performance.

## Next evaluation

Review the misses across the second and third sets, then collect a labeled set from diverse public policies with permission to quote excerpts. Have a second person review labels and report per-language precision and recall with disagreements. Keep all three synthetic sets as regression checks; any prompt change must be evaluated on a new untouched set.
