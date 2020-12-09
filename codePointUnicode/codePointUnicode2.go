package main 

import "fmt" 

// dato un simbolo unicode, mostra il corrispondente code point

func main() {
	var carattere string
	
	fmt.Print("Carattere: ")
	fmt.Scan(&carattere)
	
	runes := []rune(carattere)
	
	fmt.Println(runes[0])
}
