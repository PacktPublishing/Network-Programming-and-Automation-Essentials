package main

import (
	"fmt"
	"math"
)

func squareRoot(value float64) float64 {
	if value < 0 {
		panic("negative values are not allowed")
	}

	return math.Sqrt(value)
}

func main() {
	fmt.Println(squareRoot(-2))

	fmt.Println("done")
}
