package main 

import "fmt" 

// fattoriale di N

func main() {
	var n int
	
	fmt.Print("N: ")
	fmt.Scan(&n)
	
	fattoriale := 1
	
	for i := 0; i < n; i++ {
		fattoriale *= n - i
	}
	
	fmt.Println("Il fattoriale di", n, "è", fattoriale)
}
