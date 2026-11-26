import { existsSync, readdirSync, readFileSync } from "node:fs";
import { join, relative, resolve } from "node:path";

const root = resolve(import.meta.dirname, "..");
const enDir = resolve(root, "docs", "en");
const esDir = resolve(root, "docs", "es");

function markdownFiles(directory) {
  if (!existsSync(directory)) throw new Error(`Missing documentation directory: ${directory}`);
  return readdirSync(directory).filter((file) => file.endsWith(".md")).sort();
}

const english = markdownFiles(enDir);
const spanish = markdownFiles(esDir);

if (JSON.stringify(english) !== JSON.stringify(spanish)) {
  throw new Error(`Documentation parity failed. EN=${english.join(",")} ES=${spanish.join(",")}`);
}

for (const file of english) {
  const en = readFileSync(resolve(enDir, file), "utf8").trim();
  const es = readFileSync(resolve(esDir, file), "utf8").trim();
  if (en.length < 100 || es.length < 100) throw new Error(`${file} is incomplete in one language.`);
  if (!en.startsWith("#") || !es.startsWith("#")) throw new Error(`${file} must start with a Markdown heading.`);
}

// Video production and dependency caches are separate workspaces. Official
// project Markdown is checked here; the video project can validate its own
// production files when it is ready to submit.
const excluded = new Set([".git", ".cache", "node_modules", ".next", "dist", "build", "coverage", "video"]);
function collectMarkdown(directory) {
  return readdirSync(directory, { withFileTypes: true }).flatMap((entry) => {
    const path = join(directory, entry.name);
    if (entry.isDirectory()) return excluded.has(entry.name) ? [] : collectMarkdown(path);
    return entry.isFile() && entry.name.endsWith(".md") ? [relative(root, path).replaceAll("\\", "/")] : [];
  });
}

const files = new Set(collectMarkdown(root));
let adjacentPairs = 0;
for (const file of files) {
  if (file.startsWith("docs/en/") || file.startsWith("docs/es/")) continue;
  if (file.endsWith(".es.md")) {
    const counterpart = file.replace(/\.es\.md$/, ".md");
    if (!files.has(counterpart)) throw new Error(`Missing English counterpart: ${counterpart}`);
    continue;
  }
  const counterpart = file.replace(/\.md$/, ".es.md");
  if (!files.has(counterpart)) throw new Error(`Missing Spanish counterpart: ${counterpart}`);
  const en = readFileSync(resolve(root, file), "utf8").trim();
  const es = readFileSync(resolve(root, counterpart), "utf8").trim();
  if (!en || !es) throw new Error(`Incomplete language pair: ${file} / ${counterpart}`);
  adjacentPairs++;
}

console.log(`Documentation parity passed: ${english.length} paired documents and ${adjacentPairs} adjacent pairs.`);

