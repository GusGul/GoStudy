package main

import "fmt"

// Exemplos de uso de goroutines
func loop(size int) {
	for x := 0; x < size; x++ {
		println(x)
	}
}

func main() {
	// Exemplos de uso de goroutines
	go loop(2)
	loop(4)
	var escreva string
	fmt.Scanln(&escreva)
}
