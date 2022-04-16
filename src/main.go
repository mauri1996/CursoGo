package main

import "fmt"

func main() {
	// Slicen
	numeros := []int{1, 2, 3}          // no se colocan la cantidad de elementos, son objetos mutables
	fmt.Println(numeros, len(numeros)) // [1 2 3] 3

	// Agregar elementos
	numeros = append(numeros, 4)       // agregar 4 a numeros
	numeros = append(numeros, 5)       // agregar 5 a numeros
	fmt.Println(numeros, len(numeros)) // [1 2 3 4 5] 5

	//Sub Slice
	subSlicen := numeros[:2] // desde indice 0 hasta indice 2
	numeros[0] = 100         // modifica posicion 0

	// si se modifica el original se modificara los subSlice, ya que un Slicen maneja referencias
	fmt.Println(subSlicen) // [100 2]
	fmt.Println(numeros)   // [100 2 3 4 5]

	//Punteros
	//Longitud
	//Capacidad

	meses := []string{"Enero", "Febrero", "Marzo"}
	fmt.Printf("Len: %v, Cap: %v, %p \n", len(meses), cap(meses), meses) // Len: 3, Cap: 3, 0xc000076480

	meses = append(meses, "Abril")                                       // se agrega un nuevo Slicen cada vez que se agrega un elemento nuevo
	fmt.Printf("Len: %v, Cap: %v, %p \n", len(meses), cap(meses), meses) // Len: 4, Cap: 6, 0xc0000220c0
}
