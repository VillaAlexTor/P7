import { describe, expect, it } from "vitest";
import { copy, profiles, samples } from "./i18n";

describe("bilingual product contract", () => {
  it("keeps the same translation keys in Spanish and English", () => {
    expect(Object.keys(copy.es).sort()).toEqual(Object.keys(copy.en).sort());
    expect(Object.keys(copy.es.status).sort()).toEqual(Object.keys(copy.en.status).sort());
    expect(Object.keys(copy.es.level).sort()).toEqual(Object.keys(copy.en.level).sort());
  });

  it("exposes all seven differentiators in both languages", () => {
    expect(copy.es.pillars).toHaveLength(7);
    expect(copy.en.pillars).toHaveLength(7);
    expect(new Set(copy.es.pillars).size).toBe(7);
    expect(new Set(copy.en.pillars).size).toBe(7);
  });

  it("keeps equivalent profile identifiers", () => {
    expect(profiles.es.map((profile) => profile.value)).toEqual(profiles.en.map((profile) => profile.value));
    expect(profiles.es).toHaveLength(5);
  });

  it("ships usable sample policies for both locales", () => {
    expect(samples.es.policy.length).toBeGreaterThan(80);
    expect(samples.en.policy.length).toBeGreaterThan(80);
    expect(samples.es.title).toBe(samples.en.title);
  });
});
