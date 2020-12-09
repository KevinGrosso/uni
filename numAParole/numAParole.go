package main

import "fmt"

// dato un numero, lo stampa in lettere

func main() {
	var n int

	fmt.Print("Numero da convertire: ")
	fmt.Scan(&n)

	// centinaia
	var o rune
	switch n / 100 {
	case 0:
		o = 'o'
	case 2:
		fmt.Print("due")
	case 3:
		fmt.Print("tre")
	case 4:
		fmt.Print("quattro")
	case 5:
		fmt.Print("cinque")
	case 6:
		fmt.Print("sei")
	case 7:
		fmt.Print("sette")
	case 8:
		fmt.Print("otto")
	case 9:
		fmt.Print("nove")
	}

	if n/100 != 0 {
		fmt.Print("cento")
	}

	// casi speciali

	if n == 0 || n%100 >= 10 && n%100 <= 19 {
		switch n % 100 {
		case 0:
			fmt.Print("zero")
		case 10:
			fmt.Print("dieci")
		case 11:
			fmt.Print("undici")
		case 12:
			fmt.Print("dodici")
		case 13:
			fmt.Print("tredici")
		case 14:
			fmt.Print("quattordici")
		case 15:
			fmt.Print("quindici")
		case 16:
			fmt.Print("sedici")
		case 17:
			fmt.Print("diciassette")
		case 18:
			fmt.Print("diciotto")
		case 19:
			fmt.Print("diciannove")
		}
	} else {
		// decine

		vocale := 'a'

		switch (n / 10) % 10 {
		case 2:
			fmt.Print("vent")
			vocale = 'i'
		case 3:
			fmt.Print("trent")
		case 4:
			fmt.Print("quarant")
		case 5:
			fmt.Print("cinquant")
		case 6:
			fmt.Print("sessant")
		case 7:
			fmt.Print("settant")
		case 8:
			fmt.Print(string(o) + "ttant")
		case 9:
			fmt.Print("novant")
		}

		if n%10 != 1 && n%10 != 8 && (n/10)%10 != 1 && (n/10)%10 != 0 {
			fmt.Print(string(vocale))
		}

		// unità
		if n < 10 || n > 20 {
			switch n % 10 {
			case 1:
				fmt.Print("uno")
			case 2:
				fmt.Print("due")
			case 3:
				fmt.Print("tre")
			case 4:
				fmt.Print("quattro")
			case 5:
				fmt.Print("cinque")
			case 6:
				fmt.Print("sei")
			case 7:
				fmt.Print("sette")
			case 8:
				fmt.Print("otto")
			case 9:
				fmt.Print("nove")
			}
		}
	}
	fmt.Println()
}
