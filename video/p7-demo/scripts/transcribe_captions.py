from __future__ import annotations

import json
import textwrap
from pathlib import Path

from faster_whisper import WhisperModel


def stamp(seconds: float) -> str:
    milliseconds = int(round(seconds * 1000))
    hours, milliseconds = divmod(milliseconds, 3_600_000)
    minutes, milliseconds = divmod(milliseconds, 60_000)
    secs, milliseconds = divmod(milliseconds, 1000)
    return f"{hours:02d}:{minutes:02d}:{secs:02d},{milliseconds:03d}"


def main() -> None:
    cache = Path("../../.cache/faster-whisper").resolve()
    model = WhisperModel("tiny.en", device="cpu", compute_type="int8", download_root=str(cache))
    segments, _ = model.transcribe(
        "assets/voice/narration.wav",
        language="en",
        word_timestamps=True,
        vad_filter=True,
        beam_size=5,
    )
    words = []
    for segment in segments:
        for word in segment.words or []:
            text = word.word.strip()
            if text:
                words.append({"text": text, "start": word.start, "end": word.end})

    groups = []
    current = []
    for word in words:
        if current and (word["start"] - current[-1]["end"] >= 0.18 or len(current) >= 6):
            groups.append(current)
            current = []
        current.append(word)
        if word["text"].endswith((".", "?", "!")) and len(current) >= 3:
            groups.append(current)
            current = []
    if current:
        groups.append(current)

    srt_lines = []
    transcript = []
    for index, group in enumerate(groups, start=1):
        text = " ".join(word["text"] for word in group)
        wrapped = "\n".join(textwrap.wrap(text, width=42, break_long_words=False))
        start = max(0.0, group[0]["start"] - 0.05)
        end = group[-1]["end"] + 0.08
        srt_lines.extend([str(index), f"{stamp(start)} --> {stamp(end)}", wrapped, ""])
        transcript.append({"text": text, "start": start, "end": end})

    Path("assets/p7-demo.en.srt").write_text("\n".join(srt_lines), encoding="utf-8")
    Path("assets/transcript.json").write_text(json.dumps(transcript, indent=2), encoding="utf-8")
    print(json.dumps({"words": len(words), "cues": len(groups), "end": groups[-1][-1]["end"]}))


if __name__ == "__main__":
    main()
