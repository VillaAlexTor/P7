# Round 13: optional practices in complete privacy policies

## Method and sources

On November 25, 2026, we evaluated the unchanged p7 pipeline on four official, complete privacy-policy pages: [OpenAI EN](https://openai.com/policies/privacy-policy/), [OpenAI ES](https://openai.com/es-419/policies/privacy-policy/), [TikTok EEA EN](https://www.tiktok.com/legal/page/eea/privacy-policy/en), and [TikTok EEA ES](https://www.tiktok.com/legal/page/eea/privacy-policy/es). These are **two underlying policies**, each tested in two languages. The downloaded OpenAI text displayed a July 30, 2026 update; the TikTok EEA text displayed November 30, 2025. The exact extracts are identified by SHA-256 and character count in the [frozen manifest](../../apps/api/evaluation/real_policies_round13.json). Source pages may change or route by region.

We committed that manifest as `329672d` **before inference**, including seven-category reference labels, supporting line numbers, and eight audit questions. One reviewer assigned the labels. The complete extracted texts (22,484–53,600 characters) stay in gitignored `.cache/real_policies_round13`; we do not republish the third-party policies. The [compact machine report](../../apps/api/evaluation/real_policies_round13_result.json) contains all final categories, clause IDs, evidence-quote hashes, scores, timing, and candidate counts. We used the same local `qwen3:4b-instruct`, temperature zero, general profile, prompts, and Go pipeline as rounds ten through twelve. Each page was analyzed once, with no changes after seeing predictions.

The reference counts an optional practice as present when the policy expressly says it occurs after a user choice or consent. For `contacts`, this means **device address-book access**, not merely having a person's email or phone number. For `biometric`, the policy must affirm an identifier used to recognize a person; face detection for effects or an unspecified identity-verification service is insufficient.

## Results

| Full page | Reference categories | p7 categories | TP | FP | FN | Time |
| --- | --- | --- | ---: | ---: | ---: | ---: |
| OpenAI EN | location, contacts, third parties, advertising, AI training, retention | exactly those six | 6 | 0 | 0 | 40 s |
| OpenAI ES | same six | five; **contacts missing** | 5 | 0 | 1 | 40 s |
| TikTok EEA EN | same six | exactly those six | 6 | 0 | 0 | 85 s |
| TikTok EEA ES | same six | exactly those six | 6 | 0 | 0 | 100 s |
| **Total** | | | **23** | **0** | **1** | **267 s** |

Precision was **23/23 = 100%**, recall **23/24 = 95.8%**, and **3/4** localized pages had an exact category set. The 28 document-category decisions also include four correct `biometric` absences. There are **no positive biometric examples in this round**, so biometric recall is unmeasured. The model proposed 101 candidates, retained 23 findings with exact quotes, and counted 78 rejections. The rejection counter combines duplicates and invalid evidence; it does not identify how many quotes were invalid. Risk scores were 100, 96, 100, and 100 in page order; these are **not calibrated probabilities**.

## Eight questions fixed before inference

Each question is compared against the English and Spanish page. This makes 16 localized question checks, which **overlap** with the 28 document-category decisions above; they are not an independent sample. “Not established” means this notice does not affirm the defined practice, not that the practice can never occur.

| Question | Reference | p7 EN | p7 ES | Evidence or distinction |
| --- | --- | --- | --- | --- |
| O1. Does optional OpenAI contact connection upload the device address book? | Yes | Yes | **No** | OpenAI EN/ES line 33 explicitly says the device address book is uploaded after connecting contacts. |
| O2. Can user content be used to train the models that power ChatGPT? | Yes | Yes | Yes | OpenAI line 79 describes training with user content, subject to choices and exceptions. |
| O3. Does the policy affirm personalized ads for some users? | Yes | Yes | Yes | OpenAI line 65 describes personalization for Free/Go users. |
| O4. Does it affirm biometric identity recognition by OpenAI? | Not established | No | No | Line 35 names identity/age verification but no biometric identifier. |
| T1. Does optional TikTok contact sync collect the device phone book? | Yes | Yes | Yes | TikTok line 19 explicitly names the phone book. |
| T2. Does TikTok describe model training with platform information? | Yes | Yes | Yes | TikTok line 75 describes training machine-learning models. |
| T3. Does TikTok serve interest-based ads when permitted? | Yes | Yes | Yes | TikTok line 299 describes interest-based advertising with consent. |
| T4. Does face detection for effects establish biometric identity recognition? | Not established | No | No | TikTok line 37 describes face/body detection for effects, without identity recognition. |

p7 matched **15/16** localized question checks. The sole miss was O1 in Spanish. The Spanish policy is explicit: “si eliges conectar los contactos de tu dispositivo, subimos información de las libretas de direcciones de tu dispositivo” (line 33). The English run found this category at `clause-015`; the Spanish run produced no `contacts` finding even though the source has equivalent language. The machine report does not retain rejected candidate texts, so we cannot tell whether the model omitted the passage or p7 rejected a proposed citation.

## Interpretation and next work

This is a materially better result than round twelve on these selected pages, **not evidence that the model learned**: weights, prompts, and code were unchanged. Across rounds ten through thirteen, descriptive totals are **64 TP, 13 FP, and 6 FN** over 16 localized pages from eight underlying policies: **83.1% precision and 91.4% recall**. The pages are selected, translation pairs are correlated, the same reviewer assigned labels, and a policy-level category score does not prove that each finding's cited clause is the best or only support. No confidence interval or 90% general reliability claim is warranted.

Next, instrument separate reasons for rejected candidates and inspect the Spanish contact recall path. An independent reviewer should adjudicate labels, especially borderline AI-training and biometric cases. A later evaluation should use new publishers and include affirmative biometric identity recognition, plus policies where optional contacts, advertising, and model training are absent. The 267-second batch also leaves latency to improve for live use.

To reproduce the extraction while published content is unchanged, run `python scripts/prepare-real-policies-round13.py --fetch` from the repository root and compare the printed hashes to the manifest. With Ollama and `qwen3:4b-instruct` available, run from `apps/api`:

```powershell
go run ./cmd/evaluatepolicies -manifest evaluation/real_policies_round13.json -text-dir ../../.cache/real_policies_round13
```
