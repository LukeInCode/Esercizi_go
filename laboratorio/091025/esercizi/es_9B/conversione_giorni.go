package main
import "fmt"

func main() {
	const GIORNI_ANNO = 365;
	const GIORNI_SETT = 7;
	
	var giorni_input int;
	fmt.Print("Inserire il numero di giorni: ");
	fmt.Scan(&giorni_input);
	
	var anni, settimane, giorni int;

	anni = giorni_input / GIORNI_ANNO;
	temp := giorni_input - (GIORNI_ANNO * anni);
	
	settimane = temp / GIORNI_SETT;
	giorni = temp - (settimane * GIORNI_SETT);

	fmt.Println(anni,settimane,giorni);
}