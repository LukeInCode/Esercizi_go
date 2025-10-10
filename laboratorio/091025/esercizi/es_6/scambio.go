package main
import "fmt"

func main() {
	var a,b int;
	fmt.Print("Inserire i valori delle due variabili da scambiare: ");
	_, err := fmt.Scan(&a, &b);
	fmt.Println(err);
	
	fmt.Println("Prima dello scambio: ",a,b);
	a,b = b,a;
	fmt.Println("Dopo lo scambio: ",a,b);
}