package main

import "testing"

func ShouldFatorialCorrect(t *testing.T) {
	test := Fatorial(5)
	result := 120
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
