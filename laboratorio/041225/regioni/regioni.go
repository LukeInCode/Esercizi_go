package main
import "fmt"
import "bufio"
import "os"
import "strings"


func main() {
  mp := make(map[string][]string);

  scanner := bufio.NewScanner(os.Stdin);
  if scanner.Scan() {
    scanner.Text();
  }
  for scanner.Scan() {
    fields := strings.Split(scanner.Text(), ",");
    var regione string = fields[2];
    mp[regione] = append(mp[regione], fields[0]);
  }

  if len(os.Args) > 1 {
    for i := 1; i < len(os.Args); i++ {
      fmt.Println(os.Args[i], mp[os.Args[i]]);
    }
  }
}
