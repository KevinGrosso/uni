package main

import (
    "fmt"
)

func ripetizioni(s []int) bool {
	insieme := make(map [int]bool)
	for i := 0; i < len(s); i++ {
		if insieme[s[i]] {
			return true
		} else {
			insieme[s[i]] = true	
		}
	}
	return false
}

func main() {
	s := []int{10, 20, 20, 10, 50}
	fmt.Printf("Ci sono ripetizioni nella slice %v? %v", s, ripetizioni(s))
}
