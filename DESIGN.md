# p7 Design System

## Direction

p7 uses an evidence desk rather than a conventional software dashboard. The interface borrows the hierarchy of an annotated privacy dossier: warm paper, precise black rules, editorial serif headings, compact sans-serif controls, and color that appears only when it carries meaning.

The first viewport is an operating surface. Users can paste a policy and choose a risk profile immediately; marketing copy remains secondary to the task.

## Modes

- Primary mode: operate.
- Secondary mode: read, inside evidence and interpretation sections.
- The product should feel calm and exact. Avoid alarmist cyber-security imagery, neon, glass effects, generic AI gradients, and grids of interchangeable cards.

## Color

| Token | Value | Use |
| --- | --- | --- |
| Paper | `#f3efe4` | Main surface |
| Deep paper | `#e8e1d2` | Shoulder navigation and grouped facts |
| Ink | `#171711` | Primary text, rules, dark controls |
| Soft ink | `#5b5a50` | Supporting copy |
| Vermilion | `#d64a32` | High and critical risk, active analysis cue |
| Amber | `#b57312` | Medium risk or ambiguity |
| Emerald | `#18705b` | Verified or specified evidence |
| Ultramarine | `#244fa3` | Keyboard focus |
| Violet | `#6b4e87` | Hypothetical impact and processing |

Do not use color as the only signal. Every risk state also has a label.

## Typography

- Display and reading face: Source Serif 4 Variable.
- Interface face: Manrope Variable.
- Use the serif for titles, summaries, quotations, and evidence-led reading.
- Use the sans serif for navigation, labels, controls, metadata, and compact tables.
- Keep body text between 65 and 75 characters when the layout permits.

## Composition

- Thin black rules establish structure; avoid shadowed cards.
- Desktop task entry is a split workbench with an editorial statement on the left and the form on the right.
- Analysis uses a narrow shoulder for score and navigation, plus a wide reading column for evidence.
- Each report section starts with one strong heading and a concise caveat.
- Mobile stacks the shoulder before the report and makes long diagrams horizontally scrollable within their own boundary.

## Controls and States

- Primary actions use solid ink; vermilion appears on hover and during active analysis.
- Text inputs use a single baseline rather than a rounded container.
- Radio and checkbox controls remain native-accessible but receive custom visible marks.
- Loading text names the current operation: “Revisando cláusulas…”.
- Errors state both the problem and the recovery.
- Focus is a two-pixel ultramarine outline with visible offset.

## Motion

Motion is restrained and functional. Scroll moves to the completed report after analysis; link rules extend on hover. Reduced-motion preferences disable smooth scrolling and transitions.

## Content Rules

- Separate evidence, interpretation, and possible impact.
- Never describe missing text as proof of illegality.
- Label simulated data and hypothetical consequences.
- Never claim the product observes application traffic.
- Maintain semantic parity between the official English copy and its Spanish counterpart; retain technical terms only where they improve trust or verifiability.

