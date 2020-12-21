package main

import (
	"fmt"
	"os"
)

func èPalindroma(r []rune) bool {
	n := len(r)

	for i := 0; i < n/2; i++ {
		if r[i] != r[n-i-1] {
			return false
		}
	}

	return true
}

func main() {
	s := os.Args[1]
	if èPalindroma([]rune(s)) {
		fmt.Println(s, "è palindroma")
	} else {
		fmt.Println(s, "non è palindroma")
	}
}
