package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
	"math"
)

type Punto struct {
	etichetta string
	x, y float64
}

func main() {
	tragitto := NuovoTragitto()
	lung := Lunghezza(tragitto)

	fmt.Printf("Lunghezza percorso: %.3f\n", lung)

	tot := 0.0
	for i := 0; i < len(tragitto) - 1; i++ {
		if tot > lung/2.0 {
			fmt.Println("Punto oltre metà: ", String(tragitto[i]))
			break
		}
		tot += Distanza(tragitto[i], tragitto[i+1])
	}
}

func Distanza(p1, p2 Punto) float64 {
	x := p1.x - p2.x
	y := p1.y - p2.y

	x *= x
	y *= y

	return math.Sqrt(x + y)
}

func String(p Punto) string {
	return fmt.Sprintf("%s = (%.1f, %.1f) ", p.etichetta, p.x, p.y)
}

func Lunghezza(tragitto []Punto) float64 {
	risultato := 0.0

	for i := 0; i < len(tragitto)-1; i++ {
		risultato += Distanza(tragitto[i], tragitto[i+1])
	}

	return risultato
}

func NuovoTragitto() (tragitto []Punto) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ";")
		x, _ := strconv.ParseFloat(s[1], 64)
		y, _ := strconv.ParseFloat(s[2], 64) 

		p := Punto{s[0], x, y}
		tragitto = append(tragitto, p)
	}

	return
}