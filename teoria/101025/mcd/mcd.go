package main
import "fmt"
import "math"

func main() {
  var x,y int;
  fmt.Print("Inserire un numero intero: ");
  params,err := fmt.Scan(&x, &y);

  if params != 2 || err != nil {
    fmt.Println("Parametri passati errati");
  }else {
    var min, i, divisore = int(math.Min(float64(x),float64(y))), 2, 0;
    
    for i <= min {
      if x % i == 0 && y % i == 0 {
        divisore = i;
      }
      i++;
    }
    fmt.Printf("L'MCD tra %d e %d è %d\n",x,y,divisore);
  }
}
