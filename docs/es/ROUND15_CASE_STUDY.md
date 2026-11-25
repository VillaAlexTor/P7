# Ronda 15: políticas completas de Dropbox y GitHub

## Fuentes y método

El 25 de noviembre de 2026 p7 analizó las páginas oficiales completas de [Dropbox EN](https://www.dropbox.com/en/privacy), [Dropbox ES](https://www.dropbox.com/es/privacy), [GitHub EN](https://docs.github.com/en/site-policy/privacy-policies/github-general-privacy-statement) y [GitHub ES](https://docs.github.com/es/site-policy/privacy-policies/github-general-privacy-statement). Son **dos políticas en dos idiomas**, no cuatro políticas independientes. Dropbox indica publicación/vigencia del 30 de mayo de 2025 y enlaza una versión *futura* para el 1 de enero de 2027; evaluamos el texto vigente, no el futuro. GitHub indica vigencia del 27 de abril de 2026. Las extracciones tienen 17.774–47.638 caracteres. El texto completo permanece en `.cache/real_policies_round15`, excluida de Git; el [manifiesto congelado](../../apps/api/evaluation/real_policies_round15.json) fija URL, hashes exactos, longitudes, líneas de referencia, etiquetas y ocho preguntas. Se confirmó como `ea985ab` **antes de inferir**.

Una persona asignó las etiquetas. Dropbox dice que el usuario puede permitir acceso a sus contactos y su almacenamiento, pero el aviso vigente no afirma expresamente que datos personales seleccionen anuncios ni entrenen modelos de IA. Sus cookies de terceros promocionan los servicios; no inferimos segmentación solo de esa frase. GitHub afirma que puede usar Datos Personales para entrenar modelos y describe publicidad basada en intereses, pero no afirma acceder a la agenda del dispositivo. Ningún aviso afirma identificación biométrica por el proveedor. Analizamos cada página una vez con el mismo `qwen3:4b-instruct` local a temperatura cero, perfil general, prompts y código Go. El [informe de máquina](../../apps/api/evaluation/real_policies_round15_result.json) guarda categorías, IDs de cláusulas, hashes de citas, candidatos, puntuaciones y tiempos sin republicar textos ajenos.

## Resultados por documento y categoría

| Página completa | Referencia congelada | Categorías de p7 | Aciertos | Falsos positivos | Omisiones | Tiempo |
| --- | --- | --- | ---: | ---: | ---: | ---: |
| Dropbox EN | ubicación, contactos, terceros, retención | esas cuatro + **publicidad** | 4 | 1 | 0 | 25 s |
| Dropbox ES | las mismas cuatro | esas cuatro + **publicidad** | 4 | 1 | 0 | 28 s |
| GitHub EN | ubicación, terceros, publicidad, entrenamiento de IA, retención | primeras cuatro + **contactos**; **falta retención** | 4 | 1 | 1 | 62 s |
| GitHub ES | las mismas cinco | exactamente las cinco | 5 | 0 | 0 | 75 s |
| **Total** | | | **17** | **3** | **1** | **192 s** |

La precisión por categoría fue **17/20 = 85,0 %**, la cobertura **17/18 = 94,4 %** y **1/4** páginas tuvieron el conjunto exacto. Hubo 28 decisiones documento-categoría: 17 aciertos, 3 falsos positivos, 1 omisión y 7 ausencias correctas. El modelo propuso 54 hallazgos, conservó 20 coincidencias literales y contó 34 rechazos. Ese contador mezcla duplicados con evidencia inválida. Las puntuaciones de riesgo fueron 90, 94, 98 y 96 en ese orden; no son probabilidades calibradas.

## Ocho preguntas fijadas antes de inferir

Las ocho preguntas producen 16 comprobaciones por idioma que se superponen con las 28 decisiones. «No establecido» significa que el aviso no afirma la práctica definida.

| Pregunta | Referencia | p7 EN | p7 ES | Distinción |
| --- | --- | --- | --- | --- |
| D1. ¿Dropbox accede y almacena opcionalmente los contactos? | Sí | Sí | Sí | EN línea 23 / ES línea 25: acceso y almacenamiento tras elección del usuario. |
| D2. ¿El aviso vigente de Dropbox afirma anuncios dirigidos con datos personales? | No establecido | **Sí** | **Sí** | EN línea 29 / ES línea 31: cookies para promoción, sin segmentación explícita. |
| D3. ¿Afirma entrenamiento de IA con datos del usuario? | No establecido | No | No | No hay declaración afirmativa de entrenamiento en este aviso vigente. |
| D4. ¿Conserva datos mientras exista la cuenta? | Sí | Sí | Sí | EN línea 79 / ES línea 81; eliminar la cuenta inicia otro proceso. |
| G1. ¿GitHub dice que Datos Personales pueden entrenar modelos? | Sí | Sí | Sí | La línea 99 vincula expresamente Datos Personales y entrenamiento. |
| G2. ¿GitHub afirma importar la agenda del dispositivo? | No establecido | **Sí** | No | EN cita categorías generales de identificación/contacto, no una agenda. |
| G3. ¿GitHub describe publicidad basada en intereses mediante cookies? | Sí | Sí | Sí | La línea 59 la menciona expresamente. |
| G4. ¿Puede GitHub compartir Datos Personales con socios externos? | Sí | Sí | Sí | La línea 123 afirma divulgación a socios y revendedores. |

p7 acertó **13/16** comprobaciones; no son independientes de la puntuación por categoría. El cuarto error de categoría, retención de GitHub EN, quedaba fuera de las ocho preguntas fijadas. Su línea 195 dice que GitHub conserva Datos Personales mientras la cuenta siga activa y por necesidades contractuales o legales. No agregamos una novena pregunta tras ver la omisión.

## Auditoría de citas y fallos

Revisamos las 20 ubicaciones citadas aparte de las categorías por documento. **Al menos cuatro cláusulas citadas no bastan por sí solas** para demostrar el hallazgo:

| Página | Hallazgo y cláusula | Problema |
| --- | --- | --- |
| Dropbox EN | `advertising`, `clause-023` | Describe correos promocionales para planes pagos, no anuncios dirigidos con datos personales. |
| Dropbox ES | `advertising`, `clause-021` | Describe cookies de terceros para promocionar Dropbox, sin afirmar que datos personales seleccionen anuncios. |
| GitHub EN | `contacts`, `clause-204` | Enumera identificadores y datos de contacto recopilados; no menciona agenda del dispositivo. |
| GitHub ES | `third-parties`, `clause-039` | Describe datos **recibidos de** otros usuarios, no divulgación **hacia** socios externos. La categoría sí aparece en otro lugar (línea 123), pero la cita es débil. |

Los tres primeros son falsos positivos de categoría; el cuarto es un acierto de categoría con cita débil. Las otras ubicaciones parecían pertinentes en esta revisión de una persona. p7 comprueba que la cita esté literalmente en la cláusula, pero no que esta demuestre semánticamente el hallazgo. La auditoría no es una segunda anotación independiente. La referencia de `contacts` de Dropbox también admite una lectura más estricta: el aviso dice «acceso a tus contactos», pero no usa expresamente «agenda del dispositivo». Otra persona debe decidirlo antes de dar por resuelta la cobertura de esa categoría.

En las rondas diez a quince con políticas completas, los totales descriptivos son **102 aciertos, 18 falsos positivos y 8 omisiones** en 24 páginas localizadas de 12 políticas: **85,0 % de precisión y 92,7 % de cobertura**. Proveedores seleccionados, traducciones correlacionadas, una persona anotadora y cobertura desigual impiden afirmar fiabilidad general. El modelo y los prompts no cambiaron; las diferencias no muestran progreso por entrenamiento.

La siguiente corrección debe exigir evidencia que afirme la acción definitoria, separar promoción de selección de anuncios personalizados e instrumentar motivos de rechazo. Conservaremos estas páginas como prueba cerrada y evaluaremos cualquier corrección en políticas nuevas. Para reproducir mientras las páginas no cambien, ejecuta `python scripts/prepare-real-policies-round15.py --fetch` desde la raíz y compara hashes con el manifiesto. Con Ollama y `qwen3:4b-instruct`, ejecuta desde `apps/api`:

```powershell
go run ./cmd/evaluatepolicies -manifest evaluation/real_policies_round15.json -text-dir ../../.cache/real_policies_round15
```
