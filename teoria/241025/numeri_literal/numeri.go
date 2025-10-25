package main

import (
	"fmt"
	"os"
)

/* This function, given a positive integer x, returns the Italian literal form of the number x. */
func n2s(x int) string {
	var s string
	if x >= 10 && x <= 19 {
		switch x {
		case 10:
			return "dieci"
		case 11:
			return "undici"
		case 12:
			return "dodici"
		case 13:
			return "tredici"
		case 14:
			return "quattordici"
		case 15:
			return "quindici"
		case 16:
			return "sedici"
		case 17:
			return "diciassette"
		case 18:
			return "diciotto"
		case 19:
			return "diciannove"
		default:
			return "can't happen"
		}
	} else if x >= 100 && x <= 999 {
		switch x / 100 {
		case 1:
			s += "cento"
		case 2:
			s += "duecento"
		case 3:
			s += "trecento"
		case 4:
			s += "quattrocento"
		case 5:
			s += "cinquecento"
		case 6:
			s += "seicento"
		case 7:
			s += "settecento"
		case 8:
			s += "ottocento"
		case 9:
			s += "novecento"
		}

		unita := x % 10
		decine := (x % 100) / 10

		switch decine {
		case 0:
			s += ""
		case 1:
			switch unita {
			case 0:
				s += "dieci"
			case 1:
				s += "undici"
			case 2:
				s += "dodici"
			case 3:
				s += "tredici"
			case 4:
				s += "quattordici"
			case 5:
				s += "quindici"
			case 6:
				s += "sedici"
			case 7:
				s += "diciassette"
			case 8:
				s += "diciotto"
			case 9:
				s += "diciannove"
			}
		case 2:
			s += "venti"
		case 3:
			s += "trenta"
		case 4:
			s += "quaranta"
		case 5:
			s += "cinquanta"
		case 6:
			s += "sessanta"
		case 7:
			s += "settanta"
		case 8:
			s += "ottanta"
		case 9:
			s += "novanta"
		}

		if unita == 1 || unita == 8 {
			s = s[:len(s)-1]
		}

		switch unita {
		case 0:
			s += ""
		case 1:
      if decine != 1 {
        s += "uno"
			}
		case 2:
			s += "due"
		case 3:
			s += "tre"
		case 4:
			s += "quattro"
		case 5:
			s += "cinque"
		case 6:
			s += "sei"
		case 7:
			s += "sette"
		case 8:
			s += "otto"
		case 9:
			s += "nove"
		}

		return s
	} else { // x is either <10 or >=20
		switch x / 10 {
		case 0:
			s = ""
		case 2:
			s = "venti"
		case 3:
			s = "trenta"
		case 4:
			s = "quaranta"
		case 5:
			s = "cinquanta"
		case 6:
			s = "sessanta"
		case 7:
			s = "settanta"
		case 8:
			s = "ottanta"
		case 9:
			s = "novanta"
		}

		if x > 10 && (x%10 == 1 || x%10 == 8) {
			s = s[:len(s)-1]
		}

		switch x % 10 {
		case 1:
			s += "uno"
		case 2:
			s += "due"
		case 3:
			s += "tre"
		case 4:
			s += "quattro"
		case 5:
			s += "cinque"
		case 6:
			s += "sei"
		case 7:
			s += "sette"
		case 8:
			s += "otto"
		case 9:
			s += "nove"
		}

		return s
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	if n <= 0 {
		fmt.Println("Assegni di questo valore non esistono!")
		os.Exit(1)
	}
	s := n2s(n)
	fmt.Println(s)
}

