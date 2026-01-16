/*
=== Appuntamenti ===

Scrivere un programma (il cui file deve chiamarsi 'appuntamenti.go') dotato di:

# una struttura

	Appuntamento
	con campi giorno, oraInizio, oraFine (dichiarati in quest'ordine) di tipo int

che rappresentano un giorno dell'anno, e l'ora di inizio e di fine  dell'appuntamento. Si considerano per semplicità solo appuntamenti che iniziano e finiscono nella stessa giornata e a ore precise (intere).

# una funzione

	NuovoAppuntamento(gg, h1, h2 int) (Appuntamento, bool)

che, se i parametri sono validi (giorno tra 1 e 366, ora inizio precedente o uguale a ora di fine, entrambe nell'intervallo [0,24]) crea un appuntamento corrispondente a questi dati e lo restituisce insieme al valore 'true', altrimenti restituisce un Appuntamento "zero-value" e 'false'.

# una funzione

	SovrapposizioneAppuntamenti(app1, app2 Appuntamento) bool

che, dati due appuntamenti che si assume siano validi, restituisce
- 'true' se si sovrappongono (anche parzialmente),
- 'false' se *NON* si sovrappongono.

NotaBene: due appuntamenti consecutivi, in cui cioè il primo dei due finisce all'ora in cui inizia il secondo, non sono sovrapposti.

# una funzione

	AggiungiAppuntamento(app Appuntamento, agenda *[]Appuntamento) bool

che, se l'appuntamento app non si sovrappone con NESSUNO degli altri appuntamenti già presenti in agenda, lo aggiunge alla stessa restituendo 'true', altrimenti non fa nulla e restituisce 'false'

# un **metodo** String() per Appuntamento che restituisce una stringa nella forma:

	[il giorno <giorno> dalle <ora inizio> alle <ora fine>: -------***--------------]

	(dove la stringa in fondo è la mappa oraria dell'appuntamento stesso, gli asterischi indicano la posizione all'interno delle 24h, si veda esempio sotto, le ore sono a due cifre sempre)

# una funzione main() che crea un'agenda vuota e legge da standard input un numero arbitrario di righe terminate da "end of file" (EOF, che su standard input, su sistemi Linux, si ottiene premendo invio seguito da CTRL D).

Le righe contengono ognuna tre interi che rappresentano, nell'ordine,

	giorno ora-inizio ora-fine

(come separatore si usano degli spazi, eventualmente ripetuti)

Per ciascuna riga il programma aggiunge in agenda il corrispondente appuntamento, se valido e non in sovrapposizione con altri.

Una volta raggiunto EOF, il programma visualizza su standard output gli appuntamenti inseriti in agenda, nell'ordine in cui sono stati inseriti, uno per riga, nel formato del metodo String.

Si assuma che da standard input siano inseriti solo numeri interi (non occorre fare controlli).

esempio di esecuzione
=====================

$ ./appuntamenti < due.input
[il giorno 200 dalle 11 alle 12: -----------*------------]
[il giorno 200 dalle 13 alle 15: -------------**---------]
[il giorno 200 dalle 12 alle 13: ------------*-----------]
[il giorno 201 dalle 11 alle 12: -----------*------------]
[il giorno 200 dalle 02 alle 04: --**--------------------]
[il giorno 200 dalle 18 alle 19: ------------------*-----]
[il giorno 180 dalle 23 alle 24: -----------------------*]
[il giorno 180 dalle 00 alle 01: *-----------------------]
*/
package main

import (
	//"strings"
	//"log"
	//"fmt"
	//"os/exec"
	//"strings"
	"fmt"
	"testing"
)

var prog = "./appuntamenti"
var diffwidth = 150

func TestAppuntamentoNorm(t *testing.T) {
	ap, correct := NuovoAppuntamento(300, 11, 12)
	fmt.Println(ap.giorno)
	fmt.Println(ap.oraInizio)
	fmt.Println(ap.oraFine)
	if !correct {
		t.Fail()
	}
}

func TestAppuntamentoWrongDay(t *testing.T) {
	_, correct := NuovoAppuntamento(400, 11, 12)
	if correct {
		t.Fail()
	}
}

func TestAppuntamentoWrongTime(t *testing.T) {
	_, correct := NuovoAppuntamento(300, 11, 32)
	if correct {
		t.Fail()
	}
}

func TestSovrAppuntamentoPost(t *testing.T) {
	uno, _ := NuovoAppuntamento(300, 11, 13)
	due, _ := NuovoAppuntamento(300, 12, 14)
	if !SovrapposizioneAppuntamenti(uno, due) {
		t.Fail()
	}
	if !SovrapposizioneAppuntamenti(due, uno) {
		t.Fail()
	}
}

func TestSovrAppuntamentoFake(t *testing.T) {
	uno, _ := NuovoAppuntamento(302, 11, 13)
	due, _ := NuovoAppuntamento(300, 11, 13)
	if SovrapposizioneAppuntamenti(uno, due) {
		t.Fail()
	}
	if SovrapposizioneAppuntamenti(due, uno) {
		t.Fail()
	}
}

func TestSovrAppuntamentoUguali(t *testing.T) {
	uno, _ := NuovoAppuntamento(300, 11, 13)
	due, _ := NuovoAppuntamento(300, 11, 13)
	if !SovrapposizioneAppuntamenti(uno, due) {
		t.Fail()
	}
	if !SovrapposizioneAppuntamenti(due, uno) {
		t.Fail()
	}
}

func TestSovrAppuntamentoPre(t *testing.T) {
	uno, _ := NuovoAppuntamento(300, 11, 13)
	due, _ := NuovoAppuntamento(300, 10, 12)
	if !SovrapposizioneAppuntamenti(uno, due) {
		t.Fail()
	}
	if !SovrapposizioneAppuntamenti(due, uno) {
		t.Fail()
	}
}

func TestSovrAppuntamentoPreSameEnd(t *testing.T) {
	uno, _ := NuovoAppuntamento(300, 11, 13)
	due, _ := NuovoAppuntamento(300, 10, 13)
	if !SovrapposizioneAppuntamenti(uno, due) {
		t.Fail()
	}
	if !SovrapposizioneAppuntamenti(due, uno) {
		t.Fail()
	}
}

func TestSovrAppuntamentoFakeContigui(t *testing.T) {
	uno, _ := NuovoAppuntamento(300, 11, 13)
	due, _ := NuovoAppuntamento(300, 13, 15)
	if SovrapposizioneAppuntamenti(uno, due) {
		t.Fail()
	}
	if SovrapposizioneAppuntamenti(due, uno) {
		t.Fail()
	}
}

func TestSovrAppuntamentoContenuto(t *testing.T) {
	uno, _ := NuovoAppuntamento(300, 11, 13)
	due, _ := NuovoAppuntamento(300, 10, 16)
	fmt.Println(uno)
	fmt.Println(due)
	if !SovrapposizioneAppuntamenti(uno, due) {
		t.Fail()
	}
	if !SovrapposizioneAppuntamenti(due, uno) {
		t.Fail()
	}
}

func TestAppuntamentoReversedTime(t *testing.T) {
	_, correct := NuovoAppuntamento(300, 12, 10)
	if correct {
		t.Fail()
	}
}

func TestAgendaNorm(t *testing.T) {
	agenda := make([]Appuntamento, 0)
	for index := 300; index < 400; index += 20 {
		a, ok := NuovoAppuntamento(index, 11, 13)
		if ok {
			AggiungiAppuntamento(a, &agenda)
			success := AggiungiAppuntamento(a, &agenda)

			if success {
				t.Fail()
			}
		}
	}

	if len(agenda) != 4 {
		t.Fail()
	}

	/*
		if len(AppuntamentiGiornata(300, agenda)) != 1 {
			t.Fail()
		}
	*/
}

func TestUno(t *testing.T) {
	ConfrontaConOracolo(
		t,
		prog,
		"uno.input") // stdin
}

func TestDue(t *testing.T) {
	ConfrontaConOracolo(
		t,
		prog,
		"due.input") // stdin
}

func TestTre(t *testing.T) {
	ConfrontaConOracolo(
		t,
		prog,
		"tre.input") // stdin
}

func TestInvalid(t *testing.T) {
	ConfrontaConOracolo(
		t,
		prog,
		"invalid.input") // stdin
}

func TestVuoto(t *testing.T) {
	ConfrontaConOracolo(
		t,
		prog,
		"vuoto.input") // stdin
}

func TestString(t *testing.T) {
	a := Appuntamento{300, 4, 7}
	s := a.String()
	fmt.Println(s)
	if s != "[il giorno 300 dalle 04 alle 07: ----***-----------------]" {
		t.Fail()
	}
}
