package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n := os.Args[1]
	d, _ := strconv.Atoi(os.Args[2])

	fmt.Println(generaNumeri(n, d, 0, len(n)-d))
}

func generaNumeri(n string, d int, l int, l2 int) string {
	min := "9"

	if len(n) < d+1 {
		if l < l2 {
			for i := 0; i < len(n); i++ {
				if string(n[i]) < min {
					min = string(n[i])
				}
			}
		} else {
			return ""
		}
		return min
	}

	numero := ""
	minI := 0

	for i := 0; i < d+1; i++ {
		if string(n[i]) < min {
			min = string(n[i])
			minI = i
		}
	}
	l++
	numero += min + generaNumeri(n[minI+1:], d, l, l2)

	return numero
}
