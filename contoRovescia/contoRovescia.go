package main 

import "fmt" 

// stampa un conto alla rovescia da n

func main() {
	var num int
	
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&num)
	
	for ; num >= 0; num-- {
		fmt.Println(num)
	}
}
