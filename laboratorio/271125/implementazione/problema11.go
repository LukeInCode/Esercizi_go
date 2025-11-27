package main
import "fmt"
import "io"

func reverse (s string) string {
  var reverse []rune;
  for i := len([]byte(s)) - 1; i >= 0; i-- {
    reverse = append(reverse, rune(s[i]));
  }
  return string(reverse);
}


func main() {
 var words []string;
 var s string;
 
  for {
    _, err := fmt.Scan(&s);
    if err == io.EOF {
      break;
    }
    words = append(words, s);
  }
  
  for i := 0; i < len(words); i++ {
    fmt.Println(reverse(words[i]));
  }
}
