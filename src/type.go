package main

import "fmt"

// Creacion de type
type Dinero int

// Generacion del metodo String para el tipo Dinero
func (d Dinero) String() string {
	return fmt.Sprintf("$%d", d) //$25000
}

func main() {
	// Creacion de variable
	var sueldo Dinero
	sueldo = 25000
	// fmt.Println("Sueldo:", sueldo) // Sueldo: 25000  Sin usar la funcion

	fmt.Println("Sueldo:", sueldo) // Sueldo: $25000 Usa la funcion String al imprimir

	aumento := 10000

	// sueldo += aumento // No se puede hacer debido a que sueldo es de tipo Dinero y aumento es int

	sueldo = Dinero(aumento)                 // convierte a tipo Dinero -> es como poner int(2.56) -> convierte a int el float
	fmt.Println("Sueldo + Aumento:", sueldo) // Sueldo + Aumento: $10000
}
