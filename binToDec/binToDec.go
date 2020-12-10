package main 

import "fmt" 

// dato n, controlla che sia formato da soli 0,1 e stampa il valore del numero

func main() {
	var n, valore int
	
	fmt.Print("N: ")
	fmt.Scan(&n)
	
	for potenza := 1; n != 0; potenza *= 2 {
		bit := n % 10
		if bit == 1 || bit == 0 {
			valore += potenza * bit
			n /= 10
		} else {
			break
		}
	}
	
	if n == 0 {
		fmt.Println(valore)
	} else {
		fmt.Println("Numero binario non valido")
	}
}
