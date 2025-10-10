/*
	- Il programma prende in input tre numeri interi
	- Controlla che i valori passati siano del tipo corretto
	- Restituisce in console la media tra i tre valori
*/

package main
import "fmt"

func main() {
	var n1,n2,n3 int;
	const TERMINI = 3;
	fmt.Println("Inserire tre valori interi: ");
	_,err := fmt.Scan(&n1,&n2,&n3);

	fmt.Println("Errore: ", err);

	var media int = (n1 + n2 + n3) / TERMINI;
	fmt.Println("La media tra i valori pasati è:", media);
}