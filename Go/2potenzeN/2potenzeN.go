package main

import "fmt"

// stampa le potenze di 2 <= n

func main() {
	var num int
	
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&num)

	fmt.Println("Potenze di 2 <= di", num)

	for pow := 1; pow <= num; pow *= 2 {
		fmt.Println(pow)
	}
}
