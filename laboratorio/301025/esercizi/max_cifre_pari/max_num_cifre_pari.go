package main;
import "fmt"

func main() {
  var n string;
  var counter, max int = 0, -1; 

  for n != "000" {
    fmt.Scan(&n);

    var digit byte;
    for i := 0; i < len(n); i++ {
      digit = n[i] - '0';
      if digit % 2 == 0 {
        counter++;
      }
    }
    if max < counter {
      max = counter;
    }
    counter = 0;
  }
  fmt.Printf("Il numero che contiene più cifre pari ha %d cifre\n", max);
}
