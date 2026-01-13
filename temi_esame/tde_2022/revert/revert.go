package main
import "bufio"
import "fmt"
import "os"
import "strings"

func main() {
  scanner := bufio.NewScanner(os.Stdin);

  for scanner.Scan() {
    fields := strings.Fields(scanner.Text());
    for _,v := range fields {
      if v == "stop" {
        break;
      }
      if len(v) % 2 == 0 {
        for i := len(v) - 1; i >= 0; i-- {
          fmt.Print(string(v[i]));
        }
        fmt.Println();
      }
    }
  }
}
