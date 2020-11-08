package main

import (
	"fmt"
)

// stampa i primi n numeri di mersenne

func main() {
	var quanti int

	fmt.Print("Quanti numeri primi di mersenne vuoi? ")
	fmt.Scan(&quanti)

	for i, n := 0, uint64(2); i < quanti; n *= 2 {
		if isPrime(n - 1) {
			fmt.Println(n - 1)
			i++
		}
	}
}

func isPrime(n uint64) bool {
	if n > 1 {
		var i uint64
		for i = 2; i <= n/i; i++ {
			if n%i == 0 {
				break
			}
		}

		if i > n/i {
			return true
		}
		return false
	}
	return false
}
