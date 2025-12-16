package main
import "fmt"

func recursiveSum(list []int) int {
  if len(list) == 0 {
    return 0;
  }
  return list[0] + recursiveSum(list[1:]);
}

func main() {
  fmt.Println(recursiveSum([]int{}));
}
