package main

import (
	"fmt"
	"os"
)

const lunghezzaCombinazione = 5
const caratteri = "abcdefghijklmnopqrstuvwxyz"
const numeroCaratteri = 26

var tentativiTotali int

func win() {
	fmt.Println("Ho indovinato la parola in", tentativiTotali, "tentativi!")
	os.Exit(0)
}

func generaCombinazioni(r []rune, l int) *[]string {
	if l == 0 {
		return &[]string{""}
	}

	var combinazioni []string

	for i := 0; i < len(r); i++ {
		lista := generaCombinazioni(r, l-1)

		for _, x := range *lista {
			combinazioni = append(combinazioni, string(r[i])+x)
		}
		lista = nil
	}
	return &combinazioni
}

func chiedi(tentativo string) (inPosizione, giusti int) {
	tentativiTotali++

	fmt.Println("Il mio tentativo è", tentativo)
	fmt.Print("Quanti caratteri giusti al posto giusto ho messo? ")
	fmt.Scan(&inPosizione)

	if inPosizione == 5 {
		win()
	}

	fmt.Print("Quanti caratteri giusti al posto sbagliato ho messo? ")
	fmt.Scan(&giusti)
	return
}

// 01234
// abcde
func generaTestPosizione(t string, i int, e string) string {
	ris := ""

	for j := 0; j < i; j++ {
		ris += e
	}
	ris += t
	for j := 0; j < lunghezzaCombinazione-i-1; j++ {
		ris += e
	}
	return ris
}

func generaTestGiusto(t string, r []rune, p int, e string) (string, int) {
	for i := 0; i < len(r); i++ {
		fmt.Println("DEBUG:::r[i] =", r[i], "i =", i, "p =", p)
		if r[i] == 0 && i != p {
			return generaTestPosizione(t, i, e), i
		}
	}
	return "", -1
}

// Si potrebbe ottimizzare meglio e fare meno tentativi controllando quanti
// caratteri giusti sono stati inseriti in posizione sbagliata
// (rifacendo le combinazioni con le nuove informazioni apprese)
func completaCombinazione(r []rune, p string) {
	quanti := 0
	var dove []int

	for i := 0; i < lunghezzaCombinazione; i++ {
		if r[i] == 0 {
			quanti++
			dove = append(dove, i)
		}
	}

	combo := generaCombinazioni([]rune(p), quanti)

	for i := 0; i < len(*combo); i++ {
		for j := 0; j < quanti; j++ {
			r[dove[j]] = rune((*combo)[i][j])
		}
		chiedi(string(r)) // TODO: ottimizzare per fare meno tentativi
	}
}

func main() {
	var x string
	fmt.Println("Pensa ad una combinazione di", lunghezzaCombinazione, "valori.")
	fmt.Println("Può essere composta dai seguenti caratteri:", caratteri, ".")
	fmt.Print("Premi invio quando sei pronto...")
	fmt.Scanln(&x)

	caratteriEsclusi := ""
	comboPossibili := make(map[string][3]int)
	caratteriTrovati := ""
	risultato := make([]rune, 5)

	for i := 0; i < len(caratteri)-1; i += lunghezzaCombinazione {
		tentativo := caratteri[i : i+lunghezzaCombinazione]
		inPosizione, giusti := chiedi(tentativo)

		if inPosizione > 0 && giusti > 0 {
			comboPossibili[tentativo] = [3]int{0, inPosizione, giusti}
		} else if inPosizione > 0 {
			comboPossibili[tentativo] = [3]int{1, inPosizione, 0}
		} else if giusti > 0 {
			comboPossibili[tentativo] = [3]int{-1, giusti, 0}
		} else {
			caratteriEsclusi += tentativo
		}
	}

	/*for k, v := range comboPossibili {
		fmt.Println(k, v)
	}*/

	// z caso particolare
	//inPosizione, giusti := chiedi("z" + escluso[:4])
	// da implementare singolarmente z**** *z*** **z** etc.

	inPosizione, giusti := chiedi("zzzzz")
	if inPosizione+giusti == 0 {
		caratteriEsclusi += "z"
	} else {
		caratteriTrovati += "z"

		for j, i := 0, 0; j < lunghezzaCombinazione && i < inPosizione; j++ {
			pos, _ := chiedi(generaTestPosizione("z", j, string(caratteriEsclusi[0])))

			if pos == 1 {
				risultato[j] = 'z'
				i++
			}
		}
	}

	escluso := string(caratteriEsclusi[0])

	for k, v := range comboPossibili {
		switch v[0] {
		case 1:
			// abc -> a** *b* **c
			for j, i := 0, 0; j < lunghezzaCombinazione && i < v[1]; j++ {
				if risultato[j] == 0 {
					inPosizione, _ := chiedi(generaTestPosizione(string(k[j]), j, escluso))
					if inPosizione == 1 {
						i++
						caratteriTrovati += string(k[j])
						risultato[j] = rune(k[j])
					}
				}
			}

		case -1:
			// waplc abcde **+**
			for j, i := 0, 0; j < lunghezzaCombinazione && i < v[1]; j++ {
				tentativo, dove := generaTestGiusto(string(k[j]), risultato, j, escluso)
				if tentativo == "" {
					fmt.Println("shiit")
					chiedi(string(risultato))
				}
				inPosizione, giusti := chiedi(tentativo)
				if inPosizione == 1 {
					i++
					caratteriTrovati += string(k[j])
					risultato[dove] = rune(k[j])
				} else if giusti == 1 {
					i++
					for n := 0; n < lunghezzaCombinazione; n++ {
						if n != j && n != dove && risultato[n] == 0 {
							pos, _ := chiedi(generaTestPosizione(string(k[j]), n, escluso))
							if pos == 1 {
								caratteriTrovati += string(k[j])
								risultato[n] = rune(k[j])
								break
							}
						}
					}
				}
			}

		case 0:
			for j, i := 0, 0; j < lunghezzaCombinazione && i < v[1]; j++ {
				if risultato[j] == 0 {
					inPosizione, _ := chiedi(generaTestPosizione(string(k[j]), j, escluso))
					if inPosizione == 1 {
						i++
						caratteriTrovati += string(k[j])
						risultato[j] = rune(k[j])
					}
				}
			}

			for j, i := 0, 0; j < lunghezzaCombinazione && i < v[2]; j++ {
				// fix: wtf
				tentativo, dove := generaTestGiusto(string(k[j]), risultato, j, escluso)
				if tentativo == "" {
					fmt.Println("shiit")
					chiedi(string(risultato))
				}
				inPosizione, giusti := chiedi(tentativo)
				if inPosizione == 1 {
					i++
					caratteriTrovati += string(k[j])
					risultato[dove] = rune(k[j])
				} else if giusti == 1 {
					i++
					for n := 0; n < lunghezzaCombinazione; n++ {
						if n != j && n != dove && risultato[n] == 0 {
							pos, _ := chiedi(generaTestPosizione(string(k[j]), n, escluso))
							if pos == 1 {
								caratteriTrovati += string(k[j])
								risultato[n] = rune(k[j])
								break
							}
						}
					}
				}
			}

		}
	}

	n := 0
	for n := 0; n < lunghezzaCombinazione; n++ {
		if risultato[n] == 0 {
			break
		}
	}

	if n == lunghezzaCombinazione {
		win()
	}

	completaCombinazione(risultato, caratteriTrovati)
}
