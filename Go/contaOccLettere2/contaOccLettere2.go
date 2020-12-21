package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Inserisci una stringa: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()

	occ := contaOccorrenze(s)
	stampaOccorenze(occ)
}

func contaOccorrenze(s string) (occ ['z' - 'a' + 1]int) {
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			occ[s[i]-'a']++
		}
	}
	return
}

func stampaOccorenze(occorrenza ['z' - 'a' + 1]int) {
	for i := 0; i < 'z'-'a'+1; i++ {
		if occorrenza[i] > 0 {
			fmt.Printf("%c -> %d\n", 'a'+i, occorrenza[i])
		}
	}
	return
}
