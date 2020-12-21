package main 

import "fmt" 

/*
	n = 5
0	*****
1	 ****
2	  ***
3	   **
4	    *

*/

func main() {
	var n int
	
	fmt.Print("Inserisci la lunghezza del lato: ")
	fmt.Scan(&n)
	
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		
		for j := i; j < n; j++ {
			fmt.Print("*")
		}
		
		fmt.Println()
	}
}
