package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	numeroCarteSeme = 13
	numeroCarte     = numeroCarteSeme * 4
)

/* Un set è una sequenza di carte distinte */
type set []int

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	m := mazzoNuovo(true)

	// Stampo la mano
	for _, c := range m[:5] {
		fmt.Printf("%s\n", strCarta(c))
	}
}

func strCarta(carta int) string {
	seme := [...]string{"cuori", "quadri", "fiori", "picche"}
	valore := [...]string{"A", "2", "3", "4", "5", "6", "7", "8",
		"9", "10", "J", "Q", "K"}

	return valore[carta%numeroCarteSeme] + " di " +
		seme[carta/numeroCarteSeme]
}

/* Restituisce un mazzo. Se mescolato=true, lo mescola */
func mazzoNuovo(mescolato bool) (s set) {
	s = make(set, numeroCarte)

	for i := 0; i < numeroCarte; i++ {
		s[i] = i
	}

	if mescolato {
		for i := numeroCarte; i > 0; i-- {
			pos := rand.Intn(i)
			s[pos], s[i-1] = s[i-1], s[pos]
		}
	}
	return
}

/* Data una mano di poker (un set di 5 carte) stabilisce (true) se c'è una coppia */
func coppia()
