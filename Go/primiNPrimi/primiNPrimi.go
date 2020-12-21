package main 

import "fmt" 

// dato n, stampare i primi n numeri pirmi

func main() {
	var n, j, quanti int
	
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)
	
	for i := 2; quanti < n; i++ {
		for j = 2; j <= i/2; j++ {
			if i%j == 0 {
				break
			}
		}
		if j > i/2 {
			fmt.Println(i)
			quanti++
		}
	}
}
