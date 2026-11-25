# Ronda 16: políticas vigentes de Pinterest y Slack

## Alcance y método congelado

El 25 de noviembre de 2026 p7 analizó las páginas oficiales completas de [Pinterest EN](https://policy.pinterest.com/en/privacy-policy), [Pinterest ES](https://policy.pinterest.com/es/privacy-policy), [Slack EN](https://slack.com/trust/privacy/privacy-policy) y [Slack ES](https://slack.com/intl/es-es/trust/privacy/privacy-policy). Son **dos políticas en dos idiomas**, no cuatro políticas independientes. Slack indica vigencia del 5 de julio de 2023. Los textos extraídos tienen 39.142–58.404 caracteres. Los originales permanecen en `.cache/real_policies_round16`, excluida de Git; el [manifiesto](../../apps/api/evaluation/real_policies_round16.json) fija URL, hashes SHA-256, longitudes, líneas de referencia, etiquetas y ocho preguntas. Se confirmó como `f563e0c` **antes de inferir**. Una persona asignó las etiquetas. El mismo `qwen3:4b-instruct` local, temperatura cero, perfil general, prompts y código Go analizó cada página una vez. El [informe de máquina](../../apps/api/evaluation/real_policies_round16_result.json) registra predicciones, IDs de cláusulas, hashes de citas, puntuaciones, recuentos y tiempos sin republicar las políticas.

Pinterest dice que **continúa procesando contactos sincronizados anteriormente** (línea 37). Según la definición vigente de p7, conservar una lista anterior no afirma un acceso nuevo a la agenda; por eso congelamos `contacts` como negativo. Pinterest afirma entrenar modelos de IA con Pines (línea 83). Slack describe importar una agenda opcionalmente (línea 73), anuncios dirigidos fuera de Slack (línea 291 EN / 295 ES) y conservar Otra Información tras desactivar la cuenta (línea 225 EN / 229 ES). El aviso de Slack no afirma entrenamiento de IA.

## Resultados con la referencia congelada

| Página completa | Categorías de referencia | Categorías de p7 | Aciertos | Falsos positivos | Omisiones | Tiempo |
| --- | --- | --- | ---: | ---: | ---: | ---: |
| Pinterest EN | ubicación, terceros, publicidad, entrenamiento de IA, retención | primeras cuatro + **contactos**; **falta retención** | 4 | 1 | 1 | 87 s |
| Pinterest ES | las mismas cinco | esas cinco + **contactos y biometría** | 5 | 2 | 0 | 122 s |
| Slack EN | ubicación, contactos, terceros, publicidad, retención | exactamente las cinco | 5 | 0 | 0 | 67 s |
| Slack ES | las mismas cinco | ubicación, contactos, terceros + **biometría**; **faltan publicidad y retención** | 3 | 1 | 2 | 85 s |
| **Total** | | | **17** | **4** | **3** | **363 s** |

La precisión documento-categoría con etiquetas congeladas fue **17/21 = 81,0 %**, la cobertura **17/20 = 85,0 %** y **1/4** páginas tuvieron el conjunto exacto. En 28 decisiones hubo 17 aciertos, 4 falsos positivos, 3 omisiones y 4 ausencias correctas. El modelo propuso 138 candidatos, conservó 21 coincidencias literales y contó 117 rechazos; el contador mezcla duplicados y evidencia inválida. Las puntuaciones fueron 100, 100, 92 y 96 en ese orden; no son probabilidades calibradas.

### Corrección importante de etiqueta posterior a la inferencia

Los negativos biométricos congelados para Pinterest fueron **un error en mi anotación**. Consideré su modelo de tipos de cuerpo, que no identifica personas, pero pasé por alto otra frase en las líneas **95 y 199 de ambos idiomas**: Pinterest puede comparar un rostro con la foto de un documento oficial para verificar la edad. Según el criterio de coincidencia de identidad usado para Spotify en la ronda once, ambas páginas de Pinterest deberían ser positivas en `biometric`. El modelo español halló esa práctica en `clause-122`; el inglés la omitió. **No alteramos** el manifiesto, la salida cruda ni el resultado original. El [registro separado de sensibilidad posterior a la inferencia](../../apps/api/evaluation/real_policies_round16_adjudication.json) cambia solo esas dos etiquetas: **18 aciertos, 3 falsos positivos, 4 omisiones; 85,7 % de precisión y 81,8 % de cobertura**. Siguen 1/4 páginas exactas. Esta interpretación se hizo tras ver las predicciones y necesita revisión independiente; **no reemplaza el resultado preconfirmado**.

## Ocho preguntas fijadas antes de inferir

Producen 16 comprobaciones por idioma que se superponen con las 28 decisiones; no son otra muestra. «No establecido» significa que el aviso vigente no afirma la práctica definida.

| Pregunta | Referencia congelada | p7 EN | p7 ES | Razón |
| --- | --- | --- | --- | --- |
| P1. ¿Pinterest afirma una **nueva** importación de la agenda? | No establecido | **Sí** | **Sí** | La línea 37 habla de uso continuo de contactos sincronizados antes. |
| P2. ¿Pinterest entrena IA con Pines? | Sí | Sí | Sí | La línea 83 afirma entrenamiento de modelos. |
| P3. ¿Pinterest elige anuncios según intereses inferidos? | Sí | Sí | Sí | La línea 339 describe esa selección. |
| P4. ¿La tecnología de tipos de cuerpo demuestra identificación biométrica? | No establecido | No | **Sí** | La pregunta congelada nombró esa tecnología, pero el documento también compara rostros con documentos en las líneas 95/199. Con la etiqueta *documental* corregida, EN omite y ES encuentra biometría; la pregunta sigue siendo más estrecha. |
| S1. ¿Slack recoge una agenda tras elección del usuario? | Sí | Sí | Sí | La línea 73 menciona importar una agenda. |
| S2. ¿Este aviso de Slack afirma entrenar IA con contenido del espacio? | No establecido | No | No | No contiene una declaración afirmativa. |
| S3. ¿Slack comparte datos para anuncios dirigidos fuera de Slack? | Sí | Sí | **No** | Línea 291 EN / 295 ES. |
| S4. ¿Puede Slack conservar Otra Información después de desactivar la cuenta? | Sí | Sí | **No** | Línea 225 EN / 229 ES. |

Con el mapa original de preguntas, p7 acertó **11/16**. P4 demuestra que detectar una categoría en el documento no equivale a responder la redacción precisa de la pregunta: el hallazgo biométrico español provino de otro pasaje. La respuesta congelada P4 sigue siendo válida **solo para la tecnología de tipos de cuerpo**; la categoría biométrica del documento completo requiere la corrección indicada. No reescribimos la pregunta después de inferir.

## Auditoría de cláusulas citadas

Revisamos las 21 ubicaciones conservadas. **Al menos ocho no justifican por sí solas** el hallazgo asignado:

| Página | Hallazgo y cláusula | Problema |
| --- | --- | --- |
| Pinterest EN/ES | `third-parties`, `clause-029` / `clause-038` | Habla de datos **recibidos de** afiliados/anunciantes, no de envío hacia afuera. Las transferencias se describen en otra sección (línea 541). |
| Pinterest EN/ES | `contacts`, `clause-019` / `clause-025` | Describe uso continuo de contactos sincronizados **antes**, no importación actual. Son falsos positivos con ambas referencias. |
| Slack EN/ES | `third-parties`, `clause-043` en ambos | Dice que servicios externos **envían** información a Slack. El envío desde Slack aparece después (línea 201 EN / 205 ES). |
| Slack EN | `advertising`, `clause-102` | Los correos promocionales no son segmentación publicitaria con datos personales. La categoría aparece en la línea 291. |
| Slack ES | `biometric`, `clause-076` | Menciona el lector de huellas **del dispositivo**, no procesamiento biométrico por Slack. |

Son cinco citas débiles en categorías correctas y tres asociadas a falsos positivos según las etiquetas congeladas: los dos contactos de Pinterest y la biometría de Slack ES. El cuarto falso positivo congelado, biometría de Pinterest ES, cita el contraste de rostro con documento y resulta pertinente bajo la referencia corregida. Por tanto, calidad de citas y precisión por categoría son cosas distintas. Las otras ubicaciones parecían adecuadas en esta revisión de una persona. p7 comprueba coincidencia literal de la cita, no suficiencia semántica.

## Interpretación y siguiente trabajo

Los errores recurrentes son confundir datos **entrantes** con transferencia hacia afuera, tratar contactos antiguos como acceso actual y confundir biometría solo local con procesamiento por el proveedor. Slack ES también omitió dos prácticas explícitas. La corrección posterior muestra que la calidad de las etiquetas de referencia limita esta evaluación. Proveedores seleccionados, traducciones correlacionadas, una sola persona anotadora y una etiqueta discutida impiden afirmar fiabilidad general. No cambiaron modelo ni prompt; las diferencias de puntuación no demuestran entrenamiento.

El siguiente paso es pedir una segunda revisión de la etiqueta biométrica de Pinterest, separar motivos de rechazo y exigir que la cláusula citada afirme la acción que define la categoría. Cualquier corrección debe probarse con **políticas nuevas**, conservando esta ronda cerrada. Para reproducir mientras las páginas no cambien, ejecuta `python scripts/prepare-real-policies-round16.py --fetch` desde la raíz, compara hashes y, con Ollama y `qwen3:4b-instruct`, ejecuta desde `apps/api`:

```powershell
go run ./cmd/evaluatepolicies -manifest evaluation/real_policies_round16.json -text-dir ../../.cache/real_policies_round16
```
