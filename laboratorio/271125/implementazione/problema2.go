package main
import "fmt"

func main() {
 var words []string;
 var s string;
 
  for {
    fmt.Scan(&s);
    if s == "stop" {
      break;
    }
    words = append(words, s);
  }
  for i := len(words) - 1; i >= 0; i-- {
    fmt.Println(words[i]);
  }
}
