package main

import "testing"

//Padrão triple A - AAA
//Arrange - Organizar (preparar)
//Act - Agir (rodar)
//Assert - Afirmar (verificar as asserções)
func ShouldFatorialCorrect(t *testing.T) {
	//Arrange
	test := Fatorial(5)
	//Act
	result := 120
	//Assert
	if test != result {
		t.Errorf("Fatorial(5) = %d; want %d", test, result)
	}
}

func ShouldFatorialIncorrect(t *testing.T) {
	test := Fatorial(5)
	result := 140
	if test != result {
		t.Errorf("Fatorial(5) = %d; want %d", test, result)
	}
}
