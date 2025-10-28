package main
import "fmt"
import "math/rand"

func main() {
  const K = 10;
  var sum int = 0;
  
  for i := 0; i < K; i++ {
    generated := rand.Intn(11); 
    fmt.Print(generated," ");

    sum += generated;
  }
  fmt.Println("\nSomma numeri generati: ", sum);
}
