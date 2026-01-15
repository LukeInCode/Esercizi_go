package main
import "fmt"
import "os"

func contaMaxTot(s string) ([10]int, int, int) {
  var multi_counter [10]int;
  var max, sum int = -1, 0;

  for _,v := range s {
    if v >= '0' && v <= '9' {
      var n int = int(v - '0');
      multi_counter[n]++;
      sum += n;
      if n > max {
        max = n;
      }
    }
  }
  return multi_counter, max, sum;
}

func main() {
  if len(os.Args) <= 1 {
    fmt.Println("manca argomento");
  } else {
    fmt.Println(contaMaxTot(os.Args[1]));
  }
}


