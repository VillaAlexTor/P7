# Política de seguridad

## Versión soportada

La rama `main` actual es la versión soportada para el hackathon.

## Reportar una vulnerabilidad

No abras un issue público con políticas reales, datos personales, credenciales o detalles de explotación. Contacta de forma privada a Villa Cyber Defense y proporciona una reproducción mínima y sintética.

## Límites de seguridad

- El texto se trata como dato no confiable y nunca como instrucción ejecutable.
- El cuerpo está limitado a 1 MiB y la política a 200.000 caracteres.
- El MVP no tiene cuentas ni conserva intencionalmente las políticas analizadas.
- Cada cita mostrada se verifica contra la cláusula fuente indicada.
- Los orígenes CORS deben configurarse con `P7_ALLOWED_ORIGINS` en producción.
- El MVP no obtiene URLs arbitrarias, evitando una superficie SSRF.

Cualquier integración futura con LLM deberá usar salida estructurada, aislar el contenido, rechazar citas no verificables y conservar la puntuación determinista en Go.
