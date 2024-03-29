package main

import "fmt"

func main() {
	// Defer
	defer fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
}
