package main

import "fmt"

// dato un numero intero, stabilisce se termina con la cifra 3

func main() {
	var num int
	
	fmt.Print("Inserisci un numero intero: ")
	fmt.Scan(&num)
	
	if num % 10 == 3 {
		fmt.Println("Il numero inserito termina con la cifra 3")
	} else {
		fmt.Println("Il numero inserito non termina con la cifra 3")
	}
}