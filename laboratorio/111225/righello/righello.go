package main
import "fmt"
import "os"
import "strconv"

func righello(n int) {
  if n == 0 {
    return;
  }
  if n == 1 {
    fmt.Print("-");
    return;
  }
  righello(n - 1);
  fmt.Println();
  for i := 0; i < n; i++ {
    righello(1);
  }
  fmt.Println();
  righello(n - 1);
}

func main() {
  if len(os.Args) < 2 {
    fmt.Println("fornire la dimensione (int>=0) su linea di comando");
    return;
  }
  lung, err := strconv.Atoi(os.Args[1]);
  if err != nil || lung < 0 {
    fmt.Println("argomento non valido");
    return;
  }
  righello(lung);
  fmt.Println();
}
