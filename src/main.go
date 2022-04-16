package main

import "fmt"

func main() {

	// doc https://pkg.go.dev/fmt

	helloMessage := "Hello"
	worldMessage := "world"

	// Println -> imprime en 1 linea
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	//Printf
	nombre := "Platzi"
	cursos := 500
	// %s string
	// %d int
	// %v si no se sabe los tipos de datos
	fmt.Printf("%s tiene mas de %d cursos \n", nombre, cursos) // Platzi tiene mas de 500 cursos
	fmt.Printf("%v tiene mas de %v cursos \n", nombre, cursos) // Platzi tiene mas de 500 cursos

	//Sprintf
	// todo lo generado se devuelve a message
	message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, cursos)
	fmt.Println(message)

	//Tipo de dato
	fmt.Printf("helloMessage: %T \n", helloMessage) //helloMessage: string
	fmt.Printf("Cursos: %T\n", cursos)              // Cursos: int
}
