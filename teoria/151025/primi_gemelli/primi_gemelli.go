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

  fmt.Print("Inserire il numero n da verificare: "); fmt.Scan(&n);

  for i := 1; i < n; i++ {
    if isPrimo(i) && isPrimo(i + 2) {
      fmt.Printf("(%d,%d)\n",i,i+2);
    }
  } 
}
