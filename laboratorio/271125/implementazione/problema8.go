package main
import "fmt"
import "io"

func media(arr []string) int {
  var sum int;
  for _,v := range arr {
    sum += len(v);
  }
  return sum / len(arr);
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

  vMedia := media(words);
  for i := 0; i < len(words); i++ {
    if len(words[i]) > vMedia {
      fmt.Println(words[i]);
    }
  }
}
