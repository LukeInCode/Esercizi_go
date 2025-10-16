package main;
import "fmt";

func main() {
  var max,min int;
  
  fmt.Print("Inserire due numeri interi: ");
  _,err := fmt.Scan(&max,&min);

  if err != nil {
    fmt.Println("Inserire due numeri interi!");
  }else {
    if min > max {
      min,max = max,min;
    }
    fmt.Println("Il valore massimo è:",max);
  }
}
