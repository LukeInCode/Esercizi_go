/*
In questo esercizio prendiamo in considerazione solo caratteri ASCII (byte).
In un file funzioni.go definire le seguenti funzioni:

- hasUpper(s string) bool
La funzione riceve una stringa come parametro e restituisce true se la stringa contiene almeno una lettera latina maiuscola (tra 'A' e 'Z'), false altrimenti.

- firstUpper(s string) int
La funzione riceve una stringa come parametro e restituisce la posizione della prima lettera latina maiuscola (tra 'A' e 'Z'), se la stringa ne contiene almeno una, -1 altrimenti.

- lastUpper(s string) int
La funzione riceve una stringa come parametro e restituisce la posizione dell'ultima lettera latina maiuscola (tra 'A' e 'Z'), se la stringa ne contiene almeno una, -1 altrimenti.

- countDigitsLettersOthers(s string) int int int
La funzione riceve una stringa come parametro, conta quante cifre, quante lettere e quanti altri caratteri (né cifre né lettere) contiene e restituisce i tre risultati in questo ordine.

- isPalindrome(s string) bool
La funzione riceve una stringa come parametro e restituisce true se la stringa è palindroma, e false altrimenti. Una parola è palindroma se può essere letta  sia da sinistra a destra che da destra a sinistra. Ad esempio "osso" e "ingegni" sono palindrome, ma anche "" (la stringa vuota) e "t" (le stringhe di un solo carattere).

Scrivete infine un main che legge una stringa da standard input, usa le funzioni qui sopra per determinare
- se la stringa letta contiene maiuscole,
- in che posizione è la prima maiuscola,
- in che posizione è l'ultima maiuscola,
- quante cifre, lettere e altri caratteri contiene,
- se è palindroma,
e stampa i risultati ("ha maiuscole" / "non ha maiuscole", "prima maiuscola in posizione ...", ..., "palindroma" / "non palindroma").

Esempi di esecuzione
---------------------
su input 321DOGMAIAMGOD123 il programma stampa:

ha maiuscole
prima maiuscola in posizione 3
ultima maiuscola in posizione 13
cifre, lettere, altro: 6 11 0
palindroma

su input unastringa il programma stampa:

non ha maiuscole
prima maiuscola in posizione -1
ultima maiuscola in posizione -1
cifre, lettere, altro: 0 10 0
non palindroma

*/

package main

import (
	//"fmt"
	//"log"
	//"os/exec"

	//"strings"
	"testing"
	//"bytes"
	//"io"
	//"os"
)

var prog = "./funzioni"
var diffwidth = 120

// palindromi
func TestPalindromeSingleChar(t *testing.T) {
	if !isPalindrome("A") {
		t.Fail()
	}
}

func TestPalindromeSameChar(t *testing.T) {
	if !isPalindrome("bbbbbbbbbbbbbbbbbb") {
		t.Fail()
	}
}

func TestNonPalindromo(t *testing.T) {
	if isPalindrome("ciaoajlajdlas") {
		t.Fail()
	}
}

// hasupper
func TestHasUpperSingle(t *testing.T) {
	if !hasUpper("A") {
		t.Fail()
	}
	if hasUpper("b") {
		t.Fail()
	}
}

func TestHasUpperNormal(t *testing.T) {
	if !hasUpper("askjhdaksjsIlaskdjaslkdj") {
		t.Fail()
	}
}

// count
func TestCountNormal(t *testing.T) {
	d, l, o := countDigitsLettersOthers("sakdja87213612kjhkasjhdkas8q912as k askdhaksdhkasdkjas")
	//fmt.Println(d, " ", l, " ", o)
	if d != 12 && l != 40 && o != 2 {
		t.Fail()
	}
}

func TestCountEmpty(t *testing.T) {
	d, l, o := countDigitsLettersOthers("")
	if d != 0 && l != 0 && o != 0 {
		t.Fail()
	}
}

func TestCountOneOfEach(t *testing.T) {
	d, l, o := countDigitsLettersOthers("1a,")
	//fmt.Println(d, " ", l, " ", o)
	if d != 1 && l != 1 && o != 1 {
		t.Fail()
	}
}

// firstupper
func TestFirstUpper(t *testing.T) {
	if firstUpper("udtqABsso") != 4 {
		t.Fail()
	}
}
func TestFirstUpperNo(t *testing.T) {
	if firstUpper("udtqddsso") != -1 {
		t.Fail()
	}
}

// lastupper
func TestLastUpper(t *testing.T) {
	if lastUpper("udtqAvBsso") != 6 {
		t.Fail()
	}
}

func TestMainUnaStringa(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"unastringa\n",
		"non ha maiuscole\nprima maiuscola in posizione -1\nultima maiuscola in posizione -1\ncifre, lettere, altro: 0 10 0\nnon palindroma\n")
}

func TestMainPalindroma(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"abba\n",
		"non ha maiuscole\nprima maiuscola in posizione -1\nultima maiuscola in posizione -1\ncifre, lettere, altro: 0 4 0\npalindroma\n")
}

func TestMainPalindromaDogma(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"DOGMAIAMGOD\n",
		"ha maiuscole\nprima maiuscola in posizione 0\nultima maiuscola in posizione 10\ncifre, lettere, altro: 0 11 0\npalindroma\n")
}

func TestMainPalindromaNumeriLettere(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"321DOGMAIAMGOD123\n",
		"ha maiuscole\nprima maiuscola in posizione 3\nultima maiuscola in posizione 13\ncifre, lettere, altro: 6 11 0\npalindroma\n")
}
