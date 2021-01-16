package main

import (
	"fmt"
)

func factorial(n uint64) uint64{
	result := uint64(1)
	for ; n>1; n-- {
		result *= n
	}
	return result
}

func main() {
	e := 1.0
	for i := 1; i < 40; i++ {
		e += 1.0 / float64(factorial(uint64(i)))
	}
	fmt.Printf("Approximation of the number e: %1.100f", e)
}
