package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)
	fmt.Printf("multiply %d and %d: %d", x, y, multiply(x, y))
}

func multiply(x, y int) int {
	return x * y
}
