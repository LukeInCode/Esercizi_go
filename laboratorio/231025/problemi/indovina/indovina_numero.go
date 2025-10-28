package main
import "fmt"
import "math/rand"

func main() {
  const MAX = 10;
  tentativi := MAX/2
  generated := rand.Intn(MAX) + 1;
  var n int;
  
  for i := 1 ;tentativi > 0; i++ {
    fmt.Print("Fai un tentativo: ");
    fmt.Scan(&n);
    if n < 1 && n > generated {
      fmt.Println("fuori dall'intervallo!");
    }else {
      tentativi--;
      if n == generated { 
        fmt.Println("Hai indovinato! (con",i,"tentativi)");
        break;
      }
    }
  }
  if tentativi == 0 {
    fmt.Println("hai perso, il numero era",MAX);
  }
}
