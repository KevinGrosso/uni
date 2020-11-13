package main

import "fmt"

func main() {
	var x, y, q, r int

	fmt.Print("Inserisci dividendo e divisore: ")
	fmt.Scan(&x)
	fmt.Scan(&y)

	if dividi(x, y, &q, &r) {
		fmt.Printf("%d / %d = %d, resto: %d\n", x, y, q, r)
	} else {
		fmt.Println("Divisione per 0")
	}
}

func dividiRet(divis, divid int) (quoz int, resto int, ok bool) {
	if divid == 0 {
		return
	}

	quoz = divis / divid
	resto = divis % divid
	ok = true
	return
}

func dividi(divis int, divid int, quoz *int, resto *int) bool {
	if divid == 0 {
		return false
	}
	*quoz = divis / divid
	*resto = divis % divid
	return true
}
