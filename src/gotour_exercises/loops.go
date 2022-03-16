package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	diff := z

	// Keeps iterating until the difference is less than a billionth
	for math.Abs(diff) > 10e-9 {
		diff = (z*z - x) / (2 * z)
		z -= diff
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
