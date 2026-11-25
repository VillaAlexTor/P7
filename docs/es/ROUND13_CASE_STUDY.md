# Ronda 13: prácticas opcionales en políticas completas

## Método y fuentes

El 25 de noviembre de 2026 evaluamos p7 sin modificarlo con cuatro páginas oficiales y completas: [OpenAI EN](https://openai.com/policies/privacy-policy/), [OpenAI ES](https://openai.com/es-419/policies/privacy-policy/), [TikTok EEE EN](https://www.tiktok.com/legal/page/eea/privacy-policy/en) y [TikTok EEE ES](https://www.tiktok.com/legal/page/eea/privacy-policy/es). Son **dos políticas subyacentes**, cada una en dos idiomas. El texto descargado de OpenAI indicaba actualización del 30 de julio de 2026; el de TikTok EEE, del 30 de noviembre de 2025. El [manifiesto congelado](../../apps/api/evaluation/real_policies_round13.json) identifica las extracciones exactas por SHA-256 y longitud. Las páginas pueden cambiar o variar según la región.

Confirmamos el manifiesto en el commit `329672d` **antes de la inferencia**, con etiquetas de siete categorías, líneas de referencia y ocho preguntas de control. Una persona asignó las etiquetas. Los textos completos extraídos (22.484–53.600 caracteres) permanecen en `.cache/real_policies_round13`, excluida de Git; no republicamos las políticas ajenas. El [informe compacto de máquina](../../apps/api/evaluation/real_policies_round13_result.json) conserva categorías finales, IDs de cláusula, hashes de citas, puntuaciones, tiempos y cantidad de candidatos. Usamos el mismo `qwen3:4b-instruct` local, temperatura cero, perfil general, prompts y código Go de las rondas diez a doce. Cada página se analizó una vez, sin cambios posteriores a las predicciones.

La referencia cuenta una práctica opcional como presente cuando la política afirma que ocurre tras una elección o consentimiento. `contacts` exige **acceso a la agenda del dispositivo**, no solo tener el correo o teléfono de alguien. `biometric` exige un identificador usado para reconocer a una persona; detectar rostros para efectos o mencionar un servicio de verificación sin especificar biometría no basta.

## Resultados

| Página completa | Categorías de referencia | Categorías de p7 | Aciertos | Falsos positivos | Omisiones | Tiempo |
| --- | --- | --- | ---: | ---: | ---: | ---: |
| OpenAI EN | ubicación, contactos, terceros, publicidad, entrenamiento de IA, retención | exactamente las seis | 6 | 0 | 0 | 40 s |
| OpenAI ES | las mismas seis | cinco; **faltan contactos** | 5 | 0 | 1 | 40 s |
| TikTok EEE EN | las mismas seis | exactamente las seis | 6 | 0 | 0 | 85 s |
| TikTok EEE ES | las mismas seis | exactamente las seis | 6 | 0 | 0 | 100 s |
| **Total** | | | **23** | **0** | **1** | **267 s** |

La precisión fue **23/23 = 100%**, la cobertura **23/24 = 95,8%** y **3/4** páginas tuvieron el conjunto exacto. Las 28 decisiones página-categoría incluyen cuatro ausencias correctas de `biometric`. **No hay ejemplos biométricos positivos en esta ronda**, por lo que no medimos su cobertura. El modelo propuso 101 candidatos, conservó 23 hallazgos con citas literales y contó 78 rechazos. Ese contador mezcla duplicados y pruebas inválidas; no permite saber cuántas citas eran incorrectas. Las puntuaciones de riesgo fueron 100, 96, 100 y 100 en ese orden; **no son probabilidades calibradas**.

## Ocho preguntas fijadas antes de inferir

Cada pregunta se compara con la página inglesa y la española: 16 comprobaciones localizadas que **se superponen** con las 28 decisiones página-categoría; no son otra muestra independiente. «No establecido» significa que este aviso no afirma la práctica definida, no que sea imposible.

| Pregunta | Referencia | p7 EN | p7 ES | Prueba o distinción |
| --- | --- | --- | --- | --- |
| O1. ¿La conexión opcional de OpenAI sube la agenda del dispositivo? | Sí | Sí | **No** | OpenAI EN/ES, línea 33, afirma la carga de la agenda después de conectar contactos. |
| O2. ¿Puede usarse contenido del usuario para entrenar los modelos de ChatGPT? | Sí | Sí | Sí | OpenAI, línea 79, describe el entrenamiento con contenido del usuario y sus controles. |
| O3. ¿Afirma la política anuncios personalizados para algunos usuarios? | Sí | Sí | Sí | OpenAI, línea 65, describe personalización para Free/Go. |
| O4. ¿Afirma reconocimiento biométrico de identidad por OpenAI? | No establecido | No | No | La línea 35 menciona verificar identidad/edad, sin especificar un identificador biométrico. |
| T1. ¿La sincronización opcional de TikTok recoge la agenda telefónica? | Sí | Sí | Sí | TikTok, línea 19, menciona expresamente la agenda. |
| T2. ¿Describe TikTok entrenamiento de modelos con información de la plataforma? | Sí | Sí | Sí | TikTok, línea 75, describe entrenamiento de modelos de aprendizaje automático. |
| T3. ¿Muestra TikTok anuncios basados en intereses cuando se permite? | Sí | Sí | Sí | TikTok, línea 299, describe publicidad por intereses con consentimiento. |
| T4. ¿Detectar rostros para efectos demuestra identificación biométrica? | No establecido | No | No | TikTok, línea 37, describe detección facial/corporal para efectos, sin reconocer identidades. |

p7 acertó **15/16** comprobaciones. La única omisión fue O1 en español. La política española es explícita: «si eliges conectar los contactos de tu dispositivo, subimos información de las libretas de direcciones de tu dispositivo» (línea 33). La ejecución inglesa halló `contacts` en `clause-015`; la española no produjo ese hallazgo pese al texto equivalente. El informe no conserva los candidatos rechazados, así que no podemos distinguir si el modelo omitió el pasaje o p7 rechazó una cita propuesta.

## Interpretación y próximos pasos

El resultado de estas páginas supera al de la ronda doce, pero **no demuestra que el modelo aprendiera**: pesos, prompts y código quedaron iguales. En las rondas diez a trece, los totales descriptivos son **64 aciertos, 13 falsos positivos y 6 omisiones** en 16 páginas localizadas de ocho políticas subyacentes: **83,1% de precisión y 91,4% de cobertura**. Las páginas fueron elegidas, las traducciones están correlacionadas y una sola persona etiquetó todo. Una categoría correcta por página tampoco prueba que su cita sea la mejor evidencia. No corresponde afirmar confiabilidad general del 90%.

El siguiente trabajo es distinguir motivos de rechazo de candidatos e investigar la cobertura de contactos en español. Otra persona debería revisar las etiquetas, especialmente los casos discutibles de entrenamiento y biometría. Una evaluación posterior debe usar proveedores nuevos e incluir reconocimiento biométrico afirmativo, además de políticas sin contactos opcionales, personalización publicitaria ni entrenamiento. Los 267 segundos del lote también muestran una limitación para demostraciones en vivo.

Para reproducir la extracción mientras las páginas no cambien, ejecuta `python scripts/prepare-real-policies-round13.py --fetch` desde la raíz y compara los hashes con el manifiesto. Con Ollama y `qwen3:4b-instruct`, ejecuta desde `apps/api`:

```powershell
go run ./cmd/evaluatepolicies -manifest evaluation/real_policies_round13.json -text-dir ../../.cache/real_policies_round13
```
