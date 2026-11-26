# p7 — Privacidad, descifrada.

<p align="center"><img src="docs/assets/devpost-thumbnail-p7.png" width="360" alt="p7 — Privacidad, descifrada." /></p>

**Un lector bilingüe de políticas de privacidad, respaldado por evidencia y por un LLM local.**

[English](README.md) · **Español** · [Documentación](docs/README.md)

**MVP del hackathon STEM2Connect.** Las políticas de privacidad son largas y difíciles de comparar. p7 convierte el texto pegado en una lectura estructurada y personal: prácticas declaradas, evidencia exacta, vacíos, posibles impactos y acciones concretas. Es una aplicación local funcional diseñada para revisión humana.

## Inicio rápido y demostración

Con Docker Desktop iniciado y acceso a Docker Hub, ejecuta `iniciar.cmd` en Windows o `docker compose up --build --wait` en una terminal. La primera ejecución descarga imágenes de contenedores y el modelo local y puede tardar varios minutos. Abre **http://localhost:3000**, elige idioma y perfil, pulsa **Cargar ejemplo** y luego **Analizar política**. Revisa los hallazgos citados, el mapa de datos, los vacíos, los escenarios de impacto y las acciones. El ejemplo incluido es ficticio; también puedes pegar una política propia. La aplicación no visita URL arbitrarias. Si falla la descarga de imágenes, usa la instalación local de abajo.

Detén los servicios con `docker compose down`.

## Siete diferenciadores implementados

1. **Mapa de datos:** representa el recorrido declarado.
2. **Riesgo personal:** pondera hallazgos según cinco perfiles.
3. **Detección de vacíos:** clasifica información especificada, ambigua o ausente.
4. **Evidencia verificable:** vincula cada hallazgo con una cláusula exacta y un SHA-256.
5. **Simulación de impacto:** explica consecuencias posibles sin presentarlas como hechos.
6. **Consentimiento inteligente:** genera acciones concretas y priorizadas.
7. **Privacidad para todos:** aplicación, ejemplos y resultados completos en español e inglés.

## Evaluación de IA medida

Los casos, las URL oficiales y la **salida cruda del modelo** están en [apps/api/evaluation](apps/api/evaluation). El [historial en español](docs/es/EVALUATION.md) y el [historial en inglés](docs/en/EVALUATION.md) explican cada ronda y las salvedades de las etiquetas.

| Mismo clasificador, nuevos fragmentos oficiales | Precisión | Cobertura | Casos exactos |
| --- | ---: | ---: | ---: |
| Ronda 7 · 40 casos | 73,7 % | 70,0 % | 30/40 |
| Ronda 8 · 40 casos | 96,0 % | 85,7 % | 35/40 |
| **Ronda 9 · 40 casos** | **100,0 %** | **75,0 %** | **30/40** |

La ronda 9 usó 40 fragmentos cortos seleccionados de páginas oficiales. Esas cifras **no** miden el rendimiento con documentos completos. El prompt y los pesos del modelo siguieron iguales entre rondas.

[Casos A](apps/api/evaluation/real_round9a.tsv) · [Fuentes A](apps/api/evaluation/real_round9a_sources.tsv) · [Salida cruda A](apps/api/evaluation/real_round9a_result.json) · [Casos B](apps/api/evaluation/real_round9b.tsv) · [Fuentes B](apps/api/evaluation/real_round9b_sources.tsv) · [Salida cruda B](apps/api/evaluation/real_round9b_result.json)

**Pruebas con políticas completas:** También probamos políticas oficiales completas en inglés y español. Estas mediciones por documento se acercan más al uso real que los fragmentos elegidos; los prompts y la verificación de citas cambiaron entre algunas rondas posteriores:

| Políticas oficiales completas | Precisión | Cobertura | Páginas exactas |
| --- | ---: | ---: | ---: |
| Ronda 10 · Cuentas de Mozilla, Proton | 54,5 % | 60,0 % | 1/4 |
| Ronda 11 · Apple, Spotify | 90,5 % | 95,0 % | 2/4 |
| **Ronda 12 · Google, WhatsApp** | **72,7 %** | **100,0 %** | **1/4** |
| **Ronda 13 · OpenAI, TikTok EEE** | **100,0 %** | **95,8 %** | **3/4** |
| Ronda 14 · LinkedIn, Zoom | 91,3 % | 95,5 % | 1/4 |
| Ronda 15 · Dropbox, GitHub | 85,0 % | 94,4 % | 1/4 |
| Ronda 16 · Pinterest, Slack | 81,0 % | 85,0 % | 1/4 |
| Ronda 17 · IKEA España, tras revisar el prompt | 87,5 % | 87,5 % | 1/2 |
| **Ronda 18 · Dropbox, Airbnb, tras reparar evidencia PDF** | **87,5 %** | **70,0 %** | **0/4** |

La ronda 12 usó páginas de hasta 120.506 caracteres. Encontró las 16 categorías de referencia, pero añadió seis hallazgos no respaldados; solo WhatsApp ES fue exacta. Los cuatro análisis tardaron **10 minutos y 46 segundos** en el equipo local. La etiqueta de entrenamiento de IA de Google es discutible y se analiza en el [estudio de la ronda 12](docs/es/ROUND12_CASE_STUDY.md).

La ronda 18 volvió a probar Dropbox con una etiqueta de contactos más estricta y añadió Airbnb en dos idiomas. Un fallo con saltos de línea de PDF hizo que al principio se rechazaran todos los candidatos de Airbnb. Tras reparar la verificación textual de citas, las mismas cuatro páginas produjeron **14 aciertos, 2 falsos positivos y 6 omisiones** según las etiquetas fijadas para la ronda 18; ninguna tuvo un conjunto perfecto de categorías. Una auditoría humana separada encontró **13 de 16 citas retenidas** con respaldo semántico. La pasada reparada es una **comprobación sobre la misma muestra**, no una prueba independiente. El [estudio bilingüe de la ronda 18](docs/es/ROUND18_CASE_STUDY.md) y las [salidas crudas](apps/api/evaluation/real_policies_round18_result.json) explican los errores y la etiqueta discutible de contactos de Dropbox, que fue positiva en la ronda 15. Las políticas seleccionadas, las traducciones correlacionadas y una sola persona que etiquetó no demuestran fiabilidad general. Sigue siendo necesaria la revisión humana.

## Stack

- Frontend: Next.js 16, React 19 y TypeScript.
- Backend: Go con biblioteca estándar HTTP.
- IA: Ollama con clasificación semántica estructurada mediante `qwen3:4b-instruct`.
- Pruebas: Go `testing`, `httptest` y Vitest.
- Operación: Docker Compose y GitHub Actions.

## Ejecutar en desarrollo

Instala Ollama, descarga `qwen3:4b-instruct` una sola vez y mantén Ollama activo en `http://127.0.0.1:11434`:

```powershell
ollama pull qwen3:4b-instruct
```

Terminal 1:

```powershell
cd apps/api
$env:P7_API_PORT = "8080"
$env:P7_OLLAMA_URL = "http://127.0.0.1:11434"
$env:P7_OLLAMA_MODEL = "qwen3:4b-instruct"
go run ./cmd/server
```

Terminal 2:

```powershell
cd apps/web
npm install
$env:NEXT_PUBLIC_API_URL = "http://localhost:8080"
npm run dev
```

Abre `http://localhost:3000`, carga el ejemplo y ejecuta el análisis. El selector ES/EN cambia la experiencia completa.

Si los puertos 8080 o 3000 están ocupados, configura `P7_API_PORT`, `P7_ALLOWED_ORIGINS` y `NEXT_PUBLIC_API_URL` con alternativas compatibles, y arranca Next.js con `npm run dev -- -p <puerto-web>`. Por ejemplo, en la prueba local final se usaron API `18080` y web `13000`.

## Docker

Con Docker Desktop iniciado, ejecuta `iniciar.cmd` desde la raíz del proyecto (doble clic o desde una terminal). El archivo construye la web y la API, inicia Ollama, descarga el modelo local la primera vez, espera a que todos los servicios estén sanos y abre `http://localhost:3000`.

También puedes iniciarlo manualmente:

```powershell
docker compose up --build --wait
```

El modelo se conserva en el volumen `ollama-data`, por lo que no vuelve a descargarse en cada inicio. Para detener los servicios usa `docker compose down`. Para eliminar también el modelo descargado usa `docker compose down --volumes`.

## Verificación automática

```powershell
./scripts/test-all.ps1
```

El comando ejecuta pruebas Go, pruebas del frontend, TypeScript, build de producción, auditoría de dependencias y paridad documental. GitHub Actions repite las mismas verificaciones en cada push o pull request.

## API

`POST /api/v1/analyze`

```json
{
  "title": "SocialApp Demo",
  "profile": "journalist",
  "locale": "es",
  "text": "Recopilamos tu ubicación..."
}
```

Perfiles: `general`, `student`, `journalist`, `parent` y `activist`. Idiomas: `es` y `en`.

## Seguridad y límites

- Cuerpo máximo de 1 MiB y documento máximo de 200.000 caracteres.
- El contenido se trata como dato no confiable; las instrucciones internas se ignoran.
- Cada cita se comprueba textualmente contra su cláusula, admitiendo diferencias de espacios causadas por PDF; esto no es validación semántica.
- El MVP no conserva intencionalmente políticas ni tiene cuentas.
- CORS de producción se configura mediante `P7_ALLOWED_ORIGINS`.
- El resultado es informativo, no asesoramiento jurídico.

## Uso de IA

El producto ejecutable llama al modelo local `qwen3:4b-instruct` mediante Ollama en siete pasadas centradas en categorías por lote de cláusulas. El modelo clasifica con un esquema JSON restringido y propone citas. Go rechaza categorías desconocidas y citas que no coinciden textualmente con la cláusula indicada tras normalizar espacios; luego aplica pesos por perfil y calcula el índice. Una cita coincidente aún puede sostener una interpretación incorrecta. Si Ollama no está disponible, el análisis devuelve un error de servicio en vez de sustituir silenciosamente la IA por palabras clave. ChatGPT y Codex también apoyaron investigación, alcance, arquitectura, código, diseño, pruebas y documentación. Aquí no se almacena ninguna clave API ni secreto.

## Documentación

El inglés es el idioma oficial de la presentación y cada documento tiene copia en español. Consulta el [índice bilingüe](docs/README.md), la [auditoría de preparación](docs/es/SUBMISSION_AUDIT.md), la [evaluación de IA](docs/es/EVALUATION.md), el [protocolo para cinco personas](docs/es/USER_STUDY.md) y la [decisión de despliegue](docs/es/DEPLOYMENT.md).

La [galería para la presentación](docs/assets/devpost-gallery/ready-1500x1000/) contiene ocho capturas del producto; también hay un [ZIP listo para subir](docs/assets/devpost-gallery/p7-devpost-gallery-1500x1000.zip).

## Licencia

MIT. Consulta [LICENSE](LICENSE).
