package main

import "fmt"

func main() {
	fmt.Println("--------------------")
	multipleByThree()
	fmt.Println("--------------------")
	pinPan()
}

// Parte 1
func multipleByThree() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

// Parte 2
func pinPan() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("PinPan")
		} else if i%3 == 0 {
			fmt.Println("Pin")
		} else if i%5 == 0 {
			fmt.Println("Pan")
		} else {
			fmt.Println(i)
		}
	}
}
