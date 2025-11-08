/*

In un file operazioni.go definire due versioni di una stessa funzione operazioni, una versione con variabili di ritorno nominate (operazioniV1) e una senza (operazioniV1), che accettano due interi e restituiscono somma, prodotto e differenza in quest'ordine:

- operazioniV1(n1, n2 int) (int, int, int) //con variabili di ritorno nominate

- operazioniV2(n1, n2 int) (int, int, int) //senza variabili di ritorno nominate

Scrivere un main che legge da standard input due int e invoca le due funzioni per testarle.

nomefile: operazioni.go

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

var prog = "./operazioni"
var diffwidth = 120

func TestMinimaleV1(t *testing.T) {
	s, p, d := operazioniV1(2, 3)
	if s != 5 || p != 6 || d != -1 {
		t.Fail()
	}
}

func TestMinimaleV2(t *testing.T) {
	s, p, d := operazioniV2(2, 3)
	if s != 5 || p != 6 || d != -1 {
		t.Fail()
	}
}

func TestNegativiV1(t *testing.T) {
	s, p, d := operazioniV1(-2, -3)
	if s != -5 || p != 6 || d != 1 {
		t.Fail()
	}
}

func TestNegativiV2(t *testing.T) {
	s, p, d := operazioniV2(-2, -3)
	if s != -5 || p != 6 || d != 1 {
		t.Fail()
	}
}

func TestMain1(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"57 999\n",
		"1056 56943 -942\n1056 56943 -942\n")
}

func TestMain2(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"72346823 -374673\n",
		"71972150 -27106401213879 72721496\n71972150 -27106401213879 72721496\n")
}
