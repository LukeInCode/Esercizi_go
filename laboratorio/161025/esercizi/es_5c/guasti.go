package main
import "fmt"

func main() {
  var c1,c2,c3 int;

  fmt.Println(
    "Inserire lo stato dei tre componenti:\n",
    "0. - funzionamento corretto\n",
    "1. - Caricamento acqua\n",
    "2. - Scarico acqua\n",
    "3. - Sistema di riscaldamento",
  );

  _,err := fmt.Scan(&c1,&c2,&c3);

  if err != nil {
    fmt.Println("Valodi passati non validi");
  } else {
    if c1 != 0 || c2 != 0 || c3 != 0 {
      fmt.Println("Guasti rilevati a: ");
      if c1 == 1 || c2 == 1 || c3 == 1 {
        fmt.Println("Caricamento acqua");
      }
      if c1 == 2 || c2 == 2 || c3 == 2 {
        fmt.Println("Scarico acqua");
      }
      if c1 == 3 || c2 == 3 || c3 == 3 {
        fmt.Println("Sistema di riscaldamento");
      }
    }
  }
}
