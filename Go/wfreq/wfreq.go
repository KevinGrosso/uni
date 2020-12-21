package main

import (
    "fmt"
    "bufio"
    "os"
    "unicode"
)

func controllaParola(s string) bool {
	if len(s) == 0 {
		return false
	}
	
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func separaParole(s string) []string {
	parole := []string{}
	oldI := 0
	s += " "
	
	for i, r := range s {
		if r == ' ' {
			if controllaParola(s[oldI:i]) {
				parole = append(parole, s[oldI:i])
			}
			oldI = i + 1
		}
	}
	return parole
}

func main() {
	occorrenze := make(map[string]int)
	
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		s := scanner.Text()
		for _, p := range separaParole(s) {
			occorrenze[p]++
		}
	}

	fmt.Printf("\n\n[Numero di parole distinte: %d]\n", len(occorrenze))
	
	for k, o := range occorrenze {
		fmt.Printf("%s:\t%d\n", k, o)
	}	
}
