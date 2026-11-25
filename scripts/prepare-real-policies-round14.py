"""Extract the complete official LinkedIn and Zoom privacy statements."""

import hashlib
import json
import re
import runpy
import sys
import urllib.request
from pathlib import Path


ROOT = Path(__file__).resolve().parents[1]
CACHE = ROOT / ".cache" / "real_policies_round14"
EXTRACT = runpy.run_path(str(ROOT / "scripts" / "prepare-real-policies.py"))["extract"]
SOURCES = {
    "linkedin_en": "https://www.linkedin.com/legal/privacy-policy",
    "linkedin_es": "https://es.linkedin.com/legal/privacy-policy",
    "zoom_en": "https://www.zoom.com/en/trust/privacy/privacy-statement/",
    "zoom_es": "https://www.zoom.com/es/trust/privacy/privacy-statement/",
}


def main():
    CACHE.mkdir(parents=True, exist_ok=True)
    records = []
    for key, url in SOURCES.items():
        html_path = CACHE / f"{key}.html"
        if "--fetch" in sys.argv:
            request = urllib.request.Request(url, headers={"User-Agent": "p7-policy-evaluation/1.0"})
            html_path.write_bytes(urllib.request.urlopen(request, timeout=60).read())
        html = html_path.read_text(encoding="utf-8")
        if key.startswith("zoom_"):
            body = re.search(r'<main\b[^>]*\bid="main"[^>]*>.*?</main>', html, re.DOTALL)
            if body is None:
                raise ValueError(f"Zoom policy body not found for {key}")
            html = body.group(0)
        document = EXTRACT(html)
        if not 10000 <= len(document) <= 200000:
            raise ValueError(f"Extraction outside p7 input limits for {key}: {len(document)} characters")
        (CACHE / f"{key}.txt").write_bytes(document.encode("utf-8"))
        records.append({"id": key, "url": url, "characters": len(document), "sha256": hashlib.sha256(document.encode("utf-8")).hexdigest()})
    print(json.dumps(records, indent=2, ensure_ascii=False))


if __name__ == "__main__":
    main()
