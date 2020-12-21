package main 

import (
	"fmt" 
	"bufio"
	"os"
)

// inserisce '*' al posto dei caratteri non ASCII presenti in una riga data in input

func main() {
	fmt.Print("Inserisci una riga: ")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	riga := scan.Text()
	
	for _, r := range riga {
		if r > 127 {
			fmt.Print("*")
		} else {
			fmt.Print(string(r))
		}
	}
	fmt.Println()
}
