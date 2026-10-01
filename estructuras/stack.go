package estructuras

// Stack es una pila LIFO genérica.
type Stack[T any] struct {
	items []T
}

// Push agrega un elemento al tope.
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop quita y regresa el elemento del tope.
// Si está vacía regresa el valor cero de T y false.
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if s.IsEmpty() {
	return zero, false
}

ultimo := len(s.items) - 1
item := s.items[ultimo]
s.items[ultimo] = zero
s.items = s.items[:ultimo]
return item, true
}

// Peek regresa el tope SIN quitarlo.
func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

// IsEmpty indica si la pila está vacía.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size regresa cuántos elementos hay.
func (s *Stack[T]) Size() int {
	return len(s.items)
}
