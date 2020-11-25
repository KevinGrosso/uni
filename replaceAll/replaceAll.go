package main

import (
	"fmt"
	"os"
)

func replaceAll(s string, o, n rune) string {
	rs := []rune(s)

	for i := 0; i < len(rs); i++ {
		if rs[i] == o {
			rs[i] = n
		}
	}
	return string(rs)
}

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("Sintassi: %s [stringa] [car vecchio] [carattere nuovo]",
			os.Args[0])
		return
	}

	s := os.Args[1]
	o := []rune(os.Args[2])[0]
	n := []rune(os.Args[3])[0]

	fmt.Println(replaceAll(s, o, n))
}
