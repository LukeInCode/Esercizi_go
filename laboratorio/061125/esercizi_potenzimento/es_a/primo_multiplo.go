package main
import "fmt"

func main() {
  var n int;
  var mult_3 int = -1;
  var i = 0;
  
  for i < 5 {
    fmt.Scan(&n);
    if n < 0 {
      continue;
    }
    if n % 3 == 0 {
        mult_3 = n;
        break;
    }
    i++;
  }
  if mult_3 != -1 {
    fmt.Println(mult_3);
  }else {
    fmt.Println("no");
  }
}
