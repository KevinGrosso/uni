package main 

import "fmt" 

// potenza n-esima di x

func main() {
	var esp, base int

	fmt.Print("Base: ")
	fmt.Scan(&base)
	fmt.Print("Esponente: ")
	fmt.Scan(&esp)
	
	potenza := 1
	
	for i := 0; i < esp; i++ {
		potenza *= base
	}
	
	fmt.Println(base, "^", esp, "=", potenza)
}
