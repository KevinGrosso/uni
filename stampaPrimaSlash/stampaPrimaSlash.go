package main 

import (
	"fmt" 
	"bufio"
	"os"	
)

// data una stringa ascii, stampa tutto ciò che compare prima della prima '/'

func main() {
	var indice int
	
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	testo := scanner.Text()
	
	for i := 0; i < len(testo); i++ {
		if testo[i] == '/' {
			indice = i
			break
		}
	}
	
	fmt.Println(testo[:indice])
}
