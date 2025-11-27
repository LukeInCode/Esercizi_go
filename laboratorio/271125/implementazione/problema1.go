package main
import "fmt"

func main() {
  var s string;
  for {
    fmt.Scan(&s);
    if s == "stop" {
      break;
    }
    fmt.Println(s);
  }
}
