package main
import "fmt"

func main() {
  var n int;

  fmt.Print("Inserire un numero: ");
  _,err := fmt.Scan(&n);

  if err != nil {
    fmt.Println("Inserire un numero intero!");
  }else {
    divisibile_10 := n % 10 == 0;
    if divisibile_10 && n > 0 {
      fmt.Println("positivo divisibile per 10");
    }else {
      if divisibile_10 {
        fmt.Println("divisibile per 10");
      }else if n > 0 {
        fmt.Println("positivo");
      }
    }
  }
}
