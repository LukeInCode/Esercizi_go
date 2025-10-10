package main

import "fmt"

func main() {
	//dichiarare qui due variabili x1 e x2 di tipo float64
	var x1,x2 int;

	fmt.Print("fornire due numeri float: ")
	fmt.Scan(&x1, &x2)
	//inserire qui le istruzioni per svolgere le operazioni
	//salvando i risultati con degli short assignment
	// (i nomi delle variabili da utilizzare per i risultati sono quelli usati nelle istruzioni di stampa qui sotto, cioè somma, differenza, prodotto, quoziente, )
	somma := x1+x2
	differenza := x1-x2
	prodotto := x1 * x2
	quoziente := float64(x1) / float64(x2)
	// osservate qui sotto i diversi formati di argomenti passati alla funzione fmt.Println
	fmt.Println(x1, "+", x2, "=", somma)
	fmt.Println("differenza =", differenza)
	fmt.Println("il prodotto di", x1, "e", x2, "dà", prodotto)
	fmt.Println(x1, "/", x2, "=", quoziente)
}