package main

import "fmt"

func main() {
	fmt.Println(isPrime(2))
	fmt.Println(isPrime(1))
	fmt.Println(isPrime(4))
	fmt.Println(isPrime(9))
	fmt.Println(isPrime(17))
}

func isPrime(n int) bool {
	if n > 1 {
		var i int
		for i = 2; i <= n/i; i++ {
			if n%i == 0 {
				break
			}
		}

		if i > n/i {
			return true
		}
		return false
	}
	return false
}
