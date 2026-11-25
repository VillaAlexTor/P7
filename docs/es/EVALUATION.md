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

## Siguiente evaluación

Antes de afirmar utilidad en producción, reunir un conjunto etiquetado de políticas públicas diversas con permiso para citar fragmentos, pedir a otra persona que revise las etiquetas discutidas, medir la latencia con documentos completos e informar precisión y cobertura por idioma. Conservar los seis conjuntos sintéticos como pruebas de regresión; cualquier cambio posterior del prompt necesita otro conjunto intacto.
