package main

import (
    "fmt"
)

func removeDuplicate(s []string) []string {
	m := make(map[string]bool)
	var risultato []string

	for _, valore := range s {
		if !m[valore] {
			risultato = append(risultato, valore)
			m[valore] = true
		}
	}
	return risultato
}

func main() {
	s := []string{"elefante", "cane", "gatto", "topo", "gatto", "tartaruga", "topo", "elefante"}
	fmt.Println(removeDuplicate(s))
}
