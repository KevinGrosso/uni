package main

import "fmt"

// dati due anni di nascita, stabilire quale persona è più vecchia

func main() {
	var a1, a2 int
	
	fmt.Print("Anno di nascita persona 1: ")
	fmt.Scan(&a1)
	fmt.Print("Anno di nascita persona 2: ")
	fmt.Scan(&a2)
	
	if a1 < a2 {
		fmt.Println("La persona 1 è più vecchia della 2")
	} else if a2< a1 {
		fmt.Println("La persona 2 è più vecchia della 1")
	} else {
		fmt.Println("Le persone sono nate nello stesso anno")
	}
}