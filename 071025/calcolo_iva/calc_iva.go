package main
import "fmt"

func main() {
  var scelta int;
  var controllo_ciclo int = 1;
  
  for controllo_ciclo == 1 {
    fmt.Print(
      "Benvenuto! Fai la tua scelta\n",
      "1. Dato imponibile in euro e aliquota trova il prezzo finale\n",
      "2. Dato il prezzo finale e aliquota trova l'imponibile in euro\n",
      "3. Dato l'imponibile e il prezzo finale trova l'aliquota\n",
      "4. Esci\n\n",
    );
    fmt.Scan(&scelta);
    
    switch scelta {
      case 1:
        var imponibile, aliquota float64;
        
        fmt.Print("Inserisci l'imponibile in euro: ");
        fmt.Scan(&imponibile);
        fmt.Print("Inserisci l'aliquota: ");
        fmt.Scan(&aliquota);
        fmt.Println();

        var percentuale float64 = (imponibile * aliquota) / 100;
        fmt.Println("Il prezzo finale è:", imponibile + percentuale);
        
      case 2:
        var prezzo_finale, aliquota float64;
        
        fmt.Print("Inserisci il prezzo finale in euro: ");
        fmt.Scan(&prezzo_finale);
        fmt.Print("Inserisci l'aliquota: ");
        fmt.Scan(&aliquota);
        fmt.Println();

        fmt.Println("L'imponibile è:", (100 * prezzo_finale) / (100 + aliquota));
        
      case 3:
        var prezzo_finale, imponibile float64;

        fmt.Print("Inserisci il prezzo finale in euro: ");
        fmt.Scan(&prezzo_finale);
        fmt.Print("Inserisci l'imponibile in euro: ");
        fmt.Scan(&imponibile);
        fmt.Println(); 
        
        fmt.Println("L'imponibile è:", ((100 * prezzo_finale) / imponibile) - 100);
        
      case 4:
        fmt.Println("Grazie per averci scelto!");
        controllo_ciclo = 0;
        
      default:
        fmt.Println("Comando non trovato");
    }
  };
}
