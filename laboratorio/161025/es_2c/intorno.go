package main
import "fmt"
import "math"

func main() {
  var x,y float64;

  _,err := fmt.Scan(&x, &y);

  if err != nil {
    fmt.Println("Valori passati non validi");
  }else {
    const EPSILON = 1e-5;
    var distanza float64 = math.Sqrt(math.Pow((x-0),2) + math.Pow((y-0),2));
    
    if  distanza < EPSILON{
      fmt.Println("molto vicino all'origine");
    }
  }
}
