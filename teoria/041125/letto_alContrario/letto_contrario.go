package main
import "fmt"
import "bufio"
import "os"

func main() {
  
  scanner := bufio.NewScanner(os.Stdin)
  var txt []string;
  
  for scanner.Scan() {
    row := scanner.Text()
    txt = append(txt,row)
  }
  for i := len(txt) - 1; i >= 0; i-- {
    fmt.Println(txt[i])
  }
}
