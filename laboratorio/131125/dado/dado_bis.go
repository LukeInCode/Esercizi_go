package main

import (
    "fmt"
    "math/rand"
)

func main() {
    const NFACCE = 6
    var lancio, n int
    var frequenza [NFACCE + 1]int

    fmt.Print("quanti lanci? ")
    fmt.Scan(&n)

    for i := 0; i < n; i++ {
        lancio = rand.Intn(NFACCE) + 1
        frequenza[lancio]++;
    }
    for i := 1; i < len(frequenza); i++ {
        fmt.Print(i,": ", frequenza[i], "(", frequenza[i]*100/n, "%)\n")
    }
}
