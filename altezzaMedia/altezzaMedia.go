package main

import "fmt" 

// Stampa l'altezza media di un certo numero di persone

func main() {
	var quanti, tot int
	var media float64
	
	fmt.Print("Inserisci il numero di persone: ")
	fmt.Scan(&quanti)

	for i:=0; i < quanti; i++ {
		var altezza int

		fmt.Print("Inserisci la ", i+1, "^ altezza: ")
		fmt.Scan(&altezza)

		tot += altezza
	}

	media = float64(tot) / float64(quanti)
	fmt.Println("L'altezza media è", media)
}
