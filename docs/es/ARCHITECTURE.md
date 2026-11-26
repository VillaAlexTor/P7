# Arquitectura y modelo de confianza

## Forma del sistema

```text
Navegador (Next.js, TypeScript)
        │ POST /api/v1/analyze
        ▼
API HTTP en Go
  ├─ límites de entrada y selección de idioma
  ├─ normalización y segmentación
  ├─ clasificación semántica estructurada con Ollama
  ├─ comprobación textual de citas
  ├─ ponderación determinista por perfil
  └─ mapa, vacíos, escenarios y acciones
```

El MVP no tiene cuentas ni base de datos. El texto se procesa en memoria y la aplicación no lo conserva intencionalmente.

## Flujo de análisis

1. Normaliza saltos y espacios sin reescribir el significado.
2. Divide el documento en identificadores estables de cláusula.
3. Envía cada lote de cláusulas a `qwen3:4b-instruct` local en siete pasadas centradas en categorías, cada una con un esquema JSON restringido.
4. Recibe categorías semánticas, identificadores de cláusula y citas literales propuestas.
5. Rechaza categorías no admitidas o citas que no coincidan con su cláusula tras normalizar espacios.
6. Clasifica controles como especificados, ambiguos o ausentes.
7. Aplica multiplicadores transparentes por perfil y limita el índice a 100.
8. Deriva mapa de flujo, escenarios posibles y acciones.
9. Devuelve SHA-256 de la versión normalizada.

## Siete diferenciadores en el código

| Diferenciador | Implementación |
| --- | --- |
| Data Mapping | `buildFlow` produce nodos y conexiones desde hallazgos con citas comprobadas textualmente. |
| Risk Engine | `calculateScore` usa pesos deterministas y multiplicadores por perfil. |
| Gap Detection | `detectGaps` evalúa siete controles de información. |
| Evidence Engine | `verifyEvidence` rechaza citas vacías, cláusulas incorrectas y texto que no coincide; no valida el respaldo semántico. |
| Impact Simulation | `buildGuidance` emite escenarios etiquetados como hipotéticos. |
| Smart Consent | La misma capa crea acciones priorizadas e interactivas. |
| Privacy for Everyone | Interfaz, ejemplos, prompts, errores e informes funcionan en español e inglés. |

## Límites de confianza

El documento es una entrada no confiable. El prompt indica al modelo que trate la política como datos y no como instrucciones; la inyección de instrucciones sigue siendo un riesgo del modelo. Los patrones parecidos a instrucciones se cuentan en el informe de seguridad. El cuerpo está limitado a 1 MiB y el documento a 200.000 caracteres. Los orígenes CORS se configuran con `P7_ALLOWED_ORIGINS`.

Ollama ejecuta `qwen3:4b-instruct` localmente. Las siete pasadas cubren ubicación, contactos, biometría, terceros, publicidad, entrenamiento de IA y retención por separado. Cada una devuelve candidatos restringidos por esquema: categoría, identificador de cláusula y cita propuesta. Go rechaza categorías no admitidas y citas que no coinciden textualmente con su cláusula tras normalizar espacios, deduplica, aplica pesos por perfil y calcula el índice. Esto no demuestra que la cita sostenga la categoría. Siete llamadas al modelo por lote aumentan la latencia y deben medirse con políticas completas. Si Ollama no está disponible, el endpoint devuelve un error de servicio en vez de sustituir silenciosamente la IA por palabras clave.

## Decisiones deliberadas de alcance

- Sin autenticación para reducir superficie de ataque.
- Sin base de datos para no conservar políticas potencialmente sensibles.
- Sin lectura por URL para evitar SSRF y contenido dinámico en el MVP.
- Sin conclusiones legales: se explican prácticas declaradas y vacíos.
- Sin afirmar observación de red: el mapa refleja únicamente el texto.

