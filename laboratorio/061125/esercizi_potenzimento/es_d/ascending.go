package main

import "fmt"

func main() {
	const K = 10
	var prev, curr int
	var ascending bool
	ascending = true

  for i := 1; i < K; i++ {
    fmt.Scan(&curr)
    if curr <= prev { 
			ascending = false
      break
    }
    prev = curr
  }
	if !ascending {
		fmt.Print("non ")
	}
	fmt.Println("crescente")
}

