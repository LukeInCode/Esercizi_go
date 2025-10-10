package main
import "fmt"

func main() {
  var n int;
  fmt.Print("Inserire il numero da verificare: ");
  params,err := fmt.Scan(&n);

  if params != 1 || err != nil {
    fmt.Println("Parametri passati errati");
  }else {
    i,trovato_divisore := 2, 0;
    for i < n && trovato_divisore == 0 {
      if n % i == 0 {
        trovato_divisore = i;
      }
      i++;
    }
    if trovato_divisore != 0 {
      fmt.Println("Il numero non è primo, il suo primo divisore è:", trovato_divisore);
    }else {
      fmt.Println("Il numero è primo");
    }
  }
}
