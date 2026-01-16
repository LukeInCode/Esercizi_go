/*

Luhn
====

Scrivere un programma in Go (il file deve chiamarsi 'luhn.go') che verifica la validità di numeri di carte di credito (16 cifre che rispettano la regola di Luhn).

Il programma deve essere dotato di una funzione

	isLuhn(s string) bool

che, data una stringa costituita da 16 cifre (non sono richiesti controlli nella funzione), restituisce true se la sequenza di 16 cifre è un numero di carta di credito valido, false altrimenti.

Il programma, data su linea di comando una lista di lunghezza arbitraria di stringhe, le stampa, una per riga, ciascuna seguita da "valido" / "non valido" (virgolette escluse), a seconda che siano validi numeri di carta di credito o no.

Il programma deve stampare "non valido" anche per dati non numerici o valori con un numero di caratteri diverso da 16.

Se non sono presenti dati sulla linea di comando, il programma stampa "nessun argomento" (virgolette escluse) e termina.

Infine il programma stampa il numero di carte di credito valide e il numero di quelle non valide (vedi esempio sotto).

L'algoritmo di Luhn, che consente di controllare la validità di vari numeri, tra cui quelli delle carte di credito, si basa su tre passi:
1. moltiplicare una cifra sì e una no per due; più precisamente, per le carte di credito moltiplicare per due le cifre in posizioni pari (in posizione 0, 2, ecc.), dove la prima posizione è la posizione 0;
2. per le cifre per cui la moltiplicazione ha dato un risultato a due cifre, sottrarre 9 dal prodotto;
3. sostituire ciascuna cifra originale presa in considerazione al passo 1. con quella ottenuta con i passi 1. e 2.;
4. sommare tutte le cifre, sia quelle in posizione pari ottenute con i passi 1. e 2., sia quelle originali che si trovano in posizione dispari;
5. se la somma complessiva è divisibile per 10 (la divisione non ha resto), il numero è valido.

Ad esempio, dato il numero:
4539148803436467
si dovranno raddoppiare
4_3_1_8_0_4_6_6_

ottenendo:
8_6_2_7_0_8_3_3_

Sommando ora tutte le cifre
8 5 6 9 2 4 7 8 0 3 8 3 3 4 3 7
8+5+6+9+2+4+7+8+0+3+8+3+3+4+3+7 = 80

si ottiene un multiplo di 10, quindi questo numero è valido.

Esempi
======

$ ./luhn 4539148803436467 8273123273520569 4024007190270651 4929876177122565 bastiancontrario 123unduetre123un 492987
4539148803436467 valido
8273123273520569 non valido
4024007190270651 valido
4929876177122565 non valido
bastiancontrario non valido
123unduetre123un non valido
492987 non valido
Validi: 2
Non validi: 5

$ ./luhn
nessun argomento

*/

package main

import (
	"testing"
)

var prog = "./luhn"
var diffwidth = 100

func TestMainMultiplo(t *testing.T) {
	ConfrontaConOracolo(
		t,
		prog,
		"",
		//"4539148803436467 valido\n8273123273520569 non valido\n4024007190270651 valido\n4929876177122565 non valido\nbastiancontrario non valido\n123unduetre123un non valido\n492987 non valido\nValidi: 2\nNon validi: 5\n",
		"4539148803436467",
		"8273123273520569",
		"4024007190270651",
		"4929876177122565",
		"bastiancontrario",
		"123unduetre123un",
		"3N2D2N3X3b2l1v4b",
		"492987",
		"49298797")
}

func TestMainNoInput(t *testing.T) {
	ConfrontaConOracolo(
		t,
		prog,
		"")
}

func TestValido(t *testing.T) { //analizza sequenze di 16 cifre
	if !isLuhn("4539148803436467") || !isLuhn("4024007190270651") {
		t.Fail()
	}
}

func TestNonValido(t *testing.T) { //analizza sequenze di 16 cifre
	if isLuhn("8273123273520569") || isLuhn("4929876177122565") {
		t.Fail()
	}
}
