package main
import "fmt"
import "math/rand"

func main() {
  var n int;
  fmt.Print("Inserisci il risultato della somma: "); fmt.Scan(&n);

  var sum,generated int = 0,0;

  for {
    generated = rand.Intn(10) + 1; 
    fmt.Print(generated," ");

    sum += generated;
    
    if sum >= n {
      break;
    }
  }
  fmt.Println("\nSomma: ", sum);
}
