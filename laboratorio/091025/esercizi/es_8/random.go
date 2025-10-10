package main
import (
	"fmt"
	"math/rand"
)

func main() {
	var numero_random int = rand.Intn(101);
	fmt.Println("Numero generato randomicamente:", numero_random);
}