package main

import "fmt"

// Creacion de estructura de una persona
// Estructura es lo mas parecido a Clases
type Persona struct {
	Nombre string
	Edad   int
}

func main() {

	// Creacion de una Persona
	var p Persona
	p.Nombre = "Mauri" // Adicion de datos
	p.Edad = 25

	fmt.Println("Estructura de tipo Persona", p) // Estructura de tipo Persona {Mauri 25}
	fmt.Println("Nombre:", p.Nombre)             // Nombre: Mauri
	fmt.Println("Edad:", p.Edad)                 // Edad: 25

	// Creacion de Persona metodo 2
	p2 := Persona{Nombre: "Darwin", Edad: 25}
	fmt.Println("Estructura de tipo Persona", p2) // Estructura de tipo Persona {Darwin 25}

	// Creacion de Persona metodo 3 No recomendable
	p3 := Persona{"Daysi", 24}
	fmt.Println("Estructura de tipo Persona", p3) // Estructura de tipo Persona {Daysi 24}

}
