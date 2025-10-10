package main
import "fmt"

func main() {
  var n int;
  fmt.Print("Inserire un valore intero: ");
  params, err := fmt.Scan(&n);
  
  if params != 1 || err != nil {
    fmt.Println("Valori passati sbagliati!");
  } else {
    var temp, counter int = n, 0;
    for temp > 0 {
      temp /= 10;
      counter++;
    }
    fmt.Printf("Il numero %d ha %d cifre\n", n, counter);
  }
}
