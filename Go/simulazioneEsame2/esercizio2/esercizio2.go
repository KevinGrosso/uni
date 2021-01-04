package main

import (
	"fmt"
)

const CHARS string = "0123456789"

func main() {
	var n, L, H int

	fmt.Scan(&n)
	fmt.Scan(&L)
	fmt.Scan(&H)

	charIndex := 0
	offset := L - 1
	inverted := true
	charIndex = stampaRigaOrizzontale(L, charIndex, false)
	for i := 0; i < n; i++ {
	
		charIndex = stampaRigaVerticale(H, charIndex, offset)
		charIndex = stampaRigaOrizzontale(L, charIndex, inverted)
		inverted = !inverted
		if inverted {
			offset = L - 1
		} else {
			offset = 0
		}
	}
}

func stampaRigaVerticale(length, charIndex, offset int) int {
	for i := 0; i < length - 2; i++ {
		for j := 0; j < offset; j++ {
			fmt.Print(" ")
		}
		fmt.Print(string(CHARS[charIndex%len(CHARS)]))
		charIndex++
		fmt.Println()
	}
	return charIndex
}

func stampaRigaOrizzontale(length, charIndex int, inverted bool) int {
	if !inverted {
		for i := 0; i < length; i++ {
			fmt.Print(string(CHARS[charIndex%len(CHARS)]))
			charIndex++
		}
		fmt.Println()
		return charIndex
	}
	
	charIndex += length - 1
	for i := 0; i < length ; i++ {
		fmt.Print(string(CHARS[charIndex%len(CHARS)]))
		charIndex--
	}
	fmt.Println()
	return charIndex + length + 1
}
