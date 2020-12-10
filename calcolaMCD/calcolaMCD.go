package main 

import "fmt" 

// dati due numeri, calcola il loro MCD

func main() {
	var x, y int
	
	fmt.Print("Inserisci due numeri: ")
	fmt.Scan(&x, &y)
	
	for r := x % y; r != 0; r = x % y {
		x = y
		y = r
	}
	
	fmt.Println("Il massimo comun divisore è", y)
}
