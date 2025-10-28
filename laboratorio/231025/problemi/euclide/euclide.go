package main
import "fmt"

func mcd(a int, b int) (mcd int) {
  var r int;
  for {
    if b == 0 {
      mcd = a;
      break;
    }
    r = a % b;
    a,b = b,r;
  }
  return;
}


func main() {
  var a,b int;

  fmt.Print("Inserire a e b (a >= b): ");
  _,err := fmt.Scan(&a, &b);
  
  if err != nil || b > a {
    fmt.Println("Input non validi");
  }else {
    fmt.Printf("L'MCD tra %d e %d è: %d\n",a,b,mcd(a,b));
  }
}
