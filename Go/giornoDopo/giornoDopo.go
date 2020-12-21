package main

import "fmt"

type data struct {
	g, m, a int
}

func main() {
	var g, m, a int

	fmt.Print("Inserisci una data gg/mm/aaaa: ")
	fmt.Scan(&g)
	fmt.Scan(&m)
	fmt.Scan(&a)

	d1 := data{g, m, a}
	fmt.Println("Giorno:", stringData(d1))

	d2 := retGiornoDopo(d1)
	fmt.Println("Giorno + 1:", stringData(d2))

	giornoDopo(&d2)
	fmt.Println("Giorno + 2:", stringData(d2))
}

func retGiornoDopo(d data) data {
	if d.g == giorni(d.m, d.a) {
		d.g = 1
		if d.m == 12 {
			d.m = 1
			d.a++
		} else {
			d.m++
		}
	} else {
		d.g++
	}
	return d
}

func giornoDopo(d *data) {
	if d.g == giorni(d.m, d.a) {
		d.g = 1
		if d.m == 12 {
			d.m = 1
			d.a++
		} else {
			d.m++
		}
	} else {
		d.g++
	}
}

func giorni(m, a int) int {
	switch m {
	case 11, 4, 6, 9:
		return 30
	case 2:
		if (a%400 == 0) || (a%4 == 0 && a%100 != 0) {
			return 29
		}
		return 28
	default:
		return 31
	}
}

func stringData(d data) string {
	return fmt.Sprintf("%02d/%02d/%04d", d.g, d.m, d.a)
}
