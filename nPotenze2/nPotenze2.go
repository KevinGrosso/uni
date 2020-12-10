package main 

import "fmt" 

// stampa le prime n potenze di 2

func main() {
	var num int
	
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&num)
	
	fmt.Println("Prime n potenze di 2")
	for pow, i := 1, 0; i <= num; i, pow = i+1, pow*2 {
		fmt.Println(pow)
	}
}
