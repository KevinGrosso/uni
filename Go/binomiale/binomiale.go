package main

import (
    "fmt"
)

func Binomiale(n int) [][]int {
	risultato := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		risultato[i] = make([]int, i+1)
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				risultato[i][j] = 1
			} else {
				risultato[i][j] = risultato[i-1][j-1] + risultato[i-1][j]
			}
		}
	}
	return risultato
}

func main() {
	fmt.Println(Binomiale(10))
}
