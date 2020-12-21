package main 

import "fmt" 

// dato n, stampa un quadrato di asterischi con area n*n

func main() {
	var num int
	
	fmt.Print("Inserisci la lunghezza del lato: ")
	fmt.Scan(&num)
	
	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
