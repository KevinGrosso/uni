package main 

import "fmt" 

// somma dei primi n numeri dispari

func main() {
	var n, somma int
	
	fmt.Print("N: ")
	fmt.Scan(&n)
	
	for i := 0; i < n; i++{
		somma += 2*i+1
	}
	
	fmt.Println("La somma dei primi N numeri dispari è:", somma)
}
