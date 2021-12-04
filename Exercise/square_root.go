package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	// calculating sqr root
	for c := 0; c < 10; c++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
