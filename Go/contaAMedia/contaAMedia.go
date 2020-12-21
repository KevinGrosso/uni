package main 

import (
	"fmt" 
	"bufio"
	"os"
)

// legge righe fino a EOF, alla fine mostra la media di a per riga

func main() {
	var righe int
	var quante float64
	
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		riga := scanner.Text()
		
		for _, r := range riga {
			if r == 'a' {
				quante++
			}
		}
		righe++
	}
	fmt.Println("La media di a per riga è", quante / float64(righe))
}
