package main 

import (
	"fmt" 
	"bufio"
	"os"
)

// Legge una stringa (ascii) e stampa quante a contiene

func main() {
	var quante int
	
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	testo := scanner.Text()
	
	for _, r := range testo {
		if r == 'a' {
			quante++
		}
	}
	
	fmt.Println("Il testo inserito contiene", quante, "a")
}
