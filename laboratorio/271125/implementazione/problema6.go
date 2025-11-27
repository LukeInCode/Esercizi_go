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
    if strings.IndexAny(strings.ToLower(s), "123456789") != -1 {
      fmt.Println(s);
    }
  }
}
