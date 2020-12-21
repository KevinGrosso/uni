package main

import "fmt"

type data struct {
	g, m, a int
}

func giorniMese(m int, a int) int {
	switch m {
	case 11, 4, 6, 9:
		return 30
	case 2:
		if éBisestile(a) {
			return 29
		}
		return 28
	default:
		return 31
	}
}

func éBisestile(a int) bool {
	if a%400 == 0 || a%4 == 0 && a%100 != 0 {
		return true
	}

	return false
}

func giorniDaEpoca(d data) int {
	giorni := 0

	for i := 1970; i < d.a; i++ {
		if éBisestile(i) {
			giorni += 366
		} else {
			giorni += 365
		}
	}

	for i := 1; i < d.m; i++ {
		giorni += giorniMese(i, d.a)
	}

	giorni += d.g - 1

	return giorni
}

/* data una data, restituisce una stringa che rappresenta
   il giorno (es. 18/11/2020 -> mercoledì) */
func dataInGiorno(d data) string {
	// 1 gennaio 1970 giovedì
	settimana := [7]string{"giovedì", "venerdì", "sabato", "domenica", "lunedì", "martedì", "mercoledì"}

	giorni := giorniDaEpoca(d)
	return settimana[giorni%7]
}

func main() {
	d := data{3, 7, 1984}
	fmt.Println(dataInGiorno(d))
}
