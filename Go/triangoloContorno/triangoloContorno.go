package main 

import "fmt" 

// dato n, stampa il contorno di un triangolo isoscele di lato n

func main() {
	var n int
	
	fmt.Print("Inserisci il lato del triangolo: ")
	fmt.Scan(&n)
	
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			if j == 0 || i == 0 || j == n-i-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
