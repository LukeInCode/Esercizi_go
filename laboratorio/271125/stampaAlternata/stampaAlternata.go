package main
import "fmt"
import "bufio"
import "os"

func main() {
  var rows []string;
  scanner := bufio.NewScanner(os.Stdin);
  for i := 1; scanner.Scan(); i++ {
    row := scanner.Text();
    if i % 2 == 0 {
      fmt.Println(row);
    }else {
      rows = append(rows, row);
    }
  }

  for _,v := range rows {
    fmt.Println(v);
  }
}
