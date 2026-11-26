import type { AnalysisResponse } from "./types";

const REQUEST_TIMEOUT_MS = 5 * 60 * 1000;

export interface AnalyzePayload {
  text: string;
  profile: string;
  title: string;
  locale: string;
}

export class RequestTimeoutError extends Error {
  constructor() {
    super("Analysis request timed out");
    this.name = "RequestTimeoutError";
  }
}

export function isRequestTimeout(cause: unknown): cause is RequestTimeoutError {
  return cause instanceof RequestTimeoutError;
}

export async function analyzePolicy(apiURL: string, payload: AnalyzePayload): Promise<AnalysisResponse> {
  const controller = new AbortController();
  const timeout = window.setTimeout(() => controller.abort(), REQUEST_TIMEOUT_MS);
  try {
    const response = await fetch(`${apiURL}/api/v1/analyze`, {
      method: "POST",
      headers: { "Content-Type": "application/json", "Accept-Language": payload.locale },
      body: JSON.stringify(payload),
      signal: controller.signal,
    });
    const contentType = response.headers.get("content-type") ?? "";
    const body = contentType.includes("application/json") ? await response.json() : null;
    if (!response.ok) {
      throw new Error(body?.error || `Analysis failed (${response.status})`);
    }
    if (!body || typeof body !== "object") {
      throw new Error("The analysis service returned an invalid response.");
    }
    return body as AnalysisResponse;
  } catch (cause) {
    if (controller.signal.aborted) {
      throw new RequestTimeoutError();
    }
    throw cause;
  } finally {
    window.clearTimeout(timeout);
  }
}
