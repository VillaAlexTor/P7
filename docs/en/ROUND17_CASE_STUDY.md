# Round 17: IKEA Spain's full privacy policy

## Sources and frozen reference

On November 25, 2026, p7 analyzed IKEA Spain's official [English](https://www.ikea.com/es/en/customer-service/privacy-policy/) and [Spanish](https://www.ikea.com/es/es/customer-service/privacy-policy/) privacy pages, dated August 4, 2026. These are **two translations of one policy**, not two independent organizations. The extracted full texts contain 43,120 and 47,845 characters. [Source preparation](../../scripts/prepare-real-policies-round17.py) stores downloaded HTML and text in gitignored `.cache/real_policies_round17`; the [frozen manifest](../../apps/api/evaluation/real_policies_round17.json) records URLs, hashes, lengths, seven category labels per page, and seven questions in both languages. The manifest was committed as `b671ffb` **before inference**. One reviewer assigned the reference labels.

The reference marks `location`, `third-parties`, `advertising`, and `retention` positive in each language, and `contacts`, `ai-training`, and `biometric` negative. The policy describes optional app location access (extracted line 111), disclosure to external service providers (EN 251 / ES 253), and marketing tailored to purchases and browsing (EN 167 / ES 169). It describes AI analysis of call transcripts (EN 137 / ES 139) but does not affirm model training. It counts visitors with cameras (EN 165 / ES 167) without affirming identity recognition. Event data is retained for a legal-purpose period with no definite deadline (EN 329 / ES 331). Gift Registry recipients are entered by the user, not imported from a device address book.

## Two transparent passes

The first pass used the revised prompt developed on round 16. The [initial machine output](../../apps/api/evaluation/real_policies_round17_result_initial.json) is preserved. After examining its errors, we narrowed the contact and finite-expiry wording and ran a **second pass on the same documents**. The [second machine output](../../apps/api/evaluation/real_policies_round17_result.json) is therefore an *in-sample repair check*, **not a fresh holdout**. Both used local Ollama `qwen3:4b-instruct`, temperature zero, the general profile, full-document processing, and one retained finding per category per document. Neither run trained model weights.

| Pass | IKEA EN | IKEA ES | TP | FP | FN | Precision | Recall |
| --- | --- | --- | ---: | ---: | ---: | ---: | ---: |
| Initial | all four positives + false `contacts` | all four positives + false `contacts` | 8 | 2 | 0 | 80.0% | 100% |
| Revised | misses `retention`, still false `contacts` | exact four-category set | 7 | 1 | 1 | 87.5% | 87.5% |

Both passes match **12 of 14** localized category-question checks. The revised pass has one exact page out of two. These percentages concern **category presence only**, across correlated translations and one policy. They are not deployment reliability or calibrated confidence.

## Evidence audit

We inspected every retained clause location in both reports. Only **5 of 10** initial citations and **5 of 8** revised citations independently support their category. Both passes correctly detect advertising in English at the *document* level, but cite `clause-069`, a generic member offer. The actual tailored-marketing passage is at extracted line 167. Both passes correctly detect Spanish retention at the document level, but cite `clause-041`, a stored payment card that expires; the open-ended legal-purpose passage is at line 331. The initial English retention citation has the same finite-card defect. Initial contacts findings cite customer-service “Contact” passages; the revised English one cites store “contact points” in `clause-222`. None establishes device address-book access. The revised English retention miss shows that tightening the prompt traded one false positive for one false negative. Location and outside-sharing citations are substantively relevant in both languages; the revised Spanish advertising citation is relevant too.

p7's automatic `evidenceVerified` field means **literal quote containment**, not semantic support. This audit is a separate human reading of the cited clauses. The model can still select an early weak clause when a stronger passage exists elsewhere. These results show a prompt-level repair is insufficient for dependable evidence selection; a separately validated citation-selection stage and an independent new-policy holdout are the next concrete improvements. Do not claim 90% reliability from this sample.

## Reproduction

The cached source text has SHA-256 hashes in the manifest. If the official pages change, `python scripts/prepare-real-policies-round17.py --fetch` creates a **new source version**, not a reproduction of this frozen run. With the original cached HTML, run `python scripts/prepare-real-policies-round17.py`, compare hashes, start Ollama with `qwen3:4b-instruct`, then from `apps/api` run:

```powershell
go run ./cmd/evaluatepolicies -manifest evaluation/real_policies_round17.json -text-dir ../../.cache/real_policies_round17
```

This command uses the current prompt and may differ from either recorded pass if prompts or model behavior change. The two checked-in JSON reports are the raw outputs for this study.
