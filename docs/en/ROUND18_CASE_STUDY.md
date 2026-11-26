# Round 18: final full-policy evaluation, Dropbox and Airbnb

## Frozen sources and reference

On November 26, 2026, p7 analyzed the complete current [Dropbox English](https://www.dropbox.com/en/privacy) and [Spanish](https://www.dropbox.com/es/privacy) privacy pages and Airbnb's official 2026 privacy PDFs in [English](https://assets.airbnb.com/help/2026_Privacy_Policy-English.pdf) and [Spanish](https://assets.airbnb.com/help/2026_Privacy_Policy-Spanish_Spain.pdf). These are **two publishers, each in two languages**, not four independent policies. Dropbox's current policy is effective May 30, 2025; its banner announces a different policy effective January 1, 2027. The extraction starts at the **current** policy heading. Airbnb's PDF says it was updated February 5, 2026. The four extracted texts contain 17,456–61,064 characters. Source HTML, PDFs, and text remain in gitignored `.cache/real_policies_round18`.

The [manifest](../../apps/api/evaluation/real_policies_round18.json), committed as `0ebbc80` **before inference**, fixes URLs, SHA-256 hashes, lengths, document-category labels, and eight bilingual questions. One reviewer assigned the reference. **Dropbox had already been tested in round 15** using the same current policy URLs, so this round does not add two independent publishers. Dropbox is positive for `location`, `third-parties`, and `retention`; generic upgrade emails do not establish personalized advertising. Round 18 labels its “your contacts” wording negative because it does not explicitly identify a **device address book**. Round 15 labeled the same practice positive. This inconsistency needs independent adjudication. Airbnb is positive for all seven categories: it describes address-book import, face recognition from identity documents, use of personal data to improve AI models, tailored advertising, location, disclosure, and purpose-based retention.

## Raw results and repair

Both passes used local Ollama `qwen3:4b-instruct`, temperature zero, the same source files and frozen labels, general profile, and full-document processing. The [initial raw result](../../apps/api/evaluation/real_policies_round18_result_initial.json) exposed a PDF citation failure: the model proposed 8 English and 20 Spanish Airbnb candidates, but the literal verifier rejected **all 28**, chiefly because PDF extraction inserted line breaks inside sentences. We changed only the evidence gate to compare whitespace-normalized text **within the same cited clause**. A test confirms that altered substantive words and wrong clause IDs remain rejected. The [revised raw result](../../apps/api/evaluation/real_policies_round18_result.json) is an *in-sample repair check*, **not an independent new holdout**. No model weights were trained.

| Pass | Dropbox EN | Dropbox ES | Airbnb EN | Airbnb ES | TP | FP | FN | Precision | Recall |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: |
| Initial | 3 expected + contacts | 3 expected + contacts | none | none | 6 | 2 | 14 | 75.0% | 30.0% |
| After gate repair | 3 expected + contacts | 3 expected + contacts | 5 of 7; misses location, contacts | 3 of 7; misses contacts, AI training, biometric, retention | 14 | 2 | 6 | 87.5% | 70.0% |

The revised run has **0/4 exact page-level category sets** and 6 true negatives across 28 document-category decisions. The eight frozen questions produce 16 localized checks, of which **10/16** match the category output. The translations are correlated, and the Dropbox contact reference conflicts with round 15. If both Dropbox contact labels were instead positive, the same predictions would become **16 TP, 0 FP, 6 FN**, or **100% precision and 72.7% recall**. That is a *post-hoc sensitivity calculation*, not a replacement for the frozen result or independent validation. These figures do **not** establish reliability in deployment or calibrated confidence.

## Citation audit and remaining errors

We inspected all 16 retained cited clauses in the revised output. **13/16** support their assigned category under the frozen definitions. The two Dropbox `contacts` citations (`clause-010` in each language) say the user can give access to contacts and Dropbox stores them, but do not identify a device address book. This is the source of both reference false positives and a legitimate annotation ambiguity. Airbnb EN `third-parties` cites `clause-114`, which describes payment data **received from** an outside provider; outward disclosure is stated elsewhere in Airbnb's policy. This is a category-level true positive with an unsupportive citation. All other retained clause locations appeared substantively relevant in this one-reviewer audit. The automatic `evidenceVerified` flag still means text presence after whitespace normalization, **not semantic proof**.

The repair raised Airbnb coverage substantially, yet 3 of 8 English candidates and 17 of 20 Spanish candidates were still rejected. The Spanish Airbnb PDF missed four explicit categories, including AI-model improvement and biometric facial recognition. Longer PDF input, PDF extraction, model selection, and citation generation remain separate failure points. The next technical step would be a dedicated citation-selection and semantic-validation stage, followed by a fresh holdout and a second human annotator. We make **no 90% reliability claim**.

## Reproduction

The extracted source files have frozen hashes in the manifest. `--fetch` obtains a **new version** if an official page changes. With the original cached files, install `pypdf` if needed (`python -m pip install --target .cache/pypdf pypdf`), then run `python scripts/prepare-real-policies-round18.py` and compare hashes. Start local Ollama with `qwen3:4b-instruct`; from `apps/api`, run:

```powershell
go run ./cmd/evaluatepolicies -manifest evaluation/real_policies_round18.json -text-dir ../../.cache/real_policies_round18
```

The command uses current code and model behavior, so the two checked-in JSON files are the exact raw records for these two passes.
