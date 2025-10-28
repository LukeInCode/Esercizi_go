package main
import "fmt"
import "math/rand"

func main() {
  const K = 10;
  var even int = 0;
  
  for i := 0; i < K; i++ {
    generated := rand.Intn(11); 
    fmt.Print(generated," ");

    if generated % 2 == 0 {
      even++;
    }
  }
  fmt.Println("\nNumeri pari generati: ", even);
}
