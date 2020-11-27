package main

import (
	"fmt"
	"os"
)

func permutazioni(r []rune) []string {
	if len(r) == 0 {
		return []string{""}
	}

	var risultato []string

	for i := 0; i < len(r); i++ {
		resto := []rune(string(r[:i]) + string(r[i+1:]))
		lista := permutazioni(resto)

		for _, x := range lista {
			risultato = append(risultato, string(r[i])+x)
		}
	}
	return risultato
}

func main() {
	s := os.Args[1]
	r := []rune(s)

	lista := anagrammi(r)
	for _, x := range lista {
		fmt.Println(x)
	}
}
