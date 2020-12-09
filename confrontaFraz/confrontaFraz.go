package main

import "fmt"

// Legge due frazioni e le confronta

func main() {
	var n1, d1, n2, d2, n3 int
	
	fmt.Print("Inserire il primo numeratore e denominatore: ")
	fmt.Scan(&n1)
    fmt.Scan(&d1)
	
	fmt.Print("Inserire il secondo numeratore e denominatore: ")
    fmt.Scan(&n2)
    fmt.Scan(&d2)
	
	n3 = n1*d2 - n2*d1
	
	if n3 > 0 {
		fmt.Println("La prima frazione è maggiore della seconda")
	} else if n3 < 0 {
		fmt.Println("La seconda frazione è maggiore della prima")
	} else {
		fmt.Println("Le frazioni sono uguali")
	}	
}