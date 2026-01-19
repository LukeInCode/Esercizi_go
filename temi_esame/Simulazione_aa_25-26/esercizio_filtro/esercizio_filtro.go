package main
import "fmt"
import "os"
import "strconv"
import "slices"

func extractDigits(n int) []int {
  var digits []int;
  for n > 0 {
    digits = append(digits, n % 10);
    n /= 10;
  }
  slices.Reverse(digits);
  return digits;
}

func main() {
  n,_ := strconv.Atoi(os.Args[1]);
  dgts := extractDigits(n);

  for i,d := range dgts {
    if i != len(dgts) - 1 {
      if dgts[i + 1] > d {
        fmt.Print(d);
      }
    }
  }
  fmt.Println();
}
