package main
import "fmt"

func main() {
  var orario int;

  fmt.Print("Inserire un orario: ");
  _,err := fmt.Scan(&orario);

  if err != nil || orario < 0 || orario > 23 {
    fmt.Println("orario non valido");
  }else {
    if orario <= 6 {
      fmt.Println("notte");
    }else if orario <= 13 {
      fmt.Println("mattino");
    } else if orario <= 18 {
      fmt.Println("pomeriggio");
    }else {
      fmt.Println("sera");
    }
  }
}
