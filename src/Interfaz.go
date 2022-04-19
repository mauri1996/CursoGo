package main

import "fmt"

// Creacion de type
type Persona struct {
	nombre string
	email  string
	edad   int
}

func (p Persona) Nombre() string {
	return p.nombre
}

func (p Persona) Email() string {
	return p.email
}

type Adminsitrador struct {
	Persona
	cargo string
}

func (a Adminsitrador) AbrirForo() {
	fmt.Println("Foro Abierto")
}

type Manager struct {
	Persona
	puesto string
}

func (m Manager) AbrirTarea() {
	fmt.Println("Tarea Abierto")
}

// +++++ Redundancia de CODIGO ++++++
// func PresentarseA(a Adminsitrador) {
// 	fmt.Println("Nombre:", a.Nombre())
// 	fmt.Println("Email:", a.Email())
// }

// func PresentarseM(m Manager) {
// 	fmt.Println("Nombre:", m.Nombre())
// 	fmt.Println("Email:", m.Email())
// }

// Antigua funcion
// func Presentarse(p Persona) {
// 	fmt.Println("Nombre:", p.Nombre())
// 	fmt.Println("Email:", p.Email())
// }
////

// Creacion de Interfaz para que sea una especie de generalizacion
// Se debe agregar los metodos de los type que contendra la interfaz
// Nombre() es metodo del type Persona y esta devuelve un String
// Go automaticamente detecta y hace que ese type implemente esa interfaz
// Ahora Administrador y Manager pueden ser tratados como Usuarios, puesto que estos tienen los metodos
// intrinsicamente devido a que tienen un type tipo Persona dentro.
type Usuario interface {
	// se debe colocar aqui todos los metodos que podran compartir en la interfaz
	Nombre() string
	Email() string
}

func Presentarse(p Usuario) {
	fmt.Println("Nombre:", p.Nombre())
	fmt.Println("Email:", p.Email())
}

func main() {
	mauri := Persona{"Mauri", "aa@g.com", 10}
	Presentarse(mauri) // Nombre: Mauri
	// Email: aa@g.com

	admin1 := Adminsitrador{Persona{"daysi", "d@.com", 20}, "admin"}
	manage1 := Manager{Persona{"Caro", "c@.com", 20}, "caro"}

	// +++++ Redundancia de CODIGO ++++++
	// PresentarseM(manage1)
	// // Nombre: Caro
	// // Email: c@.com
	// PresentarseA(admin1)
	// // Nombre: daysi
	// // Email: d@.com

	// Ahora pueden ser tratados como interfaces Adminsitrador y Manager
	// se usa la misma funcion para todos los tipos
	Presentarse(admin1)
	// Nombre: daysi
	// Email: d@.com
	Presentarse(manage1)
	// Nombre: Caro
	// Email: c@.com

	// se puede valores a una interfaz
	var i Usuario

	i = mauri
	fmt.Println("i:", i)          // i: {Mauri aa@g.com 10}
	fmt.Println("i:", i.Nombre()) // i: Mauri

	i = admin1
	fmt.Println("i:", i)          // i: {Mauri aa@g.com 10}
	fmt.Println("i:", i.Nombre()) // i: daysi

	// NO se puede hacer
	// la interaz solo tiene acceso a los metodos Nombre() y Email()
	// i.AbrirForo()
}
