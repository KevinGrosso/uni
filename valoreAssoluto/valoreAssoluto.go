package main 

import "fmt" 

//

func main() {
	var num int
	
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&num)
	
	fmt.Println("Il valore assoluto di", num, "è", num*((2*num+1) % 2))
}
