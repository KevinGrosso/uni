package main

import (
	"fmt"
)

func main() {
	var m map[string]int
	m = make(map[string]int)

	m2 := map[string]int{"Test": 22, "Ciao": 33}

	m["Kevin"] = 19
	m["Greta"] = 17

	fmt.Println(m)
	fmt.Println(m2)
	fmt.Println(m["Kevin"])
	fmt.Println(m2["Ciao"])
	fmt.Println(m2["Test"])
}
