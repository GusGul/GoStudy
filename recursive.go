package main

import "fmt"

func main() {
	// Recursão fatorial
	fmt.Println(Fatorial(5))
}

func Fatorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * Fatorial(x-1)
}
