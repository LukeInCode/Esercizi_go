package main
import "fmt"
import "os"
import "strconv"

func fattoriale(n int) int {
  if n == 0 {
    return 1;
  }
  return n * fattoriale(n - 1);
} 

func main() {
  n,_ := strconv.Atoi(os.Args[1]);
  fmt.Println(fattoriale(n));
}
