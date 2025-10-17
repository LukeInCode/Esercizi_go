package main
import "fmt"

func main() {
  var guess int; const TARGET = 7;
    
  fmt.Print("Inserire un numero intero da indovinare: ");
  fmt.Scan(&guess);

  if guess >= 1 && guess <= 10 {
    if guess == TARGET {
      fmt.Println("Hai indovinato!");
    }else {
      fmt.Println("Non hai indovinato");
    }
  }else {
    fmt.Println("Valore non valido");
  }
}
