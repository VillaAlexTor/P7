"use client";

import { FormEvent, useEffect, useMemo, useRef, useState } from "react";
import { copy, type Locale, profiles, samples } from "@/lib/i18n";
import type { AnalysisResponse, GapStatus, Severity } from "@/lib/types";

const API_URL = process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080";

export default function Home() {
  const [locale, setLocale] = useState<Locale>("es");
  const [text, setText] = useState("");
  const [profile, setProfile] = useState("general");
  const [title, setTitle] = useState("");
  const [result, setResult] = useState<AnalysisResponse | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");
  const [checked, setChecked] = useState<Record<string, boolean>>({});
  const resultRef = useRef<HTMLElement>(null);
  const t = copy[locale];
  const profileOptions = profiles[locale];

  const selectedProfile = useMemo(
    () => profileOptions.find((item) => item.value === profile) ?? profileOptions[0],
    [profile, profileOptions],
  );

  useEffect(() => {
    document.documentElement.lang = locale;
  }, [locale]);

  async function submit(event: FormEvent) {
    event.preventDefault();
    setError("");
    if (text.trim().length < 80) {
      setError(t.shortError);
      return;
    }
    setLoading(true);
    try {
      const response = await fetch(`${API_URL}/api/v1/analyze`, {
        method: "POST",
        headers: { "Content-Type": "application/json", "Accept-Language": locale },
        body: JSON.stringify({ text, profile, title: title || (locale === "es" ? "Política analizada" : "Analyzed policy"), locale }),
      });
      const body = await response.json();
      if (!response.ok) throw new Error(body.error || t.fallbackError);
      setResult(body);
      setChecked({});
      window.setTimeout(() => resultRef.current?.scrollIntoView({ behavior: "smooth", block: "start" }), 80);
    } catch (cause) {
      setError(cause instanceof Error ? cause.message : t.connectionError);
    } finally {
      setLoading(false);
    }
  }

  function loadExample() {
    setText(samples[locale].policy);
    setTitle(samples[locale].title);
    setError("");
  }

  function switchLocale(next: Locale) {
    if (next === locale) return;
    const currentSample = samples[locale];
    const nextSample = samples[next];
    if (text === currentSample.policy) setText(nextSample.policy);
    if (title === currentSample.title) setTitle(nextSample.title);
    setLocale(next);
    setResult(null);
    setChecked({});
    setError("");
  }

  return (
    <main>
      <header className="topbar">
        <a className="brand" href="#inicio" aria-label={t.homeLabel}>
          <span className="brandMark">p7</span>
          <span className="brandLine">{t.tagline}</span>
        </a>
        <nav aria-label={t.mainNavLabel}>
          <a href="#analizar">{t.analyzeNav}</a>
          <a href="#metodo">{t.methodNav}</a>
          <a href="#resultado">{t.resultNav}</a>
        </nav>
        <div className="headerActions">
          <div className="languageSwitch" aria-label="Language / Idioma">
            <button type="button" aria-pressed={locale === "es"} onClick={() => switchLocale("es")}>ES</button>
            <button type="button" aria-pressed={locale === "en"} onClick={() => switchLocale("en")}>EN</button>
          </div>
          <span className="prototypeLabel">{t.prototype}</span>
        </div>
      </header>

      <section className="workbench" id="inicio">
        <div className="workbenchIntro">
          <h1>{t.heroA}<br />{t.heroB}</h1>
          <p>{t.heroBody}</p>
          <div className="principle" id="metodo">
            <span aria-hidden="true" className="principleRule" />
            <p>{t.principle}</p>
          </div>
        </div>

        <form className="analysisForm" id="analizar" onSubmit={submit}>
          <div className="formHeading">
            <div>
              <h2>{t.newReading}</h2>
              <p>{t.memory}</p>
            </div>
            <button className="textButton" type="button" onClick={loadExample}>{t.loadExample}</button>
          </div>

          <label htmlFor="policy-title">{t.serviceName} <span>{t.optional}</span></label>
          <input
            id="policy-title"
            value={title}
            onChange={(event) => setTitle(event.target.value)}
            placeholder={t.servicePlaceholder}
            maxLength={120}
          />

          <label htmlFor="policy-text">{t.policy}</label>
          <textarea
            id="policy-text"
            value={text}
            onChange={(event) => setText(event.target.value)}
            placeholder={t.policyPlaceholder}
            aria-describedby="text-help"
            maxLength={200000}
          />
          <div className="fieldMeta" id="text-help">
            <span>{text.length.toLocaleString(locale === "es" ? "es-BO" : "en-US")} {t.characters}</span>
            <span>{t.max}</span>
          </div>

          <fieldset>
            <legend>{t.protect}</legend>
            <div className="profileList">
              {profileOptions.map((item) => (
                <label className="profileOption" key={item.value}>
                  <input type="radio" name="profile" value={item.value} checked={profile === item.value} onChange={() => setProfile(item.value)} />
                  <span className="radioMark" aria-hidden="true" />
                  <span><strong>{item.label}</strong><small>{item.detail}</small></span>
                </label>
              ))}
            </div>
          </fieldset>

          {error && <div className="formError" role="alert"><strong>{t.errorTitle}</strong><span>{error}</span></div>}

          <button className="primaryButton" type="submit" disabled={loading}>
            <span>{loading ? t.analyzing : t.analyze}</span>
            <span className="buttonCue" aria-hidden="true" />
          </button>
        </form>
      </section>

      <section className="sevenPillars" aria-label={t.pillarsLabel}>
        {t.pillars.map((pillar, index) => <span key={pillar}><b>{String(index + 1).padStart(2, "0")}</b>{pillar}</span>)}
      </section>

      {result ? (
        <AnalysisReport result={result} profileLabel={selectedProfile.label} checked={checked} setChecked={setChecked} reportRef={resultRef} locale={locale} />
      ) : (
        <section className="emptyReport" id="resultado">
          <div>
            <span className="emptyIndex">p7</span>
            <h2>{t.emptyTitle}</h2>
          </div>
          <p>{t.emptyBody}</p>
        </section>
      )}

      <footer>
        <div><span className="brandMark">p7</span><p>{t.footerProject}</p></div>
        <p>{t.footerLegal}</p>
      </footer>
    </main>
  );
}

function AnalysisReport({ result, profileLabel, checked, setChecked, reportRef, locale }: {
  result: AnalysisResponse;
  profileLabel: string;
  checked: Record<string, boolean>;
  setChecked: React.Dispatch<React.SetStateAction<Record<string, boolean>>>;
  reportRef: React.RefObject<HTMLElement | null>;
  locale: Locale;
}) {
  const t = copy[locale];
  const statusLabels: Record<GapStatus, string> = t.status;
  const levelLabels: Record<Severity, string> = t.level;
  const missing = result.gaps.filter((gap) => gap.status !== "specified").length;
  return (
    <section className="report" id="resultado" ref={reportRef} tabIndex={-1} aria-live="polite">
      <aside className="reportShoulder">
        <div className={`scoreBlock risk-${result.level}`}>
          <span>{t.score}</span>
          <strong>{result.score}<small>/100</small></strong>
          <b>{levelLabels[result.level]}</b>
        </div>
        <dl className="reportFacts">
          <div><dt>{t.profile}</dt><dd>{profileLabel}</dd></div>
          <div><dt>{t.findings}</dt><dd>{result.findings.length}</dd></div>
          <div><dt>{t.gaps}</dt><dd>{missing}</dd></div>
          <div><dt>{t.evidence}</dt><dd>{result.meta.evidenceChecked} {t.verified}</dd></div>
        </dl>
        <nav aria-label={t.reportNavLabel}>
          <a href="#mapa">{t.map}</a>
          <a href="#hallazgos">{t.findings}</a>
          <a href="#vacios">{t.gaps}</a>
          <a href="#impacto">{t.impact}</a>
          <a href="#acciones">{t.actions}</a>
        </nav>
      </aside>

      <div className="reportBody">
        <header className="reportHeader">
          <div>
            <span className="reportGuide">{result.meta.language.toUpperCase()} · {result.meta.clauseCount} {t.clauses}</span>
            <h2>{result.meta.title}</h2>
            <p>{result.summary}</p>
          </div>
          <div className="integrityStamp" title={result.meta.documentHash}>
            <span>{t.documentIntegrity}</span>
            <strong>SHA-256</strong>
            <code>{result.meta.documentHash.slice(0, 12)}…{result.meta.documentHash.slice(-8)}</code>
          </div>
        </header>

        <ReportSection id="mapa" title={t.flowTitle} note={t.flowNote}>
          <DataFlow result={result} locale={locale} />
        </ReportSection>

        <ReportSection id="hallazgos" title={t.findingTitle} note={t.findingNote}>
          <div className="findingsList">
            {result.findings.length ? result.findings.map((finding, index) => (
              <article className="finding" key={finding.id}>
                <div className="findingHead">
                  <span>{String(index + 1).padStart(2, "0")}</span>
                  <div><h4>{finding.title}</h4><p>{finding.interpretation}</p></div>
                  <b className={`severity severity-${finding.severity}`}>{levelLabels[finding.severity]}</b>
                </div>
                <blockquote>
                  <span>{finding.evidence.clauseId} · {t.verifiedQuote}</span>
                  <p>“{finding.evidence.quote}”</p>
                </blockquote>
              </article>
            )) : <p className="emptyInline">{t.noFindings}</p>}
          </div>
        </ReportSection>

        <ReportSection id="vacios" title={t.gapsTitle} note={t.gapsNote}>
          <div className="gapTable" role="table" aria-label={t.controlsLabel}>
            {result.gaps.map((gap) => (
              <div className="gapRow" role="row" key={gap.control}>
                <strong role="cell">{gap.control}</strong>
                <span role="cell" className={`gapStatus status-${gap.status}`}>{statusLabels[gap.status]}</span>
                <p role="cell">{gap.note}</p>
              </div>
            ))}
          </div>
        </ReportSection>

        <ReportSection id="impacto" title={t.impactTitle} note={t.impactNote}>
          <div className="impactList">
            {result.impacts.map((impact) => (
              <article key={impact.dataType}><h4>{t.ifShare} {impact.dataType.toLowerCase()}</h4><p>{impact.scenario}</p><span>{t.possibleScenario}</span></article>
            ))}
          </div>
        </ReportSection>

        <ReportSection id="acciones" title={t.actionsTitle} note={t.actionsNote}>
          <div className="actionList">
            {result.actions.map((action) => (
              <label key={action.id} className={checked[action.id] ? "actionChecked" : ""}>
                <input type="checkbox" checked={Boolean(checked[action.id])} onChange={() => setChecked((current) => ({ ...current, [action.id]: !current[action.id] }))} />
                <span className="checkMark" aria-hidden="true" />
                <span>{action.text}</span>
                <b className={`actionPriority severity-${action.priority}`}>{levelLabels[action.priority]}</b>
              </label>
            ))}
          </div>
        </ReportSection>

        <div className="methodNote">
          <strong>{t.methodTitle}</strong>
          <p>{t.methodBody}</p>
          <div className="methodMeta">
            <code>{result.meta.engine}</code>
            <strong>{result.meta.aiProvider} · {result.meta.aiModel}</strong>
            <span>{t.aiClassified}</span>
            <span>{result.meta.aiCandidates} {t.aiCandidates}</span>
            <span>{result.safety.citationsRejected} {t.rejectedCitations}</span>
            <span>{t.inputSafety}</span>
            <span>{result.safety.ignoredInstructionPatterns} {t.ignoredInstructions}</span>
          </div>
        </div>
      </div>
    </section>
  );
}

function ReportSection({ id, title, note, children }: { id: string; title: string; note: string; children: React.ReactNode }) {
  return <section className="reportSection" id={id}><header><h3>{title}</h3><p>{note}</p></header>{children}</section>;
}

function DataFlow({ result, locale }: { result: AnalysisResponse; locale: Locale }) {
  const t = copy[locale];
  return (
    <div className="dataFlow" aria-label={t.flowLabel}>
      {result.nodes.map((node, index) => (
        <div className="flowStep" key={node.id}>
          {index > 0 && <div className="flowConnector"><span>{result.edges[index - 1]?.dataLabel ?? t.dataLabel}</span><i aria-hidden="true" /></div>}
          <div className={`flowNode node-${node.kind}`}><span>{node.kind === "person" ? t.source : node.kind === "service" ? t.control : node.kind === "processing" ? t.processing : t.destination}</span><strong>{node.label}</strong></div>
        </div>
      ))}
    </div>
  );
}

