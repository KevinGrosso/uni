package main

import "fmt"

func main() {
	var scelta int
	var valore uint64

	fmt.Print("Scegli la conversione:\n" +
		"1)	secondi	-> ore\n2)	secondi	-> minuti\n3) 	minuti	-> ore \n" +
		"4)	minuti	-> secondi\n5)	ore	-> secondi\n6)	ore	-> minuti\n" +
		"7)	minuti	-> giorni e ore\n8)	minuti	-> anni e giorni\n" +
		"9)	esci dal programma\n: ")

	for fmt.Scan(&scelta); scelta != 9; fmt.Scan(&scelta) {
		if scelta < 1 || scelta > 9 {
			fmt.Print("L'opzione ", scelta, " non è disponibile\n: ")
		} else {
			fmt.Print("Inserisci il valore da convertire: ")
			fmt.Scan(&valore)

			if scelta == 1 {
				fmt.Println(valore, "secondi corrispondono a", float64(valore)/60/60, "ore")
			} else if scelta == 2 {
				fmt.Println(valore, "secondi corrispondono a", float64(valore)/60, "minuti")
			} else if scelta == 3 {
				fmt.Println(valore, "minuti corrispondono a", float64(valore)/60, "ore")
			} else if scelta == 4 {
				fmt.Println(valore, "minuti corrispondono a", float64(valore)*60, "secondi")
			} else if scelta == 5 {
				fmt.Println(valore, "ore corrispondono a", float64(valore)*60*60, "secondi")
			} else if scelta == 6 {
				fmt.Println(valore, "ore corrispondono a", float64(valore)*60, "minuti")
			} else if scelta == 7 {
/*				per calcore in modo esatto
				giorni := (valore*60) / (24*60*60)
				ore := ((valore*60 - giorni*24*60*60)) / (60*60)
				minuti := (valore*60 - giorni*24*60*60 - ore*60*60) / 60

				if minuti < 10 {
					fmt.Print(valore, " minuti corrispondono a ", giorni, " giorni e ", ore, ".0", minuti, " ore\n")
				} else {
					fmt.Print(valore, " minuti corrispondono a ", giorni, " giorni e ", ore, ".", minuti, " ore\n")
				}
*/
			}
			fmt.Print("\n\nScegli la conversione:\n" +
			 "1)	secondi	-> ore\n2)	secondi	-> minuti\n3) 	minuti	-> ore \n" +
			 "4)	minuti	-> secondi\n5)	ore	-> secondi\n6)	ore	-> minuti\n" +
			 "7)	minuti	-> giorni e ore\n8)	minuti	-> anni e giorni\n" +
		 	 "9)	esci dal programma\n: ")
		}
	}
}
