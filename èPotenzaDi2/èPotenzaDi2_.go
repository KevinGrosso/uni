package main 

import "fmt" 

/* 
	legge una serie arbitraria di valori interi maggiori di 0 e per ciascuno di essi
	determina se è una potenza di 2. Il programma termina quando l'utente inserisce un
	valore minore o uguale a 0
*/

func main() {
	var n, i int
	
	fmt.Print("N: ")
	
	for fmt.Scan(&n); n > 0; fmt.Scan(&n) {
		for i = 1; i < n; i *= 2{
		}
		
		if i == n {
			fmt.Println("è una potenza di 2")
		}
		fmt.Print("N: ")
	}
	
	fmt.Println("Programma terminato")
}
