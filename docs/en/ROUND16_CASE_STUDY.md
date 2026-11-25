# Round 16: current Pinterest and Slack privacy policies

## Scope and frozen method

On November 25, 2026, p7 analyzed the complete official [Pinterest EN](https://policy.pinterest.com/en/privacy-policy), [Pinterest ES](https://policy.pinterest.com/es/privacy-policy), [Slack EN](https://slack.com/trust/privacy/privacy-policy), and [Slack ES](https://slack.com/intl/es-es/trust/privacy/privacy-policy) pages. These are **two policies in two languages**, not four independent policies. Slack's notice displays a July 5, 2023 effective date. The extracted texts were 39,142–58,404 characters. Source texts remain in gitignored `.cache/real_policies_round16`; the [manifest](../../apps/api/evaluation/real_policies_round16.json) fixes their URLs, SHA-256 hashes, lengths, reference lines, category labels, and eight questions. It was committed as `f563e0c` **before inference**. One reviewer assigned labels. The unchanged local `qwen3:4b-instruct`, temperature zero, general profile, prompts, and Go pipeline analyzed each page once. The [machine report](../../apps/api/evaluation/real_policies_round16_result.json) records predictions, clause IDs, quote hashes, scores, counts, and timing without republishing source policies.

Pinterest says it **continues to process contacts previously synced** (source line 37). Under p7's existing definition, merely retaining an earlier contact list does not establish new address-book access, so the frozen `contacts` label was negative. Pinterest explicitly says it trains AI models with Pins (line 83). Slack expressly describes optional address-book import (line 73), targeted ads on non-Slack sites (line 291 EN / 295 ES), and retaining Other Information after deactivation (line 225 EN / 229 ES). This Slack notice does not affirm AI-model training.

## Results under the frozen reference

| Full page | Reference categories | p7 categories | TP | FP | FN | Time |
| --- | --- | --- | ---: | ---: | ---: | ---: |
| Pinterest EN | location, third parties, advertising, AI training, retention | first four + **contacts**; **retention missing** | 4 | 1 | 1 | 87 s |
| Pinterest ES | same five | those five + **contacts, biometric** | 5 | 2 | 0 | 122 s |
| Slack EN | location, contacts, third parties, advertising, retention | exactly those five | 5 | 0 | 0 | 67 s |
| Slack ES | same five | location, contacts, third parties + **biometric**; **advertising and retention missing** | 3 | 1 | 2 | 85 s |
| **Total** | | | **17** | **4** | **3** | **363 s** |

Frozen-label document-category precision was **17/21 = 81.0%**, recall **17/20 = 85.0%**, and **1/4** localized pages had exact category sets. Across 28 decisions, that is 17 TP, 4 FP, 3 FN, and 4 true negatives. The model proposed 138 candidates, retained 21 literal-quote matches, and counted 117 rejections; this counter mixes duplicates and invalid evidence. The scores were 100, 100, 92, and 96 in page order; they are not calibrated probabilities.

### Important post-inference label correction

The frozen biometric negatives for Pinterest were **my annotation mistake**. I considered its body-type model, which does not identify people, but missed a separate statement at lines **95 and 199 in both languages**: Pinterest may compare a person's face with a government ID photo for age verification. Under the identity-matching interpretation used for Spotify in round eleven, both Pinterest pages should be `biometric` positives. The Spanish model found this practice at `clause-122`; the English model missed it. We **did not alter** the frozen manifest, raw output, or original result. The separate [post-inference sensitivity record](../../apps/api/evaluation/real_policies_round16_adjudication.json) changes only these two reference labels, yielding **18 TP, 3 FP, 4 FN; 85.7% precision and 81.8% recall**. Exact pages remain 1/4. This revised interpretation was made after seeing predictions and needs independent adjudication; it is **not a replacement preregistered result**.

## Eight questions fixed before inference

These give 16 localized checks that overlap the 28 category decisions; they are not a separate sample. “Not established” means the current notice does not affirm the defined practice.

| Question | Frozen reference | p7 EN | p7 ES | Reason |
| --- | --- | --- | --- | --- |
| P1. Does Pinterest affirm a **new** device address-book import? | Not established | **Yes** | **Yes** | Line 37 says continued use of contacts synced previously. |
| P2. Does Pinterest train AI models with Pins? | Yes | Yes | Yes | Line 83 explicitly describes model training. |
| P3. Does Pinterest select ads using inferred interests? | Yes | Yes | Yes | Line 339 describes selecting ads from interests. |
| P4. Does the body-type technology establish biometric identification? | Not established | No | **Yes** | The frozen question named body-type technology, but the document also has face-to-ID matching at lines 95/199. Under the corrected *document* label, EN misses and ES finds biometrics; the question wording itself remains narrower. |
| S1. Does Slack collect an address book after user choice? | Yes | Yes | Yes | Line 73 explicitly names importing an address book. |
| S2. Does this Slack notice affirm training on workspace content? | Not established | No | No | No affirmative training disclosure in this notice. |
| S3. Does Slack disclose data sharing for targeted ads outside Slack? | Yes | Yes | **No** | Line 291 EN / 295 ES describes that practice. |
| S4. May Slack retain Other Information after deactivation? | Yes | Yes | **No** | Line 225 EN / 229 ES says it may. |

Under the original question mapping, p7 matched **11/16** localized checks. P4 illustrates why a category-presence answer is not the same as answering its precise wording: Spanish p7's biometric finding came from a different passage. The frozen P4 answer remains correct **for body-type technology alone**, while the full Pinterest document's biometric category requires the correction above. We did not rewrite the question after inference.

## Audit of cited clauses

We inspected all 21 retained clause locations. **At least eight do not by themselves support their assigned finding**:

| Page | Finding and clause | Problem |
| --- | --- | --- |
| Pinterest EN/ES | `third-parties`, `clause-029` / `clause-038` | Describes data **received from** affiliates/advertisers, not outward sharing. The policy discusses outward transfers elsewhere (source line 541). |
| Pinterest EN/ES | `contacts`, `clause-019` / `clause-025` | Describes continued processing of **previously** synced contacts, not current device-book import. Both are category false positives under the frozen and corrected references. |
| Slack EN/ES | `third-parties`, `clause-043` in both | Says third-party services **send** information to Slack. Sharing from Slack to enabled services appears elsewhere (EN line 201 / ES line 205). |
| Slack EN | `advertising`, `clause-102` | Promotional emails are not personal-data-based targeting of ads. The document category is present at line 291. |
| Slack ES | `biometric`, `clause-076` | Mentions a **device** fingerprint reader, not provider-run biometric processing. |

Under the frozen mapping, **five weak citations accompany category-level true positives** and three accompany false positives: Pinterest contacts twice and Slack ES biometrics. The remaining fourth false positive is Pinterest ES biometrics, whose face-to-ID citation is actually relevant under the corrected reference. Thus citation quality and category accuracy must be reported separately. Other locations appeared supportive in this one-reviewer audit. p7 verifies literal substring containment, not semantic sufficiency.

## Interpretation and next work

The strongest recurring errors are confusing **incoming** data with outgoing sharing, treating retained old contacts as current access, and mistaking device-only biometrics for provider processing. Spanish Slack also missed two explicit practices. The post-inference correction shows that reference-label quality is itself a limit of this evaluation. Selected publishers, correlated translations, a single annotator, and a disputed label prevent a general reliability claim. No model or prompt changed; score differences are not evidence of training.

Next, obtain a second reviewer for the Pinterest biometric label, split citation rejection reasons, and require the cited clause to state the category's defining action. Test any fix on **new** complete policies, preserving this frozen run. To reproduce while pages remain unchanged, run `python scripts/prepare-real-policies-round16.py --fetch` from the repository root, compare hashes, then with Ollama and `qwen3:4b-instruct` from `apps/api`:

```powershell
go run ./cmd/evaluatepolicies -manifest evaluation/real_policies_round16.json -text-dir ../../.cache/real_policies_round16
```
