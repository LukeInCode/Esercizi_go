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
  
  sort(&words);
  for i := 0; i < len(words); i++ {
      fmt.Println(words[i]);
  }
}

func min(v1, v2 int) int {
  if v1 <= v2 {
    return v1;
  }
  return v2;
}


func sort(toOrder *[]string) {
  var arr []string = make([]string, len(*toOrder));
  copy(arr, *toOrder)
  for i := 0; i < len(*toOrder) - 1; i++ {
    if (*toOrder)[i] > (*toOrder)[i + 1] {
      (arr)[i], (arr)[i + 1] = (*toOrder)[i + 1], (*toOrder)[i];
    }
  }
  *toOrder = arr;
}
