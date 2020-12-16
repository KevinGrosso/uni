package main

import (
	"fmt"
	"regexp"
	"os"
	"io/ioutil"
)

const vocali = "aeiou"

func generaSillabe(c string) []string{
	sillabe := make([]string, 5)
	for i, v := range vocali {
		sillabe[i] = c + string(v)
	}
	return sillabe
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Sintassi: %s [lettere separate da spazi]", os.Args[0])
		os.Exit(1)
	}

	sillabe := make([]string, 0, 5 * (len(os.Args) - 1))
	for _, l := range os.Args[1:] {
		sillabe = append(sillabe, generaSillabe(l)...)
	}

	parole, err := ioutil.ReadFile("dizionario.txt")
	if err != nil {		
		fmt.Fprintf(os.Stderr, "Errore con il file dizionario.txt: %s", err)	
		os.Exit(1)
	}
	
	expr := `\b(`
	for _, s := range sillabe {
		expr += s + "|"
	}
	expr = expr[:len(expr)-1] + `){1,10}\b`
	
	paroleRE := regexp.MustCompile(expr) 
	trovate := paroleRE.FindAllString(string(parole), -1)

	for i, p := range trovate {
		fmt.Printf("[%d] %s\n", i, p)
	}
}
