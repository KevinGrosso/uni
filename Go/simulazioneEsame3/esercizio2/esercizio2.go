package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	N, _ := strconv.Atoi(os.Args[1])
	k, _ := strconv.Atoi(os.Args[2])

	numeriPari := FiltraNumeri(GeneraNumeri(N, k))

	for _, num := range numeriPari {
		fmt.Println(num)
	}
}

func GeneraNumeri(N, k int) []int {
	numeri := []int{}
	ns := strconv.Itoa(N)

	for i := 0; i+k-1 < len(ns); i++ {
		ni, _ := strconv.Atoi(ns[i:i+k])
		numeri = append(numeri, ni)
	}

	return numeri
}

func  FiltraNumeri(sl [] int) []int {
	pari := []int{}
	for _, n := range sl {
		if n & 1 == 0 {
			pari = append(pari, n)
		}
	}

	return pari
}
