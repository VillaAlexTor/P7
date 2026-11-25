# Ronda 14: políticas completas de LinkedIn y Zoom

## Alcance y referencia congelada

El 25 de noviembre de 2026 p7 analizó las páginas oficiales completas de [LinkedIn EN](https://www.linkedin.com/legal/privacy-policy), [LinkedIn ES](https://es.linkedin.com/legal/privacy-policy), [Zoom EN](https://www.zoom.com/en/trust/privacy/privacy-statement/) y [Zoom ES](https://www.zoom.com/es/trust/privacy/privacy-statement/). Son **dos políticas subyacentes en dos idiomas**, no cuatro proveedores independientes. Los textos extraídos tienen 42.150–64.369 caracteres. LinkedIn indica vigencia del 3 de noviembre de 2025; Zoom, actualización del 17 de septiembre de 2026. Los textos completos permanecen en `.cache/real_policies_round14`, excluida de Git; el [manifiesto publicado](../../apps/api/evaluation/real_policies_round14.json) registra hashes SHA-256 exactos, longitudes, líneas de apoyo, etiquetas y ocho preguntas. Se confirmó como `80f8536` **antes de ejecutar el modelo**.

Una persona fijó las siete categorías. `contacts` exige acceso o carga explícita de la **agenda del dispositivo**: LinkedIn menciona carga de agendas, mientras Zoom habla de listas de contactos e integración con otra aplicación, sin afirmar acceso a la agenda del dispositivo. LinkedIn dice que puede entrenar modelos de IA con datos personales. Zoom niega entrenar sus modelos o los de terceros con contenido del cliente similar a comunicaciones; procesar contenido para ofrecer funciones de IA es otra práctica. Contamos como `biometric` las huellas de voz y geometría facial opcionales de Zoom junto a la finalidad declarada de autenticación; los efectos faciales locales no cuentan. Esa interpretación biométrica merece revisión independiente. Analizamos cada página una vez con el mismo `qwen3:4b-instruct` local, temperatura cero, perfil general, prompts y código Go. El [informe compacto de máquina](../../apps/api/evaluation/real_policies_round14_result.json) conserva categorías, IDs de cláusulas, hashes de citas, candidatos, puntuaciones y tiempos sin republicar las políticas.

## Resultados por documento y categoría

| Página completa | Categorías de referencia | Categorías de p7 | Aciertos | Falsos positivos | Omisiones | Tiempo |
| --- | --- | --- | ---: | ---: | ---: | ---: |
| LinkedIn EN | ubicación, contactos, terceros, publicidad, entrenamiento de IA, retención | todas salvo **contactos** | 5 | 0 | 1 | 68 s |
| LinkedIn ES | las mismas seis | exactamente las seis | 6 | 0 | 0 | 100 s |
| Zoom EN | ubicación, terceros, publicidad, biometría, retención | las cinco **más contactos** | 5 | 1 | 0 | 114 s |
| Zoom ES | las mismas cinco | las cinco **más contactos** | 5 | 1 | 0 | 149 s |
| **Total** | | | **21** | **2** | **1** | **433 s** |

La precisión por documento-categoría fue **21/23 = 91,3 %**, la cobertura **21/22 = 95,5 %** y **1/4** páginas tuvieron el conjunto exacto. Las cuatro páginas producen 28 decisiones de categoría, incluidas cuatro ausencias correctas. El modelo propuso 103 hallazgos, conservó 23 coincidencias literales y contó 80 rechazos; ese contador mezcla duplicados con evidencia inválida. Las puntuaciones de riesgo fueron 92, 100, 100 y 100, en ese orden; no son probabilidades calibradas. Más de siete minutos por lote es una limitación práctica.

## Ocho preguntas fijadas antes de inferir

Son 16 comprobaciones por idioma, superpuestas a las 28 decisiones anteriores. «No establecido» significa que el aviso no afirma la práctica definida.

| Pregunta | Referencia | p7 EN | p7 ES | Dificultad |
| --- | --- | --- | --- | --- |
| L1. ¿LinkedIn usa agendas cargadas para sugerir contactos? | Sí | **No** | Sí | La carga explícita está en la línea 173; sincronizar calendarios o tener datos de contacto es distinto. |
| L2. ¿LinkedIn puede entrenar modelos de IA con datos personales? | Sí | Sí | Sí | La línea 161 afirma tanto el origen de datos como la finalidad de entrenamiento. |
| L3. ¿LinkedIn muestra anuncios personalizados? | Sí | Sí | Sí | La línea 207 lo afirma. |
| L4. ¿LinkedIn afirma identificación biométrica? | No establecido | No | No | Inferir edad o género no es identificación biométrica. |
| Z1. ¿Zoom entrena IA con reuniones, chats o audio del cliente? | No establecido | No | No | La línea 63 lo niega; el procesamiento para funciones de IA en la línea 73 no es entrenamiento. |
| Z2. ¿Zoom procesa huellas de voz o geometría facial opcionales para funciones que incluyen autenticación? | Sí | Sí | Sí | Las líneas 71 y 335 describen rasgos únicos y autenticación; los efectos solo locales de la línea 69 son distintos. |
| Z3. ¿Las listas de contactos de Zoom demuestran acceso a la agenda del dispositivo? | No establecido | **Sí** | **Sí** | La línea 39 habla de listas e integración con otra aplicación, no de agenda del dispositivo. |
| Z4. ¿Zoom declara anuncios dirigidos según el uso de su sitio? | Sí | Sí | Sí | La línea 97 describe publicidad dirigida por uso del sitio. |

p7 acertó **13/16** comprobaciones. La omisión de LinkedIn EN y los dos falsos positivos de contactos en Zoom son las discrepancias. Las versiones lingüísticas están correlacionadas; no son 16 observaciones independientes.

## Auditoría separada de las citas

La puntuación anterior da crédito cuando la práctica aparece **en alguna parte de la política completa**. No demuestra que la cláusula concreta citada por p7 la sostenga. Revisamos las 23 ubicaciones conservadas. Al menos **nueve ubicaciones no bastan por sí solas** para justificar el hallazgo, incluidos los dos falsos positivos de Zoom:

| Página y hallazgo | Cláusula citada | Problema |
| --- | --- | --- |
| LinkedIn ES `contacts` | `clause-043` | Sincronización de calendarios/email y datos de contacto; la carga de la agenda aparece después, en la línea 173. La categoría existe, pero la cita es débil. |
| LinkedIn ES `third-parties` | `clause-044` | Describe datos **recibidos de** clientes/socios, no datos **compartidos con** un destinatario externo. El envío sí se afirma en la línea 345. |
| Zoom EN/ES `contacts` | `clause-022` en ambos | Listas de contactos e integración con otra aplicación no prueban acceso a la agenda del dispositivo. También son falsos positivos de categoría. |
| Zoom EN/ES `third-parties` | `clause-049` / `clause-050` | Describe información **recibida de** socios. El acceso de proveedores se afirma en la línea 93. |
| Zoom EN/ES `advertising` | `clause-049` / `clause-050` | Describe datos obtenidos de socios publicitarios, incluidos clics; no afirma por sí sola que los datos personales seleccionen anuncios. La segmentación aparece en la línea 97. |
| Zoom EN `biometric` | `clause-076` | Describe TrueDepth procesando dimensiones faciales/manuales **en el dispositivo** para efectos. El procesamiento opcional por Zoom aparece en las líneas 71 y 335. |

Hay **siete citas débiles en categorías correctas y dos en los falsos positivos de contactos**. Las demás ubicaciones parecían pertinentes en esta revisión de una persona; no es una anotación independiente ni otra métrica calibrada. El verificador Go confirma que la cita sea una subcadena literal de la cláusula, pero no que demuestre semánticamente la categoría. Por eso una cita exacta puede inducir a error. El informe de máquina guarda hashes de citas; las cláusulas se pueden comprobar en la caché o en una descarga nueva cuyos hashes coincidan.

## Interpretación y trabajo siguiente

En las rondas diez a catorce con políticas completas, los totales descriptivos son **85 aciertos, 15 falsos positivos y 7 omisiones** en 20 páginas localizadas de diez políticas subyacentes: **85,0 % de precisión y 92,4 % de cobertura** por categoría. El 91,3 % de precisión de esta ronda **no significa** 91,3 % de calidad de evidencia. Las políticas se seleccionaron, las traducciones están correlacionadas, una persona asignó las etiquetas y auditó las citas, y no cambiaron prompts ni modelo. La etiqueta biométrica de Zoom requiere especialmente otra revisión. No se ha demostrado fiabilidad general del 90 % ni progreso por entrenamiento.

La siguiente corrección debe exigir que la evidencia citada afirme la acción que define la categoría: acceso a agenda para `contacts`, transferencia hacia afuera para `third-parties`, segmentación con datos personales para `advertising` y procesamiento de un identificador biométrico por el proveedor para `biometric`. Conviene separar motivos de rechazo y reevaluar en **políticas nuevas**, conservando esta ronda como prueba cerrada. Los 433 segundos también requieren mejorar el rendimiento.

Para reproducir mientras las páginas no cambien, ejecuta `python scripts/prepare-real-policies-round14.py --fetch` desde la raíz y verifica los hashes impresos. Con Ollama y `qwen3:4b-instruct`, ejecuta desde `apps/api`:

```powershell
go run ./cmd/evaluatepolicies -manifest evaluation/real_policies_round14.json -text-dir ../../.cache/real_policies_round14
```
