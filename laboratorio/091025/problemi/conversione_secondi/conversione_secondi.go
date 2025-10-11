package main
import "fmt"

func main() {
  var secondi_input int;
  fmt.Print("Inserire il numero di secondi da convertire: ");
  _,err := fmt.Scan(&secondi_input);

  if err != nil {
    fmt.Println("Inserire un numero di secondi intero");
  }else {
    const CONVERSIONE_SECONDI_ORE = 3600;
    const CONVERSIONE_MINUTI_ORE = 60;
    const CONVERSIONE_GIORNI = 24;

    var temp,minuti,ore,giorni int = secondi_input,0,0,0;

    
    giorni = secondi_input / (CONVERSIONE_SECONDI_ORE * CONVERSIONE_GIORNI);
    temp = secondi_input % (CONVERSIONE_SECONDI_ORE * CONVERSIONE_GIORNI);
    ore,temp = temp / CONVERSIONE_SECONDI_ORE, temp % CONVERSIONE_SECONDI_ORE;
    minuti,temp = temp / CONVERSIONE_MINUTI_ORE, temp % CONVERSIONE_MINUTI_ORE;

    fmt.Printf("gg : %d, hh: %d, min : %d, sec : %d\n", giorni, ore, minuti, temp);
  }
}
