package main
import "fmt"

func main() {
  var num1,den1 int;
  var num2,den2 int;

  fmt.Print("Inserire la prima frazione: ");
  _,err1 := fmt.Scan(&num1, &den1);

  fmt.Print("Inserire la seconda frazione: ");
  _,err2 := fmt.Scan(&num2, &den2);

  if err1 != nil || err2 != nil {
    fmt.Println("Valori passati non validi");
  }else {
    if num1 * den2 == num2 * den1 {
      fmt.Println("Le frazioni sono equivalenti");
    }else {
      fmt.Println("Le frazioni non sono equivalenti");
    }
  }
}
