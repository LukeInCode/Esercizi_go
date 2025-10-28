package main
import "fmt"

func sumdigits(val int) (sum int) {
  var partial int = val;
  
  for partial > 0 {
      var digit int = partial % 10;
      sum += digit;
      partial /= 10;
  }
   return;
}

func main() {
  var n int;

  fmt.Print("Inserire un numero intero: ");
  _,err := fmt.Scan(&n);

  if err != nil {
    fmt.Println("Inserire input validi");
  }else {
    digit := sumdigits(n);
    fmt.Printf("La somma delle cifre di (%d) è: %d\n",n, digit);
  }
}
