# p7 Demo Film — Director's Spec

GOAL: A judge should understand that p7 is a working, bilingual privacy-policy reader whose AI findings remain verifiable and require human review.

OUTPUT: HyperFrames, 1920×1080, 30 fps, English narration and English captions.

TIME: Target 225–235 seconds after measured narration; voice is the clock. Music sits near −31 LUFS under speech and fades over the final 2 seconds.

LOOK: Ground #f3efe5 with no texture; ink #12120f; accent #df472f only on risk, clicks, and one key phrase per chapter; secondary green #297a62 only for verified evidence. Display face EB Garamond 700 at 86–132 px with −2% tracking; body Inter 400/700 at 34–48 px. Product UI remains real screenshot material, filling 78–92% of the frame. Shadows: 0 18px 50px rgba(18,18,15,.16), no glow, no vignette, no grain.

OBJECTS: Real p7 home, input, summary, map, findings, gaps, impact, and actions screenshots; a large white cursor; one thin red evidence line; exact evaluation numbers; captions in a compact lower band.

MOTION: Scene entry 18 frames power3.out; camera push 36–54 frames power2.inOut; cursor travel 24 frames power2.inOut; click compression 4 frames and ring 12 frames; cuts happen on narration anchors. Each product scene gets one slow purposeful pan or push, never continuous floating.

SHOTS:

- 0–18 s: dense policy copy fills the frame, then collapses into “What are you accepting?”; narration opens the human problem.
- 18–40 s: real input screen; the cursor selects Journalist, visibly presses Analyze, and the button changes to “Reviewing clauses…”.
- 40–65 s: report summary and data map; camera pushes from index and SHA-256 to the declared path.
- 65–90 s: findings screen; exact verified quote receives one green evidence line, then gaps screen replaces it.
- 90–112 s: possible impact changes into actions; two checkboxes receive restrained click rings.
- 112–132 s: continuous same-policy profile comparison; General is selected, Analyze is pressed again, and the second real result changes the index from 97 to 80 while the same four findings remain.
- 132–158 s: architecture strip: local model → evidence verification → deterministic score, over real UI rather than a floating dashboard.
- 158–181 s: the name p7 resolves into its seven pillars; the seventh ends on the real Spanish interface and plain-language commitment.
- 181–213 s: evaluation story; twelve early denial clauses, then the latest revised full-policy result: 87.5% precision, 70.0% recall, and 0/4 exact localized pages.
- 213–end: action screen returns, then p7 title and the line “Evidence before confidence.”

BANNED: black or navy void; blue/purple gradients; glow, particles, glass cards; cream-paper plus decorative mono labels; three cards in a row; tiny UI floating in space; docked logo; generic tech icons; repeated fade-and-slide entrances; text drifting while read; exaggerated claims; hiding the complete-policy evaluation result.

QA: Check opening frame, every screenshot change, exact numbers at the evaluation shot, caption legibility at 360 px, no crop of product labels, final hold of at least 2.5 seconds, and audible narration/music measured after render.
