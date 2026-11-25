# Borrador de presentación para Devpost

## Nombre del proyecto

p7 — Privacidad, descifrada.

## Pista

Impacto comunitario

## Inspiración y problema

Las personas aceptan políticas extensas sin comprender qué datos entregan, quién puede recibirlos, qué deja sin aclarar el documento o qué pueden hacer después. Los resúmenes genéricos reducen texto, pero rara vez apoyan una decisión personal. Esto afecta especialmente a estudiantes y personas sin formación jurídica o técnica, y muchas herramientas no comienzan desde los usuarios latinoamericanos de habla hispana.

## Qué hace

p7 transforma una política en una evaluación personalizada y verificable. El usuario pega texto, elige uno de cinco perfiles y recibe:

- mapa del flujo de datos declarado;
- índice de riesgo personal determinista;
- controles de información ausente o ambigua;
- hallazgos vinculados a cláusulas exactas y verificadas;
- escenarios posibles claramente etiquetados;
- plan de acción priorizado para el consentimiento;
- experiencia completa en español e inglés.

El documento normalizado recibe un hash SHA-256 para identificar la versión exacta analizada.

## Cómo lo construimos

La aplicación usa Next.js, React y TypeScript. Una API Go normaliza y segmenta políticas, aplica reglas bilingües, verifica cada cita, evalúa siete controles, aplica pesos transparentes por perfil y deriva el mapa y las recomendaciones. El MVP no usa cuentas ni base de datos.

## Cómo usamos IA

ChatGPT y Codex apoyaron investigación de competencia, definición del problema, alcance, arquitectura, programación, diseño, pruebas y documentación. Esta asistencia está divulgada en el repositorio y la presentación.

La versión ejecutable sigue siendo reproducible sin créditos API pagados. La IA podrá proponer clasificaciones semánticas, pero no controlará el índice. Go seguirá verificando citas, deduplicando y aplicando pesos deterministas. Elegimos esta arquitectura porque una respuesta convincente no es suficiente en privacidad: cada hallazgo debe volver a la política.

## Desafíos

El desafío principal fue evitar que el producto se convierta en otro resumidor legal. Tuvimos que conectar siete funciones sin perder usabilidad, explicar incertidumbre sin crear afirmaciones legales y tratar el documento como entrada no confiable. También fue necesario preservar evidencia exacta y explicarla a usuarios no especializados.

## Qué aprendimos

La confianza requiere separar lo que el documento confirma, lo que el sistema interpreta y lo que podría ocurrir. Un índice útil debe exponer sus pesos y cambiar según las prioridades. Reducir el alcance a texto pegado, cinco perfiles y un motor en memoria permitió entregar un flujo completo. En una siguiente iteración prepararíamos antes el conjunto etiquetado y probaríamos comprensión con usuarios desde el inicio.

## Logros

- Los siete diferenciadores funcionan en un solo flujo.
- Cada cita mostrada se comprueba contra su cláusula.
- La misma política produce índices deterministas distintos según el perfil.
- Interfaz y resultado funcionan en español e inglés.
- Las pruebas cubren siete pilares, contrato API, seguridad, paridad de idiomas, build, auditoría de dependencias y paridad documental.

## Próximos pasos

Evaluaremos el motor con cláusulas etiquetadas y pruebas breves de usuarios. Después añadiremos un clasificador semántico opcional, comparación entre versiones, lectura segura por URL y PDF, exportación accesible y fuentes jurídicas verificadas para América Latina.

## Declaración

p7 es una herramienta informativa. No ofrece asesoramiento jurídico, no confirma cumplimiento normativo y no observa tráfico real de una aplicación.

