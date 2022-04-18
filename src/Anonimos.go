package main

import "fmt"

// tipo persona
type Persona struct {
	Nombre string
	Edad   int
}

type Estudiante struct {
	Persona // tipo anonimo, Persona tiene todo lo de la estructura Persona
	Carrera string
}

type Profesor struct {
	Estudiante        // tipo anonimo, Estudiante
	Carrera    string // igual que estudiante
}

func main() {
	mauri := Estudiante{
		Persona{
			Nombre: "mauri",
			Edad:   15,
		},
		"Informatica",
	}
	fmt.Println("mauri:", mauri) // mauri: {{mauri 15} Informatica}

	// Acceso a los campos
	// Aunque no este declarado implicitamente en Estudiante se puede acceder a los campos de Persona
	fmt.Println("Nombre:", mauri.Nombre)   // Nombre: mauri
	fmt.Println("Edad:", mauri.Edad)       // Edad: 15
	fmt.Println("Carrera:", mauri.Carrera) // Carrera: Informatica

	// Mismo campo
	// Siempre se llamra al campo de la estructura padre

	pedro := Profesor{
		Estudiante{
			Persona{
				Nombre: "Pedro",
				Edad:   22,
			},
			"Contabilidad",
		},
		"Informatica",
	}
	fmt.Println("Pedro:", pedro) // Pedro: {{{Pedro 22} Contabilidad} Informatica}

	fmt.Println("Nombre:", pedro.Nombre)   // Nombre: Pedro
	fmt.Println("Edad:", pedro.Edad)       // Edad: 22
	fmt.Println("Carrera:", pedro.Carrera) // Carrera: Informatica

	// llamada a carrera de estudiante
	fmt.Println("Carrera:", pedro.Estudiante.Carrera) // Carrera: Contabilidad

	// Otra manera de inicializar
	var profe Profesor
	profe.Nombre = "Jorge"
	profe.Edad = 23
	profe.Carrera = "Informatica"
	profe.Estudiante.Carrera = "Educacion"
	fmt.Println("Profesor:", profe) // Profesor: {{{Jorge 23} Educacion} Informatica}

}
