package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, _ := strconv.Atoi(os.Args[1])

	for i := 0; i < n; i++ {
		if i == 0 {
			stampaPrimaRiga(n)
		} else {
			stampaRune('O', n - i)
			stampaRune(' ', 3 + ((i-1)*2))
			stampaRune('O', n - i)
			fmt.Println()
		}
	}
}

func stampaPrimaRiga(n int) {
	stampaRune('O', n)
	fmt.Print("+")
	stampaRune('O', n)
	fmt.Println()
}

func stampaRune(r rune, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(r))
	}
}
