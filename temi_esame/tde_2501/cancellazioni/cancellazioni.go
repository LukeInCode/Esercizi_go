package main
import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

func cancella(lista []string) []string {
  var lista_finale []string = make([]string, len(lista))
  copy(lista_finale, lista);
  
  for i,v := range lista_finale {
    var lastIndex int;
    n,err := strconv.Atoi(v);
    if err == nil {
      lastIndex := i + n + 1;
      if lastIndex >= len(lista_finale) {
        lastIndex = -1;
      }
      if lastIndex != -1 {
        lista_finale = append(lista_finale[:i], lista_finale[lastIndex:]...);
      } else {
        lista_finale = lista_finale[:i];
      }
    }
    if lastIndex == -1 {
      break;
    }
  }
  return lista_finale;
}

func main() {
  if len(os.Args) <= 1 {
    fmt.Println("Fornire 1 nome di file");
  } else {
    file,err_file := os.Open(os.Args[1]);
    if err_file != nil {
      fmt.Println("File non accessibile");
    } else {
      var lista_originale, lista_cancellata []string;
      scanner := bufio.NewScanner(file);
      for scanner.Scan() {
        row := strings.Fields(scanner.Text());
        lista_originale = append(lista_originale, row...);
        lista_cancellata = cancella(row);
      }
      fmt.Println(lista_originale);
      fmt.Println(lista_cancellata);
    }
  }
}
