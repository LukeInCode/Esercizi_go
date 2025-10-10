package main
import "fmt"
import "math"

func main() {
  var x1,y1,x2,y2 float64;
  
  fmt.Print("Inserire la prima coppia di valori: ");
  _, err := fmt.Scan(&x1, &y1);
  
  if err == nil {
    _,err := fmt.Print("Inserire la seconda coppia di valori: ");
    fmt.Scan(&x2, &y2)
    if err == nil {
      distanza := math.Sqrt( math.Pow((x2-x1),2) + math.Pow((y2-y1),2) );
      fmt.Printf("La distanza tra i punti (%f,%f) e (%f,%f) è %f\n",x1,y1,x2,y2,distanza);
    }else {
      fmt.Println("I valori passati non rispettano i requisiti");
    }
  }else {
    fmt.Println("I valori passati non rispettano i requisiti");
  }
}
