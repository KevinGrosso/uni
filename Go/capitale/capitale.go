package main 

import "fmt" 

/* 
	dato un capitale iniziale ed un tasso di interesse (in % intera), 
	visualizzare una tabella con la rivalutazione del capitale nei primi
	quattro anni
*/

func main() {
	var capitale, tassoPerc float64
	var tasso int
	
	fmt.Print("Inserisci il capitale iniziale: ")
	fmt.Scan(&capitale)
	fmt.Print("Inserisci il tasso di interesse (in % intera): ")
	fmt.Scan(&tasso)
	
	tassoPerc = float64(tasso) / 100
	
	fmt.Println("Anno\t\tCapitale")	
	
	capitale += capitale * tassoPerc
	fmt.Println("1\t\t", capitale)
	
	capitale += capitale * tassoPerc
	fmt.Println("2\t\t", capitale)
	
	capitale += capitale * tassoPerc
	fmt.Println("3\t\t", capitale)
	
	capitale += capitale * tassoPerc
	fmt.Println("4\t\t", capitale)
}
