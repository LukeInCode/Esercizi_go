package main
import "fmt"
import "io"

func main() {
  var p, day int = -1, 1;

  for c := 1 ;; c++{
     _, err := fmt.Scan(&p);
    if err == io.EOF {
      break
    }
    if p != 0 {
      day = c;
    }
  }
  fmt.Println("L'ultiom giorno che ha piovuto è:",day);
}
