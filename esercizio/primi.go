package main

import (
	"fmt"
//	"math"
)

func main() {
	var n int = 10000000
	
	for candidato := 3; candidato <= n; candidato += 2 {
		var divisore int
		
		// sqrt del candidato arrotondata
		//var sqrtCandidato = int(math.Sqrt(float64(candidato))) + 1
		
		for divisore = 3; divisore < candidato/divisore; divisore += 2 {
			if candidato%divisore == 0 {
				break
			}
		}
		
		if divisore > candidato/divisore {
			fmt.Println(candidato) // è primo
		}
	}
}
