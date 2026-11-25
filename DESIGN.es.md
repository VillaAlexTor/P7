# Sistema de diseño de p7

## Dirección

p7 utiliza una mesa de evidencia en vez de un dashboard convencional. La interfaz toma la jerarquía de un expediente anotado: papel cálido, reglas negras precisas, títulos editoriales con serif y controles compactos sans serif. El color aparece únicamente cuando comunica riesgo, verificación o estado.

La primera pantalla es operativa. El usuario puede pegar una política y elegir perfil de inmediato; el texto promocional queda subordinado a la tarea.

## Modos

- Principal: operar.
- Secundario: leer evidencias e interpretaciones.
- Carácter sereno y exacto; sin alarmismo, neón, cristal, gradientes de IA ni rejillas de tarjetas intercambiables.

## Color

| Token | Valor | Uso |
| --- | --- | --- |
| Papel | `#f3efe4` | Superficie principal |
| Papel profundo | `#e8e1d2` | Navegación lateral y hechos |
| Tinta | `#171711` | Texto, reglas y controles |
| Tinta suave | `#5b5a50` | Texto secundario |
| Bermellón | `#d64a32` | Riesgo alto o crítico |
| Ámbar | `#b57312` | Riesgo medio o ambigüedad |
| Esmeralda | `#18705b` | Evidencia verificada |
| Ultramar | `#244fa3` | Foco de teclado |
| Violeta | `#6b4e87` | Escenario hipotético |

El color nunca es la única señal; cada estado conserva una etiqueta.

## Tipografía

- Source Serif 4 Variable para títulos, resúmenes, citas y lectura.
- Manrope Variable para navegación, controles, metadatos y tablas.
- Los cuerpos deben mantener de 65 a 75 caracteres cuando el diseño lo permita.

## Composición

- Reglas finas en vez de tarjetas con sombra.
- Entrada dividida en escritorio: declaración editorial y formulario.
- Informe con hombro estrecho para índice/navegación y columna amplia de lectura.
- Cada sección empieza con título fuerte y advertencia breve.
- En móvil, el hombro precede al informe y los diagramas largos desplazan dentro de su propio límite.

## Controles y estados

- Acciones primarias en tinta sólida; bermellón en hover y análisis activo.
- Entradas con línea base, no contenedores redondeados.
- Radios y checkboxes accesibles con marcas personalizadas visibles.
- Carga, error y recuperación nombran la operación concreta.
- Foco ultramar de dos píxeles con separación visible.

## Movimiento

Movimiento funcional: desplazamiento al informe y reglas que se extienden en hover. `prefers-reduced-motion` elimina desplazamiento suave y transiciones.

## Reglas de contenido

- Separar evidencia, interpretación e impacto posible.
- No describir una ausencia como prueba de ilegalidad.
- Etiquetar datos simulados y consecuencias hipotéticas.
- No afirmar que se observa tráfico de aplicaciones.
- Mantener paridad de significado entre inglés y español.

