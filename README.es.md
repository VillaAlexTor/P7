# p7

**Privacidad, descifrada.**

[English](README.md) · **Español**

p7 convierte políticas de privacidad en evaluaciones personales, explicables y respaldadas por evidencia. El MVP usa un LLM local mediante Ollama, sin cuentas, base de datos, claves API ni créditos pagados.

## Siete diferenciadores implementados

1. **Mapa de datos:** representa el recorrido declarado.
2. **Riesgo personal:** pondera hallazgos según cinco perfiles.
3. **Detección de vacíos:** clasifica información especificada, ambigua o ausente.
4. **Evidencia verificable:** vincula cada hallazgo con una cláusula exacta y un SHA-256.
5. **Simulación de impacto:** explica consecuencias posibles sin presentarlas como hechos.
6. **Consentimiento inteligente:** genera acciones concretas y priorizadas.
7. **Privacidad para todos:** aplicación, ejemplos y resultados completos en español e inglés.

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

```powershell
docker compose up --build
```

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

El producto ejecutable llama al modelo local `qwen3:4b-instruct` mediante Ollama. El modelo clasifica las cláusulas con un esquema JSON restringido y propone referencias exactas. Go rechaza categorías desconocidas y cualquier cita que no esté contenida en la cláusula indicada; luego aplica pesos por perfil y calcula el índice. ChatGPT y Codex también apoyaron investigación, alcance, arquitectura, código, diseño, pruebas y documentación. Aquí no se almacena ninguna clave API ni secreto.

## Documentación

El inglés es el idioma oficial de la presentación y cada documento tiene copia en español. Consulta el [índice bilingüe](docs/README.md), la [evaluación de IA](docs/es/EVALUATION.md), el [protocolo para cinco personas](docs/es/USER_STUDY.md) y la [decisión de despliegue](docs/es/DEPLOYMENT.md).

## Licencia

MIT. Consulta [LICENSE](LICENSE).
