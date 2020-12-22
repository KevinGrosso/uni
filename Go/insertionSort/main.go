package main

import (
	"fmt"
)

func main() {
	a := []int{10, 12, 70, 22, 13, -2, 82, 21, 12}
	fmt.Println(a)

	sort(a)
	fmt.Println(a)
}

func sort(a []int) {
	var j int

	for i := 0; i < len(a); i++ {
		n := a[i]
		for j = 0; j < i; j++ {
			if a[i] < a[j] {
				break
			}
		}

		for k := i-1; k >= j; k-- {
			a[k + 1] = a[k]
		}

		a[j] = n
	}
}
