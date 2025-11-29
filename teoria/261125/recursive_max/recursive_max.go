package main
import "fmt"

func recurive_max(list []int, index int) (int, int) {
  if len(list) == 1 {
    return list[0], index;
  } else {
    m, index_max := recurive_max(list[1:], index + 1);
    if list[0] < m {
      return m, index_max;
    }
    return list[0], index;
  }
}


func main() {
  el, index := recurive_max([]int{1,2,5,1}, 0);
  fmt.Println(el, index);
}
