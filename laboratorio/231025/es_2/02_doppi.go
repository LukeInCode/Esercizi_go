package main
import "fmt"

func main() {
  const K = 5;
  var input int;

  for i := 0; i < K; i++ {
    fmt.Print("Inserire un numero: "); fmt.Scan(&input);
    fmt.Print((i + 1), ". ", input, " * 2 = ", input * 2,"\n");
  }
}
