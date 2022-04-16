package main

import "fmt"

func incrementarNormal(numero int) int {
	numero++
	return numero
}

func incrementarSinPuntero(numero int) {
	numero++
	fmt.Println(numero) // 3
}

func incrementarConPuntero(numero *int) {
	*numero++
	fmt.Println(*numero) // 3
}

func main() {
	numero := 2
	numero = incrementarNormal(numero)
	fmt.Println("Numero sin puntero:", numero) // Numero sin puntero: 3

	puntero := 2
	incrementarSinPuntero(puntero)
	fmt.Println("Incrementeo sin puntero:", puntero) // Incrementeo sin puntero: 2

	punteroVerdad := 2
	incrementarConPuntero(&punteroVerdad)                  //& paso por puntero
	fmt.Println("Incrementeo con puntero:", punteroVerdad) // Incrementeo sin puntero: 2
}
