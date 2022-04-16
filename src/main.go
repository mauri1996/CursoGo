package main

import "fmt"

func main() {
	// Crear Slicen
	// numeros := make([]int, 0, 3)                     // Slicen vacio de longitud 0 y capacidad 3
	// fmt.Println(numeros, len(numeros), cap(numeros)) // [] 0 3

	// numeros[0] = 100 // ERROR
	// numeros[1] = 200 // ERROR
	// numeros[2] = 300 // ERROR

	// Se debe modificar la longitud
	numeros := make([]int, 3, 3) // Slicen vacio de longitud 0 y capacidad 3

	numeros[0] = 100
	numeros[1] = 200
	numeros[2] = 300

	numeros = append(numeros, 400)

	fmt.Println(numeros, len(numeros), cap(numeros)) // [100 200 300 400] 4 6
}
