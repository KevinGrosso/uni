package main 

import "fmt" 

// dato un numero, stabilisce se contiene la cifra 3

func main() {
	var num int
	
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&num)
	
	for cifra := num%10; num != 0; cifra=num%10 {
		if cifra == 3 {
			break;
		}
		num /= 10;
	}
	
	if num != 0 {
		fmt.Println("Il numero contiene la cifra 3")
	} else {
		fmt.Println("Il numero non contiene la cifra 3")
	}
}
