package main

import "fmt"

func main() {
	var n1, n2 float64

	fmt.Print("Inserisci due numeri: ")
	fmt.Scan(&n1, &n2)
	
	fmt.Println("La media dei due numeri è:", (n1 + n2) / 2)
}