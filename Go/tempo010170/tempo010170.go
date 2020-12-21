package main

import "fmt"

// si legge una data gg/mm/aaaa e si calcola quanti giorni sono passati dal 01/01/1970

func main() {
	var g, m, a, anno, mese, giorni int

	fmt.Print("Inserisci una data (gg/mm/aaaa): ")
	fmt.Scan(&g)
	fmt.Scan(&m)
	fmt.Scan(&a)
	
	if a >= 1970 {
		anno = 1970
		giorni += g
	} else {
		anno, a = a, anno
		mese = m
		m = 12
		
		if mese == 11 || mese == 4 || mese == 6 || mese == 9 {
			giorni += 31 - g
		} else if mese == 2{
			if (anno%400 == 0) || (anno%4 == 0 && anno%100 != 0) {
				giorni += 30 - g	
			} else {
				giorni += 29 - g
			}
		} else {
			giorni += 32 - g
		}
	}
	
	for ; anno < a; anno++ {
		giorni += 365
		if anno%400 == 0 {
			giorni++
		} else if anno%4 == 0 && anno%100 != 0 {
			giorni++
		}
	}

	for ; mese < m; mese++ {
		if mese == 11 || mese == 4 || mese == 6 || mese == 9 {
			giorni += 30
		} else if mese == 2 {
			giorni += 28
			if a%400 == 0 {
				giorni++
			} else if a%4 == 0 && a%100 != 0 {
				giorni++
			}
		} else {
			giorni += 31
		}
	}
	
	if anno < 1970 {
		giorni = -giorni
	}
	
	fmt.Println("Sono passati", giorni, "giorni dal 01/01/1970")
}
