package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"math"
)

type Punto struct {
	etichetta string
	x, y float64
}

type TriangoloRettangolo struct {
	area float64
	vertici [3]Punto
}

func main() {
	var punti []Punto
	
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		s := scanner.Text()
		tmp := strings.Split(s, ";")

		p := Punto{}
		p.etichetta = tmp[0]
		p.x, _ = strconv.ParseFloat(tmp[1], 64)
		p.y, _ = strconv.ParseFloat(tmp[2], 64)

		punti = append(punti, p)
	}
	
	var triangolo TriangoloRettangolo

	// A B C D E F G
	
	for i := 0; i < len(punti); i++ {
		for j := 0; j < len(punti); j++ {
			if i != j {
				if punti[i].y == punti[j].y { // sono paralleli all'asse X
					base := Distanza(punti[i], punti[j])
					for k := 0; k < len(punti); k++ {
						if k != i && k != j {
							altezza := 0.0
							if punti[k].x == punti[i].x {
								altezza = Distanza(punti[k], punti[i])
							} else if punti[k].x == punti[j].x {
								altezza = Distanza(punti[k], punti[j])
							}
							if altezza != 0 {
								area := base * altezza / 2
								if area > triangolo.area {
									triangolo.area = area
									triangolo.vertici[0] = punti[i]
									triangolo.vertici[1] = punti[j]
									triangolo.vertici[2] = punti[k]
								}
							}
						}
					}
				}
			}
		}
	}

	if triangolo.area != 0 {
		fmt.Println(StringTriangoloRettangolo(triangolo))
	}
}

func Distanza(p1, p2 Punto) float64 {
	x := p1.x - p2.x
	y := p1.y - p2.y

	x *= x
	y *= y

	return math.Sqrt(x + y)
}

func StringPunto(p Punto) string {
	return fmt.Sprintf("%s = (%f, %f)", p.etichetta, p.x, p.y)
}

func StringTriangoloRettangolo(t TriangoloRettangolo) string {
	return fmt.Sprintf("Triangolo rettangolo con vertici %s, %s e %s, ed area %f",
						StringPunto(t.vertici[0]), StringPunto(t.vertici[1]), 
						StringPunto(t.vertici[2]), t.area)
}
