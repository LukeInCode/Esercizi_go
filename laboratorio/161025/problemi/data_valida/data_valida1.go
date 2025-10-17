package main
import "fmt"

func main() {
  var mese,giorno int;

  fmt.Print("Inserire rispettivamente giorno e mese: ");
  _,err := fmt.Scan(&giorno,&mese);

  if err != nil {
    fmt.Println("Valori inseriti non validi!");
  }else {
    const MAX_GIORNI = 31; const MAX_MESE = 12;
    if (giorno > 0 && giorno <= MAX_GIORNI) && (mese > 0 && mese <= MAX_MESE) {
      fmt.Println("Data valida");
    }else {
      fmt.Println("Data non valida");
    }
  }
}
