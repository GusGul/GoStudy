package main

import "testing"

func ShouldSumCorrect(t *testing.T) {
	//Arrange
	test := Sum(5, 5)
	//Act
	result := 10
	//Assert
	if test != result {
		t.Errorf("Sum(5, 5) = %f; want %f", test, result)
	}
}

func ShouldSumIncorrect(t *testing.T) {
	test := Sum(5, 5)
	result := 11
	if test != result {
		t.Errorf("Sum(5, 5) = %f; want %f", test, result)
	}
}

func ShouldSubtractCorrect(t *testing.T) {
	test := Subtract(5, 5)
	result := 0
	if test != result {
		t.Errorf("Subtract(5, 5) = %f; want %f", test, result)
	}
}

func ShouldSubtractIncorrect(t *testing.T) {
	test := Subtract(5, 5)
	result := 1
	if test != result {
		t.Errorf("Subtract(5, 5) = %f; want %f", test, result)
	}
}

func ShouldMultiplyCorrect(t *testing.T) {
	test := Multiply(5, 5)
	result := 25
	if test != result {
		t.Errorf("Multiply(5, 5) = %f; want %f", test, result)
	}
}

func ShouldMultiplyIncorrect(t *testing.T) {
	test := Multiply(5, 5)
	result := 24
	if test != result {
		t.Errorf("Multiply(5, 5) = %f; want %f", test, result)
	}
}

func ShouldDivideCorrect(t *testing.T) {
	test := Divide(5, 5)
	result := 1
	if test != result {
		t.Errorf("Divide(5, 5) = %f; want %f", test, result)
	}
}

func ShouldDivideIncorrect(t *testing.T) {
	test := Divide(5, 5)
	result := 2
	if test != result {
		t.Errorf("Divide(5, 5) = %f; want %f", test, result)
	}
}
