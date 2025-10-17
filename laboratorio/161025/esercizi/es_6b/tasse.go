package main
import "fmt"

func main() {
  var reddito float64; var coniugato bool;

  fmt.Print("Inserire il reddito e dichiarare se congiugati (t/f): ");
  _,err := fmt.Scan(&reddito, &coniugato);

  if err != nil || reddito < 0{
    fmt.Println("Valori inseriti non validi!");
  }else {
    const MIN_ALIQUOTA = 0.10; const MAX_ALIQUOTA = 0.24;
    const MAX_REDDITO_SOGLIA = 64000; const MIN_REDDITO_SOGLIA = 32000;
    
    if coniugato {
      if reddito <= MAX_REDDITO_SOGLIA {
        fmt.Printf("Tasse per coniugato con reddito %.2f: %.2f\n", reddito, reddito * MIN_ALIQUOTA);
      }else {
        fmt.Printf("Tasse per coniugato con reddito %.2f: %.2f\n", reddito, reddito * MAX_ALIQUOTA);
      }
    }else {
      if reddito <= MIN_REDDITO_SOGLIA {
        fmt.Printf("Tasse per non coniugato con reddito %.2f: %.2f\n", reddito, reddito * MIN_ALIQUOTA);
      }else {
        fmt.Printf("Tasse per non coniugato con reddito %.2f: %.2f\n", reddito, reddito * MAX_ALIQUOTA);
      }
    }
  }
}
