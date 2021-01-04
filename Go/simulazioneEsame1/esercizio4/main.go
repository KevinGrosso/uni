package main

import (
	"fmt"
	"os"
	"strings"
	"sort"
)

func main() {
	s := make([]string, len(os.Args) - 1)
	copy(s, os.Args[1:])
	m := sottosequenze(s)
	
	i := 0
	chiavi := make([]string, len(m))
	for k, _ := range m {
		chiavi[i] = k
		i++
	}

	sort.Slice(chiavi, func(i, j int) bool {
		if len(chiavi[i]) < len(chiavi[j]) {
			return false
		}
		return true
	})
	
	for _, v := range chiavi {
		fmt.Printf("%s -> Occorrenze: %d\n", v, m[v])
	}
}

func sottosequenze(s []string) map[string]int {
	m := make(map[string]int)

	for i := 0; i < len(s); i++ {
		for j := len(s)-1; j > i + 1; j-- {
			if s[i] == s[j] {
				m[strings.Join(s[i:j+1], " ")]++
			}
		}
	}
	
	return m
}
