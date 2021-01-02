package main

import (
	"fmt"
	"unicode"
)

func main() {
	var pwd string
	
	var length bool // true se la lunghezza va bene
	var lower int
	var upper int
	var decimal int
	var special int
	
	fmt.Scan(&pwd)

	if len(pwd) >= 12 {
		length = true
	}

	for _, r := range pwd {
		if unicode.IsLower(r) {
			lower++
		} else if unicode.IsUpper(r) {
			upper++
		} else if unicode.IsDigit(r) {
			decimal++
		} else if !unicode.IsLetter(r) {
			special++
		}
	}

	if length && lower >= 2 && upper >= 2 && decimal >= 3 && special >= 4 {
		fmt.Println("La pw è ben definita!")
	} else {
		fmt.Println("La pw non è definita correttamente:")
		if !length {
			fmt.Println("- La pw deve avere una lunghezza minima di 12 caratteri")
		}
		if lower < 2 {
			fmt.Println("- Almeno 2 caratteri della pw devono rappresentare delle lettere minuscole")
		}
		if upper < 2 {
			fmt.Println("- Almeno 2 caratteri della pw devono rappresentare delle lettere maiuscole")
		}
		if decimal < 3 {
			fmt.Println("- Almeno 3 caratteri della pw devono rappresentare delle cifre decimali")
		}
		if special < 4 {
			fmt.Println("- Almeno 4 caratteri della pw non devono rappresentare lettere o cifre decimali")
		}
	}
}
