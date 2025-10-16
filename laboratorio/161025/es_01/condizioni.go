package main

import "fmt"

// aggiungere dopo ogni fmt.Scan il codice occorrente
func main() {
	var n int
	var x float64

	fmt.Print("int uguale a 10: ")
	fmt.Scan(&n)
	fmt.Println(n == 10)

	fmt.Print("int diverso da 10: ")
	fmt.Scan(&n)
	fmt.Println(n != 10)

	fmt.Print("int diverso da 10 e da 20: ")
	fmt.Scan(&n)
	fmt.Println(n != 10 && n != 20)

	fmt.Print("int diverso da 10 o da 20: ")
	fmt.Scan(&n)
	fmt.Println(n != 10 || n != 20)

	fmt.Print("int maggiore o uguale a 10: ")
	fmt.Scan(&n)
	fmt.Println(n >= 10)

	fmt.Print("int compreso tra 10 e 20, nell'intervallo [10,20]: ")
	fmt.Scan(&n)
	fmt.Println(n >= 10 && n <= 20)

	fmt.Print("int compreso tra 10 e 20, nell'intervallo (10,20): ")
	fmt.Scan(&n)
	fmt.Println(n > 10 && n < 20)

	fmt.Print("int compreso tra 10 e 20, nell'intervallo [10,20): ")
	fmt.Scan(&n)
	fmt.Println(n >= 10 && n < 20)

	fmt.Print("int esterno all'intervallo [10,20]: ")
	fmt.Scan(&n)
	fmt.Println(n <= 10 || n >= 20)

	fmt.Print("int tra 10 e 20 (nell'intervallo [10,20]) o tra 30 e 40 ([30,40]) ")
	fmt.Scan(&n)
	fmt.Println(n >= 10 && n <= 20 || n >= 30 && n <= 40)

	fmt.Print("int multiplo di 4 ma non di 100: ")
	fmt.Scan(&n)
	fmt.Println(n % 4 == 0 && n % 100 != 0)

	fmt.Print("int dispari e compreso tra 0 e 100 ([0,100]): ")
	fmt.Scan(&n)
	fmt.Println(n % 2 != 0 && n >= 0 && n <= 100)

	fmt.Print("float64 vicino a 10.0 (non più lontano di 10^-6): ")
	fmt.Scan(&x)
	fmt.Println(x >= 10 - 1e-6 || x <= 10 + 1e-6);
}
