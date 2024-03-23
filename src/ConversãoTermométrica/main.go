package main

import "fmt"

const boilingPointK = 373.15

func main() {
	celsius := boilingPointK - 273.15
	fahrenheit := celsius*9/5 + 32

	fmt.Printf("Ponto de ebulição da água em Kelvin: %.2f\n", boilingPointK)
	fmt.Printf("Ponto de ebulição da água em Celsius: %.2f\n", celsius)
	fmt.Printf("Ponto de ebulição da água em Fahrenheit: %.2f\n", fahrenheit)
}

