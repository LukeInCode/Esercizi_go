package main
import "fmt"

func main() {
  var n int;

  fmt.Print("Inserire n: ");
  _,err := fmt.Scan(&n);

  if err != nil {
    fmt.Println("Input errati");
  }else {
    var quadrato int;
    
    for quadrato = 1; quadrato*quadrato <= n; quadrato++ {
      fmt.Println("dentro",quadrato)
    }

    fmt.Println("Il massimo quadrato <=",n,"è:",quadrato - 1);
  }
}
