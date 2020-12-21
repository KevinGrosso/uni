package main

import "fmt"

// stampa la parte frazionaria di un numero inserito

func main() {
	var num float64
	
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&num)
	
	fmt.Println("La parte frazionaria è:", num - float64(int(num)))
}