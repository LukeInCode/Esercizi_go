package main
import "fmt"

func main() {
  var n,p,q int;

  fmt.Print("Inserire il pool di possibili multipli, p e q: ");
  fmt.Scan(&n,&p,&q);

  for i := 2; i < n; i++ {
    if p % i == 0 && q % i != 0 {
      fmt.Print(i,"\t");
    }
  }
  fmt.Println();
}
