package main

import (
	"fmt"

	"github.com/a00837334-tec/compilador/estructuras"
)

func main() {
	var pila estructuras.Stack[int]

	fmt.Println("=== STACK ===")

	pila.Push(10)
	pila.Push(20)
	pila.Push(30)
	fmt.Println("Tamaño:", pila.Size())

	tope, _ := pila.Peek()
	fmt.Println("Peek:", tope)

	for !pila.IsEmpty() {
		v, _ := pila.Pop()
		fmt.Println("Pop:", v)
	}

	v, ok := pila.Pop()
	fmt.Println("Pop en vacía:", v, ok)

	fmt.Println("=== QUEUE ===")

		var cola estructuras.Queue[string]
	cola.Enqueue("a")
	cola.Enqueue("b")
	cola.Enqueue("c")

	frente, _ := cola.Front()
	fmt.Println("Front:", frente) // a

	for !cola.IsEmpty() {
		v, _ := cola.Dequeue()
		fmt.Println("Dequeue:", v) // a, b, c → FIFO
	}

	w, ok := cola.Dequeue()
	fmt.Printf("Dequeue en vacía: %q %v\n", w, ok) // "" false

	fmt.Println("=== TABLE ===")

		simbolos := estructuras.NewTable[string, string]()
	simbolos.Set("x", "int")
	simbolos.Set("y", "float")
	simbolos.Set("z", "bool")
	simbolos.Set("x", "string") // actualiza, no duplica

	fmt.Println("Tamaño:", simbolos.Size()) // 3

	for _, nombre := range simbolos.Keys() {
		tipo, _ := simbolos.Get(nombre)
		fmt.Println(nombre, "→", tipo) // x string, y float, z bool (siempre en ese orden)
	}

	fmt.Println("¿Existe y?", simbolos.Contains("y")) // true
	simbolos.Delete("y")
	fmt.Println("¿Existe y?", simbolos.Contains("y")) // false
	fmt.Println("Llaves:", simbolos.Keys())           // [x z]

	_, existeW := simbolos.Get("w")
	fmt.Println("Get de llave inexistente:", existeW) // false
}

