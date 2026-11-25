# Decisión de despliegue para la demostración

Para esta entrega del hackathon, la ruta fiable es un repositorio público, instalación local reproducible y una demostración grabada del flujo completo. Ollama corre en la computadora del presentador; el archivo del modelo se descarga por separado y no entra en Git. Esto evita prometer una web pública que no pueda comunicarse con Ollama instalado en una laptop privada.

Ejecuta `ollama pull qwen3:4b-instruct`, inicia Ollama, la API Go y la web Next.js como indica el README de la raíz. Comprueba el procedimiento desde un checkout limpio antes de grabar. El repositorio público y el video cumplen los requisitos de entrega publicados; una URL alojada puede ayudar, pero las reglas oficiales no la enumeran como obligatoria.

Un despliegue público posterior requiere una máquina siempre disponible con memoria suficiente para el modelo, frontend y API con HTTPS, límites de peticiones, tiempos de espera, monitoreo y una decisión sobre el tratamiento de las políticas de usuarios. Docker Compose por sí solo no aporta esas garantías de producción. No expongas el puerto de Ollama a Internet sin un diseño de seguridad deliberado.
