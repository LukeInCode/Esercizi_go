/*
Scrivere un programma array.go dotato di:

- una costante DIM = 5 per la dimensione dell'array, dichiarata a livello di package

- una funzione main che inizializza a piacere un array di int di dimensione DIM (ad esempio con numeri 0, 1, 2, ...) e testa le due seguenti funzioni che lavorano **direttamente sull'array stesso**, senza quindi dichiarare e usare altri array. Il programma stampa l'array iniziale, l'array dopo aver invocato reverse e l'array dopo aver invocato switchFirstLast.

- una funzione reverse che inverte l'ordine dei valori in un array di dimensione DIM, mettendo il primo in fondo, il secondo in penultima posizione e così via, nell'array stesso

- una funzione switchFirstLast che scambia il primo con l'ultimo dei valori in un array di dimensione DIM nell'array stesso


Esempio di esecuzione

$ go run array.go

array: [0 1 2 3 4]
reverse: [4 3 2 1 0]
switchFirstLast: [0 3 2 1 4]
*/

package main;
import "fmt";

const DIM = 5;

func main() {
  arr := [DIM]int{2, 3, 5, 7, 11 }
  fmt.Println("array:", arr);
  reverse(&arr);
  fmt.Println("reverse:", arr);
  switchFirstLast(&arr);
  fmt.Println("switchFirstLast:",arr)
}

func reverse(arr *[DIM]int) {
  for i := 0; i < DIM; i++ {
    targetPos := DIM - 1 - i;
    if targetPos == i {
      break;
    }
    (*arr)[i],(*arr)[targetPos] = (*arr)[targetPos], (*arr)[i];
  }
}

func switchFirstLast(arr *[DIM]int) {
  (*arr)[0],(*arr)[DIM - 1] = (*arr)[DIM - 1], (*arr)[0];
}
