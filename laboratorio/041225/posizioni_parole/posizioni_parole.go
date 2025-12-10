package main
import "fmt"
import "bufio"
import "os"
import "strings"

type WordCounter map[string][]int;

func main() {
  counter := make(WordCounter);
  scanner := bufio.NewScanner(os.Stdin);

  for i := 0; scanner.Scan(); {
    words := strings.Fields(scanner.Text());
    for _,w := range words {
      counter[w] = append(counter[w], i);
      i++;
    }
  }
  fmt.Println(counter);
}
