package main

import (
    "fmt"
    "bufio"
    "os"
    "unicode"
    "sort"
)

func contaLettere(s string, m map[rune]int) {
	for _, r := range s {
		if unicode.IsLetter(r) {
			m[r]++	
		}
	}
}

func rigaDiAsterischi(n int) string{
	s := ""
	for i := 0; i < n; i++ {
		s += "*"
	}
	return s
}

func stampaIstogramma(m map[rune]int) {
	chiavi := make([]rune, 0, len(m))
	for k, _ := range m {
		chiavi = append(chiavi, k)
	}

	sort.Slice(chiavi, func(i, j int) bool { return chiavi[i] < chiavi[j] })

	for _, k := range chiavi {
		fmt.Printf("%c: %d\t%s\n", k, m[k], rigaDiAsterischi(m[k]))
	}
}

func main() {
	occorrenze := make(map[rune]int)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		s := scanner.Text()
		contaLettere(s, occorrenze)
	}

	stampaIstogramma(occorrenze)
}
