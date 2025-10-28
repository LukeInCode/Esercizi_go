package main
import "fmt"
import "math"

func bit2ten(val int) (dec int, success bool) {
  var partial,index int = val,0;
  success = true;
  
  for partial > 0 {
      var digit int = partial % 10;
      if digit > 1 {
        success = false;
        break;
      }
      dec += digit * int(math.Pow(2,float64(index)));
      partial /= 10;
      index++;
  }
   return;
}

func main() {
  var n int;

  fmt.Print("Inserire un numero binario: ");
  _,err := fmt.Scan(&n);

  if err != nil {
    fmt.Println("Inserire input validi");
  }else {
    digit,success := bit2ten(n);
    if !success {
      fmt.Println("dato non valido");
    }else {
      fmt.Printf("La conversione di (%d) in decimale è: %d\n",n, digit);
    }
  }
}
