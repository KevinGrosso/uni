package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	numero := os.Args[1]
	combinazioni := generaCombinazioni(numero)

	for _, v := range combinazioni {
		n, _ := strconv.Atoi(v)
		if ÈPrimo(n) {
			fmt.Println(n)
		}
	}
}

func generaCombinazioni(n string) []string{
	numeri := []string{n}

	// 5899 -> 5899 - 589 589 599 899 - 58 59 99 - 5 9
	for i := 1; i <= 3; i++ {
		for j := 0; j+i <= len(n); j++ {
			numeri = append(numeri, n[:j] + n[j+i:])
		}
	}
	
	
	return numeri
}

func ÈPrimo(n int) bool {
	if n == 0 {
		return false
	}
	for i := 2; i <= n / 2; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}
