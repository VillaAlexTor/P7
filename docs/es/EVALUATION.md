# Evaluación de IA — conjunto de cláusulas sintéticas

## Método

Enviamos individualmente cuarenta cláusulas sintéticas etiquetadas a mano al modelo local `qwen3:4b-instruct` mediante el mismo clasificador estructurado de Ollama que usa p7. Veinte están en inglés y veinte en español. Cada idioma tiene catorce prácticas afirmativas de siete categorías y seis controles negativos, incluidas negaciones y un plazo concreto de eliminación. El evaluador solo acepta una predicción si la categoría es admitida y la cita propuesta aparece literalmente en la cláusula indicada.

Ejecuta desde `apps/api`, con Ollama activo y el modelo instalado:

```powershell
go run ./cmd/evaluate -data evaluation/clauses.tsv
```

El conjunto y el evaluador están en el repositorio. La salida incluye cada caso, conteos por categoría y tiempo transcurrido. No usamos datos de personas ni políticas reales de empresas.

## Resultado observado, 25 de noviembre de 2026

| Versión | Positivos acertados | Falsos positivos | Positivos omitidos | Precisión | Cobertura |
| --- | ---: | ---: | ---: | ---: | ---: |
| Prompt inicial | 28 | 12 | 0 | 70,0 % | 100,0 % |
| Instrucción explícita sobre negaciones | 22 | 0 | 6 | 100,0 % | 78,6 % |
| Prompt actual con ejemplos afirmativos | 27 | 0 | 1 | 100,0 % | 96,4 % |

El prompt actual reconoció los doce controles negativos y 27 de 28 cláusulas afirmativas. Omitió E14: “We keep your information for as long as necessary for business purposes.” Las 27 predicciones aceptadas tenían una cita literal válida; en esta corrida no se rechazó ninguna cita propuesta. Las 40 llamadas individuales tardaron unos 13 segundos en la computadora de desarrollo.

Estos resultados son **sintéticos y ajustados sobre el mismo conjunto**. Usamos los fallos de este conjunto para revisar el prompt; los porcentajes no estiman de forma independiente la exactitud en políticas reales. Documentos largos, varias prácticas en una cláusula, formulaciones indirectas, excepciones y texto adversarial requieren otra evaluación reservada. El clasificador todavía puede omitir prácticas o interpretar mal su alcance, y una cita literal correcta no prueba por sí sola que el hallazgo sea semánticamente verdadero.

## Segunda evaluación: 40 cláusulas nuevas

Congelamos el prompt actual y creamos `evaluation/heldout.tsv` después de las revisiones descritas arriba. Este segundo conjunto sintético contiene 20 cláusulas en inglés y 20 en español: 28 casos afirmativos, dos con dos etiquetas, y 12 controles negativos. Examina consentimiento condicionado, redacción indirecta, procesadores externos, copias de seguridad y prácticas combinadas. No ajustamos el prompt con este conjunto antes de registrar el resultado.

Ejecuta desde `apps/api`:

```powershell
go run ./cmd/evaluate -data evaluation/heldout.tsv
```

| Conjunto | Etiquetas acertadas | Falsos positivos | Etiquetas omitidas | Precisión | Cobertura |
| --- | ---: | ---: | ---: | ---: | ---: |
| Primero tras el ajuste (usado para ajustar) | 27 | 0 | 1 | 100,0 % | 96,4 % |
| Segundo (reservado del ajuste) | 19 | 0 | 11 | 100,0 % | 63,3 % |

Los 12 controles negativos quedaron sin hallazgos. No se rechazó ninguna cita propuesta. Las 40 llamadas tardaron unos 11 segundos en la computadora de desarrollo. Resultados por categoría del segundo conjunto:

| Categoría | Aciertos | Omisiones |
| --- | ---: | ---: |
| Ubicación | 4 | 0 |
| Contactos | 1 | 3 |
| Terceros | 1 | 5 |
| Publicidad | 4 | 0 |
| Entrenamiento de IA | 4 | 0 |
| Biometría | 4 | 0 |
| Retención | 1 | 3 |

Por ejemplo, se omitieron H02 (importación condicionada de contactos), H03 (contratista de prevención de fraude) y P07 (almacenamiento sin fecha de borrado). H04 y P04 recibieron la etiqueta de publicidad, pero omitieron la divulgación simultánea a terceros. Son vacíos relevantes para una herramienta que ayuda a interpretar condiciones de privacidad. El cero de falsos positivos solo corresponde a estas 40 cláusulas sintéticas, y una sola persona asignó las etiquetas. Es una comprobación independiente del prompt con casos sintéticos nuevos, no una estimación del rendimiento con políticas reales o documentos completos.

## Tercera evaluación: otras 40 cláusulas

Mantuvimos el mismo modelo `qwen3:4b-instruct` y el mismo prompt. El tercer conjunto sintético, `evaluation/heldout_2.tsv`, se escribió antes de ejecutar el clasificador y contiene otras 20 cláusulas en inglés y 20 en español. Cada idioma tiene 14 casos afirmativos, incluidos dos casos con dos etiquetas, y seis controles negativos. Examina acceso condicionado, subcontratistas, prácticas combinadas, situaciones futuras hipotéticas, funciones retiradas y plazos fijos de eliminación. Una sola persona asignó las etiquetas.

Ejecuta desde `apps/api`:

```powershell
go run ./cmd/evaluate -data evaluation/heldout_2.tsv
```

| Conjunto | Etiquetas acertadas | Falsos positivos | Etiquetas omitidas | Precisión | Cobertura |
| --- | ---: | ---: | ---: | ---: | ---: |
| Primero, tras ajustar con el mismo conjunto | 27 | 0 | 1 | 100,0 % | 96,4 % |
| Segundo, conjunto sintético nuevo | 19 | 0 | 11 | 100,0 % | 63,3 % |
| Tercero, conjunto sintético nuevo | 25 | 0 | 7 | 100,0 % | 78,1 % |

Los 12 controles negativos quedaron sin hallazgos y no se rechazó ninguna cita. Las 40 llamadas tardaron unos 18 segundos en la computadora de desarrollo. En inglés hubo 12 etiquetas acertadas y cuatro omitidas; en español, 13 acertadas y tres omitidas.

| Categoría | Aciertos | Omisiones |
| --- | ---: | ---: |
| Ubicación | 3 | 1 |
| Contactos | 4 | 0 |
| Terceros | 5 | 1 |
| Publicidad | 4 | 2 |
| Entrenamiento de IA | 4 | 0 |
| Biometría | 4 | 0 |
| Retención | 1 | 3 |

T07, T14 y S07 omitieron prácticas de retención. T04 y S04 perdieron una de sus dos etiquetas esperadas, igual que T08 y S08. El tercer conjunto confirma que la retención expresada de forma amplia y las cláusulas con dos prácticas simultáneas siguen siendo puntos débiles. Estas cifras describen solo esta pequeña muestra sintética; no validan el análisis de políticas reales ni documentos completos.

## Mejora iterativa y punto de parada

Mejoramos las **instrucciones y el flujo de inferencia** alrededor del modelo local `qwen3:4b-instruct`; no actualizamos sus pesos ni lo ajustamos con entrenamiento adicional. El clasificador de una sola pasada seguía omitiendo prácticas en un conjunto nuevo de 40 cláusulas, así que dividimos la clasificación en pasadas enfocadas. Después separamos retención en su propia pasada. Comprobamos cada versión con las cláusulas de trabajo antes de redactar un conjunto final intacto.

| Corrida | Conjunto | Etiquetas acertadas | Falsos positivos | Etiquetas omitidas | Cobertura |
| --- | --- | ---: | ---: | ---: | ---: |
| Una pasada | `heldout_4.tsv`, primera corrida | 22 | 1 | 10 | 68,8 % |
| Cuatro pasadas | `heldout_4.tsv`, repetición de trabajo | 27 | 1 | 5 | 84,4 % |
| Cinco pasadas | `heldout_4.tsv`, repetición de trabajo | 30 | 0 | 2 | 93,8 % |
| **Cinco pasadas, comprobación final intacta** | **`heldout_5.tsv`** | **27** | **0** | **5** | **84,4 %** |

El conjunto final se escribió antes de su primera ejecución: 20 cláusulas en inglés y 20 en español, 28 casos afirmativos con cuatro casos de dos etiquetas, y 12 controles negativos. El modelo acertó al menos un positivo en cada una de las siete categorías y no produjo hallazgos para los controles negativos. Ningún hallazgo aceptado careció de cita literal; una cita propuesta fue rechazada. Ejecútalo desde `apps/api` con:

```powershell
go run ./cmd/evaluate -data evaluation/heldout_5.tsv
```

Las omisiones finales fueron X03 y X23 (pasarelas de pago), X06 y X26 (ritmo de tecleo) y X27 (retención abierta). Las cláusulas sobre pasarelas no dicen expresamente que sean externas, así que esas dos etiquetas de referencia pueden suscitar desacuerdo. Mantuvimos las etiquetas originales y el resultado informado. Las salidas del modelo también variaron entre corridas repetidas pese a usar temperatura cero. Cinco llamadas al modelo por lote de cláusulas tardan más que el flujo original de una sola pasada.

Dejamos de cambiar el clasificador cuando el conjunto intacto superó el umbral elegido de **80 % de cobertura** sin falsos positivos en esta corrida. Esto evidencia progreso con cláusulas sintéticas; no significa un 80 % de confianza estadística en su exactitud con políticas reales ni que esté listo para decisiones sin revisión.

## Evaluación con fragmentos de políticas reales

El 25 de noviembre de 2026 reunimos 40 fragmentos breves y literales de políticas de privacidad y explicaciones oficiales publicadas: 20 en inglés y 20 en español. El texto exacto, las etiquetas y las URL de origen están en `evaluation/real_dev.tsv`, `evaluation/real_dev_sources.tsv`, `evaluation/real_holdout.tsv` y `evaluation/real_holdout_sources.tsv`. Las URL permiten contrastar cada fragmento con su página publicada. Una sola persona asignó las etiquetas antes de ejecutar el modelo. Los dos conjuntos usan páginas de origen diferentes; no cambiamos el prompt, el modelo ni el código de clasificación después de preparar cada conjunto. Corregimos una grafía británica, acortamos un fragmento negativo y sustituimos una fuente del conjunto de trabajo que coincidía con el reservado; después repetimos ese conjunto. Ejecutamos una sola vez el conjunto reservado y no lo modificamos.

| Conjunto de fuentes reales | Etiquetas acertadas | Falsos positivos | Etiquetas omitidas | Precisión | Cobertura |
| --- | ---: | ---: | ---: | ---: | ---: |
| Trabajo, 20 fragmentos | 18 | 1 | 1 | 94,7 % | 94,7 % |
| **Reservado intacto, 20 fragmentos** | **17** | **0** | **2** | **100,0 %** | **89,5 %** |
| Ambos conjuntos, 40 fragmentos | 35 | 1 | 3 | 97,2 % | 92,1 % |

En el reservado acertó 8/9 etiquetas inglesas y 9/10 españolas; los tres controles negativos quedaron sin hallazgos. Las omisiones fueron RF06, el uso de la IP por Discord para determinar una región, y RF16, la sincronización periódica de números telefónicos del dispositivo por Discord. RD07 recibió una etiqueta de terceros sin fundamento en el conjunto de trabajo, mientras que RD12 omitió retención amplia. El clasificador congelado de cinco pasadas **superó el umbral de 80 % de cobertura con fragmentos reales publicados sin otro ajuste**, así que detuvimos las iteraciones.

Para reproducirlo desde `apps/api`, con Ollama activo:

```powershell
go run ./cmd/evaluate -data evaluation/real_dev.tsv
go run ./cmd/evaluate -data evaluation/real_holdout.tsv
```

Esto respalda la conclusión de que el proyecto va por buen camino. La muestra es pequeña, seleccionada y etiquetada por una sola persona; varios fragmentos proceden de la misma empresa. Evaluamos fragmentos breves por separado, no políticas completas, y la salida del modelo local puede variar entre corridas aun con temperatura cero. Los porcentajes **no** garantizan exactitud para cualquier política real. Antes de afirmar fiabilidad hacen falta otra persona que revise las etiquetas, más empresas, pruebas con documentos completos y medidas de latencia.

## Siguiente evaluación

Antes de afirmar utilidad en producción, ampliar la muestra de fuentes reales, pedir a otra persona que revise las etiquetas discutidas, medir la latencia con documentos completos e informar precisión y cobertura por idioma. Conservar los seis conjuntos sintéticos y los dos de fuentes reales como pruebas de regresión; cualquier cambio posterior del prompt necesita otro conjunto intacto.

## Seguimiento del 25 de noviembre: objetivo del 90 %

Reforzamos las instrucciones del clasificador local `qwen3:4b-instruct` y separamos ubicación, contactos y biometría en consultas individuales. **Esto cambia el prompt y el flujo de inferencia, no los pesos del modelo.** Preparamos otros dos conjuntos bilingües de 20 fragmentos breves cada uno, de páginas oficiales publicadas, con una URL de origen por fragmento. Sus textos, etiquetas y URL están en `evaluation/real_holdout_90*.tsv`. Ambos conjuntos ya son datos de trabajo porque sus errores guiaron cambios posteriores.

| Corrida | Etiquetas acertadas | Falsos positivos | Etiquetas omitidas | Precisión | Cobertura |
| --- | ---: | ---: | ---: | ---: | ---: |
| Primer conjunto nuevo, primera corrida | 18 | 3 | 0 | 85,7 % | 100,0 % |
| Segundo conjunto nuevo, primera corrida | 13 | 5 | 3 | 72,2 % | 81,3 % |
| Segundo conjunto, tras siete consultas enfocadas | 14 | 1 | 2 | 93,3 % | 87,5 % |

Después de la primera corrida acortamos el fragmento R9S01 del primer conjunto para cumplir el límite de 25 palabras citadas. Su etiqueta no cambió; la puntuación inicial no se puede reproducir byte por byte con el archivo confirmado. La primera puntuación del segundo conjunto se registró antes de ajustarlo. La última fila es una repetición sobre datos de trabajo, **no una prueba reservada independiente**. Sus errores restantes son una etiqueta biométrica omitida en una lista que depende del contexto (RFE09), una retención amplia omitida (RFS08) y una etiqueta publicitaria sin sustento cuando solo se menciona compartir datos con socios publicitarios (RFS09). Otra persona debería revisar la etiqueta RFE09 con el contexto original.

Los resultados observados **no** respaldan afirmar 90 % tanto de precisión como de cobertura en fragmentos reales intactos. Cargamos el modelo local más grande `qwen3-coder:latest` para comparar, pero la ejecución dependiente de CPU no completó los 20 fragmentos tras varios minutos; detuvimos esa comparación y no cambiamos el modelo del producto. La siguiente comprobación válida necesita otro conjunto real, intacto y etiquetado de forma independiente, seguido de pruebas con documentos completos y latencia. Estos fragmentos seleccionados no dan confianza estadística ni garantizan exactitud en políticas arbitrarias.

Ejecuta los conjuntos adicionales desde `apps/api`, con Ollama activo:

```powershell
go run ./cmd/evaluate -data evaluation/real_holdout_90.tsv
go run ./cmd/evaluate -data evaluation/real_holdout_90_final.tsv
```

## 25 de noviembre: cuarta ronda real e historial del prompt

Seguimos mejorando **las instrucciones y la inferencia** del modelo local `qwen3:4b-instruct`; no entrenamos ni cambiamos sus pesos. Cada conjunto tiene 10 fragmentos breves en inglés y 10 en español, etiquetas asignadas antes de la inferencia y URL de páginas oficiales por separado. El comprobador de fuentes verifica 120 fragmentos, 120 URL distintas y un máximo de 25 palabras citadas por URL en los seis conjuntos reales.

| Corrida cronológica | Etiquetas acertadas | Falsos positivos | Omisiones | Precisión | Cobertura | Estado |
| --- | ---: | ---: | ---: | ---: | ---: | --- |
| `real_round3.tsv`, primera corrida con `cf8884a` | 13 | 1 | 3 | 92,9 % | 81,3 % | Primera corrida intacta |
| `real_round3.tsv`, tras revisar el prompt en `c90c9e9` | 16 | 0 | 0 | 100,0 % | 100,0 % | Repetición de trabajo |
| `real_holdout_90_final.tsv`, misma revisión | 16 | 1 | 0 | 94,1 % | 100,0 % | Repetición de trabajo |
| **`real_round4.tsv`, primera corrida con `c90c9e9` congelado** | **18** | **2** | **2** | **90,0 %** | **90,0 %** | **Primera corrida intacta** |

La primera corrida de la tercera ronda omitió R3E01 y R3E03 (retención sin plazo) y R3E10 (uso biométrico con consentimiento); añadió publicidad a R3S10, que solo describía una transferencia a socios publicitarios. Revisamos el prompt para reconocer retención por finalidad y biometría opcional, y exigir una afirmación explícita de segmentación publicitaria. Los resultados posteriores de la tercera ronda y del conjunto anterior son de trabajo porque esas cláusulas ya habían guiado las instrucciones. En el conjunto anterior aún queda una etiqueta de publicidad sin sustento en RFS09.

Confirmamos el texto y las fuentes de la cuarta ronda en Git **antes** de su primera inferencia y ejecutamos una sola vez el clasificador congelado. R4E01 omitió la importación de contactos porque rechazó la cita propuesta; R4E04 omitió la segmentación de anuncios por terceros. R4E05 llamó erróneamente tercero a Apple Ads y R4S06 dedujo segmentación publicitaria de una transferencia a socios publicitarios. En inglés hubo 9 aciertos, 1 falso positivo y 2 omisiones; en español, 9 aciertos, 1 falso positivo y ninguna omisión. Los tres controles negativos quedaron sin hallazgos. No cambiamos las etiquetas ni el prompt después de ver estos resultados.

La precisión y cobertura observadas del 90 % alcanzan el punto de parada elegido **solo en esta muestra seleccionada de 20 fragmentos**. No miden confianza estadística, rendimiento con políticas completas ni fiabilidad general. Una sola persona etiquetó los fragmentos y algunas empresas aparecen en varios conjuntos. Aún hacen falta otra persona para revisar las etiquetas, más empresas independientes, evaluación de documentos completos y medidas de latencia. Para repetir la última corrida desde `apps/api`, con Ollama activo:

```powershell
go run ./cmd/evaluate -data evaluation/real_round4.tsv
```

## 25 de noviembre: quinta ronda con fuentes reales

Congelamos el mismo prompt y flujo de inferencia de `c90c9e9`, y confirmamos `real_round5.tsv` y `real_round5_sources.tsv` antes de la primera corrida del modelo. El conjunto nuevo contiene 10 fragmentos ingleses y 10 españoles de 20 páginas oficiales que no aparecían en los conjuntos reales anteriores. Tiene 18 etiquetas afirmativas, incluido un caso de dos etiquetas, y tres controles negativos. El comprobador de fuentes cubre ahora 140 fragmentos y 140 URL distintas, con un máximo de 25 palabras citadas por URL.

| Corrida | Etiquetas acertadas | Falsos positivos | Omisiones | Precisión | Cobertura |
| --- | ---: | ---: | ---: | ---: | ---: |
| `real_round5.tsv`, prompt congelado, primera corrida | 18 | 0 | 0 | 100,0 % | 100,0 % |

Los tres controles negativos quedaron sin hallazgos. **No cambiamos el prompt ni el modelo** después de la corrida: la muestra nueva no mostró un error que justificara otra revisión específica. Esto evalúa de nuevo el modelo local y sus instrucciones; no entrena los pesos del modelo. La primera corrida de la cuarta ronda sigue siendo 90,0 % de precisión y 90,0 % de cobertura; el resultado nuevo no borra sus cuatro errores.

La quinta muestra es pequeña, seleccionada y etiquetada por una sola persona. Su resultado observado perfecto no estima la exactitud general ni demuestra rendimiento con documentos completos. El trabajo siguiente debe dar prioridad a etiquetas revisadas por otra persona, más empresas, pruebas de documentos completos y medidas de repetibilidad y latencia. Para repetirla desde `apps/api`, con Ollama activo:

```powershell
go run ./cmd/evaluate -data evaluation/real_round5.tsv
```

## 25 de noviembre: cuarenta pruebas reales adicionales

Preparamos **40 fragmentos literales nuevos** de 40 URL oficiales no utilizadas antes: 20 ingleses y 20 españoles. Los dividimos de antemano en un conjunto de trabajo de 20 casos y otro reservado de 20, ambos equilibrados por idioma. Confirmamos el texto, las etiquetas de referencia y las URL en Git antes de cada primera inferencia. El comprobador cubre ahora 180 casos y 180 enlaces de origen, verifica que las 40 URL nuevas sean distintas y limita cada cita a 25 palabras. Un conjunto de trabajo antiguo usa la misma URL para dos casos. No cambiamos las etiquetas después de ver las respuestas.

De nuevo cambiamos **solo el prompt**, no los pesos del modelo local `qwen3:4b-instruct`. El conjunto de trabajo reveló dos avisos biométricos falsos cuando la huella o el rostro permanecían en el dispositivo del usuario, y dos omisiones de retención sin plazo cuando los chats seguían guardados hasta que el usuario los borrara. La primera revisión del prompt corrigió tres de los cuatro errores; una segunda aclaración sobre biometría exclusiva del dispositivo corrigió el último en la repetición de trabajo. Congelamos la revisión `b821d56` antes de ejecutar el reservado.

| Corrida | Etiquetas acertadas | Falsos positivos | Omisiones | Precisión | Cobertura | Interpretación |
| --- | ---: | ---: | ---: | ---: | ---: | --- |
| `real_round6_dev.tsv`, primera corrida con el prompt anterior | 12 | 2 | 2 | 85,7 % | 85,7 % | Base de trabajo |
| `real_round6_dev.tsv`, tras la primera revisión | 14 | 1 | 0 | 93,3 % | 100,0 % | Repetición de trabajo |
| `real_round6_dev.tsv`, tras la segunda revisión | 14 | 0 | 0 | 100,0 % | 100,0 % | Repetición de trabajo |
| `real_round4.tsv`, revisión `b821d56` | 18 | 1 | 2 | 94,7 % | 90,0 % | Regresión con conjunto anterior |
| `real_round5.tsv`, revisión `b821d56` | 17 | 0 | 1 | 100,0 % | 94,4 % | Regresión con conjunto anterior |
| **`real_round6_holdout.tsv`, primera corrida con `b821d56` congelado** | **16** | **0** | **1** | **100,0 %** | **94,1 %** | **Primera corrida independiente** |

Los cinco controles negativos del reservado quedaron sin hallazgos. R6HE01 fue la única omisión: su fragmento breve describe que alguien sube contactos de una agenda, pero no nombra al servicio; la etiqueta de referencia depende en parte del contexto de la página. Otra persona debería revisarla. Conservamos el texto congelado y la puntuación. En la repetición de la quinta ronda se omitió la etiqueta de subprocesadores externos de R5E08, pese a que había sido detectada en la corrida original. La salida puede variar entre corridas con temperatura cero; el resultado perfecto de la quinta ronda no debe interpretarse como garantía estable. También siguen visibles las omisiones anteriores de la cuarta ronda.

El reservado supera los umbrales elegidos del 90 % de precisión y cobertura **en estos 20 fragmentos seleccionados**. El resultado combinado de trabajo y reservado en 40 casos es de 30 etiquetas acertadas, cero falsos positivos y una omisión, pero no estima 40 casos independientes porque la mitad guió el prompt. Una persona etiquetó las muestras; las cláusulas breves y las empresas repetidas no establecen rendimiento con políticas completas. Para repetirlo desde `apps/api`, con Ollama activo:

```powershell
go run ./cmd/evaluate -data evaluation/real_round6_dev.tsv
go run ./cmd/evaluate -data evaluation/real_round6_holdout.tsv
```
