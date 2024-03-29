package main

import "testing"

func TestFatorial(t *testing.T) {
	result := Fatorial(5)
	if result != 120 {
		t.Errorf("Fatorial(5) = %d; want 120", result)
	}
}

func TestFatorial2(t *testing.T) {
	result := Fatorial(5)
	if result != 140 {
		t.Errorf("Fatorial(5) = %d; want 120", result)
	}
}
