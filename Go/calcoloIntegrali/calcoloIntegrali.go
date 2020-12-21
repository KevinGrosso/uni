package main

import (
    "fmt"
    "math"
    "math/rand"
    "time"
)

type funz func(float64) float64

// Calcola l'integrale definito da a a b della funzione f con il 
// metodo dei trapezoidi (n numero di trapezi)
func integraleTrap(f funz, a, b float64, n int) float64 {
	delta := (b - a)/float64(n) //lunghezza di ogni trapezio
	area := 0.0
	
	for i := 0.0; i < float64(n); i++ {
		xi := a + i * delta 				// estremo sinistro
		xipu := a + (i + 1) * delta			// estremo destro
		hi := f(xi)							// altezza estremo sinistro
		hipu := f(xipu)						// altezza estremo destro
		atrap := (hi + hipu) * delta / 2	// area trapezio
		area += atrap
	}
	return area
}

// Calcola l'integrale definito da a a b della funzione f con il 
// metodo di integrazone monte-carlo
func integraleMC(f funz, a, b, A float64, n int) float64 {
	areaRett := A * (b - a)
	aSegno := 0
	
	for i := 0; i < n; i++ {
		y := A * rand.Float64() 
		x := a + (b - a) * rand.Float64()

		if y < f(x) {
			aSegno++
		}
	}
	return areaRett * float64(aSegno) / float64(n)
}

func sin3x(x float64) float64 {
	return math.Sin(3 * x)
}

func sinquadrox(x float64) float64 {
	return math.Sin(x) * math.Sin(x)
}

func xquadropotto(x float64) float64 {
	return x * x + 8
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Integrali con il metodo dei trapezoidi:")	
	fmt.Println("Integrale definito di sin(x) da 3 a 7:", integraleTrap(math.Sin, 3, 7, 100000))
	fmt.Println("Integrale definito di sin(3x) da 3 a 7:", integraleTrap(sin3x, 3, 7, 100000))
	fmt.Println("Integrale definito di sin^2(x) da 3 a 7:", integraleTrap(sinquadrox, 3, 7, 1000000))
	fmt.Println("Integrale definito di x^2 + 8 da 3 a 7:", integraleTrap(xquadropotto, 3, 7, 1000000))

	fmt.Println()
 	fmt.Println("Integrali con il metodo di monte-carlo:")
	fmt.Println("Integrale definito di sin^2(x) da 3 a 7:", integraleMC(sinquadrox, 3, 7, 5, 100000000))
	fmt.Println("Integrale definito di x^2 + 8 da 3 a 7:", integraleMC(xquadropotto, 3, 7, 200, 100000000))
}
