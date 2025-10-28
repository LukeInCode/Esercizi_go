package main
import "fmt"

func isPerfectSquare(n int) bool {
  for i := 2; i < n; i++ {
    if i * i == n {
      return true;
    }
  }
  return false;
}

func main() {
  var n int;
  
  fmt.Print("Inserire n: "); fmt.Scan(&n);

  for i := 1; i <= n; i++ {
    if isPerfectSquare(i) {
      fmt.Println("Quadrato perfetto:",i);
    }
  }
}
