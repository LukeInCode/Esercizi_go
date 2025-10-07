package main
import "fmt"

func main() {
  var decimal_float float64;
  fmt.Print("Ciao Biondo! Inserisci il tuo numero decimale: ");
  fmt.Scan(&decimal_float);
  var int_input int = int(decimal_float);
  fmt.Println(decimal_float - float64(int_input));
}
