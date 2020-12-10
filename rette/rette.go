package main

import "fmt"

// date 2 rette nella "forma" y = mx + q, trovare l'intersezione

func main() {
	var m1, m2, q1, q2, x, y float64
	
	fmt.Print("Inserisci m e q della prima retta: ")
	fmt.Scan(&m1, &q1)
	fmt.Print("Inserisci m e q della seconda retta: ")
    fmt.Scan(&m2, &q2)
	
	x = (q1 - q2) / (m2 - m1)
	y = m1 * x + q1
	fmt.Println("L'intersezione si trova nel punto (", x, ",", y, ")")
}