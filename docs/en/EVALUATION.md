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

## November 25: forty more real-source tests

We prepared **40 new verbatim excerpts** from 40 previously unused official URLs: 20 English and 20 Spanish. We split them in advance into a 20-case development set and a separate 20-case holdout, each balanced by language. Text, reference labels, and source URLs were committed before either first inference. The metadata checker now covers 180 cases and 180 source links, verifies that the 40 new URLs are distinct, and limits each source quote to 25 words. An older development set reuses one source URL for two cases. Reference labels were not changed after inference.

We again changed **prompts only**, not the local `qwen3:4b-instruct` model weights. The development set revealed two false biometric findings where fingerprint or face data stayed on the user's device and two missed indefinite-retention findings where chats remained saved until a user deleted them. The first prompt revision fixed three of those four errors; a second device-only biometric clarification fixed the last in the development rerun. We froze revision `b821d56` before running the holdout.

| Run | Correct labels | False positives | Missed labels | Precision | Recall | Interpretation |
| --- | ---: | ---: | ---: | ---: | ---: | --- |
| `real_round6_dev.tsv`, first run with prior prompt | 12 | 2 | 2 | 85.7% | 85.7% | Development baseline |
| `real_round6_dev.tsv`, after first revision | 14 | 1 | 0 | 93.3% | 100.0% | Development rerun |
| `real_round6_dev.tsv`, after second revision | 14 | 0 | 0 | 100.0% | 100.0% | Development rerun |
| `real_round4.tsv`, revision `b821d56` | 18 | 1 | 2 | 94.7% | 90.0% | Prior-set regression rerun |
| `real_round5.tsv`, revision `b821d56` | 17 | 0 | 1 | 100.0% | 94.4% | Prior-set regression rerun |
| **`real_round6_holdout.tsv`, first run with frozen `b821d56`** | **16** | **0** | **1** | **100.0%** | **94.1%** | **Independent first run** |

The holdout's five negative controls produced no findings. R6HE01 was the sole miss: its short excerpt describes someone uploading address-book contacts but omits the service name; the reference label relies partly on source-page context. A second reviewer should assess that label. We kept the frozen wording and score. In the fifth-set regression rerun, R5E08's third-party subprocessor label was missed despite being found in its original run. Output can vary between runs at temperature zero; the perfect fifth-round score must not be treated as a stable guarantee. Earlier round-four misses also remain visible.

The holdout exceeds the selected 90% precision and recall thresholds **on these 20 selected excerpts**. The combined 40-case development-plus-holdout result is 30 correct labels, zero false positives, and one miss, but it is not an independent 40-case estimate because half the cases guided the prompt. One reviewer labeled the samples; short clauses and repeated publishers do not establish performance on whole policies. Reproduce from `apps/api` with Ollama running:

```powershell
go run ./cmd/evaluate -data evaluation/real_round6_dev.tsv
go run ./cmd/evaluate -data evaluation/real_round6_holdout.tsv
```

## November 25: forty further real-source checks

We froze the `b821d56` prompt and the installed `qwen3:4b-instruct` model. We selected 40 short excerpts from 40 new official-page URLs, split evenly between English and Spanish and between `real_round7a.tsv` and `real_round7b.tsv`. The labels and source links were committed before inference. The metadata checker covers 220 cases across all real-source sets and enforces at most 25 quoted words per URL. Some English and Spanish URLs are translations of the same underlying article, so 40 URLs do **not** mean 40 independent policies.

| Run | Correct labels | False positives | Missed labels | Precision | Recall | Exact cases |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| `real_round7a.tsv`, first run | 8 | 1 | 2 | 88.9% | 80.0% | 17/20 |
| `real_round7b.tsv`, first run | 6 | 4 | 4 | 60.0% | 60.0% | 13/20 |
| **Both sets together, first runs** | **14** | **5** | **6** | **73.7%** | **70.0%** | **30/40** |
| `real_round7b.tsv`, quote-corrected rerun | 6 | 4 | 4 | 60.0% | 60.0% | 13/20 |

The English cases had 9 correct labels, 2 false positives, and 2 misses (81.8% precision and recall; 16/20 exact cases). Spanish had 5 correct labels, 3 false positives, and 4 misses (62.5% precision, 55.6% recall; 14/20 exact cases). Across both sets, 10 cases differed from their reference labels. The first runs took 23 and 21 seconds respectively on the local machine; the corrected rerun took 22 seconds.

The misses were R7AE10 (archived chats remain saved), R7AS05 (location-based travel notifications), R7BE07 (conditional location access by My AI), R7BS06 (location badges), R7BS09 (customer lists for ad audiences), and R7BS10 (ad personalization from broader ChatGPT signals). False positives were R7AS07 and R7BE09 (trusted-contact details mistaken for device address-book access), R7BE03 and R7BS08 (device biometric sign-in mistaken for provider biometric collection), and R7BS09 (advertising audiences mistaken for device contacts). R7AS05, R7BS06, and R7BS10 use short fragments whose interpretation depends partly on page context; a second annotator should review them. The reference labels were not changed after seeing model output.

After the first runs, we discovered that R7BS01's Spanish quotation matched another regional version but not its recorded `es-es` URL. We corrected **only its wording** to the literal text on that recorded official page and reran set B. Its prediction and scores were unchanged. The first-run B result is retained separately and is not byte-for-byte reproducible from the corrected TSV. An earlier parser-only change to R7AS04 removed quote marks before any successful model run. The raw reports are in `real_round7a_result.json`, `real_round7b_initial_result.json`, and `real_round7b_result.json` beside the datasets.

**Candid assessment:** this new round falls well below 90% precision and recall. The prior round-six holdout's 100%/94.1% result was sample-specific and does not generalize to these selected clauses. The present classifier is useful as a review aid, but its findings should require human checking; we should not claim it is 90% reliable on unseen policies. No model weights were trained in this round, and the prompt was not tuned after these results. These short excerpts still do not test complete policy documents. Reproduce with Ollama running:

```powershell
cd apps/api
go run ./cmd/evaluate -data evaluation/real_round7a.tsv
go run ./cmd/evaluate -data evaluation/real_round7b.tsv
```

## November 25: eighth real-source round

We tested the unchanged `b821d56` prompt with the local `qwen3:4b-instruct` model on another **40 short verbatim excerpts from 40 previously unused official-page URLs**: 20 English and 20 Spanish. The text, reference labels, and URLs were committed before the first inference. The source checker now covers 260 cases and 260 source links, and limits each source quotation to 25 words. Some URLs are language versions of the same underlying article, so this is not 40 independent privacy policies. No prompt or model-weight change was made after these results.

| First run | Correct labels | False positives | Missed labels | Precision | Recall | Exact cases |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| `real_round8a.tsv` | 16 | 0 | 4 | 100.0% | 80.0% | 16/20 |
| `real_round8b.tsv` | 8 | 1 | 0 | 88.9% | 100.0% | 19/20 |
| **Both new sets** | **24** | **1** | **4** | **96.0%** | **85.7%** | **35/40** |

English had 12 correct labels, one false positive, and two misses: 92.3% precision, 85.7% recall, and 17/20 exact cases. Spanish had 12 correct labels, no false positives, and two misses: 100.0% precision, 85.7% recall, and 18/20 exact cases. The runs took 28 and 21 seconds on the local machine.

The four missed labels were R8AE08 (ad personalization from account data), R8AE09 (location information with a fixed deletion deadline), R8AS05 (the external provider processing facial data), and R8AS08 (activity saved in an account without a deletion deadline). R8BE08 was the sole false positive: Windows Hello sign-in was treated as biometric collection by the provider. R8AE09 is a context-sensitive reference label because the short excerpt describes deletion of location information rather than its initial collection; independent review should adjudicate it. No affirmative AI-training clause appears in this round, although negative training controls do.

The previous 40-case round measured 73.7% precision and 70.0% recall. This higher result is **sample variation under the same classifier**, not evidence that training improved it. Recall remains below 90% in the new round. The labels were assigned by one reviewer, and selected snippets do not establish reliability on complete policies. Keep human review for product findings. The committed machine reports are `real_round8a_result.json` and `real_round8b_result.json` beside the case and source TSVs. Reproduce with Ollama running:

```powershell
cd apps/api
go run ./cmd/evaluate -data evaluation/real_round8a.tsv
go run ./cmd/evaluate -data evaluation/real_round8b.tsv
```

## November 25: ninth real-source round

We froze the prompt from `b821d56` and ran the installed local `qwen3:4b-instruct` model once on **40 previously unseen excerpts from 40 new official URLs**. Cases, reference labels, and URLs were committed before inference (`9b973ee`–`e74f664`). The sources are substantive privacy policies and product notices from Google, Spotify, Microsoft, Mozilla, and Apple. Each set contains ten English and ten Spanish excerpts. Source text is quoted verbatim in spans of at most 25 words; the checker now covers **300 labeled excerpts and 300 source links**. Some URLs are translations of one underlying notice, so this is not a sample of 40 independent policies.

| First run | True positives | False positives | Missed labels | Precision | Recall | Exact cases | Time |
| --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `real_round9a.tsv` | 19 | 0 | 1 | 100.0% | 95.0% | 19/20 | 30 s |
| `real_round9b.tsv` | 14 | 0 | 10 | 100.0% | 58.3% | 11/20 | 24 s |
| **Combined** | **33** | **0** | **11** | **100.0%** | **75.0%** | **30/40** | **54 s** |

English: 17 correct labels, no false positives, five missed labels, **77.3% recall and 15/20 exact cases**. Spanish: 16 correct labels, no false positives, six missed labels, **72.7% recall and 15/20 exact cases**. Precision is 100% in each language *for this small selected sample*, not a general guarantee. By category, advertising was 16/16, AI training 2/2, location 9/10, retention 2/6, and third-party sharing 4/10. Contacts and biometrics were absent and cannot be assessed from this round.

The ten mismatched cases were R9AE09; R9BE03, R9BE04, R9BE07, R9BE10; and R9BS03, R9BS04, R9BS06, R9BS07, R9BS10. The model missed explicit retention language in both Mozilla policy excerpts and the 24-hour deletion in Apple's Find My notice. It also missed several outside processors, including card issuers and health-data recipients. R9AE09 and both Google Analytics excerpts (R9BE03/R9BS03) are somewhat context-sensitive: the short quotes identify third-party cookies or a named external analytics service, but do not spell out every data transfer. We retained the precommitted labels and flag those three for independent adjudication. The other misses remain visible in the raw reports.

**Current state:** this round missed the 90% recall target by a wide margin, despite zero false positives. Round eight's 96.0% precision and 85.7% recall and round seven's 73.7%/70.0% show substantial sample variation under the same classifier. P7 is a usable privacy-review aid with verifiable citations, but findings still need human review and omissions remain possible. We did not retrain model weights, edit the prompt, or change reference labels after seeing these results. One reviewer labeled selected short excerpts; neither inter-reviewer agreement nor whole-document performance has been established.

The raw machine output, including every prediction and rejected-citation count, is committed in `apps/api/evaluation/real_round9a_result.json` and `real_round9b_result.json`. Reproduce with Ollama and `qwen3:4b-instruct` running:

```powershell
cd apps/api
go run ./cmd/evaluate -data evaluation/real_round9a.tsv
go run ./cmd/evaluate -data evaluation/real_round9b.tsv
```

## November 25: tenth round, complete real privacy policies

This run changed the unit of evaluation from selected excerpts to **entire published notices**. We downloaded the complete `<main>` content of the [Mozilla Accounts notice](https://www.mozilla.org/en-US/privacy/mozilla-accounts/) and the [Proton policy](https://proton.me/legal/privacy), plus their [Mozilla Spanish](https://www.mozilla.org/es-ES/privacy/mozilla-accounts/) and [Proton Spanish](https://proton.me/es-419/legal/privacy) versions. The extracted texts were 16,843–20,873 characters and 130–157 p7 clauses each. We checked the first and last sections and retained all main-content sections. The local extracts stay in gitignored `.cache/real_policies`; the [frozen manifest](../../apps/api/evaluation/real_policies_round10.json) records each official URL, exact SHA-256, text length, reference categories, and reference line numbers. It was committed as `4eed7b5` **before inference**. These publishers appeared in earlier excerpt sets, so this is a new *full-document test*, not publisher-independent holdout data.

We ran the unchanged `qwen3:4b-instruct` and p7 `AnalyzeWithClassifier` pipeline once on each complete text, using the general profile. The seven categories were scored as present or absent **per document**, with one possible finding per category. A single reviewer assigned reference labels. The raw [machine report](../../apps/api/evaluation/real_policies_round10_result.json) preserves every document's predicted categories, score, clause IDs, citation verification counts, source hash, and timing. Evidence quotes are represented by hashes in the committed report so the third-party policy text is not republished.

| Policy | Reference labels | p7 labels | Exact? |
| --- | --- | --- | --- |
| Mozilla EN | location, third parties, retention | location, contacts, advertising | No |
| Mozilla ES | location, third parties, retention | location, contacts, advertising, retention | No |
| Proton EN | third parties, retention | third parties, advertising | No |
| Proton ES | third parties, retention | third parties, retention | Yes |

**Observed document-category result: 6 true positives, 5 false positives, 4 misses; 54.5% precision, 60.0% recall, and 1/4 documents with an exact category set.** The four runs took 125 seconds total. Twenty-nine proposed citations were rejected by p7's exact-quote verifier; the 11 retained findings had verified literal quotes, but that verification did not establish semantic correctness. The two Mozilla contact findings refer to contact information, not access to a device address book. The three advertising findings do not establish personalized ad targeting under the frozen category definition. The English Mozilla notice's explicit outside recipients and broad retention were missed; the English Proton notice's potentially permanent IP retention was also missed. The Spanish versions sometimes produced different labels for equivalent policies.

These four URLs are **two policies in two languages**, not four independent publishers. Negative labels for biometric and AI training provide limited coverage; there is no positive example of either category. The 0.3 hybrid risk scores (48–74) were generated but were **not validated** against a separate risk reference. The result is much lower than the ninth round's short-excerpt precision and recall and is the best current evidence of how this version behaves on complete policies. Do not present it as 90% reliable or use it for unsupervised privacy decisions. We did not change the model, prompts, code, or labels after this run. Independent label review, more publishers, and targeted fixes followed by a new holdout are needed.

To reproduce the exact local texts, rerun the fetch script while the published pages are unchanged and compare the recorded hashes; publishers may edit them later. Then, from `apps/api` with Ollama serving the model:

```powershell
python scripts/prepare-real-policies.py --fetch
cd apps/api
go run ./cmd/evaluatepolicies -manifest evaluation/real_policies_round10.json
```

## November 25: eleventh round, two different complete policies

We tested the **unchanged** p7 pipeline on the complete [Apple Privacy Policy (English)](https://www.apple.com/legal/privacy/en-ww/), [Apple Privacy Policy (Spanish)](https://www.apple.com/legal/privacy/es/), [Spotify Privacy Policy (UK English)](https://www.spotify.com/uk/legal/privacy-policy/), and [Spotify Privacy Policy (Spain Spanish)](https://www.spotify.com/es/legal/privacy-policy/). Apple and Spotify are different publishers from round ten. We extracted each page's full `<main>` content, checked the beginning and end of the extraction, and kept the 28,051–51,336-character texts in gitignored `.cache/real_policies_round11`. The [manifest](../../apps/api/evaluation/real_policies_round11.json) fixes official URLs, exact SHA-256 hashes, lengths, document-level labels, and supporting line numbers. It was committed as `2b0f619` **before model inference**. Apple and Spotify had appeared in earlier short-excerpt sets, so this is not a publisher-independent holdout across the entire project history.

One reviewer labeled the seven categories using the existing p7 definitions. In particular, contact information alone is not device address-book access. We count Spotify's optional face-photo comparison to an identity document as biometric use; facial age estimation by itself would be ambiguous. Spotify explicitly describes training machine-learning models in a purpose table containing user and usage data. Apple's policy explicitly excludes private personal data from foundational-model training; Applebot's crawling of public web material does not establish training with private Apple-user data. These decisions were written in the manifest before inference and need independent adjudication.

| Complete page | Reference categories | p7 categories | Exact set |
| --- | --- | --- | --- |
| Apple EN | location, third parties, advertising, retention | all four **plus contacts** | No |
| Apple ES | location, third parties, advertising, retention | location, third parties, advertising, **contacts** | No |
| Spotify EN | location, third parties, advertising, AI training, biometric, retention | all six | Yes |
| Spotify ES | location, third parties, advertising, AI training, biometric, retention | all six | Yes |

**First-run result: 19 true positives, 2 false positives, 1 miss; 90.5% precision, 95.0% recall, and 2/4 exact document-category sets.** The four runs took 291 seconds on the local machine. The unchanged `qwen3:4b-instruct` produced 99 candidate findings; p7 retained 21 and rejected 78. Its `citationsRejected` counter includes duplicate-category proposals **as well as** invalid evidence, so these 78 cannot be interpreted as 78 invalid quotes. The committed [machine report](../../apps/api/evaluation/real_policies_round11_result.json) retains every final prediction, finding category and clause ID, evidence-quote hash, metadata, score, and timing without republishing full third-party policy texts.

**Errors to fix:** Apple EN and ES both produced a `contacts` finding from ordinary contact information, repeating round ten's confusion between an email/address field and device address-book access. Apple ES omitted the broad retention period that Apple EN identified. A citation can be literally present yet semantically wrong. The unvalidated risk scores were 90 and 84 for Apple and 100 for both Spotify pages; they should not be described as calibrated risk probabilities. The result suggests a stricter address-book criterion, Spanish retention checks, and rejection-reason instrumentation before another fresh holdout.

The higher result than round ten's 54.5% precision / 60.0% recall is **sample variation under the same model and prompts**, not measured improvement or training. Combining the two full-policy rounds descriptively gives 25 true positives, 7 false positives, and 5 misses (78.1% precision, 83.3% recall) over eight localized pages from four underlying policies; language versions are correlated, and one reviewer supplied every label. No model, prompt, pipeline, or frozen label was changed after this run. Independent labeling and more publisher-independent complete policies are still needed.

To reproduce, fetch the pages from the repository root while their contents are unchanged, verify the recorded hashes, and run the model from `apps/api`:

```powershell
python scripts/prepare-real-policies-round11.py --fetch
cd apps/api
go run ./cmd/evaluatepolicies -manifest evaluation/real_policies_round11.json -text-dir ../../.cache/real_policies_round11
```

## November 25: twelfth round, longer and harder complete policies

We froze the [four source hashes, seven-category labels, and eight audit questions](../../apps/api/evaluation/real_policies_round12.json) in commit `194da5e` before inference. The new full policy pages were [Google EN](https://policies.google.com/privacy?hl=en), [Google ES](https://policies.google.com/privacy?hl=es), [WhatsApp EEA EN](https://www.whatsapp.com/legal/privacy-policy-eea?lang=en_gb), and [WhatsApp EEA ES](https://www.whatsapp.com/legal/privacy-policy-eea?lang=es_LA), totaling 349,926 extracted characters. We ran the same local model, prompts, pipeline, and general profile once per page. The [machine report](../../apps/api/evaluation/real_policies_round12_result.json) records the final predictions without republishing the policies.

| Complete-policy round | True positives | False positives | Misses | Precision | Recall | Exact pages | Time |
| --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| 10 · Mozilla, Proton | 6 | 5 | 4 | 54.5% | 60.0% | 1/4 | 125 s |
| 11 · Apple, Spotify | 19 | 2 | 1 | 90.5% | 95.0% | 2/4 | 291 s |
| **12 · Google, WhatsApp** | **16** | **6** | **0** | **72.7%** | **100.0%** | **1/4** | **646 s** |

Google EN and ES both falsely labeled account contact details as device address-book access and both labeled a public-web AI-training passage as training on private user data under our frozen interpretation. WhatsApp EN falsely inferred ad personalization from ad-related cookies and biometric identification from an age-estimation video selfie. WhatsApp ES had the exact category set. The Google AI reference is contestable: treating public personal material as user training data would raise round-twelve precision to 81.8% without changing recall. We preserve the original labels and result. The **[bilingual, question-by-question case study](ROUND12_CASE_STUDY.md)** explains all eight precommitted questions, the extra biometric error, source-line reasoning, category sensitivity, and the 274-second WhatsApp ES latency. One reviewer assigned the labels; translation pairs are correlated; this is neither independent validation nor evidence of a model improvement.
