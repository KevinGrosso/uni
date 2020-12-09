package main

import "fmt"

/*
	Stampa num asterischi e va a capo
*/
func stampaRigaAsterischi(num int) {
	for i := 0; i < num; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func main() {
	stampaRigaAsterischi(50)
	fmt.Println("titolo")
	stampaRigaAsterischi(50)
}
