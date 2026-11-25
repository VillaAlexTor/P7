# Round 12: long, difficult privacy policies

## What we tested

On November 25, 2026, we evaluated p7 against the complete published policy pages of [Google](https://policies.google.com/privacy?hl=en) and [WhatsApp for the EEA](https://www.whatsapp.com/legal/privacy-policy-eea?lang=en_gb), each in English and Spanish. The Google policy was effective October 16, 2026; the WhatsApp policy was effective July 2, 2026. These are **two underlying policies and four localized pages**, not four independent policies. The extracted pages ranged from **57,636 to 120,506 characters** and from **431 to 976 p7 clauses**. Both fit the API's 200,000-character input limit. The extracts include some page navigation and repeated disclosure/table text, as a paste of a long web page may do.

We fixed the [source URLs, SHA-256 hashes, category labels, and eight audit questions](../../apps/api/evaluation/real_policies_round12.json) in commit `194da5e` **before** running the model. One reviewer applied the same seven category definitions used in rounds ten and eleven. The entire downloaded policy text remains in gitignored `.cache/real_policies_round12`; it is not republished. The [compact machine report](../../apps/api/evaluation/real_policies_round12_result.json) records all final categories, clause IDs, evidence-quote hashes, scores, timing, and candidate counts. The classifier was the unchanged local `qwen3:4b-instruct`, with temperature zero and the general profile. We ran each complete page once, without prompt or model changes.

## What p7 returned

| Page | Reference categories | p7 categories | True positives | False positives | Misses | Time |
| --- | --- | --- | ---: | ---: | ---: | ---: |
| Google EN | location, third parties, advertising, retention | those four + contacts + AI training | 4 | 2 | 0 | 76 s |
| Google ES | location, third parties, advertising, retention | those four + contacts + AI training | 4 | 2 | 0 | 119 s |
| WhatsApp EN | location, contacts, third parties, retention | those four + advertising + biometric | 4 | 2 | 0 | 175 s |
| WhatsApp ES | location, contacts, third parties, retention | exactly those four | 4 | 0 | 0 | 274 s |
| **Total** | | | **16** | **6** | **0** | **646 s** |

**Precision was 16/22 = 72.7%; recall was 16/16 = 100%; exact document-category sets were 1/4.** This run has no affirmative reference example for AI training or biometric identification, so it cannot measure recall for those categories. The model proposed 235 candidate findings; p7 retained 22 and counted 213 rejections. That rejection counter mixes duplicate-category proposals with invalid evidence, so it cannot tell us how many literal citations failed. The risk scores were 100 for Google EN, Google ES, and WhatsApp EN, and 76 for WhatsApp ES. These scores are not validated probabilities.

## Eight questions fixed before inference

Each question below is answered from the policy reference, then compared with whether p7 returned that category. These are **16 language-specific checks drawn from the same 28 document-category decisions**; they are not an independent accuracy test. “Not established” means that this policy does not affirm the defined practice, rather than proving the practice never occurs.

| Question | Reference | p7 EN | p7 ES | Why it is difficult |
| --- | --- | --- | --- | --- |
| G1. Does Google collect location information? | Yes | Yes | Yes | Precise and approximate location appear with optional controls. |
| G2. Does Google declare upload of the device address book? | Not established | **Yes** | **Yes** | Recipient suggestions and ordinary contact details do not affirm device-book access. |
| G3. Does the public-web AI passage establish training on private account content? | Not established | **Yes** | **Yes** | It mentions model training, but the named source is publicly available material. This label needs independent review. |
| G4. Does Google retain some data for open-ended business or legal needs? | Yes | Yes | Yes | Fixed deletion examples coexist with broad retention. |
| W1. Does WhatsApp allow device address-book upload? | Yes | Yes | Yes | The feature is optional, but upload is affirmative when chosen. |
| W2. Does this WhatsApp notice affirm personal-data-based ad targeting? | Not established | **Yes** | No | Cookies for ads and ads in Status/Channels do not by themselves establish targeting. |
| W3. Does using Meta AI through WhatsApp itself establish training on chats? | Not established | No | No | Data transfer for an AI feature is not the same claim as model training. |
| W4. Does WhatsApp keep some data while the account stays active? | Yes | Yes | Yes | Fixed deadlines for other data do not remove open-ended account retention. |

p7 matched **11 of these 16** question-based checks under the frozen mapping. The unasked ninth challenge surfaced after the run: WhatsApp EN labeled `biometric` from an age-estimation provider's video selfie. That clause describes estimating age and immediate deletion, not identifying a person by a biometric identifier. WhatsApp ES did not produce this finding. We did not add that question to the precommitted eight or alter the reference labels after seeing it.

## Error analysis and sensitivity

1. **Contact details versus address book:** Google EN clause `164` and Google ES clause `181` describe managing a name, email, and phone number. Both p7 findings were false positives under the category definition. This is the third complete-policy round where contact terminology caused an error, though WhatsApp's explicit address-book upload was correctly found in both languages.
2. **Public material versus private user data:** Google EN and ES clause `377` describes collecting publicly available information to help train AI models. The frozen label treated this as insufficient evidence that *private account content* is used for training. This is the most contestable judgment. If an independent reviewer counts public personal material as user data for this category, both Google AI findings become true positives: this round would instead have **18 true positives, 4 false positives, 0 misses, 81.8% precision**. We retain the original 72.7% result and show the alternate interpretation transparently.
3. **Ads versus ad targeting:** WhatsApp EN clause `074` describes technologies used for ads in optional features; it does not say personal information selects the ads. The Spanish run did not make this inference. The policy's denial of third-party banner ads also does not imply that no other ads exist.
4. **Age estimation versus identity recognition:** WhatsApp EN clause `117` describes a provider receiving video-selfie images to estimate age and deleting them afterward. The reference category requires a biometric identifier used to recognize a person. The Spanish run avoided this false positive.
5. **Long-document usability:** WhatsApp ES took **274 seconds** and the entire four-page batch took **10 minutes 46 seconds** on this local machine. This may be too slow for a live demo or routine reading. No browser-end-to-end timing was measured.

Across rounds ten through twelve, the unchanged pipeline has **41 true positives, 13 false positives, and 5 misses** across 12 localized pages representing six underlying policies: **75.9% precision and 89.1% recall** as descriptive totals. The languages are correlated, the policies were selected rather than randomly sampled, and one reviewer assigned all labels. This is not a statistically supported reliability estimate or evidence of training. The better recall here may simply reflect that these policies contain many explicit statements. A literal quote is necessary for traceability but does not make the interpretation correct.

## What to improve next

- Require explicit access, import, sync, or upload of a **device address book** for `contacts`; account contact fields and recipient suggestions must not satisfy the category.
- Require a statement that **personal data is used to select or personalize ads** for `advertising`; mention of ad cookies or ad locations alone is insufficient.
- Separate public-web model-training disclosures from private-user-data training, and have a second reviewer adjudicate the Google label before revising the category definition.
- Require identity recognition for `biometric`; age estimation with a selfie belongs to a different concept.
- Record why each candidate was rejected (duplicate, invalid clause ID, inexact quote, or unknown category), then profile long-document latency. Re-evaluate fixes on **new complete policies**, not these pages.

To reproduce the exact extraction while the published pages are unchanged, run from the repository root:

```powershell
python scripts/prepare-real-policies-round12.py --fetch
cd apps/api
go run ./cmd/evaluatepolicies -manifest evaluation/real_policies_round12.json -text-dir ../../.cache/real_policies_round12
```
