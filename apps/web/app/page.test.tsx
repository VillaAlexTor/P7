// @vitest-environment jsdom

import { cleanup, render, screen, waitFor } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { afterEach, describe, expect, it, vi } from "vitest";
import Home from "./page";
import type { AnalysisResponse } from "@/lib/types";

const report: AnalysisResponse = {
  meta: {
    title: "Sample Service",
    profile: "journalist",
    documentHash: "0123456789abcdef0123456789abcdef",
    clauseCount: 7,
    language: "en",
    engine: "p7-hybrid/0.3",
    evidenceChecked: 4,
    locale: "en",
    aiUsed: true,
    aiProvider: "Ollama",
    aiModel: "qwen3:4b-instruct",
    aiCandidates: 4,
  },
  score: 97,
  level: "critical",
  summary: "Four relevant practices require review.",
  clauses: [],
  findings: ["location", "third-parties", "advertising", "retention"].map((category, index) => ({
    id: `finding-${index}`,
    category,
    title: category,
    severity: "high" as const,
    interpretation: "The policy describes this practice.",
    evidence: { clauseId: `clause-${index}`, quote: "Verified policy evidence.", verified: true },
    weight: 20,
  })),
  gaps: [{ control: "Account deletion", status: "missing", note: "The policy should explain deletion." }],
  nodes: [{ id: "you", label: "You", kind: "person" }, { id: "service", label: "Service", kind: "service" }],
  edges: [{ from: "you", to: "service", dataLabel: "Declared data" }],
  impacts: [{ dataType: "Location", scenario: "Location can reveal routines.", confidence: "scenario" }],
  actions: [{ id: "action-1", priority: "high", text: "Limit location access." }],
  safety: { inputTreatedAsData: true, ignoredInstructionPatterns: 0, citationsVerified: 4, citationsRejected: 0 },
};

function jsonResponse(body: unknown, status = 200): Response {
  return {
    ok: status >= 200 && status < 300,
    status,
    headers: new Headers({ "content-type": "application/json" }),
    json: async () => body,
  } as Response;
}

afterEach(() => {
  cleanup();
  vi.unstubAllGlobals();
});

describe("analysis form", () => {
  it("shows loading, sends the selected profile, and renders the report", async () => {
    let finish!: (response: Response) => void;
    const fetchMock = vi.fn<typeof fetch>(() => new Promise<Response>((resolve) => { finish = resolve; }));
    vi.stubGlobal("fetch", fetchMock);
    Element.prototype.scrollIntoView = vi.fn();
    const user = userEvent.setup();
    render(<Home />);

    await user.click(screen.getByRole("button", { name: /cargar ejemplo/i }));
    await user.click(screen.getByRole("radio", { name: /periodista/i }));
    await user.click(screen.getByRole("button", { name: /analizar política/i }));

		expect((screen.getByRole("button", { name: /revisando cláusulas/i }) as HTMLButtonElement).disabled).toBe(true);
    const sent = JSON.parse(String(fetchMock.mock.calls[0][1]?.body));
    expect(sent.profile).toBe("journalist");

    finish(jsonResponse(report));
		expect(await screen.findByText("97")).toBeTruthy();
		expect(screen.getAllByText("The policy describes this practice.")).toHaveLength(4);
  });

  it("shows an API error and restores the submit button", async () => {
    vi.stubGlobal("fetch", vi.fn(async () => jsonResponse({ error: "La IA local no está disponible." }, 503)));
    const user = userEvent.setup();
    render(<Home />);

    await user.click(screen.getByRole("button", { name: /cargar ejemplo/i }));
    await user.click(screen.getByRole("button", { name: /analizar política/i }));

		expect((await screen.findByRole("alert")).textContent).toContain("La IA local no está disponible.");
		await waitFor(() => expect((screen.getByRole("button", { name: /analizar política/i }) as HTMLButtonElement).disabled).toBe(false));
  });
});
