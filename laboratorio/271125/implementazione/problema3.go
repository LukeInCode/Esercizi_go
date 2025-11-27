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
  for i := 0; i < len(words); i++ {
    if len(words[i]) % 2 == 0 {
      fmt.Println(words[i]);
    }
  }
}
