package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"sort"
)

type data struct {
	g, m, a int
}

type operazione struct {
	tipo rune
	importo float64
}

func main() {
	file, err := os.Open("operazioni.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Errore: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var saldo float64
	scanner := bufio.NewScanner(file)
	
	// Gestisco la prima linea a parte
	scanner.Scan()
	s := scanner.Text()
	saldo, err = strconv.ParseFloat(s, 64)
	
	// Poi gestisco tutte le altre
	operazioni := make(map[data][]operazione)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ";")
		
		d := stringToData(s[0])
		t := []rune(s[1])[0]
		i, _ := strconv.ParseFloat(s[2], 64)
		operazioni[d] = append(operazioni[d], operazione{t, i})
	}

	date := []data{}
	for k, _ := range operazioni {
		date = append(date, k)
	}
	
	sort.Slice(date, func(i, j int) bool { 
		if date[i].a < date[j].a {
			return true
		} else if date[i].a > date[j].a {
			return false
		} else {
			if date[i].m < date[j].m {
				return true
			} else if date[i].m > date[j].m {
				return false
			} else {
				if date[i].g < date[j].g {
					return true
				} else if date[i].g > date[j].g {
					return false
				} else {
					return true
				}
			}
		}
	})

	// Output su schermo
	fmt.Printf("SALDO INIZIALE:%+19s\n\n", fmt.Sprintf("%0.2f", saldo))
	for i := 0; i < len(date); i++ {
		fmt.Printf("DATA: %s\n", dataToString(date[i]))
		saldoGiornaliero := 0.0
		for _, op := range operazioni[date[i]] {
			if op.tipo == 'P' {
				fmt.Printf("Prelievo:%+25s\n", fmt.Sprintf("%.2f", op.importo))
				saldoGiornaliero -= op.importo
			} else {
				fmt.Printf("Versamento:%+23s\n", fmt.Sprintf("%.2f", op.importo))
				saldoGiornaliero += op.importo
			}
		}
		saldo += saldoGiornaliero
		
		fmt.Printf("%+34s\n", "___________")
		fmt.Printf("SALDO GIORNALIERO: %+15s\n", fmt.Sprintf("%.2f", saldoGiornaliero))
		fmt.Printf("%+34s\n\n", "===========")
	}
	
	//fmt.Println(date)
	//fmt.Printf("%#v\n", operazioni)
	fmt.Printf("SALDO FINALE:%+21s\n", fmt.Sprintf("%0.2f", saldo))
	fmt.Printf("%+34s\n\n", "===========")
}

func stringToData(s string) data {
	d := data{}
	tmp := strings.Split(s, "/")

	d.m, _ = strconv.Atoi(tmp[1])
	if len(tmp[0]) == 4 {
		d.g, _ = strconv.Atoi(tmp[2])
		d.a, _ = strconv.Atoi(tmp[0])
	} else {
		d.g, _ = strconv.Atoi(tmp[0])
		d.a, _ = strconv.Atoi(tmp[2])
	}

	return d
}

func dataToString(d data) string {
	return fmt.Sprintf("%02d-%02d-%04d", d.g, d.m, d.a)
}
