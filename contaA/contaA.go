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
	
	for i := 0; i < len(testo); i++ {
		if testo[i] == 'a' {
			quante++
		}
	}
	
	fmt.Println("Il testo inserito contiene", quante, "a")
}
