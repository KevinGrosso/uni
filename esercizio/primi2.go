package main

import (
	"fmt"
//	"math"
)

func main() {
	var n int = 10000000
	
	for candidato := 2; candidato <= n; candidato++ {
		var divisore int
		
		// sqrt del candidato arrotondata
		//var sqrtCandidato = int(math.Sqrt(float64(candidato))) + 1
		
		for divisore = 2; divisore < candidato; divisore++ {
			if candidato%divisore == 0 {
				break
			}
		}
		
		if divisore >= candidato {
			fmt.Println(candidato) // è primo
		}
	}
}
