package main
import "fmt"

func main() {
	var dividendo, divisore int;
	fmt.Print("Inserire un dividendo e un divisore: ");
	fmt.Scan(&dividendo, &divisore);

	divisione_intera := dividendo / divisore;
	resto := dividendo % divisore;

	fmt.Println("Il risultato della divisione intera è:", divisione_intera);
	fmt.Println("Il resto della divisione è:", resto);
}