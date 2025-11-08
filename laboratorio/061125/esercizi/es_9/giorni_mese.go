package main
import "fmt"
import "strconv"

func giorniInMese(mese int) int {
  switch mese {
    case 11,4,6,9:
      return 30;
    case 2 :
      return 28;
    default:
      return 31;
  }
}

func main() {
  var d string;
  for {
    fmt.Scan(&d);
    if len(d) == 10 && d[2] == '-' && d[5] == '-' {
      break;
    }
  }
  mese,_ := strconv.Atoi(d[3:5]);
  fmt.Printf("Il mese %d ha %d giorni\n",mese,giorniInMese(mese));
}
