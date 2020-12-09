package main 

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	
	fmt.Print("$: ")
	fmt.Scan(&s)
	
	num, ok := numeroMisterioso(s)
	if ok {
		fmt.Println(num)
	} else {
		fmt.Println("Si è verificato un errore")
	}
}

func numeroMisterioso(s string) (int, bool) {
	var num string
	
	for _, r := range s {
		switch r {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			num += string(r)
		case '#':
		default:
			return 0, false
		}
	}
	x, err := strconv.Atoi(num)
	return x*2, err == nil
}