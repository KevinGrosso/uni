package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	var valori []float64
	
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		riga := scanner.Text()
		f, _ := strconv.ParseFloat(riga, 64)
		valori = append(valori, f)
	}

	fmt.Println("moda:", moda(valori))
	fmt.Println("mediana:", mediana(valori))
}

func moda(serie []float64) float64 {
	freq := make(map[float64]int)

	// conto le frequenze
	for _, v := range serie {
		freq[v]++
	}

	// trovo la frequenza massima
	max := 0
	for _, f := range freq {
		if f > max {
			max = f
		}
	}

	// trovo il valore piú piccolo che abbia la frequenza massima
	min := math.MaxFloat64
	for k, f := range freq {
		if f == max && k <= min {
			min = k
		}
	}
	
	return min
}

func mediana (serie []float64) float64 {
	sort.Float64s(serie)
	lungh := len(serie)
	
	if lungh % 2 != 0 {
		return serie[lungh / 2]
	}
	return (serie[lungh/2 - 1] + serie[lungh / 2]) / 2
}
