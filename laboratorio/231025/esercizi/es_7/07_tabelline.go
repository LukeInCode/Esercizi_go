package main
import "fmt"

func main() {
  var n int;

  fmt.Print("Inserire n: "); fmt.Scan(&n);

  for i := 1; i <= 10; i++ {
    fmt.Print(n,"*",i,"= ",n * i,"\t");
  }
  fmt.Println();
}
