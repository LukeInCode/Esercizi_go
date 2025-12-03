package main
import "fmt"
import "slices"
import "os"

func trovaGiorno(primoGennaio string, giorno int, settimana [7]string) string {
  var lenSettimana int = len(settimana);
  var posizioneIniziale int = slices.Index(settimana[:], primoGennaio);
  var risultato = (giorno % lenSettimana + posizioneIniziale - 1) % lenSettimana;
  return settimana[risultato];
}


func main() {
  var giorni [7]string = [7]string{"lun", "mart", "merc", "giov", "ven", "sab", "dom"};
  var primoGennaio string;
  var anno int;

  fmt.Print("Inserire un anno ed il giorno in cui cade il primo Gennario: ");
  fmt.Scanf("%s %d", &primoGennaio, &anno);
  

  if !slices.Contains(giorni[:], primoGennaio) {
    fmt.Printf("Il giorno %s non è valido\n", primoGennaio);
    os.Exit(1);
  }

  var n int;
  for {
    fmt.Print("Inserire un giorno da verificare (-1 per uscire): ");
    fmt.Scanf("%d", &n);
    if n == -1 {
      break;
    }
    fmt.Printf("\nIl giorno %d, dell'anno %d cade di: %s\n", n, anno, trovaGiorno(primoGennaio, n, giorni));
  }
}
