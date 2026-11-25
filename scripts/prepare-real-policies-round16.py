"""Extract complete official Pinterest and Slack privacy policies."""

import hashlib
import json
import runpy
import sys
import urllib.request
from pathlib import Path


ROOT = Path(__file__).resolve().parents[1]
CACHE = ROOT / ".cache" / "real_policies_round16"
EXTRACT = runpy.run_path(str(ROOT / "scripts" / "prepare-real-policies.py"))["extract"]
SOURCES = {
    "pinterest_en": "https://policy.pinterest.com/en/privacy-policy",
    "pinterest_es": "https://policy.pinterest.com/es/privacy-policy",
    "slack_en": "https://slack.com/trust/privacy/privacy-policy",
    "slack_es": "https://slack.com/intl/es-es/trust/privacy/privacy-policy",
}


def main():
    CACHE.mkdir(parents=True, exist_ok=True)
    records = []
    for key, url in SOURCES.items():
        html_path = CACHE / f"{key}.html"
        if "--fetch" in sys.argv:
            request = urllib.request.Request(url, headers={"User-Agent": "p7-policy-evaluation/1.0"})
            html_path.write_bytes(urllib.request.urlopen(request, timeout=60).read())
        document = EXTRACT(html_path.read_text(encoding="utf-8"))
        if not 10000 <= len(document) <= 200000:
            raise ValueError(f"Extraction outside p7 input limits for {key}: {len(document)} characters")
        (CACHE / f"{key}.txt").write_bytes(document.encode("utf-8"))
        records.append({"id": key, "url": url, "characters": len(document), "sha256": hashlib.sha256(document.encode("utf-8")).hexdigest()})
    print(json.dumps(records, indent=2, ensure_ascii=False))


if __name__ == "__main__":
    main()
