"""Extract the full official IKEA Spain privacy policy in English and Spanish.

Downloaded HTML and extracted text stay in gitignored .cache. Use --fetch to
refresh the source; otherwise reuse the frozen HTML from the evaluation run.
"""

import hashlib
import json
import runpy
import sys
import urllib.request
from pathlib import Path


ROOT = Path(__file__).resolve().parents[1]
CACHE = ROOT / ".cache" / "real_policies_round17"
EXTRACT = runpy.run_path(str(ROOT / "scripts" / "prepare-real-policies.py"))["extract"]
SOURCES = {
    "ikea_en": "https://www.ikea.com/es/en/customer-service/privacy-policy/",
    "ikea_es": "https://www.ikea.com/es/es/customer-service/privacy-policy/",
}
HEADINGS = {"ikea_en": "Privacy Policy", "ikea_es": "Política de privacidad"}


def main():
    CACHE.mkdir(parents=True, exist_ok=True)
    records = []
    for key, url in SOURCES.items():
        html_path = CACHE / f"{key}.html"
        if "--fetch" in sys.argv:
            request = urllib.request.Request(url, headers={"User-Agent": "p7-policy-evaluation/1.0"})
            html_path.write_bytes(urllib.request.urlopen(request, timeout=60).read())
        document = EXTRACT(html_path.read_text(encoding="utf-8"))
        heading = HEADINGS[key]
        start = document.find(heading + "\n")
        if start < 0:
            raise ValueError(f"Policy heading missing for {key}")
        document = document[start:]
        if not 10000 <= len(document) <= 200000 or "04/08/2026" not in document[-600:]:
            raise ValueError(f"Incomplete or oversized policy for {key}: {len(document)} characters")
        (CACHE / f"{key}.txt").write_bytes(document.encode("utf-8"))
        records.append({
            "id": key,
            "url": url,
            "characters": len(document),
            "sha256": hashlib.sha256(document.encode("utf-8")).hexdigest(),
        })
    print(json.dumps(records, indent=2, ensure_ascii=False))


if __name__ == "__main__":
    main()
