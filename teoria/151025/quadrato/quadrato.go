package main
import "fmt"

func main() {
  var n int;

  fmt.Print("Inserire il lato del quadrato: "); fmt.Scan(&n);

  for i := 0; i < n; i++ {
    for j := 0; j < n; j++ {
      fmt.Print("*");
    }
    fmt.Println();
  }
}
