package main
import "fmt"


func trovaDivisori(n int) (divisori []int) {
  divisori = append(divisori, 1);
  for i := 2; i < n; i++ {
    if n % i == 0 {
      divisori = append(divisori, i);
    }
  }
  return;
}

func Abbondante(n int) bool {
  if n <= 0 {
    return false;
  }
  var sum int = 0;
  var abbondante bool = false;
  for _,v := range trovaDivisori(n) {
    sum += v;
    if sum > n {
      abbondante = true;
      break;
    }
  }
  return abbondante;
}


func main() {
  var n, counter int;
  fmt.Scanf("%d", &n);
  for i := 0; counter < n; i++ {
    if Abbondante(i) {
      fmt.Println(i);
      counter++;
    }
  }
}
