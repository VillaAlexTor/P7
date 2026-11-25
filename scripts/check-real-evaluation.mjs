import { readFileSync } from "node:fs";
import { resolve } from "node:path";

const root = resolve(import.meta.dirname, "..");
const evaluation = resolve(root, "apps", "api", "evaluation");
const categories = new Set(["location", "contacts", "third-parties", "advertising", "ai-training", "biometric", "retention"]);

function rows(file, header) {
  const lines = readFileSync(resolve(evaluation, file), "utf8").trim().split(/\r?\n/);
  if (lines.shift() !== header) throw new Error(`${file}: unexpected header`);
  return lines.map((line) => line.split("\t"));
}

const urlsBySet = [];
for (const name of ["real_dev", "real_holdout", "real_holdout_90", "real_holdout_90_final", "real_round3", "real_round4", "real_round5"]) {
  const cases = rows(`${name}.tsv`, "id\tlocale\ttext\texpected");
  const sources = rows(`${name}_sources.tsv`, "id\turl");
  if (cases.length !== 20 || sources.length !== 20) throw new Error(`${name}: expected 20 cases and sources`);
  const sourceMap = new Map(sources);
  if (sourceMap.size !== 20) throw new Error(`${name}: duplicate source ID`);
  const ids = new Set();
  const locales = { en: 0, es: 0 };
  const wordsByUrl = new Map();
  for (const [id, locale, clause, expected] of cases) {
    if (!id || ids.has(id) || !clause || !expected || !(locale in locales)) throw new Error(`${name}: invalid case ${id}`);
    ids.add(id);
    locales[locale]++;
    for (const label of expected.split(",")) {
      if (label !== "-" && !categories.has(label)) throw new Error(`${name}: invalid label ${label}`);
    }
    const url = sourceMap.get(id);
    if (!url || !url.startsWith("https://")) throw new Error(`${name}: missing HTTPS source for ${id}`);
    wordsByUrl.set(url, (wordsByUrl.get(url) ?? 0) + clause.trim().split(/\s+/).length);
  }
  if (locales.en !== 10 || locales.es !== 10) throw new Error(`${name}: expected 10 cases per language`);
  for (const [url, words] of wordsByUrl) {
    if (words > 25) throw new Error(`${name}: ${words} quoted words from ${url}`);
  }
  urlsBySet.push(new Set(sourceMap.values()));
}

for (let i = 0; i < urlsBySet.length; i++) {
  for (let j = i + 1; j < urlsBySet.length; j++) {
    for (const url of urlsBySet[i]) {
      if (urlsBySet[j].has(url)) throw new Error(`Source repeated across sets: ${url}`);
    }
  }
}
console.log("Real-source evaluation metadata passed: 140 cases, 140 source links, 20 per set, 10 per language in each set.");
