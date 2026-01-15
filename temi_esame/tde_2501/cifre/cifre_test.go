/*
Cifre
------

# Scrivere un programma cifre.go che elabora le cifre contenute in un testo passato da linea di comando.

Il programma deve essere dotato di:

  - una funzione
    contaMaxTot(s string) ([10]int, int, int)

che restituisce, in quest'ordine:
  - un array di lunghezza 10 che contiene il numero di '0', '1', '2', ... e '9' (in quest'ordine,) che si trovano nella stringa s,
  - la massima cifra presente nella stringa s (se non sono presenti cifre nella stringa, il valore restituito deve essere -1)
  - la somma delle cifre contenute nella stringa s.

Non ci sono vincoli sui tipi di caratteri in s (possono cioè essere anche non ASCII).

- una funzione main() che legge una frase (racchiusa tra "") da linea di comando che
contiene caratteri di qualsiasi tipo (cifre, lettere, simboli, punteggiatura, lettere accentate, ecc.)
e stampa quante cifre di ogni tipo ci sono, qual è la cifra più grande e la somma delle cifre incontrate.
Se manca l'argomento da linea di comando, deve stampare "manca argomento" e terminare.

Esempio
-------

Input: "1, 2, 3; non c'è fiaba senza re; 1, 2, 3; venite giù da me; 4, 5, 6; siete dei babbei; 7,8,9; io sono già altrove."
Output: [0 2 2 2 1 1 1 1 1 1] 9 51
*/
package main

import (
	"fmt"

	//"strings"
	//"log"
	"testing"
)

var prog = "./cifre"
var diffwidth = 120

func Test1(t *testing.T) {
	ConfrontaConOracolo(
		t,
		prog,
		"", // file per stdin
		"1, 2, 3; non c'è fiaba senza re; 1, 2, 3; venite giù da me; 4, 5, 6; siete dei babbei; 7,8,9; io sono già altrove.") // arg1
}

func Test2(t *testing.T) {
	ConfrontaConOracolo(
		t,
		prog,
		"",                         // file per stdin
		"[02 2 2 2 11 1 1 13 1 1]") // arg1
}

func Test3(t *testing.T) {
	ConfrontaConOracolo(
		t,
		prog,
		"",        // file per stdin
		"abcdefg") // arg1
}

func Test4(t *testing.T) {
	ConfrontaConOracolo(
		t,
		prog,
		"",                             // file per stdin
		"aaaaaaa1ddd3ddddd64gggg8hhhh") // arg1
}

func Test5(t *testing.T) {
	ConfrontaConOracolo(
		t,
		prog,
		"",           // file per stdin
		"0123456789") // arg1
}

func TestUnicode(t *testing.T) {
	ConfrontaConOracolo(
		t,
		prog,
		"",                                // file per stdin
		"aaaa⚡aaa1ddd3dddÞdd6gg¢gg8h🔋hhh") // arg1
}

func TestMain(t *testing.T) {
	ConfrontaConOracolo(
		t,
		prog,
		"", // file per stdin
		"abcde45fghi8askdjhaskdjash6sakdjhasdkjsah28") // arg1
}

func TestMainNoArg(t *testing.T) {
	ConfrontaConOracolo(
		t,
		prog,
		"", // file per stdin
	) // arg1
}

// verifica che sia proprio un array di int
func TestCheckType(t *testing.T) {
	var array [10]int
	var max, tot int
	array, max, tot = contaMaxTot("23, non c’è fiaba senza re; 33, venite giù da me; 156 siete dei babbei; 789 io sono già altrove.")
	fmt.Println(array, max, tot)
}
