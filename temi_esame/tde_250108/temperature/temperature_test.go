/*
Temperature
===========

Scrivere un programma in Go (il programma deve chiamarsi 'temperature.go') che elabora temperature.

In particolare il programma legge da linea di comando
- un nome di file (che conterrà una sequenza di temperature intere lette da un sensore, separate da whitespace)
- seguito da alcune temperature (int)
e deve fare le seguenti cose:

- determinare le frequenze delle temperature lette da file e stamparle, su un'unica riga, in ordine crescente di temperatura, separate da spazi, nel formato <temperatura>:<frequenza>

- determinare le temperature (una o più) che si sono verificate con la frequenza massima e stamparle su una riga, in ordine crescente di temperatura, separate da spazi, nel formato <temperatura>:<frequenza>, precedute da "freq max:"

- stampare una riga di separazione "---"

- infine, per ogni temperatura letta da linea di comando, stampare la frequenza corrispondente (stampando 0 se non presente), il formato è sempre <temperatura>:<frequenza>. Si può assumere che le temperature passate su linea di comando siano valide.

Se il nomefile passato a linea di comando non esiste il programma esce senza elaborazione ulteriore stampando "File non esistente".

Se la linea di comando non ha abbastanza parametri (almeno un nomefile e una temperatura), il programma esce senza elaborazione ulteriore stampando "Mancano parametri".

Si può assumere che il file da leggere contenga solo dati validi (interi separati da white spaces), non sono quindi richiesti controlli sulla validità dei dati letti.

Esempi
------

$ ./temperature quattro.input 45 578 23 11 -7
-15:7 -7:7 0:6 8:5 9:1 12:1 22:7
freq max: -15:7 -7:7 22:7
---
45:0 578:0 23:0 11:0 -7:7


$ ./temperature due.input
Mancano parametri


$ ./temperature due.input 3 6
-3:1 -2:2 -1:1 0:2 1:2 2:4 3:4 4:2 5:1 6:1 7:1 8:2 11:1 12:1
freq max: 2:4 3:4
---
3:4 6:1


*/

package main

import "testing"

var prog = "./temperature"
var diffwidth = 150

func TestManca(t *testing.T) {
	ConfrontaConOracolo(t, prog, "", "uno.input")
}

func TestEsempio1(t *testing.T) {
	ConfrontaConOracolo(t, prog, "", "uno.input", "16")
}

func TestEsempio2(t *testing.T) {
	ConfrontaConOracolo(t, prog, "", "due.input", "34", "8")
}

func TestNeg(t *testing.T) {
	ConfrontaConOracolo(t, prog, "", "quattro.input", "-7", "-15")
}

func TestEsempioTempUguali(t *testing.T) {
	ConfrontaConOracolo(t, prog, "", "tre.input", "22")
}

func TestEsempioStesseFreq(t *testing.T) {
	ConfrontaConOracolo(t, prog, "", "quattro.input", "12", "45", "578", "23", "11", "-7")
}
