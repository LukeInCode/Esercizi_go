package main
import "fmt"
import "math/rand"

func main() {
  const K = 10;

  for i := 0; i < K; i++ {
    fmt.Print(rand.Intn(11)," ");
  }
  fmt.Println();
}
