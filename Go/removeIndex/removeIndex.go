package main

import (
	"fmt"
)

func removeIndex(s []int, i int) []int {
	var ris []int

	ris = append(ris, s[:i]...)
	ris = append(ris, s[i+1:]...)

	return ris
}

func removeIndexes(slice, indici []int) []int {
	ris := make([]int, len(slice))
	indiciNuovi := make([]int, len(indici))

	copy(ris, slice)
	copy(indiciNuovi, indici)

	for i := 0; i < len(indiciNuovi); i++ {
		ris = append(ris[:indiciNuovi[i]], ris[indiciNuovi[i]+1:]...)

		for j := 0; j < len(indiciNuovi); j++ {
			if indiciNuovi[j] > indiciNuovi[i] {
				indiciNuovi[j]--
			}
		}
	}
	return ris
}

func main() {
	slice := []int{10, 20, 30, 40, 50, 60}
	indici := []int{4, 1, 2}
	fmt.Println(removeIndexes(slice, indici))
}
