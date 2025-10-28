package main
import "fmt"
import "math"

func main() {
  var n,max int;
  var max_differenza float64;
  const TARGET = 10;

  for i := 0; i < 5; i++ {
    fmt.Print("Inserire un numero tra 0 e 20: ");
    fmt.Scan(&n);

    differenza := math.Abs(float64(TARGET - n));
    fmt.Println(differenza)
    if max_differenza > differenza || i == 0 {
      max_differenza = differenza;
      max = n;
    }
  }
  fmt.Println("Il valore inserito più vicino a",TARGET,"è:",max);
}
