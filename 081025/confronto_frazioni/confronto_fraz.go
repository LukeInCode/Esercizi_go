/*
  Avveduto Luca
  Calcolo soluzioni di un'equazione di secondo grado
*/

package main
import "fmt"

func main() {
  var numeratore_1, denominatore_1 int;
  var numeratore_2, denominatore_2 int;

  fmt.Print("Inserisci il primo numeratore: ");
  fmt.Scan(&numeratore_1);
  fmt.Print("Inserisci il primo denominatore: ");
  fmt.Scan(&denominatore_1);
  fmt.Print("Inserisci il secondo numeratore: ");
  fmt.Scan(&numeratore_2);
  fmt.Print("Inserisci il secondo denominatore: ");
  fmt.Scan(&denominatore_2);

  var ris_1, ris_2 float64 = float64(numeratore_1) / float64(denominatore_1), float64(numeratore_2) / float64(denominatore_2)

  if ris_1 > ris_2 {
    fmt.Println("La prima frazione è la maggiore");
  }else if ris_1 < ris_2 {
    fmt.Println("La seconda frazione è la maggiore");
  }else {
    fmt.Println("Le due frazioni sono equivalenti");
  }
}
