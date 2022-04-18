package main

import (
	"fmt"
	"math"
)

type Rectangulo struct {
	Base   float64
	Altura float64
}

type Circulo struct {
	Radio float64
}

func (c Circulo) area() float64 {
	return c.Radio * c.Radio * math.Pi
}

func (r Rectangulo) area() float64 {
	return r.Base * r.Altura
}

func (r *Rectangulo) inc(incremento float64) { // SE DEBE COLOCAR EL * PARA DEFINIR QUE LOS CAMBIOS REPERCUTAN EN EL MODELO
	r.Altura += incremento
	r.Base += incremento
}

func main() {
	// Creacion del rectangulo
	r1 := Rectangulo{Base: 2, Altura: 3}
	fmt.Println("El area del rectangulo es:", r1.area()) // El area del rectangulo es: 6

	// Incrementar Rectangulo

	r1.inc(1)
	fmt.Println("El area del rectangulo es:", r1.area()) // El area es: 12

	// Creacion del circulo
	c1 := Circulo{12}
	fmt.Println("El area del circulo es:", c1.area()) // El area del circulo es: 452.3893421169302

	c2 := Circulo{13}
	c3 := Circulo{12}

	fmt.Println("c1 y c2 son iguales:", c1 == c2)  // C1 y C2 son iguales: false
	fmt.Println("Cc1 y c3 son iguales:", c1 == c3) // C1 y C3 son iguales: true

}
