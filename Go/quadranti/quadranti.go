package main 

import "fmt" 

// dato un punto(x, y) stabilsce in che quadrante si trova

func main() {
	var x, y float64
//	var quadrante int
	
	fmt.Print("Inserisci la coordinata X del punto: ")
	fmt.Scan(&x)
	fmt.Print("Inserisci la coordinata Y del punto: ")
	fmt.Scan(&y)
	
	if x > 0 && y > 0 {
		fmt.Println("Il punto si trova nel primo quadrante")
	} else if x < 0 && y > 0 {
		fmt.Println("Il punto si trova nel secondo quadrante")
	} else if x < 0 && y < 0 {
		fmt.Println("Il punto si trova nel terzo quadrante")
	} else if x > 0 && y < 0 {
		fmt.Println("Il punto si trova nel quarto quadrante")
	} else {
		fmt.Println("Indeterminato")
	}

/*	Soluzione del prof

	if x >= 0 && y >= 0 {
		quadrante = 1
	} else if x >= 0 {
		quadrante = 4
	} else if y >= 0 {
		quadrante = 2
	} else {
		quadrante = 3
	}
	fmt.Println("Il punto si trova nel", quadrante, "° quadrante")
*/
}