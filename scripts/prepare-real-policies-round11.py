"""Prepare the Apple and Spotify full-document holdout in gitignored cache."""

import hashlib
import json
import runpy
import sys
import urllib.request
from pathlib import Path


ROOT = Path(__file__).resolve().parents[1]
CACHE = ROOT / ".cache" / "real_policies_round11"
EXTRACT = runpy.run_path(str(ROOT / "scripts" / "prepare-real-policies.py"))["extract"]
SOURCES = {
    "apple_en": "https://www.apple.com/legal/privacy/en-ww/",
    "apple_es": "https://www.apple.com/legal/privacy/es/",
    "spotify_en": "https://www.spotify.com/uk/legal/privacy-policy/",
    "spotify_es": "https://www.spotify.com/es/legal/privacy-policy/",
}


def main():
    CACHE.mkdir(parents=True, exist_ok=True)
    records = []
    for key, url in SOURCES.items():
        html_path = CACHE / f"{key}.html"
        if "--fetch" in sys.argv:
            request = urllib.request.Request(url, headers={"User-Agent": "p7-policy-evaluation/1.0"})
            html_path.write_bytes(urllib.request.urlopen(request, timeout=45).read())
        text = EXTRACT(html_path.read_text(encoding="utf-8"))
        if len(text) < 5000:
            raise ValueError(f"Extraction too short for {key}: {len(text)} characters")
        (CACHE / f"{key}.txt").write_bytes(text.encode("utf-8"))
        records.append({
            "id": key,
            "url": url,
            "characters": len(text),
            "sha256": hashlib.sha256(text.encode("utf-8")).hexdigest(),
        })
    print(json.dumps(records, indent=2, ensure_ascii=False))


if __name__ == "__main__":
    main()
