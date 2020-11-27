package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const lunghezzaCombinazione = 5
const caratteri = "abcdefghijklmnopqrstuvwxyz"
const numeroCaratteri = 26

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

//xvmsz
func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println("Caricamento...")
	combinazioni := generaCombinazioni([]rune(caratteri), lunghezzaCombinazione)
	fmt.Println("Caricamento completato\n")

	n := rand.Intn(int(math.Pow(26, 5)))
	secret := (*combinazioni)[n]
	fmt.Println(secret)
	combinazioni = nil

	var tentativo string

	for {
		fmt.Print("Prova ad indovinare (-1 per uscire) -> ")
		fmt.Scan(&tentativo)
		if tentativo == "-1" {
			break
		} else if len(tentativo) != 5 {
			fmt.Println("La combinazione segreta è lunga", lunghezzaCombinazione, "caratteri!")
			continue
		}

		if secret == tentativo {
			fmt.Println("Hai vinto!")
			break
		} else {
			inPosizione := 0
			tentativoRimasto := ""
			secretRimasto := ""

			for i := 0; i < lunghezzaCombinazione; i++ {
				if secret[i] == tentativo[i] { // aaghj abcde
					inPosizione++
				} else {
					secretRimasto += string(secret[i])
					tentativoRimasto += string(tentativo[i])
				}
			}

			giusti := 0
			for i := 0; i < len(secretRimasto); i++ {
				for j := 0; j < len(tentativoRimasto); j++ {
					if secretRimasto[i] == tentativoRimasto[j] {
						giusti++
						tmp := []rune(tentativoRimasto)
						tmp[j] = ' '
						tentativoRimasto = string(tmp)
						break
					}
				}
			}

			fmt.Println("I caratteri giusti al posto giusto ->", inPosizione)
			fmt.Println("I caratteri giusti al posto sbagliato ->", giusti)
			fmt.Println()
		}
	}

	fmt.Println("Grazie per aver giocato")
}
