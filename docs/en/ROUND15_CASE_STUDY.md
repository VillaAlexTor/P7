# Round 15: complete Dropbox and GitHub privacy policies

## Sources and method

On November 25, 2026, p7 analyzed the complete current official [Dropbox EN](https://www.dropbox.com/en/privacy), [Dropbox ES](https://www.dropbox.com/es/privacy), [GitHub EN](https://docs.github.com/en/site-policy/privacy-policies/github-general-privacy-statement), and [GitHub ES](https://docs.github.com/es/site-policy/privacy-policies/github-general-privacy-statement) pages. They are **two policies in two languages**, not four independent policies. Dropbox displays a May 30, 2025 publication/effective date and links to a *future* January 1, 2027 version; we evaluated the current text, not that future version. GitHub displays an April 27, 2026 effective date. Extracts were 17,774–47,638 characters. Complete source text remains in gitignored `.cache/real_policies_round15`; the [frozen manifest](../../apps/api/evaluation/real_policies_round15.json) fixes URLs, exact hashes, lengths, reference lines, labels, and eight audit questions. It was committed as `ea985ab` **before inference**.

One reviewer assigned the labels. Dropbox says users may let it access and store contacts, but its current notice does not expressly say personal data selects ads or trains AI models. Its third-party advertising cookies promote Dropbox services; we did not infer targeting from that alone. GitHub explicitly says Personal Data may be used to train models and describes interest-based advertising, but does not affirm device address-book access. Neither notice affirms provider-run biometric identity recognition. We ran the unchanged local `qwen3:4b-instruct` at temperature zero with the general profile, same prompts and Go pipeline, once per page. The [machine report](../../apps/api/evaluation/real_policies_round15_result.json) preserves final categories, clause IDs, evidence-quote hashes, candidate counts, scores, and timings without republishing third-party text.

## Document-category results

| Full page | Frozen reference | p7 categories | TP | FP | FN | Time |
| --- | --- | --- | ---: | ---: | ---: | ---: |
| Dropbox EN | location, contacts, third parties, retention | those four + **advertising** | 4 | 1 | 0 | 25 s |
| Dropbox ES | same four | those four + **advertising** | 4 | 1 | 0 | 28 s |
| GitHub EN | location, third parties, advertising, AI training, retention | first four + **contacts**; **retention missing** | 4 | 1 | 1 | 62 s |
| GitHub ES | same five | exactly those five | 5 | 0 | 0 | 75 s |
| **Total** | | | **17** | **3** | **1** | **192 s** |

Category precision was **17/20 = 85.0%**, recall **17/18 = 94.4%**, and **1/4** localized pages had the exact category set. There were 28 document-category decisions: 17 TP, 3 FP, 1 FN, and 7 true negatives. The model proposed 54 findings, retained 20 literal-quote matches, and counted 34 rejections. That count combines duplicates with invalid evidence. Risk scores were 90, 94, 98, and 96 in page order; they are not calibrated probabilities.

## Eight questions fixed before inference

The eight questions yield 16 language-specific checks that overlap the 28 category decisions. “Not established” means this notice does not affirm the defined practice.

| Question | Reference | p7 EN | p7 ES | Supporting distinction |
| --- | --- | --- | --- | --- |
| D1. Does Dropbox optionally access and store users' contacts? | Yes | Yes | Yes | EN line 23 / ES line 25 says access and server storage after user choice. |
| D2. Does the current Dropbox notice affirm personal-data-based ad targeting? | Not established | **Yes** | **Yes** | EN line 29 / ES line 31 mentions cookies for promotion, not targeting. |
| D3. Does it affirm training AI on user data? | Not established | No | No | No affirmative model-training disclosure in this current notice. |
| D4. Does it retain information while an account exists? | Yes | Yes | Yes | EN line 79 / ES line 81; account deletion starts a separate process. |
| G1. Does GitHub say Personal Data can train models? | Yes | Yes | Yes | Line 99 explicitly ties Personal Data to model training. |
| G2. Does GitHub affirm importing a device address book? | Not established | **Yes** | No | EN cites ordinary identifier/contact-information categories, not a device book. |
| G3. Does GitHub describe interest-based advertising through cookies? | Yes | Yes | Yes | Line 59 explicitly says interest-based advertising. |
| G4. May GitHub share Personal Data with outside partners? | Yes | Yes | Yes | Line 123 affirms disclosure to partners and resellers. |

p7 matched **13/16** localized question checks; these are not independent of the category result. The fourth category error, GitHub EN retention, was outside the eight preselected questions. Its source line 195 says GitHub retains Personal Data while an account is active and for contractual or legal needs. We did not add a ninth question after seeing the miss.

## Citation audit and failure modes

We inspected the 20 retained clause locations separately from the document-level categories. **At least four cited clauses do not by themselves establish the assigned finding**:

| Page | Finding and clause | Why insufficient |
| --- | --- | --- |
| Dropbox EN | `advertising`, `clause-023` | Describes promotional emails about paid upgrades, not personal-data-based ad targeting. |
| Dropbox ES | `advertising`, `clause-021` | Describes third-party cookies to promote Dropbox, without saying personal data selects ads. |
| GitHub EN | `contacts`, `clause-204` | Lists identifiers and contact information collected; no device address book. |
| GitHub ES | `third-parties`, `clause-039` | Describes data **received from** other users, not disclosure **to** outside partners. The category is present elsewhere (source line 123), but this citation is weak. |

The first three are category false positives; the fourth is a category true positive with a weak citation. Other retained locations appeared supportive in this one-reviewer audit. p7 checks that quotes occur literally in the cited clause, not that the clause semantically proves the finding. This audit is not an independent second set of labels. Dropbox's optional `contacts` reference is also open to a stricter reading: the notice says “access to your contacts” but does not spell out “device address book” in those words. A second reviewer should adjudicate that before treating recall for this category as settled.

Across complete-policy rounds ten through fifteen, descriptive category totals are **102 TP, 18 FP, 8 FN** over 24 localized pages representing 12 underlying policies: **85.0% precision and 92.7% recall**. Selected publishers, correlated translations, one reviewer, and uneven category coverage prevent a general reliability claim. The model and prompts did not change; differences across rounds are not training progress.

Next, require evidence that itself affirms the category's defining action, distinguish promotion from personalized ad selection, and instrument rejection reasons. Preserve these pages as a closed holdout and test any fix on new policies. To reproduce while pages remain unchanged, run `python scripts/prepare-real-policies-round15.py --fetch` from the repository root and compare hashes with the manifest. With Ollama and `qwen3:4b-instruct`, run from `apps/api`:

```powershell
go run ./cmd/evaluatepolicies -manifest evaluation/real_policies_round15.json -text-dir ../../.cache/real_policies_round15
```
