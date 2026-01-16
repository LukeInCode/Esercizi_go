package main
import "fmt"
import "os"
import "bufio"

func main() {
  scanner := bufio.NewScanner(os.Stdin);
  var max int = -1;
  
  for i := 1; scanner.Scan(); i++ {
    row := scanner.Text();
    row_len := len(row);
    fmt.Print(i, ",", row_len, ":", row, "\n");

    if row_len > max {
      max = row_len;
    }
  }
  fmt.Println("lunghezza max", max);
}
