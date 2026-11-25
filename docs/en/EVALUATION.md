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

## Iterative improvement and stop point

We improved the **instructions and inference workflow** around the local `qwen3:4b-instruct` model; we did not update its weights or fine-tune it. The one-pass classifier still missed practices in a new 40-clause set, so we split classification into focused category passes. We then separated retention into its own pass. Each version was checked against the development clauses before we wrote a final untouched set.

| Run | Set | Correct labels | False positives | Missed labels | Recall |
| --- | --- | ---: | ---: | ---: | ---: |
| One pass | `heldout_4.tsv`, first run | 22 | 1 | 10 | 68.8% |
| Four passes | `heldout_4.tsv`, development rerun | 27 | 1 | 5 | 84.4% |
| Five passes | `heldout_4.tsv`, development rerun | 30 | 0 | 2 | 93.8% |
| **Five passes, final untouched check** | **`heldout_5.tsv`** | **27** | **0** | **5** | **84.4%** |

The final set was written before its first run: 20 English and 20 Spanish clauses, 28 affirmative cases with four two-label cases, and 12 negative controls. The model found at least one positive in each of the seven categories and returned no finding for any negative control. No accepted finding lacked a verbatim quote; one proposed quote was rejected. Run it from `apps/api` with:

```powershell
go run ./cmd/evaluate -data evaluation/heldout_5.tsv
```

The final misses were X03 and X23 (payment gateways), X06 and X26 (typing rhythm), and X27 (open-ended retention). The payment-gateway wording does not explicitly say the gateway is external, so those two reference labels are open to reviewer disagreement. We kept the original labels and the reported score unchanged. Model outputs also varied between repeated runs despite temperature zero. Five model calls per clause batch cost more time than the original one-pass flow.

We stopped changing the classifier after the untouched set exceeded the chosen **80% recall** threshold with no false positives in this run. This is evidence of progress on synthetic clauses, not 80% statistical confidence in real-world accuracy or readiness for unreviewed decisions.

## Evaluation on excerpts from real policies

On November 25, 2026, we collected 40 short, verbatim excerpts from published privacy policies and official privacy explanations: 20 English and 20 Spanish. The exact text, labels, and source URLs are in `evaluation/real_dev.tsv`, `evaluation/real_dev_sources.tsv`, `evaluation/real_holdout.tsv`, and `evaluation/real_holdout_sources.tsv`. The URLs allow each excerpt to be reviewed against its published page. A single reviewer assigned the labels before model inference. The two sets use different source pages; no prompt, model, or classification-code change was made after either set was assembled. We corrected one UK spelling, shortened one negative excerpt, and replaced one development source that overlapped the holdout, then reran that set. The holdout was run once and left unchanged.

| Real-source set | Correct labels | False positive labels | Missed labels | Precision | Recall |
| --- | ---: | ---: | ---: | ---: | ---: |
| Development, 20 excerpts | 18 | 1 | 1 | 94.7% | 94.7% |
| **Untouched holdout, 20 excerpts** | **17** | **0** | **2** | **100.0%** | **89.5%** |
| Both sets, 40 excerpts | 35 | 1 | 3 | 97.2% | 92.1% |

The holdout had 8/9 correct English labels and 9/10 correct Spanish labels; all three negative controls had no finding. The misses were RF06, Discord's use of IP data to determine a region, and RF16, Discord's periodic synchronization of phone numbers from a device. RD07 in the development set gained an unsupported third-party label, while RD12's broad retention was missed. The frozen five-pass classifier **cleared the 80% recall threshold on real published excerpts without another tuning step**, so we stopped iterating.

Reproduce from `apps/api` with Ollama running:

```powershell
go run ./cmd/evaluate -data evaluation/real_dev.tsv
go run ./cmd/evaluate -data evaluation/real_holdout.tsv
```

This supports the judgment that the project is moving in the right direction. The sample is small, selected rather than random, and labeled by one reviewer; several excerpts come from the same publisher. We evaluated short excerpts individually, not complete policies, and the output of the local model can vary between runs even at temperature zero. These percentages are **not** a measured accuracy guarantee for arbitrary real policies. A second annotator, more publishers, full-document tests, and latency measurements remain necessary before any reliability claim.

## Next evaluation

Before a production claim, expand the real-source sample, have a second person review disputed labels, measure full-document latency, and report per-language precision and recall. Keep the six synthetic and two real-source sets as regression checks; any later prompt change needs a new untouched set.

## November 25 follow-up: 90% target

We strengthened the local `qwen3:4b-instruct` classifier's instructions and split location, contacts, and biometric into separate passes. **This changes the prompt and inference workflow, not the model weights.** We assembled two additional bilingual sets of 20 short excerpts each from published official pages, with one source URL per excerpt. Their texts, labels, and URLs are in `evaluation/real_holdout_90*.tsv`. Both sets are now development data because their errors informed further changes.

| Run | True positive labels | False positive labels | Missed labels | Precision | Recall |
| --- | ---: | ---: | ---: | ---: | ---: |
| First new set, first run | 18 | 3 | 0 | 85.7% | 100.0% |
| Second new set, first run | 13 | 5 | 3 | 72.2% | 81.3% |
| Second set, after seven focused passes | 14 | 1 | 2 | 93.3% | 87.5% |

The first set's first-run excerpt for R9S01 was shortened afterward to meet the 25-word quotation limit. Its label stayed the same; the first-run score cannot be reproduced byte for byte from the committed file. The second set's first-run score was recorded before any tuning on it. The last row is a development rerun, **not an independent holdout result**. Its remaining errors are a missed biometric label in a context-dependent list (RFE09), missed broad retention (RFS08), and an unsupported advertising label when sharing with advertising partners is mentioned (RFS09). A second reviewer should examine the context-dependent RFE09 label.

The observed results do **not** support a 90% claim for both precision and recall on untouched real excerpts. A larger locally installed `qwen3-coder:latest` model was loaded for comparison, but its CPU-heavy inference did not complete the 20-excerpt run in several minutes; we stopped that comparison and did not switch the product model. The next valid check needs another untouched, independently labeled real-source set, followed by full-document and latency tests. Do not infer statistical confidence or accuracy on arbitrary policies from these selected excerpts.

Run the additional sets from `apps/api` with Ollama running:

```powershell
go run ./cmd/evaluate -data evaluation/real_holdout_90.tsv
go run ./cmd/evaluate -data evaluation/real_holdout_90_final.tsv
```

## November 25: fourth real-source round and prompt history

We continued improving **prompts and inference only** with the local `qwen3:4b-instruct`; its weights were not trained or changed. Each set has 10 English and 10 Spanish short excerpts, labels assigned before inference, and separate official-page URLs. The source checker verifies 120 excerpts, 120 distinct URLs, and at most 25 quoted words per URL across all six real-source sets.

| Chronological run | True positive labels | False positive labels | Missed labels | Precision | Recall | Status |
| --- | ---: | ---: | ---: | ---: | ---: | --- |
| `real_round3.tsv`, first run with `cf8884a` | 13 | 1 | 3 | 92.9% | 81.3% | Untouched first run |
| `real_round3.tsv`, after prompt revision `c90c9e9` | 16 | 0 | 0 | 100.0% | 100.0% | Development rerun |
| `real_holdout_90_final.tsv`, same revision | 16 | 1 | 0 | 94.1% | 100.0% | Development rerun |
| **`real_round4.tsv`, first run with frozen `c90c9e9`** | **18** | **2** | **2** | **90.0%** | **90.0%** | **Untouched first run** |

The first round-three run missed R3E01 and R3E03 (open-ended retention) and R3E10 (consented biometric use); it added advertising to R3S10, which only described a transfer to advertising partners. We revised the prompt to recognize purpose-based retention and optional biometric use while requiring an affirmative claim of ad targeting. The revised results on round three and the earlier set are development results because those clauses had already informed the instructions. The older set still has an unsupported advertising finding at RFS09.

We committed the round-four text and source links **before** its first inference, then ran the frozen classifier once. R4E01 missed contact import because the proposed supporting citation was rejected; R4E04 missed ad targeting by outside parties. R4E05 incorrectly called Apple Ads an outside recipient, and R4S06 inferred ad targeting from a transfer to advertising partners. English had 9 true positives, 1 false positive, and 2 misses; Spanish had 9 true positives, 1 false positive, and no misses. The three negative controls produced no findings. We did not alter the labels or prompt after seeing these results.

The observed 90% precision and recall meet the chosen stop point **for this 20-excerpt, selected holdout only**. They do not measure statistical confidence, whole-policy performance, or general reliability. One reviewer labeled the excerpts, and several publishers appear in multiple sets. A second reviewer, more independent publishers, full-document evaluation, and latency measurements are still needed. To reproduce the last run from `apps/api` with Ollama running:

```powershell
go run ./cmd/evaluate -data evaluation/real_round4.tsv
```

## November 25: fifth real-source round

We froze the same `c90c9e9` prompt and inference workflow, then committed `real_round5.tsv` and `real_round5_sources.tsv` before the model's first run. This new set contains 10 English and 10 Spanish excerpts from 20 official pages not used in the preceding real-source sets. It has 18 affirmative labels, including one two-label case, and three negative controls. The source checker now covers 140 excerpts and 140 distinct URLs, with no more than 25 quoted words per URL.

| Run | True positive labels | False positive labels | Missed labels | Precision | Recall |
| --- | ---: | ---: | ---: | ---: | ---: |
| `real_round5.tsv`, frozen prompt, first run | 18 | 0 | 0 | 100.0% | 100.0% |

All three negative controls had no finding. We made **no prompt or model change** after this run: the new sample supplied no observed error to justify another targeted revision. This is another evaluation of the existing local model and prompts, not training of model weights. The previous round-four first-run score remains 90.0% precision and 90.0% recall; its four errors have not been erased by the new result.

The fifth set is small, selected, and labeled by one reviewer. Its perfect observed score is not a general accuracy estimate or evidence of performance on whole documents. Future work should prioritize independent annotation, broader publishers, full-document testing, and repeatability and latency measurements. To reproduce from `apps/api` with Ollama running:

```powershell
go run ./cmd/evaluate -data evaluation/real_round5.tsv
```
