package main

import (
	"testing"
	"math/rand"
)

func TestMedia1(t *testing.T) {
	var s studente
	s = studente{"Kevin_Grosso", []int{30, 18, 18}}
	actualMedia := mediaVoti(s)
	expectedMedia := 22.0
	if actualMedia != expectedMedia {
		t.Error("Expected", expectedMedia, "got", actualMedia)
	}
}

func TestMedia2(t *testing.T) {
	type testCase struct {
		s studente
		expectedMedia float64
	}
	var testCases []testCase = []testCase {
		testCase{studente{"Kevin_Grosso", []int{18, 18, 18}}, 18.0},
		testCase{studente{"Mario_Rossi", []int{32, 30}}, 30.0},
		testCase{studente{"Paolo_Bianchi", []int{20, 24}}, 22.0},
	}

	for _, tc := range testCases {
		actualMedia := mediaVoti(tc.s)
			if actualMedia != tc.expectedMedia {
				t.Error("Expected", tc.expectedMedia, "got", actualMedia)
			}
	}
}

func TestRandomMedia(t *testing.T) {
	var s studente
	s.Nome = "Kevin_Grosso"
	for i := 0; i < 100; i++ {
		v1 := 18 + rand.Intn(30 - 18 + 1)
		v2 := 18 + rand.Intn(30 - 18 + 1)
		s.Voti = []int{v1, v2}
		actual := mediaVoti(s)
		if float64(v1+v2)/2.0 != actual {
			t.Error("Expected", (v1+v2)/2, "got", actual)
		}
	}
}
