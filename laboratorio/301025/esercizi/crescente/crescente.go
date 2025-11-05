package main
import "fmt"

func main() {
  var parola string;
  

  fmt.Print("Inserire una parola: ");
  _,err := fmt.Scan(&parola);

  if err != nil {
    fmt.Println("Input errati");
  }else {
    var precedente rune;
    for _,r := range parola {
      if r < 97 || r > 122 {
        break;
      }else {
        if r < precedente && precedente != 0 {
          fmt.Print("-");
        }
        fmt.Print(string(r));
      }
      precedente = r;
    }
  }
  fmt.Println();
}
