package main
import "fmt"
import "math"

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

func isMersenne(exp int) bool {
  return isPrimo(int(math.Pow(2,float64(exp)) - 1));
}


func main() {
  var n int;

  fmt.Print("Inserire il numero n da verificare: "); fmt.Scan(&n);

  if isMersenne(n) {
    fmt.Println("E' un primo di mersenne");
  }else {
    fmt.Println("Non è un primo di mersenne");
  }
}
