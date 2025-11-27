package main
import "fmt"
import "io"

func findMaxLenght(arr []string) int {
  var max int = len(arr[0]);
  for i := 1; i < len(arr); i++ {
    if len(arr[i]) > max {
      max = len(arr[i]);
    }
  }
  return max;
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
  
  maxLenght := findMaxLenght(words);
  for i := 0; i < len(words); i++ {
    if len(words[i]) == maxLenght {
      fmt.Println(words[i]);
      break;
    }
  }
}
