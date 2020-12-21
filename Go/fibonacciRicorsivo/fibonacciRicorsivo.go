package main

//Stampa l'n-esimo valore della successione di fibonacci

import (
	"fmt"
)

var cache []uint64

func fibonacciN(n uint64) uint64 {
	// fn = fn-1 + fn-2 per n >= 2
	// 0 1 1 2 3 5 8 13 21 34

	if n < 2 {
		return n
	}

	var f1, f2 uint64

	if cache[n-1] != 0 {
		f1 = cache[n-1]
	} else {
		f1 = fibonacciN(n - 1)
		cache[n-1] = f1
	}

	if cache[n-2] != 0 {
		f2 = cache[n-2]
	} else {
		f2 = fibonacciN(n - 2)
		cache[n-2] = f2
	}

	return f1 + f2
}

func limite(successione []uint64) float64 {
	var tot float64

	for i := 1; i < len(successione)-1; i++ { // parto da 1 per evitare divisioni per 0
		tot += float64(successione[i+1]) / float64(successione[i])
		fmt.Println(float64(successione[i+1]) / float64(successione[i]))
	}
	return tot / float64(len(successione)-2) // -2 perché parto da i=1 e arrivo fino a len-1
}

func main() {
	var n uint64

	fmt.Println("Quale valore della successione di fibonacci vuoi visualizzare?")
	fmt.Print("(massimo 93 per avere un valore corretto) -> ")
	fmt.Scan(&n)

	cache = make([]uint64, n) // alloco spazio per la cache

	fmt.Printf("Il %d° valore é %d\n", n, fibonacciN(n))
	fmt.Println("Cache:", cache)
	fmt.Println("Limite di f(n+1)/f(n) per n che tende a +inf ~=", limite(cache[n/2:]))
}
