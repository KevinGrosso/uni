package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Cerchio struct {
	nome         string
	x, y, raggio float64
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	prevCerchio := Cerchio{"", 0.0, 0.0, 0.0}

	for scanner.Scan() {
		descrCerchio := scanner.Text()
		cerchio := newCircle(descrCerchio)

		fmt.Print(cerchio)
		if !tocca(cerchio, prevCerchio) {
			fmt.Print(" non")
		}
		fmt.Print(" tangente, ")
		if maggiore(cerchio, prevCerchio) {
			fmt.Print("maggiore\n")
		} else {
			fmt.Print("minore o uguale\n")
		}

		prevCerchio = cerchio
	}
}

func newCircle(descr string) Cerchio {
	var cerchio Cerchio

	valori := strings.Split(descr, " ")
	cerchio.nome = valori[0]
	cerchio.x, _ = strconv.ParseFloat(valori[1], 64)
	cerchio.y, _ = strconv.ParseFloat(valori[2], 64)
	cerchio.raggio, _ = strconv.ParseFloat(valori[3], 64)

	return cerchio
}

func distanzaPunti(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
}

func tocca(cerc1, cerc2 Cerchio) bool {
	if math.Abs(distanzaPunti(cerc1.x, cerc1.y, cerc2.x, cerc2.y)-cerc1.raggio-cerc2.raggio) < 1e-6 {
		return true
	} else if math.Abs(distanzaPunti(cerc1.x, cerc1.y, cerc2.x, cerc2.y)-math.Abs(cerc1.raggio-cerc2.raggio)) < 1e-6 {
		return true
	}

	return false
}

func maggiore(cerc1, cerc2 Cerchio) bool {
	areaCerc1 := math.Pi * cerc1.raggio * cerc1.raggio
	areaCerc2 := math.Pi * cerc2.raggio * cerc2.raggio

	if areaCerc1 > areaCerc2 {
		return true
	}
	return false
}
