package main

import "fmt"

func main() {
	// Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// Declaración de variables enteras
	base := 12 // : dice que la variable no fue creada anteriormente
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero Value -> valor por defecto
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d) // 0 0   false

	// Area de un cuadrado

	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("Area Cuadrado: ", areaCuadrado)

}
