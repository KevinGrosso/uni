package main

import "fmt"

// Dato n, stampa i numeri primi <= n

func main() {
	var n, j int

	fmt.Print("Inserisci il numero massimo: ")
	fmt.Scan(&n)

	for i := 2; i <= n; i++ {
		for j = 2; j <= i/2; j++ {
			if i%j == 0 {
				break
			}
		}

		if j > i/2 {
			fmt.Println(i, "è primo")
		}
	}
}
