package main

import "fmt"

// Legge due frazioni e ne calcola la somma (non semplificata ai minimi termini)

func main() {
	var n1, d1, n2, d2, n3, d3 int
	
	fmt.Print("Inserire il primo numeratore e denominatore: ")
	fmt.Scan(&n1)
    fmt.Scan(&d1)
	
	fmt.Print("Inserire il secondo numeratore e denominatore: ")
    fmt.Scan(&n2)
    fmt.Scan(&d2)
	
	n3 = n1*d2 + n2*d1
	d3 = d1 * d2
	
	fmt.Println("La somma delle due frazioni è", n3, "/", d3)
}