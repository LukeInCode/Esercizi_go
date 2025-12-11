package main

import (
	"fmt"
	"testing"
)

var prog = "./segmento"
var diffwidth = 120

/*
func TestTODO(t *testing.T) {
	fmt.Println("RIMPOLPARE TEST!!!", prog)
	t.Fail()
}
*/

func TestMainMinimale(t *testing.T) {
	/* da STDIN
	- il carattere per gli estremi
	- il carattere interno
	- true o false per orizzontale o vertivale
	- la lunghezza del segmento
	*/
	LanciaGenerica(t,
		prog,
		"+ - true 4",
		"+--+\n+-----+\n+\n-\n-\n-\n-\n-\n+\n")
}

func TestEsistenzaMetodoString(t *testing.T) {
	var s Segmento
	//fmt.Println(r)
	fmt.Println(s.String()) // non compila se manca
}

func TestEsistenzaFunzione(t *testing.T) {
	var s Segmento
	fmt.Println("prima:", s)
	allunga(&s, 3)
	fmt.Println("dopo allung:", s)
	if s.lunghezza != 3 {
		t.Fail()
	}
	allunga(&s, -2)
	fmt.Println("dopo accorcia:", s)
	if s.lunghezza != 1 {
		t.Fail()
	}
}
