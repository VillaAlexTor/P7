"""Fetch and extract current, long Google and WhatsApp privacy policies.

Third-party text remains in gitignored .cache. The frozen manifest stores its
hash and labels, while the committed report stores only compact predictions.
"""

import hashlib
import json
import sys
import urllib.request
from html.parser import HTMLParser
from pathlib import Path


ROOT = Path(__file__).resolve().parents[1]
CACHE = ROOT / ".cache" / "real_policies_round12"
SOURCES = {
    "google_en": "https://policies.google.com/privacy?hl=en",
    "google_es": "https://policies.google.com/privacy?hl=es",
    "whatsapp_en": "https://www.whatsapp.com/legal/privacy-policy-eea?lang=en_gb",
    "whatsapp_es": "https://www.whatsapp.com/legal/privacy-policy-eea?lang=es_LA",
}


class DocumentText(HTMLParser):
    BLOCKS = {"p", "h1", "h2", "h3", "h4", "h5", "li", "div", "section", "br", "ul", "ol", "td", "th"}
    SKIP = {"script", "style", "svg", "nav", "footer", "aside", "noscript"}

    def __init__(self, source):
        super().__init__(convert_charrefs=True)
        self.source = source
        self.root_depth = 0
        self.skip_depth = 0
        self.parts = []

    def is_root(self, tag, attrs):
        attributes = dict(attrs)
        if self.source.startswith("google_"):
            return tag == "div" and attributes.get("role") == "main"
        return tag == "div" and "_8l_f" in attributes.get("class", "").split()

    def handle_starttag(self, tag, attrs):
        if not self.root_depth and self.is_root(tag, attrs):
            self.root_depth = 1
        elif self.root_depth and tag == "div":
            self.root_depth += 1
        if self.root_depth and tag in self.SKIP:
            self.skip_depth += 1
        if self.root_depth and not self.skip_depth and tag in self.BLOCKS:
            self.parts.append("\n")

    def handle_endtag(self, tag):
        if self.root_depth and not self.skip_depth and tag in self.BLOCKS:
            self.parts.append("\n")
        if self.root_depth and tag in self.SKIP:
            self.skip_depth -= 1
        if self.root_depth and tag == "div":
            self.root_depth -= 1

    def handle_data(self, data):
        if self.root_depth and not self.skip_depth:
            self.parts.append(data)


def extract(source, html):
    parser = DocumentText(source)
    parser.feed(html)
    lines = [" ".join(line.split()) for line in "".join(parser.parts).splitlines()]
    return "\n\n".join(line for line in lines if line) + "\n"


def main():
    CACHE.mkdir(parents=True, exist_ok=True)
    records = []
    for key, url in SOURCES.items():
        html_path = CACHE / f"{key}.html"
        if "--fetch" in sys.argv:
            request = urllib.request.Request(url, headers={"User-Agent": "p7-policy-evaluation/1.0"})
            html_path.write_bytes(urllib.request.urlopen(request, timeout=50).read())
        document = extract(key, html_path.read_text(encoding="utf-8"))
        if not 12000 <= len(document) <= 200000:
            raise ValueError(f"Extraction outside p7 input limits for {key}: {len(document)} characters")
        (CACHE / f"{key}.txt").write_bytes(document.encode("utf-8"))
        records.append({
            "id": key, "url": url, "characters": len(document),
            "sha256": hashlib.sha256(document.encode("utf-8")).hexdigest(),
        })
    print(json.dumps(records, indent=2, ensure_ascii=False))


if __name__ == "__main__":
    main()
