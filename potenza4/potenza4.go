package main

import "fmt"

//Calcola la potenza quarta di un numero

func main() {
	var num int

	fmt.Print("Inserisci un numero intero: ")
	fmt.Scan(&num)
	
	fmt.Println("La potenza quarta di", num, "è",num * num *num * num)		
}
