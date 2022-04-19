package main

import (
	"fmt"
	"time"
)

func main() {
	go animation(100 * time.Millisecond) // invoca la animacion infinitamente y termianra cuando se acabe el proceso
	const n = 45
	resultado := fib(n) // no es necesario el sync.WaitGroup, porque fib() funge como wait y cuado trermine se acabara el programa
	fmt.Printf("\rFibonacci(%d) = %d\n", n, resultado)
}

func animation(retraso time.Duration) {
	for {
		for _, r := range `\|/-` { // forEach
			fmt.Printf("\r%c", r)
			time.Sleep(retraso)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
