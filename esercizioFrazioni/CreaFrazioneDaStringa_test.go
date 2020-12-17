package main

import "testing"

func TestCreaFrazioneDaStringa(t *testing.T) {
	f, ok := CreaFrazioneDaStringa("5/7")
	if f.num != 5 || f.den != 7 || !ok{
		t.Error("ERRORE:", f.num, f.den, ok)
	}
}
