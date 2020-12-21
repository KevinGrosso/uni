package main

import (
    "fmt"
)

// Stampa il carattere r t volte
func stampaRune(r rune, t int) {
	for i := 0; i < t; i++ {
		fmt.Print(string(r))
	}
}

func stampaFiguraRic(n, indent int) {
	if n == 1 {
		stampaRune(' ', indent)
		fmt.Println("*")
		return
	}
	stampaRune(' ', indent)
	stampaRune('*', n)
	fmt.Println()
	
	stampaFiguraRic(n-1, indent+1)
	
	stampaRune(' ', indent)
	stampaRune('*', n)
	fmt.Println()
}

func stampaFigura(n int) {
	stampaFiguraRic(n, 0)
}

func main() {
	stampaFigura(7)
}
