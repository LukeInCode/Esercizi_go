package main
import "fmt"
import "io"

func haDoppie(s string) bool {
  var haDoppie = false;
  var previousCh rune;
  
  for _,r := range s {
      if r == previousCh {
        haDoppie = true;
        break;
      }
      previousCh = r;
  }
  return haDoppie;
}

func main() {
  var parola string;
  var counter int = 0;
  for {
    _,err := fmt.Scan(&parola);
    if err == io.EOF {
      break;
    }
    if haDoppie(parola) {
      counter++;
    }
  }
  fmt.Printf("le parole con doppie sono: %d\n", counter);
}
