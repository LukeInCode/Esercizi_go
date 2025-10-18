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

func sommaCifre(val int) int {
  var somma int = 0;
  for val > 0 {
    somma += val % 10;
    val /= 10;
  }
  return somma;
}


func main() {
  var n int;

  fmt.Print("Inserire il numero n limite: "); fmt.Scan(&n);

  for i := 1; i <= n; i++ {
    if isPrimo(i) && sommaCifre(i) == 13 {
      fmt.Println(i);
    }
  } 
}
