# Ronda 12: políticas de privacidad largas y difíciles

## Qué probamos

El 25 de noviembre de 2026 evaluamos p7 con las páginas oficiales completas de [Google](https://policies.google.com/privacy?hl=es) y [WhatsApp para el EEE](https://www.whatsapp.com/legal/privacy-policy-eea?lang=es_LA), cada una en inglés y español. La política de Google entró en vigor el 16 de octubre de 2026 y la de WhatsApp, el 2 de julio de 2026. Son **dos políticas subyacentes y cuatro páginas localizadas**, no cuatro políticas independientes. Los textos extraídos tienen entre **57.636 y 120.506 caracteres** y entre **431 y 976 cláusulas de p7**. Todos caben en el límite actual de 125.000 caracteres de la API. Las extracciones incluyen algo de navegación y texto repetido de tablas o declaraciones, como puede ocurrir al copiar una página web larga.

Fijamos las [URL, los hashes SHA-256, las etiquetas y ocho preguntas de control](../../apps/api/evaluation/real_policies_round12.json) en el commit `194da5e` **antes** de ejecutar el modelo. Una persona aplicó las mismas definiciones de siete categorías de las rondas diez y once. Los textos descargados permanecen en `.cache/real_policies_round12`, excluida de Git; no se republican. El [informe compacto de máquina](../../apps/api/evaluation/real_policies_round12_result.json) registra todas las categorías finales, IDs de cláusulas, hashes de citas, puntuaciones, tiempos y recuentos de candidatos. Usamos el mismo `qwen3:4b-instruct` local, temperatura cero y perfil general. Analizamos cada página completa una sola vez, sin cambiar modelo ni prompt.

## Qué devolvió p7

| Página | Categorías de referencia | Categorías de p7 | Aciertos | Falsos positivos | Omisiones | Tiempo |
| --- | --- | --- | ---: | ---: | ---: | ---: |
| Google EN | ubicación, terceros, publicidad, conservación | esas cuatro + contactos + entrenamiento de IA | 4 | 2 | 0 | 76 s |
| Google ES | ubicación, terceros, publicidad, conservación | esas cuatro + contactos + entrenamiento de IA | 4 | 2 | 0 | 119 s |
| WhatsApp EN | ubicación, contactos, terceros, conservación | esas cuatro + publicidad + biometría | 4 | 2 | 0 | 175 s |
| WhatsApp ES | ubicación, contactos, terceros, conservación | exactamente esas cuatro | 4 | 0 | 0 | 274 s |
| **Total** | | | **16** | **6** | **0** | **646 s** |

**La precisión fue 16/22 = 72,7 %; la cobertura, 16/16 = 100 %; los conjuntos exactos, 1/4 documentos.** Esta ronda no contiene ejemplos afirmativos de entrenamiento de IA ni identificación biométrica según las etiquetas de referencia, por lo que no mide la cobertura de esas categorías. El modelo propuso 235 hallazgos; p7 conservó 22 y contabilizó 213 rechazos. Ese contador mezcla duplicados de categoría con evidencia inválida: no permite saber cuántas citas literales fallaron. Las puntuaciones de riesgo fueron 100 en Google EN, Google ES y WhatsApp EN, y 76 en WhatsApp ES. No son probabilidades validadas.

## Ocho preguntas fijadas antes de la inferencia

Respondemos cada pregunta según la referencia de la política y comparamos si p7 devolvió la categoría correspondiente. Son **16 comprobaciones por idioma tomadas de las mismas 28 decisiones de categoría y documento**; no constituyen una prueba de exactitud independiente. «No consta» significa que la política no confirma la práctica definida; no demuestra que la práctica jamás ocurra.

| Pregunta | Referencia | p7 EN | p7 ES | Por qué es difícil |
| --- | --- | --- | --- | --- |
| G1. ¿Google recoge información de ubicación? | Sí | Sí | Sí | Conviven ubicación precisa, aproximada y controles opcionales. |
| G2. ¿Google declara cargar la agenda del dispositivo? | No consta | **Sí** | **Sí** | Sugerir destinatarios y guardar datos de contacto no confirma acceso a la agenda. |
| G3. ¿El pasaje de datos web públicos confirma entrenamiento con contenido privado de la cuenta? | No consta | **Sí** | **Sí** | Menciona entrenamiento, pero identifica material públicamente disponible. La etiqueta requiere revisión independiente. |
| G4. ¿Google conserva algunos datos por fines comerciales o legales sin fecha fija? | Sí | Sí | Sí | Ejemplos de borrado con fecha coexisten con conservación amplia. |
| W1. ¿WhatsApp permite cargar la agenda del dispositivo? | Sí | Sí | Sí | Es una función opcional, pero confirma carga cuando se usa. |
| W2. ¿Este aviso confirma anuncios dirigidos mediante datos personales? | No consta | **Sí** | No | Cookies para anuncios y anuncios en Estados/Canales no prueban por sí solos segmentación. |
| W3. ¿Usar Meta AI en WhatsApp demuestra por sí solo entrenamiento con chats? | No consta | No | No | Enviar datos para usar una función de IA no equivale a afirmar entrenamiento. |
| W4. ¿WhatsApp conserva algunos datos mientras la cuenta esté activa? | Sí | Sí | Sí | Los plazos fijos para otros datos no eliminan esa conservación abierta. |

p7 acertó **11 de estas 16** comprobaciones bajo las etiquetas congeladas. La novena dificultad, no predefinida como pregunta, apareció tras la corrida: WhatsApp EN marcó `biometric` por una selfie en vídeo enviada a un proveedor para estimar la edad. La cláusula trata de estimación de edad y borrado inmediato, no de identificar a una persona mediante un identificador biométrico. WhatsApp ES no produjo ese hallazgo. No añadimos la pregunta a las ocho preconfirmadas ni cambiamos etiquetas al verlo.

## Errores y sensibilidad de las etiquetas

1. **Datos de contacto frente a agenda:** la cláusula `164` de Google EN y la `181` de Google ES hablan de nombre, correo y teléfono de la cuenta. Ambos hallazgos fueron falsos positivos según la definición de la categoría. Es la tercera ronda de políticas completas con este tipo de confusión; en cambio, p7 sí detectó la carga explícita de la agenda de WhatsApp en ambos idiomas.
2. **Material público frente a datos privados del usuario:** la cláusula `377` en Google EN y ES describe datos públicamente disponibles que pueden ayudar a entrenar modelos de IA. La etiqueta congelada exigió evidencia de entrenamiento con *contenido privado de la cuenta*. Es el juicio más discutible. Si otra persona considera que los datos personales públicos cuentan como datos de usuario para esta categoría, los dos hallazgos de Google pasarían a aciertos: la ronda tendría **18 aciertos, 4 falsos positivos, 0 omisiones y 81,8 % de precisión**. Conservamos el resultado original de 72,7 % y mostramos la interpretación alternativa.
3. **Anuncios frente a segmentación:** la cláusula `074` de WhatsApp EN menciona tecnologías para anuncios en funciones opcionales; no declara que los datos personales elijan los anuncios. La ejecución española no hizo esa inferencia. La negación de banners de terceros tampoco significa que no exista ningún otro tipo de anuncio.
4. **Estimación de edad frente a reconocimiento de identidad:** la cláusula `117` de WhatsApp EN describe que un proveedor recibe imágenes de una selfie en vídeo para estimar edad y las elimina después. La referencia biométrica exige un rasgo usado para reconocer la identidad. La ejecución española evitó este falso positivo.
5. **Usabilidad con textos largos:** WhatsApp ES tardó **274 segundos** y las cuatro páginas sumaron **10 minutos y 46 segundos** en este equipo local. Puede ser demasiado lento para una demo en vivo o una lectura frecuente. No medimos el tiempo completo desde el navegador.

En las rondas diez a doce, el flujo sin cambios acumula **41 aciertos, 13 falsos positivos y 5 omisiones** en 12 páginas localizadas que representan seis políticas subyacentes: **75,9 % de precisión y 89,1 % de cobertura** como totales descriptivos. Los idiomas están correlacionados, las políticas fueron seleccionadas y una sola persona asignó todas las etiquetas. No es una estimación estadística de fiabilidad ni evidencia de entrenamiento. La mejor cobertura de esta ronda puede deberse a sus declaraciones más explícitas. Una cita literal ayuda a rastrear un hallazgo, pero no hace correcta su interpretación.

## Qué mejorar después

- Exigir acceso, importación, sincronización o carga explícita de la **agenda del dispositivo** para `contacts`; los campos de contacto y sugerencias de destinatarios no deben bastar.
- Exigir que la política declare que **usa datos personales para elegir o personalizar anuncios**; mencionar cookies o lugares donde aparecen anuncios no basta.
- Separar entrenamiento con datos web públicos del entrenamiento con datos privados de usuarios y pedir a otra persona que revise la etiqueta de Google antes de redefinir la categoría.
- Exigir reconocimiento de identidad para `biometric`; estimar edad mediante una selfie es un concepto distinto.
- Registrar por qué se rechazó cada candidato (duplicado, ID de cláusula inválido, cita inexacta o categoría desconocida) y medir la latencia de documentos largos. Probar las correcciones con **políticas completas nuevas**, no con estas páginas.

Para reproducir la extracción exacta mientras no cambien las páginas publicadas, ejecuta desde la raíz del repositorio:

```powershell
python scripts/prepare-real-policies-round12.py --fetch
cd apps/api
go run ./cmd/evaluatepolicies -manifest evaluation/real_policies_round12.json -text-dir ../../.cache/real_policies_round12
```
