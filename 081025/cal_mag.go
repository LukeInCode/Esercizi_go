/*
  Avveduto Luca
  Calcolo maggiore età
*/

package main
import "fmt"

func main() {
  var anno_nascita, anno_attuale int;
  fmt.Print("Inserisci l'anno attuale: ");
  fmt.Scan(&anno_attuale);
  fmt.Print("Inserisci il tuo anno di nascita: ");
  fmt.Scan(&anno_nascita);
  
  if anno_attuale - anno_nascita >= 18 {
    fmt.Println("Sei maggiorenne!");
  } else {
    fmt.Println("Sei minorenne!");
  }
}
