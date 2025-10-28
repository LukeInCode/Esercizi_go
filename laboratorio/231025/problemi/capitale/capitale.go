package main
import "fmt"


func main() {
  var initial_amount, tax, goal float64;
  var y int;
  
  fmt.Print("Inserire un capitale iniziale, il tasso e l'obbiettivo: ");
  _,err := fmt.Scan(&initial_amount, &tax, &goal);

  if err != nil {
    fmt.Println("Input errati");
  }else {
    var partial_amount float64 = initial_amount;
    
    for partial_amount < goal {
      y++;
      partial_amount += (partial_amount * tax) / 100;
    }
  }
  fmt.Printf("Ci vogliono %d anni per superare l'obbiettivo\n", y);
}
