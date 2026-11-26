# Guion de demostración

Este es el esquema de planificación. El video terminado dura 3 minutos y 48,6 segundos; el límite del hackathon es de 6 minutos. La narración final está en `video/p7-demo/SCRIPT.md`.

Úsalo como guía natural en vez de leerlo palabra por palabra. Si el modelo local tarda al iniciar el análisis, haz una pausa y recórtala durante la edición.

**Producción final:** El video muestra capturas reales de la aplicación sin cámara del autor. La narración inglesa se sintetizó localmente con Kokoro y se mezcló con música; los subtítulos corregidos se entregan como archivo SRT separado. Declara la voz sintética en la postulación Devpost.

## 0:00–0:35 — Problema e impacto

**En pantalla:** Inicio de p7 en español; después cambia a inglés y mantenlo para el análisis del ejemplo y la comparación de 97 a 80 indicada abajo.

**Di:**

“Aceptamos políticas de privacidad todo el tiempo y casi nunca las leemos completas. Incluso cuando una herramienta las resume, puede no decirme qué importa para mi situación, qué deja sin aclarar el documento o qué puedo hacer después. Construí p7 para convertir esa decisión difícil en evidencia, posibles impactos y acciones concretas, con una experiencia completa en español e inglés.”

## 0:35–1:10 — Entrada y perfil

**En pantalla:** Carga el ejemplo, selecciona **Periodista** e inicia el análisis.

**Di:**

“El usuario pega una política en lugar de entregar a p7 acceso a una cuenta o al historial de navegación. Voy a usar el ejemplo ficticio incluido y elegir el perfil Periodista. Este perfil da más peso a ubicación, contactos y transferencia a terceros. La política se procesa localmente y la aplicación no la almacena de forma intencional.”

## 1:10–1:55 — Resultado personalizado y mapa

**En pantalla:** Muestra índice, perfil, cantidad de cláusulas, identidad SHA-256 y mapa de datos.

**Di:**

“Este índice no lo inventa directamente el modelo de lenguaje. El modelo identifica prácticas candidatas; el backend en Go verifica la evidencia y aplica pesos transparentes según el perfil. El valor SHA-256 identifica el documento normalizado exacto usado para este informe. El mapa representa el recorrido declarado en la política, desde la persona hacia el servicio y posibles destinos. No afirma observar tráfico real de red.”

## 1:55–2:45 — Evidencia y vacíos de información

**En pantalla:** Abre dos hallazgos y después la sección de vacíos.

**Di:**

“Cada hallazgo confirmado vuelve a una cláusula citada. En este ejemplo, p7 identifica ubicación precisa y posible uso del contenido enviado para entrenar IA. El backend rechaza citas que no coinciden con la cláusula indicada, admitiendo saltos de línea de PDF. Una cita coincidente todavía requiere interpretación humana. p7 también separa información especificada, ambigua y ausente; que algo no aparezca invita a preguntar, no prueba una infracción legal.”

## 2:45–3:30 — Posibles impactos y acciones

**En pantalla:** Muestra un escenario y marca dos acciones como completadas.

**Di:**

“La siguiente sección traduce los hallazgos en escenarios posibles sin presentarlos como hechos. Por ejemplo, la ubicación precisa podría revelar rutinas, y los datos compartidos con socios podrían salir del servicio original. El plan convierte esas preocupaciones en pasos prácticos, como restringir la ubicación y revisar la configuración de transferencia.”

## 3:30–3:55 — Personalización

**En pantalla:** Vuelve al mismo ejemplo, cambia de **Periodista** a **Uso general**, pulsa **Analizar política** otra vez y señala el cambio de índice. En la demo inglesa verificada, cambió de 97 a 80.

**Di:**

“Al volver a analizar la misma política con otro perfil, cambian las prioridades. En esta ejecución, el índice pasa de 97 a 80. Un periodista, un estudiante o un padre pueden afrontar consecuencias distintas. Los pesos del perfil son reproducibles y el modelo no los modifica.”

## 3:55–4:35 — Uso significativo y declarado de IA

**En pantalla:** Mantén visible el informe y señala los metadatos del modelo.

**Di:**

“El producto ejecutable usa `qwen3:4b-instruct` mediante Ollama local. Realiza pasadas enfocadas de clasificación semántica y devuelve candidatos estructurados. Go comprueba las categorías admitidas y que el texto citado aparezca en la política, elimina duplicados y calcula el índice. Que la cita coincida todavía requiere una revisión humana de su significado. ChatGPT y Codex también me ayudaron con investigación, alcance, arquitectura, código, diseño, pruebas y documentación. Revisé y probé ese trabajo, y declaro ambos tipos de asistencia de IA en el repositorio y en esta presentación.”

## 4:35–5:10 — Aprendizaje y una limitación honesta

**En pantalla:** Muestra brevemente la tabla de evaluación del README o de la galería.

**Di:**

“Mi experiencia previa me daba más confianza en backend y ciberseguridad que en IA explicable. La primera evaluación mostró que el modelo había interpretado doce negaciones como prácticas reales. Más tarde, una prueba con PDF reveló un fallo que rechazaba todos los candidatos de Airbnb. Al repararlo, la cobertura subió de 30 % a 70 % sobre las mismas cuatro páginas localizadas, pero ninguna obtuvo un conjunto perfecto y tres citas aún carecían de respaldo semántico. Aprendí que una demo funcional demuestra funcionamiento, no exactitud general.”

## 5:10–5:25 — Trabajo previo y próximos pasos

**En pantalla:** Termina en el plan de acciones o en el título del proyecto.

**Di:**

“No existía código de la aplicación antes de este hackathon; p7 se creó durante el evento usando herramientas y modelos de código abierto acreditados. Después añadiría revisión independiente de etiquetas, pruebas de comprensión con usuarios, comparación de versiones y fuentes verificadas para América Latina. p7 no reemplaza asesoramiento jurídico: ofrece una primera lectura más clara antes de aceptar.”

## Checklist de grabación

- Graba la pantalla de la aplicación; mostrar tu cara o usar webcam es opcional.
- Declara la narración sintetizada localmente con Kokoro en la postulación.
- Mantén el video final en 6:00 o menos.
- Muestra un análisis real completo, no solo capturas o diapositivas.
- Recorta en la edición la espera de carga o inferencia del modelo.
- Añade subtítulos en inglés y revisa nombres como Ollama, SHA-256 y `qwen3:4b-instruct`.
- Verifica sin iniciar sesión que los enlaces del repositorio y del video funcionen.
- No afirmes que se realizaron pruebas con usuarios hasta completar realmente las sesiones.

## Cobertura de criterios

- Problema e impacto: 0:00–0:35.
- Funcionalidad: 0:35–3:55.
- Uso de IA: 3:55–4:35.
- Aprendizaje y crecimiento: 4:35–5:10.
- Presentación, trabajo previo y próximos pasos: 5:10–5:25.
