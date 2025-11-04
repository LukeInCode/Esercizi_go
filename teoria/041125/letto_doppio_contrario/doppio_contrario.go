package main
import "fmt"
import "bufio"
import "os"

func main() {
  
  scanner := bufio.NewScanner(os.Stdin);
  var txt []string;
  
  for scanner.Scan() {
    row := scanner.Text();
    txt = append(txt,row);
  }
  for i := len(txt) - 1; i >= 0; i-- {
    row := make([]string, len(txt[i]) - 1);
    for _,r := range txt[i] {
      row = append(row, string(r));
    }
    for j := len(row) - 1; j >= 0; j-- {
      fmt.Print(row[j]);
    }
    fmt.Println();
  }
}
