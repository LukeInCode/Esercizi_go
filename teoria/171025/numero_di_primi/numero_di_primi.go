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

func quantiPrimi(n int) int {
  var sommatore int = 0;
  for i := 2; i < n; i++ {
    if isPrimo(i) {
      sommatore++;
    }
  }
  return sommatore;
}


func main() {
    var n int = 10;
    const ITERAZIONI = 3;

    for i := 1; i <= ITERAZIONI; i,n = i + 1, n * 10 {
      n_primi := float64(quantiPrimi(n));
      fmt.Print("Per n = ",n,", i numeri primi sono: ",n_primi,"\n");
      prova := n_primi / (float64(n) / math.Log(float64(n)));
      fmt.Println("π(x) / ( x/ln(x) ) =",prova);
    }
}
