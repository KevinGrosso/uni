package main

// Dato un numero positivo, stampa la sequenza di collatz

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Inserisci un numero positivo: ")
	fmt.Scan(&n)

	if n > 0 {
		fmt.Println("Sequenza di collatz con n=", n)
		for prossimo := n; prossimo != 1; prossimo = prossimoCollatz(prossimo) {
			fmt.Println(prossimo)
		}
		fmt.Println(1)
	} else {
		fmt.Println("Il numero deve essere positivo")
	}
}

/* Dato un intero positivo, restituisce
   n/2      se n è pari
   3n + 1   se n è dispari
*/
func prossimoCollatz(n int) int {
	if n > 0 {
		if n%2 == 0 {
			return n / 2
		}
		return 3*n + 1
	}
	return 0
}
