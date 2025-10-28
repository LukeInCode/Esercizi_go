package main
import "fmt"

func main() {
  var n, previous_n,sum float64;

  for {
    fmt.Print("\nInserire un numero: ");
    _,err := fmt.Scan(&n);

    if err != nil || n == -1 {
      fmt.Println("Uscita dal programma...");
      break;
    }else {
      sum += n;
      if previous_n != 0 {
        if n >= previous_n {
          fmt.Print("+");
        }else {
          fmt.Print("-");
        }
      }
      previous_n = n;
    }
  }
  fmt.Println("somma:",sum);
}
