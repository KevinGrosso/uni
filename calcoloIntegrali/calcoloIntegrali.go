package main

import (
    "fmt"
    "math"
)

type funz func(float64) float64

//Calcola l'integrale definito da a a b della funzione f con il 
//metodo dei trapezoidi (n numero di trapezi)
func integraleTrap(f funz, a, b float64, n int) float64 {
	delta := (b - a)/float64(n) //lunghezza di ogni trapezio
	area := 0.0
	
	for i := 0.0; i < float64(n); i++ {
		xi := a + i * delta 				//estremo sinistro
		xipu := a + (i + 1) * delta			//estremo destro
		hi := f(xi)							//altezza estremo sinistro
		hipu := f(xipu)						//altezza estremo destro
		atrap := (hi + hipu) * delta / 2	//area trapezio
		area += atrap
	}
	return area
}

func sin3x(x float64) float64 {
	return math.Sin(3 * x)
}

func main() {
	fmt.Println("Integrale definito di sin(x) da 3 a 7:\t", integraleTrap(math.Sin, 3, 7, 10000))
	fmt.Println("Integrale definito di sin(3x) da 3 a 7:\t", integraleTrap(sin3x, 3, 7, 10000))
}
