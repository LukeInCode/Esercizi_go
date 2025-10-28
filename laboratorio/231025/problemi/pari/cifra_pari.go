package main
import "fmt"

func findEvendigit(val int) (pos int, value int) {
  var partial int = val;
  pos = 0;
  
  for partial > 0 {
      pos++;
      var digit int = partial % 10;
      if digit % 2 == 0 {
        value = digit;
        return;
      }
      partial /= 10;
  }
   pos = -1;
   return;
}

func main() {
  var n int;

  fmt.Print("Inserire un numero intero: ");
  _,err := fmt.Scan(&n);

  if err != nil {
    fmt.Println("Inserire input validi");
  }else {
    pos, digit := findEvendigit(n);
    fmt.Printf("La posizione della prima cifra pari (%d) è: %d\n", digit, pos);
  }
}
