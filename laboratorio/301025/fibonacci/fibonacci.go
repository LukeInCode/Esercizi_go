package main
import "fmt"

func main() {
  var n, sum int = -1, 1;
  fmt.Print("inserire n: ");
  _,err := fmt.Scan(&n);
  
  if err != nil || n < 0 {
    fmt.Println("input errati!");
  }else {
    var previous, most_previous, index int;
    
    for i := 0; i < n; i++ {
      if i > 1 {
        sum = previous + most_previous;
      }
      if i == 2 || index == i - 1 {
        most_previous = previous;
        index = i;
      }
      for j := 0; j < sum; j++ {
        fmt.Print("*");
      }
      fmt.Println();
      previous = sum;
    } 
  }
}
