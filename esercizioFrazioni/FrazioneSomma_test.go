package main

import (
	"testing"
	"math"
)

func TestFrazioneSomma(t *testing.T) {
	f1 := Frazione{4, 3}
	f2 := Frazione{7, 2}

	f := f1.Somma(f2)

	if math.Abs(float64(f.num)/float64(f.den) - 29.0/6.0) > 0.000001 {
		t.Error("ERRORE:", f.num, f.den)
	}
}
