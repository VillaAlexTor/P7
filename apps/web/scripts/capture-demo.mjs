import { mkdir, writeFile } from "node:fs/promises";
import path from "node:path";
import { fileURLToPath } from "node:url";
import { chromium } from "playwright-core";

const here = path.dirname(fileURLToPath(import.meta.url));
const output = path.resolve(here, "../../../video/p7-demo/assets/screens");
const browserPath = process.env.P7_BROWSER_PATH ?? "C:\\Program Files\\BraveSoftware\\Brave-Browser\\Application\\brave.exe";
await mkdir(output, { recursive: true });

const browser = await chromium.launch({ executablePath: browserPath, headless: true });
const page = await browser.newPage({ viewport: { width: 1920, height: 1080 }, deviceScaleFactor: 1, reducedMotion: "reduce" });

async function shot(name) {
  await page.screenshot({ path: path.join(output, name), type: "jpeg", quality: 92 });
}

async function scrollTo(selector) {
  await page.locator(selector).scrollIntoViewIfNeeded();
  await page.waitForTimeout(250);
}

async function analyze(analyzingShot) {
  const responsePromise = page.waitForResponse((response) => response.url().endsWith("/api/v1/analyze") && response.request().method() === "POST", { timeout: 310_000 });
  await page.getByRole("button", { name: "Analyze policy" }).click();
  await shot(analyzingShot);
  const response = await responsePromise;
  const result = await response.json();
  if (!response.ok()) throw new Error(`Analysis failed (${response.status()}): ${JSON.stringify(result)}`);
  await page.locator(".report").waitFor({ state: "visible", timeout: 310_000 });
  return result;
}

try {
  await page.goto("http://localhost:3000", { waitUntil: "networkidle" });
  await page.getByRole("button", { name: "EN", exact: true }).click();
  await page.getByRole("button", { name: "Load example" }).click();
  await page.getByRole("radio", { name: /Journalist/ }).check({ force: true });
  await shot("02-input-en.jpg");

  const journalist = await analyze("02-analyzing-en.jpg");
  await scrollTo(".report");
  await shot("03-summary-en.jpg");
  await scrollTo("#mapa");
  await shot("04-map-en.jpg");
  await scrollTo("#hallazgos");
  await shot("05-findings-en.jpg");
  await scrollTo("#vacios");
  await shot("06-gaps-en.jpg");
  await scrollTo("#impacto");
  await shot("07-impact-en.jpg");
  await scrollTo("#acciones");
  await shot("08-actions-en.jpg");

  await page.getByRole("radio", { name: /General/ }).scrollIntoViewIfNeeded();
  await page.getByRole("radio", { name: /General/ }).check({ force: true });
  await shot("09-profile-general-en.jpg");
  const general = await analyze("09-analyzing-general-en.jpg");
  await scrollTo(".report");
  await shot("10-summary-general-en.jpg");

  const run = {
    journalist: { score: journalist.score, findings: journalist.findings.map((item) => item.category), result: journalist },
    general: { score: general.score, findings: general.findings.map((item) => item.category), result: general },
  };
  await writeFile(path.join(output, "demo-run.json"), JSON.stringify(run, null, 2) + "\n", "utf8");
  console.log(JSON.stringify({ journalist: run.journalist, general: run.general }, (key, value) => key === "result" ? undefined : value, 2));
} finally {
  await browser.close();
}
