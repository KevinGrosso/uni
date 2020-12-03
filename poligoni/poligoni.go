package poligono

import (
    "fmt"
)

type Punto struct {
	x, y float64
}
type Poligono []punto

//crea un poligono formato da n punti letti da standard input
//ritorna una coppia data dal Poligono creato e dal valore true, se i punti sono tutti diversi tra loro e n > 1
//un qualsiasi Poligono e false altrimenti
func CreaPoligono(n int) (Poligono, bool) {
	
}

func main() {

}

/*
Definire un tipo Poligono che rappresenta un insieme di punti del piano (quindi può essere conveniente
definire anche un tipo Punto)

(in realtà il tipo rappresenta una "linea poligonale")

Definire le seguenti funzioni

//crea un poligono formato da n punti letti da standard input
//ritorna una coppia data dal Poligono creato e dal valore true, se i punti sono tutti diversi tra loro e n > 1
//un qualsiasi Poligono e false altrimenti
func CreaPoligono(n int) (Poligono,bool)

//crea un poligono formato da n punti generati in modo random (v. descrizione sopra)
func CreaPoligonoRandom(n int) (Poligono,bool)

//calcola e ritorna la lunghezza di una linea poligonale, considerando i punti
//secondo l'ordine con cui sono stati letti/generati
func Lunghezza(p Poligono) float64

//definire altre funzioni che possano essere associate a una linea poligonale
*/
