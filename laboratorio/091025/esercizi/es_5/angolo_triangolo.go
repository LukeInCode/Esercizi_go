package main
import "fmt"

func main() {
	var angolo_1, angolo_2 float64;
	fmt.Print("Inserire i due angoli del triangolo: ");
	fmt.Scan(&angolo_1, &angolo_2);
	const SOMMA_ANGOLI = 180;

	angolo_finale := 180 - (angolo_1 + angolo_2);
	fmt.Println("Il terzo angolo ha ampiezza:", angolo_finale);
}