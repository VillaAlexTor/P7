"""Fetch and extract complete official privacy notices for local evaluation.

The downloaded third-party text stays in .cache (gitignored). The published
evaluation records the source URL, extraction hash, and model output.
"""

import hashlib
import json
import sys
import urllib.request
from html.parser import HTMLParser
from pathlib import Path


ROOT = Path(__file__).resolve().parents[1]
CACHE = ROOT / ".cache" / "real_policies"
SOURCES = {
    "mozilla_en": "https://www.mozilla.org/en-US/privacy/mozilla-accounts/",
    "mozilla_es": "https://www.mozilla.org/es-ES/privacy/mozilla-accounts/",
    "proton_en": "https://proton.me/legal/privacy",
    "proton_es": "https://proton.me/es-419/legal/privacy",
}


class MainText(HTMLParser):
    BLOCKS = {"p", "h1", "h2", "h3", "h4", "h5", "li", "div", "section", "br", "ul", "ol"}
    SKIP = {"script", "style", "svg", "nav", "footer", "aside", "noscript"}

    def __init__(self):
        super().__init__(convert_charrefs=True)
        self.main_depth = 0
        self.skip_depth = 0
        self.parts = []

    def handle_starttag(self, tag, attrs):
        if tag == "main":
            self.main_depth += 1
        elif self.main_depth and tag in self.SKIP:
            self.skip_depth += 1
        if self.main_depth and not self.skip_depth and tag in self.BLOCKS:
            self.parts.append("\n")

    def handle_endtag(self, tag):
        if self.main_depth and not self.skip_depth and tag in self.BLOCKS:
            self.parts.append("\n")
        if tag == "main":
            self.main_depth -= 1
        elif self.main_depth and tag in self.SKIP:
            self.skip_depth -= 1

    def handle_data(self, data):
        if self.main_depth and not self.skip_depth:
            self.parts.append(data)


def extract(html):
    parser = MainText()
    parser.feed(html)
    lines = [" ".join(line.split()) for line in "".join(parser.parts).splitlines()]
    lines = [line for line in lines if line]
    return "\n\n".join(lines) + "\n"


def main():
    CACHE.mkdir(parents=True, exist_ok=True)
    download = "--fetch" in sys.argv
    records = []
    for key, url in SOURCES.items():
        html_path = CACHE / f"{key}.html"
        if download:
            request = urllib.request.Request(url, headers={"User-Agent": "p7-policy-evaluation/1.0"})
            html_path.write_bytes(urllib.request.urlopen(request, timeout=40).read())
        text = extract(html_path.read_text(encoding="utf-8"))
        if len(text) < 3000:
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
