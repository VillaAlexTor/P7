export type Locale = "es" | "en";

export const samples: Record<Locale, { title: string; policy: string }> = {
  es: {
    title: "SocialApp Demo",
    policy: `Política de privacidad de SocialApp Demo

Recopilamos tu nombre, correo electrónico, ubicación precisa, contactos y actividad dentro de la aplicación para proporcionar y mejorar nuestros servicios.

Podemos compartir información personal con proveedores de servicios, socios comerciales y otros terceros para analítica y publicidad personalizada.

Conservaremos la información mientras sea necesario para cumplir nuestros fines comerciales y obligaciones aplicables.

Puedes solicitar acceso o corrección de ciertos datos escribiendo a privacidad@socialapp.example.

También podemos usar el contenido enviado para entrenar modelos de inteligencia artificial y mejorar funciones automatizadas.`,
  },
  en: {
    title: "SocialApp Demo",
    policy: `SocialApp Demo Privacy Policy

We collect your name, email address, precise location, contacts, and in-app activity to provide and improve our services.

We may share personal information with service providers, business partners, and other third parties for analytics and personalized advertising.

We retain information as long as necessary for our business purposes and applicable obligations.

You may request access to or correction of certain data by emailing privacy@socialapp.example.

We may also use submitted content for model training and to improve artificial intelligence features.`,
  },
};

export const profiles = {
  es: [
    { value: "general", label: "Uso general", detail: "Equilibrio entre todos los tipos de datos" },
    { value: "student", label: "Estudiante", detail: "Prioriza perfilado, contenido y entrenamiento de IA" },
    { value: "journalist", label: "Periodista", detail: "Prioriza ubicación, contactos y terceros" },
    { value: "parent", label: "Madre, padre o tutor", detail: "Prioriza ubicación, biometría y contactos" },
    { value: "activist", label: "Activista", detail: "Máxima sensibilidad a ubicación y redes personales" },
  ],
  en: [
    { value: "general", label: "General use", detail: "Balances all data categories" },
    { value: "student", label: "Student", detail: "Prioritizes profiling, content, and AI training" },
    { value: "journalist", label: "Journalist", detail: "Prioritizes location, contacts, and third parties" },
    { value: "parent", label: "Parent or guardian", detail: "Prioritizes location, biometrics, and contacts" },
    { value: "activist", label: "Activist", detail: "Highest sensitivity to location and personal networks" },
  ],
} satisfies Record<Locale, Array<{ value: string; label: string; detail: string }>>;

export const copy = {
  es: {
    homeLabel: "p7, inicio", mainNavLabel: "Navegación principal", reportNavLabel: "Secciones del informe", pillarsLabel: "Los siete pilares de p7", flowLabel: "Mapa del flujo declarado de datos", controlsLabel: "Controles de privacidad", characters: "caracteres", fallbackError: "No se pudo completar el análisis.", dataLabel: "Datos",
    tagline: "Privacidad, descifrada.", analyzeNav: "Analizar", methodNav: "Método", resultNav: "Resultado", prototype: "MVP abierto",
    heroA: "Lee lo que aceptas.", heroB: "Sigue el rastro de tus datos.",
    heroBody: "p7 convierte políticas complejas en una evaluación personal, muestra la cláusula que respalda cada hallazgo y explica qué puedes hacer antes de aceptar.",
    principle: "La IA interpreta. El motor verifica la evidencia y calcula el índice con reglas reproducibles.",
    newReading: "Nueva lectura", memory: "El texto se procesa en memoria. Este MVP no lo almacena.", loadExample: "Cargar ejemplo",
    serviceName: "Nombre del servicio", optional: "opcional", servicePlaceholder: "Ej. Plataforma educativa", policy: "Política de privacidad",
    policyPlaceholder: "Pega aquí la política o los términos que quieres comprender…", max: "Máximo 200 000", protect: "¿Qué quieres proteger más?",
    analyze: "Analizar política", analyzing: "Revisando cláusulas…", errorTitle: "No pudimos analizar todavía.", shortError: "Pega una política de al menos 80 caracteres o carga el ejemplo.", connectionError: "No se pudo conectar con el motor de análisis.",
    pillars: ["Mapa de datos", "Riesgo personal", "Vacíos", "Evidencia", "Impacto", "Acciones", "Acceso bilingüe"],
    emptyTitle: "Tu lectura aparecerá aquí", emptyBody: "Obtendrás un mapa del flujo de datos, hallazgos vinculados a cláusulas, vacíos informativos y un plan de acción personal.",
    footerProject: "Proyecto de código abierto por Villa Cyber Defense.", footerLegal: "Herramienta informativa. No sustituye asesoramiento jurídico ni observa el tráfico real de una aplicación.",
    score: "Índice personal", profile: "Perfil", findings: "Hallazgos", gaps: "Vacíos", evidence: "Evidencias", verified: "verificadas", clauses: "cláusulas",
    map: "Mapa", impact: "Impacto", actions: "Acciones", documentIntegrity: "Integridad documental",
    flowTitle: "El recorrido declarado de tus datos", flowNote: "Representa las prácticas descritas en el documento; no es una medición del tráfico real.",
    findingTitle: "Hallazgos con respaldo", findingNote: "Cada interpretación conserva la cláusula exacta utilizada como evidencia.", verifiedQuote: "Cita verificada",
    gapsTitle: "Lo que el documento no aclara", gapsNote: "Una ausencia informativa no demuestra por sí sola una infracción legal.",
    impactTitle: "Qué podría significar para ti", impactNote: "Estos son escenarios hipotéticos, no predicciones de daño.", ifShare: "Si compartes", possibleScenario: "Escenario posible",
    actionsTitle: "Decide con más información", actionsNote: "Marca los pasos que quieras revisar antes o después de aceptar.",
    methodTitle: "Cómo se obtuvo este resultado", methodBody: "p7 segmentó el documento, aplicó reglas de detección, comprobó que cada cita existe y ponderó los hallazgos según tu perfil. El índice no es una probabilidad estadística.",
    inputSafety: "El documento se trató únicamente como datos", ignoredInstructions: "patrones de instrucciones ignorados",
    noFindings: "No encontramos prácticas de alto impacto con las reglas actuales. Revisa igualmente los vacíos informativos.",
    source: "Origen", control: "Control", destination: "Destino", processing: "Proceso",
    status: { specified: "Especificado", ambiguous: "Ambiguo", missing: "No encontrado" },
    level: { low: "Bajo", medium: "Medio", high: "Alto", critical: "Crítico" },
  },
  en: {
    homeLabel: "p7, home", mainNavLabel: "Main navigation", reportNavLabel: "Report sections", pillarsLabel: "The seven p7 pillars", flowLabel: "Declared data flow map", controlsLabel: "Privacy controls", characters: "characters", fallbackError: "The analysis could not be completed.", dataLabel: "Data",
    tagline: "Privacy. Decoded.", analyzeNav: "Analyze", methodNav: "Method", resultNav: "Result", prototype: "Open MVP",
    heroA: "Read what you accept.", heroB: "Follow your data trail.",
    heroBody: "p7 turns complex privacy policies into a personal assessment, shows the clause behind every finding, and explains what you can do before accepting.",
    principle: "AI interprets. The engine verifies evidence and calculates the index with reproducible rules.",
    newReading: "New reading", memory: "Text is processed in memory. This MVP does not store it.", loadExample: "Load example",
    serviceName: "Service name", optional: "optional", servicePlaceholder: "Example: Learning platform", policy: "Privacy policy",
    policyPlaceholder: "Paste the policy or terms you want to understand…", max: "Maximum 200,000", protect: "What matters most to you?",
    analyze: "Analyze policy", analyzing: "Reviewing clauses…", errorTitle: "We could not analyze it yet.", shortError: "Paste at least 80 characters from a policy or load the example.", connectionError: "Could not connect to the analysis engine.",
    pillars: ["Data map", "Personal risk", "Disclosure gaps", "Evidence", "Impact", "Actions", "Bilingual access"],
    emptyTitle: "Your reading will appear here", emptyBody: "You will receive a data-flow map, clause-linked findings, disclosure gaps, and a personal action plan.",
    footerProject: "Open-source project by Villa Cyber Defense.", footerLegal: "Informational tool. It does not replace legal advice or observe an application's real network traffic.",
    score: "Personal index", profile: "Profile", findings: "Findings", gaps: "Gaps", evidence: "Evidence", verified: "verified", clauses: "clauses",
    map: "Map", impact: "Impact", actions: "Actions", documentIntegrity: "Document integrity",
    flowTitle: "The declared path of your data", flowNote: "Represents practices described in the document; it does not measure real network traffic.",
    findingTitle: "Evidence-backed findings", findingNote: "Every interpretation keeps the exact clause used as evidence.", verifiedQuote: "Verified quote",
    gapsTitle: "What the document does not clarify", gapsNote: "Missing information does not prove a legal violation by itself.",
    impactTitle: "What it could mean for you", impactNote: "These are hypothetical scenarios, not predictions of harm.", ifShare: "If you share", possibleScenario: "Possible scenario",
    actionsTitle: "Decide with more information", actionsNote: "Mark the steps you want to review before or after accepting.",
    methodTitle: "How this result was produced", methodBody: "p7 segmented the document, applied detection rules, verified every quote, and weighted findings for your profile. The index is not a statistical probability.",
    inputSafety: "The document was treated only as data", ignoredInstructions: "instruction patterns ignored",
    noFindings: "The current rules found no high-impact practices. Review the disclosure gaps as well.",
    source: "Source", control: "Control", destination: "Destination", processing: "Process",
    status: { specified: "Specified", ambiguous: "Ambiguous", missing: "Not found" },
    level: { low: "Low", medium: "Medium", high: "High", critical: "Critical" },
  },
} as const;
