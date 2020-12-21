package main 

import "fmt" 

// dato un numero, stabilisce se è primo

func main() {
	var num, i int
	
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&num)
	
	for i=2; i <= num/2 ; i++ {
		if num%i == 0 {
			break
		}
	}
	
	if i <= num/2 {
		fmt.Println("Il numero", num, "non è primo")
	} else {
		fmt.Println("Il numero", num, "è primo")
	}
}
