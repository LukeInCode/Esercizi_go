package main;
import "fmt";

func main() {
  var orario int;

  fmt.Print("Inserire un orario da verificare: ");
  _,err := fmt.Scan(&orario);

  if err != nil {
    fmt.Println("Inserire un orario intero!");
  }else {
    if orario >= 6 && orario < 13 {
      fmt.Println("Mattino");
    }
  }
}
