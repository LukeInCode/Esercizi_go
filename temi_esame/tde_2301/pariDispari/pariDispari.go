package main
import "fmt"
import "os"
import "strconv"

func isEven(n int) int {
  if n % 2 == 0 {
    return 0;
  }
  return 1;
}

func main() {
  var previous_status int = -1;
  var error_seq bool = false;
  var index_error int = 0;

  if len(os.Args) <= 1 {
    fmt.Println("nessun valore in ingresso");
    os.Exit(0);
  }
  for i, el := range os.Args[1:] {
    n,err := strconv.Atoi(el);
    if err != nil {
      error_seq = true;
      index_error = i;
      break;
    }
    actual_status := isEven(n);
    if previous_status != -1 {
      if actual_status == previous_status {
        error_seq = true;
        index_error = i;
        break;
      }
    }
    previous_status = actual_status;
  }

  if !error_seq {
    fmt.Println("sequenza valida");
  }else {
    fmt.Printf("elemento in posizione %d non valido\n", index_error + 1);
  }
}

