package main 

import "fmt" 

// calcola la somma dei quadrati dei primi n numeri

func main() {
	var num, sum, sum2 int
	
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&num)
	
	for i := 1; i <= num; i++ {
		sum += i * i
	}
	
	sum2 = (num*(num+1)*(2*num+1))/6
	
	fmt.Println("La somma dei primi n numeri al quadrato con il ciclo for è", sum)
	fmt.Println("La somma dei primi n numeri al quadrato con la formula è", sum2)
}
