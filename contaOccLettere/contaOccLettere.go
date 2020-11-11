package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	var s, nextS string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Inserisci una stringa: ")
	scanner.Scan()
	s = scanner.Text()
	nextS = s

	for len(s) > 0 {
		s = nextS
		nextS = ""

		for i, r1 := range s {
			if unicode.IsLetter(r1) {
				var quante = 1
				for j, r2 := range s {
					if r2 == r1 && j > i {
						quante++
					} else if r2 != r1 {
						nextS += string(r2)
					}
				}
				fmt.Println("Numero di ", string(r1), "->", quante)
				break
			}
		}
	}
}
