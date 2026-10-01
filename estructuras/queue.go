package estructuras

// Queue es una cola FIFO genérica construida sobre un slice.
// El frente es el índice 0; los elementos nuevos entran al final.
type Queue[T any] struct {
	items []T
}

// Enqueue agrega un elemento al final de la cola.
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Dequeue quita y regresa el elemento del frente.
// Si está vacía regresa el valor cero de T y false.
func (q *Queue[T]) Dequeue() (T, bool) {
	var zero T
	if q.IsEmpty() {
		return zero, false
	}
	item := q.items[0]
	q.items[0] = zero
	q.items = q.items[1:]
	return item, true
}

// Front regresa el elemento del frente sin quitarlo.
func (q *Queue[T]) Front() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}
	return q.items[0], true
}

// IsEmpty indica si la cola está vacía.
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Size regresa cuántos elementos hay.
func (q *Queue[T]) Size() int {
	return len(q.items)
}
