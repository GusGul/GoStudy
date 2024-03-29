package main

import (
	"fmt"
)

func main() {
	fmt.Println("===== Calculator =====")

	fmt.Println("Type your first number:")
	var firstNumber int
	fmt.Scanln(&firstNumber)
	fmt.Println("Type your operator: (+, -, *, /)")
	var operator string
	fmt.Scanln(&operator)
	fmt.Println("Type your second number:")
	var secondNumber int
	fmt.Scanln(&secondNumber)

	var result int
	switch operator {
	case "+":
		result = Sum(firstNumber, secondNumber)
	case "-":
		result = Subtract(firstNumber, secondNumber)
	case "*":
		result = Multiply(firstNumber, secondNumber)
	case "/":
		result = Divide(firstNumber, secondNumber)
	default:
		fmt.Println("Invalid operator")
	}

	fmt.Println("Result: ", result)
}

func Sum(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) int {
	if b == 0 {
		fmt.Println("Invalid operation: Division by zero")
		return 0
	}
	return a / b
}
