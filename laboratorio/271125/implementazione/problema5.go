package main
import "fmt"
import "strings"

func main() {
  var s string;
 
  for {
    fmt.Scan(&s);
    if s == "stop" {
      break;
    }
    if strings.IndexAny(strings.ToLower(s), "aeoiu") == 0 {
      fmt.Println(s);
    }
  }
}
