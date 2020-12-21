package main

import (
	"fmt"
)

func main() {
	var n, base int
	var conv string

	const cifre = "0123456789ABCDEF"

	fmt.Print("Inserisci un numero in base 10: ")
	fmt.Scan(&n)
	fmt.Print("In quale base lo vuoi convertire? (da 2 a 16) ")
	fmt.Scan(&base)

	if n == 0 {
		fmt.Print(0)
	}

	for n != 0 {
		conv = string(cifre[n%base]) + conv
		n /= base
	}

	fmt.Println(conv)
}
