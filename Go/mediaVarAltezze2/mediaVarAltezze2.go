package main

import (
	"fmt"
	"math"
)

func main() {
	var alt []int
	var n int

	fmt.Print("Quante persone ci sono? ")
	fmt.Scan(&n)

	alt = make([]int, n)

	for i := range alt {
		fmt.Print(i+1, "# -> ")
		fmt.Scan(&alt[i])
	}

	// Media
	somma := 0
	for i := range alt {
		somma += alt[i]
	}
	media := float64(somma) / float64(n)

	// Varianza
	sommaScartiQuadratici := 0.0
	for _, v := range alt {
		sommaScartiQuadratici += (float64(v) - media) * (float64(v) - media)
	}

	varianza := sommaScartiQuadratici / float64(n)
	sqm := math.Sqrt(varianza)

	fmt.Printf("Media: %.2f\nVarianza: %.2f\nScarto quadratico medio: %.2f\n", media, varianza, sqm)
}
