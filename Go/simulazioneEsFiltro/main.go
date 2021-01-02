package main

import (
	"fmt"
)

func main() {
	var s string

	fmt.Scan(&s)
	lastN := '9'
	ordinata := true
	
	for _, r := range s {
		if r >= '0' && r <= '9' {
			if r > lastN {
				ordinata = false
				break
			}
			lastN = r
		}
	}

	if ordinata {
		fmt.Println("Sequenza nascosta ordinata.")
	} else {
		fmt.Println("Sequenza nascosta non ordinata.")
	}
}
