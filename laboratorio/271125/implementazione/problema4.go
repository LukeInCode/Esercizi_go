package main
import "fmt"

func main() {
 var odd, even []string;
 var s string;
 
  for {
    fmt.Scan(&s);
    if s == "stop" {
      break;
    }
    casted := []byte(s);
    if len(casted) % 2 == 0 {
      even = append(even, s);
    }else {
      odd = append(odd,s);
    }
  }
  for i := 0; i < len(odd) - 1; i++ {
    fmt.Print(odd[i] + " ");
  }
  fmt.Println(odd[len(odd) - 1], "\n");


  for i := 0; i < len(even) - 1; i++ {
    fmt.Print(even[i] + " ");
  }
  fmt.Println(even[len(even) - 1], "\n");
}
