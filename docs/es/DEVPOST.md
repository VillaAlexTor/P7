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
- hallazgos vinculados a cláusulas citadas, con comprobación textual de la cita;
- escenarios posibles claramente etiquetados;
- plan de acción priorizado para el consentimiento;
- experiencia completa en español e inglés.

El documento normalizado recibe un hash SHA-256 para identificar la versión exacta analizada.

## Cómo lo construimos

La aplicación usa Next.js, React y TypeScript. Una API Go segmenta políticas y envía lotes de cláusulas a Ollama local, donde `qwen3:4b-instruct` devuelve candidatos semánticos restringidos por esquema. Go comprueba la presencia textual de cada cita en su cláusula, evalúa controles, aplica pesos transparentes y deriva el mapa y las recomendaciones. Esta comprobación no demuestra que la interpretación sea correcta. El MVP no usa cuentas, base de datos, API pagada ni clave.

## Cómo usamos IA

ChatGPT y Codex apoyaron investigación de competencia, definición del problema, alcance, arquitectura, programación, diseño, pruebas y documentación. Esta asistencia está divulgada en el repositorio y la presentación.

La versión ejecutable llama a `qwen3:4b-instruct` mediante Ollama local y solicita candidatos de clasificación semántica estructurados. Sigue siendo reproducible sin créditos API pagados ni secretos. Go comprueba las citas contra sus cláusulas, admite diferencias de espacios causadas por PDF, rechaza categorías desconocidas, deduplica y aplica pesos deterministas. Una auditoría separada de citas hecha por una persona muestra que una coincidencia textual aún puede sostener una interpretación errónea.

## Desafíos

El desafío principal fue evitar que el producto se convierta en otro resumidor legal. Tuvimos que conectar siete funciones sin perder usabilidad, explicar incertidumbre sin crear afirmaciones legales y tratar el documento como entrada no confiable. También fue necesario preservar evidencia exacta y explicarla a usuarios no especializados.

## Qué aprendimos

Llegué a p7 con más confianza en backend y ciberseguridad que en IA explicable o diseño de producto. Mi primer impulso fue concentrarme en lograr que el análisis funcionara. La primera evaluación me corrigió: el modelo interpretó con mucha seguridad las doce cláusulas negativas como si esas prácticas sí ocurrieran. Fue un error incómodo, pero útil. Revisé el prompt, añadí controles negativos, mantuve la verificación de evidencia exacta y dejé de confundir una respuesta convincente con una respuesta correcta.

Las rondas posteriores con fuentes reales fueron igual de importantes. Los resultados variaron mucho entre fragmentos seleccionados y políticas completas. En la ronda 18, los saltos de línea del PDF hicieron que el verificador literal rechazara todos los candidatos de Airbnb. Al admitir diferencias de espacios, la cobertura en esas mismas cuatro páginas localizadas subió de 30 % a 70 %, pero ninguna tuvo un conjunto perfecto y 3 de 16 citas retenidas carecían de respaldo semántico. Aprendí que un resultado fuerte no garantiza fiabilidad, que una demo funcional no es un estudio de exactitud y que una cita coincidente puede sostener una interpretación equivocada. Si comenzara de nuevo, probaría mucho antes documentos completos y comprensión con usuarios.

## Logros

- Los siete diferenciadores funcionan en un solo flujo.
- Cada cita mostrada se comprueba contra su cláusula.
- La misma política produce índices deterministas distintos según el perfil.
- Interfaz y resultado funcionan en español e inglés.
- Las pruebas cubren siete pilares, contrato API, seguridad, paridad de idiomas, build, auditoría de dependencias y paridad documental.
- Una [evaluación sintética reproducible de 40 cláusulas](EVALUATION.md) informa 27 positivos verdaderos, cero falsos positivos y un falso negativo tras ajustar el prompt con el mismo conjunto.
- Nueve rondas con políticas completas conservan fuentes y salidas crudas del modelo. La última pasada revisada sobre Dropbox y Airbnb encontró 14 de 20 categorías de referencia, añadió 2 según la etiqueta discutida de contactos de Dropbox y omitió 6: 87,5 % de precisión, 70,0 % de cobertura y 0 de 4 páginas localizadas exactas. Fue una comprobación del arreglo sobre la misma muestra, no una nueva prueba independiente. Una auditoría de una persona encontró 13 de 16 citas retenidas pertinentes; el [estudio de la ronda 18](ROUND18_CASE_STUDY.md) documenta las salvedades.

## Próximos pasos

Pediremos a otra persona que revise las etiquetas, evaluaremos más políticas completas de proveedores independientes y realizaremos el [protocolo preparado para cinco personas](USER_STUDY.md). Después añadiremos comparación entre versiones, lectura segura por URL y PDF, exportación accesible y fuentes jurídicas verificadas para América Latina.

## Declaración de trabajo previo

No existía código de la aplicación antes de este hackathon. El concepto, la implementación, la evaluación, la documentación y los materiales de presentación de p7 se crearon durante el evento. El proyecto utiliza frameworks y bibliotecas de código abierto acreditados, el entorno de Ollama y el modelo `qwen3:4b-instruct`.

## Declaración

p7 es una herramienta informativa. No ofrece asesoramiento jurídico, no confirma cumplimiento normativo y no observa tráfico real de una aplicación.

