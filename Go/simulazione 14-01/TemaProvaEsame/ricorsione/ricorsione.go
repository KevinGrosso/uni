package main

import (
	"os"
	"fmt"
)

func main() {
	pos := trova(os.Args[1][0], os.Args[2])
	if pos != -1 {
		fmt.Printf("%c si trova in posizione %d\n", os.Args[1][0], pos)
	} else {
		fmt.Printf("%c non é presente nella stringa\n", os.Args[1][0])
	}
}

func trova(c byte, s string) int {
	if s == "" {
		return -1
	}
	if s[0] == c {
		return 0
	}

	pos := trova(c, s[1:])
	if pos == -1 {
		return -1
	}
	
	return 1 + pos
}
