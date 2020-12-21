package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s, sub string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Inserisci una stringa: ")
	scanner.Scan()
	s = scanner.Text()

	fmt.Print("Insersci un'altra stringa: ")
	scanner.Scan()
	sub = scanner.Text()

	var i, j int

	for i = 0; i <= len(s)-len(sub); i++ {
		for j = 0; j < len(sub); j++ {
			if s[i+j] != sub[j] {
				break
			}
		}

		if j >= len(sub) {
			fmt.Println("La seconda stringa è contenuta nella prima")
			break
		}
	}
	if i > len(s)-len(sub) {
		fmt.Println("La seconda stringa non è contenuta nella prima")
	}
}
