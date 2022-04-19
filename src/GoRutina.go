package main

// Go rutinas crea concurrencia, manjear diversas cosas al mismo tiempo
// Concurrencia != paralelismo

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup // variable para inidicar que se debe esperar a las n goRutinas
// cuantas debe esperar para terminar la ejecucion

func imprimirCantidad(etiqueta string) {
	// se indica que la rutina ya termino
	// se restara 1 de las adicionadas
	defer wg.Done()

	for cantidad := 1; cantidad <= 10; cantidad++ {
		sleep := rand.Intn(1000)                                // devuelve un numero random de 0 a 1000
		time.Sleep(time.Duration(sleep) * time.Millisecond)     // espera que pasen n milesegundos
		fmt.Printf("Cantidad: %d  de %s\n", cantidad, etiqueta) // Imprime la candidad

	}
}

func main() {
	wg.Add(2) // adicion de 2 Gorutinas para que finalice
	fmt.Println("Iniciamos las gorutinas...")

	// con go se lanzan las rutinas
	go imprimirCantidad("A")

	go imprimirCantidad("B")

	// Espera que se finalicen las 2 gorutinas
	fmt.Println("Esperando que finalicen..")
	wg.Wait() // no va a seguir hasta que se termine las 2 gorutinas
	fmt.Println("Terminando el Programa")

	// 	EJECUCION

	// Iniciamos las gorutinas...
	// Esperando que finalicen..
	// Cantidad: 1  de A
	// Cantidad: 1  de B
	// Cantidad: 2  de A
	// Cantidad: 2  de B
	// Cantidad: 3  de A
	// Cantidad: 3  de B
	// Cantidad: 4  de A
	// Cantidad: 4  de B
	// Cantidad: 5  de A
	// Cantidad: 5  de B
	// Cantidad: 6  de A
	// Cantidad: 6  de B
	// Cantidad: 7  de B
	// Cantidad: 7  de A
	// Cantidad: 8  de A
	// Cantidad: 9  de A
	// Cantidad: 8  de B
	// Cantidad: 9  de B
	// Cantidad: 10  de A
	// Cantidad: 10  de B
	// Terminando el Programa
}
