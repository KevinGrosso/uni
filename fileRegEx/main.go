package main

import (
    "fmt"
	"os"
	"encoding/json"
	"io/ioutil"
)

func main() {
	stud, err := parseFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	
	fmt.Fprintln(os.Stdout, stud)
	
	b, _ := json.Marshal(stud)
	ioutil.WriteFile(os.Args[2], b, 0644)
}