package main
import "fmt"


func count_occurrences(s string, target rune) int {
  var counter int = 0;
  
  for _,r := range s {
    if r == target {
      counter++;
    }
  }
  return counter;
}

func main() {
  var s string;
  var target rune = '/';
  
  fmt.Print("Inserire una stringa: "); 
  _,err := fmt.Scan(&s);
  if(err != nil) {
    fmt.Println("Input non valido");
  }else {
    fmt.Println(count_occurrences(s, target));
  }
}
