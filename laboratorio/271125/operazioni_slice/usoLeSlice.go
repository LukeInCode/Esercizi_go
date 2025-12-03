/*
Scrivere un programma usoLeSlice.go che

1. legge da linea di comando una lista di almeno 3 parole, le salva in una slice e stampa la slice;
2. legge da standard input un'altra lista di parole (anche almeno 3), le salva in una slice e stampa la slice.

Quindi fa le seguenti operazioni e man mano stampa il risultato ottenuto:

 3. crea un nuova lista in cui salva gli elementi della prima lista seguiti dagli elementi della seconda lista
 4. mette in ordine alfabetico (lessicografico) la slice (usare una funzione di libreria) e stampa la slice ottenuta
 5. cancella l'ultima parola dalla slice
 6. rimuove dalla slice gli elementi di indice dal 2 (incluso) al 4 (escluso)
 7. crea una nuova slice con "aa", "bb" e "cc"
 8. inserisce la nuova slice in posizione 1 della vecchia slice ("aa" avrà indice 1, ecc. e gli elementi della vecchia slice, tranne quello di indice 0, andranno spostati a destra della slice inserita)
 9. legge una nuova parola e la inserisce in posizione 2 della slice
10. legge una nuova parola e la inserisce in fondo alla slice
11. estende la slice con una nuova slice di lunghezza 2 (di stringhe vuote)
12. inserisce alla posizione 3 una nuova slice di lunghezza 3 (di stringhe vuote)
13. copia in una nuova slice la slice ottenuta fin qui
14. dalla copia rimuove l'ultimo elemento e stampa sia la vecchia slice che la copia.

Completare il programma inserendo al posto dei "..." le istruzioni occorrenti.

*/

package main

import (
	"fmt"
	"os"
	"bufio"
	"slices"
	"strings"
)

func main() {
	//1. legge da linea di comando una lista di almeno 3 parole e le salva in una slice
	if len(os.Args) < 4 {
		fmt.Println("almeno 3 parole su CL")
		return
	}
	var wordsFromCL []string 
	wordsFromCL = append(wordsFromCL, os.Args[1:]...);
	fmt.Println(" 1.", wordsFromCL)

	//2. legge da standard input un'altra lista di parole (anche almeno 3) e le salva in una slice
	fmt.Println("scrivi almeno 3 parole (seguite da ctrl D)")
	scanner := bufio.NewScanner(os.Stdin);
	var wordsFromStdin, temp []string;
	for scanner.Scan() {
		elements := strings.Split(scanner.Text(), " ");
		temp = append(temp, elements...);
	}
	if len(temp) < 3 {
		fmt.Println("almeno 3 parole su STDIN")
		return
	}
	wordsFromStdin = append(wordsFromStdin, temp...);
	fmt.Println(" 2.", wordsFromStdin)

	//3. crea un nuova lista mySlice in cui salva gli elementi della prima lista seguiti dagli elementi della seconda lista
	var mySlice []string;
	mySlice = append(mySlice, append(wordsFromCL, wordsFromStdin...)...);
	fmt.Println(" 3.", mySlice)

	// 4. mette in ordine alfabetico (lessicografico) la slice (usare una funzione di libreria) e stampa la slice ottenuta
	slices.Sort(mySlice)
	fmt.Println(" 4.", mySlice)

	// 5. cancella l'ultima parola dalla slice
	mySlice = mySlice[:len(mySlice) - 1]
	fmt.Println(" 5.", mySlice)

	// 6. rimuove dalla slice gli elementi di indice dal 2 (incluso) al 4 (escluso)
	mySlice = slices.Delete(mySlice, 2, 4);
	fmt.Println(" 6.", mySlice)

	// 7. crea una nuova slice con "aa", "bb" e "cc" e la stampa
	myNewSlice := []string{"aa","bb","cc"}
	fmt.Println(" 7.", myNewSlice)

	// 8. inserisce la nuova slice in posizione 1 della vecchia slice
	mySlice = append(mySlice[:1], append(myNewSlice, mySlice[1:]...)...)
	fmt.Println(" 8.", mySlice)

	// 9. legge una nuova parola e la inserisce in posizione 2 della slice
	fmt.Print("scrivi una parola: ")
	var word []string = make([]string, 1);
	fmt.Scan(&word[0]);
	mySlice = append(mySlice[:2], append(word, mySlice[2:]...)...);
	fmt.Println(" 9.", mySlice)

	//10. legge una nuova parola e la inserisce in fondo alla slice
	fmt.Print("scrivi una parola: ")
	var second_word string;
	fmt.Scan(&second_word);
	mySlice = append(mySlice, second_word);
	fmt.Println("10.", mySlice)

	// 11. estende la slice con una nuova slice di lunghezza 2 (di stringhe vuote)
	var emptySlice []string = []string{"", ""};
	mySlice = append(mySlice, emptySlice...);
	fmt.Println("11.", mySlice)

	// 12. inserisce alla posizione 3 una nuova slice di lunghezza 3 (di stringhe vuote)
	emptySlice = append(emptySlice, "");
	mySlice = append(mySlice[:3], append(emptySlice, mySlice[3:]...)...);
	fmt.Println("12.", mySlice)

	// 13. copia in una nuova slice la slice ottenuta fin qui
	copyOfMySlice := make([]string, len(mySlice));
	copy(copyOfMySlice, mySlice);
	fmt.Println("13.", copyOfMySlice);

	// 14. dalla copia rimuove l'ultimo elemento e stampa sia la vecchia slice che la copia.
	copyOfMySlice = copyOfMySlice[:len(copyOfMySlice) - 1];
	fmt.Println("14.", mySlice, "\n   ", copyOfMySlice)
}

