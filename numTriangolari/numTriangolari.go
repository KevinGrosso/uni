package main 

import "fmt" 

// visualizzare i primi n numeri triangolari

func main() {
	var n, somma int
	
	fmt.Print("N: ")
	fmt.Scan(&n)
	
	for i := 1; i <= n; i++ {
		somma += i
		fmt.Println(somma)
	}
}
