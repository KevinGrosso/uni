package main 

import "fmt" 

/* 
	dato n int, stampa fizz se è divisibile per 3, 
	stampa buzz se è divisibile per 5, stampa fizz buzz
	se è divisibile sia per 3 sia per 5
*/

func main() {
	var num int
	
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&num)
	
	if num % 3 == 0 {
		fmt.Print("fizz ")
	}
	if num % 5 == 0 {
		fmt.Print("buzz")
	}
	fmt.Print("\n")
}
