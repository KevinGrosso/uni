package main

import "fmt"

func insertAtIndex(slice []int, v, i int) []int {
	ris := make([]int, len(slice)+1)

	copy(ris, slice[:i])
	ris[i] = v
	copy(ris[i+1:], slice[i:])

	return ris
}

func main() {
	slice := []int{10, 20, 40, 50, 60}
	fmt.Println(slice)

	slice = insertAtIndex(slice, 30, 2)
	fmt.Println(slice)

	slice = insertAtIndex(slice, 0, 0)
	fmt.Println(slice)
}
