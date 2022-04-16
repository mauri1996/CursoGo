package main

import "fmt"

func main() {

	// mapa -> diccionarios -> uso de clave y valor
	dias := make(map[int]string)
	fmt.Println(dias) // map[]

	// Agregar Datos
	dias[1] = "Domingo"
	dias[2] = "Lunes"
	dias[3] = "Martes"
	dias[4] = "Miercoles"
	dias[5] = "Jueves"
	dias[6] = "Viernes"
	dias[7] = "Sabado"

	fmt.Println(dias) // map[1:Domingo 2:Lunes 3:Martes 4:Miercoles 5:Jueves 6:Viernes 7:Sabado]

	dias[100] = "test"
	fmt.Println(dias) // map[1:Domingo 2:Lunes 3:Martes 4:Miercoles 5:Jueves 6:Viernes 7:Sabado 100:test]

	// Modificacion
	dias[7] = "SABADO"
	fmt.Println(dias) // map[1:Domingo 2:Lunes 3:Martes 4:Miercoles 5:Jueves 6:Viernes 7:SABADO 100:test]

	// Eliminar
	delete(dias, 100)
	fmt.Println(dias) // map[1:Domingo 2:Lunes 3:Martes 4:Miercoles 5:Jueves 6:Viernes 7:SABADO]

	fmt.Println(dias, len(dias)) // map[1:Domingo 2:Lunes 3:Martes 4:Miercoles 5:Jueves 6:Viernes 7:SABADO] 7

	// MAPAS COMPLEJOS
	estudiantes := make(map[string][]int) // clave string y valor arreglo de int

	estudiantes["Mauri"] = []int{13, 15, 16}
	estudiantes["Daysi"] = []int{14, 17, 18}

	fmt.Println(estudiantes) // map[Daysi:[14 17 18] Mauri:[13 15 16]]

	// Acceder a un elemento
	fmt.Println(estudiantes["Mauri"])    // [13 15 16]
	fmt.Println(estudiantes["Mauri"][1]) // 15
}
