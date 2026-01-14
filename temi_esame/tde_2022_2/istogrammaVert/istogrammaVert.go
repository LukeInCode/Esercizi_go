package main
import "fmt"
import "os"
import "strconv"

func max(numeri []int) (max int) {
  max = numeri[0];
  for _,v := range numeri[1:] {
    if v > max {
      max = v;
    }
  }
  return;
}

func disegnaIstogramma(numeri []int) {
  n_max := max(numeri);
  for i := 0; i < n_max; i++ {
    for j := 0; j < len(numeri); j++ {
      if i >= n_max - numeri[j] {
        fmt.Print("*");
      }else {
        fmt.Print(" ");
      }
    }
    fmt.Println();
  }
}


func argsToNumeri(args []string) (numeri []int) {
  for _,v := range args {
    n,_ := strconv.Atoi(v);
    numeri = append(numeri, n);
  }
  return;
}

func main() {
  args := argsToNumeri(os.Args[1:]);
  disegnaIstogramma(args);
}
