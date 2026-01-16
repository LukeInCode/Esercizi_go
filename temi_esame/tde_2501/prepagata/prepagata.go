package main
import "fmt"
import "errors"
import "os"

type Prepagata struct {
  numero int;
  saldo float64;
}

func ricarica(carta *Prepagata, importo int) {
  if importo > 0 {
    carta.saldo += float64(importo);
  }
}

func preleva(carta *Prepagata, importo int) (int, error) {
  if importo >= 0 {
    if carta.saldo >= float64(importo) {
      carta.saldo -= float64(importo);
      return importo, nil;
    } else {
      return 0, errors.New("saldo insufficiente");
    }
  } else {
    return 0, errors.New("importo non valido");
  }
}

func (c Prepagata) String() string {
  return fmt.Sprintf("carta n. %d, saldo %.2f", c.numero, c.saldo);
}


func main() {
  var prepagata Prepagata = Prepagata{ numero: 1709, saldo: 100 };
  var choice string;

  for {
    fmt.Println("a. saldo\nb. ricarica\nc. prelievo\ne. esci");
    fmt.Scan(&choice);

    switch choice {
      case "a":
        fmt.Println(prepagata);
      case "b":
        var versamento int = 0;
        fmt.Scan(&versamento);
        ricarica(&prepagata, versamento);
        fmt.Println(prepagata);
      case "c":
        var versamento int = 0;
        fmt.Scan(&versamento);
        _,err := preleva(&prepagata, versamento);
        if err == nil {
          fmt.Println(prepagata);
        } else {
          fmt.Println(err);
        }
      case "e":
        fmt.Println("arrivederci");
        os.Exit(0);
      default:
        fmt.Println("opzione non valida");
    }
  }
}
