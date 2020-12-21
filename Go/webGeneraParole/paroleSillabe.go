package main

import (
	_"fmt"
	_"regexp"
	_"os"
	_"io/ioutil"
	"net/http"
)

const vocali = "aeiou"

func generaSillabe(c string) []string{
	sillabe := make([]string, 5)
	for i, v := range vocali {
		sillabe[i] = c + string(v)
	}
	return sillabe
}

func pageImageBack(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "back.png")
}

func pageLanding(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "main.html")
}

func main() {
	http.HandleFunc("/", pageLanding)
	http.HandleFunc("/back.png", pageImageBack)
	http.ListenAndServe("127.0.0.1:4444", nil)

	/*
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
	}*/
}
