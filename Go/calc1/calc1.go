package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 4 {
		n1, err1 := strconv.Atoi(os.Args[1])
		op := os.Args[2]
		n2, err2 := strconv.Atoi(os.Args[3])

		if err1 != nil || err2 != nil {
			fmt.Printf("Errore: valore non numerico")
			return
		}

		switch op {
		case "+":
			fmt.Println(n1 + n2)
		case "-":
			fmt.Println(n1 - n2)
		case "x":
			fmt.Println(n1 * n2)
		case "/":
			fmt.Println(float64(n1) / float64(n2))
		default:
			fmt.Println("Errore: operazione non valida (gli operatori disponibili sono + - x /)")
		}
	} else {
		fmt.Printf("Sintassi: %s n1 [+-x/] n2", os.Args[0])
	}
}
