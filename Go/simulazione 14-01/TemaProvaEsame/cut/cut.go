package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inizioCol, _ := strconv.Atoi(os.Args[1])
	fineCol, _ := strconv.Atoi(os.Args[2])

	inizioCol--
	fineCol--

	path := os.Args[3]
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		riga := scanner.Text()
		var fine int
		if fineCol > len(riga) {
			fine = len(riga)
		} else {
			fine = fineCol + 1
		}
		if len(riga) > inizioCol {
			fmt.Println(riga[inizioCol:fine])
		}
	}
}
