package main

import (
    "fmt"
    "os"
    "bufio"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Sintassi: %s <percorso del file>\n", os.Args[0])
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprint(os.Stderr, "ERRORE: File non trovato\n")
		os.Exit(1)
	}
	defer file.Close()

	emptyFlag := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			fmt.Println(scanner.Text())
			emptyFlag = false
		} else if !emptyFlag {
			fmt.Println()
			emptyFlag = true
		}
	}
}
