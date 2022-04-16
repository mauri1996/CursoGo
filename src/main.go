package main

import "fmt"

func saludar(nombre string) {
	fmt.Println("Hola", nombre)
}

func sumar(a, b float64) float64 {
	return a + b
}

func main() {
	saludar("Mauri")
	fmt.Println("Suma de 10 y 20:", sumar(10, 20))
}
