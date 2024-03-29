package main

import "testing"

func TestFatorial(t *testing.T) {
	test := Fatorial(5)
	result := 120
	if test != result {
		t.Errorf("Fatorial(5) = %d; want %d", test, result)
	}
}

func TestFatorial2(t *testing.T) {
	test := Fatorial(5)
	result := 140
	if test != result {
		t.Errorf("Fatorial(5) = %d; want %d", test, result)
	}
}
