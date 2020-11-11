package main

import (
	"fmt"
)

func main() {
	var n int
	var x, y = 1, 1

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)

	for ; n > 0; n-- {
		fmt.Println(x)
		x += y
		x, y = y, x
	}
}
