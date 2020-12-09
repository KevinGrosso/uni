package main 

import "fmt" 

// dato un valore intero, mostra il simbolo unicode corrispondente

func main() {
	var codePoint rune
	
	fmt.Print("Code point: ")
	fmt.Scan(&codePoint)
	
	fmt.Println(string(codePoint))
}
