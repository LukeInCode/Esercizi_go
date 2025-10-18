package main
import "fmt"

func main() {
  var n int;

  fmt.Print("Inserire il seme per la sequenza di collatz: "); _,err := fmt.Scan(&n);
  
  if err == nil && n > 0{
    fmt.Println(n);
    for n != 1 {
      if n % 2 == 0 {
        n /= 2;
      }else {
        n = 3 * n + 1;
      }
      fmt.Println(n);
    }
  }
}
