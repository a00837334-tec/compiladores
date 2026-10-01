package estructuras

import "testing"

// Verifica el orden LIFO: el último en entrar es el primero en salir.
func TestStackOrdenLIFO(t *testing.T) {
	var s Stack[int]
	s.Push(1)
	s.Push(2)
	s.Push(3)

	esperados := []int{3, 2, 1}
	for _, esperado := range esperados {
		v, ok := s.Pop()
		if !ok || v != esperado {
			t.Errorf("Pop() = %v, %v; se esperaba %v, true", v, ok, esperado)
		}
	}
}

// Verifica que Pop en una pila vacía no truene y regrese false.
func TestStackPopVacia(t *testing.T) {
	var s Stack[int]
	if v, ok := s.Pop(); ok {
		t.Errorf("Pop() en vacía = %v, %v; se esperaba 0, false", v, ok)
	}
}

// Verifica que Peek regresa el tope pero no lo quita.
func TestStackPeekNoModifica(t *testing.T) {
	var s Stack[int]
	s.Push(5)
	s.Push(7)

	peek, ok := s.Peek()
	if !ok || peek != 7 {
		t.Errorf("Peek() = %v, %v; se esperaba 7, true", peek, ok)
	}

	size := s.Size()
	if size != 2 {
		t.Errorf("Size() después de Peek = %v; se esperaba 2", size)
	}
}

// Verifica el orden FIFO: el primero en entrar es el primero en salir.
func TestQueueOrdenFIFO(t *testing.T) {
	var q Queue[string]
	q.Enqueue("a")
	q.Enqueue("b")
	q.Enqueue("c")

	esperados := []string{"a", "b", "c"}
	for _, esperado := range esperados {
		v, ok := q.Dequeue()
		if !ok || v != esperado {
			t.Errorf("Dequeue() = %v, %v; se esperaba %v, true", v, ok, esperado)
		}
	}
}

// Verifica que Dequeue en una cola vacía no truene y regrese false.
func TestQueueDequeueVacia(t *testing.T) {
	var q Queue[string]
	if v, ok := q.Dequeue(); ok {
		t.Errorf("Dequeue() en vacía = %v, %v; se esperaba \"\", false", v, ok)
	}
}

// Verifica que Set con una llave repetida actualiza sin duplicar.
func TestTableSetActualiza(t *testing.T) {
	tabla := NewTable[string, int]()
	tabla.Set("x", 1)
	tabla.Set("x", 2)

	if size := tabla.Size(); size != 1 {
		t.Errorf("Size() después de Set repetido = %v; se esperaba 1", size)
	}
	if v, ok := tabla.Get("x"); !ok || v != 2 {
		t.Errorf("Get(\"x\") = %v, %v; se esperaba 2, true", v, ok)
	}
}

// Verifica que Keys conserva el orden de inserción, también después de un Delete.
func TestTableOrdenYDelete(t *testing.T) {
	tabla := NewTable[string, int]()
	tabla.Set("x", 1)
	tabla.Set("y", 2)
	tabla.Set("z", 3)
	tabla.Delete("y")

	esperadas := []string{"x", "z"}
	llaves := tabla.Keys()

	if len(llaves) != len(esperadas) {
		t.Errorf("Keys() = %v; se esperaba %v", llaves, esperadas)
	}

	for i, esperado := range esperadas {
		if llaves[i] != esperado {
			t.Errorf("Keys()[%d] = %v; se esperaba %v", i, llaves[i], esperado)
		}
	}
}

// Verifica que Get y Delete de una llave inexistente regresan false.
func TestTableLlaveInexistente(t *testing.T) {
	tabla := NewTable[string, int]()

	_, okGet := tabla.Get("w")
	if okGet {
		t.Errorf("Get(\"w\") = _, %v; se esperaba false", okGet)
	}
	
	okDelete := tabla.Delete("w")
	if okDelete {
		t.Errorf("Delete(\"w\") = %v; se esperaba false", okDelete)
	}
}
