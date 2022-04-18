package main

import "fmt"

type Persona struct {
	Nombre string
	Edad   int
}

func older(p1, p2 Persona) (personaMenor string, diferencia int) { // Otra forma de retornar valores sin la necesidad de declararlos
	if p1.Edad < p2.Edad {
		personaMenor = p1.Nombre
		diferencia = p2.Edad - p1.Edad
	} else {
		personaMenor = p2.Nombre
		diferencia = p1.Edad - p2.Edad
	}
	return
}

func procesar(p1, p2 Persona) {
	menor, diferencia := older(p1, p2)
	fmt.Printf("Entre %s y %s, %s es menor, con una diferencia de %d años\n",
		p1.Nombre, p2.Nombre, menor, diferencia)
}

func main() {
	// Ejericio, se debe pasar 2 personas a una funcion y se debe
	// decir cual de los 2 es menor y cual es la diferencia.

	p1 := Persona{Nombre: "Michael", Edad: 10}
	p2 := Persona{Nombre: "Ani", Edad: 20}
	p3 := Persona{Nombre: "Mauri", Edad: 30}

	// 1 Solucion
	// menor, diferencia := older(p1, p2)
	// fmt.Printf("Entre %s y %s, %s es menor, con una diferencia de %d años\n",
	// 	p1.Nombre, p2.Nombre, menor, diferencia)

	// menor, diferencia = older(p1, p3)
	// fmt.Printf("Entre %s y %s, %s es menor, con una diferencia de %v años\n",
	// 	p1.Nombre, p3.Nombre, menor, diferencia)

	// menor, diferencia = older(p2, p3)
	// fmt.Printf("Entre %s y %s, %s es menor, con una diferencia de %v años\n",
	// 	p2.Nombre, p3.Nombre, menor, diferencia)

	// 2 Solucion
	procesar(p1, p2)
	procesar(p1, p3)
	procesar(p2, p3)

}
