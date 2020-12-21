package main 

import "fmt" 

// dato un cateto, stampa un triangolo rettangolo isoscele fatto di asterischi

func main() {
	var lato int
	
	fmt.Print("Inserisci la lunghezza del cateto: ")
	fmt.Scan(&lato)
	
	for i := 0; i < lato; i++ {
		for j:= 0; j < lato - i; j++ {
			fmt.Print("*")
		}
		
		fmt.Println()
	}
}
