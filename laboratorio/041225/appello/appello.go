package main
import "fmt"
import "sort"
import "os"

func main() {
  people := make([]string, len(os.Args));
  for i,v := range os.Args {
    if i > 0 {
      people[i] = v;
    }
  }
  sort.Strings(people);
  fmt.Println(people);
}
