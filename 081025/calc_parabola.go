/*
  Avveduto Luca
  Calcolo soluzioni di un'equazione di secondo grado
*/

package main
import (
  "fmt"
  "math"
)

func main() {
  var a,b,c float64;

  fmt.Print("Inserisci il primo coefficiente: ");
  fmt.Scan(&a);
  fmt.Print("\nInserisci il secondo coefficiente: ");
  fmt.Scan(&b);
  fmt.Print("\nInserisci il terzo coefficiente: ");
  fmt.Scan(&c);

  var delta float64 = math.Pow(b,2) - 4 * a * c;
  var x1, x2 float64 = (b * -1 + math.Sqrt(delta)) / (2 * a), (b * -1 - math.Sqrt(delta)) / (2 * a);
  fmt.Printf("La prima soluzione reale è %.2f, e la seconda è %.2f\n",x1,x2);
}
