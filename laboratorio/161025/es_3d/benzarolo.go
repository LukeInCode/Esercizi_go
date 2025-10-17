package main
import "fmt"

func main() {
  var litri float64; var scelta int;

  fmt.Print("Inserire il numero di litri di carburante: ");
  _,err_litri := fmt.Scan(&litri);
  
  if err_litri != nil {
    fmt.Println("Inserire un numero di litri valido!")
  }else {
    fmt.Println(
      "Inserire il tipo di carburante:\n",
      "0 - Benzina\n",
      "1 - Diesel\n",
      "2 - Alcol\n",
      "3 - Cherosene",
    );
    _,err_scelta := fmt.Scan(&scelta);
    if err_scelta != nil || scelta < 0 || scelta > 3 {
      fmt.Println("Tipo di carburante non trovato!");
    }else {
      if scelta == 0 {
        const BENZINA = 1.78;
        fmt.Println("Il totale del suo rifornimento è di euro: ", litri * BENZINA);
      } else if scelta == 1 {
        const DIESEL = 1.98;
        fmt.Println("Il totale del suo rifornimento è di euro: ", litri * DIESEL);
      } else if scelta == 2 {
        const ALCOL = 1.2; 
        fmt.Println("Il totale del suo rifornimento è di euro: ", litri * ALCOL);
      } else {
        const CHEROSENE = 1.1;
        fmt.Println("Il totale del suo rifornimento è di euro: ", litri * CHEROSENE);
      }
    }
  }
}
