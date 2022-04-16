package main

import "fmt"

func main() {
	var numeros [5]int   // crear un array de 5 posiciones con 0
	fmt.Println(numeros) // [0 0 0 0 0]

	numeros[0] = 10
	numeros[1] = 20
	fmt.Println(numeros) // [10 20 0 0 0]

	//Arrays con valores
	nombres := [3]string{"Mauri", "Daysi", "Jessie"}
	fmt.Println(nombres) // [Mauri Daysi Jessie]

	// Arreglos sin longitud exacta, pero no se puede agregar mas elementos
	colores := [...]string{
		"rojo",
		"negro",
		"azul",
	}
	// len() para longitud
	fmt.Println(colores, len(colores)) // [rojo negro azul] 3

	// Indice definido
	monedas := [...]string{
		0: "Dolar",
		3: "Soles",
		5: "Pesos",
	}
	fmt.Println(monedas, len(monedas)) // [Dolar   Soles  Pesos] 6
	// Se agregan caracteres vacios a los indices que no se definen

	// Sub Array
	subMoneda := monedas[0:3]              // va desde 0 al indice 2
	fmt.Println(subMoneda, len(subMoneda)) // [Dolar  ] 3

}
