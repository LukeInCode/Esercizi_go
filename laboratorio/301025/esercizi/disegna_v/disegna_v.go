package main
import "fmt"

func stampa_spazi(spazi int) {
  for i := 0; i < spazi; i++ {
    fmt.Print(" ");
  }
}

func main() {
  var h int;

  fmt.Print("Inserire l'altezza: ");
  _,err := fmt.Scan(&h);

  if err != nil || h < 0 {
    fmt.Println("Paramtri errati");
  }else {
    for i,spazi := 0, h + 1; i < h; i,spazi = i + 1, spazi - 1 {
      stampa_spazi(i);
      fmt.Print("*");
      stampa_spazi(spazi - i);
      if spazi - i > 0 {
        fmt.Println("*");
      }
    }
  }
  fmt.Println();
}
