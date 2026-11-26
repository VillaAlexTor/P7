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
TAIL = 2.5


def silence(seconds: float, sample_rate: int) -> np.ndarray:
    return np.zeros(round(seconds * sample_rate), dtype=np.float32)


def main() -> None:
    clips: list[np.ndarray] = []
    sample_rate = 0
    for index, item_id in enumerate(IDS):
        audio, rate = sf.read(VOICE_DIR / f"{item_id}.wav", dtype="float32")
        if audio.ndim > 1:
            audio = audio.mean(axis=1)
        if sample_rate and rate != sample_rate:
            raise RuntimeError("Voice chapter sample rates differ")
        sample_rate = rate
        if index == 0:
            clips.append(silence(LEAD, sample_rate))
        clips.append(audio)
        if index < len(IDS) - 1:
            clips.append(silence(GAP, sample_rate))
    clips.append(silence(TAIL, sample_rate))
    narration = np.concatenate(clips)
    sf.write(VOICE_DIR / "narration.wav", narration, sample_rate)
    print(f"narration: {len(narration) / sample_rate:.3f}s at {sample_rate}Hz")


if __name__ == "__main__":
    main()
