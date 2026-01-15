/*

Carta prepagata
===============

Scrivere un programma 'prepagata.go' dotato di:

- una struttura 'Prepagata' che ha due campi, 'numero' di tipo int e 'saldo' di tipo float64, dichiarati in quest'ordine, che rappresentano il numero della carta prepagata e il saldo disponibile sulla carta stessa

- una funzione

    ricarica(carta *Prepagata, importo int)

  che carica importo sulla carta, aggiungendolo al saldo già disponibile. Se importo è negativo, non fa nulla.


- una funzione

    preleva(carta *Prepagata, importo int) (int, error)

  che preleva importo dalla carta, e restituisce la cifra prelevata e un errore;
  in particolare, se importo è maggiore o uguale a 0 e il saldo è maggiore o uguale all'importo,
  addebita importo sulla carta e restituisce importo e 'nil';
  altrimenti restituisce 0 e un errore "saldo insufficiente" in caso di mancanza di fondi o "importo non valido" in caso di un parametro negativo.


- un *metodo* String per Prepagata che restituisce una descrizione della carta nel seguente formato:

	carta n. <numero>, saldo <saldo>

  con due cifre decimali per il saldo. Ad esempio:

	carta n. 2341, saldo 200.15


- una funzione main() che crea una carta prepagata con numero 1709 e saldo 100, stampa un menu:
a. saldo
b. ricarica
c. prelievo
e. esci

e chiede ripetutamente all'utente un'opzione fino a che l'opzione non è 'e'.

Il main:
- se l'opzione è 'e', stampa "arrivederci" e termina;
- se l'opzione è 'a', stampa la situazione della carta (vedi il formato nelle specifiche di String());
- se l'opzione è 'b', legge da stdin l'importo digitato dall'utente e lo carica sulla carta, poi stampa la situazione della carta;
- se l'opzione è 'c', legge da stdin l'importo digitato dall'utente e lo addebita sulla carta, poi stampa la situazione della carta o un messaggio di errore;
- per ogni altra opzione, stampa "opzione non valida".

Suggerimenti:
1. per creare un errore, utilizzare la funzione New
   del package 'errors' (vedere la documentazione relativa).
2. per concatenare strighe e numeri, utilizzare le funzioni
   Sprint o Sprintf del package fmt (vedere la documentazione relativa).
3. per rappresentare un float con 2 cifre decimali, usare %.2f

Esempio
=======

data la seguente sequenza di opzioni in input:

a
c
200
b
500
c
200
c
-78
b
-86
f
e

$ ./prepagata

produce in output:

a. saldo
b. ricarica
c. prelievo
e. esci
carta n. 1709, saldo 100.00
saldo insufficiente
carta n. 1709, saldo 600.00
carta n. 1709, saldo 400.00
importo non valido
carta n. 1709, saldo 400.00
opzione non valida
arrivederci

*/

package main

import (
	"fmt"
	//"io/ioutil"
	"testing"
	//"log"
)

var prog = "./prepagata"
var diffwidth = 150

func TestMainMultiplo(t *testing.T) {
	ConfrontaConOracolo(
		t,
		prog,
		"simple.input")

	//lanciaEcontrolla("a\nc\n200\nb\n500\nc\n200\nf\ne\n", "a. saldo\nb. versamento\nc. prelievo\ne. esci\nopzione: carta n. 1601, saldo 100.00\nopzione: importo: puoi prelevare fino a 100 euro\nopzione: importo: carta n. 1601, saldo 600.00\nopzione: importo: carta n. 1601, saldo 400.00\nopzione: opzione non valida\nopzione: arrivederci\n", t)
}

func TestRicarica(t *testing.T) {
	var carta = Prepagata{1709, 100}
	ricarica(&carta, 200)

	if carta.saldo != 300 {
		t.Fail()
	}
}

func TestRicaricaNeg(t *testing.T) {
	var carta = Prepagata{1709, 300}
	ricarica(&carta, -350)

	if carta.saldo != 300 {
		t.Fail()
	}
}

func TestPrelevaNormale(t *testing.T) {
	var carta = Prepagata{1601, 100}
	fmt.Println(preleva(&carta, 20))

	if carta.saldo != 80 {
		t.Fail()
	}
}

func TestPrelevaTroppo(t *testing.T) {
	var carta = Prepagata{1601, 100}
	p, e := preleva(&carta, 200)
	fmt.Println(p, e)

	if carta.saldo != 100 {
		t.Fail()
	}

	if p != 0 {
		t.Fail()
	}

	if e == nil || e.Error() != "saldo insufficiente" {
		fmt.Println("manca errore o messaggio non corretto")
		t.Fail()
	}
}

func TestPrelevaNeg(t *testing.T) {
	var carta = Prepagata{1709, 100}
	p, e := preleva(&carta, -47)

	if carta.saldo != 100 {
		t.Fail()
	}

	if p != 0 {
		t.Fail()
	}

	if e == nil || e.Error() != "importo non valido" {
		fmt.Println("manca errore o messaggio non corretto")
		t.Fail()
	}

}

func TestString(t *testing.T) {
	carta := Prepagata{17, 20}
	fmt.Println(carta)

	s := carta.String()
	fmt.Println(s)

	if s != "carta n. 17, saldo 20.00" {
		t.Fail()
	}
}
