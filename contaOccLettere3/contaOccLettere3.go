package main

import (
	"bufio"
	"fmt"
	"os"
)

type occ ['z' - 'a' + 1]int

func main() {
	var occorrenze occ
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Inserisci delle stringhe: ")
	for scanner.Scan() {
		s := scanner.Text()
		aggiungiOccorrenze(s, &occorrenze)
	}
	stampaOccorenze(occorrenze)
}

func aggiungiOccorrenze(s string, occorrenze *occ) {
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			occorrenze[s[i]-'a']++
		}
	}
	return
}

func stampaOccorenze(occorrenze occ) {
	for i := 0; i < 'z'-'a'+1; i++ {
		if occorrenze[i] > 0 {
			fmt.Printf("%c -> %d\n", 'a'+i, occorrenze[i])
		}
	}
	return
}
