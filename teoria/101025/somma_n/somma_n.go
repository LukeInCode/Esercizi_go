package main
import "fmt"

func main() {
  var n int;
  fmt.Print("Inserire un numero intero: ");
  params,err := fmt.Scan(&n);

  if params != 1 || err != nil {
    fmt.Println("Parametri passati errati");
  }else {
    var i, result int = 1, 0;
    for i <= n {
      result += i;
      i++;
    }
    fmt.Println("Il risultato della somma è", result);
  }
}
