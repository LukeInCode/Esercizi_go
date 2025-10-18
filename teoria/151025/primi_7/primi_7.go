package main
import "fmt"

func isPrimo(val int) bool {
  var primo bool = true;
  for i := 2; i < val; i++ {
    if val % i == 0 {
      primo = false;
      break;
    }
  }
  return primo;
}


func main() {
  var n int;

  fmt.Print("Inserire il numero n limite: "); fmt.Scan(&n);

  for i := 7; i <= n; i++ {
    if isPrimo(i) && i % 10 == 7 {
      fmt.Println(i);
    }
  } 
}
