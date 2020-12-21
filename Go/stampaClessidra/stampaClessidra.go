package main 

import "fmt" 

// stampa una clessidra

func main() {
	var n int
	
	fmt.Print("N: ")
	fmt.Scan(&n)
	
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			if c >= r && c <= n-1-r || c <= r && c >= n-1-r{
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}