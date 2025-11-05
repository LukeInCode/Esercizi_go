package main
import "fmt"

func main() {
  var n, max int = 0, -1;
  var occ int;
  
  for i := 0; i < 10; i++ {
    fmt.Scan(&n);
    if max < n {
      max = n;
      occ = 0;
    }
    if n == max {
      occ++;
    }
  }
  fmt.Printf("Il numero massimo è %d con %d occorrenze\n", max, occ);
}
