package main

import (
	"fmt"
	"math"
)

func main() {
	var alt [1000]int
	var n int

	fmt.Print("Quante persone ci sono? ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print(i+1, "# -> ")
		fmt.Scan(&alt[i])
	}

	// Media
	somma := 0
	for i := 0; i < n; i++ {
		somma += alt[i]
	}
	media := float64(somma) / float64(n)

	// Varianza
	sommaScartiQuadratici := 0.0
	for i := 0; i < n; i++ {
		sommaScartiQuadratici += (float64(alt[i]) - media) * (float64(alt[i]) - media)
	}

	varianza := sommaScartiQuadratici / float64(n)
	sqm := math.Sqrt(varianza)

	fmt.Printf("Media: %.2f\nVarianza: %.2f\nScarto quadratico medio: %.2f\n", media, varianza, sqm)
}
