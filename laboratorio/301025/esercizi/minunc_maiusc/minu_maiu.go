package main
import "fmt"
import "unicode"

func main() {
  var s string;
  fmt.Print("Inserire una stringa: ");
  _,err := fmt.Scan(&s);

  if err != nil {
    fmt.Println("valori non validi");
  }else {
    var check, previous int = 0, 0;
    
    for i,r := range s { 
     if unicode.IsUpper(r) {
      check = 1;
     }else if unicode.IsLower(r) {
      check = 0;
     }
     if check != previous && i > 0{
      fmt.Println(previous);
      check = -1;
      break;
     }
     previous = check;
    }
    switch check {
      case 0:
        fmt.Println("Tutto minuscolo");

      case 1:
        fmt.Println("Tutto maiuscolo");

      case -1:
        fmt.Println("Sia maiuscolo che minuscolo");

      default:
        fmt.Println("Can't happen");
    }
  }
}
