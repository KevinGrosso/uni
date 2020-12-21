package main

import (
	"fmt"
	"io/ioutil"
	"io"
	"os"
	"bufio"
)

const FILE_PATH = "provainput.txt"

func main() {
	fmt.Print("Input con il primo metodo:\n")

	// Metodo numero 1 (da non usare)
	b, err := ioutil.ReadFile(FILE_PATH)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERRORE: %v\n", err)
		os.Exit(1)
	}
	fmt.Print(string(b))
	b = nil

	fmt.Print("\n\nInput con il secondo metodo:\n")
	
	// Metodo numero 2 (preferibile per diversi motivi, leggo tot byte alla volta)
	file, err := os.Open(FILE_PATH)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERRORE: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	buffer := make([]byte, 10) // Leggo 10 byte alla volta, fino all'EOF
	for {
		n, err := file.Read(buffer)
		if n > 0 {
			fmt.Print(string(buffer[:n]))
		}
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "ERRORE: %v\n", err)
			os.Exit(1)
		}
	}
	file.Close() // non serve in un programma normale, é solo per poter usare il terzo metodo
	
	fmt.Print("\n\nInput con il terzo metodo:\n")
	
	// Metodo numero 3 (piú comune, legge fino a un delimitatore)
	file, err = os.Open(FILE_PATH)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERRORE: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		// Leggo una stringa alla volta, la stringa é delimitata dal byte '\n'
		s, err := reader.ReadString('\n')
		
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "ERRORE: %v", err)
			os.Exit(1)
		}
		
		fmt.Print(s)
	}

	file.Close() // non serve in un programma normale, é solo per poter usare il quarto metodo
		
	fmt.Print("\n\nInput con il quarto metodo:\n")
		
	// Metodo numero 4 (piú facile da usare, legge riga per riga)
	file, err = os.Open(FILE_PATH)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERRORE: %v", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(s)
	}

	if scanner.Err() != nil {
		fmt.Fprintf(os.Stderr, "ERRORE: %v\n", err)
		os.Exit(1)
	}
}
