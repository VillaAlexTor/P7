"""Fetch and extract full OpenAI and TikTok policies for the local evaluation."""

import hashlib
import json
import runpy
import sys
import urllib.request
from html.parser import HTMLParser
from pathlib import Path


ROOT = Path(__file__).resolve().parents[1]
CACHE = ROOT / ".cache" / "real_policies_round13"
MAIN_EXTRACT = runpy.run_path(str(ROOT / "scripts" / "prepare-real-policies.py"))["extract"]
SOURCES = {
    "openai_en": "https://openai.com/policies/privacy-policy/",
    "openai_es": "https://openai.com/es-419/policies/privacy-policy/",
    "tiktok_en": "https://www.tiktok.com/legal/page/eea/privacy-policy/en",
    "tiktok_es": "https://www.tiktok.com/legal/page/eea/privacy-policy/es",
}


class TikTokBody(HTMLParser):
    BLOCKS = {"p", "h1", "h2", "h3", "h4", "h5", "li", "div", "section", "br", "ul", "ol", "td", "th"}
    SKIP = {"script", "style", "svg", "nav", "footer", "aside", "noscript"}

    def __init__(self):
        super().__init__(convert_charrefs=True)
        self.depth = 0
        self.skip = 0
        self.parts = []

    def handle_starttag(self, tag, attrs):
        if not self.depth and tag == "section" and dict(attrs).get("data-testid") in {"policy-card-pre-content", "policy-card-post-content"}:
            self.depth = 1
        elif self.depth and tag == "section":
            self.depth += 1
        if self.depth and tag in self.SKIP:
            self.skip += 1
        if self.depth and not self.skip and tag in self.BLOCKS:
            self.parts.append("\n")

    def handle_endtag(self, tag):
        if self.depth and not self.skip and tag in self.BLOCKS:
            self.parts.append("\n")
        if self.depth and tag in self.SKIP:
            self.skip -= 1
        if self.depth and tag == "section":
            self.depth -= 1

    def handle_data(self, data):
        if self.depth and not self.skip:
            self.parts.append(data)


def extract(key, html):
    if key.startswith("openai_"):
        return MAIN_EXTRACT(html)
    parser = TikTokBody()
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
        if not 10000 <= len(document) <= 200000:
            raise ValueError(f"Extraction outside p7 input limits for {key}: {len(document)} characters")
        (CACHE / f"{key}.txt").write_bytes(document.encode("utf-8"))
        records.append({"id": key, "url": url, "characters": len(document), "sha256": hashlib.sha256(document.encode("utf-8")).hexdigest()})
    print(json.dumps(records, indent=2, ensure_ascii=False))


if __name__ == "__main__":
    main()
