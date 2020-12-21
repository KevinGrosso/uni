package main

import "fmt"

func main() {
	var n1, n2, n3 float64

	fmt.Print("Inserisci tre numeri: ")
	fmt.Scan(&n1, &n2, &n3)
	
	fmt.Println("La media dei tre numeri è:", (n1 + n2 + n3) / 3)
}