package estructuras

// Table es un diccionario genérico que recuerda el orden de inserción.
// Usa un map para búsquedas O(1) y un slice para conservar el orden.
type Table[K comparable, V any] struct {
	data map[K]V
	keys []K
}

// NewTable crea una tabla vacía lista para usarse.
func NewTable[K comparable, V any]() *Table[K, V] {
	return &Table[K, V]{
		data: make(map[K]V),
	}
}

// Set inserta una llave nueva o actualiza el valor de una existente.
// Si la llave ya existía, conserva su posición original.
func (t *Table[K, V]) Set(key K, value V) {
	if _, exists := t.data[key]; !exists {
		t.keys = append(t.keys, key)
	}
	t.data[key] = value
}

// Get regresa el valor de una llave y true, o el valor cero y false si no existe.
func (t *Table[K, V]) Get(key K) (V, bool) {
	value, exists := t.data[key]
	return value, exists
}

// Contains indica si la llave existe en la tabla.
func (t *Table[K, V]) Contains(key K) bool {
	_, exists := t.data[key]
	return exists
}

// Delete elimina una llave. Regresa false si no existía.
func (t *Table[K, V]) Delete(key K) bool {
	if _, exists := t.data[key]; !exists {
		return false
	}
	delete(t.data, key)
	for i, k := range t.keys {
		if k == key {
			t.keys = append(t.keys[:i], t.keys[i+1:]...)
			break
		}
	}
	return true
}

// Keys regresa las llaves en orden de inserción.
func (t *Table[K, V]) Keys() []K {
	copia := make([]K, len(t.keys))
	copy(copia, t.keys)
	return copia
}

// Size regresa cuántas llaves hay.
func (t *Table[K, V]) Size() int {
	return len(t.keys)
}
