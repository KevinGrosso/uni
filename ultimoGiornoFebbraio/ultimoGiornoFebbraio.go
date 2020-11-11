package main

import (
	"fmt"
	"strconv"
)

type data struct {
	giorno, mese, anno int
}

func main() {
	var anno int

	fmt.Print("Inserisci un anno: ")
	fmt.Scan(&anno)

	fmt.Println(dataToString(ultimoGFebbraio(anno)))
}

func ultimoGFebbraio(anno int) data {
	if (anno%4 == 0 && anno%100 != 0) || (anno%400 == 0) {
		return data{29, 2, anno}
	}
	return data{28, 2, anno}
}

func dataToString(d data) string {
	return strconv.Itoa(d.giorno) + "/" + strconv.Itoa(d.mese) + "/" + strconv.Itoa(d.anno)
}
