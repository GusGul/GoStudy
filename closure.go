package main

import "fmt"

func main() {
	x := 0

	// closure
	numero := func() int {
		x++
		return x
	}
	fmt.Println(numero())
}
