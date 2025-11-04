package main
import "fmt"
import "bufio"
import "os"
import "strings"

func main() {
  scanner := bufio.NewScanner(os.Stdin);
  var counter int;
  var pattern string;

  if !scanner.Scan() { 
    os.Exit(1); 
  }else { 
    pattern = scanner.Text(); 
  }
  
  for scanner.Scan() {
    row := scanner.Text();

    if strings.Contains(row, pattern) {
      counter++;
    }
  }
  fmt.Printf("Il pattern %s è contenuto in %d righe\n", pattern, counter);
}
