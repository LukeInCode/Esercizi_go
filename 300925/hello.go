package main
import "fmt"

func main () {
  var n int
  fmt.Print("Quante volte vuoi esser saluto padrone? ")
  fmt.Scan(&n)

  for i := 0; i < n; i++ {
    fmt.Println("Ciao ", i + 1)
  }
}

