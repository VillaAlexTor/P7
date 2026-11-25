# Ronda 17: política de privacidad completa de IKEA España

## Fuentes y referencia fijada

El 25 de noviembre de 2026, p7 analizó las páginas oficiales de IKEA España en [inglés](https://www.ikea.com/es/en/customer-service/privacy-policy/) y [español](https://www.ikea.com/es/es/customer-service/privacy-policy/), fechadas el 4 de agosto de 2026. Son **dos traducciones de una política**, no dos organizaciones independientes. Los textos completos extraídos tienen 43.120 y 47.845 caracteres. La [preparación de fuentes](../../scripts/prepare-real-policies-round17.py) guarda HTML y texto en `.cache/real_policies_round17`, ignorada por Git; el [manifiesto fijado](../../apps/api/evaluation/real_policies_round17.json) registra URL, hash, longitud, siete etiquetas por página y siete preguntas bilingües. Se confirmó como `b671ffb` **antes de la inferencia**. Una persona asignó las etiquetas de referencia.

La referencia marca `location`, `third-parties`, `advertising` y `retention` como positivos en ambos idiomas; `contacts`, `ai-training` y `biometric` son negativos. La política describe ubicación opcional en la app (línea extraída 111), cesión a proveedores externos (EN 251 / ES 253) y mensajes comerciales adaptados a compras y navegación (EN 167 / ES 169). Describe análisis de llamadas con IA (EN 137 / ES 139), pero no afirma entrenamiento del modelo. Cuenta visitantes con cámaras (EN 165 / ES 167), sin afirmar identificación de personas. Conserva datos de eventos por un plazo ligado a obligaciones legales sin fecha definida (EN 329 / ES 331). El usuario introduce destinatarios de listas de regalos; no se importa la agenda del teléfono.

## Dos pasadas transparentes

La primera pasada usó el prompt revisado a partir de la ronda 16. Se conserva su [salida inicial](../../apps/api/evaluation/real_policies_round17_result_initial.json). Tras revisar los errores, precisé la regla de contactos y de caducidad y ejecuté una **segunda pasada sobre los mismos documentos**. La [segunda salida](../../apps/api/evaluation/real_policies_round17_result.json) es una *comprobación del ajuste sobre la muestra*, **no una prueba independiente nueva**. Ambas usaron Ollama local con `qwen3:4b-instruct`, temperatura cero, perfil general, documento completo y un hallazgo retenido por categoría y página. No se entrenaron pesos del modelo.

| Pasada | IKEA EN | IKEA ES | VP | FP | FN | Precisión | Cobertura |
| --- | --- | --- | ---: | ---: | ---: | ---: | ---: |
| Inicial | cuatro positivos + falso `contacts` | cuatro positivos + falso `contacts` | 8 | 2 | 0 | 80,0% | 100% |
| Revisada | omite `retention` y mantiene falso `contacts` | cuatro categorías exactas | 7 | 1 | 1 | 87,5% | 87,5% |

Ambas pasadas aciertan **12 de 14** comprobaciones de categoría y pregunta localizadas. La revisada acierta por completo una de las dos páginas. Estos porcentajes miden **presencia de categorías**, sobre traducciones correlacionadas de una sola política. No representan fiabilidad en producción ni confianza calibrada.

## Auditoría de evidencias

Revisé todas las ubicaciones de las citas retenidas en ambos informes. Solo **5 de 10** citas iniciales y **5 de 8** revisadas sostienen por sí mismas la categoría. Ambas pasadas detectan correctamente publicidad en el documento inglés, pero citan `clause-069`, una oferta genérica para socios. El pasaje de mensajes adaptados está en la línea extraída 167. Ambas detectan conservación en español, pero citan `clause-041`, una tarjeta que caduca; la retención sin plazo fijo ligada a obligaciones legales está en la línea 331. La cita inicial de conservación inglesa falla por el mismo motivo. Los hallazgos iniciales de contactos citan pasajes de atención al cliente llamados «Contact»; el hallazgo inglés revisado cita «contact points» de las tiendas en `clause-222`. Ninguno demuestra acceso a la agenda del dispositivo. La omisión inglesa de conservación en la segunda pasada muestra que el ajuste redujo un falso positivo a costa de un falso negativo. Las citas de ubicación y cesión a proveedores sí son pertinentes en ambos idiomas; también la de publicidad española revisada.

El campo automático `evidenceVerified` de p7 significa **coincidencia literal de la cita**, no respaldo semántico. Esta auditoría es una lectura humana separada. El modelo aún puede escoger una cláusula temprana débil aunque exista otra más clara. El ajuste de prompt no basta para seleccionar evidencias de forma fiable; hace falta una etapa de validación de citas y otra prueba independiente con políticas nuevas. Esta muestra no justifica afirmar 90% de fiabilidad.

## Reproducción

El manifiesto registra los SHA-256 de los textos cacheados. Si cambian las páginas oficiales, `python scripts/prepare-real-policies-round17.py --fetch` obtiene **otra versión de fuente**, no reproduce la pasada congelada. Con el HTML original en caché, ejecuta `python scripts/prepare-real-policies-round17.py`, comprueba los hashes, inicia Ollama con `qwen3:4b-instruct` y, desde `apps/api`, ejecuta:

```powershell
go run ./cmd/evaluatepolicies -manifest evaluation/real_policies_round17.json -text-dir ../../.cache/real_policies_round17
```

El comando utiliza el prompt actual y puede diferir de las dos pasadas registradas si cambian el prompt o el comportamiento del modelo. Los dos JSON versionados contienen las salidas crudas de este estudio.
