package main 

import (
	"fmt" 
	"math"
)

/* 
	data l'altezza di partenza, calcola il tempo di 
	caduta di un corpo e la sua velocità all'impatto
*/

func main() {
	var h, g float64
	g = 9.81
	
	fmt.Print("Inserire l'altezza in metri: ")
	fmt.Scan(&h)
	
	t := math.Sqrt(2*h/g)
	fmt.Println("Il tempo di caduta è di", t, "sec.")
	fmt.Println("La velocità del corpo all'impatto è di", t*g, "al secondo")
}
