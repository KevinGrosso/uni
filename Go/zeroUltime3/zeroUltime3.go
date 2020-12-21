package main 

import "fmt" 

// stabilisce se un numero intero contiene almeno uno zero nelle ultime tre cifre

func main() {
	var num int
	var contiene = false
	
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&num)
	
	if (num % 10 == 0) {
		contiene = true
	}
	
	num /= 10
	
	if (num % 10 == 0) {
		contiene = true
	}
	
	num /= 10
	    
	if (num % 10 == 0) {
		contiene = true
	}
	
	if contiene {
		fmt.Println("Il numero contiene almeno uno zero nelle ultime tre cifre")
	} else {
		fmt.Println("Il numero non contiene zeri nelle ultime tre cifre")
	}	                    
}
