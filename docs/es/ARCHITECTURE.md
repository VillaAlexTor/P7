# Arquitectura y modelo de confianza

## Forma del sistema

```text
Navegador (Next.js, TypeScript)
        │ POST /api/v1/analyze
        ▼
API HTTP en Go
  ├─ límites de entrada y selección de idioma
  ├─ normalización y segmentación
  ├─ catálogo bilingüe de reglas
  ├─ verificación exacta de evidencia
  ├─ ponderación determinista por perfil
  └─ mapa, vacíos, escenarios y acciones
```

El MVP no tiene cuentas ni base de datos. El texto se procesa en memoria y la aplicación no lo conserva intencionalmente.

## Flujo de análisis

1. Normaliza saltos y espacios sin reescribir el significado.
2. Divide el documento en identificadores estables de cláusula.
3. Busca señales bilingües de prácticas de privacidad declaradas.
4. Copia la cláusula fuente completa en cada hallazgo.
5. Comprueba que el identificador existe y contiene la cita exacta.
6. Clasifica controles como especificados, ambiguos o ausentes.
7. Aplica multiplicadores transparentes por perfil y limita el índice a 100.
8. Deriva mapa de flujo, escenarios posibles y acciones.
9. Devuelve SHA-256 de la versión normalizada.

## Siete diferenciadores en el código

| Diferenciador | Implementación |
| --- | --- |
| Data Mapping | `buildFlow` produce nodos y conexiones desde hallazgos verificados. |
| Risk Engine | `calculateScore` usa pesos deterministas y multiplicadores por perfil. |
| Gap Detection | `detectGaps` evalúa siete controles de información. |
| Evidence Engine | `verifyEvidence` rechaza citas vacías, desconocidas o fabricadas. |
| Impact Simulation | `buildGuidance` emite escenarios etiquetados como hipotéticos. |
| Smart Consent | La misma capa crea acciones priorizadas e interactivas. |
| Privacy for Everyone | Interfaz, ejemplos, reglas, errores e informes funcionan en español e inglés. |

## Límites de confianza

El documento es una entrada no confiable. p7 nunca ejecuta instrucciones encontradas dentro. Los patrones parecidos a instrucciones se cuentan en el informe de seguridad, pero no modifican el análisis. El cuerpo está limitado a 1 MiB y el documento a 200.000 caracteres. Los orígenes CORS se configuran con `P7_ALLOWED_ORIGINS`.

Un futuro clasificador LLM opcional deberá devolver candidatos estructurados con referencia de cláusula. Go seguirá controlando validación, deduplicación, pesos y puntuación final.

## Decisiones deliberadas de alcance

- Sin autenticación para reducir superficie de ataque.
- Sin base de datos para no conservar políticas potencialmente sensibles.
- Sin lectura por URL para evitar SSRF y contenido dinámico en el MVP.
- Sin conclusiones legales: se explican prácticas declaradas y vacíos.
- Sin afirmar observación de red: el mapa refleja únicamente el texto.

