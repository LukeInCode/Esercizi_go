package main
import "fmt"
import "math/rand"

func main() {
  var max int = 0;

  for i := 0; i < 10; i++ {
    generated := rand.Intn(21) + 10;
    fmt.Print(generated, " ");

    if generated > max {
      max = generated;
    }
  }
  fmt.Println("\nMassimo valore generato:", max);
}
