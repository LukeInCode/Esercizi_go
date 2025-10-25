package main
import "fmt"


func count_lowercase(original string) int {
  var counter int = 0;
  const START_LOWERCASE = 97; const END_LOWERCASE = 122;
  
  for _,r := range original {
    if r >= START_LOWERCASE && r <= END_LOWERCASE {
      counter++;
    }
  }
  return counter;
}

func main() {
  var s string;
  fmt.Print("Inserire una stringa: "); 
  _,err := fmt.Scan(&s);
  if(err != nil) {
    fmt.Println("Input non valido");
  }else {
    fmt.Println(count_lowercase(s));
  }
}
