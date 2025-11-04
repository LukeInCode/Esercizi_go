package main

import "fmt"

func main() {
  var n int
  fmt.Scan(&n)
  
  for y := 0; y < n; y++ {
    for x := 0; x < n; x++ {
      if y == 0 || y == n-1 || x == 0 || x == n-1 || x == y || x+y == n-1 {
        fmt.Print("*")
      } else {
        fmt.Print(" ")
      }
    }
    fmt.Println()
  }
}
