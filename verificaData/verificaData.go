package main 

import "fmt" 

// l'utente inserisce una data, il programma deve validarla

func main() {
	var g, m, a int
	var limiteG int
	
	fmt.Print("Inserisci una data gg/mm/aaaa: ")
	fmt.Scan(&g)
	fmt.Scan(&m)
	fmt.Scan(&a)
	
	if m > 0 && m <= 12 {
		if m == 11 || m == 4 || m == 6 || m == 9 {
			limiteG = 30
		} else if m == 2 {
			if (a%400 == 0) || (a%100 != 0 && a%4 == 0) {
				limiteG = 29
			} else {
				limiteG = 28
			}
		} else {
			limiteG = 31
		}
		if g > 0 && g <= limiteG {
			fmt.Println("Data valida")
		} else {
			fmt.Println("Data non valida")
		}
	} else {
		fmt.Println("Data non valida")
	}
}
