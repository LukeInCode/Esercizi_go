package main
import "fmt"

func main() {
	var eta int
	var studente bool

	fmt.Print("Inserire la propria età e dichiarare se studenti: ")
	_, err := fmt.Scan(&eta, &studente)

	if err != nil || eta < 0 {
		fmt.Println("Inserire dei valori validi!")
	} else {
		var tariffa float64 = 10
		if  eta < 9 {
			tariffa = 0
		} else if (studente && eta >= 18) || eta < 18 {
      tariffa = 5
    } else if eta < 26 || eta >= 65 {
      tariffa = 7.5
    }
    fmt.Println("Il prezzo totale è di", tariffa)
	}
}
