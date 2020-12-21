package main

import (
	"fmt"
	"os"
	"strconv"
)

type Frazione struct {num, den int}

//definire le seguenti funzioni

func CreaFrazione(n,d int) Frazione {
   return Frazione{n, d}
}

func CreaFrazioneParticolare (n int) Frazione {
	return CreaFrazione(n, 1)
} 

//CreaFrazioneDaStringa crea una Frazione da una stringa, come ad es "5/7".
//Ritorna la coppia (Frazione{0,0}, false) se la stringa s non è nel formato attesa
func CreaFrazioneDaStringa(s string) (Frazione,bool) {
	f := Frazione{}

	if s[0] == '/' {
		return f, false
	}

	index := 0
	num := ""
	
	for i, r := range s {
		if r >= '0' && r <= '9' {
			num += string(r)
		} else if r == '/' {
			index = i
			break
		} else {
			return f, false
		}
	}
	
	den := ""
	for _, r := range s[index+1:] {
		if r >= '0' && r <= '9' {
			den += string(r)
		} else {
			return f, false
		}
	}

	var err1, err2 error
	f.num, err1 = strconv.Atoi(num)
	f.den, err2 = strconv.Atoi(den)
	if err1 != nil || err2 != nil {
		return Frazione{}, false
	}
	
	return f, true
}

func (f1 Frazione) Somma (f2 Frazione) Frazione {
	f := Frazione{}
	f.den = f1.den * f2.den
	f.num = f1.num * f2.den + f2.num * f1.den 
	return f
}

//Sommatoria calcola (ricorsivamente) la Frazione che risulta sommando
//le frazioni passate come argomento
func Sommatoria (sommafrazioni []Frazione) Frazione {
	if len(sommafrazioni) == 0 {
		 return CreaFrazione(0,0) //ritorniamo una frazione non definita
	}
	if len(sommafrazioni) == 1 {
		return sommafrazioni[0]
	}
	//la lunghezza della slice è > 1: caso ricorsivo
	return sommafrazioni[0].Somma(Sommatoria(sommafrazioni[1:]))
}

//main legge da linea di comando una serie di frazioni e ne
//stampa la somma (se possibile, ma non necessariamente, in forma normalizzata)
//nel caso la serie sia vuota produrrà il messaggio "nessun argomento"
//sullo standard error
func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "nessun argomento")
		os.Exit(1)
	}

	frazioni := make([]Frazione, 0, 1)
	for _, fs := range os.Args[1:] {
		f, ok :=  CreaFrazioneDaStringa(fs); if !ok {
			fmt.Fprintf(os.Stderr, "il valore %s non rappresenta una frazione valida\n", fs)
			os.Exit(1)
		}
		frazioni = append(frazioni, f)
	}
	somma := Sommatoria(frazioni)
	fmt.Fprintf(os.Stdout, "%d/%d", somma.num, somma.den)
}
