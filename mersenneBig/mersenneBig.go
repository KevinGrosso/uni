package main

import (
	"fmt"
	"math/big"
)

// stampa i primi n numeri di mersenne

func main() {
	var quanti int

	fmt.Print("Quanti numeri primi di mersenne vuoi? ")
	fmt.Scan(&quanti)

	uno := big.NewInt(1)
	due := big.NewInt(2)
	x := big.NewInt(1)
	n := big.NewInt(2)

	for i := 0; i < quanti; {
		x.Sub(n, uno)

		if isPrime(x) {
			fmt.Println(x)
			i++
		}

		n.Mul(n, due)
	}
}

func isPrime(n *big.Int) bool {

	if n.Cmp(big.NewInt(1)) == 1 {
		i := big.NewInt(2)

		d := big.NewInt(1)
		d.Div(n, i)
		z := big.NewInt(1)

		for d := i; i.Cmp(d) == -1 || i.Cmp(d) == 0; i.Add(i, big.NewInt(1)) {
			//z.Mod(n, i)
			z.Div(n, i)
			z.Mul(z, i)

			if z.Cmp(n) == 0 {
				break
			}
		}

		d.Div(n, i)
		if i.Cmp(d) == 1 {
			return true
		}
		return false
	}
	return false
}
