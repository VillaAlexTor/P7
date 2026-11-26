from __future__ import annotations

import argparse
import inspect
import json
import re
from pathlib import Path

import kokoro_onnx
import soundfile as sf


def split_for_retry(value: str) -> tuple[str, str]:
    midpoint = len(value) // 2
    boundaries = [
        index + 1
        for index, char in enumerate(value)
        if (char in ".!?\n" or char.isspace()) and 0 < index < len(value) - 1
    ]
    split_at = min(boundaries, key=lambda index: abs(index - midpoint)) if boundaries else midpoint
    return value[:split_at], value[split_at:]


def is_token_overflow(error: IndexError) -> bool:
    return bool(re.search(r"index \d+ is out of bounds for axis 0 with size 510", str(error), re.I))


def main() -> None:
    parser = argparse.ArgumentParser()
    parser.add_argument("--model", required=True)
    parser.add_argument("--voices", required=True)
    parser.add_argument("--text-file", required=True)
    parser.add_argument("--voice", default="af_nova")
    parser.add_argument("--speed", type=float, default=1.0)
    parser.add_argument("--output", required=True)
    args = parser.parse_args()

    text = Path(args.text_file).read_text(encoding="utf-8").strip()
    model = kokoro_onnx.Kokoro(args.model, args.voices)
    create_params = inspect.signature(model.create).parameters

    def synthesize(value: str):
        kwargs = {"voice": args.voice, "speed": args.speed}
        if "lang" in create_params:
            kwargs["lang"] = "en-us"
        try:
            return model.create(value, **kwargs)
        except IndexError as error:
            if not is_token_overflow(error) or len(value) < 2:
                raise
            left, right = split_for_retry(value)
            left_samples, left_rate = synthesize(left)
            right_samples, right_rate = synthesize(right)
            if left_rate != right_rate:
                raise RuntimeError("Kokoro returned inconsistent sample rates")
            return list(left_samples) + list(right_samples), left_rate

    samples, sample_rate = synthesize(text)
    output = Path(args.output)
    output.parent.mkdir(parents=True, exist_ok=True)
    sf.write(output, samples, sample_rate)
    print(json.dumps({"output": str(output), "duration": round(len(samples) / sample_rate, 3)}))


if __name__ == "__main__":
    main()
