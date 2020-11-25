package main

import (
	"fmt"
	"os"
)

func èPalindroma(r []rune) bool {
	if len(r) <= 1 {
		return true
	}

	if r[0] != r[len(r)-1] {
		return false
	}

	return èPalindroma(r[1 : len(r)-1])
}

func main() {
	s := os.Args[1]
	if èPalindroma([]rune(s)) {
		fmt.Println(s, "è palindroma")
	} else {
		fmt.Println(s, "non è palindroma")
	}
}
