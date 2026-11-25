import { existsSync, readdirSync, readFileSync } from "node:fs";
import { resolve } from "node:path";

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

const rootPairs = [
  ["README.md", "README.es.md"],
  ["PRODUCT.md", "PRODUCT.es.md"],
  ["DESIGN.md", "DESIGN.es.md"],
  ["CONTRIBUTING.md", "CONTRIBUTING.es.md"],
  ["SECURITY.md", "SECURITY.es.md"],
];

for (const [enFile, esFile] of rootPairs) {
  if (!existsSync(resolve(root, enFile)) || !existsSync(resolve(root, esFile))) {
    throw new Error(`Missing root language pair: ${enFile} / ${esFile}`);
  }
}

console.log(`Documentation parity passed: ${english.length} paired documents and ${rootPairs.length} root pairs.`);

