package main

import (
	"fmt"
)

type data struct {
	giorno, mese, anno int
}

func main() {
	var d *data = new(data)

	fmt.Print("Inserisci una data gg/mm/aaaa: ")
	fmt.Scan(&(d.giorno))
	fmt.Scan(&(d.mese))
	fmt.Scan(&(d.anno))

	stampaData(d)
	resettaGiorno(d)
	stampaData(d)
}

func resettaGiorno(d *data) bool {
	if d == nil {
		return false
	}
	d.giorno = 1
	return true
}

func stampaData(d *data) {
	fmt.Printf("%02d/%02d/%04d\n", (*d).giorno,
		(*d).mese, (*d).anno)
}
