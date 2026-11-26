# Auditoría de preparación para la entrega — 26 de noviembre de 2026

## Requisitos oficiales

La [página del hackathon en Devpost](https://stem2connect-online-hackathon.devpost.com/) exige un proyecto funcional, un **repositorio público de GitHub** con código, instalación y declaración de IA, un **video de hasta seis minutos** en YouTube o Vimeo público u oculto, y una descripción que cubra problema, solución, uso de IA, aprendizaje y próximos pasos. Hay que indicar una categoría principal, incluir a todo el equipo y declarar trabajo anterior al hackathon si lo hubiera. Los participantes deben ser estudiantes de al menos 13 años; el equipo puede tener hasta seis personas. El plazo mostrado es el **29 de noviembre de 2026 a las 23:59 EST** (30 de noviembre a las 00:59 en Bolivia). El video entregado cuenta como presentación pregrabada; no hace falta un segundo. El Demo Day en vivo es opcional.

## Verificado en este espacio de trabajo

- Se confirmó que [VillaAlexTor/P7](https://github.com/VillaAlexTor/P7) es público y su rama predeterminada es `main`. El último commit local y los cambios de esta revisión requieren un nuevo resultado remoto de CI después del push.
- Pasó la suite completa `scripts/test-all.ps1` el 26 de noviembre: pruebas Go, seis pruebas del frontend, TypeScript, build de producción, auditoría de dependencias de producción con **cero vulnerabilidades**, paridad bilingüe de documentación del proyecto y metadatos de fuentes de evaluación.
- Una petición nueva al código Go actual utilizó Ollama `qwen3:4b-instruct` y devolvió cuatro hallazgos con cuatro citas verificadas textualmente, siete comprobaciones de información, un mapa de cinco nodos y cuatro conexiones, cuatro escenarios de impacto y cinco acciones. El ejemplo inglés obtuvo **97** para Periodista; la prueba anterior de Uso general obtuvo **80**. La salud de la API informó el modelo correcto. La prueba de humo previa de la web devolvió HTTP 200 y permitió el origen CORS configurado.
- La documentación versionada tiene parejas en inglés y español. Un escaneo de nombres versionados y patrones habituales de credenciales no encontró secretos publicados; `.cache`, el entorno virtual del video y sus medios generados están ignorados.
- El borrador Devpost, la arquitectura, el README y el guion ahora describen las siete pasadas por categoría y distinguen coincidencia textual de respaldo semántico.
- Se revisó el video final local: H.264/AAC, 1920×1080 a 30 fps y **3:48,6** de duración. Muestra la acción Analizar, un informe completo con cuatro hallazgos y un segundo análisis que cambia Periodista **97** a Uso general **80**. La diapositiva de evaluación final indica **87,5 % de precisión, 70 % de cobertura y 0/4 páginas exactas**. El audio mide **−16,2 LUFS** integrados y pico real de **−1,5 dBFS**; se transcribió automáticamente la voz de la mezcla final. Los subtítulos ingleses se entregan como archivo SRT separado. Este entorno no permitió una escucha humana del audio.

## Pendiente antes de entregar

- Subir el video revisado a YouTube o Vimeo como público u oculto, añadir el SRT inglés corregido y verificar el enlace sin iniciar sesión. Durante esta auditoría no se encontró una URL pública del video en el repositorio.
- Completar el formulario Devpost, elegir **Community Impact** como categoría principal, agregar los enlaces del repositorio y video, comprobar que cada integrante se haya unido y revisar la reflexión personal y la declaración de trabajo previo para que sean fieles a los hechos. El registro, la condición de estudiante, el equipo y el envío Devpost no se pueden verificar desde el repositorio.
- **No** se completó un arranque limpio de Docker Compose en esta máquina: Docker Hub agotó el tiempo del handshake TLS al descargar `ollama/ollama`. Volvió a pasar `docker compose config -q` y existe la etiqueta fijada de Docker Hub. La prueba de humo anterior sí comprobó Ollama instalado en el equipo con la API Go y Next.js de producción. Los puertos `8080` y `3000` pueden estar ocupados por otras aplicaciones; aquella prueba usó API `18080` y web `13000`.

## Límite de exactitud que conviene explicar

La ronda 18 reparada sobre cuatro páginas completas localizadas registró **14 VP, 2 FP y 6 FN**, es decir, 87,5 % de precisión y 70,0 % de cobertura según las etiquetas fijadas, con **0/4** páginas exactas. Una auditoría de una persona consideró pertinentes **13/16** citas retenidas. La etiqueta de contactos de Dropbox contradice la ronda 15 y la segunda pasada reutilizó las páginas de la primera. Las cifras no demuestran fiabilidad general ni 90 % de exactitud. Consulta el [estudio bilingüe](ROUND18_CASE_STUDY.md) y su versión inglesa.
