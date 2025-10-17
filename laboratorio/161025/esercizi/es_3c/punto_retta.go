package main
import "fmt"

func main() {
  var x,y,m,q float64;

  fmt.Print("Inserire le coordinate del punto da verificare, la pendenza e l'intercetta: ");
  _,err := fmt.Scan(&x,&y,&m,&q);

  if err != nil {
    fmt.Println("Valori passati non validi");
  }else {
    ordinata_attesa := m * x + q;
    const TOLLERANZA = 1e-6;

    if ordinata_attesa >= (y - TOLLERANZA) && ordinata_attesa <= (y + TOLLERANZA) {
      fmt.Println("Il punto appartiene alla retta");
    }else if ordinata_attesa > y {
      fmt.Println("Il punto si trova sotto la retta");
    }else {
      fmt.Println("Il punto si trova sopra la retta");
    }
  }
}
