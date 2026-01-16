/*
Realizzare un programma Go (nome file 'nl.go') che implementi un semplice 'nl', comando Unix che "numera" le linee di un testo in input (stdin).

Il programma legge le righe da `stdin` (fino a incontrare EOF, ctrl D da tastiera) e deve stampare su `stdout` le righe lette, ciascuna preceduta da:

- un numero progressivo (conteggio righe partendo da 1) seguito da ","

- la lunghezza (n. di byte) della riga corrente seguito da ":"

- la riga corrente

Al termine della lista numerata il programma deve stampare la lunghezza massima di riga incontrata durante l'elaborazione, scrivendo:

lunghezza max <numero>

## Esempio esecuzione

(presuppone il 'nl.go' già compilato, 'uno.input' è presente in questa directory)

$ ./nl < uno.input
1,48:Lorem Ipsum is simply dummy text of the printing
2,76:and typesetting industry. Lorem Ipsum has been the industry's standard dummy
3,55:text ever since the 1500s, when an unknown printer took
4,48:a galley of type and scrambled it to make a type
5,102:specimen book. It has survived not only five centuries, but also the leap into electronic typesetting,
6,81:remaining essentially unchanged. It was popularised in the 1960s with the release
7,60:of Letraset sheets containing Lorem Ipsum passages, and more
8,97:recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.
9,0:
10,48:Lorem Ipsum is simply dummy text of the printing
11,76:and typesetting industry. Lorem Ipsum has been the industry's standard dummy
12,55:text ever since the 1500s, when an unknown printer took
13,48:a galley of type and scrambled it to make a type
14,102:specimen book. It has survived not only five centuries, but also the leap into electronic typesetting,
15,81:remaining essentially unchanged. It was popularised in the 1960s with the release
16,60:of Letraset sheets containing Lorem Ipsum passages, and more
17,97:recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.
18,48:Lorem Ipsum is simply dummy text of the printing
19,76:and typesetting industry. Lorem Ipsum has been the industry's standard dummy
20,55:text ever since the 1500s, when an unknown printer took
21,48:a galley of type and scrambled it to make a type
22,102:specimen book. It has survived not only five centuries, but also the leap into electronic typesetting,
23,81:remaining essentially unchanged. It was popularised in the 1960s with the release
24,60:of Letraset sheets containing Lorem Ipsum passages, and more
25,97:recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.
lunghezza max 102

*/

package main

import (
	"testing"
)

var prog = "./nl"
var diffwidth = 150

func TestComeDaEsempio(t *testing.T) {
	ConfrontaConOracolo(
		t,
		prog,
		"uno.input")
}

func TestSelf(t *testing.T) {
	ConfrontaConOracolo(
		t,
		prog,
		"nl_test.go")
}
