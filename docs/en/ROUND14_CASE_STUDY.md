# Round 14: complete LinkedIn and Zoom privacy policies

## Scope and frozen reference

On November 25, 2026, p7 analyzed the complete official [LinkedIn EN](https://www.linkedin.com/legal/privacy-policy), [LinkedIn ES](https://es.linkedin.com/legal/privacy-policy), [Zoom EN](https://www.zoom.com/en/trust/privacy/privacy-statement/), and [Zoom ES](https://www.zoom.com/es/trust/privacy/privacy-statement/) pages. These represent **two underlying policies in two languages**, not four independent publishers. The extracted texts were 42,150–64,369 characters. LinkedIn displayed an effective date of November 3, 2025; Zoom displayed an update date of September 17, 2026. The full source texts remain in gitignored `.cache/real_policies_round14`; the published [manifest](../../apps/api/evaluation/real_policies_round14.json) identifies their exact SHA-256 hashes, lengths, supporting lines, labels, and eight audit questions. It was committed as `80f8536` **before model inference**.

One reviewer froze the seven category labels. `contacts` requires explicit **device address-book** access or upload: LinkedIn mentions address-book uploads, while Zoom mentions contact lists and third-party app integration without affirming device-book access. LinkedIn explicitly says it may train AI models with personal data. Zoom says it does not train its or third-party AI models on communications-like Customer Content; its AI features processing content is a different practice. Zoom's optional voiceprints and facial geometry plus its stated authentication purpose count as `biometric`; device-only face/hand effects do not. This biometric interpretation warrants independent review. We ran the unchanged local `qwen3:4b-instruct`, temperature zero, general profile, prompts, and Go pipeline once per page. The [compact machine report](../../apps/api/evaluation/real_policies_round14_result.json) contains final categories, clause IDs, evidence-quote hashes, candidate counts, scores, and timing without republishing the notices.

## Document-category results

| Complete page | Reference categories | p7 categories | TP | FP | FN | Time |
| --- | --- | --- | ---: | ---: | ---: | ---: |
| LinkedIn EN | location, contacts, third parties, advertising, AI training, retention | all except **contacts** | 5 | 0 | 1 | 68 s |
| LinkedIn ES | same six | exactly those six | 6 | 0 | 0 | 100 s |
| Zoom EN | location, third parties, advertising, biometric, retention | those five **plus contacts** | 5 | 1 | 0 | 114 s |
| Zoom ES | same five | those five **plus contacts** | 5 | 1 | 0 | 149 s |
| **Total** | | | **21** | **2** | **1** | **433 s** |

Document-category precision was **21/23 = 91.3%**, recall **21/22 = 95.5%**, and **1/4** localized pages had the exact category set. The four pages give 28 category decisions, including four correct negative decisions. The model proposed 103 findings, retained 23 literal-quote matches, and counted 80 rejections. The rejection count mixes duplicate proposals with invalid evidence. Risk scores were 92, 100, 100, and 100 in page order; they are not calibrated probabilities. A seven-minute-plus batch is a real usability limitation.

## Eight questions fixed before inference

The questions yield 16 language-specific checks, overlapping the 28 document-category decisions above. “Not established” means this notice does not affirm the defined practice.

| Question | Reference | p7 EN | p7 ES | Why it is difficult |
| --- | --- | --- | --- | --- |
| L1. Does LinkedIn use device address-book uploads for connection suggestions? | Yes | **No** | Yes | The explicit upload is at line 173; calendar sync and ordinary contact information are different. |
| L2. Does LinkedIn say it may train AI models with personal data? | Yes | Yes | Yes | Line 161 states both the data source and training purpose. |
| L3. Does LinkedIn serve tailored ads? | Yes | Yes | Yes | Line 207 affirms personalized ads. |
| L4. Does LinkedIn affirm biometric identity recognition? | Not established | No | No | Age/gender inferences are not biometric identity recognition. |
| Z1. Does Zoom train AI models on customer meetings, chat, or audio? | Not established | No | No | Line 63 explicitly denies this use; AI feature processing at line 73 is not training. |
| Z2. Does Zoom process optional voiceprints or facial geometry for features including authentication? | Yes | Yes | Yes | Lines 71 and 335 describe unique physical features and authentication; local-only effects at line 69 are distinct. |
| Z3. Do Zoom contact lists establish device address-book access? | Not established | **Yes** | **Yes** | Line 39 names contact lists and third-party app integration, not the device address book. |
| Z4. Does Zoom disclose ads targeted by website use? | Yes | Yes | Yes | Line 97 explicitly mentions targeted advertising based on website use. |

p7 matched **13/16** question checks under the frozen mapping. The LinkedIn EN miss and two Zoom contact false positives are the three disagreements. The same underlying policies appear in both languages, so these are not 16 independent observations.

## Separate audit of cited evidence

The category score above gives credit when a practice appears **somewhere in the full notice**. It does not establish that the particular clause p7 cites supports that practice. We inspected all 23 retained clause locations. At least **nine cited locations do not by themselves substantiate the assigned finding**, including both false-positive Zoom contact citations:

| Page and finding | Cited clause | Why the cited location is insufficient |
| --- | --- | --- |
| LinkedIn ES `contacts` | `clause-043` | Calendar/email synchronization and contact details; actual address-book passage is later at source line 173. The category is present, but this citation is weak. |
| LinkedIn ES `third-parties` | `clause-044` | Describes data **received from** customers/partners, not sharing data **to** an outside recipient. Sharing is affirmed elsewhere at line 345. |
| Zoom EN/ES `contacts` | `clause-022` in both | Contact lists and third-party app integration do not establish device-book access. These are category false positives as well. |
| Zoom EN/ES `third-parties` | `clause-049` / `clause-050` | Describes information **received from** partners. Vendor access is affirmed elsewhere at line 93. |
| Zoom EN/ES `advertising` | `clause-049` / `clause-050` | Describes information obtained from ad partners, including clicks; it does not itself say personal data selects ads. Targeting is affirmed elsewhere at line 97. |
| Zoom EN `biometric` | `clause-076` | Describes Apple TrueDepth processing face/hand dimensions **on the device** for effects. The provider's optional voiceprint/facial-geometry processing appears elsewhere at lines 71 and 335. |

These are **seven weak citations attached to category-level true positives plus two false-positive contact citations**. The other cited locations appeared supportive on this one-reviewer audit; this is not an independent annotation or a second calibrated metric. The Go verifier checks that a quote is literally contained in its referenced clause, not whether the quote semantically proves the category. That explains how an exact quote can still mislead. The machine report stores quote hashes rather than the third-party text; the clause IDs can be checked against the cached extracts or a fresh download whose hashes match the manifest.

## Interpretation and follow-up

Across complete-policy rounds ten through fourteen, descriptive category totals are **85 TP, 15 FP, 7 FN** over 20 localized pages from ten underlying policies: **85.0% precision and 92.4% recall**. The 91.3% precision in this round should not be read as 91.3% evidence quality. The policies were selected, language versions are correlated, one reviewer assigned all reference labels and audited citations, and no prompt or model changed between these rounds. The Zoom biometric label is particularly worth independent adjudication. These numbers do not establish general 90% reliability or model training progress.

The next fix should require evidence that itself states the category's defining action: device-address-book access for `contacts`, outward transfer for `third-parties`, personal-data-based ad targeting for `advertising`, and provider processing of a biometric identifier for `biometric`. Record separate rejection reasons and rerun on **new policies**, preserving this round as a holdout. The 433-second local runtime also merits profiling.

To reproduce while the published pages are unchanged, run `python scripts/prepare-real-policies-round14.py --fetch` from the repository root and verify the printed hashes. With Ollama and `qwen3:4b-instruct`, run from `apps/api`:

```powershell
go run ./cmd/evaluatepolicies -manifest evaluation/real_policies_round14.json -text-dir ../../.cache/real_policies_round14
```
