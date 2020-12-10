package main 

import (
	"fmt"
	"unicode"
	"bufio"
	"os"
)

func main() {
	fmt.Print("$: ")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	t := scan.Text()
	
	v, c := contaMinuscole(t)
	fmt.Println("Vocali minuscole:", v, "/ Consonanti minuscole:", c)
}

/* data una stringa, conta e restituisce il numero di vocali e consonanti */
func contaMinuscole(s string) (vocali int, consonanti int) {
	for _, r := range(s) {
		if unicode.IsLower(r) && unicode.IsLetter(r) {
			switch r {
			case 'a', 'e', 'i', 'o', 'u':
				vocali++
			default:
				consonanti++
			}
		}
	}
	return
}