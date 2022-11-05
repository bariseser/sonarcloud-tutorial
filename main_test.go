package main

import "testing"

func TestMultiply(t *testing.T) {
	expected := 2 * 3
	got := multiply(2, 3)

	if expected != got {
		t.Fatalf("Expected %d, Got: %d", expected, got)
	}
}
