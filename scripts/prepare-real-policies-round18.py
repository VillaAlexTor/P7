"""Extract current Dropbox HTML and Airbnb's 2026 PDF policies, EN and ES.

Downloaded sources and extracted text stay in gitignored .cache. Install pypdf
into .cache/pypdf if it is not otherwise available. --fetch refreshes sources;
without it, the script reuses the downloaded files from the frozen evaluation.
"""

import hashlib
import json
import runpy
import sys
import unicodedata
import urllib.request
from pathlib import Path


ROOT = Path(__file__).resolve().parents[1]
CACHE = ROOT / ".cache" / "real_policies_round18"
sys.path.insert(0, str(ROOT / ".cache" / "pypdf"))
from pypdf import PdfReader  # noqa: E402


EXTRACT_HTML = runpy.run_path(str(ROOT / "scripts" / "prepare-real-policies.py"))["extract"]
SOURCES = {
    "dropbox_en": ("https://www.dropbox.com/en/privacy", "html", "Dropbox Privacy Policy"),
    "dropbox_es": ("https://www.dropbox.com/es/privacy", "html", "Política de privacidad de Dropbox"),
    "airbnb_en": ("https://assets.airbnb.com/help/2026_Privacy_Policy-English.pdf", "pdf", "Privacy Policy"),
    "airbnb_es": ("https://assets.airbnb.com/help/2026_Privacy_Policy-Spanish_Spain.pdf", "pdf", "Política de Privacidad"),
}


def extract_pdf(path):
    reader = PdfReader(path)
    pages = [unicodedata.normalize("NFKC", page.extract_text()).replace("\xa0", " ") for page in reader.pages]
    document = "\n\n".join(pages)
    return document.rstrip() + "\n"


def main():
    CACHE.mkdir(parents=True, exist_ok=True)
    records = []
    for key, (url, kind, heading) in SOURCES.items():
        source = CACHE / f"{key}.{kind}"
        if "--fetch" in sys.argv:
            request = urllib.request.Request(url, headers={"User-Agent": "Mozilla/5.0 p7-policy-evaluation/1.0"})
            source.write_bytes(urllib.request.urlopen(request, timeout=60).read())
        if kind == "html":
            document = EXTRACT_HTML(source.read_text(encoding="utf-8"))
            start = document.find(heading + "\n")
            if start < 0:
                raise ValueError(f"Current policy heading missing for {key}")
            document = document[start:]
        else:
            document = extract_pdf(source)
        if not document.startswith(heading) or not 10000 <= len(document) <= 200000:
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
