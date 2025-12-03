package main
import "fmt"
import "os"
import "strconv"
import "math"

func isSquare(n int) bool {
  sqrt := math.Sqrt(float64(n));

  if int(sqrt) == int(sqrt + 0.9) {
    return true;
  }
  return false;
}


func main() {
  var square []int;  

  for i := 0; i < len(os.Args); i++ {
    value, err := strconv.Atoi(os.Args[i]);
    if err == nil {
      if isSquare(value) {
        square = append(square, value);
      }
    }
  }

  for _,v := range square {
    fmt.Println(v);
  }
}



