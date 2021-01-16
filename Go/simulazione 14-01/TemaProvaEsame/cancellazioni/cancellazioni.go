package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Fornire 1 nome di file")
		return // os.Exit(1)
	}

	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("File non accessibile")
		return // os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lista []string
	for scanner.Scan() {
		riga := scanner.Text()
		lista = append(lista, strings.Split(riga, " ")...)
	}

	fmt.Println(lista)
	fmt.Println(cancella(lista))
}

func cancella(lista []string) []string {
	var listaRisultato []string
	for i := 0; i < len(lista); i++ {
		if n, éNumero := strconv.Atoi(lista[i]); éNumero == nil {
			i += n
		} else {
			listaRisultato = append(listaRisultato, lista[i])
		}
	}
	return listaRisultato
}
