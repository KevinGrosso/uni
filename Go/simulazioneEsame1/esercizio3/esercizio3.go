package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"math"
)

type Punto struct {
	ascissa, ordinata float64
	etichetta string
}

type Segmento [2]Punto

func main() {
	soglia, _ := strconv.ParseFloat(os.Args[1], 64)

	var punti []Punto
	
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		punti = append(punti, GeneraPunto(s))
	}

	var seg Segmento
	for i := 0; i < len(punti); i++ {
		seg[0] = punti[i]
		for j := i; j < len(punti); j++ {
			seg[1] = punti[j]
			if !paralleloAssi(seg) && Distanza(seg[0], seg[1]) > soglia {
				fmt.Println(StringSegmento(seg))
			}
		}
	}
	
}

func paralleloAssi(seg Segmento) bool {
	if seg[0].ascissa == seg[1].ascissa || seg[0].ordinata == seg[1].ordinata {
		return true
	}
	return false
}

func GeneraPunto(s string) Punto{
	p := Punto{}
	tmp := strings.Split(s, ";")
	
	p.etichetta = tmp[0]
	p.ascissa, _ = strconv.ParseFloat(tmp[1], 64) 
	p.ordinata, _ = strconv.ParseFloat(tmp[2], 64) 
	
	return p
}

func Distanza(p1, p2 Punto) float64 {
	a := p1.ascissa - p2.ascissa
	a *= a // a^2
	b := p1.ordinata - p2.ordinata
	b *= b // b^2
	
	return math.Sqrt(a + b)
}

func StringPunto(p Punto) string {
	return fmt.Sprintf("%s = (%.6f, %.6f)", p.etichetta, p.ascissa, p.ordinata)
}

func StringSegmento(s Segmento) string {
	return "Segmento con estremi " + StringPunto(s[0]) + " e " + StringPunto(s[1])
}
