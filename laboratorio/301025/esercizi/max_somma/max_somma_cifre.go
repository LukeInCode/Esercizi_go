package main
import "fmt"

func somma_cifre(n int) (result int) {
  var temp = n;
  result = 0;
  
  for temp > 0 {
    result += temp % 10;
    temp /= 10;
  }
  return;
}


func main() {
  var max, n = -1, 0;
  
  
  for {
    _,err := fmt.Scan(&n);
    
    if n < 0 || n == 999 || err != nil {
      break;
    }
    if max < n {
      max = n;
    }
  }
  fmt.Println("Il numero la cui somma delle cifre è la maggiore è:",max,"e la somma delle sue cire è:",somma_cifre(max));
}
