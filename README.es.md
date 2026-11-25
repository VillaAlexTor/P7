# p7 — Privacidad, descifrada.

<p align="center"><img src="docs/assets/devpost-thumbnail-p7.png" width="360" alt="p7 — Privacidad, descifrada." /></p>

**Un lector bilingüe de políticas de privacidad, respaldado por evidencia y por un LLM local.**

[English](README.md) · **Español** · [Documentación](docs/README.md)

**MVP del hackathon STEM2Connect.** Las políticas de privacidad son largas y difíciles de comparar. p7 convierte el texto pegado en una lectura estructurada y personal: prácticas declaradas, evidencia exacta, vacíos, posibles impactos y acciones concretas. Es una aplicación local funcional diseñada para revisión humana.

## Inicio rápido y demostración

Con Docker Desktop iniciado, ejecuta `iniciar.cmd` en Windows o `docker compose up --build --wait` en una terminal. La primera ejecución descarga el modelo local y puede tardar varios minutos. Abre **http://localhost:3000**, elige idioma y perfil, pulsa **Cargar ejemplo** y luego **Analizar política**. Revisa los hallazgos citados, el mapa de datos, los vacíos, los escenarios de impacto y las acciones. El ejemplo incluido es ficticio; también puedes pegar una política propia. La aplicación no visita URL arbitrarias.

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

La última ronda usó **40 URL oficiales nuevas** de Google, Spotify, Microsoft, Mozilla y Apple, divididas por igual entre idiomas. Hubo 33 etiquetas acertadas, ningún falso positivo y 11 omisiones. Las mayores deficiencias fueron transferencia a terceros y retención. Son fragmentos cortos seleccionados, algunos traducidos de un mismo aviso, etiquetados por una persona. **No** miden el rendimiento en documentos completos; el 100 % observado no garantiza fiabilidad general. El prompt y los pesos del modelo siguieron iguales entre rondas. p7 ayuda a revisar, pero una persona debe comprobar hallazgos y posibles omisiones.

[Casos A](apps/api/evaluation/real_round9a.tsv) · [Fuentes A](apps/api/evaluation/real_round9a_sources.tsv) · [Salida cruda A](apps/api/evaluation/real_round9a_result.json) · [Casos B](apps/api/evaluation/real_round9b.tsv) · [Fuentes B](apps/api/evaluation/real_round9b_sources.tsv) · [Salida cruda B](apps/api/evaluation/real_round9b_result.json)

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
- Cada cita se valida contra la cláusula indicada.
- El MVP no conserva intencionalmente políticas ni tiene cuentas.
- CORS de producción se configura mediante `P7_ALLOWED_ORIGINS`.
- El resultado es informativo, no asesoramiento jurídico.

## Uso de IA

El producto ejecutable llama al modelo local `qwen3:4b-instruct` mediante Ollama en cinco pasadas centradas en categorías. El modelo clasifica las cláusulas con un esquema JSON restringido y propone referencias exactas. Go rechaza categorías desconocidas y cualquier cita que no esté contenida en la cláusula indicada; luego aplica pesos por perfil y calcula el índice. Si Ollama no está disponible, el análisis devuelve un error de servicio en vez de sustituir silenciosamente la IA por palabras clave. ChatGPT y Codex también apoyaron investigación, alcance, arquitectura, código, diseño, pruebas y documentación. Aquí no se almacena ninguna clave API ni secreto.

## Documentación

El inglés es el idioma oficial de la presentación y cada documento tiene copia en español. Consulta el [índice bilingüe](docs/README.md), la [evaluación de IA](docs/es/EVALUATION.md), el [protocolo para cinco personas](docs/es/USER_STUDY.md) y la [decisión de despliegue](docs/es/DEPLOYMENT.md).

La [galería para la presentación](docs/assets/devpost-gallery/ready-1500x1000/) contiene ocho capturas del producto; también hay un [ZIP listo para subir](docs/assets/devpost-gallery/p7-devpost-gallery-1500x1000.zip).

## Licencia

MIT. Consulta [LICENSE](LICENSE).
