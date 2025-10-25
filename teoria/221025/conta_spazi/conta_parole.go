package main
import "fmt"


func count_words(s string) int {
  var counter int = 0;
  const TARGET = ' ';
  var spaces bool = true;
  
  for _,r := range s {
    if r == TARGET {
      spaces = true;
      continue;
    }
    if spaces {
      counter ++;
      spaces = false;
    }
  }
  return counter;
}

func main() {
  var s string = "Ciao    COME Va    CIAO          ciao";  
  fmt.Println(count_words(s));
}
