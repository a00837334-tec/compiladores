# Compilador – Módulo 3: Compiladores

Proyecto del curso Desarrollo de Aplicaciones Avanzadas de Ciencias Computacionales.
Lenguaje: go version go1.27.1 darwin/arm64.

## Tarea 1: Estructuras de datos

Implementación desde cero de tres estructuras genéricas en el paquete `estructuras`:

| Estructura | Archivo | Operaciones | Uso previsto en el compilador |
|---|---|---|---|
| Stack (LIFO) | `stack.go` | Push, Pop, Peek, IsEmpty, Size | Manejo de scopes y evaluación de expresiones |
| Queue (FIFO) | `queue.go` | Enqueue, Dequeue, Front, IsEmpty, Size | Flujo de tokens entre lexer y parser |
| Table (diccionario ordenado) | `table.go` | Set, Get, Contains, Delete, Keys, Size | Tabla de símbolos |

### Decisiones de diseño

- **Genéricos** (`[T any]`, `[K comparable, V any]`) para reutilizar las estructuras con cualquier tipo.
- **Stack y Queue** se construyen sobre slices. Push/Pop y Enqueue/Dequeue son O(1) amortizado.
  Al sacar un elemento se limpia su posición para que el recolector de basura pueda liberarlo.
- **Table** combina un `map` (búsqueda O(1)) con un slice de llaves, porque el `map` de Go
  no garantiza orden al recorrerse. Así las llaves se mantienen en orden de inserción.
  Trade-off: `Delete` es O(n) por la búsqueda en el slice, aceptable para una tabla de
  símbolos donde casi no se borra.
- Las operaciones sobre estructuras vacías no generan panic: regresan el valor cero y `false`
  (patrón "comma ok").

## Cómo ejecutar

```bash
go run .              # programa de demostración
go test ./... -v      # pruebas unitarias
```

## Test cases

| Test | Qué valida | Resultado |
|---|---|---|
| TestStackOrdenLIFO | Push 1,2,3 y Pop regresa 3,2,1 | PASS |
| TestStackPopVacia | Pop en pila vacía regresa false sin panic | PASS |
| TestStackPeekNoModifica | Peek regresa el tope sin cambiar el tamaño | PASS |
| TestQueueOrdenFIFO | Enqueue a,b,c y Dequeue regresa a,b,c | PASS |
| TestQueueDequeueVacia | Dequeue en cola vacía regresa false sin panic | PASS |
| TestTableSetActualiza | Set repetido actualiza el valor sin duplicar la llave | PASS |
| TestTableOrdenYDelete | Keys conserva el orden de inserción tras borrar la llave del medio | PASS |
| TestTableLlaveInexistente | Get y Delete de una llave inexistente regresan false | PASS |

Durante el desarrollo, una primera versión de `Dequeue` sacaba el último elemento
(comportamiento LIFO). El caso de orden FIFO es precisamente el que detecta ese error.

## Uso de IA

- **Herramienta:** Claude (Anthropic), modelo Claude Opus 5.5, en claude.ai.
- **Modo de uso:** como tutor. La IA explicó conceptos y proporcionó esqueletos con
  secciones por completar; la implementación se escribió y depuró de forma guiada,
  con revisión de código y explicación de errores de compilación.
- **Conversación completa:** (https://claude.ai/share/a5b770d2-5c5d-425a-924a-ca8fd8dcb87c)
- **Prompts principales:**
  1. "Requiero hacer la tarea 1. Pero, la cosa es que nos pidió aprender un lenguaje nuevo
     para el proyecto, estoy pensando en Rust o Go. Que me recomiendas en cuanto a que se
     utiliza cada uno para y ventajas competitivas en el mercado laboral."
  2. "Mi indecisión esta basada en que a parte de este proyecto, el negocio que estamos
     haciendo incluye desarrollo de aplicaciones. [...] quiero saber si es recomendable
     aprender Rust para un mejor funcionamiento en desarrollo o si seria mejor Go"
  3. "Ok, adoptaremos Go [...] explícame de que va la tarea 1 ahora que tomamos la decisión
     de usar Go, y guíame para resolverla. Quiero aprenderlo asi que requiero que seas guía"
  4. "Requiero de tu apoyo con los TODO. Puedes guiarme paso a paso? Que esto me sirva como
     repaso para entender bien que estamos haciendo"
  5. Preguntas conceptuales sobre el comportamiento de Queue y Table, envío de mi código
     para revisión y de errores de compilación/pruebas para su diagnóstico.