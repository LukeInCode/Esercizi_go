package main
import "fmt"
import "unicode"

func puntiCarta(s string) int {
  var sum int = 0;
  for _,r := range s {
    switch {
      case r == 'A':
        sum += 11;
      case r == '3':
        sum += 10;
      case r == 'K':
        sum += 4;
      case r == 'Q':
        sum += 3;
      case r == 'J':
        sum += 2;
      case !unicode.IsDigit(r):
        sum = -1;
        break;
      default:
        sum += 0;
    }
  }
  return sum;
}

func main() {
  var mano string;
  for {
    fmt.Scan(&mano);
    if len(mano) == 3 {
      break;
    }
  }
  fmt.Printf("mano %s: %d punti\n",mano,puntiCarta(mano));
}
