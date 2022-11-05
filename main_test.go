package main

import "testing"

func TestMultiple(t *testing.T) {
	expected := 2 * 3
	got := multiple(2, 3)

	if expected != got {
		t.Fatalf("Expected %d, Got: %d", expected, got)
	}
}
