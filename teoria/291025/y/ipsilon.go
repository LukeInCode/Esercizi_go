package main

import "fmt"

/* Stampa una riga di n asterischi e va a capo */
func stampaRigaPiena(n int) {
  for i := 0; i < n; i++ {
    fmt.Print("*")
  }
  fmt.Println()
}

func stampaRigaCentrale(n int) {
  fmt.Print("*")
  for i := 0; i < (n-3)/2; i++ {
    fmt.Print(" ")
  }
  fmt.Print("*")
  for i := 0; i < (n-3)/2; i++ {
    fmt.Print(" ")
  }
  fmt.Println("*")
}

func stampaRigaIntermedia(n, r int) {
    fmt.Print("*")
    for i := 0; i < r; i++ {
      fmt.Print(" ")
    }
    fmt.Print("*")
    for i := 0; i < n-4-2*r; i++ {
      fmt.Print(" ")
    }
    fmt.Print("*")
    for i := 0; i < r; i++ {
      fmt.Print(" ")
    }
    fmt.Println("*")
}

func main() {
  var n int
  fmt.Scan(&n)

  // Riga a (iniziale)
  stampaRigaPiena(n)

  // Righe del gruppo A
  for r := 0; r < (n-3)/2; r++ {
    stampaRigaIntermedia(n, r)
  }
  
  // Riga b (centrale)
  stampaRigaCentrale(n)

  // Righe del gruppo B
  for r := (n-3)/2 - 1; r >= 0; r-- {
    stampaRigaIntermedia(n, r)
  }

  // Riga c (finale)
  stampaRigaPiena(n)
}

