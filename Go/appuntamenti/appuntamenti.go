package main 

import "fmt" 

/*
	Verifica se due appuntamenti giornalieri si sovrappongono.
	Ogni appuntamento ha un orario di inizio e di fine. 
	Un orario è una coppia hh mm
*/

func main() {
	var inizioH1, inizioM1, fineH1, fineM1 int
	var inizioH2, inizioM2, fineH2, fineM2 int
	
	fmt.Print("Inserisci l'orario di inizio del primo appuntamento (hh mm): ")
	fmt.Scan(&inizioH1)
	fmt.Scan(&inizioM1)
	fmt.Print("Inserisci l'orario di fine del primo appuntamento (hh mm): ")
	fmt.Scan(&fineH1)
	fmt.Scan(&fineM1)

	fmt.Print("Inserisci l'orario di inizio del secondo appuntamento (hh mm): ")
	fmt.Scan(&inizioH2)
	fmt.Scan(&inizioM2)
	fmt.Print("Inserisci l'orario di fine del secondo appuntamento (hh mm): ")
	fmt.Scan(&fineH2)
	fmt.Scan(&fineM2)
	
	var inizio1, fine1 int = inizioH1 * 100 + inizioM1, fineH1 * 100 + fineM1
	var inizio2, fine2 int = inizioH2 * 100 + inizioM2, fineH2 * 100 + fineM2
	
	if inizio1 > inizio2 && inizio1 < fine2 {
		fmt.Println("coincidono")
	}
	if inizio2 > inizio1 && inizio2 < fine1 {
		fmt.Println("coincidono")
	}
}
