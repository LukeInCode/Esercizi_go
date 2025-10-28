package main
import "fmt"

func main() {
  var n, previous_n float64;

  for {
    fmt.Print("Inserire un numero: ");
    _,err := fmt.Scan(&n);

    if err != nil || n == 0 {
      fmt.Println("Uscita dal programma...");
      break;
    }else {
      if previous_n != 0 {
        fmt.Printf("Differenza tra %.2f e %.2f: %.2f\n",n, previous_n, n - previous_n);
      }
      previous_n = n;
    }
  }
}
