package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	numeroCarteSeme  = 13
	numeroCarteMazzo = numeroCarteSeme * 4
	numeroCarteMano  = 5
)

type mazzo [numeroCarteMazzo]int
type set []int

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	m := mazzoNuovo(true)
	mano := m[:numeroCarteMano]

	//Stampo la mano
	for _, c := range mano {
		fmt.Printf("%s\n", strCarta(c))
	}

	if coppia(mano) {
		fmt.Println("Coppia!")
	}
}

func strCarta(carta int) string {
	seme := [...]string{"cuori", "quadri", "fiori", "picche"}
	valore := [...]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	return valore[carta%numeroCarteSeme] + " di " + seme[carta/numeroCarteSeme]
}

/* Restituisce un mazzo. Se mescolato=true, lo mescola */
func mazzoNuovo(mescolato bool) (s mazzo) {
	for i := 0; i < numeroCarteMazzo; i++ {
		s[i] = i
	}

	if mescolato {
		for i := numeroCarteMazzo; i > 0; i-- {
			pos := rand.Intn(i)
			s[pos], s[i-1] = s[i-1], s[pos]
		}
	}
	return
}

/* Dato un set di carte, stabilisce (true) se c'è almeno una coppia */
func coppia(mano set) bool {
	for i := 0; i < numeroCarteMano-1; i++ {
		for j := 1; i+j < numeroCarteMano; j++ {
			if mano[i]%numeroCarteSeme == mano[i+j]%numeroCarteSeme {
				return true
			}
		}
	}
	return false
}
