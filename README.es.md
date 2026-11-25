# p7

**Privacidad, descifrada.**

[English](README.md) · **Español**

p7 convierte políticas de privacidad en evaluaciones personales, explicables y respaldadas por evidencia. El MVP funciona sin cuentas, base de datos ni API de IA pagada.

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
- Pruebas: Go `testing`, `httptest` y Vitest.
- Operación: Docker Compose y GitHub Actions.

## Ejecutar en desarrollo

Terminal 1:

```powershell
cd apps/api
$env:P7_API_PORT = "8080"
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

ChatGPT y Codex apoyaron investigación, alcance, arquitectura, programación, diseño, pruebas y documentación. La versión actual usa un motor determinista local y no depende de créditos externos. Un futuro clasificador semántico será opcional y no controlará la puntuación final.

## Documentación

El inglés es el idioma oficial de la presentación y cada documento tiene copia en español. Consulta el [índice bilingüe](docs/README.md), la [arquitectura](docs/es/ARCHITECTURE.md), la [estrategia de evaluación](docs/es/JUDGING.md) y el [guion](docs/es/DEMO_SCRIPT.md).

## Licencia

MIT. Consulta [LICENSE](LICENSE).
