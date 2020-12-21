package main

import (
    "fmt"
    "os"
)

const FILE_PATH = "fileoutput.txt"

func main() {
	file, err := os.Create(FILE_PATH)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERRORE: %v", err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Fprintf(file, "Prima riga\nééééé\nTerza riga\n\nRiga finale")
}
