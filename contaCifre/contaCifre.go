package main 

import "fmt" 

// dato un numero intero, stampa quante cifre ha

func main() {
	var num int
	var cifre int = 1
	
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&num)
	
	for num/=10; num != 0; num/=10 {
		cifre++
	}
	
	fmt.Println("Il numero ha", cifre, "cifre")
}
