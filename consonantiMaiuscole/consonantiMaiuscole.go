package main 

import (
	"fmt" 
	"bufio"
	"os"
)

// data una stringa, la stampa con le consonanti maiuscole

func main() {
	fmt.Print("Inserisci una riga: ")
	
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	riga := scan.Text()
	
	for _, r := range riga {
		if r != 'a' && r != 'e' && r != 'i' && r != 'o' && r != 'u' && r > 96 && r < 123 {
			fmt.Print(string(r - 'a' + 'A'))
		} else {
			fmt.Print(string(r))
		}
	}
	fmt.Println()
}
