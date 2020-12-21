package main

import "fmt"

// dato un numero, lo arrotonda

func main() {
	var num float64
	
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&num)
	
	fmt.Println("Il numero arrotondato è:", int(num + 0.5))
}