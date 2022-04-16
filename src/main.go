package main

import "fmt"

func main() {

	// Area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area Cuadrado:", areaCuadrado)

	x := 10
	y := 50

	//Suma
	result := x + y
	fmt.Println("Suma:", result)

	//Resta
	result = y - x // se coloca = puesto uqe ya fue ini
	fmt.Println("Resta:", result)

	//Multiplicación
	result = x * y
	fmt.Println("Multiplicación:", result)

	//Division
	result = y / x
	fmt.Println("Division:", result)

	//Modulo
	result = y % x
	fmt.Println("Modulo: ", result)

	//Incremental
	x++
	fmt.Println("Incremental: ", x)

	//Decremental
	x--
	fmt.Println("Decremental: ", x)

	// Reto Calculo de area de:
	// Rectangulo, Trapecio y Circulo

	// Variables
	x = 2
	y = 3

	// Rectangulo
	AreaRectangulo := x * y
	fmt.Println("Area de rectangulo:", AreaRectangulo)

	// Trapecio
	const a float64 = 2
	const b float64 = 3
	const h float64 = 3

	AreaTrapecio := h * ((a + b) / 2)
	fmt.Println("Area de trapecio:", AreaTrapecio)

	// Circulo
	const pi float64 = 3.1415
	const radio = 2

	AreaCirculo := radio * pi
	fmt.Println("Area de circulo:", AreaCirculo)
}
