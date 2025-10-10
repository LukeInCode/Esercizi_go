package main
import "fmt"
import "math"

func main() {
  var n int;
  fmt.Print("Inserire un numero intero: ");
  params,err := fmt.Scan(&n);

  if params != 1 || err != nil {
    fmt.Println("Parametri passati errati");
  }else {
    var i, result int = 1, 0;
    for i <= n {
      result += int(math.Pow(float64(i),2));
      i++;
    }
    fmt.Println("Il risultato della somma è", result);
  }
}
