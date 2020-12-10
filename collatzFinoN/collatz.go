package main

import "fmt"

// dato n, stampa le sequenza di collatz

func main() {
	var x int

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&x)

	if x < 1 {
		fmt.Println("Inserire un numero >= 1")
		return
	}

	for i := 1; i <= x; i++ {
		n := i
		fmt.Print("#", n)

		for n != 1 {
			if n%2 == 0 {
				n /= 2
			} else {
				n = n*3 + 1
			}
			fmt.Print(" → ", n)
		}
		fmt.Println("\n")
	}
}
