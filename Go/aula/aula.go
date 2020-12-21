package main

import (
    "fmt"
)

/* 
	Un'aula é formata da un certo numero di file, ognuna composta da 
	un certo numero di banchi. I banchi valgono true se sono occupati,
	false altrimenti.
*/
type Aula [][]bool

/* 
	Crea un'aula delle dimensioni specificate.
	Ritorna l'aula e il valore true se le dimensioni sono corrette 
	(entrambe > 1) ritorna un'aula qualsiasi e false altrimenti.
	In un'aula appena creata tutti i banchi sono liberi.
*/
func creaAula(nfile, ncol int) (Aula, bool) {
	var a Aula
	
	if nfile < 1 || ncol < 1 {
		return a, false
	}
	
	a = make([][]bool, nfile)
	for i := 0; i < nfile; i++ {
		a[i] = make([]bool, ncol)
	}
	
	return a, true
}

/*
	Stampa la disposizione di un'aula come una griglia. Le 'x' 
	rappesentano i posti occupati, mentre i '_' rappresentano 
	quelli liberi.
*/
func stampaAula(a Aula) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] {
				fmt.Print(" x")
			} else {
				fmt.Print(" _")
			}
		}
		fmt.Println()
	}
}

/*
	occupa un banco di un'aula corrispondente a una certa posizione.
	Restituisce false se l'operazione non ha esito positivo, ovvero se
	lo stato complessivo dell'aula non cambia.
	Le file e le colonne sono numerate partendo da 0.
*/
func occupa(a Aula, fila, col int) bool {
	if fila >= 0 && fila < len(a) && col >= 0 && col < len(a[fila]) {
		if !a[fila][col] {
			a[fila][col] = true
			return true
		}
		return false
	}
	return false
}

/*
	libera un banco di un'aula che corrisponde a una certa posizione.
	Restituisce false se l'operazione non ha esito positivo, ovvero se
	lo stato complessivo dell'aula non cambia.
	Le file e le colonne sono numerate partendo da 0.
*/
func libera(a Aula, fila, col int) bool {
	if fila >= 0 && fila < len(a) && col >= 0 && col < len(a[fila]) {
		if a[fila][col] {
			a[fila][col] = false
			return true
		}
		return false
	}
	return false
}

/*
	verifica il distanziamento di un banco che corrisponde a una certa
	posizione. 
	Criterio di distanziamento: nessun banco occupato nelle posizioni
	immediatamente adiacenti che si trovano sulla stella fila/colonna.
	Restituisce true se e solo se il banco specificato rispetta il 
	criterio di distanziamento
*/
func verificaDistanziamento(a Aula, fila, col int) bool {
	if fila >= 0 && fila < len(a) && col >= 0 && col < len(a[fila]) {
		if fila == 0 {
			if col == 0 {
				return !(a[0][1] || a[1][0])
			} else if col == len(a[fila]) - 1 {
				return !(a[0][col-1] || a[1][col])
			}
			return !(a[0][col-1] || a[0][col+1] || a[1][col])
		} else if fila == len(a) - 1 {
			if col == 0 {
				return !(a[fila][1] || a[fila-1][0])
			} else if col == len(a[fila]) - 1 {
				return !(a[fila][col-1] || a[fila-1][col])
			}
			return !(a[fila][col-1] || a[fila][col+1] || 
					 a[fila-1][col])
		} else if col == 0 {
			return !(a[fila-1][0] || a[fila+1][0] || a[fila][1])
		} else if col == len(a[fila]) - 1 {
			return !(a[fila-1][col] || a[fila+1][col] || 
					 a[fila][col-1])
		}
		return !(a[fila-1][col] || a[fila+1][col] || 
				a[fila][col+1] || a[fila][col-1])
	}
	return false
}

func main() {
	a, _ := creaAula(5, 5)
	
	for i := 0; i < len(a); i ++ {
		for j := 0; j < len(a[i]); j++ {
			if verificaDistanziamento(a, i, j) {
				occupa(a, i, j)
			}
		}
	}
	
	stampaAula(a)
	fmt.Println()
	
	libera(a, 1, 1)
	stampaAula(a)
}