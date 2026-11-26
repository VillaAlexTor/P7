import difflib
import json
import re
from pathlib import Path

ids = ["01-problem", "02-input", "03-result", "04-evidence", "05-impact", "06-personalization", "07-ai", "08-pillars", "08-learning", "09-close"]
recognized = " ".join(item["text"] for item in json.loads(Path("assets/transcript.json").read_text(encoding="utf-8")))
source = " ".join((Path("assets/voice") / f"{item_id}.txt").read_text(encoding="utf-8") for item_id in ids)
normalize = lambda value: re.findall(r"[a-z0-9']+", value.lower())
left, right = normalize(recognized), normalize(source)
matcher = difflib.SequenceMatcher(None, left, right)
print({"recognized": len(left), "source": len(right), "ratio": round(matcher.ratio(), 4)})
for tag, i1, i2, j1, j2 in matcher.get_opcodes():
    if tag != "equal":
        print(f"{tag}: REC=[{' '.join(left[i1:i2])}] SRC=[{' '.join(right[j1:j2])}]")
