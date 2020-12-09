package main 

import "fmt" 

/* 
	legge una serie arbitraria di valori interi maggiori di 0 e per ciascuno di essi
	determina se è una potenza di 2. Il programma termina quando l'utente inserisce un
	valore minore o uguale a 0
*/

func main() {
	var n int
	
	fmt.Print("N: ")
	
	for fmt.Scan(&n); n > 0; fmt.Scan(&n) {
		/*
			0100 & 0011 = 0  
			1100 & 1011 = 1
			0010 & 0001 = 0
			1010 & 1001 = 1
		*/
		if n & (n - 1) == 0 {
			fmt.Println("è una potenza di 2")
		}	
		fmt.Print("N: ")
	}
	
	fmt.Println("Programma terminato")
}
