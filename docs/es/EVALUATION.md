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

## Siguiente evaluación

Revisar las 11 omisiones y luego reunir un conjunto etiquetado de políticas públicas diversas con permiso para citar fragmentos. Pedir a otra persona que revise las etiquetas e informar precisión y cobertura por idioma con ejemplos de desacuerdos. Conservar ambos conjuntos sintéticos como pruebas de regresión; cualquier cambio del prompt deberá evaluarse con otro conjunto intacto.
