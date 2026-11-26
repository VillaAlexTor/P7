# Ronda 18: evaluación final de políticas completas de Dropbox y Airbnb

## Fuentes y referencia fijada

El 26 de noviembre de 2026, p7 analizó las páginas vigentes de privacidad de Dropbox en [inglés](https://www.dropbox.com/en/privacy) y [español](https://www.dropbox.com/es/privacy), y los PDF oficiales de Airbnb de 2026 en [inglés](https://assets.airbnb.com/help/2026_Privacy_Policy-English.pdf) y [español](https://assets.airbnb.com/help/2026_Privacy_Policy-Spanish_Spain.pdf). Son **dos empresas con dos idiomas cada una**, no cuatro políticas independientes. La política vigente de Dropbox rige desde el 30 de mayo de 2025; un aviso anuncia otra para el 1 de enero de 2027. La extracción comienza en el título de la política **vigente**. El PDF de Airbnb indica actualización del 5 de febrero de 2026. Los cuatro textos extraídos contienen entre 17.456 y 61.064 caracteres. HTML, PDF y textos permanecen en `.cache/real_policies_round18`, ignorada por Git.

El [manifiesto](../../apps/api/evaluation/real_policies_round18.json), confirmado como `0ebbc80` **antes de la inferencia**, fija URL, SHA-256, longitud, etiquetas por documento y ocho preguntas bilingües. Una persona asignó la referencia. Dropbox es positivo en `location`, `third-parties` y `retention`; los correos genéricos sobre mejoras no demuestran publicidad personalizada y «tus contactos» no identifica explícitamente una **agenda del dispositivo**. Esta última etiqueta es discutible y requiere otra revisión. Airbnb es positivo en las siete categorías: describe importación de agenda, reconocimiento facial a partir de documentos de identidad, uso de datos personales para mejorar modelos de IA, publicidad adaptada, ubicación, cesión y conservación según la finalidad.

## Resultados crudos y reparación

Ambas pasadas usaron Ollama local con `qwen3:4b-instruct`, temperatura cero, las mismas fuentes y etiquetas fijadas, perfil general y documentos completos. El [resultado crudo inicial](../../apps/api/evaluation/real_policies_round18_result_initial.json) reveló un fallo de citas PDF: el modelo propuso 8 candidatos en inglés y 20 en español para Airbnb, pero el verificador literal rechazó **los 28**, principalmente porque la extracción PDF inserta saltos de línea dentro de las frases. Cambiamos únicamente la verificación de evidencias para comparar el texto sin diferencias de espacios **dentro de la misma cláusula citada**. Una prueba confirma que se siguen rechazando palabras cambiadas y números de cláusula incorrectos. El [resultado crudo revisado](../../apps/api/evaluation/real_policies_round18_result.json) es una *comprobación sobre la misma muestra*, **no otra prueba independiente**. No se entrenaron los pesos del modelo.

| Pasada | Dropbox EN | Dropbox ES | Airbnb EN | Airbnb ES | VP | FP | FN | Precisión | Cobertura |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: |
| Inicial | 3 esperadas + contactos | 3 esperadas + contactos | ninguna | ninguna | 6 | 2 | 14 | 75,0% | 30,0% |
| Tras reparar verificador | 3 esperadas + contactos | 3 esperadas + contactos | 5 de 7; omite ubicación y contactos | 3 de 7; omite contactos, IA, biometría y conservación | 14 | 2 | 6 | 87,5% | 70,0% |

La segunda pasada tiene **0/4 páginas con el conjunto exacto** y 6 verdaderos negativos entre 28 decisiones de documento y categoría. Las ocho preguntas fijadas generan 16 comprobaciones localizadas, de las que **10/16** coinciden con la salida por categoría. Las traducciones están correlacionadas y la referencia de contactos de Dropbox es discutible. Estas cifras **no** demuestran 87,5% de fiabilidad en producción ni confianza calibrada.

## Auditoría de citas y errores pendientes

Revisé las 16 cláusulas citadas que quedaron en la segunda salida. **13/16** respaldan su categoría según las definiciones fijadas. Las dos citas `contacts` de Dropbox (`clause-010` en ambos idiomas) indican acceso a contactos y que Dropbox los guarda, pero no identifican una agenda del dispositivo. Originan los dos falsos positivos frente a la referencia y también una ambigüedad legítima de anotación. Airbnb EN cita `clause-114` para `third-parties`; esa cláusula describe datos de pago **recibidos de** un proveedor externo. La cesión hacia fuera aparece en otro lugar de la política. Es un acierto de categoría con una cita inadecuada. Las demás ubicaciones parecieron pertinentes en esta revisión de una persona. El campo automático `evidenceVerified` aún significa presencia del texto tras normalizar espacios, **no prueba semántica**.

La reparación aumentó mucho la cobertura de Airbnb, pero se rechazaron todavía 3 de 8 candidatos ingleses y 17 de 20 españoles. El PDF español omitió cuatro categorías explícitas, entre ellas mejora de modelos de IA y reconocimiento facial biométrico. La longitud del PDF, su extracción, la selección del modelo y la generación de citas son puntos de fallo distintos. El siguiente paso técnico sería una etapa específica de selección y validación semántica de citas, seguida de una prueba independiente y una segunda persona que anote la referencia. **No afirmamos 90% de fiabilidad**.

## Reproducción

El manifiesto conserva los hashes de los textos extraídos. `--fetch` obtiene **otra versión** si cambia una página oficial. Con los archivos originales en caché, instala `pypdf` si hace falta (`python -m pip install --target .cache/pypdf pypdf`), ejecuta `python scripts/prepare-real-policies-round18.py` y compara hashes. Inicia Ollama local con `qwen3:4b-instruct`; desde `apps/api`, ejecuta:

```powershell
go run ./cmd/evaluatepolicies -manifest evaluation/real_policies_round18.json -text-dir ../../.cache/real_policies_round18
```

El comando usa el código actual y el comportamiento actual del modelo; los dos JSON versionados son los registros crudos exactos de estas dos pasadas.
