from __future__ import annotations

import re
import textwrap
from pathlib import Path

import numpy as np
import soundfile as sf


VOICE_DIR = Path("assets/voice")
IDS = [
    "01-problem",
    "02-input",
    "03-result",
    "04-evidence",
    "05-impact",
    "06-personalization",
    "07-ai",
    "08-pillars",
    "08-learning",
    "09-close",
]
LEAD = 0.6
GAP = 1.2


def caption_chunks(text: str) -> list[str]:
    clauses = [part.strip() for part in re.split(r"(?<=[.!?;:,])\s+", text) if part.strip()]
    chunks: list[str] = []
    for clause in clauses:
        words = clause.split()
        while len(words) > 11:
            chunks.append(" ".join(words[:9]))
            words = words[9:]
        if words:
            if chunks and len(words) < 3 and len(chunks[-1].split()) + len(words) <= 11:
                chunks[-1] += " " + " ".join(words)
            else:
                chunks.append(" ".join(words))
    merged: list[str] = []
    index = 0
    while index < len(chunks):
        current = chunks[index]
        if len(current.split()) < 4 and index + 1 < len(chunks):
            following = chunks[index + 1]
            if len(current.split()) + len(following.split()) <= 13:
                merged.append(current + " " + following)
                index += 2
                continue
        if len(current.split()) < 4 and merged and len(merged[-1].split()) + len(current.split()) <= 13:
            merged[-1] += " " + current
        else:
            merged.append(current)
        index += 1
    return merged


def silence_midpoints(audio: np.ndarray, sample_rate: int) -> list[float]:
    frame = max(1, int(sample_rate * 0.02))
    usable = len(audio) - len(audio) % frame
    blocks = audio[:usable].reshape(-1, frame)
    rms = np.sqrt(np.mean(np.square(blocks), axis=1) + 1e-12)
    silent = rms < 0.012
    minimum = 8
    mids: list[float] = []
    start = None
    for index, value in enumerate(silent):
        if value and start is None:
            start = index
        if (not value or index == len(silent) - 1) and start is not None:
            end = index if not value else index + 1
            if end - start >= minimum:
                mids.append(((start + end) / 2) * frame / sample_rate)
            start = None
    duration = len(audio) / sample_rate
    return [mid for mid in mids if 0.2 < mid < duration - 0.2]


def choose_boundaries(chunks: list[str], candidates: list[float], duration: float) -> list[float]:
    if len(chunks) == 1:
        return []
    weights = [len(chunk.split()) for chunk in chunks]
    total = sum(weights)
    cumulative = np.cumsum(weights)[:-1] / total
    targets = [float(value * duration) for value in cumulative]
    chosen: list[float] = []
    last = 0.0
    for index, target in enumerate(targets):
        remaining = len(targets) - index
        minimum = last + 0.45
        maximum = duration - 0.45 * remaining
        available = [candidate for candidate in candidates if minimum <= candidate <= maximum]
        boundary = min(available, key=lambda candidate: abs(candidate - target)) if available else target
        boundary = min(max(boundary, minimum), maximum)
        chosen.append(boundary)
        last = boundary
    return chosen


def stamp(seconds: float) -> str:
    milliseconds = int(round(seconds * 1000))
    hours, milliseconds = divmod(milliseconds, 3_600_000)
    minutes, milliseconds = divmod(milliseconds, 60_000)
    secs, milliseconds = divmod(milliseconds, 1000)
    return f"{hours:02d}:{minutes:02d}:{secs:02d},{milliseconds:03d}"


def main() -> None:
    cues: list[tuple[float, float, str]] = []
    offset = LEAD
    for item_id in IDS:
        text = (VOICE_DIR / f"{item_id}.txt").read_text(encoding="utf-8").strip()
        audio, rate = sf.read(VOICE_DIR / f"{item_id}.wav", dtype="float32")
        if audio.ndim > 1:
            audio = audio.mean(axis=1)
        duration = len(audio) / rate
        chunks = caption_chunks(text)
        boundaries = choose_boundaries(chunks, silence_midpoints(audio, rate), duration)
        starts = [0.0] + boundaries
        ends = boundaries + [duration]
        for start, end, chunk in zip(starts, ends, chunks):
            wrapped = "\n".join(textwrap.wrap(chunk, width=42, break_long_words=False))
            cues.append((offset + start, offset + end, wrapped))
        offset += duration + GAP

    lines: list[str] = []
    for index, (start, end, text) in enumerate(cues, start=1):
        lines.extend([str(index), f"{stamp(start)} --> {stamp(end)}", text, ""])
    Path("assets/p7-demo.en.srt").write_text("\n".join(lines), encoding="utf-8")
    print(f"wrote {len(cues)} cues through {stamp(cues[-1][1])}")


if __name__ == "__main__":
    main()
