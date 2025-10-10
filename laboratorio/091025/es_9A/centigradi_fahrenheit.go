package main
import "fmt"

func main() {
	const CENT_FHT = 32;
	const FHT_CENT = 9.0/5.0;
	var centigradi float64;
	fmt.Print("Inserire la temperatura in gradi centigradi: ");
	fmt.Scan(&centigradi);

	var fahrenheit float64 = centigradi * FHT_CENT + CENT_FHT;
	fmt.Println("Temperatura in gradi:",fahrenheit);
}