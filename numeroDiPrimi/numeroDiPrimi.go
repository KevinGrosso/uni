package main 

import "fmt" 
import "math"

// dato n, calcolare il numero di numeri primi <= n
// verificare che sia asintotico con n/log(n)

func main() {
	var n, j, quanti int
	
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)
	
	
	x := float64(n) / math.Log(float64(n))
	
	for i := 2; i <= n; i++ {
		for j = 2; j <= i/2; j++ {
			if i%j == 0 {
				break
			}
		}
		if j > i/2 {
			quanti++
		}
	}
	
	fmt.Println("I numeri primi <= n sono", quanti)
	fmt.Println(float64(quanti)/x)
}
