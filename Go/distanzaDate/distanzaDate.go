package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type data struct {
	g, m, a int
}

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Sintassi: %s gg/mm/aaaa gg/mm/aaaa\n", os.Args[0])
		return
	}

	d1, ok1 := stringToData(os.Args[1])
	d2, ok2 := stringToData(os.Args[2])

	if !ok1 || !ok2 {
		fmt.Println("Errore: date non valide")
		return
	}

	fmt.Println(distanzaDate(d1, d2))
}

/* Funzione di debug*/
func stampaData(d data) {
	fmt.Printf("%02d/%02d/%04d\n", d.g, d.m, d.a)
}

func stringToData(ds string) (data, bool) {
	var d data

	das := strings.Split(ds, "/")

	if len(das) != 3 {
		return d, false
	}

	var dai [3]int
	for i := 0; i < 3; i++ {
		var err error
		dai[i], err = strconv.Atoi(das[i])

		if err != nil {
			return d, false
		}
	}

	d.g = dai[0]
	d.m = dai[1]
	d.a = dai[2]

	if d.m <= 0 || d.m > 12 {
		return d, false
	}

	if d.g <= 0 || d.g > giorniMese(d.m, d.a) {
		return d, false
	}

	return d, true
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

func distanzaDate(d1, d2 data) int {
	dist := giorniDaEpoca(d1) - giorniDaEpoca(d2)

	if dist < 0 {
		dist *= -1
	}

	return dist
}
