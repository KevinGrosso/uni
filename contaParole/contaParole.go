package main

import (
	"bufio"
	"fmt"
	"os"
)

// legge una riga e stampa quante parole contiene

func main() {
	var parole int

	fmt.Print("Scrivi una riga di testo: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	riga := scanner.Text()

	if riga != "" {
		riga += " "
		for _, r := range riga {
			if r == ' ' || r == '\n' {
				parole++
			}
		}
	}
	fmt.Println("Hai scritto", parole, "parole")
}
