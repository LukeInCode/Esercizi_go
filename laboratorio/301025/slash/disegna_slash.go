package main
import "fmt"

func main() {
  var h int;

  fmt.Print("Inserire l'altezza dello slash: ");
  _,err := fmt.Scan(&h);

  if err != nil || h < 0 {
    fmt.Println("Parametri errati");
  }else {
    for i := 0; i < h; i++ {
      for j := 0; j < i; j++ {
        fmt.Print(" ");
      }
      fmt.Println("*");
    }
  }
}
