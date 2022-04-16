package main

import (
	"fmt"
	"time"
)

func getValue(nombre string) int {
	if nombre == "Mauri" {
		return 1
	}
	return 0
}

func verificarEdad(edad int, nombre string) string {
	if value := getValue(nombre); value == 1 && edad >= 18 {
		return nombre + " es mayor de edad"
	} else if value := getValue(nombre); value == 0 && edad >= 18 {
		return nombre + " no es Mauri y es mayor de edad"
	} else {
		return nombre + " no es mayor de edad"
	}
}
func saludo() {
	t := time.Now()
	fmt.Println("Hora:", t) // 2022-04-16 17:04:31.3976784 -0500 -05 m=+0.002642901
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func main() {
	fmt.Println("20:", verificarEdad(20, "Mauri"))  //  Mauri es mayor de edad
	fmt.Println("20:", verificarEdad(20, "Mauri1")) //  Mauri1 no es Mauri y es mayor de edad
	saludo()                                        //Good evening.
}
